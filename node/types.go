package node

// ===== public types =====

type ResponseNodeSettings struct {
	Result         []Results `json:"Result"`
	Errors         []string  `json:"Errors"`
	RequestId      string    `json:"RequestId"`
	ResponseMaster string    `json:"ResponseMaster"`
}

type Results struct {
	Node         string `json:"Node,omitempty"`
	Platform     string `json:"Platform,omitempty"`
	Result       Result `json:"Result,omitempty"`
	ResponseNode string `json:"Response_node,omitempty"`
}

type Result struct {
	Accounts            Accounts           `json:"Accounts,omitempty"`
	AccountsGroups      AccountsGroups     `json:"AccountsGroups,omitempty"`
	SymbolsGroups       SymbolsGroups      `json:"SymbolsGroups,omitempty"`
	MiniAccountsGroups  MiniAccountsGroups `json:"MiniAccountsGroups,omitempty"`
	SymbolsRolloverMode string             `json:"SymbolsRolloverMode,omitempty"`
	IsDemo              bool               `json:"IsDemo,omitempty"`
	Node                string             `json:"Node,omitempty"`
	Platform            string             `json:"Platform,omitempty"`
	Tags                []string           `json:"Tags,omitempty"`
	Clusters            []string           `json:"Clusters,omitempty"`
}

//Accounts
//
//Excluded List of accounts included in the exclusion list
//
//Included List of accounts included in the inclusion list
type Accounts struct {
	Included []StringInclusion `json:"Included,omitempty"`
	Excluded []Int64Exclusion  `json:"Excluded,omitempty"`
}

//AccountsGroups
//
//Excluded List of account groups included in the exclusion list
//
//Included List of account groups included in the inclusion list
type AccountsGroups struct {
	Included []StringInclusion `json:"Included,omitempty"`
	Excluded []StringExclusion `json:"Excluded,omitempty"`
}

//SymbolsGroups
//
//Excluded List of symbol groups included in the exclusion list
//
//Included List of symbol groups included in the inclusion list
type SymbolsGroups struct {
	Included []StringInclusion `json:"Included,omitempty"`
	Excluded []StringExclusion `json:"Excluded,omitempty"`
}

//MiniAccountsGroups
//
//Excluded List of mini accounts groups included in the exclusion list
//
//Included List of mini accounts groups included in the inclusion list
type MiniAccountsGroups struct {
	Included []StringInclusion `json:"Included,omitempty"`
	Excluded []StringExclusion `json:"Excluded,omitempty"`
}

type Settings struct {
	Accounts            Accounts           `json:"Accounts,omitempty"`
	AccountsGroups      AccountsGroups     `json:"AccountsGroups,omitempty"`
	SymbolsGroups       SymbolsGroups      `json:"SymbolsGroups,omitempty"`
	MiniAccountsGroups  MiniAccountsGroups `json:"MiniAccountsGroups,omitempty"`
	SymbolsRolloverMode string             `json:"SymbolsRolloverMode,omitempty"`
	IsDemo              bool               `json:"IsDemo,omitempty"`
	Tags                []string           `json:"Tags,omitempty"`
	Nodes               []string           `json:"Nodes,omitempty"`
	RequestId           string             `json:"RequestId,omitempty"`
	RequestName         string             `json:"RequestName,omitempty"`
}

type ResponseSaveNodeSettings struct {
	Result         []ResultSaveNodeSettings `json:"Result,omitempty"`
	RequestId      string                   `json:"RequestId,omitempty"`
	ResponseMaster string                   `json:"ResponseMaster,omitempty"`
	Errors         []string                 `json:"Errors,omitempty"`
}

type ResultSaveNodeSettings struct {
	Node         string `json:"Node,omitempty"`
	Platform     string `json:"Platform,omitempty"`
	Result       string `json:"Result,omitempty"`
	ResponseNode string `json:"ResponseNode,omitempty"`
}

// ===== private types =====

type StringInclusion struct {
	Item string `json:"Item,omitempty"`
}

type StringExclusion struct {
	Item string `json:"Item,omitempty"`
}

type Int64Exclusion struct {
	Item int64 `json:"Item,omitempty"`
}
