// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeConfigurationSourceDetails The details of configuration source for nodes
type NodeConfigurationSourceDetails interface {
}

type nodeconfigurationsourcedetails struct {
	JsonData            []byte
	ConfigurationSource string `json:"configurationSource"`
}

// UnmarshalJSON unmarshals json
func (m *nodeconfigurationsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernodeconfigurationsourcedetails nodeconfigurationsourcedetails
	s := struct {
		Model Unmarshalernodeconfigurationsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigurationSource = s.Model.ConfigurationSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *nodeconfigurationsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigurationSource {
	case "INSTANCE_CONFIG":
		mm := NodeConfigSourceFromInstanceConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for NodeConfigurationSourceDetails: %s.", m.ConfigurationSource)
		return *m, nil
	}
}

func (m nodeconfigurationsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m nodeconfigurationsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodeConfigurationSourceDetailsConfigurationSourceEnum Enum with underlying type: string
type NodeConfigurationSourceDetailsConfigurationSourceEnum string

// Set of constants representing the allowable values for NodeConfigurationSourceDetailsConfigurationSourceEnum
const (
	NodeConfigurationSourceDetailsConfigurationSourceInstanceConfig NodeConfigurationSourceDetailsConfigurationSourceEnum = "INSTANCE_CONFIG"
)

var mappingNodeConfigurationSourceDetailsConfigurationSourceEnum = map[string]NodeConfigurationSourceDetailsConfigurationSourceEnum{
	"INSTANCE_CONFIG": NodeConfigurationSourceDetailsConfigurationSourceInstanceConfig,
}

var mappingNodeConfigurationSourceDetailsConfigurationSourceEnumLowerCase = map[string]NodeConfigurationSourceDetailsConfigurationSourceEnum{
	"instance_config": NodeConfigurationSourceDetailsConfigurationSourceInstanceConfig,
}

// GetNodeConfigurationSourceDetailsConfigurationSourceEnumValues Enumerates the set of values for NodeConfigurationSourceDetailsConfigurationSourceEnum
func GetNodeConfigurationSourceDetailsConfigurationSourceEnumValues() []NodeConfigurationSourceDetailsConfigurationSourceEnum {
	values := make([]NodeConfigurationSourceDetailsConfigurationSourceEnum, 0)
	for _, v := range mappingNodeConfigurationSourceDetailsConfigurationSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeConfigurationSourceDetailsConfigurationSourceEnumStringValues Enumerates the set of values in String for NodeConfigurationSourceDetailsConfigurationSourceEnum
func GetNodeConfigurationSourceDetailsConfigurationSourceEnumStringValues() []string {
	return []string{
		"INSTANCE_CONFIG",
	}
}

// GetMappingNodeConfigurationSourceDetailsConfigurationSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeConfigurationSourceDetailsConfigurationSourceEnum(val string) (NodeConfigurationSourceDetailsConfigurationSourceEnum, bool) {
	enum, ok := mappingNodeConfigurationSourceDetailsConfigurationSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
