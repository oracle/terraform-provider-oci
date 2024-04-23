// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LifecycleStageSummary Provides summary information for a lifecycle stage.
type LifecycleStageSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle stage.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the lifecycle stage.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User-specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages within the lifecycle environment.
	Rank *int `mandatory:"true" json:"rank"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment that contains the lifecycle stage.
	LifecycleEnvironmentId *string `mandatory:"false" json:"lifecycleEnvironmentId"`

	// The user-friendly name for the lifecycle environment. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	LifecycleEnvironmentDisplayName *string `mandatory:"false" json:"lifecycleEnvironmentDisplayName"`

	// The operating system of the managed instances in the lifecycle stage.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The CPU architecture of the managed instances in the lifecycle stage.
	ArchType ArchTypeEnum `mandatory:"false" json:"archType,omitempty"`

	// The vendor of the operating system used by the managed instances in the lifecycle stage.
	VendorName VendorNameEnum `mandatory:"false" json:"vendorName,omitempty"`

	// The list of managed instances associated with the lifecycle stage.
	ManagedInstances *int `mandatory:"false" json:"managedInstances"`

	// The location of managed instances associated with the lifecycle stage.
	Location ManagedInstanceLocationEnum `mandatory:"false" json:"location,omitempty"`

	SoftwareSourceId *SoftwareSourceDetails `mandatory:"false" json:"softwareSourceId"`

	// The time the lifecycle stage was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the lifecycle stage was last modified (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeModified *common.SDKTime `mandatory:"false" json:"timeModified"`

	// The current state of the lifecycle environment.
	LifecycleState LifecycleStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LifecycleStageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LifecycleStageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVendorNameEnum(string(m.VendorName)); !ok && m.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", m.VendorName, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagedInstanceLocationEnum(string(m.Location)); !ok && m.Location != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", m.Location, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
