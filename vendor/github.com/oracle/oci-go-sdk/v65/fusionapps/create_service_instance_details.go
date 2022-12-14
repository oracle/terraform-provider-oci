// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CreateServiceInstanceDetails The information about new ServiceInstance.
type CreateServiceInstanceDetails interface {

	// The service instance type being provisioned
	GetDisplayName() *string

	// Comparment where the instance is to be created
	GetCompartmentId() *string
}

type createserviceinstancedetails struct {
	JsonData            []byte
	DisplayName         *string `mandatory:"true" json:"displayName"`
	CompartmentId       *string `mandatory:"false" json:"compartmentId"`
	ServiceInstanceType string  `json:"serviceInstanceType"`
}

// UnmarshalJSON unmarshals json
func (m *createserviceinstancedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateserviceinstancedetails createserviceinstancedetails
	s := struct {
		Model Unmarshalercreateserviceinstancedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.ServiceInstanceType = s.Model.ServiceInstanceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createserviceinstancedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ServiceInstanceType {
	case "ANALYTICS_WAREHOUSE":
		mm := CreateOaxServiceInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGRATION_CLOUD":
		mm := CreateOicServiceInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createserviceinstancedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m createserviceinstancedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createserviceinstancedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createserviceinstancedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateServiceInstanceDetailsServiceInstanceTypeEnum Enum with underlying type: string
type CreateServiceInstanceDetailsServiceInstanceTypeEnum string

// Set of constants representing the allowable values for CreateServiceInstanceDetailsServiceInstanceTypeEnum
const (
	CreateServiceInstanceDetailsServiceInstanceTypeIntegrationCloud   CreateServiceInstanceDetailsServiceInstanceTypeEnum = "INTEGRATION_CLOUD"
	CreateServiceInstanceDetailsServiceInstanceTypeAnalyticsWarehouse CreateServiceInstanceDetailsServiceInstanceTypeEnum = "ANALYTICS_WAREHOUSE"
)

var mappingCreateServiceInstanceDetailsServiceInstanceTypeEnum = map[string]CreateServiceInstanceDetailsServiceInstanceTypeEnum{
	"INTEGRATION_CLOUD":   CreateServiceInstanceDetailsServiceInstanceTypeIntegrationCloud,
	"ANALYTICS_WAREHOUSE": CreateServiceInstanceDetailsServiceInstanceTypeAnalyticsWarehouse,
}

var mappingCreateServiceInstanceDetailsServiceInstanceTypeEnumLowerCase = map[string]CreateServiceInstanceDetailsServiceInstanceTypeEnum{
	"integration_cloud":   CreateServiceInstanceDetailsServiceInstanceTypeIntegrationCloud,
	"analytics_warehouse": CreateServiceInstanceDetailsServiceInstanceTypeAnalyticsWarehouse,
}

// GetCreateServiceInstanceDetailsServiceInstanceTypeEnumValues Enumerates the set of values for CreateServiceInstanceDetailsServiceInstanceTypeEnum
func GetCreateServiceInstanceDetailsServiceInstanceTypeEnumValues() []CreateServiceInstanceDetailsServiceInstanceTypeEnum {
	values := make([]CreateServiceInstanceDetailsServiceInstanceTypeEnum, 0)
	for _, v := range mappingCreateServiceInstanceDetailsServiceInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateServiceInstanceDetailsServiceInstanceTypeEnumStringValues Enumerates the set of values in String for CreateServiceInstanceDetailsServiceInstanceTypeEnum
func GetCreateServiceInstanceDetailsServiceInstanceTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_CLOUD",
		"ANALYTICS_WAREHOUSE",
	}
}

// GetMappingCreateServiceInstanceDetailsServiceInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateServiceInstanceDetailsServiceInstanceTypeEnum(val string) (CreateServiceInstanceDetailsServiceInstanceTypeEnum, bool) {
	enum, ok := mappingCreateServiceInstanceDetailsServiceInstanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
