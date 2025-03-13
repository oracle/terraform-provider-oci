// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSoftwareImageDetails Detailed specification about the new SoftwareImage.
type CreateSoftwareImageDetails struct {

	// SoftwareImage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of Software Image
	ImageType CreateSoftwareImageDetailsImageTypeEnum `mandatory:"true" json:"imageType"`

	// Name of the first software image version
	VersionName *string `mandatory:"true" json:"versionName"`

	// Release for the software image version
	Release *string `mandatory:"true" json:"release"`

	MosCredentials *ValidateMosCredentialDetails `mandatory:"true" json:"mosCredentials"`

	// Description of SoftwareImage
	Description *string `mandatory:"false" json:"description"`

	// One Offs for the software image version
	OneOffs []int `mandatory:"false" json:"oneOffs"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSoftwareImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSoftwareImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSoftwareImageDetailsImageTypeEnum(string(m.ImageType)); !ok && m.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", m.ImageType, strings.Join(GetCreateSoftwareImageDetailsImageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSoftwareImageDetailsImageTypeEnum Enum with underlying type: string
type CreateSoftwareImageDetailsImageTypeEnum string

// Set of constants representing the allowable values for CreateSoftwareImageDetailsImageTypeEnum
const (
	CreateSoftwareImageDetailsImageTypeDb CreateSoftwareImageDetailsImageTypeEnum = "DB"
)

var mappingCreateSoftwareImageDetailsImageTypeEnum = map[string]CreateSoftwareImageDetailsImageTypeEnum{
	"DB": CreateSoftwareImageDetailsImageTypeDb,
}

var mappingCreateSoftwareImageDetailsImageTypeEnumLowerCase = map[string]CreateSoftwareImageDetailsImageTypeEnum{
	"db": CreateSoftwareImageDetailsImageTypeDb,
}

// GetCreateSoftwareImageDetailsImageTypeEnumValues Enumerates the set of values for CreateSoftwareImageDetailsImageTypeEnum
func GetCreateSoftwareImageDetailsImageTypeEnumValues() []CreateSoftwareImageDetailsImageTypeEnum {
	values := make([]CreateSoftwareImageDetailsImageTypeEnum, 0)
	for _, v := range mappingCreateSoftwareImageDetailsImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSoftwareImageDetailsImageTypeEnumStringValues Enumerates the set of values in String for CreateSoftwareImageDetailsImageTypeEnum
func GetCreateSoftwareImageDetailsImageTypeEnumStringValues() []string {
	return []string{
		"DB",
	}
}

// GetMappingCreateSoftwareImageDetailsImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSoftwareImageDetailsImageTypeEnum(val string) (CreateSoftwareImageDetailsImageTypeEnum, bool) {
	enum, ok := mappingCreateSoftwareImageDetailsImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
