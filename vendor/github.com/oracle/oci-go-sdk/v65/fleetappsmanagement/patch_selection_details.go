// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PatchSelectionDetails Patch Selection Details
type PatchSelectionDetails interface {
}

type patchselectiondetails struct {
	JsonData      []byte
	SelectionType string `json:"selectionType"`
}

// UnmarshalJSON unmarshals json
func (m *patchselectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchselectiondetails patchselectiondetails
	s := struct {
		Model Unmarshalerpatchselectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SelectionType = s.Model.SelectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *patchselectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SelectionType {
	case "PATCH_LEVEL":
		mm := PatchLevelSelectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATCH_RELEASE_DATE":
		mm := PatchReleaseDateSelectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATCH_NAME":
		mm := PatchNameSelectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PatchSelectionDetails: %s.", m.SelectionType)
		return *m, nil
	}
}

func (m patchselectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchselectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchSelectionDetailsSelectionTypeEnum Enum with underlying type: string
type PatchSelectionDetailsSelectionTypeEnum string

// Set of constants representing the allowable values for PatchSelectionDetailsSelectionTypeEnum
const (
	PatchSelectionDetailsSelectionTypeLevel       PatchSelectionDetailsSelectionTypeEnum = "PATCH_LEVEL"
	PatchSelectionDetailsSelectionTypeName        PatchSelectionDetailsSelectionTypeEnum = "PATCH_NAME"
	PatchSelectionDetailsSelectionTypeReleaseDate PatchSelectionDetailsSelectionTypeEnum = "PATCH_RELEASE_DATE"
)

var mappingPatchSelectionDetailsSelectionTypeEnum = map[string]PatchSelectionDetailsSelectionTypeEnum{
	"PATCH_LEVEL":        PatchSelectionDetailsSelectionTypeLevel,
	"PATCH_NAME":         PatchSelectionDetailsSelectionTypeName,
	"PATCH_RELEASE_DATE": PatchSelectionDetailsSelectionTypeReleaseDate,
}

var mappingPatchSelectionDetailsSelectionTypeEnumLowerCase = map[string]PatchSelectionDetailsSelectionTypeEnum{
	"patch_level":        PatchSelectionDetailsSelectionTypeLevel,
	"patch_name":         PatchSelectionDetailsSelectionTypeName,
	"patch_release_date": PatchSelectionDetailsSelectionTypeReleaseDate,
}

// GetPatchSelectionDetailsSelectionTypeEnumValues Enumerates the set of values for PatchSelectionDetailsSelectionTypeEnum
func GetPatchSelectionDetailsSelectionTypeEnumValues() []PatchSelectionDetailsSelectionTypeEnum {
	values := make([]PatchSelectionDetailsSelectionTypeEnum, 0)
	for _, v := range mappingPatchSelectionDetailsSelectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSelectionDetailsSelectionTypeEnumStringValues Enumerates the set of values in String for PatchSelectionDetailsSelectionTypeEnum
func GetPatchSelectionDetailsSelectionTypeEnumStringValues() []string {
	return []string{
		"PATCH_LEVEL",
		"PATCH_NAME",
		"PATCH_RELEASE_DATE",
	}
}

// GetMappingPatchSelectionDetailsSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSelectionDetailsSelectionTypeEnum(val string) (PatchSelectionDetailsSelectionTypeEnum, bool) {
	enum, ok := mappingPatchSelectionDetailsSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
