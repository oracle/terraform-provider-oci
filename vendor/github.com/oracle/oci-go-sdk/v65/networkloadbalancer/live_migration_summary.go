// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LiveMigrationSummary Network load balancer live migration object to be used for list operations.
type LiveMigrationSummary struct {

	// Unique identifier for the live migration operation.
	// Example: `ocid1.coreservicesworkrequest.oc1.iad.aaaaaaaa2dzekghuaygw5vjslup47huywt6omihmfilp75xcnp2rlktkdd2q`
	WorkRequestId *string `mandatory:"true" json:"workRequestId"`

	// The source pod id from where the live migration operation is performed.
	// Example: `10.0.0.1`
	SourcePodId *string `mandatory:"true" json:"sourcePodId"`

	// The destination pod id to where the live migration is performed.
	// Example: `10.0.0.2`
	DestinationPodId *string `mandatory:"true" json:"destinationPodId"`

	// The general status of the live migration operation.
	// *   **SUCCESS:**  Successfully completed the live migration operation.
	// *   **PARTIAL_SUCCESS:**  Completed the live migration operation with pending source detach.
	// *   **UPDATING:** The live migration operation is in progress.
	// *   **FAILED:** The live migration operation failed in processing.
	Status LiveMigrationSummaryStatusEnum `mandatory:"true" json:"status"`

	// The general status of the live migration operation.
	// *   **NOT_INITIATED:**  Rollback operation not initiated..
	// *   **SUCCESS:**  Successfully completed the live migration rollback operation on failure.
	// *   **UPDATING:** The live migration rollback operation is in progress.
	// *   **FAILED:** The live migration rollback operation failed in processing.
	// *   **NOT_SUPPORTED:** The live migration rollback operation not possible at current step of processing.
	RollbackStatus LiveMigrationSummaryRollbackStatusEnum `mandatory:"true" json:"rollbackStatus"`

	// The step on which the live migration failure happened.
	FailureStep *string `mandatory:"true" json:"failureStep"`

	// Unique identifier for the nlb resource.
	// Example: `ocid1.networkloadbalancer.oc1.iad.aaaaaaaa2dzekghuaygw5vjslup47huywt6omihmfilp75xcnp2rlktkdd2q`
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Unique identifier for the service VNIC.
	// Example: `ocid1.vnic.oc1.iad.aaaaaaaa2dzekghuaygw5vjslup47huywt6omihmfilp75xcnp2rlktkdd2q`
	ServiceVnicId *string `mandatory:"false" json:"serviceVnicId"`

	// The date and time the data was created, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the data was created, in the format defined by RFC3339.
	// Example: `2020-05-01T18:28:11+00:00`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The destination pod name to where the live migration is performed.
	// Example: `pod2`
	DestinationPodName *string `mandatory:"false" json:"destinationPodName"`

	// The destination shard name to where the live migration is performed.
	// Example: `shard2`
	DestinationShardName *string `mandatory:"false" json:"destinationShardName"`

	// The source slot id.
	// Example: `1233`
	SourceSlotId *int `mandatory:"false" json:"sourceSlotId"`

	// The destination slot id.
	// Example: `1244`
	DestinationSlotId *int `mandatory:"false" json:"destinationSlotId"`

	// Detailed error reason.
	Details *string `mandatory:"false" json:"details"`
}

func (m LiveMigrationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LiveMigrationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLiveMigrationSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLiveMigrationSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLiveMigrationSummaryRollbackStatusEnum(string(m.RollbackStatus)); !ok && m.RollbackStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RollbackStatus: %s. Supported values are: %s.", m.RollbackStatus, strings.Join(GetLiveMigrationSummaryRollbackStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LiveMigrationSummaryStatusEnum Enum with underlying type: string
type LiveMigrationSummaryStatusEnum string

// Set of constants representing the allowable values for LiveMigrationSummaryStatusEnum
const (
	LiveMigrationSummaryStatusSuccess        LiveMigrationSummaryStatusEnum = "SUCCESS"
	LiveMigrationSummaryStatusPartialSuccess LiveMigrationSummaryStatusEnum = "PARTIAL_SUCCESS"
	LiveMigrationSummaryStatusUpdating       LiveMigrationSummaryStatusEnum = "UPDATING"
	LiveMigrationSummaryStatusFailed         LiveMigrationSummaryStatusEnum = "FAILED"
)

var mappingLiveMigrationSummaryStatusEnum = map[string]LiveMigrationSummaryStatusEnum{
	"SUCCESS":         LiveMigrationSummaryStatusSuccess,
	"PARTIAL_SUCCESS": LiveMigrationSummaryStatusPartialSuccess,
	"UPDATING":        LiveMigrationSummaryStatusUpdating,
	"FAILED":          LiveMigrationSummaryStatusFailed,
}

var mappingLiveMigrationSummaryStatusEnumLowerCase = map[string]LiveMigrationSummaryStatusEnum{
	"success":         LiveMigrationSummaryStatusSuccess,
	"partial_success": LiveMigrationSummaryStatusPartialSuccess,
	"updating":        LiveMigrationSummaryStatusUpdating,
	"failed":          LiveMigrationSummaryStatusFailed,
}

// GetLiveMigrationSummaryStatusEnumValues Enumerates the set of values for LiveMigrationSummaryStatusEnum
func GetLiveMigrationSummaryStatusEnumValues() []LiveMigrationSummaryStatusEnum {
	values := make([]LiveMigrationSummaryStatusEnum, 0)
	for _, v := range mappingLiveMigrationSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLiveMigrationSummaryStatusEnumStringValues Enumerates the set of values in String for LiveMigrationSummaryStatusEnum
func GetLiveMigrationSummaryStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"PARTIAL_SUCCESS",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingLiveMigrationSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLiveMigrationSummaryStatusEnum(val string) (LiveMigrationSummaryStatusEnum, bool) {
	enum, ok := mappingLiveMigrationSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LiveMigrationSummaryRollbackStatusEnum Enum with underlying type: string
type LiveMigrationSummaryRollbackStatusEnum string

// Set of constants representing the allowable values for LiveMigrationSummaryRollbackStatusEnum
const (
	LiveMigrationSummaryRollbackStatusNotInitiated LiveMigrationSummaryRollbackStatusEnum = "NOT_INITIATED"
	LiveMigrationSummaryRollbackStatusSuccess      LiveMigrationSummaryRollbackStatusEnum = "SUCCESS"
	LiveMigrationSummaryRollbackStatusUpdating     LiveMigrationSummaryRollbackStatusEnum = "UPDATING"
	LiveMigrationSummaryRollbackStatusFailed       LiveMigrationSummaryRollbackStatusEnum = "FAILED"
	LiveMigrationSummaryRollbackStatusNotSupported LiveMigrationSummaryRollbackStatusEnum = "NOT_SUPPORTED"
)

var mappingLiveMigrationSummaryRollbackStatusEnum = map[string]LiveMigrationSummaryRollbackStatusEnum{
	"NOT_INITIATED": LiveMigrationSummaryRollbackStatusNotInitiated,
	"SUCCESS":       LiveMigrationSummaryRollbackStatusSuccess,
	"UPDATING":      LiveMigrationSummaryRollbackStatusUpdating,
	"FAILED":        LiveMigrationSummaryRollbackStatusFailed,
	"NOT_SUPPORTED": LiveMigrationSummaryRollbackStatusNotSupported,
}

var mappingLiveMigrationSummaryRollbackStatusEnumLowerCase = map[string]LiveMigrationSummaryRollbackStatusEnum{
	"not_initiated": LiveMigrationSummaryRollbackStatusNotInitiated,
	"success":       LiveMigrationSummaryRollbackStatusSuccess,
	"updating":      LiveMigrationSummaryRollbackStatusUpdating,
	"failed":        LiveMigrationSummaryRollbackStatusFailed,
	"not_supported": LiveMigrationSummaryRollbackStatusNotSupported,
}

// GetLiveMigrationSummaryRollbackStatusEnumValues Enumerates the set of values for LiveMigrationSummaryRollbackStatusEnum
func GetLiveMigrationSummaryRollbackStatusEnumValues() []LiveMigrationSummaryRollbackStatusEnum {
	values := make([]LiveMigrationSummaryRollbackStatusEnum, 0)
	for _, v := range mappingLiveMigrationSummaryRollbackStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLiveMigrationSummaryRollbackStatusEnumStringValues Enumerates the set of values in String for LiveMigrationSummaryRollbackStatusEnum
func GetLiveMigrationSummaryRollbackStatusEnumStringValues() []string {
	return []string{
		"NOT_INITIATED",
		"SUCCESS",
		"UPDATING",
		"FAILED",
		"NOT_SUPPORTED",
	}
}

// GetMappingLiveMigrationSummaryRollbackStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLiveMigrationSummaryRollbackStatusEnum(val string) (LiveMigrationSummaryRollbackStatusEnum, bool) {
	enum, ok := mappingLiveMigrationSummaryRollbackStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
