// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeSqlPlanBaselinesAttributesDetails The details required to change SQL plan baseline attributes.
// It takes either credentials or databaseCredential. It's recommended to provide databaseCredential
type ChangeSqlPlanBaselinesAttributesDetails struct {

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

	Credentials ManagedDatabaseCredential `mandatory:"false" json:"credentials"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`
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
		SqlHandle          *string                   `json:"sqlHandle"`
		PlanName           *string                   `json:"planName"`
		IsEnabled          *bool                     `json:"isEnabled"`
		IsFixed            *bool                     `json:"isFixed"`
		IsAutoPurged       *bool                     `json:"isAutoPurged"`
		Credentials        manageddatabasecredential `json:"credentials"`
		DatabaseCredential databasecredentialdetails `json:"databaseCredential"`
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

	nn, e = model.DatabaseCredential.UnmarshalPolymorphicJSON(model.DatabaseCredential.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseCredential = nn.(DatabaseCredentialDetails)
	} else {
		m.DatabaseCredential = nil
	}

	return
}
