// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MediaWorkflowJobFact One fact of a list of facts associated to a MediaWorkflowJob that presents a point-in-time
// snapshot of the resources, data and events that were composed to generate a runnable job.
// This information will be used internally to trouble-shoot problematic workflows or jobs.
type MediaWorkflowJobFact struct {

	// Reference to the parent job.
	MediaWorkflowJobId *string `mandatory:"true" json:"mediaWorkflowJobId"`

	// System generated serial number to uniquely identify a detail in order within a MediaWorkflowJob.
	Key *int64 `mandatory:"true" json:"key"`

	// Unique name. It is read-only and generated for the fact.
	Name *string `mandatory:"true" json:"name"`

	// The type of information contained in this detail.
	Type *string `mandatory:"true" json:"type"`

	// The body of the detail captured as JSON.
	Detail map[string]interface{} `mandatory:"true" json:"detail"`
}

func (m MediaWorkflowJobFact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaWorkflowJobFact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
