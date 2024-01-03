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

// CreateSkillDetails Properties that are required to create a Skill.
type CreateSkillDetails interface {

	// The resource's category.  This is used to group resource's together.
	GetCategory() *string

	// A short description of the resource.
	GetDescription() *string

	// The ODA Platform Version for this resource.
	GetPlatformVersion() *string

	// The multilingual mode for the resource.
	GetMultilingualMode() BotMultilingualModeEnum

	// The primary language for the resource.
	GetPrimaryLanguageTag() *string

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createskilldetails struct {
	JsonData           []byte
	Category           *string                           `mandatory:"false" json:"category"`
	Description        *string                           `mandatory:"false" json:"description"`
	PlatformVersion    *string                           `mandatory:"false" json:"platformVersion"`
	MultilingualMode   BotMultilingualModeEnum           `mandatory:"false" json:"multilingualMode,omitempty"`
	PrimaryLanguageTag *string                           `mandatory:"false" json:"primaryLanguageTag"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Kind               string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *createskilldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateskilldetails createskilldetails
	s := struct {
		Model Unmarshalercreateskilldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Category = s.Model.Category
	m.Description = s.Model.Description
	m.PlatformVersion = s.Model.PlatformVersion
	m.MultilingualMode = s.Model.MultilingualMode
	m.PrimaryLanguageTag = s.Model.PrimaryLanguageTag
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createskilldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "CLONE":
		mm := CloneSkillDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NEW":
		mm := CreateNewSkillDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSION":
		mm := CreateSkillVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTEND":
		mm := ExtendSkillDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateSkillDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetCategory returns Category
func (m createskilldetails) GetCategory() *string {
	return m.Category
}

// GetDescription returns Description
func (m createskilldetails) GetDescription() *string {
	return m.Description
}

// GetPlatformVersion returns PlatformVersion
func (m createskilldetails) GetPlatformVersion() *string {
	return m.PlatformVersion
}

// GetMultilingualMode returns MultilingualMode
func (m createskilldetails) GetMultilingualMode() BotMultilingualModeEnum {
	return m.MultilingualMode
}

// GetPrimaryLanguageTag returns PrimaryLanguageTag
func (m createskilldetails) GetPrimaryLanguageTag() *string {
	return m.PrimaryLanguageTag
}

// GetFreeformTags returns FreeformTags
func (m createskilldetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createskilldetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createskilldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createskilldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBotMultilingualModeEnum(string(m.MultilingualMode)); !ok && m.MultilingualMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MultilingualMode: %s. Supported values are: %s.", m.MultilingualMode, strings.Join(GetBotMultilingualModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
