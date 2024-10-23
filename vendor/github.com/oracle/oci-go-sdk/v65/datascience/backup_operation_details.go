// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupOperationDetails Backup operation details of the model.
type BackupOperationDetails struct {

	// The backup status of the model.
	BackupState ModelSettingActionStateEnum `mandatory:"true" json:"backupState"`

	// The backup execution status details of the model.
	BackupStateDetails *string `mandatory:"true" json:"backupStateDetails"`

	// The last backup execution time of the model.
	TimeLastBackup *common.SDKTime `mandatory:"false" json:"timeLastBackup"`
}

func (m BackupOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelSettingActionStateEnum(string(m.BackupState)); !ok && m.BackupState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupState: %s. Supported values are: %s.", m.BackupState, strings.Join(GetModelSettingActionStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
