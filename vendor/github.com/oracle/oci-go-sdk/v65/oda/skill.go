// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Skill Skill metadata.
type Skill struct {

	// Unique immutable identifier that was assigned when the resource was created.
	Id *string `mandatory:"true" json:"id"`

	// The reource's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"true" json:"name"`

	// The resource's version. The version can only contain numbers, letters, periods, underscores, dashes or spaces.  The version must begin with a letter or a number.
	Version *string `mandatory:"true" json:"version"`

	// The resource's display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The resource's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The resource's publish state.
	LifecycleDetails BotPublishStateEnum `mandatory:"true" json:"lifecycleDetails"`

	// The ODA Platform Version for this resource.
	PlatformVersion *string `mandatory:"true" json:"platformVersion"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The resource's category.  This is used to group resource's together.
	Category *string `mandatory:"false" json:"category"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The resource's namespace.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The unique identifier for the base reource (when this resource extends another).
	BaseId *string `mandatory:"false" json:"baseId"`

	// The multilingual mode for the resource.
	MultilingualMode BotMultilingualModeEnum `mandatory:"false" json:"multilingualMode,omitempty"`

	// The primary language for the resource.
	PrimaryLanguageTag *string `mandatory:"false" json:"primaryLanguageTag"`

	// A list of native languages supported by this resource.
	NativeLanguageTags []string `mandatory:"false" json:"nativeLanguageTags"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Skill) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Skill) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBotPublishStateEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetBotPublishStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBotMultilingualModeEnum(string(m.MultilingualMode)); !ok && m.MultilingualMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MultilingualMode: %s. Supported values are: %s.", m.MultilingualMode, strings.Join(GetBotMultilingualModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
