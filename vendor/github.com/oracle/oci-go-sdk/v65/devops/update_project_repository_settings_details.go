// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateProjectRepositorySettingsDetails Information to update custom project repository settings.
type UpdateProjectRepositorySettingsDetails struct {
	MergeSettings *MergeSettings `mandatory:"false" json:"mergeSettings"`

	ApprovalRules *UpdateApprovalRuleDetailsCollection `mandatory:"false" json:"approvalRules"`

	RepositoryAccessMode RepositoryAccessMode `mandatory:"false" json:"repositoryAccessMode"`
}

func (m UpdateProjectRepositorySettingsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateProjectRepositorySettingsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateProjectRepositorySettingsDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MergeSettings        *MergeSettings                       `json:"mergeSettings"`
		ApprovalRules        *UpdateApprovalRuleDetailsCollection `json:"approvalRules"`
		RepositoryAccessMode repositoryaccessmode                 `json:"repositoryAccessMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MergeSettings = model.MergeSettings

	m.ApprovalRules = model.ApprovalRules

	nn, e = model.RepositoryAccessMode.UnmarshalPolymorphicJSON(model.RepositoryAccessMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RepositoryAccessMode = nn.(RepositoryAccessMode)
	} else {
		m.RepositoryAccessMode = nil
	}

	return
}
