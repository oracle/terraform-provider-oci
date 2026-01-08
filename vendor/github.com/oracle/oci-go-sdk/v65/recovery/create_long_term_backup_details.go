// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateLongTermBackupDetails Describes the parameters required to create an on demand long term backup.
type CreateLongTermBackupDetails struct {

	// The OCID of the protected database for which you want to create the long-term backup.
	ProtectedDatabaseId *string `mandatory:"true" json:"protectedDatabaseId"`

	// The maximum period to retain the long-term backup. Specify the retention period type and the duration for the long-term backup. If you have chosen the retention period type as 'DAYS', then specify a duration ranging from 90 days to 3650 days. If you have chosen the retention period type as 'YEARS', then specify a duration ranging from 1 year to 10 years.
	RetentionPeriod []RetentionPeriodValue `mandatory:"true" json:"retentionPeriod"`

	// A user provided name for the long term backup. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The desired target point (SCN) at which you want to create the long-term backup of the database.For example, specify the value as 1000 if you want to create the long-term backup until SCN 1000. If you want to specify the target point as a time expression instead of the SCN value, then use the longTermRetentionPointInTime parameter.
	RetentionScn *int `mandatory:"false" json:"retentionScn"`

	// An RFC3339 formatted datetime string that indicates the desired target point in time in the database at which you want to create the long-term backup. For example, if you want the long-term backup to include all the changes until May 22 at 9:10 PM, then specify the value as, '2020-05-22T21:10:00.000Z'. If you want to specify the target point as an SCN value instead of the target time, then use the databaseSCN parameter.
	RetentionPointInTime *common.SDKTime `mandatory:"false" json:"retentionPointInTime"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. For more information, see Resource Tags (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm)
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateLongTermBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLongTermBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
