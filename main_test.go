package main

import "testing"

func TestSbox(t *testing.T) {
	originalValue := "long string to check if the encryption algorithm is working correctly"

	encryptedValue := sboxencrypt(originalValue)

	decryptedValue := sboxdecrypt(encryptedValue)

	if originalValue != decryptedValue {
		t.Errorf("Test failed originalValue: %s, decryptValue: %s", originalValue, decryptedValue)
	}
}

func TestPbox(t *testing.T) {
	originalValue := "qwertyui"

	encryptedValue := pboxencrypt(originalValue)

	decryptedValue := pboxdecrypt(encryptedValue)

	if originalValue != decryptedValue {
		t.Errorf("Test failed originalValue: %s, decryptValue: %s", originalValue, decryptedValue)
	}
}
