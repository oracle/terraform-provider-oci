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

// DistributedAutonomousDatabaseRuChangeLogSummary One move RU invocation record for the Globally distributed autonomous database.
type DistributedAutonomousDatabaseRuChangeLogSummary struct {

	// Source shard for move RU invocation.
	SourceShardName *string `mandatory:"false" json:"sourceShardName"`

	// Target shard for move RU invocation.
	TargetShardName *string `mandatory:"false" json:"targetShardName"`

	// Replication unit identifier associated with the invocation.
	RuId *string `mandatory:"false" json:"ruId"`

	// Current or terminal status of the move RU invocation.
	Status DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Actor who triggered the move RU invocation.
	TriggeredBy *string `mandatory:"false" json:"triggeredBy"`

	// Entry creation timestamp in UTC.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Entry update timestamp in UTC.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DistributedAutonomousDatabaseRuChangeLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseRuChangeLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum Enum with underlying type: string
type DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum
const (
	DistributedAutonomousDatabaseRuChangeLogSummaryStatusFailed     DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum = "FAILED"
	DistributedAutonomousDatabaseRuChangeLogSummaryStatusQueued     DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum = "QUEUED"
	DistributedAutonomousDatabaseRuChangeLogSummaryStatusInProgress DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum = "IN_PROGRESS"
	DistributedAutonomousDatabaseRuChangeLogSummaryStatusCompleted  DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum = "COMPLETED"
)

var mappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum = map[string]DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum{
	"FAILED":      DistributedAutonomousDatabaseRuChangeLogSummaryStatusFailed,
	"QUEUED":      DistributedAutonomousDatabaseRuChangeLogSummaryStatusQueued,
	"IN_PROGRESS": DistributedAutonomousDatabaseRuChangeLogSummaryStatusInProgress,
	"COMPLETED":   DistributedAutonomousDatabaseRuChangeLogSummaryStatusCompleted,
}

var mappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumLowerCase = map[string]DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum{
	"failed":      DistributedAutonomousDatabaseRuChangeLogSummaryStatusFailed,
	"queued":      DistributedAutonomousDatabaseRuChangeLogSummaryStatusQueued,
	"in_progress": DistributedAutonomousDatabaseRuChangeLogSummaryStatusInProgress,
	"completed":   DistributedAutonomousDatabaseRuChangeLogSummaryStatusCompleted,
}

// GetDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumValues Enumerates the set of values for DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum
func GetDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumValues() []DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum {
	values := make([]DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum
func GetDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"QUEUED",
		"IN_PROGRESS",
		"COMPLETED",
	}
}

// GetMappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum(val string) (DistributedAutonomousDatabaseRuChangeLogSummaryStatusEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseRuChangeLogSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
