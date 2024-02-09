// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AclGroupSummary The summary of an ACL Group.
type AclGroupSummary struct {

	// Unique identifier for the ACL Group.
	Id *string `mandatory:"true" json:"id"`

	// User-specified name for the ACL Group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User-specified description of the ACL Group.
	Description *string `mandatory:"true" json:"description"`

	// The compartment containing this ACL Group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether this ACL Group is associated with a bucket, compartment, or tenancy.
	AclGroupType AclGroupSummaryAclGroupTypeEnum `mandatory:"true" json:"aclGroupType"`

	// An ordered list of ACL IDs. The ACLs are evaluated first to last when determining network restrictions.
	AclIds []string `mandatory:"true" json:"aclIds"`

	// The date and time the ACL Group was created as per RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the ACL Group was last modified as per RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// Specifies request types that are not subject to any ACLs in the ACL group.
	AclGroupExceptions []AclGroupSummaryAclGroupExceptionsEnum `mandatory:"false" json:"aclGroupExceptions,omitempty"`
}

func (m AclGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AclGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAclGroupSummaryAclGroupTypeEnum(string(m.AclGroupType)); !ok && m.AclGroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupType: %s. Supported values are: %s.", m.AclGroupType, strings.Join(GetAclGroupSummaryAclGroupTypeEnumStringValues(), ",")))
	}

	for _, val := range m.AclGroupExceptions {
		if _, ok := GetMappingAclGroupSummaryAclGroupExceptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupExceptions: %s. Supported values are: %s.", val, strings.Join(GetAclGroupSummaryAclGroupExceptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AclGroupSummaryAclGroupTypeEnum Enum with underlying type: string
type AclGroupSummaryAclGroupTypeEnum string

// Set of constants representing the allowable values for AclGroupSummaryAclGroupTypeEnum
const (
	AclGroupSummaryAclGroupTypeTenancy     AclGroupSummaryAclGroupTypeEnum = "TENANCY"
	AclGroupSummaryAclGroupTypeCompartment AclGroupSummaryAclGroupTypeEnum = "COMPARTMENT"
	AclGroupSummaryAclGroupTypeBucket      AclGroupSummaryAclGroupTypeEnum = "BUCKET"
)

var mappingAclGroupSummaryAclGroupTypeEnum = map[string]AclGroupSummaryAclGroupTypeEnum{
	"TENANCY":     AclGroupSummaryAclGroupTypeTenancy,
	"COMPARTMENT": AclGroupSummaryAclGroupTypeCompartment,
	"BUCKET":      AclGroupSummaryAclGroupTypeBucket,
}

var mappingAclGroupSummaryAclGroupTypeEnumLowerCase = map[string]AclGroupSummaryAclGroupTypeEnum{
	"tenancy":     AclGroupSummaryAclGroupTypeTenancy,
	"compartment": AclGroupSummaryAclGroupTypeCompartment,
	"bucket":      AclGroupSummaryAclGroupTypeBucket,
}

// GetAclGroupSummaryAclGroupTypeEnumValues Enumerates the set of values for AclGroupSummaryAclGroupTypeEnum
func GetAclGroupSummaryAclGroupTypeEnumValues() []AclGroupSummaryAclGroupTypeEnum {
	values := make([]AclGroupSummaryAclGroupTypeEnum, 0)
	for _, v := range mappingAclGroupSummaryAclGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupSummaryAclGroupTypeEnumStringValues Enumerates the set of values in String for AclGroupSummaryAclGroupTypeEnum
func GetAclGroupSummaryAclGroupTypeEnumStringValues() []string {
	return []string{
		"TENANCY",
		"COMPARTMENT",
		"BUCKET",
	}
}

// GetMappingAclGroupSummaryAclGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupSummaryAclGroupTypeEnum(val string) (AclGroupSummaryAclGroupTypeEnum, bool) {
	enum, ok := mappingAclGroupSummaryAclGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AclGroupSummaryAclGroupExceptionsEnum Enum with underlying type: string
type AclGroupSummaryAclGroupExceptionsEnum string

// Set of constants representing the allowable values for AclGroupSummaryAclGroupExceptionsEnum
const (
	AclGroupSummaryAclGroupExceptionsCopy        AclGroupSummaryAclGroupExceptionsEnum = "COPY"
	AclGroupSummaryAclGroupExceptionsReplication AclGroupSummaryAclGroupExceptionsEnum = "REPLICATION"
)

var mappingAclGroupSummaryAclGroupExceptionsEnum = map[string]AclGroupSummaryAclGroupExceptionsEnum{
	"COPY":        AclGroupSummaryAclGroupExceptionsCopy,
	"REPLICATION": AclGroupSummaryAclGroupExceptionsReplication,
}

var mappingAclGroupSummaryAclGroupExceptionsEnumLowerCase = map[string]AclGroupSummaryAclGroupExceptionsEnum{
	"copy":        AclGroupSummaryAclGroupExceptionsCopy,
	"replication": AclGroupSummaryAclGroupExceptionsReplication,
}

// GetAclGroupSummaryAclGroupExceptionsEnumValues Enumerates the set of values for AclGroupSummaryAclGroupExceptionsEnum
func GetAclGroupSummaryAclGroupExceptionsEnumValues() []AclGroupSummaryAclGroupExceptionsEnum {
	values := make([]AclGroupSummaryAclGroupExceptionsEnum, 0)
	for _, v := range mappingAclGroupSummaryAclGroupExceptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupSummaryAclGroupExceptionsEnumStringValues Enumerates the set of values in String for AclGroupSummaryAclGroupExceptionsEnum
func GetAclGroupSummaryAclGroupExceptionsEnumStringValues() []string {
	return []string{
		"COPY",
		"REPLICATION",
	}
}

// GetMappingAclGroupSummaryAclGroupExceptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupSummaryAclGroupExceptionsEnum(val string) (AclGroupSummaryAclGroupExceptionsEnum, bool) {
	enum, ok := mappingAclGroupSummaryAclGroupExceptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
