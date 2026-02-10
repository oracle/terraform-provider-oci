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

// AdvisorReportCheckSummary Pre-Migration extended advisor report check item.
type AdvisorReportCheckSummary struct {

	// Pre-Migration сheck id.
	Key *string `mandatory:"true" json:"key"`

	// Pre-Migration сheck display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pre-Migration advisor result.
	ResultType AdvisorResultsEnum `mandatory:"true" json:"resultType"`

	// User flag for advisor report check.
	IsReviewed *bool `mandatory:"true" json:"isReviewed"`

	// Description of the issue.
	Issue *string `mandatory:"true" json:"issue"`

	// Impact of the issue on data migration.
	Impact *string `mandatory:"true" json:"impact"`

	// Fixing the issue.
	Action *string `mandatory:"true" json:"action"`

	// If false, objects cannot be excluded from migration.
	IsExclusionAllowed *bool `mandatory:"true" json:"isExclusionAllowed"`

	// Array of the column of the objects table.
	Columns []AdvisorReportCheckColumn `mandatory:"true" json:"columns"`

	// Number of database objects to migrate.
	ObjectCount *int `mandatory:"true" json:"objectCount"`

	// The path to the fixup script for this check.
	FixupScriptLocation *string `mandatory:"false" json:"fixupScriptLocation"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m AdvisorReportCheckSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdvisorReportCheckSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdvisorResultsEnum(string(m.ResultType)); !ok && m.ResultType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResultType: %s. Supported values are: %s.", m.ResultType, strings.Join(GetAdvisorResultsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
