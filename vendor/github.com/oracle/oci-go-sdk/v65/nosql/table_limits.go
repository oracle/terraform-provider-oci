// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TableLimits Throughput and storage limits configuration of a table.
type TableLimits struct {

	// Maximum sustained read throughput limit for the table.
	MaxReadUnits *int `mandatory:"true" json:"maxReadUnits"`

	// Maximum sustained write throughput limit for the table.
	MaxWriteUnits *int `mandatory:"true" json:"maxWriteUnits"`

	// Maximum size of storage used by the table.
	MaxStorageInGBs *int `mandatory:"true" json:"maxStorageInGBs"`

	// The capacity mode of the table.  If capacityMode = ON_DEMAND,
	// maxReadUnits and maxWriteUnits are not used, and both will have
	// the value of zero.
	CapacityMode TableLimitsCapacityModeEnum `mandatory:"false" json:"capacityMode,omitempty"`
}

func (m TableLimits) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableLimits) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTableLimitsCapacityModeEnum(string(m.CapacityMode)); !ok && m.CapacityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CapacityMode: %s. Supported values are: %s.", m.CapacityMode, strings.Join(GetTableLimitsCapacityModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TableLimitsCapacityModeEnum Enum with underlying type: string
type TableLimitsCapacityModeEnum string

// Set of constants representing the allowable values for TableLimitsCapacityModeEnum
const (
	TableLimitsCapacityModeProvisioned TableLimitsCapacityModeEnum = "PROVISIONED"
	TableLimitsCapacityModeOnDemand    TableLimitsCapacityModeEnum = "ON_DEMAND"
)

var mappingTableLimitsCapacityModeEnum = map[string]TableLimitsCapacityModeEnum{
	"PROVISIONED": TableLimitsCapacityModeProvisioned,
	"ON_DEMAND":   TableLimitsCapacityModeOnDemand,
}

var mappingTableLimitsCapacityModeEnumLowerCase = map[string]TableLimitsCapacityModeEnum{
	"provisioned": TableLimitsCapacityModeProvisioned,
	"on_demand":   TableLimitsCapacityModeOnDemand,
}

// GetTableLimitsCapacityModeEnumValues Enumerates the set of values for TableLimitsCapacityModeEnum
func GetTableLimitsCapacityModeEnumValues() []TableLimitsCapacityModeEnum {
	values := make([]TableLimitsCapacityModeEnum, 0)
	for _, v := range mappingTableLimitsCapacityModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTableLimitsCapacityModeEnumStringValues Enumerates the set of values in String for TableLimitsCapacityModeEnum
func GetTableLimitsCapacityModeEnumStringValues() []string {
	return []string{
		"PROVISIONED",
		"ON_DEMAND",
	}
}

// GetMappingTableLimitsCapacityModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTableLimitsCapacityModeEnum(val string) (TableLimitsCapacityModeEnum, bool) {
	enum, ok := mappingTableLimitsCapacityModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
