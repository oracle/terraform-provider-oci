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

// OccCapacityRequestBaseDetails The details of the create capacity request. This model serves as a base for different namespaces.
type OccCapacityRequestBaseDetails struct {

	// The type of the resource against which the user wants to place a capacity request.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The type of the workload (Generic/ROW).
	WorkloadType *string `mandatory:"true" json:"workloadType"`

	// The name of the COMPUTE server shape for which the request is made. Do not use CAPACITY_CONSTRAINT as the resource name.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The number of compute server's with name <resourceName> required by the user.
	DemandQuantity *int64 `mandatory:"true" json:"demandQuantity"`

	// The WorkloadType from where capacity request are to be transferred.
	SourceWorkloadType *string `mandatory:"false" json:"sourceWorkloadType"`

	// The incremental quantity of resources supplied as the provisioning is underway.
	ExpectedHandoverQuantity *int64 `mandatory:"false" json:"expectedHandoverQuantity"`

	// The date on which the latest increment to supplied quantity of resources was delivered.
	DateExpectedHandover *common.SDKTime `mandatory:"false" json:"dateExpectedHandover"`

	// The actual handed over quantity of resources at the time of request resolution.
	ActualHandoverQuantity *int64 `mandatory:"false" json:"actualHandoverQuantity"`

	// The date on which the actual handover quantity of resources is delivered.
	DateActualHandover *common.SDKTime `mandatory:"false" json:"dateActualHandover"`
}

func (m OccCapacityRequestBaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCapacityRequestBaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
