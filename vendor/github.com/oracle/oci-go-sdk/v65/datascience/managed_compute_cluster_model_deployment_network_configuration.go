// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManagedComputeClusterModelDeploymentNetworkConfiguration The network configuration details for model deploy on managed compute cluster type compute target.
type ManagedComputeClusterModelDeploymentNetworkConfiguration interface {
}

type managedcomputeclustermodeldeploymentnetworkconfiguration struct {
	JsonData    []byte
	NetworkType string `json:"networkType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclustermodeldeploymentnetworkconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclustermodeldeploymentnetworkconfiguration managedcomputeclustermodeldeploymentnetworkconfiguration
	s := struct {
		Model Unmarshalermanagedcomputeclustermodeldeploymentnetworkconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkType = s.Model.NetworkType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclustermodeldeploymentnetworkconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkType {
	case "CUSTOM_NETWORK":
		mm := ManagedComputeClusterModelDeploymentCustomNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFAULT_NETWORK":
		mm := ManagedComputeClusterModelDeploymentDefaultNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterModelDeploymentNetworkConfiguration: %s.", m.NetworkType)
		return *m, nil
	}
}

func (m managedcomputeclustermodeldeploymentnetworkconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclustermodeldeploymentnetworkconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum Enum with underlying type: string
type ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum
const (
	ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeCustomNetwork  ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum = "CUSTOM_NETWORK"
	ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeDefaultNetwork ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum = "DEFAULT_NETWORK"
)

var mappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum = map[string]ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum{
	"CUSTOM_NETWORK":  ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeCustomNetwork,
	"DEFAULT_NETWORK": ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeDefaultNetwork,
}

var mappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumLowerCase = map[string]ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum{
	"custom_network":  ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeCustomNetwork,
	"default_network": ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeDefaultNetwork,
}

// GetManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumValues Enumerates the set of values for ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum
func GetManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumValues() []ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum {
	values := make([]ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum
func GetManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_NETWORK",
		"DEFAULT_NETWORK",
	}
}

// GetMappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum(val string) (ManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterModelDeploymentNetworkConfigurationNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
