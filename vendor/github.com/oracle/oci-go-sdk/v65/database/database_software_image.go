// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseSoftwareImage Database software images are created by specifying a patch set, one-off patches and patches for the database home (listed by `ls inventory`).
type DatabaseSoftwareImage struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database software image.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database version with which the database software image is to be built.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the database software image.
	LifecycleState DatabaseSoftwareImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database software image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of software image. Can be grid or database.
	ImageType DatabaseSoftwareImageImageTypeEnum `mandatory:"true" json:"imageType"`

	// To what shape the image is meant for.
	ImageShapeFamily DatabaseSoftwareImageImageShapeFamilyEnum `mandatory:"true" json:"imageShapeFamily"`

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

	// True if this Database software image is supported for Upgrade.
	IsUpgradeSupported *bool `mandatory:"false" json:"isUpgradeSupported"`
}

func (m DatabaseSoftwareImage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseSoftwareImage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSoftwareImageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseSoftwareImageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageImageTypeEnum(string(m.ImageType)); !ok && m.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", m.ImageType, strings.Join(GetDatabaseSoftwareImageImageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageImageShapeFamilyEnum(string(m.ImageShapeFamily)); !ok && m.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", m.ImageShapeFamily, strings.Join(GetDatabaseSoftwareImageImageShapeFamilyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseSoftwareImageLifecycleStateEnum Enum with underlying type: string
type DatabaseSoftwareImageLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageLifecycleStateEnum
const (
	DatabaseSoftwareImageLifecycleStateProvisioning DatabaseSoftwareImageLifecycleStateEnum = "PROVISIONING"
	DatabaseSoftwareImageLifecycleStateAvailable    DatabaseSoftwareImageLifecycleStateEnum = "AVAILABLE"
	DatabaseSoftwareImageLifecycleStateDeleting     DatabaseSoftwareImageLifecycleStateEnum = "DELETING"
	DatabaseSoftwareImageLifecycleStateDeleted      DatabaseSoftwareImageLifecycleStateEnum = "DELETED"
	DatabaseSoftwareImageLifecycleStateFailed       DatabaseSoftwareImageLifecycleStateEnum = "FAILED"
	DatabaseSoftwareImageLifecycleStateTerminating  DatabaseSoftwareImageLifecycleStateEnum = "TERMINATING"
	DatabaseSoftwareImageLifecycleStateTerminated   DatabaseSoftwareImageLifecycleStateEnum = "TERMINATED"
	DatabaseSoftwareImageLifecycleStateUpdating     DatabaseSoftwareImageLifecycleStateEnum = "UPDATING"
)

var mappingDatabaseSoftwareImageLifecycleStateEnum = map[string]DatabaseSoftwareImageLifecycleStateEnum{
	"PROVISIONING": DatabaseSoftwareImageLifecycleStateProvisioning,
	"AVAILABLE":    DatabaseSoftwareImageLifecycleStateAvailable,
	"DELETING":     DatabaseSoftwareImageLifecycleStateDeleting,
	"DELETED":      DatabaseSoftwareImageLifecycleStateDeleted,
	"FAILED":       DatabaseSoftwareImageLifecycleStateFailed,
	"TERMINATING":  DatabaseSoftwareImageLifecycleStateTerminating,
	"TERMINATED":   DatabaseSoftwareImageLifecycleStateTerminated,
	"UPDATING":     DatabaseSoftwareImageLifecycleStateUpdating,
}

var mappingDatabaseSoftwareImageLifecycleStateEnumLowerCase = map[string]DatabaseSoftwareImageLifecycleStateEnum{
	"provisioning": DatabaseSoftwareImageLifecycleStateProvisioning,
	"available":    DatabaseSoftwareImageLifecycleStateAvailable,
	"deleting":     DatabaseSoftwareImageLifecycleStateDeleting,
	"deleted":      DatabaseSoftwareImageLifecycleStateDeleted,
	"failed":       DatabaseSoftwareImageLifecycleStateFailed,
	"terminating":  DatabaseSoftwareImageLifecycleStateTerminating,
	"terminated":   DatabaseSoftwareImageLifecycleStateTerminated,
	"updating":     DatabaseSoftwareImageLifecycleStateUpdating,
}

// GetDatabaseSoftwareImageLifecycleStateEnumValues Enumerates the set of values for DatabaseSoftwareImageLifecycleStateEnum
func GetDatabaseSoftwareImageLifecycleStateEnumValues() []DatabaseSoftwareImageLifecycleStateEnum {
	values := make([]DatabaseSoftwareImageLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageLifecycleStateEnum
func GetDatabaseSoftwareImageLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"DELETING",
		"DELETED",
		"FAILED",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingDatabaseSoftwareImageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageLifecycleStateEnum(val string) (DatabaseSoftwareImageLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseSoftwareImageImageTypeEnum Enum with underlying type: string
type DatabaseSoftwareImageImageTypeEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageImageTypeEnum
const (
	DatabaseSoftwareImageImageTypeGridImage     DatabaseSoftwareImageImageTypeEnum = "GRID_IMAGE"
	DatabaseSoftwareImageImageTypeDatabaseImage DatabaseSoftwareImageImageTypeEnum = "DATABASE_IMAGE"
)

var mappingDatabaseSoftwareImageImageTypeEnum = map[string]DatabaseSoftwareImageImageTypeEnum{
	"GRID_IMAGE":     DatabaseSoftwareImageImageTypeGridImage,
	"DATABASE_IMAGE": DatabaseSoftwareImageImageTypeDatabaseImage,
}

var mappingDatabaseSoftwareImageImageTypeEnumLowerCase = map[string]DatabaseSoftwareImageImageTypeEnum{
	"grid_image":     DatabaseSoftwareImageImageTypeGridImage,
	"database_image": DatabaseSoftwareImageImageTypeDatabaseImage,
}

// GetDatabaseSoftwareImageImageTypeEnumValues Enumerates the set of values for DatabaseSoftwareImageImageTypeEnum
func GetDatabaseSoftwareImageImageTypeEnumValues() []DatabaseSoftwareImageImageTypeEnum {
	values := make([]DatabaseSoftwareImageImageTypeEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageImageTypeEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageImageTypeEnum
func GetDatabaseSoftwareImageImageTypeEnumStringValues() []string {
	return []string{
		"GRID_IMAGE",
		"DATABASE_IMAGE",
	}
}

// GetMappingDatabaseSoftwareImageImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageImageTypeEnum(val string) (DatabaseSoftwareImageImageTypeEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseSoftwareImageImageShapeFamilyEnum Enum with underlying type: string
type DatabaseSoftwareImageImageShapeFamilyEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageImageShapeFamilyEnum
const (
	DatabaseSoftwareImageImageShapeFamilyVmBmShape    DatabaseSoftwareImageImageShapeFamilyEnum = "VM_BM_SHAPE"
	DatabaseSoftwareImageImageShapeFamilyExadataShape DatabaseSoftwareImageImageShapeFamilyEnum = "EXADATA_SHAPE"
	DatabaseSoftwareImageImageShapeFamilyExaccShape   DatabaseSoftwareImageImageShapeFamilyEnum = "EXACC_SHAPE"
)

var mappingDatabaseSoftwareImageImageShapeFamilyEnum = map[string]DatabaseSoftwareImageImageShapeFamilyEnum{
	"VM_BM_SHAPE":   DatabaseSoftwareImageImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": DatabaseSoftwareImageImageShapeFamilyExadataShape,
	"EXACC_SHAPE":   DatabaseSoftwareImageImageShapeFamilyExaccShape,
}

var mappingDatabaseSoftwareImageImageShapeFamilyEnumLowerCase = map[string]DatabaseSoftwareImageImageShapeFamilyEnum{
	"vm_bm_shape":   DatabaseSoftwareImageImageShapeFamilyVmBmShape,
	"exadata_shape": DatabaseSoftwareImageImageShapeFamilyExadataShape,
	"exacc_shape":   DatabaseSoftwareImageImageShapeFamilyExaccShape,
}

// GetDatabaseSoftwareImageImageShapeFamilyEnumValues Enumerates the set of values for DatabaseSoftwareImageImageShapeFamilyEnum
func GetDatabaseSoftwareImageImageShapeFamilyEnumValues() []DatabaseSoftwareImageImageShapeFamilyEnum {
	values := make([]DatabaseSoftwareImageImageShapeFamilyEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageImageShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageImageShapeFamilyEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageImageShapeFamilyEnum
func GetDatabaseSoftwareImageImageShapeFamilyEnumStringValues() []string {
	return []string{
		"VM_BM_SHAPE",
		"EXADATA_SHAPE",
		"EXACC_SHAPE",
	}
}

// GetMappingDatabaseSoftwareImageImageShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageImageShapeFamilyEnum(val string) (DatabaseSoftwareImageImageShapeFamilyEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageImageShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
