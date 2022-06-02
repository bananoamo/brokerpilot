package node

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"test/brokerpilot/connector"
)

/*
GetNodeSettings request node settings

Example:

 bp, err := con.NewBrokerPilot(&connector.Settings{
 	Token: "OAuth-token", 		// does not need by default
 	HostURL: "https://esplanade.brokerpilot.net/bpapi",
 	ProxyURL: "localhost:8079", 	// MUST set for `socket proxy` or `proxy` requests
 	SocketTunnel: true, 		// MUST set for `socket proxy` request ONLY
 })

 nodes := []string{"node1", "node2"} // settings for `node1` and `node2` only
 nodes := []string{}                 // settings for all available `nodes`

 ResponseNodeSettings, status, err := node.GetNodeSettings(bp, nodes)

 if err != nil {
 	log.Fatal(err)
 }

 API https://indigosoft.atlassian.net/wiki/spaces/BA/pages/2116585185/RequestNodesSettings
*/
func GetNodeSettings(bp *connector.BrokerPilot, nodes []string) (*ResponseNodeSettings, int, error) {

	var resp []byte
	var status int
	var err error

	nodeSettings := &Settings{
		RequestId:   uuid.New().String(),
		Nodes:       nodes,
		RequestName: requestNodeSettings,
	}

	headers := map[string]string{
		"Accept": `application/json`,
	}

	if bp.Token() != `` {
		headers[`Authorization`] = fmt.Sprintf("Bearer %s", bp.Token())
	}

	raw, err := json.Marshal(nodeSettings)
	if err != nil {
		return nil, -1, err
	}

	uri := fmt.Sprintf("%s/?request=%s", bp.Host(), url.QueryEscape(string(raw)))

	switch bp.IsActiveSocket() {
	case true:
		resp, status, err = bp.GetQueryBySocketProxy(headers, nil, uri)
	default:
		resp, status, err = bp.GetQuery(headers, nil, uri)
	}

	if err != nil {
		return nil, status, err
	}

	res := &ResponseNodeSettings{}
	if err = json.Unmarshal(resp, res); err != nil {
		return nil, status, err
	}

	if len(res.Errors) != 0 {
		return res, status, fmt.Errorf("%+v", res.Errors)
	}

	return res, status, nil
}

/*SaveNodeSettings request to save node settings

Example:

 bp, err := con.NewBrokerPilot(&connector.Settings{
 	Token: "OAuth-token", 		// does not need by default
 	HostURL: "https://esplanade.brokerpilot.net/bpapi",
 	ProxyURL: "localhost:8079", 	// MUST set for `socket proxy` or `proxy` requests
 	SocketTunnel: true, 		// MUST set for `socket proxy` request ONLY
 })

 setting := &node.Settings{
 	Nodes: []string{}, 	// only one `node` in slice working
 	SymbolsRolloverMode: 	// one of these values [node.ChangeOpenPrice|node.SwapCommission],
 	// the rest struct fields is up to you
 }

 ResponseSaveNodeSettings, status, err := node.SaveNodeSettings(bp, Settings)

 if err != nil {
 	log.Fatal(err)
 }

 API https://indigosoft.atlassian.net/wiki/spaces/BA/pages/2116585200/RequestSaveNodeSettings
*/
func SaveNodeSettings(bp *connector.BrokerPilot, s *Settings) (*ResponseSaveNodeSettings, int, error) {
	// TODO: not tested on real api
	var resp []byte
	var status int
	var err error

	headers := map[string]string{
		"Accept": `application/json`,
	}

	if bp.Token() != `` {
		headers[`Authorization`] = fmt.Sprintf("Bearer %s", bp.Token())
	}

	s.RequestId = uuid.New().String()
	s.RequestName = requestSaveNodeSettings

	raw, err := json.Marshal(s)
	if err != nil {
		return nil, -1, err
	}

	uri := fmt.Sprintf("%s/?request=%s", bp.Host(), url.QueryEscape(string(raw)))

	switch bp.IsActiveSocket() {
	case true:
		resp, status, err = bp.GetQueryBySocketProxy(headers, nil, uri)
	default:
		resp, status, err = bp.GetQuery(headers, nil, uri)
	}

	if err != nil {
		return nil, status, err
	}

	res := &ResponseSaveNodeSettings{}
	if err = json.Unmarshal(resp, res); err != nil {
		return nil, status, err
	}

	if len(res.Errors) != 0 {
		return res, status, fmt.Errorf("%+v", res.Errors)
	}

	return res, status, nil
}
