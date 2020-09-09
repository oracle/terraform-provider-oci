// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDatabaseSoftwareImageDetails Parameters for creating a database software image in the specified compartment.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseSoftwareImageDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment the database software image  belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database version with which the database software image is to be built.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The PSU or PBP or Release Updates. To get a list of supported versions, use the ListDbVersions operation.
	PatchSet *string `mandatory:"true" json:"patchSet"`

	// To what shape the image is meant for.
	ImageShapeFamily CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum `mandatory:"false" json:"imageShapeFamily,omitempty"`

	// List of the Fault Domains in which this DB system is provisioned.
	ImageType CreateDatabaseSoftwareImageDetailsImageTypeEnum `mandatory:"false" json:"imageType,omitempty"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageOneOffPatches []string `mandatory:"false" json:"databaseSoftwareImageOneOffPatches"`

	// output from lsinventory which will get passed as a string
	LsInventory *string `mandatory:"false" json:"lsInventory"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDatabaseSoftwareImageDetails) String() string {
	return common.PointerString(m)
}

// CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum Enum with underlying type: string
type CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum string

// Set of constants representing the allowable values for CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum
const (
	CreateDatabaseSoftwareImageDetailsImageShapeFamilyVmBmShape    CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "VM_BM_SHAPE"
	CreateDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum = "EXADATA_SHAPE"
)

var mappingCreateDatabaseSoftwareImageDetailsImageShapeFamily = map[string]CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum{
	"VM_BM_SHAPE":   CreateDatabaseSoftwareImageDetailsImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": CreateDatabaseSoftwareImageDetailsImageShapeFamilyExadataShape,
}

// GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues Enumerates the set of values for CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum
func GetCreateDatabaseSoftwareImageDetailsImageShapeFamilyEnumValues() []CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum {
	values := make([]CreateDatabaseSoftwareImageDetailsImageShapeFamilyEnum, 0)
	for _, v := range mappingCreateDatabaseSoftwareImageDetailsImageShapeFamily {
		values = append(values, v)
	}
	return values
}

// CreateDatabaseSoftwareImageDetailsImageTypeEnum Enum with underlying type: string
type CreateDatabaseSoftwareImageDetailsImageTypeEnum string

// Set of constants representing the allowable values for CreateDatabaseSoftwareImageDetailsImageTypeEnum
const (
	CreateDatabaseSoftwareImageDetailsImageTypeGridImage     CreateDatabaseSoftwareImageDetailsImageTypeEnum = "GRID_IMAGE"
	CreateDatabaseSoftwareImageDetailsImageTypeDatabaseImage CreateDatabaseSoftwareImageDetailsImageTypeEnum = "DATABASE_IMAGE"
)

var mappingCreateDatabaseSoftwareImageDetailsImageType = map[string]CreateDatabaseSoftwareImageDetailsImageTypeEnum{
	"GRID_IMAGE":     CreateDatabaseSoftwareImageDetailsImageTypeGridImage,
	"DATABASE_IMAGE": CreateDatabaseSoftwareImageDetailsImageTypeDatabaseImage,
}

// GetCreateDatabaseSoftwareImageDetailsImageTypeEnumValues Enumerates the set of values for CreateDatabaseSoftwareImageDetailsImageTypeEnum
func GetCreateDatabaseSoftwareImageDetailsImageTypeEnumValues() []CreateDatabaseSoftwareImageDetailsImageTypeEnum {
	values := make([]CreateDatabaseSoftwareImageDetailsImageTypeEnum, 0)
	for _, v := range mappingCreateDatabaseSoftwareImageDetailsImageType {
		values = append(values, v)
	}
	return values
}
