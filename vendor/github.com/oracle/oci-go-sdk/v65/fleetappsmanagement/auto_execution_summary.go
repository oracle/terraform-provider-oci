// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoExecutionSummary Summary of the Auto Executions.
type AutoExecutionSummary struct {

	// The unique id of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SchedulerJob.
	LifecycleState AutoExecutionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the compartment in which resource exist.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// Lifecycle operation of given execution.
	LifecycleOperation *string `mandatory:"false" json:"lifecycleOperation"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The scheduled date and time for the Job.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// FleetId associated with the execution.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Display Name of the Fleet associated with the execution.
	ResourceDisplayName *string `mandatory:"false" json:"resourceDisplayName"`

	// Display Name of the product.
	ProductName *string `mandatory:"false" json:"productName"`

	// RunbookId associated with the execution.
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the Runbook version associated with the execution.
	RunbookVersionName *string `mandatory:"false" json:"runbookVersionName"`

	// Display name of Runbook associated with the execution.
	RunbookDisplayName *string `mandatory:"false" json:"runbookDisplayName"`

	// Actual start date and time for the Execution.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Actual end date and time for the Execution.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Initiation type for Execution.
	InitiationType AutoExecutionInitiationTypeEnum `mandatory:"false" json:"initiationType,omitempty"`
}

func (m AutoExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoExecutionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutoExecutionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutoExecutionInitiationTypeEnum(string(m.InitiationType)); !ok && m.InitiationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InitiationType: %s. Supported values are: %s.", m.InitiationType, strings.Join(GetAutoExecutionInitiationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
