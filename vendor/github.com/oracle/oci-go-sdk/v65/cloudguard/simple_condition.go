// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SimpleCondition Simple Condition object.
type SimpleCondition struct {

	// parameter Key
	Parameter *string `mandatory:"false" json:"parameter"`

	// type of operator
	Value *string `mandatory:"false" json:"value"`

	// type of operator
	Operator OperatorTypeEnum `mandatory:"false" json:"operator,omitempty"`

	// type of value
	ValueType ConditionValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

func (m SimpleCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SimpleCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOperatorTypeEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetOperatorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConditionValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetConditionValueTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SimpleCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSimpleCondition SimpleCondition
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeSimpleCondition
	}{
		"SIMPLE",
		(MarshalTypeSimpleCondition)(m),
	}

	return json.Marshal(&s)
}
