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

// DefaultConfigParams Default database configuration.
type DefaultConfigParams struct {

	// The configuration variable name.
	ConfigKey *string `mandatory:"true" json:"configKey"`

	// Default value for the variable.
	DefaultConfigValue *string `mandatory:"true" json:"defaultConfigValue"`

	// Range or list of allowed values.
	AllowedValues *string `mandatory:"true" json:"allowedValues"`

	// If true, modifying this configuration value will require a restart.
	IsRestartRequired *bool `mandatory:"true" json:"isRestartRequired"`

	// Data type of the variable.
	DataType *string `mandatory:"true" json:"dataType"`

	// Whether the value can be overridden or not.
	IsOverridable *bool `mandatory:"true" json:"isOverridable"`

	// Details about the PostgreSQL variable.
	Description *string `mandatory:"true" json:"description"`
}

func (m DefaultConfigParams) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultConfigParams) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
