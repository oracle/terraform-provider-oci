// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalClusterNetworkConfiguration The details of a network address configuration in an external cluster.
type ExternalClusterNetworkConfiguration struct {

	// The network number.
	NetworkNumber *int `mandatory:"false" json:"networkNumber"`

	// The network type.
	NetworkType ExternalClusterNetworkConfigurationNetworkTypeEnum `mandatory:"false" json:"networkType,omitempty"`

	// The subnet for the network.
	Subnet *string `mandatory:"false" json:"subnet"`
}

func (m ExternalClusterNetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalClusterNetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExternalClusterNetworkConfigurationNetworkTypeEnum(string(m.NetworkType)); !ok && m.NetworkType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkType: %s. Supported values are: %s.", m.NetworkType, strings.Join(GetExternalClusterNetworkConfigurationNetworkTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalClusterNetworkConfigurationNetworkTypeEnum Enum with underlying type: string
type ExternalClusterNetworkConfigurationNetworkTypeEnum string

// Set of constants representing the allowable values for ExternalClusterNetworkConfigurationNetworkTypeEnum
const (
	ExternalClusterNetworkConfigurationNetworkTypeAutoconfig ExternalClusterNetworkConfigurationNetworkTypeEnum = "AUTOCONFIG"
	ExternalClusterNetworkConfigurationNetworkTypeDhcp       ExternalClusterNetworkConfigurationNetworkTypeEnum = "DHCP"
	ExternalClusterNetworkConfigurationNetworkTypeStatic     ExternalClusterNetworkConfigurationNetworkTypeEnum = "STATIC"
	ExternalClusterNetworkConfigurationNetworkTypeMixed      ExternalClusterNetworkConfigurationNetworkTypeEnum = "MIXED"
)

var mappingExternalClusterNetworkConfigurationNetworkTypeEnum = map[string]ExternalClusterNetworkConfigurationNetworkTypeEnum{
	"AUTOCONFIG": ExternalClusterNetworkConfigurationNetworkTypeAutoconfig,
	"DHCP":       ExternalClusterNetworkConfigurationNetworkTypeDhcp,
	"STATIC":     ExternalClusterNetworkConfigurationNetworkTypeStatic,
	"MIXED":      ExternalClusterNetworkConfigurationNetworkTypeMixed,
}

var mappingExternalClusterNetworkConfigurationNetworkTypeEnumLowerCase = map[string]ExternalClusterNetworkConfigurationNetworkTypeEnum{
	"autoconfig": ExternalClusterNetworkConfigurationNetworkTypeAutoconfig,
	"dhcp":       ExternalClusterNetworkConfigurationNetworkTypeDhcp,
	"static":     ExternalClusterNetworkConfigurationNetworkTypeStatic,
	"mixed":      ExternalClusterNetworkConfigurationNetworkTypeMixed,
}

// GetExternalClusterNetworkConfigurationNetworkTypeEnumValues Enumerates the set of values for ExternalClusterNetworkConfigurationNetworkTypeEnum
func GetExternalClusterNetworkConfigurationNetworkTypeEnumValues() []ExternalClusterNetworkConfigurationNetworkTypeEnum {
	values := make([]ExternalClusterNetworkConfigurationNetworkTypeEnum, 0)
	for _, v := range mappingExternalClusterNetworkConfigurationNetworkTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalClusterNetworkConfigurationNetworkTypeEnumStringValues Enumerates the set of values in String for ExternalClusterNetworkConfigurationNetworkTypeEnum
func GetExternalClusterNetworkConfigurationNetworkTypeEnumStringValues() []string {
	return []string{
		"AUTOCONFIG",
		"DHCP",
		"STATIC",
		"MIXED",
	}
}

// GetMappingExternalClusterNetworkConfigurationNetworkTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalClusterNetworkConfigurationNetworkTypeEnum(val string) (ExternalClusterNetworkConfigurationNetworkTypeEnum, bool) {
	enum, ok := mappingExternalClusterNetworkConfigurationNetworkTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
