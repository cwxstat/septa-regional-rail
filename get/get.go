package get

import (
	"errors"

	"context"
	"io/ioutil"
	"net/http"
	"time"
)

var debug = false

// Headers contains all HTTP headers to send
var Headers = make(map[string]string)

// Cookies contains all HTTP cookies to send
var Cookies = make(map[string]string)

// SetDebug sets the debug status
// Setting this to true causes the panics to be thrown and logged onto the console.
// Setting this to false causes the errors to be saved in the Error field in the returned struct.
func SetDebug(d bool) {
	debug = d
}

// Header sets a new HTTP header
func Header(n string, v string) {
	Headers[n] = v
}

func Cookie(n string, v string) {
	Cookies[n] = v
}

// GetWithClient returns the HTML returned by the url using a provided HTTP client
func GetWithClient(url string, client *http.Client) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		if debug {
			panic("Couldn't perform GET request to " + url + err.Error())
		}
		return "", errors.New("couldn't perform GET request to " + url)
	}
	// Set headers
	for hName, hValue := range Headers {
		req.Header.Set(hName, hValue)
	}
	// Set cookies
	for cName, cValue := range Cookies {
		req.AddCookie(&http.Cookie{
			Name:  cName,
			Value: cValue,
		})
	}
	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		if debug {
			panic("Couldn't perform GET request to " + url)
		}
		return "", errors.New("couldn't perform GET request to " + url)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if debug {
			panic("Unable to read the response body")
		}
		return "", errors.New("unable to read the response body")
	}
	return string(bytes), nil
}

type HTTP struct {
	client *http.Client
}

func Get(url string, client ...*http.Client) (string, error) {

	var newclient *http.Client
	if client == nil {
		newclient = &http.Client{}
	} else {
		newclient = client[0]
	}

	return GetWithClient(url, newclient)
}
