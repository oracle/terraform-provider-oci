// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// The date and time the Database Home was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m DbHomeSummary) String() string {
	return common.PointerString(m)
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

var mappingDbHomeSummaryLifecycleState = map[string]DbHomeSummaryLifecycleStateEnum{
	"PROVISIONING": DbHomeSummaryLifecycleStateProvisioning,
	"AVAILABLE":    DbHomeSummaryLifecycleStateAvailable,
	"UPDATING":     DbHomeSummaryLifecycleStateUpdating,
	"TERMINATING":  DbHomeSummaryLifecycleStateTerminating,
	"TERMINATED":   DbHomeSummaryLifecycleStateTerminated,
	"FAILED":       DbHomeSummaryLifecycleStateFailed,
}

// GetDbHomeSummaryLifecycleStateEnumValues Enumerates the set of values for DbHomeSummaryLifecycleStateEnum
func GetDbHomeSummaryLifecycleStateEnumValues() []DbHomeSummaryLifecycleStateEnum {
	values := make([]DbHomeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbHomeSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
