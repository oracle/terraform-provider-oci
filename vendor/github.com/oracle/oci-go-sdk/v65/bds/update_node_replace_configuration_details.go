// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNodeReplaceConfigurationDetails The information about the NodeReplaceConfiguration
type UpdateNodeReplaceConfigurationDetails struct {
	LevelTypeDetails LevelTypeDetails `mandatory:"false" json:"levelTypeDetails"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of compute instance health metric to use for node replacement
	MetricType NodeReplaceConfigurationMetricTypeEnum `mandatory:"false" json:"metricType,omitempty"`

	// This value is the pending duration time to wait for metric emission before triggering node replacement. The value is in minutes.
	DurationInMinutes *int `mandatory:"false" json:"durationInMinutes"`
}

func (m UpdateNodeReplaceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNodeReplaceConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeReplaceConfigurationMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetNodeReplaceConfigurationMetricTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateNodeReplaceConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LevelTypeDetails  leveltypedetails                       `json:"levelTypeDetails"`
		DisplayName       *string                                `json:"displayName"`
		MetricType        NodeReplaceConfigurationMetricTypeEnum `json:"metricType"`
		DurationInMinutes *int                                   `json:"durationInMinutes"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.DisplayName = model.DisplayName

	m.MetricType = model.MetricType

	m.DurationInMinutes = model.DurationInMinutes

	return
}
