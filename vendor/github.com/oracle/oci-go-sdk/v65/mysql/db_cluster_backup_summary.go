// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbClusterBackupSummary A summary of shared-storage DB cluster backups.
type DbClusterBackupSummary struct {

	// OCID of the shared-storage DB cluster backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the shared-storage DB cluster backup.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current lifecycle state of the backup.
	LifecycleState DbClusterBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Source DbClusterBackupSource `mandatory:"true" json:"source"`

	// The date and time the backup was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The size of the backup in GiBs.
	BackupSizeInGBs *int `mandatory:"false" json:"backupSizeInGBs"`

	DbClusterSnapshotSummary *DbClusterSnapshotSummary `mandatory:"false" json:"dbClusterSnapshotSummary"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster backup.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Additional information about the current lifecycle state of the backup.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The MySQL version of the shared-storage DB cluster at the time the backup was taken.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The number of days to retain this backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DbClusterBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbClusterBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbClusterBackupSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupSizeInGBs          *int                              `json:"backupSizeInGBs"`
		DbClusterSnapshotSummary *DbClusterSnapshotSummary         `json:"dbClusterSnapshotSummary"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		Description              *string                           `json:"description"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		LifecycleDetails         *string                           `json:"lifecycleDetails"`
		MysqlVersion             *string                           `json:"mysqlVersion"`
		RetentionInDays          *int                              `json:"retentionInDays"`
		SystemTags               map[string]map[string]interface{} `json:"systemTags"`
		Id                       *string                           `json:"id"`
		CompartmentId            *string                           `json:"compartmentId"`
		DisplayName              *string                           `json:"displayName"`
		LifecycleState           DbClusterBackupLifecycleStateEnum `json:"lifecycleState"`
		Source                   dbclusterbackupsource             `json:"source"`
		TimeCreated              *common.SDKTime                   `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BackupSizeInGBs = model.BackupSizeInGBs

	m.DbClusterSnapshotSummary = model.DbClusterSnapshotSummary

	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.LifecycleDetails = model.LifecycleDetails

	m.MysqlVersion = model.MysqlVersion

	m.RetentionInDays = model.RetentionInDays

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(DbClusterBackupSource)
	} else {
		m.Source = nil
	}

	m.TimeCreated = model.TimeCreated

	return
}
