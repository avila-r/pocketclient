package collections

type CollectionResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`

	Type   string                    `json:"type"`
	Name   string                    `json:"name"`
	System bool                      `json:"system"`
	Schema []CollectionResponseField `json:"schema"`

	ListRule   string  `json:"listRule"`
	ViewRule   string  `json:"viewRule"`
	CreateRule string  `json:"createRule"`
	UpdateRule string  `json:"updateRule"`
	DeleteRule *string `json:"deleteRule,omitempty"`

	Options *struct {
		ManageRule         *string  `json:"manageRule"`
		AllowOAuth2Auth    bool     `json:"allowOAuth2Auth"`
		AllowUsernameAuth  bool     `json:"allowUsernameAuth"`
		AllowEmailAuth     bool     `json:"allowEmailAuth"`
		RequireEmail       bool     `json:"requireEmail"`
		ExceptEmailDomains []string `json:"exceptEmailDomains"`
		OnlyEmailDomains   []string `json:"onlyEmailDomains"`
		MinPasswordLength  int      `json:"minPasswordLength"`
	} `json:"options,omitempty"`

	Indexes []string `json:"indexes"`
}

type CollectionResponseField struct {
	System   bool   `json:"system"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Unique   bool   `json:"unique"`
	Options  struct {
		Min       *int     `json:"min,omitempty"`
		Max       *int     `json:"max,omitempty"`
		Pattern   string   `json:"pattern"`
		MaxSelect *int     `json:"maxSelect,omitempty"`
		MaxSize   *int     `json:"maxSize,omitempty"`
		MimeTypes []string `json:"mimeTypes,omitempty"`
		Thumbs    *string  `json:"thumbs,omitempty"`
	} `json:"options"`
}
