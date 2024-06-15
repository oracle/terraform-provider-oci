// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkEditZprOperationDetails The representation of BulkEditZprOperationDetails
type BulkEditZprOperationDetails struct {

	// An enum-like description of the type of operation.
	// * `ADD_WHERE_ABSENT` adds a zpr tag only if the tag does not already exist on the resource.
	// * `SET_WHERE_PRESENT` updates the value for a zpr tag only if the tag is present on the resource.
	// * `ADD_OR_SET` combines the first two operations to add a zpr tag if it does not already exist on the resource
	// or update the value for a zpr tag only if the tag is present on the resource.
	// * `REMOVE` removes the zpr tag from the resource. The tag is removed from the resource regardless of the tag value.
	OperationType BulkEditZprOperationDetailsOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Zpr tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: ``{"oracle-zpr": {"td": {"value": "42", "mode": "audit"}}}``
	ZprTags map[string]map[string]interface{} `mandatory:"true" json:"zprTags"`
}

func (m BulkEditZprOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkEditZprOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBulkEditZprOperationDetailsOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetBulkEditZprOperationDetailsOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BulkEditZprOperationDetailsOperationTypeEnum Enum with underlying type: string
type BulkEditZprOperationDetailsOperationTypeEnum string

// Set of constants representing the allowable values for BulkEditZprOperationDetailsOperationTypeEnum
const (
	BulkEditZprOperationDetailsOperationTypeAddWhereAbsent  BulkEditZprOperationDetailsOperationTypeEnum = "ADD_WHERE_ABSENT"
	BulkEditZprOperationDetailsOperationTypeSetWherePresent BulkEditZprOperationDetailsOperationTypeEnum = "SET_WHERE_PRESENT"
	BulkEditZprOperationDetailsOperationTypeAddOrSet        BulkEditZprOperationDetailsOperationTypeEnum = "ADD_OR_SET"
	BulkEditZprOperationDetailsOperationTypeRemove          BulkEditZprOperationDetailsOperationTypeEnum = "REMOVE"
)

var mappingBulkEditZprOperationDetailsOperationTypeEnum = map[string]BulkEditZprOperationDetailsOperationTypeEnum{
	"ADD_WHERE_ABSENT":  BulkEditZprOperationDetailsOperationTypeAddWhereAbsent,
	"SET_WHERE_PRESENT": BulkEditZprOperationDetailsOperationTypeSetWherePresent,
	"ADD_OR_SET":        BulkEditZprOperationDetailsOperationTypeAddOrSet,
	"REMOVE":            BulkEditZprOperationDetailsOperationTypeRemove,
}

var mappingBulkEditZprOperationDetailsOperationTypeEnumLowerCase = map[string]BulkEditZprOperationDetailsOperationTypeEnum{
	"add_where_absent":  BulkEditZprOperationDetailsOperationTypeAddWhereAbsent,
	"set_where_present": BulkEditZprOperationDetailsOperationTypeSetWherePresent,
	"add_or_set":        BulkEditZprOperationDetailsOperationTypeAddOrSet,
	"remove":            BulkEditZprOperationDetailsOperationTypeRemove,
}

// GetBulkEditZprOperationDetailsOperationTypeEnumValues Enumerates the set of values for BulkEditZprOperationDetailsOperationTypeEnum
func GetBulkEditZprOperationDetailsOperationTypeEnumValues() []BulkEditZprOperationDetailsOperationTypeEnum {
	values := make([]BulkEditZprOperationDetailsOperationTypeEnum, 0)
	for _, v := range mappingBulkEditZprOperationDetailsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBulkEditZprOperationDetailsOperationTypeEnumStringValues Enumerates the set of values in String for BulkEditZprOperationDetailsOperationTypeEnum
func GetBulkEditZprOperationDetailsOperationTypeEnumStringValues() []string {
	return []string{
		"ADD_WHERE_ABSENT",
		"SET_WHERE_PRESENT",
		"ADD_OR_SET",
		"REMOVE",
	}
}

// GetMappingBulkEditZprOperationDetailsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBulkEditZprOperationDetailsOperationTypeEnum(val string) (BulkEditZprOperationDetailsOperationTypeEnum, bool) {
	enum, ok := mappingBulkEditZprOperationDetailsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
