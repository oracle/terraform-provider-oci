// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring (APM) Control Plane API
//
// Provide a set of APIs for tenant to perform operations like create, update, delete and list APM domains, and also
// work request APIs to monitor progress of these operations.
//

package apmcontrolplane

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// BaseDomainDetails Basic details for an APM Domain.
type BaseDomainDetails struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// APM Domain display name, can be updated.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment corresponding to the APM Domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of the APM Domain.
	Description *string `mandatory:"false" json:"description"`

	// The current lifecycle state of the APM Domain.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates if this is an Always Free resource.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The time the the APM Domain was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the APM Domain was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BaseDomainDetails) String() string {
	return common.PointerString(m)
}
