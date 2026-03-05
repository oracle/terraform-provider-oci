// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SwitchoverDeploymentPeerDetails The information about switching to deployment peer.
type SwitchoverDeploymentPeerDetails struct {

	// The availability domain of a placement.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The fault domain of a placement.
	FaultDomain *string `mandatory:"true" json:"faultDomain"`

	// Specifies the switchover mode. SWITCHOVER performs a planned role transition to the standby; FAILOVER performs a forced promotion when the primary is not available.
	Mode SwitchoverDeploymentPeerDetailsModeEnum `mandatory:"false" json:"mode,omitempty"`
}

func (m SwitchoverDeploymentPeerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SwitchoverDeploymentPeerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSwitchoverDeploymentPeerDetailsModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetSwitchoverDeploymentPeerDetailsModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SwitchoverDeploymentPeerDetailsModeEnum Enum with underlying type: string
type SwitchoverDeploymentPeerDetailsModeEnum string

// Set of constants representing the allowable values for SwitchoverDeploymentPeerDetailsModeEnum
const (
	SwitchoverDeploymentPeerDetailsModeSwitchover SwitchoverDeploymentPeerDetailsModeEnum = "SWITCHOVER"
	SwitchoverDeploymentPeerDetailsModeFailover   SwitchoverDeploymentPeerDetailsModeEnum = "FAILOVER"
)

var mappingSwitchoverDeploymentPeerDetailsModeEnum = map[string]SwitchoverDeploymentPeerDetailsModeEnum{
	"SWITCHOVER": SwitchoverDeploymentPeerDetailsModeSwitchover,
	"FAILOVER":   SwitchoverDeploymentPeerDetailsModeFailover,
}

var mappingSwitchoverDeploymentPeerDetailsModeEnumLowerCase = map[string]SwitchoverDeploymentPeerDetailsModeEnum{
	"switchover": SwitchoverDeploymentPeerDetailsModeSwitchover,
	"failover":   SwitchoverDeploymentPeerDetailsModeFailover,
}

// GetSwitchoverDeploymentPeerDetailsModeEnumValues Enumerates the set of values for SwitchoverDeploymentPeerDetailsModeEnum
func GetSwitchoverDeploymentPeerDetailsModeEnumValues() []SwitchoverDeploymentPeerDetailsModeEnum {
	values := make([]SwitchoverDeploymentPeerDetailsModeEnum, 0)
	for _, v := range mappingSwitchoverDeploymentPeerDetailsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSwitchoverDeploymentPeerDetailsModeEnumStringValues Enumerates the set of values in String for SwitchoverDeploymentPeerDetailsModeEnum
func GetSwitchoverDeploymentPeerDetailsModeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"FAILOVER",
	}
}

// GetMappingSwitchoverDeploymentPeerDetailsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSwitchoverDeploymentPeerDetailsModeEnum(val string) (SwitchoverDeploymentPeerDetailsModeEnum, bool) {
	enum, ok := mappingSwitchoverDeploymentPeerDetailsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
