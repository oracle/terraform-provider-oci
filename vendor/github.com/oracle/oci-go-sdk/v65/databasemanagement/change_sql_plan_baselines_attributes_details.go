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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeSqlPlanBaselinesAttributesDetails The details required to change SQL plan baseline attributes.
type ChangeSqlPlanBaselinesAttributesDetails struct {
	Credentials ManagedDatabaseCredential `mandatory:"true" json:"credentials"`

	// The SQL statement handle. It identifies plans associated with a SQL statement
	// for attribute changes. If `null` then `planName` must be specified.
	SqlHandle *string `mandatory:"false" json:"sqlHandle"`

	// Then plan name. It identifies a specific plan. If `null' then all plans associated
	// with a SQL statement identified by `sqlHandle' are considered for attribute changes.
	PlanName *string `mandatory:"false" json:"planName"`

	// Indicates whether the plan is available for use by the optimizer.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Indicates whether the plan baseline is fixed. A fixed plan takes precedence over a non-fixed plan.
	IsFixed *bool `mandatory:"false" json:"isFixed"`

	// Indicates whether the plan is purged if it is not used for a time period.
	IsAutoPurged *bool `mandatory:"false" json:"isAutoPurged"`
}

func (m ChangeSqlPlanBaselinesAttributesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeSqlPlanBaselinesAttributesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChangeSqlPlanBaselinesAttributesDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SqlHandle    *string                   `json:"sqlHandle"`
		PlanName     *string                   `json:"planName"`
		IsEnabled    *bool                     `json:"isEnabled"`
		IsFixed      *bool                     `json:"isFixed"`
		IsAutoPurged *bool                     `json:"isAutoPurged"`
		Credentials  manageddatabasecredential `json:"credentials"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SqlHandle = model.SqlHandle

	m.PlanName = model.PlanName

	m.IsEnabled = model.IsEnabled

	m.IsFixed = model.IsFixed

	m.IsAutoPurged = model.IsAutoPurged

	nn, e = model.Credentials.UnmarshalPolymorphicJSON(model.Credentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Credentials = nn.(ManagedDatabaseCredential)
	} else {
		m.Credentials = nil
	}

	return
}
