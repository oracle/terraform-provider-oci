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

// AssessorCheckSummary Assessor Check Summary
type AssessorCheckSummary struct {

	// The Name of the Check.
	Name *string `mandatory:"true" json:"name"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	AssessorCheckGroup *AssessorCheckGroup `mandatory:"true" json:"assessorCheckGroup"`

	// The current state of the Assessor Check.
	AssessorCheckState AssessorCheckStatesEnum `mandatory:"true" json:"assessorCheckState"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The Help URL.
	HelpLinkUrl *string `mandatory:"false" json:"helpLinkUrl"`

	// The Help link text.
	HelpLinkText *string `mandatory:"false" json:"helpLinkText"`

	CheckAction *AssessorCheckAction `mandatory:"false" json:"checkAction"`

	// Pre-Migration сheck id.
	Key *string `mandatory:"false" json:"key"`

	// Description of the issue.
	Issue *string `mandatory:"false" json:"issue"`

	// Impact of the issue on data migration.
	Impact *string `mandatory:"false" json:"impact"`

	// Fixing the issue.
	Action *string `mandatory:"false" json:"action"`

	// The path to the fixup script for this check.
	FixupScriptLocation *string `mandatory:"false" json:"fixupScriptLocation"`

	// If false, objects cannot be excluded from migration.
	IsExclusionAllowed *bool `mandatory:"false" json:"isExclusionAllowed"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// Array of the column of the objects table.
	Columns []AdvisorReportCheckColumn `mandatory:"false" json:"columns"`

	// Number of database objects to migrate.
	ObjectCount *int `mandatory:"false" json:"objectCount"`

	// The objects display name.
	ObjectsDisplayName *string `mandatory:"false" json:"objectsDisplayName"`

	LogLocation *LogLocationBucketDetails `mandatory:"false" json:"logLocation"`
}

func (m AssessorCheckSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssessorCheckSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssessorCheckStatesEnum(string(m.AssessorCheckState)); !ok && m.AssessorCheckState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessorCheckState: %s. Supported values are: %s.", m.AssessorCheckState, strings.Join(GetAssessorCheckStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
