// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PhaseStatus Job phase status details.
type PhaseStatus struct {

	// Phase name
	Name OdmsJobPhasesEnum `mandatory:"true" json:"name"`

	// Phase status
	Status JobPhaseStatusEnum `mandatory:"true" json:"status"`

	// Duration of the phase in milliseconds
	DurationInMs *int `mandatory:"true" json:"durationInMs"`

	// True if a Pre-Migration Advisor report is available for this phase. False or null if no report is available.
	IsAdvisorReportAvailable *bool `mandatory:"false" json:"isAdvisorReportAvailable"`

	// Summary of phase status results.
	Extract []PhaseExtractEntry `mandatory:"false" json:"extract"`

	LogLocation *LogLocationBucketDetails `mandatory:"false" json:"logLocation"`

	// Percent progress of job phase.
	Progress *int `mandatory:"false" json:"progress"`
}

func (m PhaseStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PhaseStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdmsJobPhasesEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetOdmsJobPhasesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobPhaseStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobPhaseStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
