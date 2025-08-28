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

// StreamGroup Details for a Stream Group
type StreamGroup struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamGroup.
	Id *string `mandatory:"true" json:"id"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A human-friendly name for the streamGroup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// When the streamGroup was created, as an RFC3339 datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the streamGroup was updated, as an RFC3339 datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the streamGroup.
	LifecycleState StreamGroupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Stream
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// List of streamSource OCIDs associated with the stream group
	StreamSourceIds []string `mandatory:"false" json:"streamSourceIds"`

	// List of streamSource OCIDs where the streamSource overlaps in field of view.
	StreamOverlaps []StreamGroupOverlap `mandatory:"false" json:"streamOverlaps"`

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

func (m StreamGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStreamGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamGroupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamGroupLifecycleStateEnum Enum with underlying type: string
type StreamGroupLifecycleStateEnum string

// Set of constants representing the allowable values for StreamGroupLifecycleStateEnum
const (
	StreamGroupLifecycleStateCreating StreamGroupLifecycleStateEnum = "CREATING"
	StreamGroupLifecycleStateUpdating StreamGroupLifecycleStateEnum = "UPDATING"
	StreamGroupLifecycleStateActive   StreamGroupLifecycleStateEnum = "ACTIVE"
	StreamGroupLifecycleStateDeleting StreamGroupLifecycleStateEnum = "DELETING"
	StreamGroupLifecycleStateDeleted  StreamGroupLifecycleStateEnum = "DELETED"
	StreamGroupLifecycleStateFailed   StreamGroupLifecycleStateEnum = "FAILED"
)

var mappingStreamGroupLifecycleStateEnum = map[string]StreamGroupLifecycleStateEnum{
	"CREATING": StreamGroupLifecycleStateCreating,
	"UPDATING": StreamGroupLifecycleStateUpdating,
	"ACTIVE":   StreamGroupLifecycleStateActive,
	"DELETING": StreamGroupLifecycleStateDeleting,
	"DELETED":  StreamGroupLifecycleStateDeleted,
	"FAILED":   StreamGroupLifecycleStateFailed,
}

var mappingStreamGroupLifecycleStateEnumLowerCase = map[string]StreamGroupLifecycleStateEnum{
	"creating": StreamGroupLifecycleStateCreating,
	"updating": StreamGroupLifecycleStateUpdating,
	"active":   StreamGroupLifecycleStateActive,
	"deleting": StreamGroupLifecycleStateDeleting,
	"deleted":  StreamGroupLifecycleStateDeleted,
	"failed":   StreamGroupLifecycleStateFailed,
}

// GetStreamGroupLifecycleStateEnumValues Enumerates the set of values for StreamGroupLifecycleStateEnum
func GetStreamGroupLifecycleStateEnumValues() []StreamGroupLifecycleStateEnum {
	values := make([]StreamGroupLifecycleStateEnum, 0)
	for _, v := range mappingStreamGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamGroupLifecycleStateEnumStringValues Enumerates the set of values in String for StreamGroupLifecycleStateEnum
func GetStreamGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingStreamGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamGroupLifecycleStateEnum(val string) (StreamGroupLifecycleStateEnum, bool) {
	enum, ok := mappingStreamGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
