// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GranularityCapability Describes time granularity behavior for time-series Insight.
type GranularityCapability struct {

	// Granularity selection strategy.
	Type GranularityCapabilityTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Minimum supported granularity in seconds.
	MinSeconds *int `mandatory:"false" json:"minSeconds"`

	// Maximum supported granularity in seconds.
	MaxSeconds *int `mandatory:"false" json:"maxSeconds"`
}

func (m GranularityCapability) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GranularityCapability) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGranularityCapabilityTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGranularityCapabilityTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GranularityCapabilityTypeEnum Enum with underlying type: string
type GranularityCapabilityTypeEnum string

// Set of constants representing the allowable values for GranularityCapabilityTypeEnum
const (
	GranularityCapabilityTypeFixed             GranularityCapabilityTypeEnum = "FIXED"
	GranularityCapabilityTypeClientInput       GranularityCapabilityTypeEnum = "CLIENT_INPUT"
	GranularityCapabilityTypeBackendDetermined GranularityCapabilityTypeEnum = "BACKEND_DETERMINED"
)

var mappingGranularityCapabilityTypeEnum = map[string]GranularityCapabilityTypeEnum{
	"FIXED":              GranularityCapabilityTypeFixed,
	"CLIENT_INPUT":       GranularityCapabilityTypeClientInput,
	"BACKEND_DETERMINED": GranularityCapabilityTypeBackendDetermined,
}

var mappingGranularityCapabilityTypeEnumLowerCase = map[string]GranularityCapabilityTypeEnum{
	"fixed":              GranularityCapabilityTypeFixed,
	"client_input":       GranularityCapabilityTypeClientInput,
	"backend_determined": GranularityCapabilityTypeBackendDetermined,
}

// GetGranularityCapabilityTypeEnumValues Enumerates the set of values for GranularityCapabilityTypeEnum
func GetGranularityCapabilityTypeEnumValues() []GranularityCapabilityTypeEnum {
	values := make([]GranularityCapabilityTypeEnum, 0)
	for _, v := range mappingGranularityCapabilityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGranularityCapabilityTypeEnumStringValues Enumerates the set of values in String for GranularityCapabilityTypeEnum
func GetGranularityCapabilityTypeEnumStringValues() []string {
	return []string{
		"FIXED",
		"CLIENT_INPUT",
		"BACKEND_DETERMINED",
	}
}

// GetMappingGranularityCapabilityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGranularityCapabilityTypeEnum(val string) (GranularityCapabilityTypeEnum, bool) {
	enum, ok := mappingGranularityCapabilityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
