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

// ShapeSummary Summary of the database system shape.
type ShapeSummary struct {

	// The name of the Compute VM shape.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"true" json:"shape"`

	// The number of OCPUs.
	OcpuCount *int `mandatory:"true" json:"ocpuCount"`

	// The amount of memory in gigabytes.
	MemorySizeInGBs *int `mandatory:"true" json:"memorySizeInGBs"`

	// A unique identifier for the shape.
	Id *string `mandatory:"false" json:"id"`

	// Indicates if the shape is a flex shape.
	IsFlexible *bool `mandatory:"false" json:"isFlexible"`

	ShapeOcpuOptions *ShapeOcpuOptions `mandatory:"false" json:"shapeOcpuOptions"`

	ShapeMemoryOptions *ShapeMemoryOptions `mandatory:"false" json:"shapeMemoryOptions"`
}

func (m ShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
