// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledFleetSummary Summary of Fleet part of the Schedule.
type ScheduledFleetSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the tenancy to which the resource belongs to.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Count of Resources affected by the Schedule
	CountOfAffectedResources *int `mandatory:"false" json:"countOfAffectedResources"`

	// Count of Targets affected by the Schedule
	CountOfAffectedTargets *int `mandatory:"false" json:"countOfAffectedTargets"`

	// All ActionGroup Types part of the schedule.
	ActionGroupTypes []LifeCycleActionGroupTypeEnum `mandatory:"false" json:"actionGroupTypes,omitempty"`

	// All application types part of the schedule.
	ApplicationTypes []string `mandatory:"false" json:"applicationTypes"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ScheduledFleetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledFleetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ActionGroupTypes {
		if _, ok := GetMappingLifeCycleActionGroupTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionGroupTypes: %s. Supported values are: %s.", val, strings.Join(GetLifeCycleActionGroupTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
