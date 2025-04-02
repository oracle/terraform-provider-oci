// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateLustreFileSystemDetails The details to create a Lustre file system.
type CreateLustreFileSystemDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Lustre file system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain the file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The Lustre file system name. This is used in mount commands and other aspects of the client command line interface.
	// The file system name is limited to 8 characters. Allowed characters are lower and upper case English letters, numbers, and '_'.
	// If you have multiple Lustre file systems mounted on the same clients, this name can help distinguish them.
	FileSystemName *string `mandatory:"true" json:"fileSystemName"`

	// Capacity of the Lustre file system in GB. You can increase capacity only in multiples of 5 TB.
	CapacityInGBs *int `mandatory:"true" json:"capacityInGBs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the Lustre file system is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Lustre file system performance tier. A value of `MBPS_PER_TB_125` represents 125 megabytes per second per terabyte.
	PerformanceTier CreateLustreFileSystemDetailsPerformanceTierEnum `mandatory:"true" json:"performanceTier"`

	RootSquashConfiguration *RootSquashConfiguration `mandatory:"true" json:"rootSquashConfiguration"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Lustre file system`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Short description of the Lustre file system.
	// Avoid entering confidential information.
	FileSystemDescription *string `mandatory:"false" json:"fileSystemDescription"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A list of Network Security Group OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this lustre file system.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the lustre file system from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group in which the Lustre file system exists.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`
}

func (m CreateLustreFileSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLustreFileSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateLustreFileSystemDetailsPerformanceTierEnum(string(m.PerformanceTier)); !ok && m.PerformanceTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceTier: %s. Supported values are: %s.", m.PerformanceTier, strings.Join(GetCreateLustreFileSystemDetailsPerformanceTierEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateLustreFileSystemDetailsPerformanceTierEnum Enum with underlying type: string
type CreateLustreFileSystemDetailsPerformanceTierEnum string

// Set of constants representing the allowable values for CreateLustreFileSystemDetailsPerformanceTierEnum
const (
	CreateLustreFileSystemDetailsPerformanceTier125  CreateLustreFileSystemDetailsPerformanceTierEnum = "MBPS_PER_TB_125"
	CreateLustreFileSystemDetailsPerformanceTier250  CreateLustreFileSystemDetailsPerformanceTierEnum = "MBPS_PER_TB_250"
	CreateLustreFileSystemDetailsPerformanceTier500  CreateLustreFileSystemDetailsPerformanceTierEnum = "MBPS_PER_TB_500"
	CreateLustreFileSystemDetailsPerformanceTier1000 CreateLustreFileSystemDetailsPerformanceTierEnum = "MBPS_PER_TB_1000"
)

var mappingCreateLustreFileSystemDetailsPerformanceTierEnum = map[string]CreateLustreFileSystemDetailsPerformanceTierEnum{
	"MBPS_PER_TB_125":  CreateLustreFileSystemDetailsPerformanceTier125,
	"MBPS_PER_TB_250":  CreateLustreFileSystemDetailsPerformanceTier250,
	"MBPS_PER_TB_500":  CreateLustreFileSystemDetailsPerformanceTier500,
	"MBPS_PER_TB_1000": CreateLustreFileSystemDetailsPerformanceTier1000,
}

var mappingCreateLustreFileSystemDetailsPerformanceTierEnumLowerCase = map[string]CreateLustreFileSystemDetailsPerformanceTierEnum{
	"mbps_per_tb_125":  CreateLustreFileSystemDetailsPerformanceTier125,
	"mbps_per_tb_250":  CreateLustreFileSystemDetailsPerformanceTier250,
	"mbps_per_tb_500":  CreateLustreFileSystemDetailsPerformanceTier500,
	"mbps_per_tb_1000": CreateLustreFileSystemDetailsPerformanceTier1000,
}

// GetCreateLustreFileSystemDetailsPerformanceTierEnumValues Enumerates the set of values for CreateLustreFileSystemDetailsPerformanceTierEnum
func GetCreateLustreFileSystemDetailsPerformanceTierEnumValues() []CreateLustreFileSystemDetailsPerformanceTierEnum {
	values := make([]CreateLustreFileSystemDetailsPerformanceTierEnum, 0)
	for _, v := range mappingCreateLustreFileSystemDetailsPerformanceTierEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateLustreFileSystemDetailsPerformanceTierEnumStringValues Enumerates the set of values in String for CreateLustreFileSystemDetailsPerformanceTierEnum
func GetCreateLustreFileSystemDetailsPerformanceTierEnumStringValues() []string {
	return []string{
		"MBPS_PER_TB_125",
		"MBPS_PER_TB_250",
		"MBPS_PER_TB_500",
		"MBPS_PER_TB_1000",
	}
}

// GetMappingCreateLustreFileSystemDetailsPerformanceTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateLustreFileSystemDetailsPerformanceTierEnum(val string) (CreateLustreFileSystemDetailsPerformanceTierEnum, bool) {
	enum, ok := mappingCreateLustreFileSystemDetailsPerformanceTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
