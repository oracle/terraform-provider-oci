// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssessorSummary Assessor Summary
type AssessorSummary struct {

	// The OCID of the resource being referenced.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// The Assessor Name.
	Name *string `mandatory:"true" json:"name"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	// The Help URL.
	HelpLinkUrl *string `mandatory:"true" json:"helpLinkUrl"`

	// The Help link text.
	HelpLinkText *string `mandatory:"true" json:"helpLinkText"`

	// The current state of the Assessor.
	LifecycleState AssessorLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Assessor actions.
	Actions []AssessorAction `mandatory:"true" json:"actions"`

	AssessorGroup *AssessorGroup `mandatory:"false" json:"assessorGroup"`

	// The Assessor Result text.
	AssessorResult *string `mandatory:"false" json:"assessorResult"`

	// The Summary of all Checks.
	ChecksSummary *string `mandatory:"false" json:"checksSummary"`

	// True if script is available either from 'script' property of through download, false otherwise.
	HasScript *bool `mandatory:"false" json:"hasScript"`

	// The generated SQL script. Can be empty if the script exceeds maxLength.
	// In this case the property 'hasScript' indicates that the script is available for download.
	Script *string `mandatory:"false" json:"script"`

	// True if DB restart required after running the script, false otherwise.
	DoesScriptRequireRestart *bool `mandatory:"false" json:"doesScriptRequireRestart"`
}

func (m AssessorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssessorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssessorLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssessorLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
