// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceConfiguration The model deployment instance configuration.
type InstanceConfiguration struct {

	// The shape used to launch the model deployment instances.
	// When using service managed open source foundation model, the supported shapes can be retrieved using get model api /models/{modelId}/definedMetadata/deploymentConfiguration/artifact/content.
	InstanceShapeName *string `mandatory:"true" json:"instanceShapeName"`

	ModelDeploymentInstanceShapeConfigDetails *ModelDeploymentInstanceShapeConfigDetails `mandatory:"false" json:"modelDeploymentInstanceShapeConfigDetails"`

	// A model deployment instance is provided with a VNIC for network access.  This specifies the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT/SGW gateway for egress.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID of a Data Science private endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// Network Access type of model deployment.
	NetworkAccessType InstanceConfigurationNetworkAccessTypeEnum `mandatory:"false" json:"networkAccessType,omitempty"`
}

func (m InstanceConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstanceConfigurationNetworkAccessTypeEnum(string(m.NetworkAccessType)); !ok && m.NetworkAccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkAccessType: %s. Supported values are: %s.", m.NetworkAccessType, strings.Join(GetInstanceConfigurationNetworkAccessTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceConfigurationNetworkAccessTypeEnum Enum with underlying type: string
type InstanceConfigurationNetworkAccessTypeEnum string

// Set of constants representing the allowable values for InstanceConfigurationNetworkAccessTypeEnum
const (
	InstanceConfigurationNetworkAccessTypeManagedNetworkingNoInternetAccess InstanceConfigurationNetworkAccessTypeEnum = "MANAGED_NETWORKING_NO_INTERNET_ACCESS"
	InstanceConfigurationNetworkAccessTypeManagedNetworkingInternetAccess   InstanceConfigurationNetworkAccessTypeEnum = "MANAGED_NETWORKING_INTERNET_ACCESS"
	InstanceConfigurationNetworkAccessTypeCustomNetworking                  InstanceConfigurationNetworkAccessTypeEnum = "CUSTOM_NETWORKING"
)

var mappingInstanceConfigurationNetworkAccessTypeEnum = map[string]InstanceConfigurationNetworkAccessTypeEnum{
	"MANAGED_NETWORKING_NO_INTERNET_ACCESS": InstanceConfigurationNetworkAccessTypeManagedNetworkingNoInternetAccess,
	"MANAGED_NETWORKING_INTERNET_ACCESS":    InstanceConfigurationNetworkAccessTypeManagedNetworkingInternetAccess,
	"CUSTOM_NETWORKING":                     InstanceConfigurationNetworkAccessTypeCustomNetworking,
}

var mappingInstanceConfigurationNetworkAccessTypeEnumLowerCase = map[string]InstanceConfigurationNetworkAccessTypeEnum{
	"managed_networking_no_internet_access": InstanceConfigurationNetworkAccessTypeManagedNetworkingNoInternetAccess,
	"managed_networking_internet_access":    InstanceConfigurationNetworkAccessTypeManagedNetworkingInternetAccess,
	"custom_networking":                     InstanceConfigurationNetworkAccessTypeCustomNetworking,
}

// GetInstanceConfigurationNetworkAccessTypeEnumValues Enumerates the set of values for InstanceConfigurationNetworkAccessTypeEnum
func GetInstanceConfigurationNetworkAccessTypeEnumValues() []InstanceConfigurationNetworkAccessTypeEnum {
	values := make([]InstanceConfigurationNetworkAccessTypeEnum, 0)
	for _, v := range mappingInstanceConfigurationNetworkAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationNetworkAccessTypeEnumStringValues Enumerates the set of values in String for InstanceConfigurationNetworkAccessTypeEnum
func GetInstanceConfigurationNetworkAccessTypeEnumStringValues() []string {
	return []string{
		"MANAGED_NETWORKING_NO_INTERNET_ACCESS",
		"MANAGED_NETWORKING_INTERNET_ACCESS",
		"CUSTOM_NETWORKING",
	}
}

// GetMappingInstanceConfigurationNetworkAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationNetworkAccessTypeEnum(val string) (InstanceConfigurationNetworkAccessTypeEnum, bool) {
	enum, ok := mappingInstanceConfigurationNetworkAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
