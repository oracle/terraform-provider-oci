// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OkeClusterBackupConfig The details of backup performed on OKE Cluster.
type OkeClusterBackupConfig struct {

	// A list of namespaces to be included in the backup.
	// The default value is null. If a list of namespaces to include is not provided, all namespaces will be backed up.
	// Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both.
	// This property applies to the OKE cluster member in primary region.
	// Example: ["default", "pv-nginx"]
	Namespaces []string `mandatory:"false" json:"namespaces"`

	// A list of namespaces to be excluded from the backup.
	// The default value is null. If a list of namespaces to exclude is not provided, all namespaces will be backed up.
	// Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both.
	// This property applies to OKE cluster members in the primary region.
	// Example: ["namespace_string_3", "namespace_string_4"]
	ExcludeNamespaces []string `mandatory:"false" json:"excludeNamespaces"`

	// The schedule for backing up namespaces to the destination region. If a backup schedule is not specified, only a single backup will be created. This format of the string specifying the backup schedule must conform with RFC-5545.
	// This schedule will use the UTC timezone.
	// This property applies to the OKE cluster member in primary region.
	// Example: FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;INTERVAL=1
	BackupSchedule *string `mandatory:"false" json:"backupSchedule"`

	// Controls the behaviour of image replication across regions.
	// This property applies to the OKE cluster member in primary region.
	ReplicateImages OkeClusterImageReplicationEnum `mandatory:"false" json:"replicateImages,omitempty"`

	// The maximum number of backups that should be retained.
	// This property applies to the OKE cluster member in primary region.
	MaxNumberOfBackupsRetained *int `mandatory:"false" json:"maxNumberOfBackupsRetained"`

	// The OCID of the vault secret that stores the image credential.
	// This property applies to the OKE cluster member in both the primary and standby region.
	ImageReplicationVaultSecretId *string `mandatory:"false" json:"imageReplicationVaultSecretId"`
}

func (m OkeClusterBackupConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OkeClusterBackupConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOkeClusterImageReplicationEnum(string(m.ReplicateImages)); !ok && m.ReplicateImages != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicateImages: %s. Supported values are: %s.", m.ReplicateImages, strings.Join(GetOkeClusterImageReplicationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
