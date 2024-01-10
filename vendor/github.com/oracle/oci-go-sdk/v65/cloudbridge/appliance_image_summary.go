// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplianceImageSummary Description of the ApplianceImage.
type ApplianceImageSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of the appliance Image file.
	FileName *string `mandatory:"true" json:"fileName"`

	// The name of the image to be displayed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The version of the image file.
	Version *string `mandatory:"true" json:"version"`

	// The size of the image file in megabytes.
	SizeInMBs *string `mandatory:"true" json:"sizeInMBs"`

	// The checksum of the image file.
	Checksum *string `mandatory:"true" json:"checksum"`

	// The virtualization platform that the image file supports.
	Platform *string `mandatory:"true" json:"platform"`

	// The file format of the image file.
	Format *string `mandatory:"true" json:"format"`

	// The time when the appliance image was created.An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the appliance image was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The URL from which the appliance image can be downloaded.
	DownloadUrl *string `mandatory:"true" json:"downloadUrl"`

	// The current state of the appliance image.
	LifecycleState ApplianceImageSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ApplianceImageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplianceImageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplianceImageSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplianceImageSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplianceImageSummaryLifecycleStateEnum Enum with underlying type: string
type ApplianceImageSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ApplianceImageSummaryLifecycleStateEnum
const (
	ApplianceImageSummaryLifecycleStateCreating ApplianceImageSummaryLifecycleStateEnum = "CREATING"
	ApplianceImageSummaryLifecycleStateUpdating ApplianceImageSummaryLifecycleStateEnum = "UPDATING"
	ApplianceImageSummaryLifecycleStateActive   ApplianceImageSummaryLifecycleStateEnum = "ACTIVE"
	ApplianceImageSummaryLifecycleStateDeleting ApplianceImageSummaryLifecycleStateEnum = "DELETING"
	ApplianceImageSummaryLifecycleStateDeleted  ApplianceImageSummaryLifecycleStateEnum = "DELETED"
	ApplianceImageSummaryLifecycleStateFailed   ApplianceImageSummaryLifecycleStateEnum = "FAILED"
)

var mappingApplianceImageSummaryLifecycleStateEnum = map[string]ApplianceImageSummaryLifecycleStateEnum{
	"CREATING": ApplianceImageSummaryLifecycleStateCreating,
	"UPDATING": ApplianceImageSummaryLifecycleStateUpdating,
	"ACTIVE":   ApplianceImageSummaryLifecycleStateActive,
	"DELETING": ApplianceImageSummaryLifecycleStateDeleting,
	"DELETED":  ApplianceImageSummaryLifecycleStateDeleted,
	"FAILED":   ApplianceImageSummaryLifecycleStateFailed,
}

var mappingApplianceImageSummaryLifecycleStateEnumLowerCase = map[string]ApplianceImageSummaryLifecycleStateEnum{
	"creating": ApplianceImageSummaryLifecycleStateCreating,
	"updating": ApplianceImageSummaryLifecycleStateUpdating,
	"active":   ApplianceImageSummaryLifecycleStateActive,
	"deleting": ApplianceImageSummaryLifecycleStateDeleting,
	"deleted":  ApplianceImageSummaryLifecycleStateDeleted,
	"failed":   ApplianceImageSummaryLifecycleStateFailed,
}

// GetApplianceImageSummaryLifecycleStateEnumValues Enumerates the set of values for ApplianceImageSummaryLifecycleStateEnum
func GetApplianceImageSummaryLifecycleStateEnumValues() []ApplianceImageSummaryLifecycleStateEnum {
	values := make([]ApplianceImageSummaryLifecycleStateEnum, 0)
	for _, v := range mappingApplianceImageSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplianceImageSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ApplianceImageSummaryLifecycleStateEnum
func GetApplianceImageSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApplianceImageSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplianceImageSummaryLifecycleStateEnum(val string) (ApplianceImageSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingApplianceImageSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
