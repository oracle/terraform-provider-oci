// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Action Action details
type Action interface {

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information.
	GetDescription() *string

	// Unique identifier of the object that represents the action
	GetReferenceKey() *string

	// A string that describes whether the change is applied hot or cold
	GetState() ActionStateEnum
}

type action struct {
	JsonData     []byte
	ReferenceKey *string         `mandatory:"false" json:"referenceKey"`
	State        ActionStateEnum `mandatory:"false" json:"state,omitempty"`
	Description  *string         `mandatory:"true" json:"description"`
	ActionType   string          `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.ReferenceKey = s.Model.ReferenceKey
	m.State = s.Model.State
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "PATCH":
		mm := PatchAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "QUARTERLY_UPGRADE":
		mm := UpgradeAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERTEX":
		mm := VertexAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Action: %s.", m.ActionType)
		return *m, nil
	}
}

// GetReferenceKey returns ReferenceKey
func (m action) GetReferenceKey() *string {
	return m.ReferenceKey
}

// GetState returns State
func (m action) GetState() ActionStateEnum {
	return m.State
}

// GetDescription returns Description
func (m action) GetDescription() *string {
	return m.Description
}

func (m action) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m action) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingActionStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetActionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionStateEnum Enum with underlying type: string
type ActionStateEnum string

// Set of constants representing the allowable values for ActionStateEnum
const (
	ActionStateAccepted   ActionStateEnum = "ACCEPTED"
	ActionStateInProgress ActionStateEnum = "IN_PROGRESS"
	ActionStateSucceeded  ActionStateEnum = "SUCCEEDED"
	ActionStateFailed     ActionStateEnum = "FAILED"
	ActionStateCanceled   ActionStateEnum = "CANCELED"
)

var mappingActionStateEnum = map[string]ActionStateEnum{
	"ACCEPTED":    ActionStateAccepted,
	"IN_PROGRESS": ActionStateInProgress,
	"SUCCEEDED":   ActionStateSucceeded,
	"FAILED":      ActionStateFailed,
	"CANCELED":    ActionStateCanceled,
}

var mappingActionStateEnumLowerCase = map[string]ActionStateEnum{
	"accepted":    ActionStateAccepted,
	"in_progress": ActionStateInProgress,
	"succeeded":   ActionStateSucceeded,
	"failed":      ActionStateFailed,
	"canceled":    ActionStateCanceled,
}

// GetActionStateEnumValues Enumerates the set of values for ActionStateEnum
func GetActionStateEnumValues() []ActionStateEnum {
	values := make([]ActionStateEnum, 0)
	for _, v := range mappingActionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetActionStateEnumStringValues Enumerates the set of values in String for ActionStateEnum
func GetActionStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

// GetMappingActionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionStateEnum(val string) (ActionStateEnum, bool) {
	enum, ok := mappingActionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ActionActionTypeEnum Enum with underlying type: string
type ActionActionTypeEnum string

// Set of constants representing the allowable values for ActionActionTypeEnum
const (
	ActionActionTypeQuarterlyUpgrade ActionActionTypeEnum = "QUARTERLY_UPGRADE"
	ActionActionTypePatch            ActionActionTypeEnum = "PATCH"
	ActionActionTypeVertex           ActionActionTypeEnum = "VERTEX"
)

var mappingActionActionTypeEnum = map[string]ActionActionTypeEnum{
	"QUARTERLY_UPGRADE": ActionActionTypeQuarterlyUpgrade,
	"PATCH":             ActionActionTypePatch,
	"VERTEX":            ActionActionTypeVertex,
}

var mappingActionActionTypeEnumLowerCase = map[string]ActionActionTypeEnum{
	"quarterly_upgrade": ActionActionTypeQuarterlyUpgrade,
	"patch":             ActionActionTypePatch,
	"vertex":            ActionActionTypeVertex,
}

// GetActionActionTypeEnumValues Enumerates the set of values for ActionActionTypeEnum
func GetActionActionTypeEnumValues() []ActionActionTypeEnum {
	values := make([]ActionActionTypeEnum, 0)
	for _, v := range mappingActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionActionTypeEnumStringValues Enumerates the set of values in String for ActionActionTypeEnum
func GetActionActionTypeEnumStringValues() []string {
	return []string{
		"QUARTERLY_UPGRADE",
		"PATCH",
		"VERTEX",
	}
}

// GetMappingActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionActionTypeEnum(val string) (ActionActionTypeEnum, bool) {
	enum, ok := mappingActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
