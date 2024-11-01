package pocketclient

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Error(message string) error {
	return errors.New(message)
}

type (
	Json map[string]interface{}

	QueryParams map[string]string

	Pagination[T any] struct {
		Page       int `json:"page"`
		PerPage    int `json:"perPage"`
		TotalItems int `json:"totalItems"`
		Items      []T `json:"items"`
	}

	PaginationParams struct {
		Page    int
		PerPage int

		// Specify the ORDER BY fields.
		//
		// Add - / + (default) in front of the attribute for DESC / ASC order, eg.:
		//
		// -> // DESC by the insertion rowid and ASC by level
		// -> ?sort=-rowid,level
		//
		// Supported log sort fields:
		// rowid, id, created, updated, level, message and any data.* attribute.
		Sort string

		// Filter expression to filter/search the returned logs list, eg.:
		// ?filter=(data.url~'test.com' && level>0)
		//
		// Supported log filter fields:
		// id, created, updated, level, message and any data.* attribute.
		//
		// The syntax basically follows the format OPERAND OPERATOR OPERAND, where:
		//
		// OPERAND - could be any of the above field literal, string (single or double quoted), number, null, true, false
		// OPERATOR - is one of:
		//
		//	= Equal
		//	!= NOT equal
		//	> Greater than
		//	>= Greater than or equal
		//	< Less than
		//	<= Less than or equal
		//	~ Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard match)
		//	!~ NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard match)
		//	?= Any/At least one of Equal
		//	?!= Any/At least one of NOT equal
		//	?> Any/At least one of Greater than
		//	?>= Any/At least one of Greater than or equal
		//	?< Any/At least one of Less than
		//	?<= Any/At least one of Less than or equal
		//	?~ Any/At least one of Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard match)
		//	?!~ Any/At least one of NOT Like/Contains (if not specified auto wraps the right string OPERAND in a "%" for wildcard match)
		//	To group and combine several expressions you could use parenthesis (...), && (AND) and || (OR) tokens.
		//
		// Single line comments are also supported: '// Example comment'.
		Filter string
	}
)

func Marshal[T any](t T) ([]byte, error) {
	return json.Marshal(t)
}

func Unmarshal[T any](b []byte, r T) error {
	return json.Unmarshal(b, &r)
}

func (p PaginationParams) ToQueryParams() QueryParams {
	return QueryParams{
		"page":    fmt.Sprint(p.Page),
		"perPage": fmt.Sprint(p.PerPage),
		"sort":    p.Sort,
		"filter":  p.Filter,
	}
}
