/*
SiGG-DIDs-Infrastructure
*/

package encoder

import "encoding/base64"

// EncodeToString encodes the bytes to string.
func EncodeToString(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// DecodeString decodes the encoded content to Bytes.
func DecodeString(encodedContent string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(encodedContent)
}
