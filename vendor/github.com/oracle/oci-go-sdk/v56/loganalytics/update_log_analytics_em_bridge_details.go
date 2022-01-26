// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateLogAnalyticsEmBridgeDetails Log analytics entity type definition to be updated.
type UpdateLogAnalyticsEmBridgeDetails struct {

	// Log analytics enterprise manager bridge display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description for log analytics enterprise manager bridge.
	Description *string `mandatory:"false" json:"description"`

	// Object store bucket name where enterprise manager harvested entities will be uploaded.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateLogAnalyticsEmBridgeDetails) String() string {
	return common.PointerString(m)
}
