package connector

import "time"

/*Settings configurations

Example:

 Token does not need by default
 ProxyURL MUST set for `socket proxy` or `proxy` request
 ProxyDialTimeout MUST set for `proxy` requests ONLY
 SocketTunnel MUST set for `socket proxy` request ONLY

*/
type Settings struct {
	Token            string
	HostURL          string
	ProxyURL         string
	SocketTunnel     bool
	ProxyDialTimeout time.Duration
}

//BrokerPilot ...
type BrokerPilot struct {
	settings *Settings
}

//headers http header map
type headers map[string]string
