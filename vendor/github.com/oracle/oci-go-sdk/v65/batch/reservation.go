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

// Reservation Reservation selection used for initial fleet provisioning.
type Reservation interface {
}

type reservation struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *reservation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreservation reservation
	s := struct {
		Model Unmarshalerreservation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *reservation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE_CAPACITY_RESERVATION":
		mm := ComputeCapacityReservation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Reservation: %s.", m.Type)
		return *m, nil
	}
}

func (m reservation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m reservation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReservationTypeEnum Enum with underlying type: string
type ReservationTypeEnum string

// Set of constants representing the allowable values for ReservationTypeEnum
const (
	ReservationTypeComputeCapacityReservation ReservationTypeEnum = "COMPUTE_CAPACITY_RESERVATION"
)

var mappingReservationTypeEnum = map[string]ReservationTypeEnum{
	"COMPUTE_CAPACITY_RESERVATION": ReservationTypeComputeCapacityReservation,
}

var mappingReservationTypeEnumLowerCase = map[string]ReservationTypeEnum{
	"compute_capacity_reservation": ReservationTypeComputeCapacityReservation,
}

// GetReservationTypeEnumValues Enumerates the set of values for ReservationTypeEnum
func GetReservationTypeEnumValues() []ReservationTypeEnum {
	values := make([]ReservationTypeEnum, 0)
	for _, v := range mappingReservationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReservationTypeEnumStringValues Enumerates the set of values in String for ReservationTypeEnum
func GetReservationTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_CAPACITY_RESERVATION",
	}
}

// GetMappingReservationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReservationTypeEnum(val string) (ReservationTypeEnum, bool) {
	enum, ok := mappingReservationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
