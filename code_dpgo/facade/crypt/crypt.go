package crypt

type Xor struct {
}

func (s *Xor) Encrypt(key byte, data *[]byte) *[]byte {
	// create empty array to store bytes
	arr := []byte{}

	for _, item := range *data {
		// XOR operation to encrypt
		arr = append(arr, item^key)
	}

	return &arr
}

func (s *Xor) Decrypt(key byte, data *[]byte) *[]byte {
	// XOR again on the encrypted data to
	// give us back the original data
	return s.Encrypt(key, data)
}
