```
func main() {
	url := "http://192.168.163.128:8181/test"
	// resp := Get(url)
	// fmt.Println(resp.Text())

	headers := &Headers{Data: map[string]string{"data": "sadf"}}
	//, Files: map[string]string{"file1": "gorequests.go", "file2": "go-requests.exe"}}
	resp := Post(url, headers)
	fmt.Println(resp.Text())
	fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header)
}

```
