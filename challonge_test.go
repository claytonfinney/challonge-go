package challonge

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
)

// This testClient and testClientFile pattern is adapted from zmb3's Spotify API client.
// It is, of course, also free to be used as seen fit.
// https://github.com/zmb3/spotify
func testClient(code int, body io.Reader, validators ...func(*http.Request)) (*ChallongeClient, *httptest.Server) {
	// This URL can be changed for remote testing etc
	url := "localhost:52504"
	// New Unstarted Server so we can add tls config
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.WriteHeader(code)
		io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	// tls config and listener for mock server
	l, _ := net.Listen("tcp", url)
	server.Listener = l
	config := &tls.Config{InsecureSkipVerify: true}
	server.TLS = config
	// Configure our test challonge client
	client := NewChallongeTestClient(url+"/", os.Getenv("CHALLONGE_USERNAME"), os.Getenv("CHALLONGE_API_KEY"))
	return client, server
}

func testClientFile(code int, filename string, validators ...func(*http.Request)) (*ChallongeClient, *httptest.Server) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testClient(code, f)
}
