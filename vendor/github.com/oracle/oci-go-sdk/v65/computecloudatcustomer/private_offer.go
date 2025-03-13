// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateOffer The private offer for the application
type PrivateOffer struct {

	// Defines whether the listing supports private offers.
	IsPrivateOfferEnabled *bool `mandatory:"true" json:"isPrivateOfferEnabled"`

	// The unit of measure for the private offer.
	UnitOfMeasure PrivateOfferUnitOfMeasureEnum `mandatory:"true" json:"unitOfMeasure"`

	// The contact email for the publisher of the private offer.
	ContactEmail *string `mandatory:"true" json:"contactEmail"`

	// Whether a private offer subscription is active for this customer.
	IsPrivateOfferActive *bool `mandatory:"true" json:"isPrivateOfferActive"`

	// The contact url for the publisher of the private offer.
	ContactUrl *string `mandatory:"false" json:"contactUrl"`
}

func (m PrivateOffer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateOffer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateOfferUnitOfMeasureEnum(string(m.UnitOfMeasure)); !ok && m.UnitOfMeasure != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitOfMeasure: %s. Supported values are: %s.", m.UnitOfMeasure, strings.Join(GetPrivateOfferUnitOfMeasureEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateOfferUnitOfMeasureEnum Enum with underlying type: string
type PrivateOfferUnitOfMeasureEnum string

// Set of constants representing the allowable values for PrivateOfferUnitOfMeasureEnum
const (
	PrivateOfferUnitOfMeasureOcpuPerHour     PrivateOfferUnitOfMeasureEnum = "OCPU_PER_HOUR"
	PrivateOfferUnitOfMeasureInstancePerHour PrivateOfferUnitOfMeasureEnum = "INSTANCE_PER_HOUR"
	PrivateOfferUnitOfMeasureCredits         PrivateOfferUnitOfMeasureEnum = "CREDITS"
	PrivateOfferUnitOfMeasureInstances       PrivateOfferUnitOfMeasureEnum = "INSTANCES"
	PrivateOfferUnitOfMeasureNodes           PrivateOfferUnitOfMeasureEnum = "NODES"
)

var mappingPrivateOfferUnitOfMeasureEnum = map[string]PrivateOfferUnitOfMeasureEnum{
	"OCPU_PER_HOUR":     PrivateOfferUnitOfMeasureOcpuPerHour,
	"INSTANCE_PER_HOUR": PrivateOfferUnitOfMeasureInstancePerHour,
	"CREDITS":           PrivateOfferUnitOfMeasureCredits,
	"INSTANCES":         PrivateOfferUnitOfMeasureInstances,
	"NODES":             PrivateOfferUnitOfMeasureNodes,
}

var mappingPrivateOfferUnitOfMeasureEnumLowerCase = map[string]PrivateOfferUnitOfMeasureEnum{
	"ocpu_per_hour":     PrivateOfferUnitOfMeasureOcpuPerHour,
	"instance_per_hour": PrivateOfferUnitOfMeasureInstancePerHour,
	"credits":           PrivateOfferUnitOfMeasureCredits,
	"instances":         PrivateOfferUnitOfMeasureInstances,
	"nodes":             PrivateOfferUnitOfMeasureNodes,
}

// GetPrivateOfferUnitOfMeasureEnumValues Enumerates the set of values for PrivateOfferUnitOfMeasureEnum
func GetPrivateOfferUnitOfMeasureEnumValues() []PrivateOfferUnitOfMeasureEnum {
	values := make([]PrivateOfferUnitOfMeasureEnum, 0)
	for _, v := range mappingPrivateOfferUnitOfMeasureEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateOfferUnitOfMeasureEnumStringValues Enumerates the set of values in String for PrivateOfferUnitOfMeasureEnum
func GetPrivateOfferUnitOfMeasureEnumStringValues() []string {
	return []string{
		"OCPU_PER_HOUR",
		"INSTANCE_PER_HOUR",
		"CREDITS",
		"INSTANCES",
		"NODES",
	}
}

// GetMappingPrivateOfferUnitOfMeasureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateOfferUnitOfMeasureEnum(val string) (PrivateOfferUnitOfMeasureEnum, bool) {
	enum, ok := mappingPrivateOfferUnitOfMeasureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
