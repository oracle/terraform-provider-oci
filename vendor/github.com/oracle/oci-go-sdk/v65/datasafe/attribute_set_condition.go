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

// AttributeSetCondition The audit policy provisioning conditions.
type AttributeSetCondition struct {

	// The OCID of the attribute set.
	AttributeSetId *string `mandatory:"true" json:"attributeSetId"`

	// Specifies whether to include or exclude the specified users or roles.
	EntitySelection PolicyConditionEntitySelectionEnum `mandatory:"true" json:"entitySelection"`

	// The operation status that the policy must be enabled for.
	OperationStatus PolicyConditionOperationStatusEnum `mandatory:"true" json:"operationStatus"`
}

// GetEntitySelection returns EntitySelection
func (m AttributeSetCondition) GetEntitySelection() PolicyConditionEntitySelectionEnum {
	return m.EntitySelection
}

// GetOperationStatus returns OperationStatus
func (m AttributeSetCondition) GetOperationStatus() PolicyConditionOperationStatusEnum {
	return m.OperationStatus
}

func (m AttributeSetCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeSetCondition) ValidateEnumValue() (bool, error) {
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
func (m AttributeSetCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAttributeSetCondition AttributeSetCondition
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeAttributeSetCondition
	}{
		"ATTRIBUTE_SET",
		(MarshalTypeAttributeSetCondition)(m),
	}

	return json.Marshal(&s)
}
