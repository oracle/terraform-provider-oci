// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Join) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJoinPolicyEnum(string(m.Policy)); !ok && m.Policy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Policy: %s. Supported values are: %s.", m.Policy, strings.Join(GetJoinPolicyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingJoinPolicyEnum = map[string]JoinPolicyEnum{
	"INNER_JOIN": JoinPolicyInnerJoin,
	"LEFT_JOIN":  JoinPolicyLeftJoin,
	"RIGHT_JOIN": JoinPolicyRightJoin,
	"FULL_JOIN":  JoinPolicyFullJoin,
}

var mappingJoinPolicyEnumLowerCase = map[string]JoinPolicyEnum{
	"inner_join": JoinPolicyInnerJoin,
	"left_join":  JoinPolicyLeftJoin,
	"right_join": JoinPolicyRightJoin,
	"full_join":  JoinPolicyFullJoin,
}

// GetJoinPolicyEnumValues Enumerates the set of values for JoinPolicyEnum
func GetJoinPolicyEnumValues() []JoinPolicyEnum {
	values := make([]JoinPolicyEnum, 0)
	for _, v := range mappingJoinPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetJoinPolicyEnumStringValues Enumerates the set of values in String for JoinPolicyEnum
func GetJoinPolicyEnumStringValues() []string {
	return []string{
		"INNER_JOIN",
		"LEFT_JOIN",
		"RIGHT_JOIN",
		"FULL_JOIN",
	}
}

// GetMappingJoinPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJoinPolicyEnum(val string) (JoinPolicyEnum, bool) {
	enum, ok := mappingJoinPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
