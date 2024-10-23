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

// RetentionOperationDetails Retention operation details for the model.
type RetentionOperationDetails struct {

	// The archival status of model.
	ArchiveState ModelSettingActionStateEnum `mandatory:"true" json:"archiveState"`

	// The archival state details of the model.
	ArchiveStateDetails *string `mandatory:"true" json:"archiveStateDetails"`

	// The estimated archival time of the model based on the provided retention setting.
	TimeArchivalScheduled *common.SDKTime `mandatory:"true" json:"timeArchivalScheduled"`

	// The deletion status of the archived model.
	DeleteState ModelSettingActionStateEnum `mandatory:"true" json:"deleteState"`

	// The deletion status details of the archived model.
	DeleteStateDetails *string `mandatory:"true" json:"deleteStateDetails"`

	// The estimated deletion time of the model based on the provided retention setting.
	TimeDeletionScheduled *common.SDKTime `mandatory:"true" json:"timeDeletionScheduled"`
}

func (m RetentionOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RetentionOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelSettingActionStateEnum(string(m.ArchiveState)); !ok && m.ArchiveState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchiveState: %s. Supported values are: %s.", m.ArchiveState, strings.Join(GetModelSettingActionStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModelSettingActionStateEnum(string(m.DeleteState)); !ok && m.DeleteState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeleteState: %s. Supported values are: %s.", m.DeleteState, strings.Join(GetModelSettingActionStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
