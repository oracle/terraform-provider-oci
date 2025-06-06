// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateLifecycleEnvironmentDetails Provides the information used to create a lifecycle environment. A lifecycle environment is a user-defined pipeline to deliver curated, versioned content in a prescribed, methodical manner.
type CreateLifecycleEnvironmentDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle environment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the lifecycle environment. Does not have to be unique and you can change the name later. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User-specified list of ranked lifecycle stages used within the lifecycle environment.
	Stages []CreateLifecycleStageDetails `mandatory:"true" json:"stages"`

	// The CPU architecture of the managed instances in the lifecycle environment.
	ArchType ArchTypeEnum `mandatory:"true" json:"archType"`

	// The operating system of the managed instances in the lifecycle environment.
	OsFamily OsFamilyEnum `mandatory:"true" json:"osFamily"`

	// The vendor of the operating system used by the managed instances in the lifecycle environment.
	VendorName VendorNameEnum `mandatory:"true" json:"vendorName"`

	// User-specified information about the lifecycle environment. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The location of managed instances attached to the lifecycle environment. If no location is provided, the default is 'ON_PREMISE.'
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateLifecycleEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLifecycleEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingManagedInstanceLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
