// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAutomaticDrConfigurationDetails The details for updating an Automatic DR configuration.
type UpdateAutomaticDrConfigurationDetails struct {

	// The display name of the Automatic DR configuration being updated.
	// Example: `Automatic DR Configuration`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A list of members for Automatic DR configuration.
	Members []UpdateAutomaticDrConfigurationMemberDetails `mandatory:"false" json:"members"`

	// The unique id of a Switchover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultSwitchoverDrPlanId *string `mandatory:"false" json:"defaultSwitchoverDrPlanId"`

	// The unique id of a Failover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultFailoverDrPlanId *string `mandatory:"false" json:"defaultFailoverDrPlanId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAutomaticDrConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutomaticDrConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateAutomaticDrConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                                       `json:"displayName"`
		Members                   []updateautomaticdrconfigurationmemberdetails `json:"members"`
		DefaultSwitchoverDrPlanId *string                                       `json:"defaultSwitchoverDrPlanId"`
		DefaultFailoverDrPlanId   *string                                       `json:"defaultFailoverDrPlanId"`
		FreeformTags              map[string]string                             `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{}             `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Members = make([]UpdateAutomaticDrConfigurationMemberDetails, len(model.Members))
	for i, n := range model.Members {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Members[i] = nn.(UpdateAutomaticDrConfigurationMemberDetails)
		} else {
			m.Members[i] = nil
		}
	}
	m.DefaultSwitchoverDrPlanId = model.DefaultSwitchoverDrPlanId

	m.DefaultFailoverDrPlanId = model.DefaultFailoverDrPlanId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
