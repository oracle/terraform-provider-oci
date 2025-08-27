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

// MonitoringTemplateSummary Summary information about Monitoring Template.
type MonitoringTemplateSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoringTemplate
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the monitoring template.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Tenant Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The OCID of the compartment containing the monitoringTemplate.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current status of the monitoring template i.e. whether it is Applied or Not Applied
	Status MonitoringTemplateLifeCycleDetailsEnum `mandatory:"true" json:"status"`

	// The current lifecycle state of the monitoring template
	LifecycleState MonitoringTemplateLifeCycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// A list of destinations for alarm notifications. Each destination is represented by the OCID of a related resource
	Destinations []string `mandatory:"true" json:"destinations"`

	// Total Alarm Conditions
	TotalAlarmConditions *float32 `mandatory:"true" json:"totalAlarmConditions"`

	// Total Applied Alarm Conditions
	TotalAppliedAlarmConditions *float32 `mandatory:"true" json:"totalAppliedAlarmConditions"`

	// The date and time the monitoringTemplate was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the monitoringTemplate was last updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A user-friendly description for the monitoring template
	Description *string `mandatory:"false" json:"description"`

	// List of members of this monitoring template
	Members []MemberReference `mandatory:"false" json:"members"`

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

func (m MonitoringTemplateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringTemplateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonitoringTemplateLifeCycleDetailsEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMonitoringTemplateLifeCycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMonitoringTemplateLifeCycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMonitoringTemplateLifeCycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
