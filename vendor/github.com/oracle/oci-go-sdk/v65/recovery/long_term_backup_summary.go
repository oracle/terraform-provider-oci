// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LongTermBackupSummary Generate an on-demand backup and retain it long-term as per the LTR backup retention period specified.
type LongTermBackupSummary struct {

	// The Long Term Backup OCID.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the long term backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The protected database ocid.
	ProtectedDatabaseId *string `mandatory:"true" json:"protectedDatabaseId"`

	// Retain the long-term backup as per the backup retention period specified
	RetentionPeriod []RetentionPeriodValue `mandatory:"true" json:"retentionPeriod"`

	// An RFC3339 formatted datetime string that indicates the time after which the long term backup will be deleted. For example '2020-05-22T21:10:29.600Z'. This timestamp has to be 95 days-10 years from the current timestamp
	RetentionUntilDateTime *common.SDKTime `mandatory:"true" json:"retentionUntilDateTime"`

	// The current state of the long term backup.
	LifecycleState LongTermBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The long term backup name. You can change the displayName. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// System Change Number (SCN) to which backup is consistent.
	RetentionScn *int `mandatory:"false" json:"retentionScn"`

	// An RFC3339 formatted datetime string that indicates the time for a long term backup to be initiated.
	RetentionPointInTime *common.SDKTime `mandatory:"false" json:"retentionPointInTime"`

	// An RFC3339 formatted datetime string that indicates the time when the backup was actually taken. For example '2020-05-22T21:10:29.600Z'
	TimeBackupInitiated *common.SDKTime `mandatory:"false" json:"timeBackupInitiated"`

	// An RFC3339 formatted datetime string that indicates the time when the backup was actually completed. For example '2020-05-22T21:10:29.600Z'
	TimeBackupCompleted *common.SDKTime `mandatory:"false" json:"timeBackupCompleted"`

	// The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud.
	DatabaseIdentifier *string `mandatory:"false" json:"databaseIdentifier"`

	// Database size
	DatabaseSizeInGBs *int `mandatory:"false" json:"databaseSizeInGBs"`

	// An RFC3339 formatted datetime string that indicates the created time for the long term backup. For example: '2020-05-22T21:10:29.600Z'.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// An RFC3339 formatted datetime string that indicates the updated time for the long term backup. For example: '2020-05-22T21:10:29.600Z'.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// More details on the state of the backup when it is in Creating lifecycleState.
	LifecycleSubstate LongTermBackupLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

	// Detailed description about the current lifecycle state of the long term backup. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Recovery Manager (RMAN) assigned unique identifier for the long-term backup.
	RmanTag *string `mandatory:"false" json:"rmanTag"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LongTermBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LongTermBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLongTermBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLongTermBackupLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLongTermBackupLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetLongTermBackupLifecycleSubstateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
