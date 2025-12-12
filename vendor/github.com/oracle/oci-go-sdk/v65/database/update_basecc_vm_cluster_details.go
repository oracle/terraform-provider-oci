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

// UpdateBaseccVmClusterDetails Details for updating the VM cluster. Applies to Basecc Cloud@Customer instances only.
type UpdateBaseccVmClusterDetails struct {

	// The number of CPU cores to enable for the Basecc VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The memory to be allocated per VM in GBs. The default is 8GB per core.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The Data Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The Reco Disk Group size in GB for the Basecc VM cluster on BICC Infrastructure.
	RecoStorageSizeInGBs *int `mandatory:"false" json:"recoStorageSizeInGBs"`

	// The Oracle license model that applies to the Basecc VM cluster on BICC Infrastructure. The default is LICENSE_INCLUDED.
	LicenseModel UpdateBaseccVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the Basecc VM cluster on BICC Infrastructure.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	// The user-friendly name for the Basecc VM cluster on BICC Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Additional Vm Storage Size in GB, to be allocated for the /u01 partition for the Basecc VM cluster on BICC Infrastructure.
	AdditionalVmStorageSizeInGBs *int `mandatory:"false" json:"additionalVmStorageSizeInGBs"`

	UpdateDetails *BaseccVmClusterUpdateDetails `mandatory:"false" json:"updateDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DataCollectionOptions *DataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	CloudAutomationUpdateDetails *CloudAutomationUpdateDetails `mandatory:"false" json:"cloudAutomationUpdateDetails"`
}

func (m UpdateBaseccVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBaseccVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateBaseccVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateBaseccVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateBaseccVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateBaseccVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateBaseccVmClusterDetailsLicenseModelEnum
const (
	UpdateBaseccVmClusterDetailsLicenseModelLicenseIncluded     UpdateBaseccVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense UpdateBaseccVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateBaseccVmClusterDetailsLicenseModelEnum = map[string]UpdateBaseccVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateBaseccVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateBaseccVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateBaseccVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateBaseccVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateBaseccVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateBaseccVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateBaseccVmClusterDetailsLicenseModelEnum
func GetUpdateBaseccVmClusterDetailsLicenseModelEnumValues() []UpdateBaseccVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateBaseccVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateBaseccVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateBaseccVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateBaseccVmClusterDetailsLicenseModelEnum
func GetUpdateBaseccVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateBaseccVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateBaseccVmClusterDetailsLicenseModelEnum(val string) (UpdateBaseccVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateBaseccVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
