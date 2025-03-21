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

// LustreFileSystemSummary Summary information about a Lustre file system.
type LustreFileSystemSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Lustre file system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain the file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Lustre file system`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Short description of the Lustre file system.
	// Avoid entering confidential information.
	FileSystemDescription *string `mandatory:"true" json:"fileSystemDescription"`

	// The date and time the Lustre file system was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2024-04-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Lustre file system was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2024-04-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Lustre file system.
	LifecycleState LustreFileSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The Lustre file system name. This is used in mount commands and other aspects of the client command line interface.
	// The default file system name is 'lustre'. The file system name is limited to 8 characters. Allowed characters are lower and upper case English letters, numbers, and '_'.
	FileSystemName *string `mandatory:"true" json:"fileSystemName"`

	// Capacity of the Lustre file system in GB.
	CapacityInGBs *int `mandatory:"true" json:"capacityInGBs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the Lustre file system is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Lustre file system performance tier. A value of `MBPS_PER_TB_125` represents 125 megabytes per second per terabyte.
	PerformanceTier LustreFileSystemSummaryPerformanceTierEnum `mandatory:"true" json:"performanceTier"`

	// The date and time the LustreFileSystem current billing cycle will end, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format. Once a cycle ends,
	// it is updated automatically to next timestamp which is after 30 days.
	// File systems deleted earlier will still incur charges till this date.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeBillingCycleEnd *common.SDKTime `mandatory:"true" json:"timeBillingCycleEnd"`

	// The IPv4 address of MGS (Lustre Management Service) used by clients to mount the file system.
	// Example: `10.0.0.4`
	ManagementServiceAddress *string `mandatory:"true" json:"managementServiceAddress"`

	// Type of network used by clients to mount the file system.
	// Example: `tcp`
	Lnet *string `mandatory:"true" json:"lnet"`

	// Major version of Lustre running in the Lustre file system.
	// Example: `2.15`
	MajorVersion *string `mandatory:"true" json:"majorVersion"`

	RootSquashConfiguration *RootSquashConfiguration `mandatory:"true" json:"rootSquashConfiguration"`

	// A message that describes the current state of the Lustre file system in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m LustreFileSystemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LustreFileSystemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLustreFileSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLustreFileSystemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLustreFileSystemSummaryPerformanceTierEnum(string(m.PerformanceTier)); !ok && m.PerformanceTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceTier: %s. Supported values are: %s.", m.PerformanceTier, strings.Join(GetLustreFileSystemSummaryPerformanceTierEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LustreFileSystemSummaryPerformanceTierEnum Enum with underlying type: string
type LustreFileSystemSummaryPerformanceTierEnum string

// Set of constants representing the allowable values for LustreFileSystemSummaryPerformanceTierEnum
const (
	LustreFileSystemSummaryPerformanceTier125  LustreFileSystemSummaryPerformanceTierEnum = "MBPS_PER_TB_125"
	LustreFileSystemSummaryPerformanceTier250  LustreFileSystemSummaryPerformanceTierEnum = "MBPS_PER_TB_250"
	LustreFileSystemSummaryPerformanceTier500  LustreFileSystemSummaryPerformanceTierEnum = "MBPS_PER_TB_500"
	LustreFileSystemSummaryPerformanceTier1000 LustreFileSystemSummaryPerformanceTierEnum = "MBPS_PER_TB_1000"
)

var mappingLustreFileSystemSummaryPerformanceTierEnum = map[string]LustreFileSystemSummaryPerformanceTierEnum{
	"MBPS_PER_TB_125":  LustreFileSystemSummaryPerformanceTier125,
	"MBPS_PER_TB_250":  LustreFileSystemSummaryPerformanceTier250,
	"MBPS_PER_TB_500":  LustreFileSystemSummaryPerformanceTier500,
	"MBPS_PER_TB_1000": LustreFileSystemSummaryPerformanceTier1000,
}

var mappingLustreFileSystemSummaryPerformanceTierEnumLowerCase = map[string]LustreFileSystemSummaryPerformanceTierEnum{
	"mbps_per_tb_125":  LustreFileSystemSummaryPerformanceTier125,
	"mbps_per_tb_250":  LustreFileSystemSummaryPerformanceTier250,
	"mbps_per_tb_500":  LustreFileSystemSummaryPerformanceTier500,
	"mbps_per_tb_1000": LustreFileSystemSummaryPerformanceTier1000,
}

// GetLustreFileSystemSummaryPerformanceTierEnumValues Enumerates the set of values for LustreFileSystemSummaryPerformanceTierEnum
func GetLustreFileSystemSummaryPerformanceTierEnumValues() []LustreFileSystemSummaryPerformanceTierEnum {
	values := make([]LustreFileSystemSummaryPerformanceTierEnum, 0)
	for _, v := range mappingLustreFileSystemSummaryPerformanceTierEnum {
		values = append(values, v)
	}
	return values
}

// GetLustreFileSystemSummaryPerformanceTierEnumStringValues Enumerates the set of values in String for LustreFileSystemSummaryPerformanceTierEnum
func GetLustreFileSystemSummaryPerformanceTierEnumStringValues() []string {
	return []string{
		"MBPS_PER_TB_125",
		"MBPS_PER_TB_250",
		"MBPS_PER_TB_500",
		"MBPS_PER_TB_1000",
	}
}

// GetMappingLustreFileSystemSummaryPerformanceTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLustreFileSystemSummaryPerformanceTierEnum(val string) (LustreFileSystemSummaryPerformanceTierEnum, bool) {
	enum, ok := mappingLustreFileSystemSummaryPerformanceTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
