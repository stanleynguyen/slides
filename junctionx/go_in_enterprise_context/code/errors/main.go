value, err := QueryDatabase()
if err != nil {
	fmt.Print(err.Error())
}
response, err := MakeAPICall(value)
if err != nil {
	fmt.Print(err.Error())
}
fmt.Print(response)
