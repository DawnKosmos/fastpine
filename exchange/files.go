package exchange

/*Checking files working like this
Exchange sends a request if a file is already existing. If the anwser is true. An ohclv chart gets parsed
and sent back
if the anwser is false, the exchanges sends back a ohclv which than gets written in a newly created file
*/
/*
func fileExistCheck(str string) bool {
	src := "src/" + str
	_, err := os.Open(src) // For read access.
	if err != nil {
		return false
	}
	return true
}

func checkFolder() bool {
	return false
}

func createFolder() {
}

func CheckExisting(exchange string, ticker string, resolution int, startTime, endTime time.Time) ([]Candle, error) {
	checkFolder()
	fileName := exchange + ticker + strconv.Itoa(resolution)

	if !fileExistCheck(fileName) {
		return []Candle{}, errors.New("File not existing")
	}

	file,


}
*/
