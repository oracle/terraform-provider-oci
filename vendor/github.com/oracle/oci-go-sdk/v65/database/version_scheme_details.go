// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VersionSchemeDetails The update should be applied on the database for the selected version scheme.
type VersionSchemeDetails interface {
}

type versionschemedetails struct {
	JsonData []byte
	Source   string `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *versionschemedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerversionschemedetails versionschemedetails
	s := struct {
		Model Unmarshalerversionschemedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *versionschemedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "VERSION_SERIES":
		mm := VersionSchemeFromVersionSeriesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for VersionSchemeDetails: %s.", m.Source)
		return *m, nil
	}
}

func (m versionschemedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m versionschemedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VersionSchemeDetailsSourceEnum Enum with underlying type: string
type VersionSchemeDetailsSourceEnum string

// Set of constants representing the allowable values for VersionSchemeDetailsSourceEnum
const (
	VersionSchemeDetailsSourceVersionSeries VersionSchemeDetailsSourceEnum = "VERSION_SERIES"
)

var mappingVersionSchemeDetailsSourceEnum = map[string]VersionSchemeDetailsSourceEnum{
	"VERSION_SERIES": VersionSchemeDetailsSourceVersionSeries,
}

var mappingVersionSchemeDetailsSourceEnumLowerCase = map[string]VersionSchemeDetailsSourceEnum{
	"version_series": VersionSchemeDetailsSourceVersionSeries,
}

// GetVersionSchemeDetailsSourceEnumValues Enumerates the set of values for VersionSchemeDetailsSourceEnum
func GetVersionSchemeDetailsSourceEnumValues() []VersionSchemeDetailsSourceEnum {
	values := make([]VersionSchemeDetailsSourceEnum, 0)
	for _, v := range mappingVersionSchemeDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetVersionSchemeDetailsSourceEnumStringValues Enumerates the set of values in String for VersionSchemeDetailsSourceEnum
func GetVersionSchemeDetailsSourceEnumStringValues() []string {
	return []string{
		"VERSION_SERIES",
	}
}

// GetMappingVersionSchemeDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVersionSchemeDetailsSourceEnum(val string) (VersionSchemeDetailsSourceEnum, bool) {
	enum, ok := mappingVersionSchemeDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
