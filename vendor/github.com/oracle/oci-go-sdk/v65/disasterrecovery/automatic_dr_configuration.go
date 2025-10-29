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

// AutomaticDrConfiguration The details of Automatic DR configuration.
type AutomaticDrConfiguration struct {

	// The OCID of the Automatic DR configuration.
	// Example: `ocid1.automaticdrconfiguration.oc1..uniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The display name of the Automatic DR configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the Automatic DR configuration.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the Automatic DR configuration was created. An RFC3339 formatted datetime string.
	// Example: `2024-03-29T09:36:42Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Automatic DR configuration was updated. An RFC3339 formatted datetime string.
	// Example: `2024-03-29T09:36:42Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the DR protection group to which this Automatic DR configuration belongs.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" json:"drProtectionGroupId"`

	// The list of members in this Automatic DR configuration.
	Members []AutomaticDrConfigurationMember `mandatory:"true" json:"members"`

	// The current state of the Automatic DR configuration.
	LifecycleState AutomaticDrConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The unique id of a Switchover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultSwitchoverDrPlanId *string `mandatory:"false" json:"defaultSwitchoverDrPlanId"`

	// The unique id of a Failover DR Plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DefaultFailoverDrPlanId *string `mandatory:"false" json:"defaultFailoverDrPlanId"`

	// A message describing the Automatic DR configuration's current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The status of most recent attempt to submit Automatic DR plan execution.
	LastAutomaticDrExecutionSubmitStatus AutomaticDrPlanExecutionSubmissionStatusEnum `mandatory:"false" json:"lastAutomaticDrExecutionSubmitStatus,omitempty"`

	// A message describing the result of the most recent attempt made to submit an Automatic DR plan execution.
	LastAutomaticDrExecutionSubmitDetails *string `mandatory:"false" json:"lastAutomaticDrExecutionSubmitDetails"`

	// The date and time of the most recent attempt made to submit an Automatic DR plan execution. An RFC3339 formatted datetime string.
	// Example: `2025-06-30T09:36:42Z`
	TimeLastAutomaticDrExecutionSubmitAttempt *common.SDKTime `mandatory:"false" json:"timeLastAutomaticDrExecutionSubmitAttempt"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AutomaticDrConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomaticDrConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutomaticDrConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutomaticDrConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutomaticDrPlanExecutionSubmissionStatusEnum(string(m.LastAutomaticDrExecutionSubmitStatus)); !ok && m.LastAutomaticDrExecutionSubmitStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAutomaticDrExecutionSubmitStatus: %s. Supported values are: %s.", m.LastAutomaticDrExecutionSubmitStatus, strings.Join(GetAutomaticDrPlanExecutionSubmissionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutomaticDrConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefaultSwitchoverDrPlanId                 *string                                      `json:"defaultSwitchoverDrPlanId"`
		DefaultFailoverDrPlanId                   *string                                      `json:"defaultFailoverDrPlanId"`
		LifecycleDetails                          *string                                      `json:"lifecycleDetails"`
		LastAutomaticDrExecutionSubmitStatus      AutomaticDrPlanExecutionSubmissionStatusEnum `json:"lastAutomaticDrExecutionSubmitStatus"`
		LastAutomaticDrExecutionSubmitDetails     *string                                      `json:"lastAutomaticDrExecutionSubmitDetails"`
		TimeLastAutomaticDrExecutionSubmitAttempt *common.SDKTime                              `json:"timeLastAutomaticDrExecutionSubmitAttempt"`
		FreeformTags                              map[string]string                            `json:"freeformTags"`
		DefinedTags                               map[string]map[string]interface{}            `json:"definedTags"`
		SystemTags                                map[string]map[string]interface{}            `json:"systemTags"`
		Id                                        *string                                      `json:"id"`
		DisplayName                               *string                                      `json:"displayName"`
		CompartmentId                             *string                                      `json:"compartmentId"`
		TimeCreated                               *common.SDKTime                              `json:"timeCreated"`
		TimeUpdated                               *common.SDKTime                              `json:"timeUpdated"`
		DrProtectionGroupId                       *string                                      `json:"drProtectionGroupId"`
		Members                                   []automaticdrconfigurationmember             `json:"members"`
		LifecycleState                            AutomaticDrConfigurationLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefaultSwitchoverDrPlanId = model.DefaultSwitchoverDrPlanId

	m.DefaultFailoverDrPlanId = model.DefaultFailoverDrPlanId

	m.LifecycleDetails = model.LifecycleDetails

	m.LastAutomaticDrExecutionSubmitStatus = model.LastAutomaticDrExecutionSubmitStatus

	m.LastAutomaticDrExecutionSubmitDetails = model.LastAutomaticDrExecutionSubmitDetails

	m.TimeLastAutomaticDrExecutionSubmitAttempt = model.TimeLastAutomaticDrExecutionSubmitAttempt

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.DrProtectionGroupId = model.DrProtectionGroupId

	m.Members = make([]AutomaticDrConfigurationMember, len(model.Members))
	for i, n := range model.Members {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Members[i] = nn.(AutomaticDrConfigurationMember)
		} else {
			m.Members[i] = nil
		}
	}
	m.LifecycleState = model.LifecycleState

	return
}
