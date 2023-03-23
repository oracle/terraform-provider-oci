// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AclGroup A list of ACLs. This can be associated with a group of objects to restrict the networks from which requests
// against those objects can originate.
// For an incoming request, all FORCE_DENY rules are evaluated first. If none match, all other rules are evaluated in
// order: first the rules in the bucket level ACL Group, then compartment level ACL Groups, then the tenancy level
// ACL Group. If any ACL Rule matches the request traffic, evaluation stops and the request is either ALLOWED or
// DENIED based on the matching ACL Rule's action field.
// If the requestor has insufficient IAM permissions, the request is denied and ACL group restrictions are not
// checked.
type AclGroup struct {

	// Unique identifier for the ACL Group.
	Id *string `mandatory:"true" json:"id"`

	// User-specified name for the ACL Group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User-specified description of the ACL Group.
	Description *string `mandatory:"true" json:"description"`

	// The compartment containing this ACL Group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether this ACL Group is associated with a bucket, compartment, or tenancy.
	AclType AclGroupAclTypeEnum `mandatory:"true" json:"aclType"`

	// An ordered list of ACL IDs. The ACLs are evaluated first to last when determining network restrictions.
	AclIds []Acl `mandatory:"true" json:"aclIds"`

	// The date and time the ACL Group was created as per RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the ACL Group was last modified as per RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// Specifies request types that are not subject to any ACLs in the ACL group.
	AclGroupExceptions []AclGroupAclGroupExceptionsEnum `mandatory:"false" json:"aclGroupExceptions,omitempty"`
}

func (m AclGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AclGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAclGroupAclTypeEnum(string(m.AclType)); !ok && m.AclType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclType: %s. Supported values are: %s.", m.AclType, strings.Join(GetAclGroupAclTypeEnumStringValues(), ",")))
	}

	for _, val := range m.AclGroupExceptions {
		if _, ok := GetMappingAclGroupAclGroupExceptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupExceptions: %s. Supported values are: %s.", val, strings.Join(GetAclGroupAclGroupExceptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AclGroupAclTypeEnum Enum with underlying type: string
type AclGroupAclTypeEnum string

// Set of constants representing the allowable values for AclGroupAclTypeEnum
const (
	AclGroupAclTypeTenancy     AclGroupAclTypeEnum = "TENANCY"
	AclGroupAclTypeCompartment AclGroupAclTypeEnum = "COMPARTMENT"
	AclGroupAclTypeBucket      AclGroupAclTypeEnum = "BUCKET"
)

var mappingAclGroupAclTypeEnum = map[string]AclGroupAclTypeEnum{
	"TENANCY":     AclGroupAclTypeTenancy,
	"COMPARTMENT": AclGroupAclTypeCompartment,
	"BUCKET":      AclGroupAclTypeBucket,
}

var mappingAclGroupAclTypeEnumLowerCase = map[string]AclGroupAclTypeEnum{
	"tenancy":     AclGroupAclTypeTenancy,
	"compartment": AclGroupAclTypeCompartment,
	"bucket":      AclGroupAclTypeBucket,
}

// GetAclGroupAclTypeEnumValues Enumerates the set of values for AclGroupAclTypeEnum
func GetAclGroupAclTypeEnumValues() []AclGroupAclTypeEnum {
	values := make([]AclGroupAclTypeEnum, 0)
	for _, v := range mappingAclGroupAclTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupAclTypeEnumStringValues Enumerates the set of values in String for AclGroupAclTypeEnum
func GetAclGroupAclTypeEnumStringValues() []string {
	return []string{
		"TENANCY",
		"COMPARTMENT",
		"BUCKET",
	}
}

// GetMappingAclGroupAclTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupAclTypeEnum(val string) (AclGroupAclTypeEnum, bool) {
	enum, ok := mappingAclGroupAclTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AclGroupAclGroupExceptionsEnum Enum with underlying type: string
type AclGroupAclGroupExceptionsEnum string

// Set of constants representing the allowable values for AclGroupAclGroupExceptionsEnum
const (
	AclGroupAclGroupExceptionsCopy        AclGroupAclGroupExceptionsEnum = "COPY"
	AclGroupAclGroupExceptionsReplication AclGroupAclGroupExceptionsEnum = "REPLICATION"
)

var mappingAclGroupAclGroupExceptionsEnum = map[string]AclGroupAclGroupExceptionsEnum{
	"COPY":        AclGroupAclGroupExceptionsCopy,
	"REPLICATION": AclGroupAclGroupExceptionsReplication,
}

var mappingAclGroupAclGroupExceptionsEnumLowerCase = map[string]AclGroupAclGroupExceptionsEnum{
	"copy":        AclGroupAclGroupExceptionsCopy,
	"replication": AclGroupAclGroupExceptionsReplication,
}

// GetAclGroupAclGroupExceptionsEnumValues Enumerates the set of values for AclGroupAclGroupExceptionsEnum
func GetAclGroupAclGroupExceptionsEnumValues() []AclGroupAclGroupExceptionsEnum {
	values := make([]AclGroupAclGroupExceptionsEnum, 0)
	for _, v := range mappingAclGroupAclGroupExceptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupAclGroupExceptionsEnumStringValues Enumerates the set of values in String for AclGroupAclGroupExceptionsEnum
func GetAclGroupAclGroupExceptionsEnumStringValues() []string {
	return []string{
		"COPY",
		"REPLICATION",
	}
}

// GetMappingAclGroupAclGroupExceptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupAclGroupExceptionsEnum(val string) (AclGroupAclGroupExceptionsEnum, bool) {
	enum, ok := mappingAclGroupAclGroupExceptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
