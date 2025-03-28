// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProblemHistorySummary Problem history definition.
type ProblemHistorySummary struct {

	// Unique identifier for the history record
	Id *string `mandatory:"true" json:"id"`

	// Problem ID with which history is associated
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Type of actor who performed the operation
	ActorType ActorTypeEnum `mandatory:"true" json:"actorType"`

	// Resource name who performed the activity
	ActorName *string `mandatory:"true" json:"actorName"`

	// Activity explanation details
	Explanation *string `mandatory:"true" json:"explanation"`

	// Additional details on the substate of the lifecycle state
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"true" json:"lifecycleDetail"`

	// Date and time the problem was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Impacted resource names in a comma-separated string
	Delta *string `mandatory:"true" json:"delta"`

	// Event status
	EventStatus EventStatusEnum `mandatory:"false" json:"eventStatus,omitempty"`

	// User-defined comments
	Comment *string `mandatory:"false" json:"comment"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m ProblemHistorySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProblemHistorySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActorTypeEnum(string(m.ActorType)); !ok && m.ActorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActorType: %s. Supported values are: %s.", m.ActorType, strings.Join(GetActorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProblemLifecycleDetailEnum(string(m.LifecycleDetail)); !ok && m.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", m.LifecycleDetail, strings.Join(GetProblemLifecycleDetailEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEventStatusEnum(string(m.EventStatus)); !ok && m.EventStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EventStatus: %s. Supported values are: %s.", m.EventStatus, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
