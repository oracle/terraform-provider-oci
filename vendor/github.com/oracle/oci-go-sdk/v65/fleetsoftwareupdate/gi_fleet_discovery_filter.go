// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GiFleetDiscoveryFilter Possible Discovery filters.
type GiFleetDiscoveryFilter interface {

	// INCLUDE or EXCLUDE the filter results in the discovery for GI targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	GetMode() GiFleetDiscoveryFilterModeEnum
}

type gifleetdiscoveryfilter struct {
	JsonData []byte
	Mode     GiFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
	Type     string                         `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *gifleetdiscoveryfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergifleetdiscoveryfilter gifleetdiscoveryfilter
	s := struct {
		Model Unmarshalergifleetdiscoveryfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Mode = s.Model.Mode
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *gifleetdiscoveryfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFINED_TAG":
		mm := GiDefinedTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_ID":
		mm := GiCompartmentIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FREEFORM_TAG":
		mm := GiFreeformTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_ID":
		mm := GiResourceIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSION":
		mm := GiVersionFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for GiFleetDiscoveryFilter: %s.", m.Type)
		return *m, nil
	}
}

// GetMode returns Mode
func (m gifleetdiscoveryfilter) GetMode() GiFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m gifleetdiscoveryfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m gifleetdiscoveryfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGiFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGiFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GiFleetDiscoveryFilterModeEnum Enum with underlying type: string
type GiFleetDiscoveryFilterModeEnum string

// Set of constants representing the allowable values for GiFleetDiscoveryFilterModeEnum
const (
	GiFleetDiscoveryFilterModeInclude GiFleetDiscoveryFilterModeEnum = "INCLUDE"
	GiFleetDiscoveryFilterModeExclude GiFleetDiscoveryFilterModeEnum = "EXCLUDE"
)

var mappingGiFleetDiscoveryFilterModeEnum = map[string]GiFleetDiscoveryFilterModeEnum{
	"INCLUDE": GiFleetDiscoveryFilterModeInclude,
	"EXCLUDE": GiFleetDiscoveryFilterModeExclude,
}

var mappingGiFleetDiscoveryFilterModeEnumLowerCase = map[string]GiFleetDiscoveryFilterModeEnum{
	"include": GiFleetDiscoveryFilterModeInclude,
	"exclude": GiFleetDiscoveryFilterModeExclude,
}

// GetGiFleetDiscoveryFilterModeEnumValues Enumerates the set of values for GiFleetDiscoveryFilterModeEnum
func GetGiFleetDiscoveryFilterModeEnumValues() []GiFleetDiscoveryFilterModeEnum {
	values := make([]GiFleetDiscoveryFilterModeEnum, 0)
	for _, v := range mappingGiFleetDiscoveryFilterModeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiFleetDiscoveryFilterModeEnumStringValues Enumerates the set of values in String for GiFleetDiscoveryFilterModeEnum
func GetGiFleetDiscoveryFilterModeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingGiFleetDiscoveryFilterModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiFleetDiscoveryFilterModeEnum(val string) (GiFleetDiscoveryFilterModeEnum, bool) {
	enum, ok := mappingGiFleetDiscoveryFilterModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GiFleetDiscoveryFilterTypeEnum Enum with underlying type: string
type GiFleetDiscoveryFilterTypeEnum string

// Set of constants representing the allowable values for GiFleetDiscoveryFilterTypeEnum
const (
	GiFleetDiscoveryFilterTypeCompartmentId GiFleetDiscoveryFilterTypeEnum = "COMPARTMENT_ID"
	GiFleetDiscoveryFilterTypeVersion       GiFleetDiscoveryFilterTypeEnum = "VERSION"
	GiFleetDiscoveryFilterTypeFreeformTag   GiFleetDiscoveryFilterTypeEnum = "FREEFORM_TAG"
	GiFleetDiscoveryFilterTypeDefinedTag    GiFleetDiscoveryFilterTypeEnum = "DEFINED_TAG"
	GiFleetDiscoveryFilterTypeResourceId    GiFleetDiscoveryFilterTypeEnum = "RESOURCE_ID"
)

var mappingGiFleetDiscoveryFilterTypeEnum = map[string]GiFleetDiscoveryFilterTypeEnum{
	"COMPARTMENT_ID": GiFleetDiscoveryFilterTypeCompartmentId,
	"VERSION":        GiFleetDiscoveryFilterTypeVersion,
	"FREEFORM_TAG":   GiFleetDiscoveryFilterTypeFreeformTag,
	"DEFINED_TAG":    GiFleetDiscoveryFilterTypeDefinedTag,
	"RESOURCE_ID":    GiFleetDiscoveryFilterTypeResourceId,
}

var mappingGiFleetDiscoveryFilterTypeEnumLowerCase = map[string]GiFleetDiscoveryFilterTypeEnum{
	"compartment_id": GiFleetDiscoveryFilterTypeCompartmentId,
	"version":        GiFleetDiscoveryFilterTypeVersion,
	"freeform_tag":   GiFleetDiscoveryFilterTypeFreeformTag,
	"defined_tag":    GiFleetDiscoveryFilterTypeDefinedTag,
	"resource_id":    GiFleetDiscoveryFilterTypeResourceId,
}

// GetGiFleetDiscoveryFilterTypeEnumValues Enumerates the set of values for GiFleetDiscoveryFilterTypeEnum
func GetGiFleetDiscoveryFilterTypeEnumValues() []GiFleetDiscoveryFilterTypeEnum {
	values := make([]GiFleetDiscoveryFilterTypeEnum, 0)
	for _, v := range mappingGiFleetDiscoveryFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiFleetDiscoveryFilterTypeEnumStringValues Enumerates the set of values in String for GiFleetDiscoveryFilterTypeEnum
func GetGiFleetDiscoveryFilterTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT_ID",
		"VERSION",
		"FREEFORM_TAG",
		"DEFINED_TAG",
		"RESOURCE_ID",
	}
}

// GetMappingGiFleetDiscoveryFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiFleetDiscoveryFilterTypeEnum(val string) (GiFleetDiscoveryFilterTypeEnum, bool) {
	enum, ok := mappingGiFleetDiscoveryFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
