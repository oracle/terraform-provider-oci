// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HttpHeader Http header to add.
type HttpHeader struct {

	// Header name.
	Key *string `mandatory:"true" json:"key"`

	// Header value.
	Value *string `mandatory:"false" json:"value"`

	// If a header with the same name already exists in the request, OVERWRITE will overwrite the value,
	// APPEND will append to the existing value, or SKIP will keep the existing value.
	IfExists HttpHeaderIfExistsEnum `mandatory:"false" json:"ifExists,omitempty"`
}

func (m HttpHeader) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpHeader) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHttpHeaderIfExistsEnum(string(m.IfExists)); !ok && m.IfExists != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IfExists: %s. Supported values are: %s.", m.IfExists, strings.Join(GetHttpHeaderIfExistsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HttpHeaderIfExistsEnum Enum with underlying type: string
type HttpHeaderIfExistsEnum string

// Set of constants representing the allowable values for HttpHeaderIfExistsEnum
const (
	HttpHeaderIfExistsOverwrite HttpHeaderIfExistsEnum = "OVERWRITE"
	HttpHeaderIfExistsAppend    HttpHeaderIfExistsEnum = "APPEND"
	HttpHeaderIfExistsSkip      HttpHeaderIfExistsEnum = "SKIP"
)

var mappingHttpHeaderIfExistsEnum = map[string]HttpHeaderIfExistsEnum{
	"OVERWRITE": HttpHeaderIfExistsOverwrite,
	"APPEND":    HttpHeaderIfExistsAppend,
	"SKIP":      HttpHeaderIfExistsSkip,
}

var mappingHttpHeaderIfExistsEnumLowerCase = map[string]HttpHeaderIfExistsEnum{
	"overwrite": HttpHeaderIfExistsOverwrite,
	"append":    HttpHeaderIfExistsAppend,
	"skip":      HttpHeaderIfExistsSkip,
}

// GetHttpHeaderIfExistsEnumValues Enumerates the set of values for HttpHeaderIfExistsEnum
func GetHttpHeaderIfExistsEnumValues() []HttpHeaderIfExistsEnum {
	values := make([]HttpHeaderIfExistsEnum, 0)
	for _, v := range mappingHttpHeaderIfExistsEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpHeaderIfExistsEnumStringValues Enumerates the set of values in String for HttpHeaderIfExistsEnum
func GetHttpHeaderIfExistsEnumStringValues() []string {
	return []string{
		"OVERWRITE",
		"APPEND",
		"SKIP",
	}
}

// GetMappingHttpHeaderIfExistsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpHeaderIfExistsEnum(val string) (HttpHeaderIfExistsEnum, bool) {
	enum, ok := mappingHttpHeaderIfExistsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
