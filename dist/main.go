package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"project/cipher"
	"project/registry"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		printUsage()
		return nil
	}

	switch args[0] {
	case "list":
		fmt.Println(strings.Join(registry.CipherRegistry.Names(), "\n"))
		return nil
	case "encrypt", "decrypt":
		return runOp(args[0], args[1:])
	default:
		printUsage()
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runOp(mode string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s requires <cipher>", mode)
	}
	cipherName := args[0]

	c, err := registry.CipherRegistry.Get(cipherName)
	if err != nil {
		return err
	}

	params, inputB64, outputB64, err := parseFlags(args[1:], c)
	if err != nil {
		return err
	}

	parsed, err := c.ParseParams(params)
	if err != nil {
		return err
	}

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	if inputB64 {
		decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(in)))
		if err != nil {
			return fmt.Errorf("base64 decode stdin: %w", err)
		}
		in = decoded
	}

	var out []byte
	if mode == "encrypt" {
		out, err = c.Encrypt(in, parsed)
	} else {
		out, err = c.Decrypt(in, parsed)
	}
	if err != nil {
		return err
	}

	if outputB64 {
		enc := base64.StdEncoding.EncodeToString(out)
		_, err = os.Stdout.WriteString(enc + "\n")
		return err
	}
	_, err = os.Stdout.Write(out)
	return err
}

func parseFlags(args []string, c cipher.Cipher) (params map[string]string, inputB64 bool, outputB64 bool, err error) {
	params = map[string]string{}

	for i := 0; i < len(args); i++ {
		a := args[i]
		switch a {
		case "--input-base64":
			inputB64 = true
			continue
		case "--output-base64":
			outputB64 = true
			continue
		}

		if !strings.HasPrefix(a, "--") {
			return nil, false, false, fmt.Errorf("unexpected argument %q (flags must be --name value)", a)
		}

		nameVal := strings.TrimPrefix(a, "--")
		var name, val string
		if eq := strings.IndexByte(nameVal, '='); eq >= 0 {
			name = nameVal[:eq]
			val = nameVal[eq+1:]
		} else {
			if i+1 >= len(args) {
				return nil, false, false, fmt.Errorf("flag --%s requires a value", nameVal)
			}
			name = nameVal
			val = args[i+1]
			i++
		}

		params[name] = val
	}

	// Enforce required params using cipher-owned schema.
	for _, ps := range c.ParamSpecs() {
		if ps.Required {
			if _, ok := params[ps.Name]; !ok {
				return nil, false, false, fmt.Errorf("missing required --%s for cipher %q", ps.Name, c.Name())
			}
		}
	}

	return params, inputB64, outputB64, nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  crosseyed-gophers list\n")
	fmt.Fprintf(os.Stderr, "  crosseyed-gophers encrypt <cipher> [--param value ...] [--input-base64] [--output-base64]\n")
	fmt.Fprintf(os.Stderr, "  crosseyed-gophers decrypt <cipher> [--param value ...] [--input-base64] [--output-base64]\n\n")

	fmt.Fprintf(os.Stderr, "Registered ciphers (only ciphers that have been registered appear here):\n")
	names := registry.CipherRegistry.Names()
	if len(names) == 0 {
		fmt.Fprintf(os.Stderr, "  (none)\n\n")
		return
	}
	for _, name := range names {
		c, err := registry.CipherRegistry.Get(name)
		if err != nil {
			continue
		}
		fmt.Fprintf(os.Stderr, "  %s\n", c.Name())
		for _, ps := range c.ParamSpecs() {
			req := "optional"
			if ps.Required {
				req = "required"
			}
			fmt.Fprintf(os.Stderr, "    --%s (%s, %s) %s\n", ps.Name, ps.Type, req, ps.Help)
		}
	}
	fmt.Fprintln(os.Stderr)

	fmt.Fprintf(os.Stderr, "Notes:\n")
	fmt.Fprintf(os.Stderr, "  - Cipher parameters are owned by the cipher implementation (ParamSpecs/ParseParams).\n")
	fmt.Fprintf(os.Stderr, "  - The registry only stores the list of registered ciphers.\n")
	fmt.Fprintf(os.Stderr, "  - Use --output-base64 when ciphertext may contain non-printable bytes (e.g. XOR).\n")
}
