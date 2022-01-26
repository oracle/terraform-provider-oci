// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ServiceCatalogSummary The model for a summary of an Oracle Cloud Infrastructure service catalog.
type ServiceCatalogSummary struct {

	// The unique identifier for the Service catalog.
	Id *string `mandatory:"true" json:"id"`

	// The lifecycle state of the service catalog.
	LifecycleState ServiceCatalogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Compartment id where the service catalog exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the service catalog.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time this service catalog was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m ServiceCatalogSummary) String() string {
	return common.PointerString(m)
}
