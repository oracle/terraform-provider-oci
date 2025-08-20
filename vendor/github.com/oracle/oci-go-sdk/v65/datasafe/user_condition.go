// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserCondition The audit policy provisioning conditions.
type UserCondition struct {

	// The list of users that the unified audit policy is enabled for.
	UserNames []string `mandatory:"true" json:"userNames"`

	// Specifies whether to include or exclude the specified users or roles.
	EntitySelection PolicyConditionEntitySelectionEnum `mandatory:"true" json:"entitySelection"`

	// The operation status that the policy must be enabled for.
	OperationStatus PolicyConditionOperationStatusEnum `mandatory:"true" json:"operationStatus"`
}

// GetEntitySelection returns EntitySelection
func (m UserCondition) GetEntitySelection() PolicyConditionEntitySelectionEnum {
	return m.EntitySelection
}

// GetOperationStatus returns OperationStatus
func (m UserCondition) GetOperationStatus() PolicyConditionOperationStatusEnum {
	return m.OperationStatus
}

func (m UserCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPolicyConditionEntitySelectionEnum(string(m.EntitySelection)); !ok && m.EntitySelection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntitySelection: %s. Supported values are: %s.", m.EntitySelection, strings.Join(GetPolicyConditionEntitySelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPolicyConditionOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetPolicyConditionOperationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UserCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUserCondition UserCondition
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeUserCondition
	}{
		"USER",
		(MarshalTypeUserCondition)(m),
	}

	return json.Marshal(&s)
}
