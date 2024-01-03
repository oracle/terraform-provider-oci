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

// OpensearchClusterBackupSummary The summary of information about an OpenSearch cluster backup.
type OpensearchClusterBackupSummary struct {

	// The OCID of the cluster backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment where the cluster backup is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies whether the cluster backup was created manually, or automatically as a scheduled backup.
	BackupType OpensearchClusterBackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The OCID of the source OpenSearch cluster for the cluster backup.
	SourceClusterId *string `mandatory:"true" json:"sourceClusterId"`

	// The current state of the cluster backup.
	LifecycleState OpensearchClusterBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the cluster backup. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the cluster backup was created. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cluster backup was updated. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional information about the current lifecycle state of the cluster backup.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the cluster backup expires. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeExpired *common.SDKTime `mandatory:"false" json:"timeExpired"`

	// The size in GB of the cluster backup.
	BackupSize *float64 `mandatory:"false" json:"backupSize"`

	// The name of the source OpenSearch cluster for the cluster backup.
	SourceClusterDisplayName *string `mandatory:"false" json:"sourceClusterDisplayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OpensearchClusterBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchClusterBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetOpensearchClusterBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOpensearchClusterBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
