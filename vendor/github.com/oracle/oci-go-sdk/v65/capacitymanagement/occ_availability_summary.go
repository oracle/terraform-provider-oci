// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccAvailabilitySummary The details about the available capacity and constraints for different resource types present in the availability catalog.
type OccAvailabilitySummary struct {

	// The OCID of the availability catalog.
	CatalogId *string `mandatory:"true" json:"catalogId"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The date by which the customer must place the order to have their capacity requirements met by the customer handover date.
	DateFinalCustomerOrder *common.SDKTime `mandatory:"true" json:"dateFinalCustomerOrder"`

	// The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	DateExpectedCapacityHandover *common.SDKTime `mandatory:"true" json:"dateExpectedCapacityHandover"`

	// The different types of resources against which customers can place capacity requests.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The type of workload (Generic/ROW).
	WorkloadType *string `mandatory:"true" json:"workloadType"`

	// The name of the resource that the customer can request.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The quantity of resource currently available that the customer can request.
	AvailableQuantity *int64 `mandatory:"true" json:"availableQuantity"`

	// The total quantity of resource that the customer can request.
	TotalAvailableQuantity *int64 `mandatory:"true" json:"totalAvailableQuantity"`

	// The quantity of resource currently demanded by the customer.
	DemandedQuantity *int64 `mandatory:"true" json:"demandedQuantity"`

	// The unit in which the resource available is measured.
	Unit *string `mandatory:"true" json:"unit"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccAvailabilitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccAvailabilitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
