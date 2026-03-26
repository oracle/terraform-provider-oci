// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoverCloudExadataInfrastructureDetails The connection details and the discovery options for the Exadata discovery.
type DiscoverCloudExadataInfrastructureDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of discovery.
	DiscoveryType DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum `mandatory:"true" json:"discoveryType"`

	// The list of VM Clusters in the Exadata infrastructure.
	VmClusterIds []string `mandatory:"true" json:"vmClusterIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure. This is applicable for rediscovery only.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`
}

func (m DiscoverCloudExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoverCloudExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum Enum with underlying type: string
type DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum
const (
	DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeNew      DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum = "NEW"
	DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeOverride DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum = "OVERRIDE"
)

var mappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum = map[string]DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum{
	"NEW":      DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeNew,
	"OVERRIDE": DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeOverride,
}

var mappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumLowerCase = map[string]DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum{
	"new":      DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeNew,
	"override": DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeOverride,
}

// GetDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumValues Enumerates the set of values for DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum
func GetDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumValues() []DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum {
	values := make([]DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum
func GetDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumStringValues() []string {
	return []string{
		"NEW",
		"OVERRIDE",
	}
}

// GetMappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum(val string) (DiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoverCloudExadataInfrastructureDetailsDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
