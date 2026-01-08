// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRepositoryBackupSettingsDetails The information about configuration of a Repository Backup.
type CreateRepositoryBackupSettingsDetails struct {

	// The name of the Settings.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The frequency at which backup of the repository will be created.
	Frequency RepositoryBackupSettingsFrequencyEnum `mandatory:"true" json:"frequency"`

	// The number of days or months after which the backup file produced by the rule will be deleted
	RetentionPeriodInDays *int `mandatory:"true" json:"retentionPeriodInDays"`

	// The start date after which backups will be created. An RFC3339 formatted datetime string
	TimeBackupCreationStart *common.SDKTime `mandatory:"true" json:"timeBackupCreationStart"`

	// The textual description for the Repository Backup
	Description *string `mandatory:"false" json:"description"`

	// Enable backup. The backup is enabled by default.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Enforce a minimum count of backups
	MinimumBackupsToRetain *int `mandatory:"false" json:"minimumBackupsToRetain"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRepositoryBackupSettingsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRepositoryBackupSettingsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryBackupSettingsFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetRepositoryBackupSettingsFrequencyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
