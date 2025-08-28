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

// DefinedMonitoringTemplateSummary Summary information about defined Monitoring Template for specified resourceType.
type DefinedMonitoringTemplateSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the definedMonitoringTemplate.
	Id *string `mandatory:"true" json:"id"`

	// The name of the definedMonitoringTemplate.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The stack monitoring service or application emitting the metric that is evaluated by the alarm.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The resource types OCID.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Defined Monitoring template alarm conditions
	DefinedAlarmConditions []DefinedAlarmCondition `mandatory:"true" json:"definedAlarmConditions"`

	// The date and time the monitoringTemplate was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Type of composite resource type OCID like EBS/PEOPLE_SOFT.
	CompositeType *string `mandatory:"false" json:"compositeType"`

	// The date and time the monitoringTemplate was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DefinedMonitoringTemplateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefinedMonitoringTemplateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
