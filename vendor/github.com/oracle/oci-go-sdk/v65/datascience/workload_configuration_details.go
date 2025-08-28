// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// WorkloadConfigurationDetails The workload configuration details
type WorkloadConfigurationDetails interface {
}

type workloadconfigurationdetails struct {
	JsonData     []byte
	WorkloadType string `json:"workloadType"`
}

// UnmarshalJSON unmarshals json
func (m *workloadconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerworkloadconfigurationdetails workloadconfigurationdetails
	s := struct {
		Model Unmarshalerworkloadconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkloadType = s.Model.WorkloadType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *workloadconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WorkloadType {
	case "MODEL_DEPLOYMENT":
		mm := ModelDeployWorkloadConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JOB_RUN":
		mm := JobRunWorkloadConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for WorkloadConfigurationDetails: %s.", m.WorkloadType)
		return *m, nil
	}
}

func (m workloadconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m workloadconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkloadConfigurationDetailsWorkloadTypeEnum Enum with underlying type: string
type WorkloadConfigurationDetailsWorkloadTypeEnum string

// Set of constants representing the allowable values for WorkloadConfigurationDetailsWorkloadTypeEnum
const (
	WorkloadConfigurationDetailsWorkloadTypeModelDeployment WorkloadConfigurationDetailsWorkloadTypeEnum = "MODEL_DEPLOYMENT"
	WorkloadConfigurationDetailsWorkloadTypeJobRun          WorkloadConfigurationDetailsWorkloadTypeEnum = "JOB_RUN"
)

var mappingWorkloadConfigurationDetailsWorkloadTypeEnum = map[string]WorkloadConfigurationDetailsWorkloadTypeEnum{
	"MODEL_DEPLOYMENT": WorkloadConfigurationDetailsWorkloadTypeModelDeployment,
	"JOB_RUN":          WorkloadConfigurationDetailsWorkloadTypeJobRun,
}

var mappingWorkloadConfigurationDetailsWorkloadTypeEnumLowerCase = map[string]WorkloadConfigurationDetailsWorkloadTypeEnum{
	"model_deployment": WorkloadConfigurationDetailsWorkloadTypeModelDeployment,
	"job_run":          WorkloadConfigurationDetailsWorkloadTypeJobRun,
}

// GetWorkloadConfigurationDetailsWorkloadTypeEnumValues Enumerates the set of values for WorkloadConfigurationDetailsWorkloadTypeEnum
func GetWorkloadConfigurationDetailsWorkloadTypeEnumValues() []WorkloadConfigurationDetailsWorkloadTypeEnum {
	values := make([]WorkloadConfigurationDetailsWorkloadTypeEnum, 0)
	for _, v := range mappingWorkloadConfigurationDetailsWorkloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkloadConfigurationDetailsWorkloadTypeEnumStringValues Enumerates the set of values in String for WorkloadConfigurationDetailsWorkloadTypeEnum
func GetWorkloadConfigurationDetailsWorkloadTypeEnumStringValues() []string {
	return []string{
		"MODEL_DEPLOYMENT",
		"JOB_RUN",
	}
}

// GetMappingWorkloadConfigurationDetailsWorkloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkloadConfigurationDetailsWorkloadTypeEnum(val string) (WorkloadConfigurationDetailsWorkloadTypeEnum, bool) {
	enum, ok := mappingWorkloadConfigurationDetailsWorkloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
