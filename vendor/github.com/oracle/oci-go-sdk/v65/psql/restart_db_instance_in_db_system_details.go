// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// RestartDbInstanceInDbSystemDetails Database instance node restart parameters.
type RestartDbInstanceInDbSystemDetails struct {

	// A unique identifier for the database instance, or node.
	DbInstanceId *string `mandatory:"true" json:"dbInstanceId"`

	// The restart type for the database instance.
	RestartType RestartDbInstanceInDbSystemDetailsRestartTypeEnum `mandatory:"true" json:"restartType"`
}

func (m RestartDbInstanceInDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestartDbInstanceInDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRestartDbInstanceInDbSystemDetailsRestartTypeEnum(string(m.RestartType)); !ok && m.RestartType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RestartType: %s. Supported values are: %s.", m.RestartType, strings.Join(GetRestartDbInstanceInDbSystemDetailsRestartTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RestartDbInstanceInDbSystemDetailsRestartTypeEnum Enum with underlying type: string
type RestartDbInstanceInDbSystemDetailsRestartTypeEnum string

// Set of constants representing the allowable values for RestartDbInstanceInDbSystemDetailsRestartTypeEnum
const (
	RestartDbInstanceInDbSystemDetailsRestartTypeNormal     RestartDbInstanceInDbSystemDetailsRestartTypeEnum = "NORMAL"
	RestartDbInstanceInDbSystemDetailsRestartTypeNodeReboot RestartDbInstanceInDbSystemDetailsRestartTypeEnum = "NODE_REBOOT"
)

var mappingRestartDbInstanceInDbSystemDetailsRestartTypeEnum = map[string]RestartDbInstanceInDbSystemDetailsRestartTypeEnum{
	"NORMAL":      RestartDbInstanceInDbSystemDetailsRestartTypeNormal,
	"NODE_REBOOT": RestartDbInstanceInDbSystemDetailsRestartTypeNodeReboot,
}

var mappingRestartDbInstanceInDbSystemDetailsRestartTypeEnumLowerCase = map[string]RestartDbInstanceInDbSystemDetailsRestartTypeEnum{
	"normal":      RestartDbInstanceInDbSystemDetailsRestartTypeNormal,
	"node_reboot": RestartDbInstanceInDbSystemDetailsRestartTypeNodeReboot,
}

// GetRestartDbInstanceInDbSystemDetailsRestartTypeEnumValues Enumerates the set of values for RestartDbInstanceInDbSystemDetailsRestartTypeEnum
func GetRestartDbInstanceInDbSystemDetailsRestartTypeEnumValues() []RestartDbInstanceInDbSystemDetailsRestartTypeEnum {
	values := make([]RestartDbInstanceInDbSystemDetailsRestartTypeEnum, 0)
	for _, v := range mappingRestartDbInstanceInDbSystemDetailsRestartTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRestartDbInstanceInDbSystemDetailsRestartTypeEnumStringValues Enumerates the set of values in String for RestartDbInstanceInDbSystemDetailsRestartTypeEnum
func GetRestartDbInstanceInDbSystemDetailsRestartTypeEnumStringValues() []string {
	return []string{
		"NORMAL",
		"NODE_REBOOT",
	}
}

// GetMappingRestartDbInstanceInDbSystemDetailsRestartTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRestartDbInstanceInDbSystemDetailsRestartTypeEnum(val string) (RestartDbInstanceInDbSystemDetailsRestartTypeEnum, bool) {
	enum, ok := mappingRestartDbInstanceInDbSystemDetailsRestartTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
