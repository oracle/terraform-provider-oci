// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ServiceCatalogAssociation The detailed model for service catalog association.
type ServiceCatalogAssociation struct {

	// Identifier of the association.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the service catalog.
	ServiceCatalogId *string `mandatory:"true" json:"serviceCatalogId"`

	// Identifier of the entity being associated with service catalog.
	EntityId *string `mandatory:"true" json:"entityId"`

	// Timestamp of when the resource was associated with service catalog.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of the entity that is associated with the service catalog.
	EntityType *string `mandatory:"false" json:"entityType"`
}

func (m ServiceCatalogAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceCatalogAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
