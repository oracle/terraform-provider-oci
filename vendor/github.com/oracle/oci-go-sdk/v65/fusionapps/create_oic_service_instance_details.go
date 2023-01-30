// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOicServiceInstanceDetails The information about new Integration Cloud instance being provisioned.
type CreateOicServiceInstanceDetails struct {

	// The service instance type being provisioned
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Comparment where the instance is to be created
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Number of 5K message packs per hour
	MessagePacks *int `mandatory:"false" json:"messagePacks"`

	// The   Oracle Integration edition
	Edition CreateOicServiceInstanceDetailsEditionEnum `mandatory:"false" json:"edition,omitempty"`
}

//GetDisplayName returns DisplayName
func (m CreateOicServiceInstanceDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateOicServiceInstanceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m CreateOicServiceInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOicServiceInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateOicServiceInstanceDetailsEditionEnum(string(m.Edition)); !ok && m.Edition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Edition: %s. Supported values are: %s.", m.Edition, strings.Join(GetCreateOicServiceInstanceDetailsEditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOicServiceInstanceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOicServiceInstanceDetails CreateOicServiceInstanceDetails
	s := struct {
		DiscriminatorParam string `json:"serviceInstanceType"`
		MarshalTypeCreateOicServiceInstanceDetails
	}{
		"INTEGRATION_CLOUD",
		(MarshalTypeCreateOicServiceInstanceDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateOicServiceInstanceDetailsEditionEnum Enum with underlying type: string
type CreateOicServiceInstanceDetailsEditionEnum string

// Set of constants representing the allowable values for CreateOicServiceInstanceDetailsEditionEnum
const (
	CreateOicServiceInstanceDetailsEditionStandard   CreateOicServiceInstanceDetailsEditionEnum = "STANDARD"
	CreateOicServiceInstanceDetailsEditionEnterprise CreateOicServiceInstanceDetailsEditionEnum = "ENTERPRISE"
)

var mappingCreateOicServiceInstanceDetailsEditionEnum = map[string]CreateOicServiceInstanceDetailsEditionEnum{
	"STANDARD":   CreateOicServiceInstanceDetailsEditionStandard,
	"ENTERPRISE": CreateOicServiceInstanceDetailsEditionEnterprise,
}

var mappingCreateOicServiceInstanceDetailsEditionEnumLowerCase = map[string]CreateOicServiceInstanceDetailsEditionEnum{
	"standard":   CreateOicServiceInstanceDetailsEditionStandard,
	"enterprise": CreateOicServiceInstanceDetailsEditionEnterprise,
}

// GetCreateOicServiceInstanceDetailsEditionEnumValues Enumerates the set of values for CreateOicServiceInstanceDetailsEditionEnum
func GetCreateOicServiceInstanceDetailsEditionEnumValues() []CreateOicServiceInstanceDetailsEditionEnum {
	values := make([]CreateOicServiceInstanceDetailsEditionEnum, 0)
	for _, v := range mappingCreateOicServiceInstanceDetailsEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOicServiceInstanceDetailsEditionEnumStringValues Enumerates the set of values in String for CreateOicServiceInstanceDetailsEditionEnum
func GetCreateOicServiceInstanceDetailsEditionEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE",
	}
}

// GetMappingCreateOicServiceInstanceDetailsEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOicServiceInstanceDetailsEditionEnum(val string) (CreateOicServiceInstanceDetailsEditionEnum, bool) {
	enum, ok := mappingCreateOicServiceInstanceDetailsEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
