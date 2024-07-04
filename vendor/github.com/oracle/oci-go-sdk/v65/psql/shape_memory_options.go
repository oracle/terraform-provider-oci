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

// ShapeMemoryOptions Options for the the shape memory
type ShapeMemoryOptions struct {

	// Default per OCPU configuration in GBs
	DefaultPerOcpuInGBs *int `mandatory:"true" json:"defaultPerOcpuInGBs"`

	// Minimum Memory configuration in GBs
	MinInGBs *int `mandatory:"true" json:"minInGBs"`

	// Minimum Memory configuration per OCPU in GBs
	MinPerOcpuInGBs *int `mandatory:"true" json:"minPerOcpuInGBs"`

	// Maximum Memory configuration in GBs
	MaxInGBs *int `mandatory:"true" json:"maxInGBs"`

	// Maximum Memory configuration per OCPU in GBs
	MaxPerOcpuInGBs *int `mandatory:"true" json:"maxPerOcpuInGBs"`
}

func (m ShapeMemoryOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeMemoryOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
