// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateNodeReplaceConfigurationDetails The information about the NodeReplaceConfiguration
type CreateNodeReplaceConfigurationDetails struct {
	LevelTypeDetails LevelTypeDetails `mandatory:"true" json:"levelTypeDetails"`

	// Base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// Type of compute instance health metric to use for node replacement
	MetricType NodeReplaceConfigurationMetricTypeEnum `mandatory:"true" json:"metricType"`

	// This value is the minimum period of time to wait before triggering node replacement. The value is in minutes.
	DurationInMinutes *int `mandatory:"true" json:"durationInMinutes"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateNodeReplaceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateNodeReplaceConfigurationDetails) ValidateEnumValue() (bool, error) {
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
func (m *CreateNodeReplaceConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                                `json:"displayName"`
		LevelTypeDetails     leveltypedetails                       `json:"levelTypeDetails"`
		ClusterAdminPassword *string                                `json:"clusterAdminPassword"`
		MetricType           NodeReplaceConfigurationMetricTypeEnum `json:"metricType"`
		DurationInMinutes    *int                                   `json:"durationInMinutes"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.LevelTypeDetails.UnmarshalPolymorphicJSON(model.LevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LevelTypeDetails = nn.(LevelTypeDetails)
	} else {
		m.LevelTypeDetails = nil
	}

	m.ClusterAdminPassword = model.ClusterAdminPassword

	m.MetricType = model.MetricType

	m.DurationInMinutes = model.DurationInMinutes

	return
}
