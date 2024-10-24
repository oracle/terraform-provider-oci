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

// ArtifactDetails Patch artifact description and content details.
type ArtifactDetails interface {
}

type artifactdetails struct {
	JsonData []byte
	Category string `json:"category"`
}

// UnmarshalJSON unmarshals json
func (m *artifactdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerartifactdetails artifactdetails
	s := struct {
		Model Unmarshalerartifactdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Category = s.Model.Category

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *artifactdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Category {
	case "PLATFORM_SPECIFIC":
		mm := PlatformSpecificArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC":
		mm := GenericArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ArtifactDetails: %s.", m.Category)
		return *m, nil
	}
}

func (m artifactdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m artifactdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArtifactDetailsCategoryEnum Enum with underlying type: string
type ArtifactDetailsCategoryEnum string

// Set of constants representing the allowable values for ArtifactDetailsCategoryEnum
const (
	ArtifactDetailsCategoryGeneric          ArtifactDetailsCategoryEnum = "GENERIC"
	ArtifactDetailsCategoryPlatformSpecific ArtifactDetailsCategoryEnum = "PLATFORM_SPECIFIC"
)

var mappingArtifactDetailsCategoryEnum = map[string]ArtifactDetailsCategoryEnum{
	"GENERIC":           ArtifactDetailsCategoryGeneric,
	"PLATFORM_SPECIFIC": ArtifactDetailsCategoryPlatformSpecific,
}

var mappingArtifactDetailsCategoryEnumLowerCase = map[string]ArtifactDetailsCategoryEnum{
	"generic":           ArtifactDetailsCategoryGeneric,
	"platform_specific": ArtifactDetailsCategoryPlatformSpecific,
}

// GetArtifactDetailsCategoryEnumValues Enumerates the set of values for ArtifactDetailsCategoryEnum
func GetArtifactDetailsCategoryEnumValues() []ArtifactDetailsCategoryEnum {
	values := make([]ArtifactDetailsCategoryEnum, 0)
	for _, v := range mappingArtifactDetailsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactDetailsCategoryEnumStringValues Enumerates the set of values in String for ArtifactDetailsCategoryEnum
func GetArtifactDetailsCategoryEnumStringValues() []string {
	return []string{
		"GENERIC",
		"PLATFORM_SPECIFIC",
	}
}

// GetMappingArtifactDetailsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactDetailsCategoryEnum(val string) (ArtifactDetailsCategoryEnum, bool) {
	enum, ok := mappingArtifactDetailsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
