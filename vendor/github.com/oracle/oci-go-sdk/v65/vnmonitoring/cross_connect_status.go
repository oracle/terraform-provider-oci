// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CrossConnectStatus The status of the cross-connect.
type CrossConnectStatus struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.
	CrossConnectId *string `mandatory:"true" json:"crossConnectId"`

	// Indicates whether Oracle's side of the interface is up or down.
	InterfaceState CrossConnectStatusInterfaceStateEnum `mandatory:"false" json:"interfaceState,omitempty"`

	// The light level of the cross-connect (in dBm).
	// Example: `14.0`
	LightLevelIndBm *float32 `mandatory:"false" json:"lightLevelIndBm"`

	// Status indicator corresponding to the light level.
	//   * **NO_LIGHT:** No measurable light
	//   * **LOW_WARN:** There's measurable light but it's too low
	//   * **HIGH_WARN:** Light level is too high
	//   * **BAD:** There's measurable light but the signal-to-noise ratio is bad
	//   * **GOOD:** Good light level
	LightLevelIndicator CrossConnectStatusLightLevelIndicatorEnum `mandatory:"false" json:"lightLevelIndicator,omitempty"`
}

func (m CrossConnectStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CrossConnectStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCrossConnectStatusInterfaceStateEnum(string(m.InterfaceState)); !ok && m.InterfaceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InterfaceState: %s. Supported values are: %s.", m.InterfaceState, strings.Join(GetCrossConnectStatusInterfaceStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCrossConnectStatusLightLevelIndicatorEnum(string(m.LightLevelIndicator)); !ok && m.LightLevelIndicator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LightLevelIndicator: %s. Supported values are: %s.", m.LightLevelIndicator, strings.Join(GetCrossConnectStatusLightLevelIndicatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CrossConnectStatusInterfaceStateEnum Enum with underlying type: string
type CrossConnectStatusInterfaceStateEnum string

// Set of constants representing the allowable values for CrossConnectStatusInterfaceStateEnum
const (
	CrossConnectStatusInterfaceStateUp   CrossConnectStatusInterfaceStateEnum = "UP"
	CrossConnectStatusInterfaceStateDown CrossConnectStatusInterfaceStateEnum = "DOWN"
)

var mappingCrossConnectStatusInterfaceStateEnum = map[string]CrossConnectStatusInterfaceStateEnum{
	"UP":   CrossConnectStatusInterfaceStateUp,
	"DOWN": CrossConnectStatusInterfaceStateDown,
}

var mappingCrossConnectStatusInterfaceStateEnumLowerCase = map[string]CrossConnectStatusInterfaceStateEnum{
	"up":   CrossConnectStatusInterfaceStateUp,
	"down": CrossConnectStatusInterfaceStateDown,
}

// GetCrossConnectStatusInterfaceStateEnumValues Enumerates the set of values for CrossConnectStatusInterfaceStateEnum
func GetCrossConnectStatusInterfaceStateEnumValues() []CrossConnectStatusInterfaceStateEnum {
	values := make([]CrossConnectStatusInterfaceStateEnum, 0)
	for _, v := range mappingCrossConnectStatusInterfaceStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCrossConnectStatusInterfaceStateEnumStringValues Enumerates the set of values in String for CrossConnectStatusInterfaceStateEnum
func GetCrossConnectStatusInterfaceStateEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
	}
}

// GetMappingCrossConnectStatusInterfaceStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCrossConnectStatusInterfaceStateEnum(val string) (CrossConnectStatusInterfaceStateEnum, bool) {
	enum, ok := mappingCrossConnectStatusInterfaceStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CrossConnectStatusLightLevelIndicatorEnum Enum with underlying type: string
type CrossConnectStatusLightLevelIndicatorEnum string

// Set of constants representing the allowable values for CrossConnectStatusLightLevelIndicatorEnum
const (
	CrossConnectStatusLightLevelIndicatorNoLight  CrossConnectStatusLightLevelIndicatorEnum = "NO_LIGHT"
	CrossConnectStatusLightLevelIndicatorLowWarn  CrossConnectStatusLightLevelIndicatorEnum = "LOW_WARN"
	CrossConnectStatusLightLevelIndicatorHighWarn CrossConnectStatusLightLevelIndicatorEnum = "HIGH_WARN"
	CrossConnectStatusLightLevelIndicatorBad      CrossConnectStatusLightLevelIndicatorEnum = "BAD"
	CrossConnectStatusLightLevelIndicatorGood     CrossConnectStatusLightLevelIndicatorEnum = "GOOD"
)

var mappingCrossConnectStatusLightLevelIndicatorEnum = map[string]CrossConnectStatusLightLevelIndicatorEnum{
	"NO_LIGHT":  CrossConnectStatusLightLevelIndicatorNoLight,
	"LOW_WARN":  CrossConnectStatusLightLevelIndicatorLowWarn,
	"HIGH_WARN": CrossConnectStatusLightLevelIndicatorHighWarn,
	"BAD":       CrossConnectStatusLightLevelIndicatorBad,
	"GOOD":      CrossConnectStatusLightLevelIndicatorGood,
}

var mappingCrossConnectStatusLightLevelIndicatorEnumLowerCase = map[string]CrossConnectStatusLightLevelIndicatorEnum{
	"no_light":  CrossConnectStatusLightLevelIndicatorNoLight,
	"low_warn":  CrossConnectStatusLightLevelIndicatorLowWarn,
	"high_warn": CrossConnectStatusLightLevelIndicatorHighWarn,
	"bad":       CrossConnectStatusLightLevelIndicatorBad,
	"good":      CrossConnectStatusLightLevelIndicatorGood,
}

// GetCrossConnectStatusLightLevelIndicatorEnumValues Enumerates the set of values for CrossConnectStatusLightLevelIndicatorEnum
func GetCrossConnectStatusLightLevelIndicatorEnumValues() []CrossConnectStatusLightLevelIndicatorEnum {
	values := make([]CrossConnectStatusLightLevelIndicatorEnum, 0)
	for _, v := range mappingCrossConnectStatusLightLevelIndicatorEnum {
		values = append(values, v)
	}
	return values
}

// GetCrossConnectStatusLightLevelIndicatorEnumStringValues Enumerates the set of values in String for CrossConnectStatusLightLevelIndicatorEnum
func GetCrossConnectStatusLightLevelIndicatorEnumStringValues() []string {
	return []string{
		"NO_LIGHT",
		"LOW_WARN",
		"HIGH_WARN",
		"BAD",
		"GOOD",
	}
}

// GetMappingCrossConnectStatusLightLevelIndicatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCrossConnectStatusLightLevelIndicatorEnum(val string) (CrossConnectStatusLightLevelIndicatorEnum, bool) {
	enum, ok := mappingCrossConnectStatusLightLevelIndicatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
