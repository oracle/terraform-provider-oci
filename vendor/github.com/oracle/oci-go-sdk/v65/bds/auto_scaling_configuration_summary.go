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

// AutoScalingConfigurationSummary The information about the autoscale configuration.
type AutoScalingConfigurationSummary struct {

	// The OCID of the autoscale configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the autoscale configuration.
	LifecycleState AutoScalingConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A node type that is managed by an autoscale configuration. The only supported types are WORKER and COMPUTE_ONLY_WORKER.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the autoscale configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Policy *AutoScalePolicy `mandatory:"true" json:"policy"`

	PolicyDetails AutoScalePolicyDetails `mandatory:"false" json:"policyDetails"`
}

func (m AutoScalingConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalingConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoScalingConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutoScalingConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AutoScalingConfigurationSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PolicyDetails  autoscalepolicydetails                     `json:"policyDetails"`
		Id             *string                                    `json:"id"`
		DisplayName    *string                                    `json:"displayName"`
		LifecycleState AutoScalingConfigurationLifecycleStateEnum `json:"lifecycleState"`
		NodeType       NodeNodeTypeEnum                           `json:"nodeType"`
		TimeCreated    *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                            `json:"timeUpdated"`
		Policy         *AutoScalePolicy                           `json:"policy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PolicyDetails.UnmarshalPolymorphicJSON(model.PolicyDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PolicyDetails = nn.(AutoScalePolicyDetails)
	} else {
		m.PolicyDetails = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.NodeType = model.NodeType

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Policy = model.Policy

	return
}
