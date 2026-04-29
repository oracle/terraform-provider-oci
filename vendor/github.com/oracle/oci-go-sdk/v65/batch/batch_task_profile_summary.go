// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTaskProfileSummary Summary information about a batch task profile.
type BatchTaskProfileSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the batch task profile.
	LifecycleState BatchTaskProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the batch task profile was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The batch task profile description.
	Description *string `mandatory:"false" json:"description"`

	// The minimum required OCPUs.
	MinOcpus *int `mandatory:"false" json:"minOcpus"`

	// The minimum required memory.
	MinMemoryInGBs *int `mandatory:"false" json:"minMemoryInGBs"`

	// The minimum required size of disk space in GBs.
	MinDiskSizeInGBs *int `mandatory:"false" json:"minDiskSizeInGBs"`

	ExtendedInformation BatchTaskProfileExtendedInformation `mandatory:"false" json:"extendedInformation"`

	// The date and time the batch task profile was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BatchTaskProfileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchTaskProfileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskProfileLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BatchTaskProfileSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description         *string                             `json:"description"`
		MinOcpus            *int                                `json:"minOcpus"`
		MinMemoryInGBs      *int                                `json:"minMemoryInGBs"`
		MinDiskSizeInGBs    *int                                `json:"minDiskSizeInGBs"`
		ExtendedInformation batchtaskprofileextendedinformation `json:"extendedInformation"`
		TimeUpdated         *common.SDKTime                     `json:"timeUpdated"`
		Id                  *string                             `json:"id"`
		CompartmentId       *string                             `json:"compartmentId"`
		DisplayName         *string                             `json:"displayName"`
		LifecycleState      BatchTaskProfileLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated         *common.SDKTime                     `json:"timeCreated"`
		DefinedTags         map[string]map[string]interface{}   `json:"definedTags"`
		FreeformTags        map[string]string                   `json:"freeformTags"`
		SystemTags          map[string]map[string]interface{}   `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.MinOcpus = model.MinOcpus

	m.MinMemoryInGBs = model.MinMemoryInGBs

	m.MinDiskSizeInGBs = model.MinDiskSizeInGBs

	nn, e = model.ExtendedInformation.UnmarshalPolymorphicJSON(model.ExtendedInformation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExtendedInformation = nn.(BatchTaskProfileExtendedInformation)
	} else {
		m.ExtendedInformation = nil
	}

	m.TimeUpdated = model.TimeUpdated

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	return
}
