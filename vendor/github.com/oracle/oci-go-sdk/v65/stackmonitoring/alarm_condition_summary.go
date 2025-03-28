// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmConditionSummary Alarm condition summary information about alarm condition
type AlarmConditionSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Alarm Condition
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
	MonitoringTemplateId *string `mandatory:"true" json:"monitoringTemplateId"`

	// The stack monitoring service or application emitting the metric that is evaluated by the alarm.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The resource group OCID.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The metric name.
	MetricName *string `mandatory:"true" json:"metricName"`

	// Alarm conditions
	Conditions []Condition `mandatory:"true" json:"conditions"`

	// The current status of the alarm condition i.e. whether it is Applied or Not Applied
	Status AlarmConditionLifeCycleDetailsEnum `mandatory:"true" json:"status"`

	// The current lifecycle state of the alarm condition.
	LifecycleState AlarmConditionLifeCycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the composite resource type like EBS/PEOPLE_SOFT.
	CompositeType *string `mandatory:"false" json:"compositeType"`

	// Type of alarm condition type.
	ConditionType ConditionTypeEnum `mandatory:"false" json:"conditionType,omitempty"`

	// The date and time the monitoringTemplate was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the monitoringTemplate was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AlarmConditionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmConditionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmConditionLifeCycleDetailsEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAlarmConditionLifeCycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlarmConditionLifeCycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlarmConditionLifeCycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConditionTypeEnum(string(m.ConditionType)); !ok && m.ConditionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionType: %s. Supported values are: %s.", m.ConditionType, strings.Join(GetConditionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
