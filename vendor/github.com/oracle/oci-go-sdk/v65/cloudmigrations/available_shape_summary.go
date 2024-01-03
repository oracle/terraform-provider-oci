// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AvailableShapeSummary Sumarized information about a shape.
type AvailableShapeSummary struct {

	// Availability domain of the shape.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// Shape name and availability domain.  Used for pagination.
	PaginationToken *string `mandatory:"true" json:"paginationToken"`

	// Name of the shape.
	Shape *string `mandatory:"true" json:"shape"`

	// Description of the processor.
	ProcessorDescription *string `mandatory:"true" json:"processorDescription"`

	// Number of CPUs.
	Ocpus *float32 `mandatory:"true" json:"ocpus"`

	// Amount of memory for the shape.
	MemoryInGBs *float32 `mandatory:"true" json:"memoryInGBs"`

	// Minimum CPUs required.
	MinTotalBaselineOcpusRequired *float32 `mandatory:"false" json:"minTotalBaselineOcpusRequired"`

	// Shape bandwidth.
	NetworkingBandwidthInGbps *float32 `mandatory:"false" json:"networkingBandwidthInGbps"`

	// Maximum number of virtual network interfaces that can be attached.
	MaxVnicAttachments *int `mandatory:"false" json:"maxVnicAttachments"`

	// Number of GPUs.
	Gpus *int `mandatory:"false" json:"gpus"`

	// Description of the GPUs.
	GpuDescription *string `mandatory:"false" json:"gpuDescription"`

	// Number of local disks.
	LocalDisks *int `mandatory:"false" json:"localDisks"`

	// Total size of local disks for shape.
	LocalDisksTotalSizeInGBs *float32 `mandatory:"false" json:"localDisksTotalSizeInGBs"`

	// Description of local disks.
	LocalDiskDescription *string `mandatory:"false" json:"localDiskDescription"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AvailableShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailableShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
