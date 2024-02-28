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

// OccCapacityRequestBaseDetails The details of the create capacity request. This model serves as a base for different namespaces.
type OccCapacityRequestBaseDetails interface {

	// The type of the workload (Generic/ROW).
	GetWorkloadType() OccAvailabilitySummaryWorkloadTypeEnum

	// The incremental quantity of resources supplied as the provisioning is underway.
	GetExpectedHandoverQuantity() *int64

	// The date on which the latest increment to supplied quantity of resources was delivered.
	GetDateExpectedHandover() *common.SDKTime

	// The actual handed over quantity of resources at the time of request resolution.
	GetActualHandoverQuantity() *int64

	// The date on which the actual handover quantity of resources is delivered.
	GetDateActualHandover() *common.SDKTime
}

type occcapacityrequestbasedetails struct {
	JsonData                 []byte
	ExpectedHandoverQuantity *int64                                 `mandatory:"false" json:"expectedHandoverQuantity"`
	DateExpectedHandover     *common.SDKTime                        `mandatory:"false" json:"dateExpectedHandover"`
	ActualHandoverQuantity   *int64                                 `mandatory:"false" json:"actualHandoverQuantity"`
	DateActualHandover       *common.SDKTime                        `mandatory:"false" json:"dateActualHandover"`
	WorkloadType             OccAvailabilitySummaryWorkloadTypeEnum `mandatory:"true" json:"workloadType"`
	ResourceType             string                                 `json:"resourceType"`
}

// UnmarshalJSON unmarshals json
func (m *occcapacityrequestbasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerocccapacityrequestbasedetails occcapacityrequestbasedetails
	s := struct {
		Model Unmarshalerocccapacityrequestbasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkloadType = s.Model.WorkloadType
	m.ExpectedHandoverQuantity = s.Model.ExpectedHandoverQuantity
	m.DateExpectedHandover = s.Model.DateExpectedHandover
	m.ActualHandoverQuantity = s.Model.ActualHandoverQuantity
	m.DateActualHandover = s.Model.DateActualHandover
	m.ResourceType = s.Model.ResourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *occcapacityrequestbasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResourceType {
	case "SERVER_HW":
		mm := OccCapacityRequestComputeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OccCapacityRequestBaseDetails: %s.", m.ResourceType)
		return *m, nil
	}
}

// GetExpectedHandoverQuantity returns ExpectedHandoverQuantity
func (m occcapacityrequestbasedetails) GetExpectedHandoverQuantity() *int64 {
	return m.ExpectedHandoverQuantity
}

// GetDateExpectedHandover returns DateExpectedHandover
func (m occcapacityrequestbasedetails) GetDateExpectedHandover() *common.SDKTime {
	return m.DateExpectedHandover
}

// GetActualHandoverQuantity returns ActualHandoverQuantity
func (m occcapacityrequestbasedetails) GetActualHandoverQuantity() *int64 {
	return m.ActualHandoverQuantity
}

// GetDateActualHandover returns DateActualHandover
func (m occcapacityrequestbasedetails) GetDateActualHandover() *common.SDKTime {
	return m.DateActualHandover
}

// GetWorkloadType returns WorkloadType
func (m occcapacityrequestbasedetails) GetWorkloadType() OccAvailabilitySummaryWorkloadTypeEnum {
	return m.WorkloadType
}

func (m occcapacityrequestbasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m occcapacityrequestbasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccAvailabilitySummaryWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetOccAvailabilitySummaryWorkloadTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccCapacityRequestBaseDetailsResourceTypeEnum Enum with underlying type: string
type OccCapacityRequestBaseDetailsResourceTypeEnum string

// Set of constants representing the allowable values for OccCapacityRequestBaseDetailsResourceTypeEnum
const (
	OccCapacityRequestBaseDetailsResourceTypeServerHw OccCapacityRequestBaseDetailsResourceTypeEnum = "SERVER_HW"
)

var mappingOccCapacityRequestBaseDetailsResourceTypeEnum = map[string]OccCapacityRequestBaseDetailsResourceTypeEnum{
	"SERVER_HW": OccCapacityRequestBaseDetailsResourceTypeServerHw,
}

var mappingOccCapacityRequestBaseDetailsResourceTypeEnumLowerCase = map[string]OccCapacityRequestBaseDetailsResourceTypeEnum{
	"server_hw": OccCapacityRequestBaseDetailsResourceTypeServerHw,
}

// GetOccCapacityRequestBaseDetailsResourceTypeEnumValues Enumerates the set of values for OccCapacityRequestBaseDetailsResourceTypeEnum
func GetOccCapacityRequestBaseDetailsResourceTypeEnumValues() []OccCapacityRequestBaseDetailsResourceTypeEnum {
	values := make([]OccCapacityRequestBaseDetailsResourceTypeEnum, 0)
	for _, v := range mappingOccCapacityRequestBaseDetailsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCapacityRequestBaseDetailsResourceTypeEnumStringValues Enumerates the set of values in String for OccCapacityRequestBaseDetailsResourceTypeEnum
func GetOccCapacityRequestBaseDetailsResourceTypeEnumStringValues() []string {
	return []string{
		"SERVER_HW",
	}
}

// GetMappingOccCapacityRequestBaseDetailsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCapacityRequestBaseDetailsResourceTypeEnum(val string) (OccCapacityRequestBaseDetailsResourceTypeEnum, bool) {
	enum, ok := mappingOccCapacityRequestBaseDetailsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
