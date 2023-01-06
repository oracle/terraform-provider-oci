// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatabaseSoftwareImageDetails Parameters for creating a database software image in the specified compartment.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseSoftwareImageDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment the database software image  belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The database version with which the database software image is to be built.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// To what shape the image is meant for.
	ImageShapeFamily CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum `mandatory:"false" json:"imageShapeFamily,omitempty"`

	// The type of software image. Can be grid or database.
	ImageType CreateDatabaseSoftwareImageDetailsImageTypeEnum `mandatory:"false" json:"imageType,omitempty"`

	// The PSU or PBP or Release Updates. To get a list of supported versions, use the ListDbVersions operation.
	PatchSet *string `mandatory:"false" json:"patchSet"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageOneOffPatches []string `mandatory:"false" json:"databaseSoftwareImageOneOffPatches"`

	// The output from the OPatch lsInventory command, which is passed as a string.
	LsInventory *string `mandatory:"false" json:"lsInventory"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	SourceDbHomeId *string `mandatory:"false" json:"sourceDbHomeId"`
}

func (m CreateDatabaseSoftwareImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseSoftwareImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum(string(m.ImageShapeFamily)); !ok && m.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", m.ImageShapeFamily, strings.Join(GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDatabaseSoftwareImageDetailsImageTypeEnum(string(m.ImageType)); !ok && m.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", m.ImageType, strings.Join(GetCreateDatabaseSoftwareImageDetailsImageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum Enum with underlying type: string
type CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum string

// Set of constants representing the allowable values for CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum
const (
	CreateDatabaseSoftwareImageDetailsImageShapeFamilyVmBmShape    CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "VM_BM_SHAPE"
	CreateDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "EXADATA_SHAPE"
	CreateDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape   CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "EXACC_SHAPE"
)

var mappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = map[string]CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum{
	"VM_BM_SHAPE":   CreateDatabaseSoftwareImageDetailsImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": CreateDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape,
	"EXACC_SHAPE":   CreateDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape,
}

var mappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumLowerCase = map[string]CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum{
	"vm_bm_shape":   CreateDatabaseSoftwareImageDetailsImageShapeFamilyVmBmShape,
	"exadata_shape": CreateDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape,
	"exacc_shape":   CreateDatabaseSoftwareImageDetailsImageShapeFamilyExaccShape,
}

// GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues Enumerates the set of values for CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum
func GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues() []CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum {
	values := make([]CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum, 0)
	for _, v := range mappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues Enumerates the set of values in String for CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum
func GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumStringValues() []string {
	return []string{
		"VM_BM_SHAPE",
		"EXADATA_SHAPE",
		"EXACC_SHAPE",
	}
}

// GetMappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum(val string) (CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum, bool) {
	enum, ok := mappingCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDatabaseSoftwareImageDetailsImageTypeEnum Enum with underlying type: string
type CreateDatabaseSoftwareImageDetailsImageTypeEnum string

// Set of constants representing the allowable values for CreateDatabaseSoftwareImageDetailsImageTypeEnum
const (
	CreateDatabaseSoftwareImageDetailsImageTypeGridImage     CreateDatabaseSoftwareImageDetailsImageTypeEnum = "GRID_IMAGE"
	CreateDatabaseSoftwareImageDetailsImageTypeDatabaseImage CreateDatabaseSoftwareImageDetailsImageTypeEnum = "DATABASE_IMAGE"
)

var mappingCreateDatabaseSoftwareImageDetailsImageTypeEnum = map[string]CreateDatabaseSoftwareImageDetailsImageTypeEnum{
	"GRID_IMAGE":     CreateDatabaseSoftwareImageDetailsImageTypeGridImage,
	"DATABASE_IMAGE": CreateDatabaseSoftwareImageDetailsImageTypeDatabaseImage,
}

var mappingCreateDatabaseSoftwareImageDetailsImageTypeEnumLowerCase = map[string]CreateDatabaseSoftwareImageDetailsImageTypeEnum{
	"grid_image":     CreateDatabaseSoftwareImageDetailsImageTypeGridImage,
	"database_image": CreateDatabaseSoftwareImageDetailsImageTypeDatabaseImage,
}

// GetCreateDatabaseSoftwareImageDetailsImageTypeEnumValues Enumerates the set of values for CreateDatabaseSoftwareImageDetailsImageTypeEnum
func GetCreateDatabaseSoftwareImageDetailsImageTypeEnumValues() []CreateDatabaseSoftwareImageDetailsImageTypeEnum {
	values := make([]CreateDatabaseSoftwareImageDetailsImageTypeEnum, 0)
	for _, v := range mappingCreateDatabaseSoftwareImageDetailsImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseSoftwareImageDetailsImageTypeEnumStringValues Enumerates the set of values in String for CreateDatabaseSoftwareImageDetailsImageTypeEnum
func GetCreateDatabaseSoftwareImageDetailsImageTypeEnumStringValues() []string {
	return []string{
		"GRID_IMAGE",
		"DATABASE_IMAGE",
	}
}

// GetMappingCreateDatabaseSoftwareImageDetailsImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseSoftwareImageDetailsImageTypeEnum(val string) (CreateDatabaseSoftwareImageDetailsImageTypeEnum, bool) {
	enum, ok := mappingCreateDatabaseSoftwareImageDetailsImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
