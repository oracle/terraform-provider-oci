// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDbConnectionBundleDetails The details for creating a database connection bundle.
type CreateDbConnectionBundleDetails struct {

	// The OCID of the compartment where you want to create the database connection bundle.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name for the connection bundle.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the database connection bundle.
	DbConnectionBundleType CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum `mandatory:"true" json:"dbConnectionBundleType"`

	// Details about the resources to associate with the connection bundle.
	AssociatedResourceDetails []AssociatedResourceDetails `mandatory:"true" json:"associatedResourceDetails"`

	// The tenancy ID for the connection bundle.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]string `mandatory:"false" json:"definedTags"`
}

func (m CreateDbConnectionBundleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbConnectionBundleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum(string(m.DbConnectionBundleType)); !ok && m.DbConnectionBundleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbConnectionBundleType: %s. Supported values are: %s.", m.DbConnectionBundleType, strings.Join(GetCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum Enum with underlying type: string
type CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum string

// Set of constants representing the allowable values for CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum
const (
	CreateDbConnectionBundleDetailsDbConnectionBundleTypeTls  CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum = "TLS"
	CreateDbConnectionBundleDetailsDbConnectionBundleTypeMtls CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum = "MTLS"
)

var mappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum = map[string]CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum{
	"TLS":  CreateDbConnectionBundleDetailsDbConnectionBundleTypeTls,
	"MTLS": CreateDbConnectionBundleDetailsDbConnectionBundleTypeMtls,
}

var mappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumLowerCase = map[string]CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum{
	"tls":  CreateDbConnectionBundleDetailsDbConnectionBundleTypeTls,
	"mtls": CreateDbConnectionBundleDetailsDbConnectionBundleTypeMtls,
}

// GetCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumValues Enumerates the set of values for CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum
func GetCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumValues() []CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum {
	values := make([]CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum, 0)
	for _, v := range mappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumStringValues Enumerates the set of values in String for CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum
func GetCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumStringValues() []string {
	return []string{
		"TLS",
		"MTLS",
	}
}

// GetMappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum(val string) (CreateDbConnectionBundleDetailsDbConnectionBundleTypeEnum, bool) {
	enum, ok := mappingCreateDbConnectionBundleDetailsDbConnectionBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
