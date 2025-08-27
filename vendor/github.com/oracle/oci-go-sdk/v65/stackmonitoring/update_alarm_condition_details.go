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

// UpdateAlarmConditionDetails The information about updating a alarm condition in the same monitoringTemplate in a compartment.
type UpdateAlarmConditionDetails struct {

	// The stack monitoring service or application emitting the metric that is evaluated by the alarm.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The OCID of composite resource type like EBS/PEOPLE_SOFT.
	CompositeType *string `mandatory:"false" json:"compositeType"`

	// The resource type OCID.
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// The metric name.
	MetricName *string `mandatory:"false" json:"metricName"`

	// Type of defined monitoring template.
	ConditionType ConditionTypeEnum `mandatory:"false" json:"conditionType,omitempty"`

	// Monitoring template conditions
	Conditions []Condition `mandatory:"false" json:"conditions"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAlarmConditionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAlarmConditionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConditionTypeEnum(string(m.ConditionType)); !ok && m.ConditionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionType: %s. Supported values are: %s.", m.ConditionType, strings.Join(GetConditionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
