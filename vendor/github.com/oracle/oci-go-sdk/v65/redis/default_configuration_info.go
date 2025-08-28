// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefaultConfigurationInfo Details of a configuration setting in the OCI Cache Default Config Set.
type DefaultConfigurationInfo struct {

	// The key of the configuration setting.
	ConfigKey *string `mandatory:"true" json:"configKey"`

	// The default value for the configuration setting.
	DefaultConfigValue *string `mandatory:"true" json:"defaultConfigValue"`

	// The data type of the configuration setting.
	DataType *string `mandatory:"true" json:"dataType"`

	// Indicates if the configuration is modifiable.
	IsModifiable *bool `mandatory:"true" json:"isModifiable"`

	// Allowed values for the configuration setting.
	AllowedValues *string `mandatory:"false" json:"allowedValues"`

	// Description of the configuration setting.
	Description *string `mandatory:"false" json:"description"`
}

func (m DefaultConfigurationInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultConfigurationInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
