// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VisionPrivateEndpoint Vision private endpoint.
type VisionPrivateEndpoint struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of private endpoint
	Id *string `mandatory:"true" json:"id"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of subnet
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// When the visionPrivateEndpoint was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the visionPrivateEndpoint.
	LifecycleState VisionPrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A human-friendly name for the visionPrivateEndpoint, which can be changed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the visionPrivateEndpoint.
	Description *string `mandatory:"false" json:"description"`

	// When the visionPrivateEndpoint was updated, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail, that can provide actionable information if creation failed.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// For example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m VisionPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VisionPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVisionPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVisionPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VisionPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type VisionPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for VisionPrivateEndpointLifecycleStateEnum
const (
	VisionPrivateEndpointLifecycleStateCreating VisionPrivateEndpointLifecycleStateEnum = "CREATING"
	VisionPrivateEndpointLifecycleStateUpdating VisionPrivateEndpointLifecycleStateEnum = "UPDATING"
	VisionPrivateEndpointLifecycleStateActive   VisionPrivateEndpointLifecycleStateEnum = "ACTIVE"
	VisionPrivateEndpointLifecycleStateDeleting VisionPrivateEndpointLifecycleStateEnum = "DELETING"
	VisionPrivateEndpointLifecycleStateDeleted  VisionPrivateEndpointLifecycleStateEnum = "DELETED"
	VisionPrivateEndpointLifecycleStateFailed   VisionPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingVisionPrivateEndpointLifecycleStateEnum = map[string]VisionPrivateEndpointLifecycleStateEnum{
	"CREATING": VisionPrivateEndpointLifecycleStateCreating,
	"UPDATING": VisionPrivateEndpointLifecycleStateUpdating,
	"ACTIVE":   VisionPrivateEndpointLifecycleStateActive,
	"DELETING": VisionPrivateEndpointLifecycleStateDeleting,
	"DELETED":  VisionPrivateEndpointLifecycleStateDeleted,
	"FAILED":   VisionPrivateEndpointLifecycleStateFailed,
}

var mappingVisionPrivateEndpointLifecycleStateEnumLowerCase = map[string]VisionPrivateEndpointLifecycleStateEnum{
	"creating": VisionPrivateEndpointLifecycleStateCreating,
	"updating": VisionPrivateEndpointLifecycleStateUpdating,
	"active":   VisionPrivateEndpointLifecycleStateActive,
	"deleting": VisionPrivateEndpointLifecycleStateDeleting,
	"deleted":  VisionPrivateEndpointLifecycleStateDeleted,
	"failed":   VisionPrivateEndpointLifecycleStateFailed,
}

// GetVisionPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for VisionPrivateEndpointLifecycleStateEnum
func GetVisionPrivateEndpointLifecycleStateEnumValues() []VisionPrivateEndpointLifecycleStateEnum {
	values := make([]VisionPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingVisionPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVisionPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for VisionPrivateEndpointLifecycleStateEnum
func GetVisionPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVisionPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVisionPrivateEndpointLifecycleStateEnum(val string) (VisionPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingVisionPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
