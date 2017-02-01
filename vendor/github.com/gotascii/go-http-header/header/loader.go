// Package header implements loading structs into http.Headers.
//
// As a simple example:
//
// 	type Options struct {
// 		IfMatch    string `header:"if-match"`
// 		RetryToken string `header:"retry-token"`
// 	}
//
// opt := Options{ "6d82cbb050ddc7fa9cbb659014546e59", "my-custom-token" }

// req, _ := http.NewRequest(http.MethodGet, url, nil)
// header.LoadStruct(&req.Header, opts)
//
// The exact mapping between Go values and Header values is described in the
// documentation for the LoadStruct() function.

package header

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

// LoadStruct encodes the field values of h and loads them into the provided
// http.Header.
//
// LoadStruct expects to be passed a struct, and traverses it recursively using
// the following encoding rules.
//
// Each exported struct field is encoded and then loaded into the header unless
//
//	- the field's tag is "-", or
//	- the field is empty and its tag specifies the "omitempty" option
//
// The empty values are false, 0, any nil pointer or interface value, any array
// slice, map, or string of length zero, and any time.Time that returns true
// for IsZero().
//
// The header field name defaults to the struct field name but can be
// specified in the struct field's tag value.  The "header" key in the struct
// field's tag value is the key name, followed by an optional comma and
// options.  For example:
//
// 	// Field is ignored by this package.
// 	Field int `header:"-"`
//
// 	// Field appears as header field "myName".
// 	Field int `header:"myName"`
//
// 	// Field appears as header field "myName" and the field is omitted if
// 	// its value is empty
// 	Field int `header:"myName,omitempty"`
//
// 	// Field appears as header field "Field" (the default), but the field
// 	// is skipped if empty.  Note the leading comma.
// 	Field int `header:",omitempty"`
//
// For encoding individual field values, the following type-dependent rules
// apply:
//
// Boolean values default to encoding as the strings "true" or "false".
// Including the "int" option signals that the field should be encoded as the
// strings "1" or "0".
//
// time.Time values default to encoding as RFC3339 timestamps.  Including the
// "unix" option signals that the field should be encoded as a Unix time (see
// time.Unix())
//
// Slice and Array values are ignored.
//
// Anonymous struct fields are usually encoded as if their inner exported fields
// were fields in the outer struct, subject to the standard Go visibility rules.
// An anonymous struct field with a name given in its header tag is treated as
// having that name, rather than being anonymous.
//
// Non-nil pointer values are encoded as the value pointed to.
//
// Nested structs are ignored.
//
// All other values are encoded using their default string representation.
//
// Multiple fields that encode to the same header field will result in
// unpredictable behavior.
func NewFromStruct(v interface{}) (http.Header, error) {
	header := make(http.Header)
	if err := LoadStruct(&header, v); err != nil {
		return nil, err
	}
	return header, nil
}

func LoadStruct(header *http.Header, v interface{}) error {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	if v == nil {
		return nil
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("header: NewFromStruct() expects struct input. Got %v", val.Kind())
	}

	err := reflectValue(header, val)
	return err
}

// reflectValue populates the header field from the struct fields in val.
// Embedded structs are followed recursively (using the rules defined in the
// Values function documentation) breadth-first.
func reflectValue(header *http.Header, val reflect.Value) error {
	var embedded []reflect.Value

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("header")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if name == "" {
			if sf.Anonymous && sv.Kind() == reflect.Struct {
				// save embedded struct for later processing
				embedded = append(embedded, sv)
				continue
			}

			name = sf.Name
		}

		if opts.Contains("omitempty") && isEmptyValue(sv) {
			continue
		}

		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			continue
		}

		if sv.Type() == timeType {
			header.Set(name, valueString(sv, opts))
			continue
		}

		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}

		if sv.Kind() == reflect.Struct {
			continue
		}

		header.Set(name, valueString(sv, opts))
	}

	for _, f := range embedded {
		if err := reflectValue(header, f); err != nil {
			return err
		}
	}

	return nil
}

// valueString returns the string representation of a value.
func valueString(v reflect.Value, opts tagOptions) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	if v.Kind() == reflect.Bool && opts.Contains("int") {
		if v.Bool() {
			return "1"
		}
		return "0"
	}

	if v.Type() == timeType {
		t := v.Interface().(time.Time)
		if opts.Contains("unix") {
			return strconv.FormatInt(t.Unix(), 10)
		}
		return t.Format(time.RFC3339)
	}

	return fmt.Sprint(v.Interface())
}

// isEmptyValue checks if a value should be considered empty for the purposes
// of omitting fields with the "omitempty" option.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	if v.Type() == timeType {
		return v.Interface().(time.Time).IsZero()
	}

	return false
}

// tagOptions is the string following a comma in a struct field's "header" tag, or
// the empty string. It does not include the leading comma.
type tagOptions []string

// parseTag splits a struct field's header tag into its name and comma-separated
// options.
func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

// Contains checks whether the tagOptions contains the specified option.
func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}
