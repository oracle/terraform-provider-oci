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

// FleetAdvancedFeatureConfiguration Metadata for the advanced features in the Fleet.
type FleetAdvancedFeatureConfiguration struct {

	// Namespace for the Fleet advanced feature.
	AnalyticNamespace *string `mandatory:"true" json:"analyticNamespace"`

	// Bucket name required to store JFR and related data.
	AnalyticBucketName *string `mandatory:"true" json:"analyticBucketName"`

	Lcm *Lcm `mandatory:"true" json:"lcm"`

	CryptoEventAnalysis *CryptoEventAnalysis `mandatory:"true" json:"cryptoEventAnalysis"`

	AdvancedUsageTracking *AdvancedUsageTracking `mandatory:"true" json:"advancedUsageTracking"`

	JfrRecording *JfrRecording `mandatory:"true" json:"jfrRecording"`

	PerformanceTuningAnalysis *PerformanceTuningAnalysis `mandatory:"true" json:"performanceTuningAnalysis"`

	JavaMigrationAnalysis *JavaMigrationAnalysis `mandatory:"true" json:"javaMigrationAnalysis"`

	// The date and time of the last modification to the Fleet Agent Configuration (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastModified *common.SDKTime `mandatory:"true" json:"timeLastModified"`
}

func (m FleetAdvancedFeatureConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetAdvancedFeatureConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
