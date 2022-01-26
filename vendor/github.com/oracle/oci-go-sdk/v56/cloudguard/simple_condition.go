// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
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
