// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedDatabaseRuChangeLogSummary One move RU invocation record for the Globally distributed database.
type DistributedDatabaseRuChangeLogSummary struct {

	// Source shard for move RU invocation.
	SourceShardName *string `mandatory:"false" json:"sourceShardName"`

	// Target shard for move RU invocation.
	TargetShardName *string `mandatory:"false" json:"targetShardName"`

	// Replication unit identifier associated with the invocation.
	RuId *string `mandatory:"false" json:"ruId"`

	// Current or terminal status of the move RU invocation.
	Status DistributedDatabaseRuChangeLogSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Actor who triggered the move RU invocation.
	TriggeredBy *string `mandatory:"false" json:"triggeredBy"`

	// Entry creation timestamp in UTC.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Entry update timestamp in UTC.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DistributedDatabaseRuChangeLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseRuChangeLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedDatabaseRuChangeLogSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseRuChangeLogSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseRuChangeLogSummaryStatusEnum Enum with underlying type: string
type DistributedDatabaseRuChangeLogSummaryStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseRuChangeLogSummaryStatusEnum
const (
	DistributedDatabaseRuChangeLogSummaryStatusFailed     DistributedDatabaseRuChangeLogSummaryStatusEnum = "FAILED"
	DistributedDatabaseRuChangeLogSummaryStatusQueued     DistributedDatabaseRuChangeLogSummaryStatusEnum = "QUEUED"
	DistributedDatabaseRuChangeLogSummaryStatusInProgress DistributedDatabaseRuChangeLogSummaryStatusEnum = "IN_PROGRESS"
	DistributedDatabaseRuChangeLogSummaryStatusCompleted  DistributedDatabaseRuChangeLogSummaryStatusEnum = "COMPLETED"
)

var mappingDistributedDatabaseRuChangeLogSummaryStatusEnum = map[string]DistributedDatabaseRuChangeLogSummaryStatusEnum{
	"FAILED":      DistributedDatabaseRuChangeLogSummaryStatusFailed,
	"QUEUED":      DistributedDatabaseRuChangeLogSummaryStatusQueued,
	"IN_PROGRESS": DistributedDatabaseRuChangeLogSummaryStatusInProgress,
	"COMPLETED":   DistributedDatabaseRuChangeLogSummaryStatusCompleted,
}

var mappingDistributedDatabaseRuChangeLogSummaryStatusEnumLowerCase = map[string]DistributedDatabaseRuChangeLogSummaryStatusEnum{
	"failed":      DistributedDatabaseRuChangeLogSummaryStatusFailed,
	"queued":      DistributedDatabaseRuChangeLogSummaryStatusQueued,
	"in_progress": DistributedDatabaseRuChangeLogSummaryStatusInProgress,
	"completed":   DistributedDatabaseRuChangeLogSummaryStatusCompleted,
}

// GetDistributedDatabaseRuChangeLogSummaryStatusEnumValues Enumerates the set of values for DistributedDatabaseRuChangeLogSummaryStatusEnum
func GetDistributedDatabaseRuChangeLogSummaryStatusEnumValues() []DistributedDatabaseRuChangeLogSummaryStatusEnum {
	values := make([]DistributedDatabaseRuChangeLogSummaryStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseRuChangeLogSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseRuChangeLogSummaryStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseRuChangeLogSummaryStatusEnum
func GetDistributedDatabaseRuChangeLogSummaryStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"QUEUED",
		"IN_PROGRESS",
		"COMPLETED",
	}
}

// GetMappingDistributedDatabaseRuChangeLogSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseRuChangeLogSummaryStatusEnum(val string) (DistributedDatabaseRuChangeLogSummaryStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseRuChangeLogSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
