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

// DatabaseSoftwareImageSummary The Database service supports the creation of database software images for use in creating and patching DB systems and databases.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see Overview of the Identity Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
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

	// The patches included in the image and the version of the image.
	IncludedPatchesSummary *string `mandatory:"false" json:"includedPatchesSummary"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageOneOffPatches []string `mandatory:"false" json:"databaseSoftwareImageOneOffPatches"`

	// The output from the OPatch lsInventory command, which is passed as a string.
	LsInventory *string `mandatory:"false" json:"lsInventory"`

	// True if this Database software image is supported for Upgrade.
	IsUpgradeSupported *bool `mandatory:"false" json:"isUpgradeSupported"`
}

func (m DatabaseSoftwareImageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseSoftwareImageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseSoftwareImageSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryImageTypeEnum(string(m.ImageType)); !ok && m.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", m.ImageType, strings.Join(GetDatabaseSoftwareImageSummaryImageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum(string(m.ImageShapeFamily)); !ok && m.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", m.ImageShapeFamily, strings.Join(GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	DatabaseSoftwareImageSummaryLifecycleStateTerminating  DatabaseSoftwareImageSummaryLifecycleStateEnum = "TERMINATING"
	DatabaseSoftwareImageSummaryLifecycleStateTerminated   DatabaseSoftwareImageSummaryLifecycleStateEnum = "TERMINATED"
	DatabaseSoftwareImageSummaryLifecycleStateUpdating     DatabaseSoftwareImageSummaryLifecycleStateEnum = "UPDATING"
)

var mappingDatabaseSoftwareImageSummaryLifecycleStateEnum = map[string]DatabaseSoftwareImageSummaryLifecycleStateEnum{
	"PROVISIONING": DatabaseSoftwareImageSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DatabaseSoftwareImageSummaryLifecycleStateAvailable,
	"DELETING":     DatabaseSoftwareImageSummaryLifecycleStateDeleting,
	"DELETED":      DatabaseSoftwareImageSummaryLifecycleStateDeleted,
	"FAILED":       DatabaseSoftwareImageSummaryLifecycleStateFailed,
	"TERMINATING":  DatabaseSoftwareImageSummaryLifecycleStateTerminating,
	"TERMINATED":   DatabaseSoftwareImageSummaryLifecycleStateTerminated,
	"UPDATING":     DatabaseSoftwareImageSummaryLifecycleStateUpdating,
}

var mappingDatabaseSoftwareImageSummaryLifecycleStateEnumLowerCase = map[string]DatabaseSoftwareImageSummaryLifecycleStateEnum{
	"provisioning": DatabaseSoftwareImageSummaryLifecycleStateProvisioning,
	"available":    DatabaseSoftwareImageSummaryLifecycleStateAvailable,
	"deleting":     DatabaseSoftwareImageSummaryLifecycleStateDeleting,
	"deleted":      DatabaseSoftwareImageSummaryLifecycleStateDeleted,
	"failed":       DatabaseSoftwareImageSummaryLifecycleStateFailed,
	"terminating":  DatabaseSoftwareImageSummaryLifecycleStateTerminating,
	"terminated":   DatabaseSoftwareImageSummaryLifecycleStateTerminated,
	"updating":     DatabaseSoftwareImageSummaryLifecycleStateUpdating,
}

// GetDatabaseSoftwareImageSummaryLifecycleStateEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryLifecycleStateEnum
func GetDatabaseSoftwareImageSummaryLifecycleStateEnumValues() []DatabaseSoftwareImageSummaryLifecycleStateEnum {
	values := make([]DatabaseSoftwareImageSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageSummaryLifecycleStateEnum
func GetDatabaseSoftwareImageSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingDatabaseSoftwareImageSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageSummaryLifecycleStateEnum(val string) (DatabaseSoftwareImageSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseSoftwareImageSummaryImageTypeEnum Enum with underlying type: string
type DatabaseSoftwareImageSummaryImageTypeEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageSummaryImageTypeEnum
const (
	DatabaseSoftwareImageSummaryImageTypeGridImage     DatabaseSoftwareImageSummaryImageTypeEnum = "GRID_IMAGE"
	DatabaseSoftwareImageSummaryImageTypeDatabaseImage DatabaseSoftwareImageSummaryImageTypeEnum = "DATABASE_IMAGE"
)

var mappingDatabaseSoftwareImageSummaryImageTypeEnum = map[string]DatabaseSoftwareImageSummaryImageTypeEnum{
	"GRID_IMAGE":     DatabaseSoftwareImageSummaryImageTypeGridImage,
	"DATABASE_IMAGE": DatabaseSoftwareImageSummaryImageTypeDatabaseImage,
}

var mappingDatabaseSoftwareImageSummaryImageTypeEnumLowerCase = map[string]DatabaseSoftwareImageSummaryImageTypeEnum{
	"grid_image":     DatabaseSoftwareImageSummaryImageTypeGridImage,
	"database_image": DatabaseSoftwareImageSummaryImageTypeDatabaseImage,
}

// GetDatabaseSoftwareImageSummaryImageTypeEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryImageTypeEnum
func GetDatabaseSoftwareImageSummaryImageTypeEnumValues() []DatabaseSoftwareImageSummaryImageTypeEnum {
	values := make([]DatabaseSoftwareImageSummaryImageTypeEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryImageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageSummaryImageTypeEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageSummaryImageTypeEnum
func GetDatabaseSoftwareImageSummaryImageTypeEnumStringValues() []string {
	return []string{
		"GRID_IMAGE",
		"DATABASE_IMAGE",
	}
}

// GetMappingDatabaseSoftwareImageSummaryImageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageSummaryImageTypeEnum(val string) (DatabaseSoftwareImageSummaryImageTypeEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageSummaryImageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseSoftwareImageSummaryImageShapeFamilyEnum Enum with underlying type: string
type DatabaseSoftwareImageSummaryImageShapeFamilyEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageSummaryImageShapeFamilyEnum
const (
	DatabaseSoftwareImageSummaryImageShapeFamilyVmBmShape    DatabaseSoftwareImageSummaryImageShapeFamilyEnum = "VM_BM_SHAPE"
	DatabaseSoftwareImageSummaryImageShapeFamilyExadataShape DatabaseSoftwareImageSummaryImageShapeFamilyEnum = "EXADATA_SHAPE"
	DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape   DatabaseSoftwareImageSummaryImageShapeFamilyEnum = "EXACC_SHAPE"
)

var mappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum = map[string]DatabaseSoftwareImageSummaryImageShapeFamilyEnum{
	"VM_BM_SHAPE":   DatabaseSoftwareImageSummaryImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": DatabaseSoftwareImageSummaryImageShapeFamilyExadataShape,
	"EXACC_SHAPE":   DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape,
}

var mappingDatabaseSoftwareImageSummaryImageShapeFamilyEnumLowerCase = map[string]DatabaseSoftwareImageSummaryImageShapeFamilyEnum{
	"vm_bm_shape":   DatabaseSoftwareImageSummaryImageShapeFamilyVmBmShape,
	"exadata_shape": DatabaseSoftwareImageSummaryImageShapeFamilyExadataShape,
	"exacc_shape":   DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape,
}

// GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumValues Enumerates the set of values for DatabaseSoftwareImageSummaryImageShapeFamilyEnum
func GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumValues() []DatabaseSoftwareImageSummaryImageShapeFamilyEnum {
	values := make([]DatabaseSoftwareImageSummaryImageShapeFamilyEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumStringValues Enumerates the set of values in String for DatabaseSoftwareImageSummaryImageShapeFamilyEnum
func GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumStringValues() []string {
	return []string{
		"VM_BM_SHAPE",
		"EXADATA_SHAPE",
		"EXACC_SHAPE",
	}
}

// GetMappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum(val string) (DatabaseSoftwareImageSummaryImageShapeFamilyEnum, bool) {
	enum, ok := mappingDatabaseSoftwareImageSummaryImageShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
