// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicateSchemaChange Options required for pipeline Initial Data Load. If enabled, copies existing data from source to target before replication.
type ReplicateSchemaChange struct {

	// If ENABLED, then addition or removal of schema is also replicated, apart from individual tables and records when creating or updating the pipeline.
	CanReplicateSchemaChange ReplicateSchemaChangeCanReplicateSchemaChangeEnum `mandatory:"true" json:"canReplicateSchemaChange"`

	// Action upon DDL Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true
	ActionOnDdlError ReplicateDdlErrorActionEnum `mandatory:"false" json:"actionOnDdlError,omitempty"`

	// Action upon DML Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true
	ActionOnDmlError ReplicateDmlErrorActionEnum `mandatory:"false" json:"actionOnDmlError,omitempty"`
}

func (m ReplicateSchemaChange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicateSchemaChange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicateSchemaChangeCanReplicateSchemaChangeEnum(string(m.CanReplicateSchemaChange)); !ok && m.CanReplicateSchemaChange != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CanReplicateSchemaChange: %s. Supported values are: %s.", m.CanReplicateSchemaChange, strings.Join(GetReplicateSchemaChangeCanReplicateSchemaChangeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReplicateDdlErrorActionEnum(string(m.ActionOnDdlError)); !ok && m.ActionOnDdlError != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionOnDdlError: %s. Supported values are: %s.", m.ActionOnDdlError, strings.Join(GetReplicateDdlErrorActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReplicateDmlErrorActionEnum(string(m.ActionOnDmlError)); !ok && m.ActionOnDmlError != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionOnDmlError: %s. Supported values are: %s.", m.ActionOnDmlError, strings.Join(GetReplicateDmlErrorActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicateSchemaChangeCanReplicateSchemaChangeEnum Enum with underlying type: string
type ReplicateSchemaChangeCanReplicateSchemaChangeEnum string

// Set of constants representing the allowable values for ReplicateSchemaChangeCanReplicateSchemaChangeEnum
const (
	ReplicateSchemaChangeCanReplicateSchemaChangeEnabled  ReplicateSchemaChangeCanReplicateSchemaChangeEnum = "ENABLED"
	ReplicateSchemaChangeCanReplicateSchemaChangeDisabled ReplicateSchemaChangeCanReplicateSchemaChangeEnum = "DISABLED"
)

var mappingReplicateSchemaChangeCanReplicateSchemaChangeEnum = map[string]ReplicateSchemaChangeCanReplicateSchemaChangeEnum{
	"ENABLED":  ReplicateSchemaChangeCanReplicateSchemaChangeEnabled,
	"DISABLED": ReplicateSchemaChangeCanReplicateSchemaChangeDisabled,
}

var mappingReplicateSchemaChangeCanReplicateSchemaChangeEnumLowerCase = map[string]ReplicateSchemaChangeCanReplicateSchemaChangeEnum{
	"enabled":  ReplicateSchemaChangeCanReplicateSchemaChangeEnabled,
	"disabled": ReplicateSchemaChangeCanReplicateSchemaChangeDisabled,
}

// GetReplicateSchemaChangeCanReplicateSchemaChangeEnumValues Enumerates the set of values for ReplicateSchemaChangeCanReplicateSchemaChangeEnum
func GetReplicateSchemaChangeCanReplicateSchemaChangeEnumValues() []ReplicateSchemaChangeCanReplicateSchemaChangeEnum {
	values := make([]ReplicateSchemaChangeCanReplicateSchemaChangeEnum, 0)
	for _, v := range mappingReplicateSchemaChangeCanReplicateSchemaChangeEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicateSchemaChangeCanReplicateSchemaChangeEnumStringValues Enumerates the set of values in String for ReplicateSchemaChangeCanReplicateSchemaChangeEnum
func GetReplicateSchemaChangeCanReplicateSchemaChangeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingReplicateSchemaChangeCanReplicateSchemaChangeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicateSchemaChangeCanReplicateSchemaChangeEnum(val string) (ReplicateSchemaChangeCanReplicateSchemaChangeEnum, bool) {
	enum, ok := mappingReplicateSchemaChangeCanReplicateSchemaChangeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
