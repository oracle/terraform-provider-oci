// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ProtectedDatabase A protected database is an Oracle Cloud Database whose
// backups are managed by Oracle Database Autonomous Recovery Service. Each protected database
// requires a recovery service subnet and a protection policy to use Recovery Service as
// the backup destination for centralized backup and recovery
type ProtectedDatabase struct {

	// The OCID of the protected database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the protected database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The dbUniqueName for the protected database in Recovery Service. You cannot change the unique name.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName"`

	// The virtual private catalog (VPC) user credentials that authenticates the protected database to access Recovery Service.
	VpcUserName *string `mandatory:"true" json:"vpcUserName"`

	// The size of the protected database. XS - Less than 5GB, S - 5GB to 50GB, M - 50GB to 500GB, L - 500GB to 1TB, XL - 1TB to 5TB, XXL - Greater than 5TB.
	DatabaseSize DatabaseSizesEnum `mandatory:"true" json:"databaseSize"`

	// The OCID of the protection policy associated with the protected database.
	ProtectionPolicyId *string `mandatory:"true" json:"protectionPolicyId"`

	// List of recovery service subnet resources associated with the protected database.
	RecoveryServiceSubnets []RecoveryServiceSubnetDetails `mandatory:"true" json:"recoveryServiceSubnets"`

	// The protected database name. You can change the displayName. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy.
	// The retention lock feature controls whether Recovery Service strictly preserves backups for the duration defined in a policy. Retention lock is useful to enforce recovery window compliance and to prevent unintentional modifications to protected database backups.
	// Recovery Service enforces a 14-day delay before the retention lock set for a policy can take effect.
	PolicyLockedDateTime *string `mandatory:"false" json:"policyLockedDateTime"`

	// The OCID of the protected database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The size of the database in GBs, in gigabytes.
	DatabaseSizeInGBs *int `mandatory:"false" json:"databaseSizeInGBs"`

	// The percentage of data changes that exist in the database between successive incremental backups.
	ChangeRate *float64 `mandatory:"false" json:"changeRate"`

	// The compression ratio of the protected database. The compression ratio represents the ratio of compressed block size to expanded block size.
	CompressionRatio *float64 `mandatory:"false" json:"compressionRatio"`

	// The value TRUE indicates that the protected database is configured to use Real-time data protection, and redo-data is sent from the protected database to Recovery Service.
	// Real-time data protection substantially reduces the window of potential data loss that exists between successive archived redo log backups. For this to be effective, additional
	// configuration is needed on client side.
	IsRedoLogsShipped *bool `mandatory:"false" json:"isRedoLogsShipped"`

	// An RFC3339 formatted datetime string that indicates the created time for a protected database. For example: '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// An RFC3339 formatted datetime string that indicates the last updated time for a protected database. For example: '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Protected Database.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates the protection status of the database.
	// A 'PROTECTED' status indicates that Recovery Service can ensure database recovery to any point in time within the entire recovery window. The potential data loss exposure since the last backup is:
	//  - Less than 10 seconds, if Real-time data protection is enabled
	//  - Less than 70 minutes if Real-time data protection is disabled
	// A 'WARNING' status indicates that Recovery Service can ensure database recovery within the current recovery window - 1 day. The potential data loss exposure since the last backup is:
	//  - Greater than 10 seconds, if Real-time data protection is enabled
	//  - Greater than 60 minutes, if if Real-time data protection is disabled
	// An 'ALERT' status indicates that Recovery Service cannot recover the database within the current recovery window.
	Health HealthEnum `mandatory:"false" json:"health,omitempty"`

	// Indicates whether the protected database is created by Recovery Service or created manually.
	// Set to <b>TRUE</b> for a service-defined protected database. When you enable the OCI-managed automatic backups option for a database and set Recovery Service as the backup destination, then Recovery Service creates the associated protected database resource.
	// Set to <b>FALSE</b> for a user-defined protected database.
	IsReadOnlyResource *bool `mandatory:"false" json:"isReadOnlyResource"`

	// Detailed description about the current lifecycle state of the protected database. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A message describing the current health of the protected database.
	HealthDetails *string `mandatory:"false" json:"healthDetails"`

	Metrics *Metrics `mandatory:"false" json:"metrics"`

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

func (m ProtectedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectedDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSizesEnum(string(m.DatabaseSize)); !ok && m.DatabaseSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSize: %s. Supported values are: %s.", m.DatabaseSize, strings.Join(GetDatabaseSizesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHealthEnum(string(m.Health)); !ok && m.Health != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Health: %s. Supported values are: %s.", m.Health, strings.Join(GetHealthEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
