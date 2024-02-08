// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourcePowerAction A power action against a resource.
type ResourcePowerAction struct {
	Action ResourcePowerActionActionEnum `mandatory:"true" json:"action"`
}

func (m ResourcePowerAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcePowerAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourcePowerActionActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetResourcePowerActionActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ResourcePowerAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeResourcePowerAction ResourcePowerAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeResourcePowerAction
	}{
		"power",
		(MarshalTypeResourcePowerAction)(m),
	}

	return json.Marshal(&s)
}

// ResourcePowerActionActionEnum Enum with underlying type: string
type ResourcePowerActionActionEnum string

// Set of constants representing the allowable values for ResourcePowerActionActionEnum
const (
	ResourcePowerActionActionStop      ResourcePowerActionActionEnum = "STOP"
	ResourcePowerActionActionStart     ResourcePowerActionActionEnum = "START"
	ResourcePowerActionActionSoftreset ResourcePowerActionActionEnum = "SOFTRESET"
	ResourcePowerActionActionReset     ResourcePowerActionActionEnum = "RESET"
)

var mappingResourcePowerActionActionEnum = map[string]ResourcePowerActionActionEnum{
	"STOP":      ResourcePowerActionActionStop,
	"START":     ResourcePowerActionActionStart,
	"SOFTRESET": ResourcePowerActionActionSoftreset,
	"RESET":     ResourcePowerActionActionReset,
}

var mappingResourcePowerActionActionEnumLowerCase = map[string]ResourcePowerActionActionEnum{
	"stop":      ResourcePowerActionActionStop,
	"start":     ResourcePowerActionActionStart,
	"softreset": ResourcePowerActionActionSoftreset,
	"reset":     ResourcePowerActionActionReset,
}

// GetResourcePowerActionActionEnumValues Enumerates the set of values for ResourcePowerActionActionEnum
func GetResourcePowerActionActionEnumValues() []ResourcePowerActionActionEnum {
	values := make([]ResourcePowerActionActionEnum, 0)
	for _, v := range mappingResourcePowerActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetResourcePowerActionActionEnumStringValues Enumerates the set of values in String for ResourcePowerActionActionEnum
func GetResourcePowerActionActionEnumStringValues() []string {
	return []string{
		"STOP",
		"START",
		"SOFTRESET",
		"RESET",
	}
}

// GetMappingResourcePowerActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourcePowerActionActionEnum(val string) (ResourcePowerActionActionEnum, bool) {
	enum, ok := mappingResourcePowerActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
