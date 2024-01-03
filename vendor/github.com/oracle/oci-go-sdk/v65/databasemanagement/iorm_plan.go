// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IormPlan The IORM plan from an Exadata storage server.
type IormPlan struct {

	// The status of the IORM plan.
	PlanStatus IormPlanStatusEnumEnum `mandatory:"true" json:"planStatus"`

	// The objective of the IORM plan.
	PlanObjective IormPlanObjectiveEnumEnum `mandatory:"true" json:"planObjective"`

	DbPlan *DatabasePlan `mandatory:"false" json:"dbPlan"`
}

func (m IormPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IormPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIormPlanStatusEnumEnum(string(m.PlanStatus)); !ok && m.PlanStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanStatus: %s. Supported values are: %s.", m.PlanStatus, strings.Join(GetIormPlanStatusEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIormPlanObjectiveEnumEnum(string(m.PlanObjective)); !ok && m.PlanObjective != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanObjective: %s. Supported values are: %s.", m.PlanObjective, strings.Join(GetIormPlanObjectiveEnumEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
