// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShapeSummary The shape of the DB System. The shape determines resources to allocate
// to the DB System - CPU cores and memory for VM shapes; CPU cores, memory
// and storage for non-VM (or bare metal) shapes.  For a description of
// shapes, see DB System Shape Options (https://docs.cloud.oracle.com/mysql-database/doc/shapes.htm).
type ShapeSummary struct {

	// The name of the shape used for the DB System.
	Name *string `mandatory:"true" json:"name"`

	// The number of CPU Cores the Instance provides. These are "OCPU"s.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The amount of RAM the Instance provides. This is an IEC base-2 number.
	MemorySizeInGBs *int `mandatory:"true" json:"memorySizeInGBs"`

	// What service features the shape is supported for.
	IsSupportedFor []ShapeSummaryIsSupportedForEnum `mandatory:"false" json:"isSupportedFor,omitempty"`
}

func (m ShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IsSupportedFor {
		if _, ok := GetMappingShapeSummaryIsSupportedForEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsSupportedFor: %s. Supported values are: %s.", val, strings.Join(GetShapeSummaryIsSupportedForEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShapeSummaryIsSupportedForEnum Enum with underlying type: string
type ShapeSummaryIsSupportedForEnum string

// Set of constants representing the allowable values for ShapeSummaryIsSupportedForEnum
const (
	ShapeSummaryIsSupportedForDbsystem        ShapeSummaryIsSupportedForEnum = "DBSYSTEM"
	ShapeSummaryIsSupportedForHeatwavecluster ShapeSummaryIsSupportedForEnum = "HEATWAVECLUSTER"
)

var mappingShapeSummaryIsSupportedForEnum = map[string]ShapeSummaryIsSupportedForEnum{
	"DBSYSTEM":        ShapeSummaryIsSupportedForDbsystem,
	"HEATWAVECLUSTER": ShapeSummaryIsSupportedForHeatwavecluster,
}

var mappingShapeSummaryIsSupportedForEnumLowerCase = map[string]ShapeSummaryIsSupportedForEnum{
	"dbsystem":        ShapeSummaryIsSupportedForDbsystem,
	"heatwavecluster": ShapeSummaryIsSupportedForHeatwavecluster,
}

// GetShapeSummaryIsSupportedForEnumValues Enumerates the set of values for ShapeSummaryIsSupportedForEnum
func GetShapeSummaryIsSupportedForEnumValues() []ShapeSummaryIsSupportedForEnum {
	values := make([]ShapeSummaryIsSupportedForEnum, 0)
	for _, v := range mappingShapeSummaryIsSupportedForEnum {
		values = append(values, v)
	}
	return values
}

// GetShapeSummaryIsSupportedForEnumStringValues Enumerates the set of values in String for ShapeSummaryIsSupportedForEnum
func GetShapeSummaryIsSupportedForEnumStringValues() []string {
	return []string{
		"DBSYSTEM",
		"HEATWAVECLUSTER",
	}
}

// GetMappingShapeSummaryIsSupportedForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShapeSummaryIsSupportedForEnum(val string) (ShapeSummaryIsSupportedForEnum, bool) {
	enum, ok := mappingShapeSummaryIsSupportedForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
