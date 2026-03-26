package cipher

import (
	"fmt"
	"sort"
	"strings"
)

type Registry struct {
	byName map[string]Cipher
}

func NewRegistry(ciphers ...Cipher) *Registry {
	r := &Registry{byName: map[string]Cipher{}}
	for _, c := range ciphers {
		if c == nil {
			panic("cipher.NewRegistry: nil cipher")
		}
		n := strings.ToLower(strings.TrimSpace(c.Name()))
		if n == "" {
			panic("cipher.NewRegistry: empty cipher name")
		}
		if _, exists := r.byName[n]; exists {
			panic("cipher.NewRegistry: duplicate cipher " + n)
		}
		r.byName[n] = c
	}
	return r
}

func (r *Registry) Names() []string {
	out := make([]string, 0, len(r.byName))
	for k := range r.byName {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (r *Registry) Get(name string) (Cipher, error) {
	n := strings.ToLower(strings.TrimSpace(name))
	c, ok := r.byName[n]
	if !ok {
		return nil, fmt.Errorf("unknown cipher %q (registered: %s)", name, strings.Join(r.Names(), ", "))
	}
	return c, nil
}
