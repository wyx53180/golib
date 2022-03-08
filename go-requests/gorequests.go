package gorequests

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

type Response struct {
	body *[]byte
	http.Response
}

func new_Response(resp *http.Response, body *[]byte) *Response {
	response := Response{body: body}
	response.StatusCode = resp.StatusCode
	response.Header = resp.Header
	response.Request = resp.Request
	return &response
}

func (r *Response) Content() *[]byte {
	return r.body
}
func (r *Response) Text() string {
	return string(*r.body)
}

func (r *Response) Json() string {
	return ""
}

func Get(url string) *Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkerr(err)
	req.Close = true

	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	return new_Response(resp, &content)
}

type Headers struct {
	Data        map[string]string
	Files       map[string]string
	ContentType string
	Cookie      string
	UserAgent   string
}

func Post(url string, headers *Headers) *Response {
	// resp, err := http.Post(url, headers.ContentType, bytes.NewBuffer([]byte(body)))
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	defer bodyWriter.Close()
	// file.
	for k, v := range headers.Files {
		fileWriter, _ := bodyWriter.CreateFormFile(k, v)
		file, err := os.Open(v)
		checkerr(err)
		defer file.Close()
		io.Copy(fileWriter, file)
	}
	for k, v := range headers.Data {
		bodyWriter.WriteField(k, v)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bodyBuffer)
	checkerr(err)
	// rest by peer, Connection reset by peer 问题
	req.Close = true

	if headers.ContentType == "" {
		headers.ContentType = bodyWriter.FormDataContentType()
	}
	req.Header.Set("Content-Type", headers.ContentType)
	if headers.Cookie != "" {
		req.Header.Set("Cookie", headers.Cookie)
	}
	if headers.UserAgent != "" {
		req.Header.Set("user-agent", headers.UserAgent)
	}

	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	return new_Response(resp, &content)
}

// func main() {
// 	url := "http://192.168.163.128:8181/test"
// 	// resp := Get(url)
// 	// fmt.Println(resp.Text())

// 	headers := &Headers{Data: map[string]string{"data": "sadf"}}
// 	//, Files: map[string]string{"file1": "gorequests.go", "file2": "go-requests.exe"}}
// 	resp := Post(url, headers)
// 	fmt.Println(resp.Text())
// 	fmt.Println(resp.StatusCode)
// 	// fmt.Println(resp.Header)
// }
