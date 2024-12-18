package collections

type CollectionPatch struct {
	Name       *string                 `json:"name"`
	Type       *CollectionType         `json:"type"`
	System     *bool                   `json:"system,omitempty"`
	Schema     []CollectionPatchField  `json:"schema,omitempty"`
	ListRule   *Rule                   `json:"listRule,omitempty"`
	ViewRule   *Rule                   `json:"viewRule,omitempty"`
	CreateRule *Rule                   `json:"createRule,omitempty"`
	UpdateRule *Rule                   `json:"updateRule,omitempty"`
	DeleteRule *Rule                   `json:"deleteRule,omitempty"`
	Indexes    *Indexes                `json:"indexes,omitempty"`
	Options    *CollectionPatchOptions `json:"options,omitempty"`
}

type CollectionPatchOptions struct {
	ManageRule         *string  `json:"manageRule,omitempty"`
	AllowOAuth2Auth    *bool    `json:"allowOAuth2Auth,omitempty"`
	AllowUsernameAuth  *bool    `json:"allowUsernameAuth,omitempty"`
	AllowEmailAuth     *bool    `json:"allowEmailAuth,omitempty"`
	RequireEmail       *bool    `json:"requireEmail,omitempty"`
	ExceptEmailDomains []string `json:"exceptEmailDomains,omitempty"`
	OnlyEmailDomains   []string `json:"onlyEmailDomains,omitempty"`
	OnlyVerified       *bool    `json:"onlyVerified,omitempty"`
	MinPasswordLength  *int     `json:"minPasswordLength,omitempty"`
	Query              *string  `json:"query,omitempty"`
}

type CollectionPatchField struct {
	System   *bool                        `json:"system,omitempty"`
	ID       *string                      `json:"id,omitempty"`
	Name     *string                      `json:"name,omitempty"`
	Type     *FieldType                   `json:"type,omitempty"`
	Required *bool                        `json:"required,omitempty"`
	Unique   *bool                        `json:"unique,omitempty"`
	Options  *CollectionPatchFieldOptions `json:"options,omitempty"`
}

type CollectionPatchFieldOptions struct {
	Min       *int     `json:"min,omitempty"`
	Max       *int     `json:"max,omitempty"`
	Pattern   *string  `json:"pattern,omitempty"`
	MaxSelect *int     `json:"maxSelect,omitempty"`
	MaxSize   *int     `json:"maxSize,omitempty"`
	MimeTypes []string `json:"mimeTypes,omitempty"`
	Thumbs    *string  `json:"thumbs,omitempty"`
}
