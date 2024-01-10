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

// ImplementOptimizerStatisticsAdvisorRecommendationsJob The job request details to implement the Optimizer Statistics Advisor task recommendations.
type ImplementOptimizerStatisticsAdvisorRecommendationsJob struct {

	// The name of the job. Valid characters are uppercase or lowercase letters,
	// numbers, and "_". The name of the job cannot be modified. It must be unique
	// in the compartment and must begin with an alphabetic character.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ResultLocation JobExecutionResultLocation `mandatory:"true" json:"resultLocation"`

	// The name of the execution.
	Description *string `mandatory:"false" json:"description"`

	Credentials ManagedDatabaseCredential `mandatory:"false" json:"credentials"`
}

func (m ImplementOptimizerStatisticsAdvisorRecommendationsJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImplementOptimizerStatisticsAdvisorRecommendationsJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ImplementOptimizerStatisticsAdvisorRecommendationsJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                    `json:"description"`
		Credentials    manageddatabasecredential  `json:"credentials"`
		Name           *string                    `json:"name"`
		CompartmentId  *string                    `json:"compartmentId"`
		ResultLocation jobexecutionresultlocation `json:"resultLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.Credentials.UnmarshalPolymorphicJSON(model.Credentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Credentials = nn.(ManagedDatabaseCredential)
	} else {
		m.Credentials = nil
	}

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(JobExecutionResultLocation)
	} else {
		m.ResultLocation = nil
	}

	return
}
