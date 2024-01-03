// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobNodeGroupConfigurationDetails Details of Job Node Group Configuration
type JobNodeGroupConfigurationDetails struct {

	// node group name.
	Name *string `mandatory:"true" json:"name"`

	// The number of nodes.
	Replicas *int `mandatory:"false" json:"replicas"`

	JobInfrastructureConfigurationDetails JobInfrastructureConfigurationDetails `mandatory:"false" json:"jobInfrastructureConfigurationDetails"`

	JobConfigurationDetails JobConfigurationDetails `mandatory:"false" json:"jobConfigurationDetails"`

	JobEnvironmentConfigurationDetails JobEnvironmentConfigurationDetails `mandatory:"false" json:"jobEnvironmentConfigurationDetails"`
}

func (m JobNodeGroupConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobNodeGroupConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *JobNodeGroupConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Replicas                              *int                                  `json:"replicas"`
		JobInfrastructureConfigurationDetails jobinfrastructureconfigurationdetails `json:"jobInfrastructureConfigurationDetails"`
		JobConfigurationDetails               jobconfigurationdetails               `json:"jobConfigurationDetails"`
		JobEnvironmentConfigurationDetails    jobenvironmentconfigurationdetails    `json:"jobEnvironmentConfigurationDetails"`
		Name                                  *string                               `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Replicas = model.Replicas

	nn, e = model.JobInfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.JobInfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobInfrastructureConfigurationDetails = nn.(JobInfrastructureConfigurationDetails)
	} else {
		m.JobInfrastructureConfigurationDetails = nil
	}

	nn, e = model.JobConfigurationDetails.UnmarshalPolymorphicJSON(model.JobConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobConfigurationDetails = nn.(JobConfigurationDetails)
	} else {
		m.JobConfigurationDetails = nil
	}

	nn, e = model.JobEnvironmentConfigurationDetails.UnmarshalPolymorphicJSON(model.JobEnvironmentConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobEnvironmentConfigurationDetails = nn.(JobEnvironmentConfigurationDetails)
	} else {
		m.JobEnvironmentConfigurationDetails = nil
	}

	m.Name = model.Name

	return
}
