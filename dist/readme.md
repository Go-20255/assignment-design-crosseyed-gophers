# Crosseyed Gophers -- Cipher Hunt

Hello studnet. 

There is a secret message in `start.json`, and nobody at the NSA has been able to figure it out. 

Luckily, we confiscated this project from the message sender.  This seems to be some sort of encryption program they were working on.

However, this version is incomplete.  It looks like there are several ciphers that need to be implemented before the message can be read.

## INSTRUCTIONS

Use the specifications in each cipher implementation to complete the cipher.  Register them in `registry/cipher_registry.go` so that they can be used in the CLI, and try your best to decode the final message. 

Whoever was sending this message was not messing around -- it looks like we will need every cipher to decode the final message.

### CLI

The CLI is fully intact -- no need to change this.  Run `go run main.go` to see how to use this.

The CLI works entirely in stdin/stdout, so you can paste things in, or pipe input/output with `<` or `>`.

Never heard of stdin/out? Are you on Windows??  You shouldn't be. 

## Testing ciphers

Cipher unit tests are provided in `cipher/*_test.go`. They will **fail** until you implement each cipher.

```bash
go test ./...
``` 