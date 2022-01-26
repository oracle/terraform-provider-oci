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

// ServiceCatalogAssociationSummary The model for a summary of a service catalog association.
type ServiceCatalogAssociationSummary struct {

	// The unique identifier of the service catalog association.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier of the service catalog.
	ServiceCatalogId *string `mandatory:"true" json:"serviceCatalogId"`

	// The unique identifier of the resource being associated to service catalog.
	EntityId *string `mandatory:"true" json:"entityId"`

	// Timestamp of when the resource was associated with service catalog.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of the entity that is associated with the service catalog.
	EntityType *string `mandatory:"false" json:"entityType"`
}

func (m ServiceCatalogAssociationSummary) String() string {
	return common.PointerString(m)
}
