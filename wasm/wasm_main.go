package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"math/rand"
	"net/url"
	"strings"
	"syscall/js"
	"time"
)

const minLength float32 = 60
const hourLength = minLength * 60
const dayLength = hourLength * 24
const monthLength = dayLength * 31
const yearLength = dayLength * 365

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmGeneratePassword", js.FuncOf(generatePassword))
	global.Set("wasmHashCalculation", js.FuncOf(hashCalculation))
	global.Set("wasmHumanReadableTimediff", js.FuncOf(humanReadableTimediff))
	global.Set("wasmUnixTimeConverter", js.FuncOf(unixTimeConverter))
	global.Set("wasmEncodeDecode", js.FuncOf(encodeDecode))
	<-done
}

// Check whether a slice (string) contains specific element (string)
func sliceContains(slice []string, elementToCheck string) bool {
	for _, e := range slice {
		if e == elementToCheck {
			return true
		}
	}
	return false
}

// select N elements from a pool ([]string) and push them into a channel
func subGenerator(length int, pool []string, c chan string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		c <- pool[rand.Intn(len(pool))]
	}
}

func generatePassword(this js.Value, args []js.Value) interface{} {
	if len(args) != 5 {
		return "ERROR: number of arguments doesn't match"
	}

	var numberLen int
	var lowerLetterLen int
	var upperLetterLen int
	var symbolLen int
	var mustElements string

	numberLen = args[0].Int()
	lowerLetterLen = args[1].Int()
	upperLetterLen = args[2].Int()
	symbolLen = args[3].Int()
	mustElements = args[4].String()

	var lettersUpper []string
	var lettersLower []string

	for i := 0; i < 26; i++ {
		lettersUpper = append(lettersUpper, string(rune('A'+i)))
		lettersLower = append(lettersLower, string(rune('a'+i)))
	}

	numbers := strings.Split("0123456789", "")
	symbols := strings.Split("~!@#$%^&*()_-+=?<>,.:;", "")

	var result []string

	// Check if the "must have" element is available in the pool
	conflictErrorMsg := "ERROR: your compulsory element request conflicts with length request"
	for _, e := range strings.Split(mustElements, "") {
		if sliceContains(lettersUpper[:], e) {
			if upperLetterLen < 1 {
				return conflictErrorMsg
			}
			result = append(result, e)
			upperLetterLen = upperLetterLen - 1
		} else if sliceContains(lettersLower, e) {
			if lowerLetterLen < 1 {
				return conflictErrorMsg
			}
			result = append(result, e)
			lowerLetterLen = lowerLetterLen - 1
		} else if sliceContains(numbers, e) {
			if numberLen < 1 {
				return conflictErrorMsg
			}
			result = append(result, e)
			numberLen = numberLen - 1
		} else if sliceContains(symbols, e) {
			if symbolLen < 1 {
				return conflictErrorMsg
			}
			result = append(result, e)
			symbolLen = symbolLen - 1
		} else {
			return fmt.Sprintf("ERROR: element '%s' is not available", e)
		}
	}

	c := make(chan string)
	go subGenerator(numberLen, numbers, c)
	go subGenerator(lowerLetterLen, lettersLower, c)
	go subGenerator(upperLetterLen, lettersUpper, c)
	go subGenerator(symbolLen, symbols, c)

	for i := 0; i < (numberLen + lowerLetterLen + upperLetterLen + symbolLen); i++ {
		result = append(result, <-c)
	}

	// Shuffle the order of elements in the result
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return strings.Join(result, "")
}

func hashCalculation(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "ERROR: number of arguments doesn't match"
	}

	var plainText string
	var algo string

	plainText = args[0].String()
	algo = args[1].String()

	var h hash.Hash

	switch algo {
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	case "md5":
		h = md5.New()
	default:
		return fmt.Sprintf("ERROR: algo %s not supported", algo)
	}

	h.Write([]byte(plainText))
	hashed := h.Sum([]byte{})

	return hex.EncodeToString(hashed)
}

func humanReadableTimediff(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "ERROR: number of arguments doesn't match"
	}

	var timediff float32
	timediff = float32(args[0].Float())

	var agoOrLater string
	if timediff >= 0 {
		agoOrLater = "ago"
	} else {
		timediff = -timediff
		agoOrLater = "later"
	}

	var humanReadableResult string

	if timediff < minLength {
		humanReadableResult = fmt.Sprintf("%.1f seconds %s", timediff, agoOrLater)
	} else if timediff < hourLength {
		humanReadableResult = fmt.Sprintf("%.1f minutes %s", timediff/minLength, agoOrLater)
	} else if timediff < dayLength {
		humanReadableResult = fmt.Sprintf("%.1f hours %s", timediff/hourLength, agoOrLater)
	} else if timediff < monthLength {
		humanReadableResult = fmt.Sprintf("%.1f days %s", timediff/dayLength, agoOrLater)
	} else if timediff < yearLength {
		humanReadableResult = fmt.Sprintf("%.1f months %s", timediff/monthLength, agoOrLater)
	} else {
		humanReadableResult = fmt.Sprintf("%.1f years %s", timediff/yearLength, agoOrLater)
	}

	return humanReadableResult
}

func unixTimeConverter(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "ERROR: number of arguments doesn't match"
	}

	var unixTime int64
	unixTime = int64(args[0].Int())

	value := time.Unix(unixTime, 0).UTC()

	return fmt.Sprintf("%s", value)
}

type requestStruct struct {
	Value, Action, Type string
}

func encodeDecode(this js.Value, args []js.Value) interface{} {
	if len(args) != 3 {
		return "ERROR: number of arguments doesn't match"
	}

	var value, action, encodingType string
	var result string

	value = args[0].String()
	action = args[1].String()
	encodingType = args[2].String()

	t := requestStruct{value, action, encodingType}

	switch t {
	case requestStruct{t.Value, "encode", "url"}:
		result = url.PathEscape(t.Value)
	case requestStruct{t.Value, "encode", "base64"}:
		result = base64.StdEncoding.EncodeToString([]byte(t.Value))
	case requestStruct{t.Value, "decode", "url"}:
		result, _ = url.PathUnescape(t.Value)
	case requestStruct{t.Value, "decode", "base64"}:
		decoded, err := base64.StdEncoding.DecodeString(t.Value)
		if err != nil {
			result = "ERROR: " + err.Error()
		} else {
			result = string(decoded)
		}
	default:
		result = "ERROR: 'Action' should be either 'encode' or 'decode' and 'Type' should be either 'url' or 'base64'"
	}

	return result
}
