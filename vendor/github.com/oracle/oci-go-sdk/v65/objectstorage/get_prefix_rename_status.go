// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GetPrefixRenameStatus A description of workrequest status.
type GetPrefixRenameStatus struct {

	// The type of transaction.
	TransactionType GetPrefixRenameStatusTransactionTypeEnum `mandatory:"false" json:"transactionType,omitempty"`

	// The Object Storage namespace associated with the prefix rename bucket.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The Object Storage bucketOcid associated with the prefix rename bucket.
	BucketId *string `mandatory:"false" json:"bucketId"`

	// Field to specify source path of the objects that has to be renamed.
	SourcePrefix *string `mandatory:"false" json:"sourcePrefix"`

	// Field to specify destination path of the objects to which it has to be renamed.
	DestinationPrefix *string `mandatory:"false" json:"destinationPrefix"`

	// Describes the  state at which the transaction is in.
	State GetPrefixRenameStatusStateEnum `mandatory:"false" json:"state,omitempty"`

	// Percentage of the prefix rename completed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Number of objects processed.
	ProcessedCount *int `mandatory:"false" json:"processedCount"`

	// Total number of objects.
	TotalCount *int `mandatory:"false" json:"totalCount"`

	// The date and time the prefix rename was started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the prefix rename was committed.
	TimeCommitted *common.SDKTime `mandatory:"false" json:"timeCommitted"`

	// The date and time the prefix rename was completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The date and time the prefix rename was last updated.
	TimeLastUpdated *common.SDKTime `mandatory:"false" json:"timeLastUpdated"`
}

func (m GetPrefixRenameStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GetPrefixRenameStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGetPrefixRenameStatusTransactionTypeEnum(string(m.TransactionType)); !ok && m.TransactionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransactionType: %s. Supported values are: %s.", m.TransactionType, strings.Join(GetGetPrefixRenameStatusTransactionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetPrefixRenameStatusStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetGetPrefixRenameStatusStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPrefixRenameStatusTransactionTypeEnum Enum with underlying type: string
type GetPrefixRenameStatusTransactionTypeEnum string

// Set of constants representing the allowable values for GetPrefixRenameStatusTransactionTypeEnum
const (
	GetPrefixRenameStatusTransactionTypePrefixRename GetPrefixRenameStatusTransactionTypeEnum = "PREFIX_RENAME"
)

var mappingGetPrefixRenameStatusTransactionTypeEnum = map[string]GetPrefixRenameStatusTransactionTypeEnum{
	"PREFIX_RENAME": GetPrefixRenameStatusTransactionTypePrefixRename,
}

var mappingGetPrefixRenameStatusTransactionTypeEnumLowerCase = map[string]GetPrefixRenameStatusTransactionTypeEnum{
	"prefix_rename": GetPrefixRenameStatusTransactionTypePrefixRename,
}

// GetGetPrefixRenameStatusTransactionTypeEnumValues Enumerates the set of values for GetPrefixRenameStatusTransactionTypeEnum
func GetGetPrefixRenameStatusTransactionTypeEnumValues() []GetPrefixRenameStatusTransactionTypeEnum {
	values := make([]GetPrefixRenameStatusTransactionTypeEnum, 0)
	for _, v := range mappingGetPrefixRenameStatusTransactionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPrefixRenameStatusTransactionTypeEnumStringValues Enumerates the set of values in String for GetPrefixRenameStatusTransactionTypeEnum
func GetGetPrefixRenameStatusTransactionTypeEnumStringValues() []string {
	return []string{
		"PREFIX_RENAME",
	}
}

// GetMappingGetPrefixRenameStatusTransactionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPrefixRenameStatusTransactionTypeEnum(val string) (GetPrefixRenameStatusTransactionTypeEnum, bool) {
	enum, ok := mappingGetPrefixRenameStatusTransactionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetPrefixRenameStatusStateEnum Enum with underlying type: string
type GetPrefixRenameStatusStateEnum string

// Set of constants representing the allowable values for GetPrefixRenameStatusStateEnum
const (
	GetPrefixRenameStatusStateInit       GetPrefixRenameStatusStateEnum = "INIT"
	GetPrefixRenameStatusStateEnqueued   GetPrefixRenameStatusStateEnum = "ENQUEUED"
	GetPrefixRenameStatusStateInProgress GetPrefixRenameStatusStateEnum = "IN_PROGRESS"
	GetPrefixRenameStatusStateCommitted  GetPrefixRenameStatusStateEnum = "COMMITTED"
	GetPrefixRenameStatusStateCompleted  GetPrefixRenameStatusStateEnum = "COMPLETED"
	GetPrefixRenameStatusStateCleared    GetPrefixRenameStatusStateEnum = "CLEARED"
	GetPrefixRenameStatusStateFailed     GetPrefixRenameStatusStateEnum = "FAILED"
	GetPrefixRenameStatusStateRollback   GetPrefixRenameStatusStateEnum = "ROLLBACK"
	GetPrefixRenameStatusStateAborted    GetPrefixRenameStatusStateEnum = "ABORTED"
)

var mappingGetPrefixRenameStatusStateEnum = map[string]GetPrefixRenameStatusStateEnum{
	"INIT":        GetPrefixRenameStatusStateInit,
	"ENQUEUED":    GetPrefixRenameStatusStateEnqueued,
	"IN_PROGRESS": GetPrefixRenameStatusStateInProgress,
	"COMMITTED":   GetPrefixRenameStatusStateCommitted,
	"COMPLETED":   GetPrefixRenameStatusStateCompleted,
	"CLEARED":     GetPrefixRenameStatusStateCleared,
	"FAILED":      GetPrefixRenameStatusStateFailed,
	"ROLLBACK":    GetPrefixRenameStatusStateRollback,
	"ABORTED":     GetPrefixRenameStatusStateAborted,
}

var mappingGetPrefixRenameStatusStateEnumLowerCase = map[string]GetPrefixRenameStatusStateEnum{
	"init":        GetPrefixRenameStatusStateInit,
	"enqueued":    GetPrefixRenameStatusStateEnqueued,
	"in_progress": GetPrefixRenameStatusStateInProgress,
	"committed":   GetPrefixRenameStatusStateCommitted,
	"completed":   GetPrefixRenameStatusStateCompleted,
	"cleared":     GetPrefixRenameStatusStateCleared,
	"failed":      GetPrefixRenameStatusStateFailed,
	"rollback":    GetPrefixRenameStatusStateRollback,
	"aborted":     GetPrefixRenameStatusStateAborted,
}

// GetGetPrefixRenameStatusStateEnumValues Enumerates the set of values for GetPrefixRenameStatusStateEnum
func GetGetPrefixRenameStatusStateEnumValues() []GetPrefixRenameStatusStateEnum {
	values := make([]GetPrefixRenameStatusStateEnum, 0)
	for _, v := range mappingGetPrefixRenameStatusStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPrefixRenameStatusStateEnumStringValues Enumerates the set of values in String for GetPrefixRenameStatusStateEnum
func GetGetPrefixRenameStatusStateEnumStringValues() []string {
	return []string{
		"INIT",
		"ENQUEUED",
		"IN_PROGRESS",
		"COMMITTED",
		"COMPLETED",
		"CLEARED",
		"FAILED",
		"ROLLBACK",
		"ABORTED",
	}
}

// GetMappingGetPrefixRenameStatusStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPrefixRenameStatusStateEnum(val string) (GetPrefixRenameStatusStateEnum, bool) {
	enum, ok := mappingGetPrefixRenameStatusStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
