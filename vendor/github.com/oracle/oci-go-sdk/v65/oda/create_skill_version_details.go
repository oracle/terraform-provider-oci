// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSkillVersionDetails Properties that are required to create a new version of an existing Skill.
type CreateSkillVersionDetails struct {

	// The unique identifier of the Skill to create a new version of.
	Id *string `mandatory:"true" json:"id"`

	// The resource's version. The version can only contain numbers, letters, periods, underscores, dashes or spaces.  The version must begin with a letter or a number.
	Version *string `mandatory:"true" json:"version"`

	// The resource's category.  This is used to group resource's together.
	Category *string `mandatory:"false" json:"category"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The ODA Platform Version for this resource.
	PlatformVersion *string `mandatory:"false" json:"platformVersion"`

	// The resource's dialog version.
	DialogVersion *string `mandatory:"false" json:"dialogVersion"`

	// The primary language for the resource.
	PrimaryLanguageTag *string `mandatory:"false" json:"primaryLanguageTag"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The multilingual mode for the resource.
	MultilingualMode BotMultilingualModeEnum `mandatory:"false" json:"multilingualMode,omitempty"`
}

// GetCategory returns Category
func (m CreateSkillVersionDetails) GetCategory() *string {
	return m.Category
}

// GetDescription returns Description
func (m CreateSkillVersionDetails) GetDescription() *string {
	return m.Description
}

// GetPlatformVersion returns PlatformVersion
func (m CreateSkillVersionDetails) GetPlatformVersion() *string {
	return m.PlatformVersion
}

// GetDialogVersion returns DialogVersion
func (m CreateSkillVersionDetails) GetDialogVersion() *string {
	return m.DialogVersion
}

// GetMultilingualMode returns MultilingualMode
func (m CreateSkillVersionDetails) GetMultilingualMode() BotMultilingualModeEnum {
	return m.MultilingualMode
}

// GetPrimaryLanguageTag returns PrimaryLanguageTag
func (m CreateSkillVersionDetails) GetPrimaryLanguageTag() *string {
	return m.PrimaryLanguageTag
}

// GetFreeformTags returns FreeformTags
func (m CreateSkillVersionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateSkillVersionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSkillVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSkillVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBotMultilingualModeEnum(string(m.MultilingualMode)); !ok && m.MultilingualMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MultilingualMode: %s. Supported values are: %s.", m.MultilingualMode, strings.Join(GetBotMultilingualModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSkillVersionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSkillVersionDetails CreateSkillVersionDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCreateSkillVersionDetails
	}{
		"VERSION",
		(MarshalTypeCreateSkillVersionDetails)(m),
	}

	return json.Marshal(&s)
}
