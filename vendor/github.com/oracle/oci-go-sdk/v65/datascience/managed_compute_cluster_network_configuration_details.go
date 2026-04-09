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

// ManagedComputeClusterNetworkConfigurationDetails The network configuration details for managed compute cluster type resource.
type ManagedComputeClusterNetworkConfigurationDetails interface {
}

type managedcomputeclusternetworkconfigurationdetails struct {
	JsonData    []byte
	NetworkType string `json:"networkType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusternetworkconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusternetworkconfigurationdetails managedcomputeclusternetworkconfigurationdetails
	s := struct {
		Model Unmarshalermanagedcomputeclusternetworkconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkType = s.Model.NetworkType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusternetworkconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkType {
	case "DEFAULT_NETWORK":
		mm := ManagedComputeClusterDefaultNetworkConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_NETWORK":
		mm := ManagedComputeClusterCustomNetworkConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterNetworkConfigurationDetails: %s.", m.NetworkType)
		return *m, nil
	}
}

func (m managedcomputeclusternetworkconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusternetworkconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum Enum with underlying type: string
type ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum
const (
	ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeCustomNetwork  ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum = "CUSTOM_NETWORK"
	ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeDefaultNetwork ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum = "DEFAULT_NETWORK"
)

var mappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum = map[string]ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum{
	"CUSTOM_NETWORK":  ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeCustomNetwork,
	"DEFAULT_NETWORK": ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeDefaultNetwork,
}

var mappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumLowerCase = map[string]ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum{
	"custom_network":  ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeCustomNetwork,
	"default_network": ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeDefaultNetwork,
}

// GetManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumValues Enumerates the set of values for ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum
func GetManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumValues() []ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum {
	values := make([]ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum
func GetManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_NETWORK",
		"DEFAULT_NETWORK",
	}
}

// GetMappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum(val string) (ManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterNetworkConfigurationDetailsNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
