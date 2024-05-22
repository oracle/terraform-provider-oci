// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeBackupSummary The information about the nodeBackupSummary.
type NodeBackupSummary struct {

	// The id of the node backup.
	Id *string `mandatory:"true" json:"id"`

	// BDS generated name for the backup. Format is nodeHostName_timeCreated.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The instance OCID of the node, which is the resource from which the node backup was acquired.
	NodeInstanceId *string `mandatory:"true" json:"nodeInstanceId"`

	// Host name of the node that the backup belongs to.
	NodeHostName *string `mandatory:"true" json:"nodeHostName"`

	// type based on how backup action was initiated.
	BackupTriggerType NodeBackupBackupTriggerTypeEnum `mandatory:"true" json:"backupTriggerType"`

	// Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created
	BackupType NodeBackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The state of NodeBackup.
	LifecycleState NodeBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m NodeBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeBackupBackupTriggerTypeEnum(string(m.BackupTriggerType)); !ok && m.BackupTriggerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupTriggerType: %s. Supported values are: %s.", m.BackupTriggerType, strings.Join(GetNodeBackupBackupTriggerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetNodeBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
