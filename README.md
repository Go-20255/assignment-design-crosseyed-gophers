# Crosseyed Gophers -- Cipher Hunt

Authors: Kyle Zaffram, Kirk Burrill

## INSTRUCTOR README

This is an assignment in which students will implement basic ciphers for encrypting and decrypting text.  Students will also learn about *regustry design patterns*, as they will register all of the ciphers they have implemented in the main program. 

The purpose of a registry is to make a system modular and easily expandable, without making changes in multiple places.  In this case, sutdents implement various Ciphers, and register them in a central registry. 

This registry is used by a simple CLI, which dynamically shows what ciphers have been implemented.  In theory, one could add a cipher by implementing a Cipher, and registering it in a single place.  One could also disable a Cipher by removing it from the registry. 

Ciphers themselves own their names, parameters, and encrypt/decrypt implementations.  This enforces a Single State of Truth, so that parameter names and cipher names do not have to be hardcoded in multiple places. 


### How does this assignment work?

Students are given stubs for various ciphers, and an incomplete registry.  Students are also given a `start.json` file, which contains an encrypted message.  This message can only be fully decoded after every cipher has been implemented. 

The initial `start.json` contains the first cipher name and parameter, as well as the ciper text.  The cipher text, once decoded, contains the cipher name, parameters, and ciphertext for the next step.  This is like Russian nesting dolls -- until the last decryption reveals the final message. 


### INSTRUCTOR: Creating the initial clue

To create the start.json, use `tools/puzzle-gen/main.go`.  This will use the registry to discover all implemented ciphers, and will create a Russian nesting doll in which every cipher is used once. 

For a sanity check, use the built-in solver: `tools/puzzle-solve/main.go`.  This will dynamically perform all the steps needed to decrypt the final message, and make sure everything is working. 

*Important*  Students are not given tools/, so they can't just solve the puzzle or generate a different one.


### INSTRUCTOR: Grading

If all tests pass on student code, it means all ciphers have been properly implemented. 

Additionally, check the registry to make sure all ciphers have been registered.  You can also give different messages to different students, and automate generating `start.json`. 


### Additional Instructor Note:

Due to this being a basic CLI, we did not include a `Makefile` or run script.

Students can complete this assignment using the CLI `go run main.go`, and the provided test suite.

The `dist/` directory contains a student-ready version of this assignment with stubbed implementations.