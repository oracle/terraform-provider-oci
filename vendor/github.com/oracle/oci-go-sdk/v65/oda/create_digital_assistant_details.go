// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDigitalAssistantDetails Properties that are required to create a Digital Assistant.
type CreateDigitalAssistantDetails interface {

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

type createdigitalassistantdetails struct {
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
func (m *createdigitalassistantdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedigitalassistantdetails createdigitalassistantdetails
	s := struct {
		Model Unmarshalercreatedigitalassistantdetails
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
func (m *createdigitalassistantdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "VERSION":
		mm := CreateDigitalAssistantVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLONE":
		mm := CloneDigitalAssistantDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NEW":
		mm := CreateNewDigitalAssistantDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTEND":
		mm := ExtendDigitalAssistantDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDigitalAssistantDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetCategory returns Category
func (m createdigitalassistantdetails) GetCategory() *string {
	return m.Category
}

// GetDescription returns Description
func (m createdigitalassistantdetails) GetDescription() *string {
	return m.Description
}

// GetPlatformVersion returns PlatformVersion
func (m createdigitalassistantdetails) GetPlatformVersion() *string {
	return m.PlatformVersion
}

// GetMultilingualMode returns MultilingualMode
func (m createdigitalassistantdetails) GetMultilingualMode() BotMultilingualModeEnum {
	return m.MultilingualMode
}

// GetPrimaryLanguageTag returns PrimaryLanguageTag
func (m createdigitalassistantdetails) GetPrimaryLanguageTag() *string {
	return m.PrimaryLanguageTag
}

// GetFreeformTags returns FreeformTags
func (m createdigitalassistantdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createdigitalassistantdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createdigitalassistantdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdigitalassistantdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBotMultilingualModeEnum(string(m.MultilingualMode)); !ok && m.MultilingualMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MultilingualMode: %s. Supported values are: %s.", m.MultilingualMode, strings.Join(GetBotMultilingualModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
