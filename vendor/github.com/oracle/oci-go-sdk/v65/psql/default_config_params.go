// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefaultConfigParams Default DB Configuration
type DefaultConfigParams struct {

	// Key is the configuration key.
	ConfigKey *string `mandatory:"true" json:"configKey"`

	// Default value
	DefaultConfigValue *string `mandatory:"true" json:"defaultConfigValue"`

	// Range or list of allowed values
	AllowedValues *string `mandatory:"true" json:"allowedValues"`

	// If true, modfying this configuration value will requires restart.
	IsRestartRequired *bool `mandatory:"true" json:"isRestartRequired"`

	// Describes about the Datatype value.
	DataType *string `mandatory:"true" json:"dataType"`

	// This flags tells whether the value is overridable or not.
	IsOverridable *bool `mandatory:"true" json:"isOverridable"`

	// Details about the Postgresql params.
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
