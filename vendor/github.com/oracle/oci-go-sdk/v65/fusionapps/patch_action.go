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

// PatchAction Monthly patch details.
type PatchAction struct {

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	// Unique identifier of the object that represents the action
	ReferenceKey *string `mandatory:"false" json:"referenceKey"`

	// patch bundle name
	Artifact *string `mandatory:"false" json:"artifact"`

	// A string that describeds whether the change is applied hot or cold
	Mode PatchActionModeEnum `mandatory:"false" json:"mode,omitempty"`

	// patch artifact category
	Category PatchActionCategoryEnum `mandatory:"false" json:"category,omitempty"`

	// A string that describes whether the change is applied hot or cold
	State ActionStateEnum `mandatory:"false" json:"state,omitempty"`
}

// GetReferenceKey returns ReferenceKey
func (m PatchAction) GetReferenceKey() *string {
	return m.ReferenceKey
}

// GetState returns State
func (m PatchAction) GetState() ActionStateEnum {
	return m.State
}

// GetDescription returns Description
func (m PatchAction) GetDescription() *string {
	return m.Description
}

func (m PatchAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchActionModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetPatchActionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchActionCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetPatchActionCategoryEnumStringValues(), ",")))
	}

	if _, ok := GetMappingActionStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetActionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchAction PatchAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypePatchAction
	}{
		"PATCH",
		(MarshalTypePatchAction)(m),
	}

	return json.Marshal(&s)
}

// PatchActionModeEnum Enum with underlying type: string
type PatchActionModeEnum string

// Set of constants representing the allowable values for PatchActionModeEnum
const (
	PatchActionModeHot  PatchActionModeEnum = "HOT"
	PatchActionModeCold PatchActionModeEnum = "COLD"
)

var mappingPatchActionModeEnum = map[string]PatchActionModeEnum{
	"HOT":  PatchActionModeHot,
	"COLD": PatchActionModeCold,
}

var mappingPatchActionModeEnumLowerCase = map[string]PatchActionModeEnum{
	"hot":  PatchActionModeHot,
	"cold": PatchActionModeCold,
}

// GetPatchActionModeEnumValues Enumerates the set of values for PatchActionModeEnum
func GetPatchActionModeEnumValues() []PatchActionModeEnum {
	values := make([]PatchActionModeEnum, 0)
	for _, v := range mappingPatchActionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchActionModeEnumStringValues Enumerates the set of values in String for PatchActionModeEnum
func GetPatchActionModeEnumStringValues() []string {
	return []string{
		"HOT",
		"COLD",
	}
}

// GetMappingPatchActionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchActionModeEnum(val string) (PatchActionModeEnum, bool) {
	enum, ok := mappingPatchActionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchActionCategoryEnum Enum with underlying type: string
type PatchActionCategoryEnum string

// Set of constants representing the allowable values for PatchActionCategoryEnum
const (
	PatchActionCategoryMonthly PatchActionCategoryEnum = "MONTHLY"
	PatchActionCategoryWeekly  PatchActionCategoryEnum = "WEEKLY"
	PatchActionCategoryOneoff  PatchActionCategoryEnum = "ONEOFF"
)

var mappingPatchActionCategoryEnum = map[string]PatchActionCategoryEnum{
	"MONTHLY": PatchActionCategoryMonthly,
	"WEEKLY":  PatchActionCategoryWeekly,
	"ONEOFF":  PatchActionCategoryOneoff,
}

var mappingPatchActionCategoryEnumLowerCase = map[string]PatchActionCategoryEnum{
	"monthly": PatchActionCategoryMonthly,
	"weekly":  PatchActionCategoryWeekly,
	"oneoff":  PatchActionCategoryOneoff,
}

// GetPatchActionCategoryEnumValues Enumerates the set of values for PatchActionCategoryEnum
func GetPatchActionCategoryEnumValues() []PatchActionCategoryEnum {
	values := make([]PatchActionCategoryEnum, 0)
	for _, v := range mappingPatchActionCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchActionCategoryEnumStringValues Enumerates the set of values in String for PatchActionCategoryEnum
func GetPatchActionCategoryEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"WEEKLY",
		"ONEOFF",
	}
}

// GetMappingPatchActionCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchActionCategoryEnum(val string) (PatchActionCategoryEnum, bool) {
	enum, ok := mappingPatchActionCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
