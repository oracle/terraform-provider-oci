// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportedPackage An imported/instantiated package within an instance.
type ImportedPackage struct {

	// ID of the host instance.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// ID of the package.
	CurrentPackageId *string `mandatory:"true" json:"currentPackageId"`

	// Stable name of the package (the same across versions).
	Name *string `mandatory:"true" json:"name"`

	// Display name of the package (can change across versions).
	DisplayName *string `mandatory:"true" json:"displayName"`

	// version of the package.
	Version *string `mandatory:"true" json:"version"`

	// Status of the imported package.
	Status ImportedPackageStatusEnum `mandatory:"true" json:"status"`

	// When the imported package was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// When the imported package was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Short message explaining the status of this imported package.
	StatusMessage *string `mandatory:"true" json:"statusMessage"`

	// A list of parameter values used to import the package.
	ParameterValues map[string]string `mandatory:"true" json:"parameterValues"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ImportedPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportedPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportedPackageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetImportedPackageStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportedPackageStatusEnum Enum with underlying type: string
type ImportedPackageStatusEnum string

// Set of constants representing the allowable values for ImportedPackageStatusEnum
const (
	ImportedPackageStatusReady            ImportedPackageStatusEnum = "READY"
	ImportedPackageStatusOperationPending ImportedPackageStatusEnum = "OPERATION_PENDING"
	ImportedPackageStatusFailed           ImportedPackageStatusEnum = "FAILED"
)

var mappingImportedPackageStatusEnum = map[string]ImportedPackageStatusEnum{
	"READY":             ImportedPackageStatusReady,
	"OPERATION_PENDING": ImportedPackageStatusOperationPending,
	"FAILED":            ImportedPackageStatusFailed,
}

var mappingImportedPackageStatusEnumLowerCase = map[string]ImportedPackageStatusEnum{
	"ready":             ImportedPackageStatusReady,
	"operation_pending": ImportedPackageStatusOperationPending,
	"failed":            ImportedPackageStatusFailed,
}

// GetImportedPackageStatusEnumValues Enumerates the set of values for ImportedPackageStatusEnum
func GetImportedPackageStatusEnumValues() []ImportedPackageStatusEnum {
	values := make([]ImportedPackageStatusEnum, 0)
	for _, v := range mappingImportedPackageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetImportedPackageStatusEnumStringValues Enumerates the set of values in String for ImportedPackageStatusEnum
func GetImportedPackageStatusEnumStringValues() []string {
	return []string{
		"READY",
		"OPERATION_PENDING",
		"FAILED",
	}
}

// GetMappingImportedPackageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportedPackageStatusEnum(val string) (ImportedPackageStatusEnum, bool) {
	enum, ok := mappingImportedPackageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
