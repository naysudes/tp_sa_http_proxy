package proxy

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	OKHeader = "HTTP/1.1 200 OK\r\n\r\n"
)

type Config struct {
	Listen string `yaml:"listen"`
	//Db      DBConfig `yaml:"db"`
}

type Server struct {
	httpClient *http.Client
	//db         *database.DB
}

func NewServer() (*Server, error) {
	var err error
	server := Server{
		httpClient: new(http.Client),
		//db:       db,
	}
	server.httpClient.Timeout = 5 * time.Second
	if err != nil {
		return nil, err
	}
	return &server, nil
}

func (s *Server) Run() {

	if err := http.ListenAndServe("127.0.0.1:8080", s); err != http.ErrServerClosed {
		return
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	//requestLogger(r, s.log)
	if r.Method == http.MethodConnect {
		s.HandleHttps(w, r)
	} else {
		s.HandleHttp(w, r)
	}
}

func (s *Server) HandleHttps(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Server) HandleHttp(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
