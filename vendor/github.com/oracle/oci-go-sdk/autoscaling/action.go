// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Action The action to take if a scale event has been triggered. Positive values indicate scale out
// and negative value indicate scale in.
type Action struct {

	// Action type to take
	Type ActionTypeEnum `mandatory:"true" json:"type"`

	Value *int `mandatory:"true" json:"value"`
}

func (m Action) String() string {
	return common.PointerString(m)
}

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeBy ActionTypeEnum = "CHANGE_COUNT_BY"
)

var mappingActionType = map[string]ActionTypeEnum{
	"CHANGE_COUNT_BY": ActionTypeBy,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionType {
		values = append(values, v)
	}
	return values
}
