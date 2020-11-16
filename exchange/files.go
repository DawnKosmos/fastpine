package exchange

import (
	"os"
)

/*Checking files working like this
Bybit sends a request if a file is already existing. If the anwser is true. An ohclv chart gets parsed
and sent back
if the anwser is false, the exchanges sends back a ohclv which than gets written in a newly created file
*/
func FileExistCheck(str string) bool {
	src := "src/" + str
	_, err := os.Open(src) // For read access.
	if err != nil {
		return false
	}
	return true
}
