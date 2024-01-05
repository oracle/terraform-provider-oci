// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAnnouncementsPreferencesDetails The object used to create announcement email preferences.
type CreateAnnouncementsPreferencesDetails struct {

	// A Boolean value to indicate whether the specified compartment chooses to not to receive informational announcements by email.
	// (Manage preferences for receiving announcements by email by specifying the `preferenceType` attribute instead.)
	IsUnsubscribed *bool `mandatory:"false" json:"isUnsubscribed"`

	// The OCID of the compartment for which you want to manage announcement email preferences. (Specify the tenancy by providing the
	// root compartment OCID.)
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time zone in which the user prefers to receive announcements. Specify the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles
	PreferredTimeZone *string `mandatory:"false" json:"preferredTimeZone"`

	// The string representing the user's preference, whether to opt in to only required announcements, to opt in to all announcements, including informational announcements, or to opt out of all announcements.
	PreferenceType BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum `mandatory:"true" json:"preferenceType"`
}

// GetIsUnsubscribed returns IsUnsubscribed
func (m CreateAnnouncementsPreferencesDetails) GetIsUnsubscribed() *bool {
	return m.IsUnsubscribed
}

// GetCompartmentId returns CompartmentId
func (m CreateAnnouncementsPreferencesDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPreferenceType returns PreferenceType
func (m CreateAnnouncementsPreferencesDetails) GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	return m.PreferenceType
}

// GetPreferredTimeZone returns PreferredTimeZone
func (m CreateAnnouncementsPreferencesDetails) GetPreferredTimeZone() *string {
	return m.PreferredTimeZone
}

func (m CreateAnnouncementsPreferencesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAnnouncementsPreferencesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum(string(m.PreferenceType)); !ok && m.PreferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferenceType: %s. Supported values are: %s.", m.PreferenceType, strings.Join(GetBaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAnnouncementsPreferencesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAnnouncementsPreferencesDetails CreateAnnouncementsPreferencesDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateAnnouncementsPreferencesDetails
	}{
		"CreateAnnouncementsPreferencesDetails",
		(MarshalTypeCreateAnnouncementsPreferencesDetails)(m),
	}

	return json.Marshal(&s)
}
