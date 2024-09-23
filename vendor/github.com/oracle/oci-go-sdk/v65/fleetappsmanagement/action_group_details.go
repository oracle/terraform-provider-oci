// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActionGroupDetails Action Group.
type ActionGroupDetails struct {

	// The ID of the ActionGroup resource .
	// Ex:fleetId.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// ID of the runbook
	RunbookId *string `mandatory:"true" json:"runbookId"`

	// Name of the ActionGroup.
	Name *string `mandatory:"false" json:"name"`

	// Type of the ActionGroup
	Type LifeCycleActionGroupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Application Type associated.
	// Only applicable if actionGroup type is ENVIRONMENT.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// Product associated.
	// Only applicable if actionGroup type is PRODUCT.
	Product *string `mandatory:"false" json:"product"`

	// LifeCycle Operation
	LifecycleOperation *string `mandatory:"false" json:"lifecycleOperation"`

	// Unique producer Id at Action Group Level
	ActivityId *string `mandatory:"false" json:"activityId"`

	// Status of the Job at Action Group Level
	Status JobStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The time the the Scheduler Job started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Scheduler Job ended. An RFC3339 formatted datetime string
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m ActionGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActionGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifeCycleActionGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetLifeCycleActionGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
