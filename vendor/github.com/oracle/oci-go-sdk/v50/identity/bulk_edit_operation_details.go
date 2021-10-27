// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BulkEditOperationDetails The representation of BulkEditOperationDetails
type BulkEditOperationDetails struct {

	// An enum-like description of the type of operation.
	// * `ADD_WHERE_ABSENT` adds a defined tag only if the tag does not already exist on the resource.
	// * `SET_WHERE_PRESENT` updates the value for a defined tag only if the tag is present on the resource.
	// * `ADD_OR_SET` combines the first two operations to add a defined tag if it does not already exist on the resource
	// or update the value for a defined tag only if the tag is present on the resource.
	// * `REMOVE` removes the defined tag from the resource. The tag is removed from the resource regardless of the tag value.
	OperationType BulkEditOperationDetailsOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
}

func (m BulkEditOperationDetails) String() string {
	return common.PointerString(m)
}

// BulkEditOperationDetailsOperationTypeEnum Enum with underlying type: string
type BulkEditOperationDetailsOperationTypeEnum string

// Set of constants representing the allowable values for BulkEditOperationDetailsOperationTypeEnum
const (
	BulkEditOperationDetailsOperationTypeAddWhereAbsent  BulkEditOperationDetailsOperationTypeEnum = "ADD_WHERE_ABSENT"
	BulkEditOperationDetailsOperationTypeSetWherePresent BulkEditOperationDetailsOperationTypeEnum = "SET_WHERE_PRESENT"
	BulkEditOperationDetailsOperationTypeAddOrSet        BulkEditOperationDetailsOperationTypeEnum = "ADD_OR_SET"
	BulkEditOperationDetailsOperationTypeRemove          BulkEditOperationDetailsOperationTypeEnum = "REMOVE"
)

var mappingBulkEditOperationDetailsOperationType = map[string]BulkEditOperationDetailsOperationTypeEnum{
	"ADD_WHERE_ABSENT":  BulkEditOperationDetailsOperationTypeAddWhereAbsent,
	"SET_WHERE_PRESENT": BulkEditOperationDetailsOperationTypeSetWherePresent,
	"ADD_OR_SET":        BulkEditOperationDetailsOperationTypeAddOrSet,
	"REMOVE":            BulkEditOperationDetailsOperationTypeRemove,
}

// GetBulkEditOperationDetailsOperationTypeEnumValues Enumerates the set of values for BulkEditOperationDetailsOperationTypeEnum
func GetBulkEditOperationDetailsOperationTypeEnumValues() []BulkEditOperationDetailsOperationTypeEnum {
	values := make([]BulkEditOperationDetailsOperationTypeEnum, 0)
	for _, v := range mappingBulkEditOperationDetailsOperationType {
		values = append(values, v)
	}
	return values
}
