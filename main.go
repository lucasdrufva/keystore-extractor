package main

import (
	"fmt"
	"os"
)

func main() {
	keystoreFile := os.Args[1]
	keystore, _ := NewFileKeystoreWithPasswordAndStrictPerms(keystoreFile, NewSecureString([]byte("")), false)
	l, _ := keystore.(ListingKeystore)
	keylist, _ := l.List();
	
	for _, key := range keylist {
		value_secure_string, _ := keystore.Retrieve(key)
		value, _ := value_secure_string.Get()
		fmt.Print(key)
		fmt.Print(" : ")
		fmt.Println(string(value[:]))
	}
}