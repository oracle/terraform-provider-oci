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

// CreateConfigurationDetails The information to create a new Configuration.
type CreateConfigurationDetails struct {

	// configuration display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Compute Shape Name like VM.Standard3.Flex.
	Shape *string `mandatory:"true" json:"shape"`

	// Version of the Postgresql DB
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// CPU cpuCoreCount. Min value is 1. Max value depends on the shape.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	DbConfigurationOverrides *DbConfigurationOverrideCollection `mandatory:"true" json:"dbConfigurationOverrides"`

	// Details about the Configuration Set.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CreateConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
