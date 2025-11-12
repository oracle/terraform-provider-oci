// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchemaDocument Schema Document representing Schema.yaml (/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager_topic-schema.htm)
type SchemaDocument struct {

	// The version of the schema definition format in use for this document.
	SchemaVersion SchemaDocumentSchemaVersionEnum `mandatory:"true" json:"schemaVersion"`

	// Key-value map of input variables defined for use by the stack.
	Variables map[string]BaseVariable `mandatory:"true" json:"variables"`

	// The display name or title for this schema document.
	Title *string `mandatory:"false" json:"title"`

	// A detailed description of the stack or schema.
	Description *string `mandatory:"false" json:"description"`

	// Additional details describing the stack's purpose or use-case.
	StackDescription *string `mandatory:"false" json:"stackDescription"`

	// The version of the package associated with this schema.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// The version identifier for this schema document.
	Version *string `mandatory:"false" json:"version"`

	// The locale/language for the schema user interface (default is EN).
	Locale SchemaDocumentLocaleEnum `mandatory:"false" json:"locale,omitempty"`

	// logo url.
	LogoUrl *string `mandatory:"false" json:"logoUrl"`

	Source *StackSource `mandatory:"false" json:"source"`

	// Informational text or notes relevant to the stack or its use.
	InformationalText *string `mandatory:"false" json:"informationalText"`

	// Setup or usage instructions for this stack.
	Instructions *string `mandatory:"false" json:"instructions"`

	// Troubleshooting tips, guidance, or steps for stack usage.
	Troubleshooting *string `mandatory:"false" json:"troubleshooting"`

	// Indicates if the stack allows users to view state information.
	CanAllowViewState *bool `mandatory:"false" json:"canAllowViewState"`

	Groupings *VariableGroups `mandatory:"false" json:"groupings"`

	// An array of variable group definitions for organizing variables together.
	VariableGroups []VariableGroup `mandatory:"false" json:"variableGroups"`

	// A mapping of output variable names to their definitions.
	Outputs map[string]BaseOutput `mandatory:"false" json:"outputs"`

	// Array of output group objects to group outputs for display or logical purposes.
	OutputGroups []OutputGroup `mandatory:"false" json:"outputGroups"`

	// primary output button value.
	PrimaryOutputButton *string `mandatory:"false" json:"primaryOutputButton"`
}

func (m SchemaDocument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchemaDocument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchemaDocumentSchemaVersionEnum(string(m.SchemaVersion)); !ok && m.SchemaVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaVersion: %s. Supported values are: %s.", m.SchemaVersion, strings.Join(GetSchemaDocumentSchemaVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSchemaDocumentLocaleEnum(string(m.Locale)); !ok && m.Locale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Locale: %s. Supported values are: %s.", m.Locale, strings.Join(GetSchemaDocumentLocaleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SchemaDocument) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Title               *string                         `json:"title"`
		Description         *string                         `json:"description"`
		StackDescription    *string                         `json:"stackDescription"`
		PackageVersion      *string                         `json:"packageVersion"`
		Version             *string                         `json:"version"`
		Locale              SchemaDocumentLocaleEnum        `json:"locale"`
		LogoUrl             *string                         `json:"logoUrl"`
		Source              *StackSource                    `json:"source"`
		InformationalText   *string                         `json:"informationalText"`
		Instructions        *string                         `json:"instructions"`
		Troubleshooting     *string                         `json:"troubleshooting"`
		CanAllowViewState   *bool                           `json:"canAllowViewState"`
		Groupings           *VariableGroups                 `json:"groupings"`
		VariableGroups      []VariableGroup                 `json:"variableGroups"`
		Outputs             map[string]baseoutput           `json:"outputs"`
		OutputGroups        []OutputGroup                   `json:"outputGroups"`
		PrimaryOutputButton *string                         `json:"primaryOutputButton"`
		SchemaVersion       SchemaDocumentSchemaVersionEnum `json:"schemaVersion"`
		Variables           map[string]basevariable         `json:"variables"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Title = model.Title

	m.Description = model.Description

	m.StackDescription = model.StackDescription

	m.PackageVersion = model.PackageVersion

	m.Version = model.Version

	m.Locale = model.Locale

	m.LogoUrl = model.LogoUrl

	m.Source = model.Source

	m.InformationalText = model.InformationalText

	m.Instructions = model.Instructions

	m.Troubleshooting = model.Troubleshooting

	m.CanAllowViewState = model.CanAllowViewState

	m.Groupings = model.Groupings

	m.VariableGroups = make([]VariableGroup, len(model.VariableGroups))
	copy(m.VariableGroups, model.VariableGroups)
	m.Outputs = make(map[string]BaseOutput)
	for k, v := range model.Outputs {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Outputs[k] = nn.(BaseOutput)
		} else {
			m.Outputs[k] = nil
		}
	}

	m.OutputGroups = make([]OutputGroup, len(model.OutputGroups))
	copy(m.OutputGroups, model.OutputGroups)
	m.PrimaryOutputButton = model.PrimaryOutputButton

	m.SchemaVersion = model.SchemaVersion

	m.Variables = make(map[string]BaseVariable)
	for k, v := range model.Variables {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Variables[k] = nn.(BaseVariable)
		} else {
			m.Variables[k] = nil
		}
	}

	return
}

// SchemaDocumentSchemaVersionEnum Enum with underlying type: string
type SchemaDocumentSchemaVersionEnum string

// Set of constants representing the allowable values for SchemaDocumentSchemaVersionEnum
const (
	SchemaDocumentSchemaVersion00 SchemaDocumentSchemaVersionEnum = "V_1_0_0"
	SchemaDocumentSchemaVersion10 SchemaDocumentSchemaVersionEnum = "V_1_1_0"
)

var mappingSchemaDocumentSchemaVersionEnum = map[string]SchemaDocumentSchemaVersionEnum{
	"V_1_0_0": SchemaDocumentSchemaVersion00,
	"V_1_1_0": SchemaDocumentSchemaVersion10,
}

var mappingSchemaDocumentSchemaVersionEnumLowerCase = map[string]SchemaDocumentSchemaVersionEnum{
	"v_1_0_0": SchemaDocumentSchemaVersion00,
	"v_1_1_0": SchemaDocumentSchemaVersion10,
}

// GetSchemaDocumentSchemaVersionEnumValues Enumerates the set of values for SchemaDocumentSchemaVersionEnum
func GetSchemaDocumentSchemaVersionEnumValues() []SchemaDocumentSchemaVersionEnum {
	values := make([]SchemaDocumentSchemaVersionEnum, 0)
	for _, v := range mappingSchemaDocumentSchemaVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaDocumentSchemaVersionEnumStringValues Enumerates the set of values in String for SchemaDocumentSchemaVersionEnum
func GetSchemaDocumentSchemaVersionEnumStringValues() []string {
	return []string{
		"V_1_0_0",
		"V_1_1_0",
	}
}

// GetMappingSchemaDocumentSchemaVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaDocumentSchemaVersionEnum(val string) (SchemaDocumentSchemaVersionEnum, bool) {
	enum, ok := mappingSchemaDocumentSchemaVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SchemaDocumentLocaleEnum Enum with underlying type: string
type SchemaDocumentLocaleEnum string

// Set of constants representing the allowable values for SchemaDocumentLocaleEnum
const (
	SchemaDocumentLocaleEn SchemaDocumentLocaleEnum = "EN"
)

var mappingSchemaDocumentLocaleEnum = map[string]SchemaDocumentLocaleEnum{
	"EN": SchemaDocumentLocaleEn,
}

var mappingSchemaDocumentLocaleEnumLowerCase = map[string]SchemaDocumentLocaleEnum{
	"en": SchemaDocumentLocaleEn,
}

// GetSchemaDocumentLocaleEnumValues Enumerates the set of values for SchemaDocumentLocaleEnum
func GetSchemaDocumentLocaleEnumValues() []SchemaDocumentLocaleEnum {
	values := make([]SchemaDocumentLocaleEnum, 0)
	for _, v := range mappingSchemaDocumentLocaleEnum {
		values = append(values, v)
	}
	return values
}

// GetSchemaDocumentLocaleEnumStringValues Enumerates the set of values in String for SchemaDocumentLocaleEnum
func GetSchemaDocumentLocaleEnumStringValues() []string {
	return []string{
		"EN",
	}
}

// GetMappingSchemaDocumentLocaleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchemaDocumentLocaleEnum(val string) (SchemaDocumentLocaleEnum, bool) {
	enum, ok := mappingSchemaDocumentLocaleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
