// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Index) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIndexLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIndexLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingIndexLifecycleStateEnum = map[string]IndexLifecycleStateEnum{
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
	for _, v := range mappingIndexLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIndexLifecycleStateEnumStringValues Enumerates the set of values in String for IndexLifecycleStateEnum
func GetIndexLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingIndexLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIndexLifecycleStateEnum(val string) (IndexLifecycleStateEnum, bool) {
	mappingIndexLifecycleStateEnumIgnoreCase := make(map[string]IndexLifecycleStateEnum)
	for k, v := range mappingIndexLifecycleStateEnum {
		mappingIndexLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIndexLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
