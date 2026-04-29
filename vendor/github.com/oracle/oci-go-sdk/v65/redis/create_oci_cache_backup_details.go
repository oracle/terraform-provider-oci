// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciCacheBackupDetails The information to create a new OCI Cache Backup.
type CreateOciCacheBackupDetails struct {

	// Backup display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCI Cache cluster identifier
	SourceClusterId *string `mandatory:"true" json:"sourceClusterId"`

	// Backup description
	Description *string `mandatory:"false" json:"description"`

	// Backup retention period in days.
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Specifies whether the backup was created from a replica or primary node
	BackupSource OciCacheBackupBackupSourceEnum `mandatory:"false" json:"backupSource,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOciCacheBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOciCacheBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOciCacheBackupBackupSourceEnum(string(m.BackupSource)); !ok && m.BackupSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupSource: %s. Supported values are: %s.", m.BackupSource, strings.Join(GetOciCacheBackupBackupSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
