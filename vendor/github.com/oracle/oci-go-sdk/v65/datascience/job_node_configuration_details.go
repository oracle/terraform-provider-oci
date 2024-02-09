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

// JobNodeConfigurationDetails The job node configuration details
type JobNodeConfigurationDetails interface {
}

type jobnodeconfigurationdetails struct {
	JsonData    []byte
	JobNodeType string `json:"jobNodeType"`
}

// UnmarshalJSON unmarshals json
func (m *jobnodeconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobnodeconfigurationdetails jobnodeconfigurationdetails
	s := struct {
		Model Unmarshalerjobnodeconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobNodeType = s.Model.JobNodeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobnodeconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobNodeType {
	case "MULTI_NODE":
		mm := MultiNodeJobNodeConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for JobNodeConfigurationDetails: %s.", m.JobNodeType)
		return *m, nil
	}
}

func (m jobnodeconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobnodeconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobNodeConfigurationDetailsJobNodeTypeEnum Enum with underlying type: string
type JobNodeConfigurationDetailsJobNodeTypeEnum string

// Set of constants representing the allowable values for JobNodeConfigurationDetailsJobNodeTypeEnum
const (
	JobNodeConfigurationDetailsJobNodeTypeMultiNode JobNodeConfigurationDetailsJobNodeTypeEnum = "MULTI_NODE"
)

var mappingJobNodeConfigurationDetailsJobNodeTypeEnum = map[string]JobNodeConfigurationDetailsJobNodeTypeEnum{
	"MULTI_NODE": JobNodeConfigurationDetailsJobNodeTypeMultiNode,
}

var mappingJobNodeConfigurationDetailsJobNodeTypeEnumLowerCase = map[string]JobNodeConfigurationDetailsJobNodeTypeEnum{
	"multi_node": JobNodeConfigurationDetailsJobNodeTypeMultiNode,
}

// GetJobNodeConfigurationDetailsJobNodeTypeEnumValues Enumerates the set of values for JobNodeConfigurationDetailsJobNodeTypeEnum
func GetJobNodeConfigurationDetailsJobNodeTypeEnumValues() []JobNodeConfigurationDetailsJobNodeTypeEnum {
	values := make([]JobNodeConfigurationDetailsJobNodeTypeEnum, 0)
	for _, v := range mappingJobNodeConfigurationDetailsJobNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobNodeConfigurationDetailsJobNodeTypeEnumStringValues Enumerates the set of values in String for JobNodeConfigurationDetailsJobNodeTypeEnum
func GetJobNodeConfigurationDetailsJobNodeTypeEnumStringValues() []string {
	return []string{
		"MULTI_NODE",
	}
}

// GetMappingJobNodeConfigurationDetailsJobNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobNodeConfigurationDetailsJobNodeTypeEnum(val string) (JobNodeConfigurationDetailsJobNodeTypeEnum, bool) {
	enum, ok := mappingJobNodeConfigurationDetailsJobNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
