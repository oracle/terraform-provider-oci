// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Operations Each patch operation object MUST have exactly one "op" member, whose value indicates the operation to perform and MAY be one of "add", "remove", or "replace". See Section 3.5.2 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.5.2) for details.
type Operations struct {

	// Defines the operation to be performed for this Patch. If op=remove, value is not required.
	Op OperationsOpEnum `mandatory:"true" json:"op"`

	// String containing an attribute path describing the target of the operation. The "path" attribute is OPTIONAL for "add" and "replace" and is REQUIRED for "remove" operations. See Section 3.5.2 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.5.2) for details
	Path *string `mandatory:"true" json:"path"`

	// The value could be either a simple value attribute e.g. string or number OR complex like map of the attributes to be added or replaced OR multivalues complex attributes.q1
	Value *interface{} `mandatory:"false" json:"value"`
}

func (m Operations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Operations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationsOpEnum(string(m.Op)); !ok && m.Op != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Op: %s. Supported values are: %s.", m.Op, strings.Join(GetOperationsOpEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperationsOpEnum Enum with underlying type: string
type OperationsOpEnum string

// Set of constants representing the allowable values for OperationsOpEnum
const (
	OperationsOpAdd     OperationsOpEnum = "ADD"
	OperationsOpRemove  OperationsOpEnum = "REMOVE"
	OperationsOpReplace OperationsOpEnum = "REPLACE"
)

var mappingOperationsOpEnum = map[string]OperationsOpEnum{
	"ADD":     OperationsOpAdd,
	"REMOVE":  OperationsOpRemove,
	"REPLACE": OperationsOpReplace,
}

var mappingOperationsOpEnumLowerCase = map[string]OperationsOpEnum{
	"add":     OperationsOpAdd,
	"remove":  OperationsOpRemove,
	"replace": OperationsOpReplace,
}

// GetOperationsOpEnumValues Enumerates the set of values for OperationsOpEnum
func GetOperationsOpEnumValues() []OperationsOpEnum {
	values := make([]OperationsOpEnum, 0)
	for _, v := range mappingOperationsOpEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationsOpEnumStringValues Enumerates the set of values in String for OperationsOpEnum
func GetOperationsOpEnumStringValues() []string {
	return []string{
		"ADD",
		"REMOVE",
		"REPLACE",
	}
}

// GetMappingOperationsOpEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationsOpEnum(val string) (OperationsOpEnum, bool) {
	enum, ok := mappingOperationsOpEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
