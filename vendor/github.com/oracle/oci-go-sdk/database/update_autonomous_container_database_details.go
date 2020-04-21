// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutonomousContainerDatabaseDetails Describes the modification parameters for the Autonomous Container Database.
type UpdateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Database Patch model preference.
	PatchModel UpdateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m UpdateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates         UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingUpdateAutonomousContainerDatabaseDetailsPatchModel = map[string]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []UpdateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsPatchModel {
		values = append(values, v)
	}
	return values
}
