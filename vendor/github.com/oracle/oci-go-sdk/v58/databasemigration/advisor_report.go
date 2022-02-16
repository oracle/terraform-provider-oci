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

// AdvisorReport Pre-Migration advisor report details.
type AdvisorReport struct {

	// Pre-Migration advisor result.
	Result AdvisorResultsEnum `mandatory:"true" json:"result"`

	// Number of Fatal results in the advisor report.
	NumberOfFatal *int `mandatory:"true" json:"numberOfFatal"`

	// Number of Fatal Blocker results in the advisor report.
	NumberOfFatalBlockers *int `mandatory:"true" json:"numberOfFatalBlockers"`

	// Number of Warning results in the advisor report.
	NumberOfWarnings *int `mandatory:"true" json:"numberOfWarnings"`

	// Number of Informational results in the advisor report.
	NumberOfInformationalResults *int `mandatory:"true" json:"numberOfInformationalResults"`

	ReportLocationDetails *AdvisorReportLocationDetails `mandatory:"false" json:"reportLocationDetails"`
}

func (m AdvisorReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdvisorReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdvisorResultsEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetAdvisorResultsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
