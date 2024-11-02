package collections

import "errors"

type CollectionBuilder struct {
	name    *string
	typ     *CollectionType
	schema  *Schema
	sys     *bool
	rules   *Rules
	indexes *Indexes
	options *CollectionOptions
}

func New(name string) *CollectionBuilder {
	return &CollectionBuilder{
		name: &name,
	}
}

func (b *CollectionBuilder) Type(t CollectionType) *CollectionBuilder {
	b.typ = &t
	return b
}

func (b *CollectionBuilder) Schema(s Schema) *CollectionBuilder {
	b.schema = &s
	return b
}

func (b *CollectionBuilder) Immutable() *CollectionBuilder {
	sys := true
	b.sys = &sys
	return b
}

func (b *CollectionBuilder) Rules(r Rules) *CollectionBuilder {
	b.rules = &r
	return b
}

func (b *CollectionBuilder) Indexes(i Indexes) *CollectionBuilder {
	b.indexes = &i
	return b
}

func (b *CollectionBuilder) Options(o CollectionOptions) *CollectionBuilder {
	b.options = &o
	return b
}

func (b *CollectionBuilder) Build() (*CollectionRequest, error) {
	b.defaults()
	if err := b.validate(); err != nil {
		return nil, err
	}

	return b.build(), nil
}

func (b *CollectionBuilder) defaults() {
	// Must be base (default), auth or view.
	if b.typ == nil || *b.typ == "" {
		*b.typ = TypeBase
	}

	if b.sys == nil {
		*b.sys = false
	}
}

func (b *CollectionBuilder) validate() error {
	if *b.typ == TypeBase && (b.schema == nil || len(*b.schema) <= 0) {
		return errors.New("schema is required for base collections")
	}

	if b.sys == nil {
		return errors.New("system must be true or false")
	}

	if r := b.rules.CreateRule; *b.typ == TypeView && (*r != "" || r != nil) {
		return errors.New("create rule must be null for view collections")
	}

	if r := b.rules.UpdateRule; *b.typ == TypeView && (*r != "" || r != nil) {
		return errors.New("update rule must be null for view collections")
	}

	if r := b.rules.DeleteRule; *b.typ == TypeView && (*r != "" || r != nil) {
		return errors.New("delete rule must be null for view collections")
	}

	if i := b.indexes; *b.typ == TypeView && (len(*i) > 0 || i != nil) {
		return errors.New("view collections don't support indexes")
	}

	if q := b.options.Query; *b.typ == TypeView && (q == nil || *q == "") {
		return errors.New("query is required for view collections")
	}

	return nil
}

func (b *CollectionBuilder) build() *CollectionRequest {
	fields := []CollectionRequestField{}
	for _, field := range *b.schema {
		options := CollectionRequestFieldOptions{
			Min:       field.O.Min,
			Max:       field.O.Max,
			Pattern:   field.O.Pattern,
			MaxSelect: field.O.MaxSelect,
			MaxSize:   field.O.MaxSize,
			MimeTypes: field.O.MimeTypes,
			Thumbs:    field.O.Thumbs,
		}

		fields = append(fields, CollectionRequestField{
			Name:     field.N,
			Type:     field.T,
			Required: field.R,
			Unique:   field.U,
			Options:  options,
		})
	}

	options := &CollectionRequestOptions{
		ViewOptions: &ViewOptions{
			Query: *b.options.Query,
		},

		AuthOptions: &AuthOptions{
			ManageRule:         b.options.ManageRule,
			AllowOAuth2Auth:    b.options.AllowOAuth2Auth,
			AllowUsernameAuth:  b.options.AllowUsernameAuth,
			AllowEmailAuth:     b.options.AllowEmailAuth,
			RequireEmail:       b.options.RequireEmail,
			ExceptEmailDomains: b.options.ExceptEmailDomains,
			OnlyEmailDomains:   b.options.OnlyEmailDomains,
			OnlyVerified:       b.options.OnlyVerified,
			MinPasswordLength:  b.options.MinPasswordLength,
		},
	}

	request := CollectionRequest{
		Name:       *b.name,
		Type:       *b.typ,
		System:     b.sys,
		Schema:     fields,
		ListRule:   b.rules.ListRule,
		ViewRule:   b.rules.ViewRule,
		CreateRule: b.rules.CreateRule,
		UpdateRule: b.rules.UpdateRule,
		DeleteRule: b.rules.DeleteRule,
		Indexes:    *b.indexes,
		Options:    *options,
	}

	return &request
}

func (b *CollectionBuilder) BuildPatch() *CollectionPatch {
	fields := []CollectionPatchField{}
	for _, field := range *b.schema {
		options := CollectionPatchFieldOptions{
			Min:       field.O.Min,
			Max:       field.O.Max,
			Pattern:   field.O.Pattern,
			MaxSelect: field.O.MaxSelect,
			MaxSize:   field.O.MaxSize,
			MimeTypes: field.O.MimeTypes,
			Thumbs:    field.O.Thumbs,
		}

		fields = append(fields, CollectionPatchField{
			Name:     &field.N,
			Type:     &field.T,
			Required: &field.R,
			Unique:   &field.U,
			Options:  &options,
		})
	}

	options := &CollectionPatchOptions{
		Query: b.options.Query,

		ManageRule:         b.options.ManageRule,
		AllowOAuth2Auth:    b.options.AllowOAuth2Auth,
		AllowUsernameAuth:  b.options.AllowUsernameAuth,
		AllowEmailAuth:     b.options.AllowEmailAuth,
		RequireEmail:       b.options.RequireEmail,
		ExceptEmailDomains: b.options.ExceptEmailDomains,
		OnlyEmailDomains:   b.options.OnlyEmailDomains,
		OnlyVerified:       b.options.OnlyVerified,
		MinPasswordLength:  b.options.MinPasswordLength,
	}

	patch := CollectionPatch{
		Name:       b.name,
		Type:       b.typ,
		System:     b.sys,
		Schema:     fields,
		ListRule:   b.rules.ListRule,
		ViewRule:   b.rules.ViewRule,
		CreateRule: b.rules.CreateRule,
		UpdateRule: b.rules.UpdateRule,
		DeleteRule: b.rules.DeleteRule,
		Indexes:    b.indexes,
		Options:    options,
	}

	return &patch
}
