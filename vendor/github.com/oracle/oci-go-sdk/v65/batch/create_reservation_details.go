// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateReservationDetails Reservation selection for initial provisioning of a fleet.
type CreateReservationDetails interface {
}

type createreservationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createreservationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatereservationdetails createreservationdetails
	s := struct {
		Model Unmarshalercreatereservationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createreservationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE_CAPACITY_RESERVATION":
		mm := CreateComputeCapacityReservationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateReservationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createreservationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createreservationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateReservationDetailsTypeEnum Enum with underlying type: string
type CreateReservationDetailsTypeEnum string

// Set of constants representing the allowable values for CreateReservationDetailsTypeEnum
const (
	CreateReservationDetailsTypeComputeCapacityReservation CreateReservationDetailsTypeEnum = "COMPUTE_CAPACITY_RESERVATION"
)

var mappingCreateReservationDetailsTypeEnum = map[string]CreateReservationDetailsTypeEnum{
	"COMPUTE_CAPACITY_RESERVATION": CreateReservationDetailsTypeComputeCapacityReservation,
}

var mappingCreateReservationDetailsTypeEnumLowerCase = map[string]CreateReservationDetailsTypeEnum{
	"compute_capacity_reservation": CreateReservationDetailsTypeComputeCapacityReservation,
}

// GetCreateReservationDetailsTypeEnumValues Enumerates the set of values for CreateReservationDetailsTypeEnum
func GetCreateReservationDetailsTypeEnumValues() []CreateReservationDetailsTypeEnum {
	values := make([]CreateReservationDetailsTypeEnum, 0)
	for _, v := range mappingCreateReservationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateReservationDetailsTypeEnumStringValues Enumerates the set of values in String for CreateReservationDetailsTypeEnum
func GetCreateReservationDetailsTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_CAPACITY_RESERVATION",
	}
}

// GetMappingCreateReservationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateReservationDetailsTypeEnum(val string) (CreateReservationDetailsTypeEnum, bool) {
	enum, ok := mappingCreateReservationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
