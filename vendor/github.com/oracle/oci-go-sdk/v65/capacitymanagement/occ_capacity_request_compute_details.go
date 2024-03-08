// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccCapacityRequestComputeDetails The details about the compute servers required for creating a capacity request.
type OccCapacityRequestComputeDetails struct {

	// The name of the COMPUTE server shape for which the request is made. Do not use CAPACITY_CONSTRAINT as the resource name.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The number of compute server's with name <resourceName> required by the user.
	DemandQuantity *int64 `mandatory:"true" json:"demandQuantity"`

	// The incremental quantity of resources supplied as the provisioning is underway.
	ExpectedHandoverQuantity *int64 `mandatory:"false" json:"expectedHandoverQuantity"`

	// The date on which the latest increment to supplied quantity of resources was delivered.
	DateExpectedHandover *common.SDKTime `mandatory:"false" json:"dateExpectedHandover"`

	// The actual handed over quantity of resources at the time of request resolution.
	ActualHandoverQuantity *int64 `mandatory:"false" json:"actualHandoverQuantity"`

	// The date on which the actual handover quantity of resources is delivered.
	DateActualHandover *common.SDKTime `mandatory:"false" json:"dateActualHandover"`

	// The type of the workload (Generic/ROW).
	WorkloadType OccAvailabilitySummaryWorkloadTypeEnum `mandatory:"true" json:"workloadType"`
}

// GetWorkloadType returns WorkloadType
func (m OccCapacityRequestComputeDetails) GetWorkloadType() OccAvailabilitySummaryWorkloadTypeEnum {
	return m.WorkloadType
}

// GetExpectedHandoverQuantity returns ExpectedHandoverQuantity
func (m OccCapacityRequestComputeDetails) GetExpectedHandoverQuantity() *int64 {
	return m.ExpectedHandoverQuantity
}

// GetDateExpectedHandover returns DateExpectedHandover
func (m OccCapacityRequestComputeDetails) GetDateExpectedHandover() *common.SDKTime {
	return m.DateExpectedHandover
}

// GetActualHandoverQuantity returns ActualHandoverQuantity
func (m OccCapacityRequestComputeDetails) GetActualHandoverQuantity() *int64 {
	return m.ActualHandoverQuantity
}

// GetDateActualHandover returns DateActualHandover
func (m OccCapacityRequestComputeDetails) GetDateActualHandover() *common.SDKTime {
	return m.DateActualHandover
}

func (m OccCapacityRequestComputeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCapacityRequestComputeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOccAvailabilitySummaryWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetOccAvailabilitySummaryWorkloadTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OccCapacityRequestComputeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOccCapacityRequestComputeDetails OccCapacityRequestComputeDetails
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeOccCapacityRequestComputeDetails
	}{
		"SERVER_HW",
		(MarshalTypeOccCapacityRequestComputeDetails)(m),
	}

	return json.Marshal(&s)
}
