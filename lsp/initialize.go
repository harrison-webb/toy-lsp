package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// lots more that goes here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	TextDocumentSync int  `json:"textDocumentSync"`
	HoverProvider    bool `json:"hoverProvider"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1, // 1 means send the full content of the document
				HoverProvider:    true,
			},
			ServerInfo: ServerInfo{
				Name:    "toylsp",
				Version: "0.1",
			},
		},
	}
}
