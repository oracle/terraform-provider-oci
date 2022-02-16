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

// MigrationJobProgressResource Progress details of a Migration Job.
type MigrationJobProgressResource struct {

	// Current status of the job.
	CurrentStatus JobPhaseStatusEnum `mandatory:"true" json:"currentStatus"`

	// Current phase of the job.
	CurrentPhase OdmsJobPhasesEnum `mandatory:"true" json:"currentPhase"`

	// List of phase status for the job.
	Phases []PhaseStatus `mandatory:"true" json:"phases"`
}

func (m MigrationJobProgressResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationJobProgressResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobPhaseStatusEnum(string(m.CurrentStatus)); !ok && m.CurrentStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentStatus: %s. Supported values are: %s.", m.CurrentStatus, strings.Join(GetJobPhaseStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdmsJobPhasesEnum(string(m.CurrentPhase)); !ok && m.CurrentPhase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentPhase: %s. Supported values are: %s.", m.CurrentPhase, strings.Join(GetOdmsJobPhasesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
