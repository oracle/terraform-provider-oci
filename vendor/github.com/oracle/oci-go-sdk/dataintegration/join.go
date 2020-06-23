// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Join The information about the join operator. The join operator links data from multiple inbound sources.
type Join struct {

	// The join condition.
	Condition *string `mandatory:"false" json:"condition"`

	// The type of join.
	Policy JoinPolicyEnum `mandatory:"false" json:"policy,omitempty"`
}

func (m Join) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Join) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJoin Join
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeJoin
	}{
		"JOIN",
		(MarshalTypeJoin)(m),
	}

	return json.Marshal(&s)
}

// JoinPolicyEnum Enum with underlying type: string
type JoinPolicyEnum string

// Set of constants representing the allowable values for JoinPolicyEnum
const (
	JoinPolicyInnerJoin JoinPolicyEnum = "INNER_JOIN"
	JoinPolicyLeftJoin  JoinPolicyEnum = "LEFT_JOIN"
	JoinPolicyRightJoin JoinPolicyEnum = "RIGHT_JOIN"
	JoinPolicyFullJoin  JoinPolicyEnum = "FULL_JOIN"
)

var mappingJoinPolicy = map[string]JoinPolicyEnum{
	"INNER_JOIN": JoinPolicyInnerJoin,
	"LEFT_JOIN":  JoinPolicyLeftJoin,
	"RIGHT_JOIN": JoinPolicyRightJoin,
	"FULL_JOIN":  JoinPolicyFullJoin,
}

// GetJoinPolicyEnumValues Enumerates the set of values for JoinPolicyEnum
func GetJoinPolicyEnumValues() []JoinPolicyEnum {
	values := make([]JoinPolicyEnum, 0)
	for _, v := range mappingJoinPolicy {
		values = append(values, v)
	}
	return values
}
