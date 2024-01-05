// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDbConfigParams Configuration for the PostgreSQL database instance.
type UpdateDbConfigParams struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration.
	ConfigId *string `mandatory:"true" json:"configId"`

	// Whether a configuration update requires a restart of the database instance or a reload of the configuration.
	// Some configuration changes require a restart of database instances to be applied.
	ApplyConfig UpdateDbConfigParamsApplyConfigEnum `mandatory:"false" json:"applyConfig,omitempty"`
}

func (m UpdateDbConfigParams) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbConfigParams) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDbConfigParamsApplyConfigEnum(string(m.ApplyConfig)); !ok && m.ApplyConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApplyConfig: %s. Supported values are: %s.", m.ApplyConfig, strings.Join(GetUpdateDbConfigParamsApplyConfigEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDbConfigParamsApplyConfigEnum Enum with underlying type: string
type UpdateDbConfigParamsApplyConfigEnum string

// Set of constants representing the allowable values for UpdateDbConfigParamsApplyConfigEnum
const (
	UpdateDbConfigParamsApplyConfigRestart UpdateDbConfigParamsApplyConfigEnum = "RESTART"
	UpdateDbConfigParamsApplyConfigReload  UpdateDbConfigParamsApplyConfigEnum = "RELOAD"
)

var mappingUpdateDbConfigParamsApplyConfigEnum = map[string]UpdateDbConfigParamsApplyConfigEnum{
	"RESTART": UpdateDbConfigParamsApplyConfigRestart,
	"RELOAD":  UpdateDbConfigParamsApplyConfigReload,
}

var mappingUpdateDbConfigParamsApplyConfigEnumLowerCase = map[string]UpdateDbConfigParamsApplyConfigEnum{
	"restart": UpdateDbConfigParamsApplyConfigRestart,
	"reload":  UpdateDbConfigParamsApplyConfigReload,
}

// GetUpdateDbConfigParamsApplyConfigEnumValues Enumerates the set of values for UpdateDbConfigParamsApplyConfigEnum
func GetUpdateDbConfigParamsApplyConfigEnumValues() []UpdateDbConfigParamsApplyConfigEnum {
	values := make([]UpdateDbConfigParamsApplyConfigEnum, 0)
	for _, v := range mappingUpdateDbConfigParamsApplyConfigEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDbConfigParamsApplyConfigEnumStringValues Enumerates the set of values in String for UpdateDbConfigParamsApplyConfigEnum
func GetUpdateDbConfigParamsApplyConfigEnumStringValues() []string {
	return []string{
		"RESTART",
		"RELOAD",
	}
}

// GetMappingUpdateDbConfigParamsApplyConfigEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDbConfigParamsApplyConfigEnum(val string) (UpdateDbConfigParamsApplyConfigEnum, bool) {
	enum, ok := mappingUpdateDbConfigParamsApplyConfigEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
