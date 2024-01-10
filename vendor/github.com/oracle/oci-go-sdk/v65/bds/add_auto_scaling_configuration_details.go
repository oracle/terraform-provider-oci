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

// AddAutoScalingConfigurationDetails The information about the autoscale configuration.
type AddAutoScalingConfigurationDetails struct {

	// A node type that is managed by an autoscale configuration. The only supported types are WORKER and COMPUTE_ONLY_WORKER.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// Whether the autoscale configuration is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Policy *AutoScalePolicy `mandatory:"false" json:"policy"`

	PolicyDetails AddAutoScalePolicyDetails `mandatory:"false" json:"policyDetails"`
}

func (m AddAutoScalingConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddAutoScalingConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AddAutoScalingConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                   `json:"displayName"`
		Policy               *AutoScalePolicy          `json:"policy"`
		PolicyDetails        addautoscalepolicydetails `json:"policyDetails"`
		NodeType             NodeNodeTypeEnum          `json:"nodeType"`
		IsEnabled            *bool                     `json:"isEnabled"`
		ClusterAdminPassword *string                   `json:"clusterAdminPassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Policy = model.Policy

	nn, e = model.PolicyDetails.UnmarshalPolymorphicJSON(model.PolicyDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PolicyDetails = nn.(AddAutoScalePolicyDetails)
	} else {
		m.PolicyDetails = nil
	}

	m.NodeType = model.NodeType

	m.IsEnabled = model.IsEnabled

	m.ClusterAdminPassword = model.ClusterAdminPassword

	return
}
