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

// UpdateRowDetails Specifications for the putting of a table row.
type UpdateRowDetails struct {

	// The map of values from a row.
	Value map[string]interface{} `mandatory:"true" json:"value"`

	// The OCID of the table's compartment.  Required
	// if the tableNameOrId path parameter is a table name.
	// Optional if tableNameOrId is an OCID.  If tableNameOrId
	// is an OCID, and compartmentId is supplied, the latter
	// must match the identified table's compartmentId.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Specifies a condition for the put operation.
	Option UpdateRowDetailsOptionEnum `mandatory:"false" json:"option,omitempty"`

	// If true, and the put fails due to an option setting, then
	// the existing row will be returned.
	IsGetReturnRow *bool `mandatory:"false" json:"isGetReturnRow"`

	// Timeout setting for the put.
	TimeoutInMs *int `mandatory:"false" json:"timeoutInMs"`

	// Time-to-live for the row, in days.
	Ttl *int `mandatory:"false" json:"ttl"`

	// If true, set time-to-live for this row to the table's default.
	IsTtlUseTableDefault *bool `mandatory:"false" json:"isTtlUseTableDefault"`

	// Sets the number of generated identity values that are
	// requested from the server during a put. If present and greater than 0,
	// this value takes precedence over a default value for the table.
	IdentityCacheSize *int `mandatory:"false" json:"identityCacheSize"`

	// If present and true, the presented row value must exactly
	// match the table's schema.  Otherwise, rows with missing
	// non-key fields or extra fields can be written successfully.
	IsExactMatch *bool `mandatory:"false" json:"isExactMatch"`
}

func (m UpdateRowDetails) String() string {
	return common.PointerString(m)
}

// UpdateRowDetailsOptionEnum Enum with underlying type: string
type UpdateRowDetailsOptionEnum string

// Set of constants representing the allowable values for UpdateRowDetailsOptionEnum
const (
	UpdateRowDetailsOptionAbsent  UpdateRowDetailsOptionEnum = "IF_ABSENT"
	UpdateRowDetailsOptionPresent UpdateRowDetailsOptionEnum = "IF_PRESENT"
)

var mappingUpdateRowDetailsOption = map[string]UpdateRowDetailsOptionEnum{
	"IF_ABSENT":  UpdateRowDetailsOptionAbsent,
	"IF_PRESENT": UpdateRowDetailsOptionPresent,
}

// GetUpdateRowDetailsOptionEnumValues Enumerates the set of values for UpdateRowDetailsOptionEnum
func GetUpdateRowDetailsOptionEnumValues() []UpdateRowDetailsOptionEnum {
	values := make([]UpdateRowDetailsOptionEnum, 0)
	for _, v := range mappingUpdateRowDetailsOption {
		values = append(values, v)
	}
	return values
}
