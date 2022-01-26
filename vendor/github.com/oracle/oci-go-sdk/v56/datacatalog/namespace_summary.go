// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// NamespaceSummary Summary of a namespace
type NamespaceSummary struct {

	// Unique namespace key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// Name of the namespace
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the namespace.
	Description *string `mandatory:"false" json:"description"`

	// If this field is defined by service or by a user
	IsServiceDefined *bool `mandatory:"false" json:"isServiceDefined"`

	// The current state of the namespace.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the namespace was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m NamespaceSummary) String() string {
	return common.PointerString(m)
}
