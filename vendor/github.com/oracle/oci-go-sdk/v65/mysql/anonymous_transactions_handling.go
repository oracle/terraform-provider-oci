// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnonymousTransactionsHandling Specifies how the replication channel handles replicated transactions without an identifier, enabling replication
// from a source that does not use transaction-id-based replication to a replica that does.
type AnonymousTransactionsHandling interface {
}

type anonymoustransactionshandling struct {
	JsonData []byte
	Policy   string `json:"policy"`
}

// UnmarshalJSON unmarshals json
func (m *anonymoustransactionshandling) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleranonymoustransactionshandling anonymoustransactionshandling
	s := struct {
		Model Unmarshaleranonymoustransactionshandling
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Policy = s.Model.Policy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *anonymoustransactionshandling) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Policy {
	case "ERROR_ON_ANONYMOUS":
		mm := ErrorOnAnonymousHandling{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASSIGN_MANUAL_UUID":
		mm := AssignManualUuidHandling{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASSIGN_TARGET_UUID":
		mm := AssignTargetUuidHandling{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m anonymoustransactionshandling) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m anonymoustransactionshandling) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnonymousTransactionsHandlingPolicyEnum Enum with underlying type: string
type AnonymousTransactionsHandlingPolicyEnum string

// Set of constants representing the allowable values for AnonymousTransactionsHandlingPolicyEnum
const (
	AnonymousTransactionsHandlingPolicyErrorOnAnonymous AnonymousTransactionsHandlingPolicyEnum = "ERROR_ON_ANONYMOUS"
	AnonymousTransactionsHandlingPolicyAssignTargetUuid AnonymousTransactionsHandlingPolicyEnum = "ASSIGN_TARGET_UUID"
	AnonymousTransactionsHandlingPolicyAssignManualUuid AnonymousTransactionsHandlingPolicyEnum = "ASSIGN_MANUAL_UUID"
)

var mappingAnonymousTransactionsHandlingPolicyEnum = map[string]AnonymousTransactionsHandlingPolicyEnum{
	"ERROR_ON_ANONYMOUS": AnonymousTransactionsHandlingPolicyErrorOnAnonymous,
	"ASSIGN_TARGET_UUID": AnonymousTransactionsHandlingPolicyAssignTargetUuid,
	"ASSIGN_MANUAL_UUID": AnonymousTransactionsHandlingPolicyAssignManualUuid,
}

var mappingAnonymousTransactionsHandlingPolicyEnumLowerCase = map[string]AnonymousTransactionsHandlingPolicyEnum{
	"error_on_anonymous": AnonymousTransactionsHandlingPolicyErrorOnAnonymous,
	"assign_target_uuid": AnonymousTransactionsHandlingPolicyAssignTargetUuid,
	"assign_manual_uuid": AnonymousTransactionsHandlingPolicyAssignManualUuid,
}

// GetAnonymousTransactionsHandlingPolicyEnumValues Enumerates the set of values for AnonymousTransactionsHandlingPolicyEnum
func GetAnonymousTransactionsHandlingPolicyEnumValues() []AnonymousTransactionsHandlingPolicyEnum {
	values := make([]AnonymousTransactionsHandlingPolicyEnum, 0)
	for _, v := range mappingAnonymousTransactionsHandlingPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetAnonymousTransactionsHandlingPolicyEnumStringValues Enumerates the set of values in String for AnonymousTransactionsHandlingPolicyEnum
func GetAnonymousTransactionsHandlingPolicyEnumStringValues() []string {
	return []string{
		"ERROR_ON_ANONYMOUS",
		"ASSIGN_TARGET_UUID",
		"ASSIGN_MANUAL_UUID",
	}
}

// GetMappingAnonymousTransactionsHandlingPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnonymousTransactionsHandlingPolicyEnum(val string) (AnonymousTransactionsHandlingPolicyEnum, bool) {
	enum, ok := mappingAnonymousTransactionsHandlingPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
