// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateScheduleDetails The details for creating a schedule.
type CreateScheduleDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Generated key that can be used in API calls to identify schedule. On scenarios where reference to the schedule is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	FrequencyDetails AbstractFrequencyDetails `mandatory:"false" json:"frequencyDetails"`

	// The timezone for the schedule.
	Timezone *string `mandatory:"false" json:"timezone"`

	// A flag to indicate whether daylight adjustment should be considered or not.
	IsDaylightAdjustmentEnabled *bool `mandatory:"false" json:"isDaylightAdjustmentEnabled"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                         *string                  `json:"key"`
		ModelVersion                *string                  `json:"modelVersion"`
		Description                 *string                  `json:"description"`
		ObjectVersion               *int                     `json:"objectVersion"`
		ObjectStatus                *int                     `json:"objectStatus"`
		FrequencyDetails            abstractfrequencydetails `json:"frequencyDetails"`
		Timezone                    *string                  `json:"timezone"`
		IsDaylightAdjustmentEnabled *bool                    `json:"isDaylightAdjustmentEnabled"`
		RegistryMetadata            *RegistryMetadata        `json:"registryMetadata"`
		Name                        *string                  `json:"name"`
		Identifier                  *string                  `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ObjectStatus = model.ObjectStatus

	nn, e = model.FrequencyDetails.UnmarshalPolymorphicJSON(model.FrequencyDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FrequencyDetails = nn.(AbstractFrequencyDetails)
	} else {
		m.FrequencyDetails = nil
	}

	m.Timezone = model.Timezone

	m.IsDaylightAdjustmentEnabled = model.IsDaylightAdjustmentEnabled

	m.RegistryMetadata = model.RegistryMetadata

	m.Name = model.Name

	m.Identifier = model.Identifier

	return
}
