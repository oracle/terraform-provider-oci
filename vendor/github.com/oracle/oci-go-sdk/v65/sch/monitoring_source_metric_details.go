// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringSourceMetricDetails The metrics to query for the specified metric namespace.
type MonitoringSourceMetricDetails interface {
}

type monitoringsourcemetricdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *monitoringsourcemetricdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermonitoringsourcemetricdetails monitoringsourcemetricdetails
	s := struct {
		Model Unmarshalermonitoringsourcemetricdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *monitoringsourcemetricdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "all":
		mm := MonitoringSourceAllMetrics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MonitoringSourceMetricDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m monitoringsourcemetricdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m monitoringsourcemetricdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoringSourceMetricDetailsKindEnum Enum with underlying type: string
type MonitoringSourceMetricDetailsKindEnum string

// Set of constants representing the allowable values for MonitoringSourceMetricDetailsKindEnum
const (
	MonitoringSourceMetricDetailsKindAll MonitoringSourceMetricDetailsKindEnum = "all"
)

var mappingMonitoringSourceMetricDetailsKindEnum = map[string]MonitoringSourceMetricDetailsKindEnum{
	"all": MonitoringSourceMetricDetailsKindAll,
}

var mappingMonitoringSourceMetricDetailsKindEnumLowerCase = map[string]MonitoringSourceMetricDetailsKindEnum{
	"all": MonitoringSourceMetricDetailsKindAll,
}

// GetMonitoringSourceMetricDetailsKindEnumValues Enumerates the set of values for MonitoringSourceMetricDetailsKindEnum
func GetMonitoringSourceMetricDetailsKindEnumValues() []MonitoringSourceMetricDetailsKindEnum {
	values := make([]MonitoringSourceMetricDetailsKindEnum, 0)
	for _, v := range mappingMonitoringSourceMetricDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoringSourceMetricDetailsKindEnumStringValues Enumerates the set of values in String for MonitoringSourceMetricDetailsKindEnum
func GetMonitoringSourceMetricDetailsKindEnumStringValues() []string {
	return []string{
		"all",
	}
}

// GetMappingMonitoringSourceMetricDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoringSourceMetricDetailsKindEnum(val string) (MonitoringSourceMetricDetailsKindEnum, bool) {
	enum, ok := mappingMonitoringSourceMetricDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
