// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdentityConnectorDetails Details of the identity connector
type IdentityConnectorDetails struct {

	// The OCID of the identity connector
	Id *string `mandatory:"false" json:"id"`

	// Cloud provider
	CloudProvider IdentityConnectorDetailsCloudProviderEnum `mandatory:"false" json:"cloudProvider,omitempty"`
}

func (m IdentityConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIdentityConnectorDetailsCloudProviderEnum(string(m.CloudProvider)); !ok && m.CloudProvider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloudProvider: %s. Supported values are: %s.", m.CloudProvider, strings.Join(GetIdentityConnectorDetailsCloudProviderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityConnectorDetailsCloudProviderEnum Enum with underlying type: string
type IdentityConnectorDetailsCloudProviderEnum string

// Set of constants representing the allowable values for IdentityConnectorDetailsCloudProviderEnum
const (
	IdentityConnectorDetailsCloudProviderAzure IdentityConnectorDetailsCloudProviderEnum = "AZURE"
	IdentityConnectorDetailsCloudProviderGcp   IdentityConnectorDetailsCloudProviderEnum = "GCP"
)

var mappingIdentityConnectorDetailsCloudProviderEnum = map[string]IdentityConnectorDetailsCloudProviderEnum{
	"AZURE": IdentityConnectorDetailsCloudProviderAzure,
	"GCP":   IdentityConnectorDetailsCloudProviderGcp,
}

var mappingIdentityConnectorDetailsCloudProviderEnumLowerCase = map[string]IdentityConnectorDetailsCloudProviderEnum{
	"azure": IdentityConnectorDetailsCloudProviderAzure,
	"gcp":   IdentityConnectorDetailsCloudProviderGcp,
}

// GetIdentityConnectorDetailsCloudProviderEnumValues Enumerates the set of values for IdentityConnectorDetailsCloudProviderEnum
func GetIdentityConnectorDetailsCloudProviderEnumValues() []IdentityConnectorDetailsCloudProviderEnum {
	values := make([]IdentityConnectorDetailsCloudProviderEnum, 0)
	for _, v := range mappingIdentityConnectorDetailsCloudProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityConnectorDetailsCloudProviderEnumStringValues Enumerates the set of values in String for IdentityConnectorDetailsCloudProviderEnum
func GetIdentityConnectorDetailsCloudProviderEnumStringValues() []string {
	return []string{
		"AZURE",
		"GCP",
	}
}

// GetMappingIdentityConnectorDetailsCloudProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityConnectorDetailsCloudProviderEnum(val string) (IdentityConnectorDetailsCloudProviderEnum, bool) {
	enum, ok := mappingIdentityConnectorDetailsCloudProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
