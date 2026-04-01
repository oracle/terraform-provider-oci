// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmPowerManagement OLVM Power management definitions
type OlvmPowerManagement struct {

	// Address of power management
	Address *string `mandatory:"false" json:"address"`

	// Supported sources of random number generator.
	Agents []OlvmAgent `mandatory:"false" json:"agents"`

	// Toggles the automated power control of the host in order to save energy.
	IsAutomaticPmEnabled *bool `mandatory:"false" json:"isAutomaticPmEnabled"`

	// Indicates whether power management configuration is enabled or disabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Toggles whether to determine if kdump is running on the host before it is shut down.
	IsKDumpDetection *bool `mandatory:"false" json:"isKDumpDetection"`

	// Determines the power management proxy.
	PmProxies []OlvmPmProxy `mandatory:"false" json:"pmProxies"`

	// Determines the power status of the host.
	Status OlvmPowerManagementStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Fencing device code.
	Type *string `mandatory:"false" json:"type"`

	// A valid user name for power management.
	Username *string `mandatory:"false" json:"username"`
}

func (m OlvmPowerManagement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmPowerManagement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmPowerManagementStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOlvmPowerManagementStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmPowerManagementStatusEnum Enum with underlying type: string
type OlvmPowerManagementStatusEnum string

// Set of constants representing the allowable values for OlvmPowerManagementStatusEnum
const (
	OlvmPowerManagementStatusOff     OlvmPowerManagementStatusEnum = "OFF"
	OlvmPowerManagementStatusOn      OlvmPowerManagementStatusEnum = "ON"
	OlvmPowerManagementStatusUnknown OlvmPowerManagementStatusEnum = "UNKNOWN"
)

var mappingOlvmPowerManagementStatusEnum = map[string]OlvmPowerManagementStatusEnum{
	"OFF":     OlvmPowerManagementStatusOff,
	"ON":      OlvmPowerManagementStatusOn,
	"UNKNOWN": OlvmPowerManagementStatusUnknown,
}

var mappingOlvmPowerManagementStatusEnumLowerCase = map[string]OlvmPowerManagementStatusEnum{
	"off":     OlvmPowerManagementStatusOff,
	"on":      OlvmPowerManagementStatusOn,
	"unknown": OlvmPowerManagementStatusUnknown,
}

// GetOlvmPowerManagementStatusEnumValues Enumerates the set of values for OlvmPowerManagementStatusEnum
func GetOlvmPowerManagementStatusEnumValues() []OlvmPowerManagementStatusEnum {
	values := make([]OlvmPowerManagementStatusEnum, 0)
	for _, v := range mappingOlvmPowerManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmPowerManagementStatusEnumStringValues Enumerates the set of values in String for OlvmPowerManagementStatusEnum
func GetOlvmPowerManagementStatusEnumStringValues() []string {
	return []string{
		"OFF",
		"ON",
		"UNKNOWN",
	}
}

// GetMappingOlvmPowerManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmPowerManagementStatusEnum(val string) (OlvmPowerManagementStatusEnum, bool) {
	enum, ok := mappingOlvmPowerManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
