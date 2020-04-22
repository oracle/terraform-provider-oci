// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Index Information about an index.
type Index struct {

	// Index name.
	Name *string `mandatory:"false" json:"name"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the table to which this index belongs.
	TableName *string `mandatory:"false" json:"tableName"`

	// the OCID of the table to which this index belongs.
	TableId *string `mandatory:"false" json:"tableId"`

	// A set of keys for a secondary index.
	Keys []IndexKey `mandatory:"false" json:"keys"`

	// The state of an index.
	LifecycleState IndexLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m Index) String() string {
	return common.PointerString(m)
}

// IndexLifecycleStateEnum Enum with underlying type: string
type IndexLifecycleStateEnum string

// Set of constants representing the allowable values for IndexLifecycleStateEnum
const (
	IndexLifecycleStateCreating IndexLifecycleStateEnum = "CREATING"
	IndexLifecycleStateUpdating IndexLifecycleStateEnum = "UPDATING"
	IndexLifecycleStateActive   IndexLifecycleStateEnum = "ACTIVE"
	IndexLifecycleStateDeleting IndexLifecycleStateEnum = "DELETING"
	IndexLifecycleStateDeleted  IndexLifecycleStateEnum = "DELETED"
	IndexLifecycleStateFailed   IndexLifecycleStateEnum = "FAILED"
)

var mappingIndexLifecycleState = map[string]IndexLifecycleStateEnum{
	"CREATING": IndexLifecycleStateCreating,
	"UPDATING": IndexLifecycleStateUpdating,
	"ACTIVE":   IndexLifecycleStateActive,
	"DELETING": IndexLifecycleStateDeleting,
	"DELETED":  IndexLifecycleStateDeleted,
	"FAILED":   IndexLifecycleStateFailed,
}

// GetIndexLifecycleStateEnumValues Enumerates the set of values for IndexLifecycleStateEnum
func GetIndexLifecycleStateEnumValues() []IndexLifecycleStateEnum {
	values := make([]IndexLifecycleStateEnum, 0)
	for _, v := range mappingIndexLifecycleState {
		values = append(values, v)
	}
	return values
}
