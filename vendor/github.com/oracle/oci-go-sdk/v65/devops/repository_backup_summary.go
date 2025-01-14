// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RepositoryBackupSummary Summary of the Repository Backup.
type RepositoryBackupSummary struct {

	// Unique identifier for backup that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// The time when backup was completed. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// DevOps Repository Identifier
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The name of repository's default branch.
	DefaultBranch *string `mandatory:"true" json:"defaultBranch"`

	// Latest commit id of the default branch
	CommitId *string `mandatory:"true" json:"commitId"`

	// Lifecycle state of backup
	LifecycleState RepositoryBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Repository Backup Settings resource ocid associated with this backup.
	RepositoryBackupSettingsId *string `mandatory:"true" json:"repositoryBackupSettingsId"`

	// Size of file or directory.
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// display name of backup.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Name of DevOps Repository when the backup was taken
	RepositoryName *string `mandatory:"false" json:"repositoryName"`

	// Name of Project to which the repository for which the backup was taken belongs to.
	ProjectName *string `mandatory:"false" json:"projectName"`

	// The textual description for the Repository Backup
	Description *string `mandatory:"false" json:"description"`

	// The description of lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RepositoryBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRepositoryBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
