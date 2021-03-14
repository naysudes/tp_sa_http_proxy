package proxy

import (
	"io"
	"log"
	"net/http"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	request.RequestURI = ""
	request.Header.Del("Proxy-Connection")

	httpClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	proxyResponse, err := httpClient.Do(request)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer proxyResponse.Body.Close()

	copyHeader(responseWriter.Header(), proxyResponse.Header)
	responseWriter.WriteHeader(proxyResponse.StatusCode)
	io.Copy(responseWriter, proxyResponse.Body)
}

func copyHeader(to, from http.Header) {
	for header, values := range from {
		for _, value := range values {
			to.Add(header, value)
		}
	}
}
