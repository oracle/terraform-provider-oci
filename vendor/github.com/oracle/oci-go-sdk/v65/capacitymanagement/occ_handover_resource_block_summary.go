// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccHandoverResourceBlockSummary Details about the association of capacity requests with the corresponding resources handed over by oracle.
type OccHandoverResourceBlockSummary struct {

	// The OCID of the resource block.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment where the resource block's are placed.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the OCI service in consideration.
	// For example Compute, Exadata and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The OCID of the customer group for which the resources were provisioned.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The date on which the resource was handed over to the customer.
	HandoverDate *common.SDKTime `mandatory:"true" json:"handoverDate"`

	// The name of the resource handed over by oracle.
	// For instance for compute namespace this will be the name of the bare metal hardware resource.
	HandoverResourceName *string `mandatory:"true" json:"handoverResourceName"`

	// The total quantity of the resource that was made available to the customer by Oracle.
	TotalHandoverQuantity *int64 `mandatory:"true" json:"totalHandoverQuantity"`

	// A list containing details about the capacity requests against which the resources were provisioned by oracle.
	AssociatedCapacityRequests []AssociatedCapacityRequestDetails `mandatory:"true" json:"associatedCapacityRequests"`

	PlacementDetails *PlacementDetails `mandatory:"true" json:"placementDetails"`
}

func (m OccHandoverResourceBlockSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccHandoverResourceBlockSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
