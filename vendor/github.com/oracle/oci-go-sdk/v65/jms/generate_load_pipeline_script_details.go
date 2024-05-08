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

// GenerateLoadPipelineScriptDetails Attributes to generate load pipeline script.
type GenerateLoadPipelineScriptDetails struct {

	// The name of the bucket where data will be exported.
	TargetBucketName *string `mandatory:"true" json:"targetBucketName"`

	// The namespace of the bucket where data will be exported.
	TargetBucketNamespace *string `mandatory:"true" json:"targetBucketNamespace"`

	// The id of the region of the target bucket.
	TargetBucketRegion *string `mandatory:"true" json:"targetBucketRegion"`

	// The time internal in minutes between consecutive executions of scheduled pipeline job.
	IntervalMinutes *int `mandatory:"false" json:"intervalMinutes"`
}

func (m GenerateLoadPipelineScriptDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateLoadPipelineScriptDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
