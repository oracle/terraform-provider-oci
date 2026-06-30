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

// PitrDetails Point-in-time recovery details
type PitrDetails struct {

	// The current state of the point-in-time recovery of the db system.
	PitrState PitrDetailsPitrStateEnum `mandatory:"true" json:"pitrState"`

	// List of point-in-time recovery windows.
	RecoveryTimeWindows []PitrTimeWindow `mandatory:"false" json:"recoveryTimeWindows"`
}

func (m PitrDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PitrDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPitrDetailsPitrStateEnum(string(m.PitrState)); !ok && m.PitrState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PitrState: %s. Supported values are: %s.", m.PitrState, strings.Join(GetPitrDetailsPitrStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PitrDetailsPitrStateEnum Enum with underlying type: string
type PitrDetailsPitrStateEnum string

// Set of constants representing the allowable values for PitrDetailsPitrStateEnum
const (
	PitrDetailsPitrStateActive   PitrDetailsPitrStateEnum = "ACTIVE"
	PitrDetailsPitrStateUpdating PitrDetailsPitrStateEnum = "UPDATING"
	PitrDetailsPitrStateInactive PitrDetailsPitrStateEnum = "INACTIVE"
)

var mappingPitrDetailsPitrStateEnum = map[string]PitrDetailsPitrStateEnum{
	"ACTIVE":   PitrDetailsPitrStateActive,
	"UPDATING": PitrDetailsPitrStateUpdating,
	"INACTIVE": PitrDetailsPitrStateInactive,
}

var mappingPitrDetailsPitrStateEnumLowerCase = map[string]PitrDetailsPitrStateEnum{
	"active":   PitrDetailsPitrStateActive,
	"updating": PitrDetailsPitrStateUpdating,
	"inactive": PitrDetailsPitrStateInactive,
}

// GetPitrDetailsPitrStateEnumValues Enumerates the set of values for PitrDetailsPitrStateEnum
func GetPitrDetailsPitrStateEnumValues() []PitrDetailsPitrStateEnum {
	values := make([]PitrDetailsPitrStateEnum, 0)
	for _, v := range mappingPitrDetailsPitrStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPitrDetailsPitrStateEnumStringValues Enumerates the set of values in String for PitrDetailsPitrStateEnum
func GetPitrDetailsPitrStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
	}
}

// GetMappingPitrDetailsPitrStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPitrDetailsPitrStateEnum(val string) (PitrDetailsPitrStateEnum, bool) {
	enum, ok := mappingPitrDetailsPitrStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
