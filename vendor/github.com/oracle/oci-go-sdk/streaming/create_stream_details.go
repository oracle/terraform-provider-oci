// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateStreamDetails Object used to create a stream.
type CreateStreamDetails struct {

	// The name of the stream. Avoid entering confidential information.
	// Example: `TelemetryEvents`
	Name *string `mandatory:"true" json:"name"`

	// The number of partitions in the stream.
	Partitions *int `mandatory:"true" json:"partitions"`

	// The OCID of the compartment that contains the stream.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the stream pool that contains the stream.
	StreamPoolId *string `mandatory:"false" json:"streamPoolId"`

	// The retention period of the stream, in hours. Accepted values are between 24 and 168 (7 days).
	// If not specified, the stream will have a retention period of 24 hours.
	RetentionInHours *int `mandatory:"false" json:"retentionInHours"`

	// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateStreamDetails) String() string {
	return common.PointerString(m)
}
