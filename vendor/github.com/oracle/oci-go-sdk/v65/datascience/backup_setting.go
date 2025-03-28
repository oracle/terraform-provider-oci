// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// BackupSetting Back up setting details of the model.
type BackupSetting struct {

	// Boolean flag representing whether backup needs to be enabled/disabled for the model.
	IsBackupEnabled *bool `mandatory:"true" json:"isBackupEnabled"`

	// OCI backup region for the model.
	BackupRegion *string `mandatory:"true" json:"backupRegion"`

	// Customer notification on backup success/failure events.
	CustomerNotificationType ModelSettingCustomerNotificationTypeEnum `mandatory:"false" json:"customerNotificationType,omitempty"`
}

func (m BackupSetting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupSetting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModelSettingCustomerNotificationTypeEnum(string(m.CustomerNotificationType)); !ok && m.CustomerNotificationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CustomerNotificationType: %s. Supported values are: %s.", m.CustomerNotificationType, strings.Join(GetModelSettingCustomerNotificationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
