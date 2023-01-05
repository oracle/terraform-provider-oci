// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// LiveMigration The status entry of live migration operation.
type LiveMigration struct {

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
	Status LiveMigrationStatusEnum `mandatory:"true" json:"status"`

	// The general status of the live migration operation.
	// *   **NOT_INITIATED:**  Rollback operation not initiated..
	// *   **SUCCESS:**  Successfully completed the live migration rollback operation on failure.
	// *   **UPDATING:** The live migration rollback operation is in progress.
	// *   **FAILED:** The live migration rollback operation failed in processing.
	// *   **NOT_SUPPORTED:** The live migration rollback operation not possible at current step of processing.
	RollbackStatus LiveMigrationRollbackStatusEnum `mandatory:"true" json:"rollbackStatus"`

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

	// The source slot id.
	// Example: `1233`
	SourceSlotId *int `mandatory:"false" json:"sourceSlotId"`

	// The destination slot id.
	// Example: `1244`
	DestinationSlotId *int `mandatory:"false" json:"destinationSlotId"`

	// Detailed error reason.
	Details *string `mandatory:"false" json:"details"`
}

func (m LiveMigration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LiveMigration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLiveMigrationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetLiveMigrationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLiveMigrationRollbackStatusEnum(string(m.RollbackStatus)); !ok && m.RollbackStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RollbackStatus: %s. Supported values are: %s.", m.RollbackStatus, strings.Join(GetLiveMigrationRollbackStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LiveMigrationStatusEnum Enum with underlying type: string
type LiveMigrationStatusEnum string

// Set of constants representing the allowable values for LiveMigrationStatusEnum
const (
	LiveMigrationStatusSuccess        LiveMigrationStatusEnum = "SUCCESS"
	LiveMigrationStatusPartialSuccess LiveMigrationStatusEnum = "PARTIAL_SUCCESS"
	LiveMigrationStatusUpdating       LiveMigrationStatusEnum = "UPDATING"
	LiveMigrationStatusFailed         LiveMigrationStatusEnum = "FAILED"
)

var mappingLiveMigrationStatusEnum = map[string]LiveMigrationStatusEnum{
	"SUCCESS":         LiveMigrationStatusSuccess,
	"PARTIAL_SUCCESS": LiveMigrationStatusPartialSuccess,
	"UPDATING":        LiveMigrationStatusUpdating,
	"FAILED":          LiveMigrationStatusFailed,
}

var mappingLiveMigrationStatusEnumLowerCase = map[string]LiveMigrationStatusEnum{
	"success":         LiveMigrationStatusSuccess,
	"partial_success": LiveMigrationStatusPartialSuccess,
	"updating":        LiveMigrationStatusUpdating,
	"failed":          LiveMigrationStatusFailed,
}

// GetLiveMigrationStatusEnumValues Enumerates the set of values for LiveMigrationStatusEnum
func GetLiveMigrationStatusEnumValues() []LiveMigrationStatusEnum {
	values := make([]LiveMigrationStatusEnum, 0)
	for _, v := range mappingLiveMigrationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLiveMigrationStatusEnumStringValues Enumerates the set of values in String for LiveMigrationStatusEnum
func GetLiveMigrationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"PARTIAL_SUCCESS",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingLiveMigrationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLiveMigrationStatusEnum(val string) (LiveMigrationStatusEnum, bool) {
	enum, ok := mappingLiveMigrationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LiveMigrationRollbackStatusEnum Enum with underlying type: string
type LiveMigrationRollbackStatusEnum string

// Set of constants representing the allowable values for LiveMigrationRollbackStatusEnum
const (
	LiveMigrationRollbackStatusNotInitiated LiveMigrationRollbackStatusEnum = "NOT_INITIATED"
	LiveMigrationRollbackStatusSuccess      LiveMigrationRollbackStatusEnum = "SUCCESS"
	LiveMigrationRollbackStatusUpdating     LiveMigrationRollbackStatusEnum = "UPDATING"
	LiveMigrationRollbackStatusFailed       LiveMigrationRollbackStatusEnum = "FAILED"
	LiveMigrationRollbackStatusNotSupported LiveMigrationRollbackStatusEnum = "NOT_SUPPORTED"
)

var mappingLiveMigrationRollbackStatusEnum = map[string]LiveMigrationRollbackStatusEnum{
	"NOT_INITIATED": LiveMigrationRollbackStatusNotInitiated,
	"SUCCESS":       LiveMigrationRollbackStatusSuccess,
	"UPDATING":      LiveMigrationRollbackStatusUpdating,
	"FAILED":        LiveMigrationRollbackStatusFailed,
	"NOT_SUPPORTED": LiveMigrationRollbackStatusNotSupported,
}

var mappingLiveMigrationRollbackStatusEnumLowerCase = map[string]LiveMigrationRollbackStatusEnum{
	"not_initiated": LiveMigrationRollbackStatusNotInitiated,
	"success":       LiveMigrationRollbackStatusSuccess,
	"updating":      LiveMigrationRollbackStatusUpdating,
	"failed":        LiveMigrationRollbackStatusFailed,
	"not_supported": LiveMigrationRollbackStatusNotSupported,
}

// GetLiveMigrationRollbackStatusEnumValues Enumerates the set of values for LiveMigrationRollbackStatusEnum
func GetLiveMigrationRollbackStatusEnumValues() []LiveMigrationRollbackStatusEnum {
	values := make([]LiveMigrationRollbackStatusEnum, 0)
	for _, v := range mappingLiveMigrationRollbackStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLiveMigrationRollbackStatusEnumStringValues Enumerates the set of values in String for LiveMigrationRollbackStatusEnum
func GetLiveMigrationRollbackStatusEnumStringValues() []string {
	return []string{
		"NOT_INITIATED",
		"SUCCESS",
		"UPDATING",
		"FAILED",
		"NOT_SUPPORTED",
	}
}

// GetMappingLiveMigrationRollbackStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLiveMigrationRollbackStatusEnum(val string) (LiveMigrationRollbackStatusEnum, bool) {
	enum, ok := mappingLiveMigrationRollbackStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
