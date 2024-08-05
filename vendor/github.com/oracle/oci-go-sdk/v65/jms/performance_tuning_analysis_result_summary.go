// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PerformanceTuningAnalysisResultSummary Summary of a performance tuning analysis result. The actual output of the analysis is stored in the Object Storage object.
type PerformanceTuningAnalysisResultSummary struct {

	// The OCID to identify this analysis results.
	Id *string `mandatory:"true" json:"id"`

	// The fleet OCID.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The OCID of the application for which the report has been generated.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The internal identifier of the application installation for which the report has been generated.
	ApplicationInstallationId *string `mandatory:"true" json:"applicationInstallationId"`

	// The installation path of the application for which the report has been generated.
	ApplicationInstallationPath *string `mandatory:"true" json:"applicationInstallationPath"`

	// Total number of warnings reported by the analysis.
	WarningCount *int `mandatory:"true" json:"warningCount"`

	// Result of the analysis based on whether warnings have been found or not.
	Result PerformanceTuningResultStatusEnum `mandatory:"true" json:"result"`

	// The managed instance OCID.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The hostname of the managed instance.
	HostName *string `mandatory:"true" json:"hostName"`

	// The name of the application for which the report has been generated.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// The Object Storage namespace of this analysis result.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Object Storage bucket name of this analysis result.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage object name of this analysis result.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The time the result is compiled.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the JFR recording has started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The time the JFR recording has finished.
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// The OCID of the work request to start the analysis.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`
}

func (m PerformanceTuningAnalysisResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PerformanceTuningAnalysisResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPerformanceTuningResultStatusEnum(string(m.Result)); !ok && m.Result != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Result: %s. Supported values are: %s.", m.Result, strings.Join(GetPerformanceTuningResultStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
