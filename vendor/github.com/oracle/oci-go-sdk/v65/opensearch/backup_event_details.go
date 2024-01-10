// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupEventDetails Details about a cluster backup event.
type BackupEventDetails struct {

	// The OCID of the OpenSearch cluster for the cluster backup.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// The result of the cluster backup operation.
	BackupState BackupStateEnum `mandatory:"true" json:"backupState"`

	// The date and time the cluster backup event started. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the cluster backup event started. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// The name of the cluster backup.
	SnapshotName *string `mandatory:"false" json:"snapshotName"`

	// The cluster backup size in GB.
	BackupSize *float64 `mandatory:"false" json:"backupSize"`
}

func (m BackupEventDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupEventDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupStateEnum(string(m.BackupState)); !ok && m.BackupState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupState: %s. Supported values are: %s.", m.BackupState, strings.Join(GetBackupStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
