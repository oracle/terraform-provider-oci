// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicApiEndpointDecommissionStatus Information regarding a cluster's public api endpoint decommission.
type PublicApiEndpointDecommissionStatus struct {

	// The date and time of rollback deadline for public api endpoint decommission.
	// Once the date is passed, rollback is not able to be launched.
	TimeDecommissionRollbackDeadline *common.SDKTime `mandatory:"true" json:"timeDecommissionRollbackDeadline"`

	// The current public api endpoint decommission status of the cluster.
	Status PublicApiEndpointDecommissionStatusStatusEnum `mandatory:"true" json:"status"`
}

func (m PublicApiEndpointDecommissionStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicApiEndpointDecommissionStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPublicApiEndpointDecommissionStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPublicApiEndpointDecommissionStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicApiEndpointDecommissionStatusStatusEnum Enum with underlying type: string
type PublicApiEndpointDecommissionStatusStatusEnum string

// Set of constants representing the allowable values for PublicApiEndpointDecommissionStatusStatusEnum
const (
	PublicApiEndpointDecommissionStatusStatusPending            PublicApiEndpointDecommissionStatusStatusEnum = "PENDING"
	PublicApiEndpointDecommissionStatusStatusInProgress         PublicApiEndpointDecommissionStatusStatusEnum = "IN_PROGRESS"
	PublicApiEndpointDecommissionStatusStatusRollingBack        PublicApiEndpointDecommissionStatusStatusEnum = "ROLLING_BACK"
	PublicApiEndpointDecommissionStatusStatusDecommissioned     PublicApiEndpointDecommissionStatusStatusEnum = "DECOMMISSIONED"
	PublicApiEndpointDecommissionStatusStatusFinalized          PublicApiEndpointDecommissionStatusStatusEnum = "FINALIZED"
	PublicApiEndpointDecommissionStatusStatusDecommissionFailed PublicApiEndpointDecommissionStatusStatusEnum = "DECOMMISSION_FAILED"
	PublicApiEndpointDecommissionStatusStatusRollbackFailed     PublicApiEndpointDecommissionStatusStatusEnum = "ROLLBACK_FAILED"
)

var mappingPublicApiEndpointDecommissionStatusStatusEnum = map[string]PublicApiEndpointDecommissionStatusStatusEnum{
	"PENDING":             PublicApiEndpointDecommissionStatusStatusPending,
	"IN_PROGRESS":         PublicApiEndpointDecommissionStatusStatusInProgress,
	"ROLLING_BACK":        PublicApiEndpointDecommissionStatusStatusRollingBack,
	"DECOMMISSIONED":      PublicApiEndpointDecommissionStatusStatusDecommissioned,
	"FINALIZED":           PublicApiEndpointDecommissionStatusStatusFinalized,
	"DECOMMISSION_FAILED": PublicApiEndpointDecommissionStatusStatusDecommissionFailed,
	"ROLLBACK_FAILED":     PublicApiEndpointDecommissionStatusStatusRollbackFailed,
}

var mappingPublicApiEndpointDecommissionStatusStatusEnumLowerCase = map[string]PublicApiEndpointDecommissionStatusStatusEnum{
	"pending":             PublicApiEndpointDecommissionStatusStatusPending,
	"in_progress":         PublicApiEndpointDecommissionStatusStatusInProgress,
	"rolling_back":        PublicApiEndpointDecommissionStatusStatusRollingBack,
	"decommissioned":      PublicApiEndpointDecommissionStatusStatusDecommissioned,
	"finalized":           PublicApiEndpointDecommissionStatusStatusFinalized,
	"decommission_failed": PublicApiEndpointDecommissionStatusStatusDecommissionFailed,
	"rollback_failed":     PublicApiEndpointDecommissionStatusStatusRollbackFailed,
}

// GetPublicApiEndpointDecommissionStatusStatusEnumValues Enumerates the set of values for PublicApiEndpointDecommissionStatusStatusEnum
func GetPublicApiEndpointDecommissionStatusStatusEnumValues() []PublicApiEndpointDecommissionStatusStatusEnum {
	values := make([]PublicApiEndpointDecommissionStatusStatusEnum, 0)
	for _, v := range mappingPublicApiEndpointDecommissionStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicApiEndpointDecommissionStatusStatusEnumStringValues Enumerates the set of values in String for PublicApiEndpointDecommissionStatusStatusEnum
func GetPublicApiEndpointDecommissionStatusStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"IN_PROGRESS",
		"ROLLING_BACK",
		"DECOMMISSIONED",
		"FINALIZED",
		"DECOMMISSION_FAILED",
		"ROLLBACK_FAILED",
	}
}

// GetMappingPublicApiEndpointDecommissionStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicApiEndpointDecommissionStatusStatusEnum(val string) (PublicApiEndpointDecommissionStatusStatusEnum, bool) {
	enum, ok := mappingPublicApiEndpointDecommissionStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
