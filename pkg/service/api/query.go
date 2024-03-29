package api

type TokenListQuery struct {
	All           *string `form:"all" json:"all" structs:"all"`
	TokenNames    string  `form:"token_names" json:"token_name" structs:"token_name"`
	IsNativeToken *bool   `form:"is_native_token" json:"is_native_token" structs:"is_native_token"`
	Symbols       string  `form:"symbols" json:"symbols" structs:"symbols"`
	Chain         string  `form:"chains" json:"chains" structs:"chains"`
	Contracts     string  `form:"contracts" json:"contracts" structs:"contracts"`
	Sort          string  `form:"sort" json:"sort" structs:"sort"`
	Limit         int64   `form:"limit" json:"limit" structs:"limit" default:"20"`
	Page          int64   `form:"page" json:"page" structs:"page" default:"1"`
	Offset        int64   `form:"offset" json:"offset" structs:"offset" default:"0"`
}

type TokenQuery struct {
	TokenName     string `form:"token_name" json:"token_name" structs:"token_name"`
	IsNativeToken *bool  `form:"is_native_token" json:"is_native_token" structs:"is_native_token"`
	Chain         string `form:"chain" json:"chain" structs:"chain"`
	SymbolName    string `form:"symbol" json:"symbol" structs:"symbol"`
	Contract      string `form:"contract" json:"contracts" structs:"contract"`
}

type ValidatorQuery struct {
	Name   string `form:"name" json:"name" structs:"name"`
	NodeID string `form:"node_id" json:"node_id" structs:"node_id"`
}

type ValidatorListQuery struct {
	Name    string `form:"name" json:"name" structs:"name"`
	NodeIDs string `form:"node_ids" json:"node_ids" structs:"node_ids"`

	Sort   string `form:"sort" json:"sort" structs:"sort"`
	Limit  int64  `form:"limit" json:"limit" structs:"limit" default:"20"`
	Page   int64  `form:"page" json:"page" structs:"page" default:"1"`
	Offset int64  `form:"offset" json:"offset" structs:"offset" default:"0"`
}
