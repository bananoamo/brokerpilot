package connector

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"time"
)

//NewBrokerPilot create BrokerPilot object
func NewBrokerPilot(s *Settings) (*BrokerPilot, error) {
	if s.HostURL == `` {
		return nil, fmt.Errorf(`empty host url`)
	}
	return &BrokerPilot{s}, nil
}

//Token returns the assigned token from Settings
func (b *BrokerPilot) Token() string {
	return b.settings.Token
}

//Host returns the assigned host from Settings
func (b *BrokerPilot) Host() string {
	return b.settings.HostURL
}

//Proxy returns the assigned proxy from Settings
func (b *BrokerPilot) Proxy() string {
	return b.settings.ProxyURL
}

//IsActiveSocket returns the assigned IsActiveSocket from Settings
func (b *BrokerPilot) IsActiveSocket() bool {
	return b.settings.SocketTunnel
}

//ProxyDialTimeout returns the assigned ProxyDialTimeout from Settings
func (b *BrokerPilot) ProxyDialTimeout() time.Duration {
	return b.settings.ProxyDialTimeout
}

/*GetQuery does http.MethodGet request for simple and proxy requests

Example:

 headers := map[string]string{
		"Accept": `text/plain; charset=utf-8`,
 }

 body := []byte("request body") // use nil instead empty body
 uri  := "http://example.com"

 raw, status, err := GetQuery(headers, body, uri)
*/
func (b *BrokerPilot) GetQuery(h headers, body []byte, uri string) ([]byte, int, error) {

	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// sets http headers from map
	for k, v := range h {
		request.Header.Set(k, v)
	}

	request.Header.SetMethod("GET")
	request.SetRequestURI(uri)
	request.SetBody(body)

	client := &fasthttp.Client{}

	if b.Proxy() != `` {
		client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(b.Proxy(), b.ProxyDialTimeout())
	}

	if err := client.Do(request, response); err != nil {
		return nil, 0, err
	}

	return response.Body(), response.StatusCode(), nil
}

/*GetQueryBySocketProxy does http.MethodGet request for socket proxy requests ONLY

Example:

 headers := map[string]string{
		"Accept": `text/plain; charset=utf-8`,
 }

 body := []byte("request body") // use nil instead empty body
 uri  := "http://example.com"

 raw, status, err := GetQueryBySocketProxy(headers, body, uri)
*/
func (b *BrokerPilot) GetQueryBySocketProxy(h headers, body []byte, uri string) ([]byte, int, error) {

	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// sets http headers from map
	for k, v := range h {
		request.Header.Set(k, v)
	}

	request.Header.SetMethod("GET")
	request.SetRequestURI(uri)
	request.SetBody(body)

	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpSocksDialer("socks5://" + b.Proxy()),
	}

	if err := client.Do(request, response); err != nil {
		return nil, 0, err
	}

	return response.Body(), response.StatusCode(), nil
}
