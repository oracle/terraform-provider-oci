// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BackupPolicy The Backup policy for the DB System.
type BackupPolicy struct {

	// If automated backups are enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The start of a 30-minute window of time in which daily, automated backups occur.
	// This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.
	// At some point in the window, the system may incur a brief service disruption as the backup is performed.
	// If not defined, a window is selected from the following Region-based time-spans:
	// - eu-frankfurt-1: 20:00 - 04:00 UTC
	// - us-ashburn-1: 03:00 - 11:00 UTC
	// - uk-london-1: 06:00 - 14:00 UTC
	// - ap-tokyo-1: 13:00 - 21:00
	// - us-phoenix-1: 06:00 - 14:00
	WindowStartTime *string `mandatory:"true" json:"windowStartTime"`

	// The number of days automated backups are retained.
	RetentionInDays *int `mandatory:"true" json:"retentionInDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BackupPolicy) String() string {
	return common.PointerString(m)
}
