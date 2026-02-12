// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// FleetResourceBasedActionGroup A string variable that holds a value
type FleetResourceBasedActionGroup struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ID of the fleet
	FleetId *string `mandatory:"false" json:"fleetId"`

	// Sequence of the Action Group.
	// Action groups will be executed in a seuential order.
	// All Action Groups having the same sequence will be executed parallely.
	// If no value is provided a default value of 1 will be given.
	Sequence *int `mandatory:"false" json:"sequence"`

	// ID of the runbook
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the runbook version
	RunbookVersionName *string `mandatory:"false" json:"runbookVersionName"`

	// Emergency override flag.
	IsEmergencyOverride *bool `mandatory:"false" json:"isEmergencyOverride"`

	// List of selected fleet resources.
	SelectedFleetResources []string `mandatory:"false" json:"selectedFleetResources"`

	// Subject Type
	SubjectType SubjectTypeEnum `mandatory:"false" json:"subjectType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m FleetResourceBasedActionGroup) GetDisplayName() *string {
	return m.DisplayName
}

func (m FleetResourceBasedActionGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetResourceBasedActionGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubjectTypeEnum(string(m.SubjectType)); !ok && m.SubjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubjectType: %s. Supported values are: %s.", m.SubjectType, strings.Join(GetSubjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FleetResourceBasedActionGroup) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFleetResourceBasedActionGroup FleetResourceBasedActionGroup
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeFleetResourceBasedActionGroup
	}{
		"FLEET_RESOURCES_USING_RUNBOOK",
		(MarshalTypeFleetResourceBasedActionGroup)(m),
	}

	return json.Marshal(&s)
}
