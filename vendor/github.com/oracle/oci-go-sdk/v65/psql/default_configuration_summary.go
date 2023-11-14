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

// DefaultConfigurationSummary Summary of the Configuration.
type DefaultConfigurationSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Configuration identifier name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the Configuration was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Configuration.
	LifecycleState DefaultConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Compute Shape Name like VM.Standard3.Flex.
	Shape *string `mandatory:"true" json:"shape"`

	// Version of the Postgresql DB
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// CPU cpuCoreCount. Min value is 1. Max value depends on the shape.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape.
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
