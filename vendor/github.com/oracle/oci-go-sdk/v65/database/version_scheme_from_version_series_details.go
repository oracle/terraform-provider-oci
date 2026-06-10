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

// VersionSchemeFromVersionSeriesDetails The update should be applied on the database for the selected version series and preference.
type VersionSchemeFromVersionSeriesDetails struct {

	// The update should be applied on the database for the selected major version series.
	// The value can be provided as 23.X.X.X then 23 major version series will be considered.
	// The list of supported versions can be obtained using the API for the provided shape
	// /20160918/dbVersions?compartmentId=<compartmentId>&dbSystemShape=ExaDbXS
	MajorVersion *string `mandatory:"true" json:"majorVersion"`

	// The update should be applied on the database for the selected version preference. *_N represents the LATEST version
	// For Ex: The current latest version is 23.7.0.0.0,
	// If versionPreference selects option as ORACLE_DB_N then oracle applies the db update with LATEST version (i.e. 23.7.0.0.0)
	// If versionPreference selects option as ORACLE_DB_N_1 then oracle applies the db update with LATEST-1 version (i.e. 23.6.0.0.0)
	// If versionPreference selects option as ORACLE_DB_N_2 then oracle applies the db update with LATEST-2 version (i.e. 23.5.0.0.0)
	VersionPreference VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum `mandatory:"true" json:"versionPreference"`
}

func (m VersionSchemeFromVersionSeriesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VersionSchemeFromVersionSeriesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VersionSchemeFromVersionSeriesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVersionSchemeFromVersionSeriesDetails VersionSchemeFromVersionSeriesDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeVersionSchemeFromVersionSeriesDetails
	}{
		"VERSION_SERIES",
		(MarshalTypeVersionSchemeFromVersionSeriesDetails)(m),
	}

	return json.Marshal(&s)
}

// VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum Enum with underlying type: string
type VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum string

// Set of constants representing the allowable values for VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum
const (
	VersionSchemeFromVersionSeriesDetailsVersionPreferenceN  VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum = "ORACLE_DB_N"
	VersionSchemeFromVersionSeriesDetailsVersionPreferenceN1 VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum = "ORACLE_DB_N_1"
	VersionSchemeFromVersionSeriesDetailsVersionPreferenceN2 VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum = "ORACLE_DB_N_2"
	VersionSchemeFromVersionSeriesDetailsVersionPreferenceN3 VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum = "ORACLE_DB_N_3"
)

var mappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum = map[string]VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum{
	"ORACLE_DB_N":   VersionSchemeFromVersionSeriesDetailsVersionPreferenceN,
	"ORACLE_DB_N_1": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN1,
	"ORACLE_DB_N_2": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN2,
	"ORACLE_DB_N_3": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN3,
}

var mappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumLowerCase = map[string]VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum{
	"oracle_db_n":   VersionSchemeFromVersionSeriesDetailsVersionPreferenceN,
	"oracle_db_n_1": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN1,
	"oracle_db_n_2": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN2,
	"oracle_db_n_3": VersionSchemeFromVersionSeriesDetailsVersionPreferenceN3,
}

// GetVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumValues Enumerates the set of values for VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum
func GetVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumValues() []VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum {
	values := make([]VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum, 0)
	for _, v := range mappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumStringValues Enumerates the set of values in String for VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum
func GetVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumStringValues() []string {
	return []string{
		"ORACLE_DB_N",
		"ORACLE_DB_N_1",
		"ORACLE_DB_N_2",
		"ORACLE_DB_N_3",
	}
}

// GetMappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum(val string) (VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum, bool) {
	enum, ok := mappingVersionSchemeFromVersionSeriesDetailsVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
