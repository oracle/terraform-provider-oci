// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Use the Service Catalog API to manage solutions in Oracle Cloud Infrastructure Service Catalog.
// For more information, see Overview of Service Catalog (https://docs.oracle.com/iaas/Content/service-catalog/overview_of_service_catalog.htm).
//

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ServiceCatalogAssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceCatalogAssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
