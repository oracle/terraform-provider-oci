// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCccProvisionedPackageDetails Updates the Compute Cloud@Customer marketplace provisioned package details. A provisioned package
// cannot be updated on the Compute Cloud@Customer infrastructure. Only metadata for the provisioned
// package in OCI can be updated.
type UpdateCccProvisionedPackageDetails struct {

	// Compute Cloud@Customer marketplace provisioned package display name.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the Compute Cloud@Customer marketplace provisioned package.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The total number of OCPUs that can be used for the compute instances that use this
	// image on the Compute Cloud@Customer infrastructure.
	// If the value is higher then the total number of OCPUs available,
	// the value will accepted but the maximum will be the total number of OCPUs on the
	// Compute Cloud@Customer infrastructure.
	TotalOcpuLimit *int `mandatory:"false" json:"totalOcpuLimit"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCccProvisionedPackageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCccProvisionedPackageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
