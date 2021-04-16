package hash_utils

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func GetBytes(key interface{}) ([]byte, error) {
	buf, ok := key.([]byte)
	if !ok {
		fmt.Println("err")
		return nil, fmt.Errorf("ooops, did not work")
	}

	return buf, nil
}

func CryptoHash(objs ...interface{}) string {
	digester := crypto.SHA256.New()
	for _, ob := range objs {
		fmt.Fprint(digester, ob)
	}
	sha := hex.EncodeToString(digester.Sum(nil))

	return sha
}

func HexToBinary(hex string) string {
	var buf bytes.Buffer

	for i := 0; i < len(hex); i++ {
		switch hex[i] {
		case '0':
			buf.WriteString("0000")
			break
		case '1':
			buf.WriteString("0001")
			break
		case '2':
			buf.WriteString("0010")
			break
		case '3':
			buf.WriteString("0011")
			break
		case '4':
			buf.WriteString("0100")
			break
		case '5':
			buf.WriteString("0101")
			break
		case '6':
			buf.WriteString("0110")
			break
		case '7':
			buf.WriteString("0111")
			break
		case '8':
			buf.WriteString("1000")
			break
		case '9':
			buf.WriteString("1001")
			break
		case 'a':
		case 'A':
			buf.WriteString("1010")
			break
		case 'b':
		case 'B':
			buf.WriteString("1011")
			break
		case 'c':
		case 'C':
			buf.WriteString("1100")
			break
		case 'd':
		case 'D':
			buf.WriteString("1101")
			break
		case 'e':
		case 'E':
			buf.WriteString("1110")
			break
		case 'f':
		case 'F':
			buf.WriteString("1111")
			break
		}
	}

	return buf.String()
}
