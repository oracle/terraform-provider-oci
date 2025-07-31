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

// CreateBaseccVmClusterDetails Details for the create Basecc VM cluster operation. Applies to Base Cloud@Customer instances only.
type CreateBaseccVmClusterDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores to enable for the Basecc VM cluster.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Oracle Base Infrastructure.
	BaseInfrastructureId *string `mandatory:"true" json:"baseInfrastructureId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Base Virtual Machine Network Id.
	BaseVmClusterNetworkId *string `mandatory:"true" json:"baseVmClusterNetworkId"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// The user-friendly name for the Basecc VM cluster on BICC Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The public key portion of one or more key pairs used for SSH access to the Basecc VM cluster on BICC Infrastructure.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The Oracle Database Edition that applies to all the databases on the DB system.
	// Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition CreateBaseccVmClusterDetailsDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The number of nodes in the Basecc VM cluster on BICC Infrastructure.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute node on which Oracle Base virtual cluster single instance is launched.
	// Note keep it blank for RAC Oracle Base virtual machine cluster.
	DbServerId *string `mandatory:"false" json:"dbServerId"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The description for Basecc VM Cluster.
	Description *string `mandatory:"false" json:"description"`

	// The time zone to use for the Basecc VM cluster on BICC Infrastructure. For details, see Time Zones (https://docs.oracle.com/iaas/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The cluster name for Basecc VM cluster on BICC Infrastructure. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The Boot Additional Vm Storage Size in GB, to be allocated for the /u01 partition for the Basecc VM cluster on BICC Infrastructure.
	AdditionalVmStorageSizeInGBs *int `mandatory:"false" json:"additionalVmStorageSizeInGBs"`

	// The Data Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The Reco Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	RecoStorageSizeInGBs *int `mandatory:"false" json:"recoStorageSizeInGBs"`

	// The memory to be allocated per VM in GBs. The default is 8GB per core.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The Oracle license model that applies to the Basecc VM cluster on BICC Infrastructure. The default is LICENSE_INCLUDED.
	LicenseModel CreateBaseccVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The vmcluster type for the Base Cloud@Customer VM cluster.
	VmClusterType CreateBaseccVmClusterDetailsVmClusterTypeEnum `mandatory:"false" json:"vmClusterType,omitempty"`

	CloudAutomationUpdateDetails *CloudAutomationUpdateDetails `mandatory:"false" json:"cloudAutomationUpdateDetails"`
}

func (m CreateBaseccVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBaseccVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateBaseccVmClusterDetailsDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetCreateBaseccVmClusterDetailsDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateBaseccVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateBaseccVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateBaseccVmClusterDetailsVmClusterTypeEnum(string(m.VmClusterType)); !ok && m.VmClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmClusterType: %s. Supported values are: %s.", m.VmClusterType, strings.Join(GetCreateBaseccVmClusterDetailsVmClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBaseccVmClusterDetailsDatabaseEditionEnum Enum with underlying type: string
type CreateBaseccVmClusterDetailsDatabaseEditionEnum string

// Set of constants representing the allowable values for CreateBaseccVmClusterDetailsDatabaseEditionEnum
const (
	CreateBaseccVmClusterDetailsDatabaseEditionStandardEdition                     CreateBaseccVmClusterDetailsDatabaseEditionEnum = "STANDARD_EDITION"
	CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEdition                   CreateBaseccVmClusterDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION"
	CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionHighPerformance    CreateBaseccVmClusterDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionExtremePerformance CreateBaseccVmClusterDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionDeveloper          CreateBaseccVmClusterDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingCreateBaseccVmClusterDetailsDatabaseEditionEnum = map[string]CreateBaseccVmClusterDetailsDatabaseEditionEnum{
	"STANDARD_EDITION":                       CreateBaseccVmClusterDetailsDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingCreateBaseccVmClusterDetailsDatabaseEditionEnumLowerCase = map[string]CreateBaseccVmClusterDetailsDatabaseEditionEnum{
	"standard_edition":                       CreateBaseccVmClusterDetailsDatabaseEditionStandardEdition,
	"enterprise_edition":                     CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           CreateBaseccVmClusterDetailsDatabaseEditionEnterpriseEditionDeveloper,
}

// GetCreateBaseccVmClusterDetailsDatabaseEditionEnumValues Enumerates the set of values for CreateBaseccVmClusterDetailsDatabaseEditionEnum
func GetCreateBaseccVmClusterDetailsDatabaseEditionEnumValues() []CreateBaseccVmClusterDetailsDatabaseEditionEnum {
	values := make([]CreateBaseccVmClusterDetailsDatabaseEditionEnum, 0)
	for _, v := range mappingCreateBaseccVmClusterDetailsDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBaseccVmClusterDetailsDatabaseEditionEnumStringValues Enumerates the set of values in String for CreateBaseccVmClusterDetailsDatabaseEditionEnum
func GetCreateBaseccVmClusterDetailsDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingCreateBaseccVmClusterDetailsDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBaseccVmClusterDetailsDatabaseEditionEnum(val string) (CreateBaseccVmClusterDetailsDatabaseEditionEnum, bool) {
	enum, ok := mappingCreateBaseccVmClusterDetailsDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateBaseccVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateBaseccVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateBaseccVmClusterDetailsLicenseModelEnum
const (
	CreateBaseccVmClusterDetailsLicenseModelLicenseIncluded     CreateBaseccVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense CreateBaseccVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateBaseccVmClusterDetailsLicenseModelEnum = map[string]CreateBaseccVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateBaseccVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateBaseccVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateBaseccVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateBaseccVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateBaseccVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateBaseccVmClusterDetailsLicenseModelEnum
func GetCreateBaseccVmClusterDetailsLicenseModelEnumValues() []CreateBaseccVmClusterDetailsLicenseModelEnum {
	values := make([]CreateBaseccVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateBaseccVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBaseccVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateBaseccVmClusterDetailsLicenseModelEnum
func GetCreateBaseccVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateBaseccVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBaseccVmClusterDetailsLicenseModelEnum(val string) (CreateBaseccVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateBaseccVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateBaseccVmClusterDetailsVmClusterTypeEnum Enum with underlying type: string
type CreateBaseccVmClusterDetailsVmClusterTypeEnum string

// Set of constants representing the allowable values for CreateBaseccVmClusterDetailsVmClusterTypeEnum
const (
	CreateBaseccVmClusterDetailsVmClusterTypeRegular   CreateBaseccVmClusterDetailsVmClusterTypeEnum = "REGULAR"
	CreateBaseccVmClusterDetailsVmClusterTypeDeveloper CreateBaseccVmClusterDetailsVmClusterTypeEnum = "DEVELOPER"
)

var mappingCreateBaseccVmClusterDetailsVmClusterTypeEnum = map[string]CreateBaseccVmClusterDetailsVmClusterTypeEnum{
	"REGULAR":   CreateBaseccVmClusterDetailsVmClusterTypeRegular,
	"DEVELOPER": CreateBaseccVmClusterDetailsVmClusterTypeDeveloper,
}

var mappingCreateBaseccVmClusterDetailsVmClusterTypeEnumLowerCase = map[string]CreateBaseccVmClusterDetailsVmClusterTypeEnum{
	"regular":   CreateBaseccVmClusterDetailsVmClusterTypeRegular,
	"developer": CreateBaseccVmClusterDetailsVmClusterTypeDeveloper,
}

// GetCreateBaseccVmClusterDetailsVmClusterTypeEnumValues Enumerates the set of values for CreateBaseccVmClusterDetailsVmClusterTypeEnum
func GetCreateBaseccVmClusterDetailsVmClusterTypeEnumValues() []CreateBaseccVmClusterDetailsVmClusterTypeEnum {
	values := make([]CreateBaseccVmClusterDetailsVmClusterTypeEnum, 0)
	for _, v := range mappingCreateBaseccVmClusterDetailsVmClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBaseccVmClusterDetailsVmClusterTypeEnumStringValues Enumerates the set of values in String for CreateBaseccVmClusterDetailsVmClusterTypeEnum
func GetCreateBaseccVmClusterDetailsVmClusterTypeEnumStringValues() []string {
	return []string{
		"REGULAR",
		"DEVELOPER",
	}
}

// GetMappingCreateBaseccVmClusterDetailsVmClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBaseccVmClusterDetailsVmClusterTypeEnum(val string) (CreateBaseccVmClusterDetailsVmClusterTypeEnum, bool) {
	enum, ok := mappingCreateBaseccVmClusterDetailsVmClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
