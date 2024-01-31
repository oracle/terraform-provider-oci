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

// LoadSqlPlanBaselinesFromAwrDetails The details required to load plans from Automatic Workload Repository (AWR).
// It takes either credentials or databaseCredential. It's recommended to provide databaseCredential
type LoadSqlPlanBaselinesFromAwrDetails struct {

	// The name of the database job used for loading SQL plan baselines.
	JobName *string `mandatory:"true" json:"jobName"`

	// The begin snapshot.
	BeginSnapshot *int `mandatory:"true" json:"beginSnapshot"`

	// The end snapshot.
	EndSnapshot *int `mandatory:"true" json:"endSnapshot"`

	// The description of the job.
	JobDescription *string `mandatory:"false" json:"jobDescription"`

	// A filter applied to AWR to select only qualifying plans to be loaded.
	// By default all plans in AWR are selected. The filter can take the form of
	// any `WHERE` clause predicate that can be specified against the column
	// `DBA_HIST_SQLTEXT.SQL_TEXT`. An example is `sql_text like 'SELECT %'`.
	SqlTextFilter *string `mandatory:"false" json:"sqlTextFilter"`

	// Indicates whether the plans are loaded as fixed plans (`true`) or non-fixed plans (`false`).
	// By default, they are loaded as non-fixed plans.
	IsFixed *bool `mandatory:"false" json:"isFixed"`

	// Indicates whether the loaded plans are enabled (`true`) or not (`false`).
	// By default, they are enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Credentials ManagedDatabaseCredential `mandatory:"false" json:"credentials"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`
}

func (m LoadSqlPlanBaselinesFromAwrDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadSqlPlanBaselinesFromAwrDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LoadSqlPlanBaselinesFromAwrDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		JobDescription     *string                   `json:"jobDescription"`
		SqlTextFilter      *string                   `json:"sqlTextFilter"`
		IsFixed            *bool                     `json:"isFixed"`
		IsEnabled          *bool                     `json:"isEnabled"`
		Credentials        manageddatabasecredential `json:"credentials"`
		DatabaseCredential databasecredentialdetails `json:"databaseCredential"`
		JobName            *string                   `json:"jobName"`
		BeginSnapshot      *int                      `json:"beginSnapshot"`
		EndSnapshot        *int                      `json:"endSnapshot"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.JobDescription = model.JobDescription

	m.SqlTextFilter = model.SqlTextFilter

	m.IsFixed = model.IsFixed

	m.IsEnabled = model.IsEnabled

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

	m.JobName = model.JobName

	m.BeginSnapshot = model.BeginSnapshot

	m.EndSnapshot = model.EndSnapshot

	return
}
