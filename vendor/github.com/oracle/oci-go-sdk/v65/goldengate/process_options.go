// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProcessOptions Required pipeline options to configure the replication process (Extract or Replicat).
type ProcessOptions struct {
	InitialDataLoad *InitialDataLoad `mandatory:"true" json:"initialDataLoad"`

	ReplicateSchemaChange *ReplicateSchemaChange `mandatory:"true" json:"replicateSchemaChange"`

	// If ENABLED, then the replication process restarts itself upon failure. This option applies when creating or updating a pipeline.
	ShouldRestartOnFailure ProcessOptionsShouldRestartOnFailureEnum `mandatory:"true" json:"shouldRestartOnFailure"`

	// If ENABLED, then the pipeline is started as part of pipeline creation. It uses default mapping. This option applies when creating or updating a pipeline.
	StartUsingDefaultMapping ProcessOptionsStartUsingDefaultMappingEnum `mandatory:"false" json:"startUsingDefaultMapping,omitempty"`
}

func (m ProcessOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProcessOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProcessOptionsShouldRestartOnFailureEnum(string(m.ShouldRestartOnFailure)); !ok && m.ShouldRestartOnFailure != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShouldRestartOnFailure: %s. Supported values are: %s.", m.ShouldRestartOnFailure, strings.Join(GetProcessOptionsShouldRestartOnFailureEnumStringValues(), ",")))
	}

	if _, ok := GetMappingProcessOptionsStartUsingDefaultMappingEnum(string(m.StartUsingDefaultMapping)); !ok && m.StartUsingDefaultMapping != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StartUsingDefaultMapping: %s. Supported values are: %s.", m.StartUsingDefaultMapping, strings.Join(GetProcessOptionsStartUsingDefaultMappingEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProcessOptionsShouldRestartOnFailureEnum Enum with underlying type: string
type ProcessOptionsShouldRestartOnFailureEnum string

// Set of constants representing the allowable values for ProcessOptionsShouldRestartOnFailureEnum
const (
	ProcessOptionsShouldRestartOnFailureEnabled  ProcessOptionsShouldRestartOnFailureEnum = "ENABLED"
	ProcessOptionsShouldRestartOnFailureDisabled ProcessOptionsShouldRestartOnFailureEnum = "DISABLED"
)

var mappingProcessOptionsShouldRestartOnFailureEnum = map[string]ProcessOptionsShouldRestartOnFailureEnum{
	"ENABLED":  ProcessOptionsShouldRestartOnFailureEnabled,
	"DISABLED": ProcessOptionsShouldRestartOnFailureDisabled,
}

var mappingProcessOptionsShouldRestartOnFailureEnumLowerCase = map[string]ProcessOptionsShouldRestartOnFailureEnum{
	"enabled":  ProcessOptionsShouldRestartOnFailureEnabled,
	"disabled": ProcessOptionsShouldRestartOnFailureDisabled,
}

// GetProcessOptionsShouldRestartOnFailureEnumValues Enumerates the set of values for ProcessOptionsShouldRestartOnFailureEnum
func GetProcessOptionsShouldRestartOnFailureEnumValues() []ProcessOptionsShouldRestartOnFailureEnum {
	values := make([]ProcessOptionsShouldRestartOnFailureEnum, 0)
	for _, v := range mappingProcessOptionsShouldRestartOnFailureEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessOptionsShouldRestartOnFailureEnumStringValues Enumerates the set of values in String for ProcessOptionsShouldRestartOnFailureEnum
func GetProcessOptionsShouldRestartOnFailureEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingProcessOptionsShouldRestartOnFailureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessOptionsShouldRestartOnFailureEnum(val string) (ProcessOptionsShouldRestartOnFailureEnum, bool) {
	enum, ok := mappingProcessOptionsShouldRestartOnFailureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ProcessOptionsStartUsingDefaultMappingEnum Enum with underlying type: string
type ProcessOptionsStartUsingDefaultMappingEnum string

// Set of constants representing the allowable values for ProcessOptionsStartUsingDefaultMappingEnum
const (
	ProcessOptionsStartUsingDefaultMappingEnabled  ProcessOptionsStartUsingDefaultMappingEnum = "ENABLED"
	ProcessOptionsStartUsingDefaultMappingDisabled ProcessOptionsStartUsingDefaultMappingEnum = "DISABLED"
)

var mappingProcessOptionsStartUsingDefaultMappingEnum = map[string]ProcessOptionsStartUsingDefaultMappingEnum{
	"ENABLED":  ProcessOptionsStartUsingDefaultMappingEnabled,
	"DISABLED": ProcessOptionsStartUsingDefaultMappingDisabled,
}

var mappingProcessOptionsStartUsingDefaultMappingEnumLowerCase = map[string]ProcessOptionsStartUsingDefaultMappingEnum{
	"enabled":  ProcessOptionsStartUsingDefaultMappingEnabled,
	"disabled": ProcessOptionsStartUsingDefaultMappingDisabled,
}

// GetProcessOptionsStartUsingDefaultMappingEnumValues Enumerates the set of values for ProcessOptionsStartUsingDefaultMappingEnum
func GetProcessOptionsStartUsingDefaultMappingEnumValues() []ProcessOptionsStartUsingDefaultMappingEnum {
	values := make([]ProcessOptionsStartUsingDefaultMappingEnum, 0)
	for _, v := range mappingProcessOptionsStartUsingDefaultMappingEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessOptionsStartUsingDefaultMappingEnumStringValues Enumerates the set of values in String for ProcessOptionsStartUsingDefaultMappingEnum
func GetProcessOptionsStartUsingDefaultMappingEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingProcessOptionsStartUsingDefaultMappingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessOptionsStartUsingDefaultMappingEnum(val string) (ProcessOptionsStartUsingDefaultMappingEnum, bool) {
	enum, ok := mappingProcessOptionsStartUsingDefaultMappingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
