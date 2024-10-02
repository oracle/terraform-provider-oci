// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeDeleteSecurityAttributeDefinition        WorkRequestOperationTypeEnum = "DELETE_SECURITY_ATTRIBUTE_DEFINITION"
	WorkRequestOperationTypeDeleteNonEmptySecurityAttributeNamespace WorkRequestOperationTypeEnum = "DELETE_NON_EMPTY_SECURITY_ATTRIBUTE_NAMESPACE"
	WorkRequestOperationTypeBulkDeleteSecurityAttributes             WorkRequestOperationTypeEnum = "BULK_DELETE_SECURITY_ATTRIBUTES"
	WorkRequestOperationTypeBulkEditOfSecurityAttributes             WorkRequestOperationTypeEnum = "BULK_EDIT_OF_SECURITY_ATTRIBUTES"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"DELETE_SECURITY_ATTRIBUTE_DEFINITION":          WorkRequestOperationTypeDeleteSecurityAttributeDefinition,
	"DELETE_NON_EMPTY_SECURITY_ATTRIBUTE_NAMESPACE": WorkRequestOperationTypeDeleteNonEmptySecurityAttributeNamespace,
	"BULK_DELETE_SECURITY_ATTRIBUTES":               WorkRequestOperationTypeBulkDeleteSecurityAttributes,
	"BULK_EDIT_OF_SECURITY_ATTRIBUTES":              WorkRequestOperationTypeBulkEditOfSecurityAttributes,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"delete_security_attribute_definition":          WorkRequestOperationTypeDeleteSecurityAttributeDefinition,
	"delete_non_empty_security_attribute_namespace": WorkRequestOperationTypeDeleteNonEmptySecurityAttributeNamespace,
	"bulk_delete_security_attributes":               WorkRequestOperationTypeBulkDeleteSecurityAttributes,
	"bulk_edit_of_security_attributes":              WorkRequestOperationTypeBulkEditOfSecurityAttributes,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"DELETE_SECURITY_ATTRIBUTE_DEFINITION",
		"DELETE_NON_EMPTY_SECURITY_ATTRIBUTE_NAMESPACE",
		"BULK_DELETE_SECURITY_ATTRIBUTES",
		"BULK_EDIT_OF_SECURITY_ATTRIBUTES",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
