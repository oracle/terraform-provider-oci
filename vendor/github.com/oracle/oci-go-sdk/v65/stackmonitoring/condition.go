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

// Condition The Monitoring Template Alarm Condition.
type Condition struct {

	// Severity - Critical/Warning
	Severity AlarmConditionSeverityEnum `mandatory:"true" json:"severity"`

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm.
	Query *string `mandatory:"true" json:"query"`

	// The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.
	Body *string `mandatory:"false" json:"body"`

	// Whether the note need to add into bottom of the body for mapping the alarms information with template or not.
	ShouldAppendNote *bool `mandatory:"false" json:"shouldAppendNote"`

	// Whether the URL need to add into bottom of the body for mapping the alarms information with template or not.
	ShouldAppendUrl *bool `mandatory:"false" json:"shouldAppendUrl"`

	// The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING".
	TriggerDelay *string `mandatory:"false" json:"triggerDelay"`
}

func (m Condition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Condition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmConditionSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmConditionSeverityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
