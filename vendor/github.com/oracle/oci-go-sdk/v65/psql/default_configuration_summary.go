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

// DefaultConfigurationSummary Summary of the configuration.
type DefaultConfigurationSummary struct {

	// A unique identifier for the configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time that the configuration was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the configuration.
	LifecycleState DefaultConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the shape for the configuration.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"true" json:"shape"`

	// Version of the PostgreSQL database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// CPU core count. Minimum value is 1.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory size in gigabytes with 1GB increment.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m DefaultConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDefaultConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDefaultConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
