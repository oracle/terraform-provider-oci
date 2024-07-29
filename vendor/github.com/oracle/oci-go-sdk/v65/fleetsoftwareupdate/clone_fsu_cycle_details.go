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

// CloneFsuCycleDetails Details for cloning an existing Exadata Fleet Update Cycle resource.
type CloneFsuCycleDetails struct {
	GoalVersionDetails FsuGoalVersionDetails `mandatory:"true" json:"goalVersionDetails"`

	// Exadata Fleet Update Cycle display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle will be assigned to.
	// If not specified, it will be assigned to the same Collection as the source Exadata Fleet Update Cycle.
	FsuCollectionId *string `mandatory:"false" json:"fsuCollectionId"`

	BatchingStrategy CreateBatchingStrategyDetails `mandatory:"false" json:"batchingStrategy"`

	StageActionSchedule CreateScheduleDetails `mandatory:"false" json:"stageActionSchedule"`

	ApplyActionSchedule CreateScheduleDetails `mandatory:"false" json:"applyActionSchedule"`
}

func (m CloneFsuCycleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloneFsuCycleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CloneFsuCycleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                       `json:"displayName"`
		CompartmentId       *string                       `json:"compartmentId"`
		FsuCollectionId     *string                       `json:"fsuCollectionId"`
		BatchingStrategy    createbatchingstrategydetails `json:"batchingStrategy"`
		StageActionSchedule createscheduledetails         `json:"stageActionSchedule"`
		ApplyActionSchedule createscheduledetails         `json:"applyActionSchedule"`
		GoalVersionDetails  fsugoalversiondetails         `json:"goalVersionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.FsuCollectionId = model.FsuCollectionId

	nn, e = model.BatchingStrategy.UnmarshalPolymorphicJSON(model.BatchingStrategy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BatchingStrategy = nn.(CreateBatchingStrategyDetails)
	} else {
		m.BatchingStrategy = nil
	}

	nn, e = model.StageActionSchedule.UnmarshalPolymorphicJSON(model.StageActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StageActionSchedule = nn.(CreateScheduleDetails)
	} else {
		m.StageActionSchedule = nil
	}

	nn, e = model.ApplyActionSchedule.UnmarshalPolymorphicJSON(model.ApplyActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ApplyActionSchedule = nn.(CreateScheduleDetails)
	} else {
		m.ApplyActionSchedule = nil
	}

	nn, e = model.GoalVersionDetails.UnmarshalPolymorphicJSON(model.GoalVersionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GoalVersionDetails = nn.(FsuGoalVersionDetails)
	} else {
		m.GoalVersionDetails = nil
	}

	return
}
