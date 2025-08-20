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

// DefinedAlarmCondition Defined Alarm Condition.
type DefinedAlarmCondition struct {

	// The metric name.
	MetricName *string `mandatory:"true" json:"metricName"`

	// Type of defined monitoring template.
	ConditionType ConditionTypeEnum `mandatory:"true" json:"conditionType"`

	// Monitoring template conditions.
	Conditions []Condition `mandatory:"true" json:"conditions"`
}

func (m DefinedAlarmCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefinedAlarmCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConditionTypeEnum(string(m.ConditionType)); !ok && m.ConditionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionType: %s. Supported values are: %s.", m.ConditionType, strings.Join(GetConditionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
