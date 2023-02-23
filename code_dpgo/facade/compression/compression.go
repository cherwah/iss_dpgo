package compression

type Rle struct {
}

// implements run-length encoding for a stream of 0s and 1s
func (r *Rle) Compress(plain *[]byte) *[]byte {
	// create empty array to store bytes
	zip := []byte{}

	// init to first character of 'plain'
	var last byte

	if len(*plain) > 0 {
		last = (*plain)[0]
	}

	rle_ct := 0

	for _, curr := range *plain {
		if curr != last {
			zip = append(zip, byte(last))
			zip = append(zip, byte(rle_ct))

			// switch "fence-marker"
			last = curr

			// reset counter
			rle_ct = 1
		} else {
			// increment run-length for current character
			rle_ct++
		}
	}

	// record final block
	zip = append(zip, byte(last))
	zip = append(zip, byte(rle_ct))

	return &zip
}

func (r *Rle) Uncompress(zip *[]byte) *[]byte {
	// create empty array to store bytes
	unzip := []byte{}
	var curr byte

	for i, item := range *zip {
		// the character and run-length are stored
		// in alternate fashion within 'zip'
		if i%2 == 0 {
			curr = item
			unzip = append(unzip, curr)
		} else {
			rlen := item - 1
			for rlen > 0 {
				// append the required run-length
				// of the current character
				unzip = append(unzip, curr)
				rlen--
			}
		}
	}

	return &unzip
}
