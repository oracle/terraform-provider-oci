// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAlarmSuppressionDetails The configuration details for creating a dimension-specific alarm suppression.
type CreateAlarmSuppressionDetails struct {
	AlarmSuppressionTarget AlarmSuppressionTarget `mandatory:"true" json:"alarmSuppressionTarget"`

	// A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A filter to suppress only alarm state entries that include the set of specified dimension key-value pairs.
	// If you specify {"availabilityDomain": "phx-ad-1"}
	// and the alarm state entry corresponds to the set {"availabilityDomain": "phx-ad-1" and "resourceId": "ocid1.instance.region1.phx.exampleuniqueID"},
	// then this alarm will be included for suppression.
	// The value cannot be an empty object.
	// Only a single value is allowed per key. No grouping of multiple values is allowed under the same key.
	// Maximum characters (after serialization): 4000. This maximum satisfies typical use cases.
	// The response for an exceeded maximum is `HTTP 400` with an "dimensions values are too long" message.
	Dimensions map[string]string `mandatory:"true" json:"dimensions"`

	// The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeSuppressFrom *common.SDKTime `mandatory:"true" json:"timeSuppressFrom"`

	// The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.
	// Example: `2023-02-01T02:02:29.600Z`
	TimeSuppressUntil *common.SDKTime `mandatory:"true" json:"timeSuppressUntil"`

	// Human-readable reason for this alarm suppression.
	// It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Oracle recommends including tracking information for the event or associated work,
	// such as a ticket number.
	// Example: `Planned outage due to change IT-1234.`
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAlarmSuppressionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAlarmSuppressionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateAlarmSuppressionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                           `json:"description"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		AlarmSuppressionTarget alarmsuppressiontarget            `json:"alarmSuppressionTarget"`
		DisplayName            *string                           `json:"displayName"`
		Dimensions             map[string]string                 `json:"dimensions"`
		TimeSuppressFrom       *common.SDKTime                   `json:"timeSuppressFrom"`
		TimeSuppressUntil      *common.SDKTime                   `json:"timeSuppressUntil"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.AlarmSuppressionTarget.UnmarshalPolymorphicJSON(model.AlarmSuppressionTarget.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AlarmSuppressionTarget = nn.(AlarmSuppressionTarget)
	} else {
		m.AlarmSuppressionTarget = nil
	}

	m.DisplayName = model.DisplayName

	m.Dimensions = model.Dimensions

	m.TimeSuppressFrom = model.TimeSuppressFrom

	m.TimeSuppressUntil = model.TimeSuppressUntil

	return
}
