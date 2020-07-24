package parser

import (
	"errors"
	"strconv"
	"strings"
)

// DecodeGreyScale decodes message hidden steganographically on a grey scale image
// by getting the last bit from every byte
func DecodeGreyScale(bytes []byte, bitLenPerChar int) (string, error) {
	//return error if length do not match
	if len(bytes)%bitLenPerChar != 0 {
		return "", errors.New("invalid-input-length")
	}
	decoded := ""
	currPos := 0
	var parsedBit byte = 0
	for _, part := range bytes {
		currPos++
		//append the last bit from a byte
		//by shifting the previous parsed byte
		//then add the last bit of current byte
		parsedBit = (parsedBit << 1) + (part % byte(2))
		if currPos == bitLenPerChar {
			//convert the parsedBit to string then append it to decoded message
			//then reset the parsedBit and position tracker
			decoded += string(parsedBit)
			currPos, parsedBit = 0, 0
		}
	}
	return decoded, nil
}

//bytesStringToBytesSlice helper func to convert BytesString to BytesSlice
func bytesStringToBytesSlice(bytesString, separator string) ([]byte, error) {
	bytes := []byte{}
	sliceOfByteStrings := strings.Split(bytesString, separator)
	for _, byteString := range sliceOfByteStrings {
		//parse to Uint64 then cast to byte
		parsedUint64, err := strconv.ParseUint(byteString, 10, 7)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, byte(parsedUint64))
	}
	return bytes, nil
}

//Parse7BitStringGreyScale is Jumat Hek solution function
//This func will parse steganographed byte string separated by spaces into messages
func Parse7BitStringGreyScale(bytesStrings string) (string, error) {
	//Parse Byte string into Byte slices to make it easier to decode the message
	bytes, err := bytesStringToBytesSlice(bytesStrings, " ")
	if err != nil {
		return "", err
	}
	//Decode the message with 7 bit / character
	decoded, err := DecodeGreyScale(bytes, 7)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
