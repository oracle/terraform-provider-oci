// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PerformanceTuningAnalysisSummary Summary information about a Performance Tuning Analysis.
type PerformanceTuningAnalysisSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Performance Tuning Analysis.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the  Work Request.
	WorkRequestId *string `mandatory:"true" json:"workRequestId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the analysis project.
	AnalysisProjectName *string `mandatory:"true" json:"analysisProjectName"`

	// Number of warnings in the Performance Tuning Analysis.
	WarningCount *int `mandatory:"true" json:"warningCount"`

	// Possible Performance Tuning Result statuses.
	Result PerformanceTuningAnalysisResultEnum `mandatory:"true" json:"result"`

	// Object storage path to the analysis.
	ResultObjectStoragePath *string `mandatory:"true" json:"resultObjectStoragePath"`

	// Object storage path to the artifact.
	ArtifactObjectStoragePath *string `mandatory:"true" json:"artifactObjectStoragePath"`

	// The date and time the Performance Tuning Analysis was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Performance Tuning Analysis was started, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The date and time the Performance Tuning Analysis was finished, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	CreatedBy *Principal `mandatory:"true" json:"createdBy"`
}

func (m PerformanceTuningAnalysisSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PerformanceTuningAnalysisSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPerformanceTuningAnalysisResultEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetPerformanceTuningAnalysisResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
