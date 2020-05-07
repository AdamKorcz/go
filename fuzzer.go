// 1: Include the library we are fuzzing
package jsoniter

import(
	"github.com/json-iterator/go"
)
// 2: Write the fuzzer
func FuzzParseString(data []byte) int {
	s = string(data)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, s)
	iter.ReadArray()
	return 1
}
