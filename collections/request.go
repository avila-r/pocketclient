package collections

// CollectionRequest represents the body parameters for creating or updating a collection.
type CollectionRequest struct {
	// Optional.
	//
	// 15-character unique identifier for the collection.
	// If omitted, it will be auto-generated.
	ID *string `json:"id,omitempty"`

	// Required.
	//
	// Unique collection name, used as
	// the table name in the database.
	Name string `json:"name"`

	// Required.
	//
	// Specifies the collection type.
	// Can be "base", "auth", or "view".
	Type CollectionType `json:"type"`

	// Schema fields list.
	//
	// Required for "base" collections;
	// Optional for "auth" collections;
	// Autopopulated for "view" collections.
	Schema []CollectionRequestField `json:"schema,omitempty"`

	// Optional. Marks the collection as "system",
	// indicating it cannot be renamed or deleted.
	System *bool `json:"system,omitempty"`

	ListRule   *Rule `json:"listRule,omitempty"`   // Optional API rule for list action. Defines restrictions or filters.
	ViewRule   *Rule `json:"viewRule,omitempty"`   // Optional API rule for view action. Specifies view restrictions.
	CreateRule *Rule `json:"createRule,omitempty"` // Optional API rule for create action. For "view" collections, this rule must be null.
	UpdateRule *Rule `json:"updateRule,omitempty"` // Optional API rule for update action. For "view" collections, this rule must be null.
	DeleteRule *Rule `json:"deleteRule,omitempty"` // Optional API rule for delete action. For "view" collections, this rule must be null.

	// Optional.
	//
	// Defines indexes and unique constraints for the collection.
	// "View" collections do not support indexes.
	Indexes []string `json:"indexes,omitempty"`

	Options CollectionRequestOptions `json:"options,omitempty"` // Optional configuration for collection-specific options, including view and auth settings.
}

// CollectionRequestField represents an individual field within a collection schema.
type CollectionRequestField struct {
	ID       string                        `json:"id"`       // Unique identifier for the field within the collection.
	Name     string                        `json:"name"`     // Required name of the field.
	Type     FieldType                     `json:"type"`     // Required type of the field, e.g., string, number, etc.
	Required bool                          `json:"required"` // Specifies if the field is mandatory.
	Unique   bool                          `json:"unique"`   // Specifies if the field should contain unique values.
	Options  CollectionRequestFieldOptions `json:"options"`  // Options for the field, including constraints and validation.
}

// CollectionRequestFieldOptions provides additional configuration for a field in a collection schema.
type CollectionRequestFieldOptions struct {
	Min       *int     `json:"min,omitempty"`       // Optional minimum value constraint for numeric fields.
	Max       *int     `json:"max,omitempty"`       // Optional maximum value constraint for numeric fields.
	Pattern   *string  `json:"pattern"`             // Pattern that the field value should match (e.g., regex for strings).
	MaxSelect *int     `json:"maxSelect,omitempty"` // Optional maximum number of selections allowed for multiselect fields.
	MaxSize   *int     `json:"maxSize,omitempty"`   // Optional maximum size constraint for the field, typically for files.
	MimeTypes []string `json:"mimeTypes,omitempty"` // Optional list of allowed MIME types for file fields.
	Thumbs    *string  `json:"thumbs,omitempty"`    // Optional thumbnail configuration for images.
}

// CollectionRequestOptions holds configuration settings for either "view" or "auth" collection options.
type CollectionRequestOptions struct {
	ViewOptions *ViewOptions `json:"view,omitempty"` // Configuration options specific to "view" collections.
	AuthOptions *AuthOptions `json:"auth,omitempty"` // Configuration options specific to "auth" collections.
}

// ViewOptions defines settings specific to view collections, including query configuration.
type ViewOptions struct {
	Query string `json:"query"` // Required SQL SELECT statement for defining the view of the collection.
}

// AuthOptions defines settings specific to auth collections, including authentication and security configurations.
type AuthOptions struct {
	ManageRule         *string  `json:"manageRule,omitempty"`         // Optional rule for admin-level permissions, allowing management of auth records (e.g., changing password).
	AllowOAuth2Auth    *bool    `json:"allowOAuth2Auth,omitempty"`    // Allows OAuth2-based authentication if set to true.
	AllowUsernameAuth  *bool    `json:"allowUsernameAuth,omitempty"`  // Allows username-password-based authentication if set to true.
	AllowEmailAuth     *bool    `json:"allowEmailAuth,omitempty"`     // Allows email-password-based authentication if set to true.
	RequireEmail       *bool    `json:"requireEmail,omitempty"`       // Ensures email is required when creating/updating auth records.
	ExceptEmailDomains []string `json:"exceptEmailDomains,omitempty"` // List of email domains to exclude from sign-ups.
	OnlyEmailDomains   []string `json:"onlyEmailDomains,omitempty"`   // List of allowed email domains for sign-ups.
	OnlyVerified       *bool    `json:"onlyVerified,omitempty"`       // Requires users to be verified for API access.
	MinPasswordLength  *int     `json:"minPasswordLength,omitempty"`  // Minimum length for user passwords.
}
