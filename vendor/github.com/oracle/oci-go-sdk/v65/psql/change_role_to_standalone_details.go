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

// ChangeRoleToStandaloneDetails The information required to change a replica database system's role to standalone.
type ChangeRoleToStandaloneDetails struct {

	// Type of the mode choose during change role operation.
	// REPLAY_PENDING_UPDATES (Default value): In this mode, the role change is delayed until replica database system has processed all Write-Ahead log (WAL) records that were archived before this API call is made.
	// IMMEDIATELY: In this mode, the role change is applied right away, without waiting for any pending WAL records to be processed. This allows for an immediate transition.
	ChangeMode ChangeRoleToStandaloneDetailsChangeModeEnum `mandatory:"true" json:"changeMode"`
}

func (m ChangeRoleToStandaloneDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeRoleToStandaloneDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChangeRoleToStandaloneDetailsChangeModeEnum(string(m.ChangeMode)); !ok && m.ChangeMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChangeMode: %s. Supported values are: %s.", m.ChangeMode, strings.Join(GetChangeRoleToStandaloneDetailsChangeModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeRoleToStandaloneDetailsChangeModeEnum Enum with underlying type: string
type ChangeRoleToStandaloneDetailsChangeModeEnum string

// Set of constants representing the allowable values for ChangeRoleToStandaloneDetailsChangeModeEnum
const (
	ChangeRoleToStandaloneDetailsChangeModeReplayPendingUpdates ChangeRoleToStandaloneDetailsChangeModeEnum = "REPLAY_PENDING_UPDATES"
	ChangeRoleToStandaloneDetailsChangeModeImmediately          ChangeRoleToStandaloneDetailsChangeModeEnum = "IMMEDIATELY"
)

var mappingChangeRoleToStandaloneDetailsChangeModeEnum = map[string]ChangeRoleToStandaloneDetailsChangeModeEnum{
	"REPLAY_PENDING_UPDATES": ChangeRoleToStandaloneDetailsChangeModeReplayPendingUpdates,
	"IMMEDIATELY":            ChangeRoleToStandaloneDetailsChangeModeImmediately,
}

var mappingChangeRoleToStandaloneDetailsChangeModeEnumLowerCase = map[string]ChangeRoleToStandaloneDetailsChangeModeEnum{
	"replay_pending_updates": ChangeRoleToStandaloneDetailsChangeModeReplayPendingUpdates,
	"immediately":            ChangeRoleToStandaloneDetailsChangeModeImmediately,
}

// GetChangeRoleToStandaloneDetailsChangeModeEnumValues Enumerates the set of values for ChangeRoleToStandaloneDetailsChangeModeEnum
func GetChangeRoleToStandaloneDetailsChangeModeEnumValues() []ChangeRoleToStandaloneDetailsChangeModeEnum {
	values := make([]ChangeRoleToStandaloneDetailsChangeModeEnum, 0)
	for _, v := range mappingChangeRoleToStandaloneDetailsChangeModeEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeRoleToStandaloneDetailsChangeModeEnumStringValues Enumerates the set of values in String for ChangeRoleToStandaloneDetailsChangeModeEnum
func GetChangeRoleToStandaloneDetailsChangeModeEnumStringValues() []string {
	return []string{
		"REPLAY_PENDING_UPDATES",
		"IMMEDIATELY",
	}
}

// GetMappingChangeRoleToStandaloneDetailsChangeModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeRoleToStandaloneDetailsChangeModeEnum(val string) (ChangeRoleToStandaloneDetailsChangeModeEnum, bool) {
	enum, ok := mappingChangeRoleToStandaloneDetailsChangeModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
