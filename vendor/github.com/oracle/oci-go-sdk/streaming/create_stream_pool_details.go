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

// CreateStreamPoolDetails Object used to create a stream pool.
type CreateStreamPoolDetails struct {

	// The OCID of the compartment that contains the stream.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the stream pool. Avoid entering confidential information.
	// Example: `MyStreamPool`
	Name *string `mandatory:"true" json:"name"`

	KafkaSettings *KafkaSettings `mandatory:"false" json:"kafkaSettings"`

	CustomEncryptionKeyDetails *CustomEncryptionKeyDetails `mandatory:"false" json:"customEncryptionKeyDetails"`

	PrivateEndpointDetails *PrivateEndpointDetails `mandatory:"false" json:"privateEndpointDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateStreamPoolDetails) String() string {
	return common.PointerString(m)
}
