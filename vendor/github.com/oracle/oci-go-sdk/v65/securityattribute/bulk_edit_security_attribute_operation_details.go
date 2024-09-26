// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkEditSecurityAttributeOperationDetails The representation of BulkEditSecurityAttributeOperationDetails
type BulkEditSecurityAttributeOperationDetails struct {

	// An enum-like description of the type of operation.
	// * `ADD_WHERE_ABSENT` adds a security attribute only if it does not already exist on the resource.
	// * `SET_WHERE_PRESENT` updates the value for a security attribute only if it is present on the resource.
	// * `ADD_OR_SET` combines the first two operations to add a security attribute if it does not already exist on the resource
	// or update the value if it is present on the resource.
	// * `REMOVE` removes the security attribute from the resource. It's removed from the resource regardless of the value.
	OperationType BulkEditSecurityAttributeOperationDetailsOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: ``{"oracle-zpr": {"td": {"value": "42", "mode": "audit"}}}``
	SecurityAttributes map[string]map[string]interface{} `mandatory:"true" json:"securityAttributes"`
}

func (m BulkEditSecurityAttributeOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkEditSecurityAttributeOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkEditSecurityAttributeOperationDetailsOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkEditSecurityAttributeOperationDetailsOperationTypeEnum Enum with underlying type: string
type BulkEditSecurityAttributeOperationDetailsOperationTypeEnum string

// Set of constants representing the allowable values for BulkEditSecurityAttributeOperationDetailsOperationTypeEnum
const (
	BulkEditSecurityAttributeOperationDetailsOperationTypeAddWhereAbsent  BulkEditSecurityAttributeOperationDetailsOperationTypeEnum = "ADD_WHERE_ABSENT"
	BulkEditSecurityAttributeOperationDetailsOperationTypeSetWherePresent BulkEditSecurityAttributeOperationDetailsOperationTypeEnum = "SET_WHERE_PRESENT"
	BulkEditSecurityAttributeOperationDetailsOperationTypeAddOrSet        BulkEditSecurityAttributeOperationDetailsOperationTypeEnum = "ADD_OR_SET"
	BulkEditSecurityAttributeOperationDetailsOperationTypeRemove          BulkEditSecurityAttributeOperationDetailsOperationTypeEnum = "REMOVE"
)

var mappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnum = map[string]BulkEditSecurityAttributeOperationDetailsOperationTypeEnum{
	"ADD_WHERE_ABSENT":  BulkEditSecurityAttributeOperationDetailsOperationTypeAddWhereAbsent,
	"SET_WHERE_PRESENT": BulkEditSecurityAttributeOperationDetailsOperationTypeSetWherePresent,
	"ADD_OR_SET":        BulkEditSecurityAttributeOperationDetailsOperationTypeAddOrSet,
	"REMOVE":            BulkEditSecurityAttributeOperationDetailsOperationTypeRemove,
}

var mappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnumLowerCase = map[string]BulkEditSecurityAttributeOperationDetailsOperationTypeEnum{
	"add_where_absent":  BulkEditSecurityAttributeOperationDetailsOperationTypeAddWhereAbsent,
	"set_where_present": BulkEditSecurityAttributeOperationDetailsOperationTypeSetWherePresent,
	"add_or_set":        BulkEditSecurityAttributeOperationDetailsOperationTypeAddOrSet,
	"remove":            BulkEditSecurityAttributeOperationDetailsOperationTypeRemove,
}

// GetBulkEditSecurityAttributeOperationDetailsOperationTypeEnumValues Enumerates the set of values for BulkEditSecurityAttributeOperationDetailsOperationTypeEnum
func GetBulkEditSecurityAttributeOperationDetailsOperationTypeEnumValues() []BulkEditSecurityAttributeOperationDetailsOperationTypeEnum {
	values := make([]BulkEditSecurityAttributeOperationDetailsOperationTypeEnum, 0)
	for _, v := range mappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkEditSecurityAttributeOperationDetailsOperationTypeEnumStringValues Enumerates the set of values in String for BulkEditSecurityAttributeOperationDetailsOperationTypeEnum
func GetBulkEditSecurityAttributeOperationDetailsOperationTypeEnumStringValues() []string {
	return []string{
		"ADD_WHERE_ABSENT",
		"SET_WHERE_PRESENT",
		"ADD_OR_SET",
		"REMOVE",
	}
}

// GetMappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnum(val string) (BulkEditSecurityAttributeOperationDetailsOperationTypeEnum, bool) {
	enum, ok := mappingBulkEditSecurityAttributeOperationDetailsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
