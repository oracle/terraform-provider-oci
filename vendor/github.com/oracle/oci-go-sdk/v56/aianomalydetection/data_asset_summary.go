// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DataAssetSummary Summary of the DataAsset.
type DataAssetSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// DataAsset Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the the DataAsset was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Unique identifier for a project that is immutable on creation
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Data source type where actually data asset is being stored
	DataSourceType DataSourceTypeEnum `mandatory:"true" json:"dataSourceType"`

	// A short description of the Ai data asset
	Description *string `mandatory:"false" json:"description"`

	// The time the the DataAsset was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the data asset.
	LifecycleState DataAssetLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// OCID of Private Endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DataAssetSummary) String() string {
	return common.PointerString(m)
}
