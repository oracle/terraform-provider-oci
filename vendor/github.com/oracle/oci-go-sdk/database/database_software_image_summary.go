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

// DatabaseSoftwareImageSummary The Database service supports the creation of database software images for use in creating and patching DB systems and databases.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DatabaseSoftwareImageSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database software image.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database version with which the database software image is to be built.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the database software image.
	LifecycleState DatabaseSoftwareImageSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database software image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of software image. Can be grid or database.
	ImageType DatabaseSoftwareImageSummaryImageTypeEnum `mandatory:"true" json:"imageType"`

	// To what shape the image is meant for.
	ImageShapeFamily DatabaseSoftwareImageSummaryImageShapeFamilyEnum `mandatory:"true" json:"imageShapeFamily"`

	// The PSU or PBP or Release Updates. To get a list of supported versions, use the ListDbVersions operation.
	PatchSet *string `mandatory:"true" json:"patchSet"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageIncludedPatches []string `mandatory:"false" json:"databaseSoftwareImageIncludedPatches"`

	// The patches included in the image and the version of the image
	IncludedPatchesSummary *string `mandatory:"false" json:"includedPatchesSummary"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageOneOffPatches []string `mandatory:"false" json:"databaseSoftwareImageOneOffPatches"`

	// output from lsinventory which will get passed as a string
	LsInventory *string `mandatory:"false" json:"lsInventory"`
}

func (m DatabaseSoftwareImageSummary) String() string {
	return common.PointerString(m)
}

// DatabaseSoftwareImageSummaryLifecycleStateEnum Enum with underlying type: string
type DatabaseSoftwareImageSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageSummaryLifecycleStateEnum
const (
	DatabaseSoftwareImageSummaryLifecycleStateProvisioning DatabaseSoftwareImageSummaryLifecycleStateEnum = "PROVISIONING"
	DatabaseSoftwareImageSummaryLifecycleStateAvailable    DatabaseSoftwareImageSummaryLifecycleStateEnum = "AVAILABLE"
	DatabaseSoftwareImageSummaryLifecycleStateDeleting     DatabaseSoftwareImageSummaryLifecycleStateEnum = "DELETING"
	DatabaseSoftwareImageSummaryLifecycleStateDeleted      DatabaseSoftwareImageSummaryLifecycleStateEnum = "DELETED"
	DatabaseSoftwareImageSummaryLifecycleStateFailed       DatabaseSoftwareImageSummaryLifecycleStateEnum = "FAILED"
	DatabaseSoftwareImageSummaryLifecycleStateUpdating     DatabaseSoftwareImageSummaryLifecycleStateEnum = "UPDATING"
)

var mappingDatabaseSoftwareImageSummaryLifecycleState = map[string]DatabaseSoftwareImageSummaryLifecycleStateEnum{
	"PROVISIONING": DatabaseSoftwareImageSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DatabaseSoftwareImageSummaryLifecycleStateAvailable,
	"DELETING":     DatabaseSoftwareImageSummaryLifecycleStateDeleting,
	"DELETED":      DatabaseSoftwareImageSummaryLifecycleStateDeleted,
	"FAILED":       DatabaseSoftwareImageSummaryLifecycleStateFailed,
	"UPDATING":     DatabaseSoftwareImageSummaryLifecycleStateUpdating,
}

// GetDatabaseSoftwareImageSummaryLifecycleStateEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryLifecycleStateEnum
func GetDatabaseSoftwareImageSummaryLifecycleStateEnumValues() []DatabaseSoftwareImageSummaryLifecycleStateEnum {
	values := make([]DatabaseSoftwareImageSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// DatabaseSoftwareImageSummaryImageTypeEnum Enum with underlying type: string
type DatabaseSoftwareImageSummaryImageTypeEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageSummaryImageTypeEnum
const (
	DatabaseSoftwareImageSummaryImageTypeGridImage     DatabaseSoftwareImageSummaryImageTypeEnum = "GRID_IMAGE"
	DatabaseSoftwareImageSummaryImageTypeDatabaseImage DatabaseSoftwareImageSummaryImageTypeEnum = "DATABASE_IMAGE"
)

var mappingDatabaseSoftwareImageSummaryImageType = map[string]DatabaseSoftwareImageSummaryImageTypeEnum{
	"GRID_IMAGE":     DatabaseSoftwareImageSummaryImageTypeGridImage,
	"DATABASE_IMAGE": DatabaseSoftwareImageSummaryImageTypeDatabaseImage,
}

// GetDatabaseSoftwareImageSummaryImageTypeEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryImageTypeEnum
func GetDatabaseSoftwareImageSummaryImageTypeEnumValues() []DatabaseSoftwareImageSummaryImageTypeEnum {
	values := make([]DatabaseSoftwareImageSummaryImageTypeEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryImageType {
		values = append(values, v)
	}
	return values
}

// DatabaseSoftwareImageSummaryImageShapeFamilyEnum Enum with underlying type: string
type DatabaseSoftwareImageSummaryImageShapeFamilyEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageSummaryImageShapeFamilyEnum
const (
	DatabaseSoftwareImageSummaryImageShapeFamilyVmBmShape    DatabaseSoftwareImageSummaryImageShapeFamilyEnum = "VM_BM_SHAPE"
	DatabaseSoftwareImageSummaryImageShapeFamilyExadataShape DatabaseSoftwareImageSummaryImageShapeFamilyEnum = "EXADATA_SHAPE"
)

var mappingDatabaseSoftwareImageSummaryImageShapeFamily = map[string]DatabaseSoftwareImageSummaryImageShapeFamilyEnum{
	"VM_BM_SHAPE":   DatabaseSoftwareImageSummaryImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": DatabaseSoftwareImageSummaryImageShapeFamilyExadataShape,
}

// GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryImageShapeFamilyEnum
func GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumValues() []DatabaseSoftwareImageSummaryImageShapeFamilyEnum {
	values := make([]DatabaseSoftwareImageSummaryImageShapeFamilyEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryImageShapeFamily {
		values = append(values, v)
	}
	return values
}
