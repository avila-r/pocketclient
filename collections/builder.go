package collections

import (
	"errors"
)

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
		b.typ = ptr(TypeBase)
	}

	if b.sys == nil {
		b.sys = ptr(false)
	}
}

func (b *CollectionBuilder) validate() error {
	if b.typ == ptr(TypeBase) && (b.schema == nil || len(*b.schema) <= 0) {
		return errors.New("schema is required for base collections")
	}

	if b.sys == nil {
		return errors.New("system must be true or false")
	}

	if i := b.indexes; b.typ == ptr(TypeView) && i != nil && len(*i) > 0 {
		return errors.New("view collections don't support indexes")
	}

	if b.options != nil && b.typ == ptr(TypeView) && (b.options.Query == nil || b.options.Query == ptr("")) {
		return errors.New("query is required for view collections")
	}

	if b.rules != nil && b.typ == ptr(TypeView) && b.rules.CreateRule != nil && b.rules.CreateRule != ptr(Rule("")) {
		return errors.New("create rule must be null for view collections")
	}

	if b.rules != nil && b.typ == ptr(TypeView) && b.rules.UpdateRule != nil && b.rules.UpdateRule != ptr(Rule("")) {
		return errors.New("update rule must be null for view collections")
	}

	if b.rules != nil && b.typ == ptr(TypeView) && b.rules.DeleteRule != nil && b.rules.DeleteRule != ptr(Rule("")) {
		return errors.New("delete rule must be null for view collections")
	}

	return nil
}

func (b *CollectionBuilder) build() *CollectionRequest {
	fields := []CollectionRequestField{}
	for _, field := range *b.schema {
		options := CollectionRequestFieldOptions{
			Min:       field.Options.Min,
			Max:       field.Options.Max,
			Pattern:   field.Options.Pattern,
			MaxSelect: field.Options.MaxSelect,
			MaxSize:   field.Options.MaxSize,
			MimeTypes: field.Options.MimeTypes,
			Thumbs:    field.Options.Thumbs,
		}

		fields = append(fields, CollectionRequestField{
			Name:     field.Name,
			Type:     field.Type,
			Required: field.Required,
			Unique:   field.Unique,
			Options:  options,
		})
	}

	var view *ViewOptions = &ViewOptions{}
	if b.options != nil && b.options.Query != nil {
		view.Query = *b.options.Query
	} else {
		view = nil
	}

	var auth *AuthOptions
	if b.options != nil {
		auth = &AuthOptions{
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
	}

	options := &CollectionRequestOptions{
		ViewOptions: view,
		AuthOptions: auth,
	}

	var (
		listRule   *Rule
		viewRule   *Rule
		createRule *Rule
		updateRule *Rule
		deleteRule *Rule
	)

	if b.rules != nil {
		listRule = b.rules.ListRule
		viewRule = b.rules.ViewRule
		createRule = b.rules.CreateRule
		updateRule = b.rules.UpdateRule
		deleteRule = b.rules.DeleteRule
	}

	indexes := Indexes{}
	if b.indexes != nil {
		indexes = *b.indexes
	}

	request := CollectionRequest{
		Name:       *b.name,
		Type:       *b.typ,
		System:     b.sys,
		Schema:     fields,
		ListRule:   listRule,
		ViewRule:   viewRule,
		CreateRule: createRule,
		UpdateRule: updateRule,
		DeleteRule: deleteRule,
		Indexes:    indexes,
		Options:    *options,
	}

	return &request
}

func (b *CollectionBuilder) BuildPatch() *CollectionPatch {
	fields := []CollectionPatchField{}
	for _, field := range *b.schema {
		options := CollectionPatchFieldOptions{
			Min:       field.Options.Min,
			Max:       field.Options.Max,
			Pattern:   field.Options.Pattern,
			MaxSelect: field.Options.MaxSelect,
			MaxSize:   field.Options.MaxSize,
			MimeTypes: field.Options.MimeTypes,
			Thumbs:    field.Options.Thumbs,
		}

		fields = append(fields, CollectionPatchField{
			Name:     &field.Name,
			Type:     &field.Type,
			Required: &field.Required,
			Unique:   &field.Unique,
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

func ptr[T any](v T) *T {
	return &v
}
