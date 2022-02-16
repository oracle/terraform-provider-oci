// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ProblemHistorySummary Problem History Definition.
type ProblemHistorySummary struct {

	// Unique identifier for the history record
	Id *string `mandatory:"true" json:"id"`

	// problemId for which history is associated to.
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Actor type who performed the operation
	ActorType ActorTypeEnum `mandatory:"true" json:"actorType"`

	// Resource Name who performed activity
	ActorName *string `mandatory:"true" json:"actorName"`

	// Activity explanation details
	Explanation *string `mandatory:"true" json:"explanation"`

	// Problem Lifecycle Detail Status
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"true" json:"lifecycleDetail"`

	// Type of the Entity
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Impacted Resource Names in a comma-separated string.
	Delta *string `mandatory:"true" json:"delta"`

	// Event status
	EventStatus EventStatusEnum `mandatory:"false" json:"eventStatus,omitempty"`

	// User Defined Comments
	Comment *string `mandatory:"false" json:"comment"`
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
