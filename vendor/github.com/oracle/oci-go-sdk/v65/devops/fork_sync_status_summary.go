// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ForkSyncStatusSummary Object that contains the sync status for a specific branch name.
type ForkSyncStatusSummary struct {

	// Sync status for the provided branch.
	SyncStatus ForkSyncStatusSummarySyncStatusEnum `mandatory:"true" json:"syncStatus"`

	// The OCID of the child repository.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The branch in the child repository we are checking the sync status of.
	BranchName *string `mandatory:"true" json:"branchName"`
}

func (m ForkSyncStatusSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ForkSyncStatusSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingForkSyncStatusSummarySyncStatusEnum(string(m.SyncStatus)); !ok && m.SyncStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyncStatus: %s. Supported values are: %s.", m.SyncStatus, strings.Join(GetForkSyncStatusSummarySyncStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ForkSyncStatusSummarySyncStatusEnum Enum with underlying type: string
type ForkSyncStatusSummarySyncStatusEnum string

// Set of constants representing the allowable values for ForkSyncStatusSummarySyncStatusEnum
const (
	ForkSyncStatusSummarySyncStatusInSync         ForkSyncStatusSummarySyncStatusEnum = "IN_SYNC"
	ForkSyncStatusSummarySyncStatusSyncInProgress ForkSyncStatusSummarySyncStatusEnum = "SYNC_IN_PROGRESS"
	ForkSyncStatusSummarySyncStatusOutOfSync      ForkSyncStatusSummarySyncStatusEnum = "OUT_OF_SYNC"
)

var mappingForkSyncStatusSummarySyncStatusEnum = map[string]ForkSyncStatusSummarySyncStatusEnum{
	"IN_SYNC":          ForkSyncStatusSummarySyncStatusInSync,
	"SYNC_IN_PROGRESS": ForkSyncStatusSummarySyncStatusSyncInProgress,
	"OUT_OF_SYNC":      ForkSyncStatusSummarySyncStatusOutOfSync,
}

var mappingForkSyncStatusSummarySyncStatusEnumLowerCase = map[string]ForkSyncStatusSummarySyncStatusEnum{
	"in_sync":          ForkSyncStatusSummarySyncStatusInSync,
	"sync_in_progress": ForkSyncStatusSummarySyncStatusSyncInProgress,
	"out_of_sync":      ForkSyncStatusSummarySyncStatusOutOfSync,
}

// GetForkSyncStatusSummarySyncStatusEnumValues Enumerates the set of values for ForkSyncStatusSummarySyncStatusEnum
func GetForkSyncStatusSummarySyncStatusEnumValues() []ForkSyncStatusSummarySyncStatusEnum {
	values := make([]ForkSyncStatusSummarySyncStatusEnum, 0)
	for _, v := range mappingForkSyncStatusSummarySyncStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetForkSyncStatusSummarySyncStatusEnumStringValues Enumerates the set of values in String for ForkSyncStatusSummarySyncStatusEnum
func GetForkSyncStatusSummarySyncStatusEnumStringValues() []string {
	return []string{
		"IN_SYNC",
		"SYNC_IN_PROGRESS",
		"OUT_OF_SYNC",
	}
}

// GetMappingForkSyncStatusSummarySyncStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingForkSyncStatusSummarySyncStatusEnum(val string) (ForkSyncStatusSummarySyncStatusEnum, bool) {
	enum, ok := mappingForkSyncStatusSummarySyncStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
