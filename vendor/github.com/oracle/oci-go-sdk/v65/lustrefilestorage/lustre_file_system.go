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

// LustreFileSystem A Lustre file system is a parallel file system that is used as a storage solution for HPC/AI/ML workloads.
// For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type LustreFileSystem struct {

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
	//
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

	// Capacity of the Lustre file system in GB.
	CapacityInGBs *int `mandatory:"true" json:"capacityInGBs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the Lustre file system is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Lustre file system performance tier. A value of `MBPS_PER_TB_125` represents 125 megabytes per second per terabyte.
	PerformanceTier LustreFileSystemPerformanceTierEnum `mandatory:"true" json:"performanceTier"`

	// The IPv4 address of MGS (Lustre Management Service) used by clients to mount the file system. For example '10.0.0.4'.
	ManagementServiceAddress *string `mandatory:"true" json:"managementServiceAddress"`

	// The Lustre file system name. This is used in mount commands and other aspects of the client command line interface.
	// The default file system name is 'lustre'. The file system name is limited to 8 characters. Allowed characters are lower and upper case English letters, numbers, and '_'.
	FileSystemName *string `mandatory:"true" json:"fileSystemName"`

	// Type of network used by clients to mount the file system.
	// Example: `tcp`
	Lnet *string `mandatory:"true" json:"lnet"`

	// Major version of Lustre running in the Lustre file system.
	// Example: `2.15`
	MajorVersion *string `mandatory:"true" json:"majorVersion"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"true" json:"maintenanceWindow"`

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

	// The date and time that the current billing cycle for the file system will end, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format. After the current cycle ends,
	// this date is updated automatically to the next timestamp, which is 30 days later.
	// File systems deleted earlier than this time will still incur charges until the billing cycle ends.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeBillingCycleEnd *common.SDKTime `mandatory:"false" json:"timeBillingCycleEnd"`
}

func (m LustreFileSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LustreFileSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLustreFileSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLustreFileSystemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLustreFileSystemPerformanceTierEnum(string(m.PerformanceTier)); !ok && m.PerformanceTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceTier: %s. Supported values are: %s.", m.PerformanceTier, strings.Join(GetLustreFileSystemPerformanceTierEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LustreFileSystemLifecycleStateEnum Enum with underlying type: string
type LustreFileSystemLifecycleStateEnum string

// Set of constants representing the allowable values for LustreFileSystemLifecycleStateEnum
const (
	LustreFileSystemLifecycleStateCreating LustreFileSystemLifecycleStateEnum = "CREATING"
	LustreFileSystemLifecycleStateUpdating LustreFileSystemLifecycleStateEnum = "UPDATING"
	LustreFileSystemLifecycleStateActive   LustreFileSystemLifecycleStateEnum = "ACTIVE"
	LustreFileSystemLifecycleStateInactive LustreFileSystemLifecycleStateEnum = "INACTIVE"
	LustreFileSystemLifecycleStateDeleting LustreFileSystemLifecycleStateEnum = "DELETING"
	LustreFileSystemLifecycleStateDeleted  LustreFileSystemLifecycleStateEnum = "DELETED"
	LustreFileSystemLifecycleStateFailed   LustreFileSystemLifecycleStateEnum = "FAILED"
)

var mappingLustreFileSystemLifecycleStateEnum = map[string]LustreFileSystemLifecycleStateEnum{
	"CREATING": LustreFileSystemLifecycleStateCreating,
	"UPDATING": LustreFileSystemLifecycleStateUpdating,
	"ACTIVE":   LustreFileSystemLifecycleStateActive,
	"INACTIVE": LustreFileSystemLifecycleStateInactive,
	"DELETING": LustreFileSystemLifecycleStateDeleting,
	"DELETED":  LustreFileSystemLifecycleStateDeleted,
	"FAILED":   LustreFileSystemLifecycleStateFailed,
}

var mappingLustreFileSystemLifecycleStateEnumLowerCase = map[string]LustreFileSystemLifecycleStateEnum{
	"creating": LustreFileSystemLifecycleStateCreating,
	"updating": LustreFileSystemLifecycleStateUpdating,
	"active":   LustreFileSystemLifecycleStateActive,
	"inactive": LustreFileSystemLifecycleStateInactive,
	"deleting": LustreFileSystemLifecycleStateDeleting,
	"deleted":  LustreFileSystemLifecycleStateDeleted,
	"failed":   LustreFileSystemLifecycleStateFailed,
}

// GetLustreFileSystemLifecycleStateEnumValues Enumerates the set of values for LustreFileSystemLifecycleStateEnum
func GetLustreFileSystemLifecycleStateEnumValues() []LustreFileSystemLifecycleStateEnum {
	values := make([]LustreFileSystemLifecycleStateEnum, 0)
	for _, v := range mappingLustreFileSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLustreFileSystemLifecycleStateEnumStringValues Enumerates the set of values in String for LustreFileSystemLifecycleStateEnum
func GetLustreFileSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLustreFileSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLustreFileSystemLifecycleStateEnum(val string) (LustreFileSystemLifecycleStateEnum, bool) {
	enum, ok := mappingLustreFileSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LustreFileSystemPerformanceTierEnum Enum with underlying type: string
type LustreFileSystemPerformanceTierEnum string

// Set of constants representing the allowable values for LustreFileSystemPerformanceTierEnum
const (
	LustreFileSystemPerformanceTier125  LustreFileSystemPerformanceTierEnum = "MBPS_PER_TB_125"
	LustreFileSystemPerformanceTier250  LustreFileSystemPerformanceTierEnum = "MBPS_PER_TB_250"
	LustreFileSystemPerformanceTier500  LustreFileSystemPerformanceTierEnum = "MBPS_PER_TB_500"
	LustreFileSystemPerformanceTier1000 LustreFileSystemPerformanceTierEnum = "MBPS_PER_TB_1000"
)

var mappingLustreFileSystemPerformanceTierEnum = map[string]LustreFileSystemPerformanceTierEnum{
	"MBPS_PER_TB_125":  LustreFileSystemPerformanceTier125,
	"MBPS_PER_TB_250":  LustreFileSystemPerformanceTier250,
	"MBPS_PER_TB_500":  LustreFileSystemPerformanceTier500,
	"MBPS_PER_TB_1000": LustreFileSystemPerformanceTier1000,
}

var mappingLustreFileSystemPerformanceTierEnumLowerCase = map[string]LustreFileSystemPerformanceTierEnum{
	"mbps_per_tb_125":  LustreFileSystemPerformanceTier125,
	"mbps_per_tb_250":  LustreFileSystemPerformanceTier250,
	"mbps_per_tb_500":  LustreFileSystemPerformanceTier500,
	"mbps_per_tb_1000": LustreFileSystemPerformanceTier1000,
}

// GetLustreFileSystemPerformanceTierEnumValues Enumerates the set of values for LustreFileSystemPerformanceTierEnum
func GetLustreFileSystemPerformanceTierEnumValues() []LustreFileSystemPerformanceTierEnum {
	values := make([]LustreFileSystemPerformanceTierEnum, 0)
	for _, v := range mappingLustreFileSystemPerformanceTierEnum {
		values = append(values, v)
	}
	return values
}

// GetLustreFileSystemPerformanceTierEnumStringValues Enumerates the set of values in String for LustreFileSystemPerformanceTierEnum
func GetLustreFileSystemPerformanceTierEnumStringValues() []string {
	return []string{
		"MBPS_PER_TB_125",
		"MBPS_PER_TB_250",
		"MBPS_PER_TB_500",
		"MBPS_PER_TB_1000",
	}
}

// GetMappingLustreFileSystemPerformanceTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLustreFileSystemPerformanceTierEnum(val string) (LustreFileSystemPerformanceTierEnum, bool) {
	enum, ok := mappingLustreFileSystemPerformanceTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
