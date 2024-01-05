// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoAnalysisResult Metadata for the result of a crypto event analysis. The analysis result is stored in an Object Storage bucket.
type CryptoAnalysisResult struct {

	// The OCID to identify this analysis results.
	Id *string `mandatory:"true" json:"id"`

	// The result aggregation mode
	AggregationMode CryptoAnalysisResultModeEnum `mandatory:"true" json:"aggregationMode"`

	// The fleet OCID.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// Total number of events in the analysis.
	TotalEventCount *int `mandatory:"true" json:"totalEventCount"`

	// Total number of summarized events. Summarized events are deduplicated events of interest.
	SummarizedEventCount *int `mandatory:"true" json:"summarizedEventCount"`

	// Total number of findings with the analysis.
	FindingCount *int `mandatory:"true" json:"findingCount"`

	// Total number of non-compliant findings with the analysis. A non-compliant finding means the
	// application won't work properly with the changes introduced by the Crypto Roadmap version
	// used by the analysis.
	NonCompliantFindingCount *int `mandatory:"true" json:"nonCompliantFindingCount"`

	// The Crypto Roadmap version used to perform the analysis.
	CryptoRoadmapVersion *string `mandatory:"true" json:"cryptoRoadmapVersion"`

	// The Object Storage namespace of this analysis result.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Object Storage bucket name of this analysis result.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage object name of this analysis result.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The OCID of the work request to start the analysis.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The managed instance OCID.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The hostname of the managed instance.
	HostName *string `mandatory:"false" json:"hostName"`

	// Time of the first event in the analysis.
	TimeFirstEvent *common.SDKTime `mandatory:"false" json:"timeFirstEvent"`

	// Time of the last event in the analysis.
	TimeLastEvent *common.SDKTime `mandatory:"false" json:"timeLastEvent"`

	// The time the result is compiled.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m CryptoAnalysisResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAnalysisResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAnalysisResultModeEnum(string(m.AggregationMode)); !ok && m.AggregationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AggregationMode: %s. Supported values are: %s.", m.AggregationMode, strings.Join(GetCryptoAnalysisResultModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
