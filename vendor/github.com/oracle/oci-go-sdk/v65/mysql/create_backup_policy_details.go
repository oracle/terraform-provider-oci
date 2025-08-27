// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBackupPolicyDetails Backup policy as optionally used for DB System Creation.
type CreateBackupPolicyDetails struct {

	// Specifies if automatic backups are enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Retains the backup to be deleted due to the retention policy in DELETE SCHEDULED
	// state for 7 days before permanently deleting it.
	SoftDelete SoftDeleteEnum `mandatory:"false" json:"softDelete,omitempty"`

	// List of policies of a DB system to schedule cross-region DB system backup copy.
	// The policy includes the name of the destination region to which the DB system backup will be copied, and
	// an optional parameter which specifies the retention period of the copied DB system backup in days.
	// **Note:** Currently, only one policy can be specified in the list.
	CopyPolicies []CopyPolicy `mandatory:"false" json:"copyPolicies"`

	// The start of a 30-minute window of time in which daily, automated backups occur.
	// This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.
	// At some point in the window, the system may incur a brief service disruption as the backup is performed.
	WindowStartTime *string `mandatory:"false" json:"windowStartTime"`

	// Number of days to retain an automatic backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	PitrPolicy *PitrPolicy `mandatory:"false" json:"pitrPolicy"`
}

func (m CreateBackupPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBackupPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftDeleteEnum(string(m.SoftDelete)); !ok && m.SoftDelete != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftDelete: %s. Supported values are: %s.", m.SoftDelete, strings.Join(GetSoftDeleteEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
