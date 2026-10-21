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

// UpgradeDbSystemDetails Upgrade details for the database system.
type UpgradeDbSystemDetails struct {

	// The target major version of the database system software.
	// Example: "17"
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration for target database
	// version to be associated with database system.
	// Optional, if database system is currently running with default configuration object.
	// Mandatory, if database system is currently running with custom configuration object.
	ConfigId *string `mandatory:"false" json:"configId"`

	// Optional post-upgrade operations to run. Verification and vacuum are skipped unless
	// their corresponding options are included. If omitted or empty, both operations are skipped.
	Options []UpgradeDbSystemDetailsOptionsEnum `mandatory:"false" json:"options,omitempty"`
}

func (m UpgradeDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpgradeDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Options {
		if _, ok := GetMappingUpgradeDbSystemDetailsOptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Options: %s. Supported values are: %s.", val, strings.Join(GetUpgradeDbSystemDetailsOptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpgradeDbSystemDetailsOptionsEnum Enum with underlying type: string
type UpgradeDbSystemDetailsOptionsEnum string

// Set of constants representing the allowable values for UpgradeDbSystemDetailsOptionsEnum
const (
	UpgradeDbSystemDetailsOptionsVerification UpgradeDbSystemDetailsOptionsEnum = "RUN_POST_UPGRADE_VERIFICATION"
	UpgradeDbSystemDetailsOptionsVacuum       UpgradeDbSystemDetailsOptionsEnum = "RUN_POST_UPGRADE_VACUUM"
)

var mappingUpgradeDbSystemDetailsOptionsEnum = map[string]UpgradeDbSystemDetailsOptionsEnum{
	"RUN_POST_UPGRADE_VERIFICATION": UpgradeDbSystemDetailsOptionsVerification,
	"RUN_POST_UPGRADE_VACUUM":       UpgradeDbSystemDetailsOptionsVacuum,
}

var mappingUpgradeDbSystemDetailsOptionsEnumLowerCase = map[string]UpgradeDbSystemDetailsOptionsEnum{
	"run_post_upgrade_verification": UpgradeDbSystemDetailsOptionsVerification,
	"run_post_upgrade_vacuum":       UpgradeDbSystemDetailsOptionsVacuum,
}

// GetUpgradeDbSystemDetailsOptionsEnumValues Enumerates the set of values for UpgradeDbSystemDetailsOptionsEnum
func GetUpgradeDbSystemDetailsOptionsEnumValues() []UpgradeDbSystemDetailsOptionsEnum {
	values := make([]UpgradeDbSystemDetailsOptionsEnum, 0)
	for _, v := range mappingUpgradeDbSystemDetailsOptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpgradeDbSystemDetailsOptionsEnumStringValues Enumerates the set of values in String for UpgradeDbSystemDetailsOptionsEnum
func GetUpgradeDbSystemDetailsOptionsEnumStringValues() []string {
	return []string{
		"RUN_POST_UPGRADE_VERIFICATION",
		"RUN_POST_UPGRADE_VACUUM",
	}
}

// GetMappingUpgradeDbSystemDetailsOptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpgradeDbSystemDetailsOptionsEnum(val string) (UpgradeDbSystemDetailsOptionsEnum, bool) {
	enum, ok := mappingUpgradeDbSystemDetailsOptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
