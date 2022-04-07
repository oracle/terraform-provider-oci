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

// DbHomeSummary A directory where Oracle Database software is installed. A bare metal or Exadata DB system can have multiple Database Homes
// and each Database Home can run a different supported version of Oracle Database. A virtual machine DB system can have only one Database Home.
// For more information, see Bare Metal and Virtual Machine DB Systems (https://docs.cloud.oracle.com/Content/Database/Concepts/overview.htm) and Exadata DB Systems (https://docs.cloud.oracle.com/Content/Database/Concepts/exaoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DbHomeSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Database Home. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Database Home.
	LifecycleState DbHomeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Oracle Database version.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The location of the Oracle Database Home.
	DbHomeLocation *string `mandatory:"true" json:"dbHomeLocation"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Database Home was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// List of one-off patches for Database Homes.
	OneOffPatches []string `mandatory:"false" json:"oneOffPatches"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`
}

func (m DbHomeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbHomeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbHomeSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbHomeSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbHomeSummaryLifecycleStateEnum Enum with underlying type: string
type DbHomeSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbHomeSummaryLifecycleStateEnum
const (
	DbHomeSummaryLifecycleStateProvisioning DbHomeSummaryLifecycleStateEnum = "PROVISIONING"
	DbHomeSummaryLifecycleStateAvailable    DbHomeSummaryLifecycleStateEnum = "AVAILABLE"
	DbHomeSummaryLifecycleStateUpdating     DbHomeSummaryLifecycleStateEnum = "UPDATING"
	DbHomeSummaryLifecycleStateTerminating  DbHomeSummaryLifecycleStateEnum = "TERMINATING"
	DbHomeSummaryLifecycleStateTerminated   DbHomeSummaryLifecycleStateEnum = "TERMINATED"
	DbHomeSummaryLifecycleStateFailed       DbHomeSummaryLifecycleStateEnum = "FAILED"
)

var mappingDbHomeSummaryLifecycleStateEnum = map[string]DbHomeSummaryLifecycleStateEnum{
	"PROVISIONING": DbHomeSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DbHomeSummaryLifecycleStateAvailable,
	"UPDATING":     DbHomeSummaryLifecycleStateUpdating,
	"TERMINATING":  DbHomeSummaryLifecycleStateTerminating,
	"TERMINATED":   DbHomeSummaryLifecycleStateTerminated,
	"FAILED":       DbHomeSummaryLifecycleStateFailed,
}

var mappingDbHomeSummaryLifecycleStateEnumLowerCase = map[string]DbHomeSummaryLifecycleStateEnum{
	"provisioning": DbHomeSummaryLifecycleStateProvisioning,
	"available":    DbHomeSummaryLifecycleStateAvailable,
	"updating":     DbHomeSummaryLifecycleStateUpdating,
	"terminating":  DbHomeSummaryLifecycleStateTerminating,
	"terminated":   DbHomeSummaryLifecycleStateTerminated,
	"failed":       DbHomeSummaryLifecycleStateFailed,
}

// GetDbHomeSummaryLifecycleStateEnumValues Enumerates the set of values for DbHomeSummaryLifecycleStateEnum
func GetDbHomeSummaryLifecycleStateEnumValues() []DbHomeSummaryLifecycleStateEnum {
	values := make([]DbHomeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbHomeSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbHomeSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DbHomeSummaryLifecycleStateEnum
func GetDbHomeSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingDbHomeSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbHomeSummaryLifecycleStateEnum(val string) (DbHomeSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDbHomeSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
