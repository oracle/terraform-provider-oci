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

// DiscoverExternalExadataInfrastructureDetails The connection details and the discovery options for the Exadata discovery.
type DiscoverExternalExadataInfrastructureDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of discovery.
	DiscoveryType DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum `mandatory:"true" json:"discoveryType"`

	// The list of the DB system identifiers.
	DbSystemIds []string `mandatory:"true" json:"dbSystemIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure. This is applicable for rediscovery only.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`
}

func (m DiscoverExternalExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoverExternalExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum(string(m.DiscoveryType)); !ok && m.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", m.DiscoveryType, strings.Join(GetDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum Enum with underlying type: string
type DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum string

// Set of constants representing the allowable values for DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum
const (
	DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeNew      DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum = "NEW"
	DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeOverride DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum = "OVERRIDE"
)

var mappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum = map[string]DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum{
	"NEW":      DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeNew,
	"OVERRIDE": DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeOverride,
}

var mappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumLowerCase = map[string]DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum{
	"new":      DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeNew,
	"override": DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeOverride,
}

// GetDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumValues Enumerates the set of values for DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum
func GetDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumValues() []DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum {
	values := make([]DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum, 0)
	for _, v := range mappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumStringValues Enumerates the set of values in String for DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum
func GetDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumStringValues() []string {
	return []string{
		"NEW",
		"OVERRIDE",
	}
}

// GetMappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum(val string) (DiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnum, bool) {
	enum, ok := mappingDiscoverExternalExadataInfrastructureDetailsDiscoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
