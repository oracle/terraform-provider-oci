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

// JobNetworkConfiguration The job network configuration details
type JobNetworkConfiguration interface {
}

type jobnetworkconfiguration struct {
	JsonData       []byte
	JobNetworkType string `json:"jobNetworkType"`
}

// UnmarshalJSON unmarshals json
func (m *jobnetworkconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobnetworkconfiguration jobnetworkconfiguration
	s := struct {
		Model Unmarshalerjobnetworkconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.JobNetworkType = s.Model.JobNetworkType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobnetworkconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobNetworkType {
	case "CUSTOM_NETWORK":
		mm := JobCustomNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFAULT_NETWORK":
		mm := JobDefaultNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for JobNetworkConfiguration: %s.", m.JobNetworkType)
		return *m, nil
	}
}

func (m jobnetworkconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m jobnetworkconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobNetworkConfigurationJobNetworkTypeEnum Enum with underlying type: string
type JobNetworkConfigurationJobNetworkTypeEnum string

// Set of constants representing the allowable values for JobNetworkConfigurationJobNetworkTypeEnum
const (
	JobNetworkConfigurationJobNetworkTypeCustomNetwork  JobNetworkConfigurationJobNetworkTypeEnum = "CUSTOM_NETWORK"
	JobNetworkConfigurationJobNetworkTypeDefaultNetwork JobNetworkConfigurationJobNetworkTypeEnum = "DEFAULT_NETWORK"
)

var mappingJobNetworkConfigurationJobNetworkTypeEnum = map[string]JobNetworkConfigurationJobNetworkTypeEnum{
	"CUSTOM_NETWORK":  JobNetworkConfigurationJobNetworkTypeCustomNetwork,
	"DEFAULT_NETWORK": JobNetworkConfigurationJobNetworkTypeDefaultNetwork,
}

var mappingJobNetworkConfigurationJobNetworkTypeEnumLowerCase = map[string]JobNetworkConfigurationJobNetworkTypeEnum{
	"custom_network":  JobNetworkConfigurationJobNetworkTypeCustomNetwork,
	"default_network": JobNetworkConfigurationJobNetworkTypeDefaultNetwork,
}

// GetJobNetworkConfigurationJobNetworkTypeEnumValues Enumerates the set of values for JobNetworkConfigurationJobNetworkTypeEnum
func GetJobNetworkConfigurationJobNetworkTypeEnumValues() []JobNetworkConfigurationJobNetworkTypeEnum {
	values := make([]JobNetworkConfigurationJobNetworkTypeEnum, 0)
	for _, v := range mappingJobNetworkConfigurationJobNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobNetworkConfigurationJobNetworkTypeEnumStringValues Enumerates the set of values in String for JobNetworkConfigurationJobNetworkTypeEnum
func GetJobNetworkConfigurationJobNetworkTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_NETWORK",
		"DEFAULT_NETWORK",
	}
}

// GetMappingJobNetworkConfigurationJobNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobNetworkConfigurationJobNetworkTypeEnum(val string) (JobNetworkConfigurationJobNetworkTypeEnum, bool) {
	enum, ok := mappingJobNetworkConfigurationJobNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
