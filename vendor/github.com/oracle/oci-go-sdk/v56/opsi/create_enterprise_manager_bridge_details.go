// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateEnterpriseManagerBridgeDetails The information about a Enterprise Manager bridge resource to be created
type CreateEnterpriseManagerBridgeDetails struct {

	// Compartment identifier of the Enterprise Manager bridge
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friedly name of Enterprise Manager Bridge that does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Object Storage Bucket Name
	ObjectStorageBucketName *string `mandatory:"true" json:"objectStorageBucketName"`

	// Description of Enterprise Manager Bridge
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEnterpriseManagerBridgeDetails) String() string {
	return common.PointerString(m)
}
