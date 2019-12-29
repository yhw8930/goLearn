package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

//根据网页链接获取到网页内容
func Fetch(url string) ([]byte, error) {
	/*resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()*/
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("wrong http request: %s", err.Error())
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong ststus code: %d", resp.StatusCode)
	}
	newReader := bufio.NewReader(resp.Body)
	e := determineEncoding(newReader)
	utf8Reader := transform.NewReader(newReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//url的html转码
func determineEncoding(newReader *bufio.Reader) encoding.Encoding {
	bytes, err := newReader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
