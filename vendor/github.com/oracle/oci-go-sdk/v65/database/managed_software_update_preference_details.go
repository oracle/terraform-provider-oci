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

// ManagedSoftwareUpdatePreferenceDetails Oracle Managed Database Software Updates schedule will be created based on the provided update preferences
type ManagedSoftwareUpdatePreferenceDetails struct {

	// The update should be applied on the database for the selected days of the week.
	DaysOfWeek []ManagedSoftwareUpdateDayOfWeek `mandatory:"true" json:"daysOfWeek"`

	// The update should be applied on the database for the selected hour of the day.
	HourOfDay *int `mandatory:"true" json:"hourOfDay"`

	VersionSchemeDetails VersionSchemeDetails `mandatory:"true" json:"versionSchemeDetails"`

	// Oracle Managed Database Software update method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	// *IMPORTANT*: Non-rolling Database Software update update involves system down time.
	UpdateMode ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum `mandatory:"false" json:"updateMode,omitempty"`
}

func (m ManagedSoftwareUpdatePreferenceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedSoftwareUpdatePreferenceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ManagedSoftwareUpdatePreferenceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		UpdateMode           ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum `json:"updateMode"`
		DaysOfWeek           []ManagedSoftwareUpdateDayOfWeek                     `json:"daysOfWeek"`
		HourOfDay            *int                                                 `json:"hourOfDay"`
		VersionSchemeDetails versionschemedetails                                 `json:"versionSchemeDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.UpdateMode = model.UpdateMode

	m.DaysOfWeek = make([]ManagedSoftwareUpdateDayOfWeek, len(model.DaysOfWeek))
	copy(m.DaysOfWeek, model.DaysOfWeek)
	m.HourOfDay = model.HourOfDay

	nn, e = model.VersionSchemeDetails.UnmarshalPolymorphicJSON(model.VersionSchemeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.VersionSchemeDetails = nn.(VersionSchemeDetails)
	} else {
		m.VersionSchemeDetails = nil
	}

	return
}

// ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum Enum with underlying type: string
type ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum string

// Set of constants representing the allowable values for ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum
const (
	ManagedSoftwareUpdatePreferenceDetailsUpdateModeRolling    ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum = "ROLLING"
	ManagedSoftwareUpdatePreferenceDetailsUpdateModeNonrolling ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum = "NONROLLING"
)

var mappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum = map[string]ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum{
	"ROLLING":    ManagedSoftwareUpdatePreferenceDetailsUpdateModeRolling,
	"NONROLLING": ManagedSoftwareUpdatePreferenceDetailsUpdateModeNonrolling,
}

var mappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumLowerCase = map[string]ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum{
	"rolling":    ManagedSoftwareUpdatePreferenceDetailsUpdateModeRolling,
	"nonrolling": ManagedSoftwareUpdatePreferenceDetailsUpdateModeNonrolling,
}

// GetManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumValues Enumerates the set of values for ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum
func GetManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumValues() []ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum {
	values := make([]ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum, 0)
	for _, v := range mappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumStringValues Enumerates the set of values in String for ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum
func GetManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum(val string) (ManagedSoftwareUpdatePreferenceDetailsUpdateModeEnum, bool) {
	enum, ok := mappingManagedSoftwareUpdatePreferenceDetailsUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
