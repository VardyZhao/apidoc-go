package dto

type CollectionDTO struct {
	Collection struct {
		Info     Info           `json:"info"`
		Item     []Item         `json:"item,omitempty"`
		Event    []Event        `json:"event,omitempty"`
		Variable []VariableItem `json:"variable,omitempty"`
	} `json:"collection"`
}

type Info struct {
	PostmanID     string `json:"_postman_id"`
	Name          string `json:"name"`
	Schema        string `json:"schema"`
	UpdatedAt     string `json:"updatedAt"`
	CreatedAt     string `json:"createdAt"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	UID           string `json:"uid"`
}

type Item struct {
	Name                    string          `json:"name"`
	Item                    []Item          `json:"item,omitempty"`
	ID                      string          `json:"id"`
	ProtocolProfileBehavior ProfileBehavior `json:"protocolProfileBehavior,omitempty"`
	Request                 *Request        `json:"request,omitempty"` // 指针以区分是否为空
	Response                []Response      `json:"response,omitempty"`
	UID                     string          `json:"uid"`
}

type ProfileBehavior struct {
	DisableBodyPruning bool `json:"disableBodyPruning"`
}

type Request struct {
	Method      string   `json:"method"`
	Header      []Header `json:"header"`
	URL         URL      `json:"url"`
	Body        *Body    `json:"body,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type URL struct {
	Raw  string   `json:"raw"`
	Host []string `json:"host"`
	Path []string `json:"path"`
}

type Body struct {
	Mode    string       `json:"mode"`
	Raw     string       `json:"raw,omitempty"`
	Options *BodyOptions `json:"options,omitempty"`
}

type BodyOptions struct {
	Raw struct {
		Language string `json:"language"`
	} `json:"raw"`
}

type Response struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	OriginalRequest Request          `json:"originalRequest"`
	Status          string           `json:"status"`
	Code            int              `json:"code"`
	Headers         []ResponseHeader `json:"header"`
	Body            *string          `json:"body,omitempty"`
}

type Event struct {
	Listen string `json:"listen"`
	Script Script `json:"script"`
}

type Script struct {
	ID   string   `json:"id"`
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}

type VariableItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResponseHeader struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}
