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

// ManagedComputeClusterJobNetworkConfiguration The network configuration details for the compute target job.
type ManagedComputeClusterJobNetworkConfiguration interface {
}

type managedcomputeclusterjobnetworkconfiguration struct {
	JsonData    []byte
	NetworkType string `json:"networkType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusterjobnetworkconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusterjobnetworkconfiguration managedcomputeclusterjobnetworkconfiguration
	s := struct {
		Model Unmarshalermanagedcomputeclusterjobnetworkconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkType = s.Model.NetworkType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusterjobnetworkconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkType {
	case "DEFAULT_NETWORK":
		mm := ManagedComputeClusterJobDefaultNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_NETWORK":
		mm := ManagedComputeClusterJobCustomNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterJobNetworkConfiguration: %s.", m.NetworkType)
		return *m, nil
	}
}

func (m managedcomputeclusterjobnetworkconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusterjobnetworkconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum Enum with underlying type: string
type ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum
const (
	ManagedComputeClusterJobNetworkConfigurationNetworkTypeCustomNetwork  ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum = "CUSTOM_NETWORK"
	ManagedComputeClusterJobNetworkConfigurationNetworkTypeDefaultNetwork ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum = "DEFAULT_NETWORK"
)

var mappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum = map[string]ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum{
	"CUSTOM_NETWORK":  ManagedComputeClusterJobNetworkConfigurationNetworkTypeCustomNetwork,
	"DEFAULT_NETWORK": ManagedComputeClusterJobNetworkConfigurationNetworkTypeDefaultNetwork,
}

var mappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumLowerCase = map[string]ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum{
	"custom_network":  ManagedComputeClusterJobNetworkConfigurationNetworkTypeCustomNetwork,
	"default_network": ManagedComputeClusterJobNetworkConfigurationNetworkTypeDefaultNetwork,
}

// GetManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumValues Enumerates the set of values for ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum
func GetManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumValues() []ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum {
	values := make([]ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum
func GetManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_NETWORK",
		"DEFAULT_NETWORK",
	}
}

// GetMappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum(val string) (ManagedComputeClusterJobNetworkConfigurationNetworkTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterJobNetworkConfigurationNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
