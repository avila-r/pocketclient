package collections

type CollectionType string

type Field struct {
	Name     string
	Type     FieldType
	Required bool
	Unique   bool
	Options  FieldOptions
}

type FieldType string

type FieldOptions struct {
	Min       *int
	Max       *int
	Pattern   *string
	MaxSelect *int
	MaxSize   *int
	MimeTypes []string
	Thumbs    *string
}

type Schema []Field

type Rule string

type Rules struct {
	ListRule   *Rule
	ViewRule   *Rule
	CreateRule *Rule
	UpdateRule *Rule
	DeleteRule *Rule
}

type Indexes []string

type CollectionOptions struct {
	// Only for view collections
	Query *string

	ManageRule         *string  // Optional rule for admin-level permissions, allowing management of auth records (e.g., changing password).
	AllowOAuth2Auth    *bool    // Allows OAuth2-based authentication if set to true.
	AllowUsernameAuth  *bool    // Allows username-password-based authentication if set to true.
	AllowEmailAuth     *bool    // Allows email-password-based authentication if set to true.
	RequireEmail       *bool    // Ensures email is required when creating/updating auth records.
	ExceptEmailDomains []string // List of email domains to exclude from sign-ups.
	OnlyEmailDomains   []string // List of allowed email domains for sign-ups.
	OnlyVerified       *bool    // Requires users to be verified for API access.
	MinPasswordLength  *int     // Minimum length for user passwords.
}
