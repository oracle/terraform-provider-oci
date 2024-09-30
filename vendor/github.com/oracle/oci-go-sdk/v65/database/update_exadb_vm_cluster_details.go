// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateExadbVmClusterDetails Details for updating the Exadata VM cluster on Exascale Infrastructure. Applies to Exadata Database Service on Exascale Infrastructure only.
type UpdateExadbVmClusterDetails struct {

	// The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure.
	TotalECpuCount *int `mandatory:"false" json:"totalECpuCount"`

	// The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure.
	EnabledECpuCount *int `mandatory:"false" json:"enabledECpuCount"`

	VmFileSystemStorage *ExadbVmClusterStorageDetails `mandatory:"false" json:"vmFileSystemStorage"`

	// The number of nodes to be added in the Exadata VM cluster on Exascale Infrastructure.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel UpdateExadbVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// Operating system version of the image.
	SystemVersion *string `mandatory:"false" json:"systemVersion"`

	// Grid Setup will be done using this grid image id.
	// The grid image id can be extracted from
	// 1. Obtain the supported major versions using API /20160918/giVersions?compartmentId=<compartmentId>&shape=EXADB_XS&availabilityDomain=<AD name>
	// 2. Replace {version} with one of the supported major versions and obtain the supported minor versions using
	// API /20160918/giVersions/{version}/minorVersions?compartmentId=<compartmentId>&shapeFamily=EXADB_XS&availabilityDomain=<AD name>
	GridImageId *string `mandatory:"false" json:"gridImageId"`

	// The update action.
	UpdateAction UpdateExadbVmClusterDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`
}

func (m UpdateExadbVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExadbVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateExadbVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateExadbVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateExadbVmClusterDetailsUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetUpdateExadbVmClusterDetailsUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExadbVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateExadbVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateExadbVmClusterDetailsLicenseModelEnum
const (
	UpdateExadbVmClusterDetailsLicenseModelLicenseIncluded     UpdateExadbVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateExadbVmClusterDetailsLicenseModelBringYourOwnLicense UpdateExadbVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateExadbVmClusterDetailsLicenseModelEnum = map[string]UpdateExadbVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateExadbVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateExadbVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateExadbVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateExadbVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateExadbVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateExadbVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateExadbVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateExadbVmClusterDetailsLicenseModelEnum
func GetUpdateExadbVmClusterDetailsLicenseModelEnumValues() []UpdateExadbVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateExadbVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateExadbVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExadbVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateExadbVmClusterDetailsLicenseModelEnum
func GetUpdateExadbVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateExadbVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExadbVmClusterDetailsLicenseModelEnum(val string) (UpdateExadbVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateExadbVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateExadbVmClusterDetailsUpdateActionEnum Enum with underlying type: string
type UpdateExadbVmClusterDetailsUpdateActionEnum string

// Set of constants representing the allowable values for UpdateExadbVmClusterDetailsUpdateActionEnum
const (
	UpdateExadbVmClusterDetailsUpdateActionRollingApply    UpdateExadbVmClusterDetailsUpdateActionEnum = "ROLLING_APPLY"
	UpdateExadbVmClusterDetailsUpdateActionNonRollingApply UpdateExadbVmClusterDetailsUpdateActionEnum = "NON_ROLLING_APPLY"
	UpdateExadbVmClusterDetailsUpdateActionPrecheck        UpdateExadbVmClusterDetailsUpdateActionEnum = "PRECHECK"
	UpdateExadbVmClusterDetailsUpdateActionRollback        UpdateExadbVmClusterDetailsUpdateActionEnum = "ROLLBACK"
)

var mappingUpdateExadbVmClusterDetailsUpdateActionEnum = map[string]UpdateExadbVmClusterDetailsUpdateActionEnum{
	"ROLLING_APPLY":     UpdateExadbVmClusterDetailsUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateExadbVmClusterDetailsUpdateActionNonRollingApply,
	"PRECHECK":          UpdateExadbVmClusterDetailsUpdateActionPrecheck,
	"ROLLBACK":          UpdateExadbVmClusterDetailsUpdateActionRollback,
}

var mappingUpdateExadbVmClusterDetailsUpdateActionEnumLowerCase = map[string]UpdateExadbVmClusterDetailsUpdateActionEnum{
	"rolling_apply":     UpdateExadbVmClusterDetailsUpdateActionRollingApply,
	"non_rolling_apply": UpdateExadbVmClusterDetailsUpdateActionNonRollingApply,
	"precheck":          UpdateExadbVmClusterDetailsUpdateActionPrecheck,
	"rollback":          UpdateExadbVmClusterDetailsUpdateActionRollback,
}

// GetUpdateExadbVmClusterDetailsUpdateActionEnumValues Enumerates the set of values for UpdateExadbVmClusterDetailsUpdateActionEnum
func GetUpdateExadbVmClusterDetailsUpdateActionEnumValues() []UpdateExadbVmClusterDetailsUpdateActionEnum {
	values := make([]UpdateExadbVmClusterDetailsUpdateActionEnum, 0)
	for _, v := range mappingUpdateExadbVmClusterDetailsUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExadbVmClusterDetailsUpdateActionEnumStringValues Enumerates the set of values in String for UpdateExadbVmClusterDetailsUpdateActionEnum
func GetUpdateExadbVmClusterDetailsUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateExadbVmClusterDetailsUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExadbVmClusterDetailsUpdateActionEnum(val string) (UpdateExadbVmClusterDetailsUpdateActionEnum, bool) {
	enum, ok := mappingUpdateExadbVmClusterDetailsUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
