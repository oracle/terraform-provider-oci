// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClientTenantStripe This contains details to fetch client tenant stripe for the API caller.
// Currently, only function is allowed, but in future, we may add new ways to fetch tenant stripe (say, we add support for a webserver instead of OCI function).
type ClientTenantStripe interface {
}

type clienttenantstripe struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *clienttenantstripe) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerclienttenantstripe clienttenantstripe
	s := struct {
		Model Unmarshalerclienttenantstripe
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *clienttenantstripe) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STRIPE_MANAGER_FUNCTION":
		mm := StripeManagerFunctionClientTenantStripe{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ClientTenantStripe: %s.", m.Type)
		return *m, nil
	}
}

func (m clienttenantstripe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m clienttenantstripe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClientTenantStripeTypeEnum Enum with underlying type: string
type ClientTenantStripeTypeEnum string

// Set of constants representing the allowable values for ClientTenantStripeTypeEnum
const (
	ClientTenantStripeTypeStripeManagerFunction ClientTenantStripeTypeEnum = "STRIPE_MANAGER_FUNCTION"
)

var mappingClientTenantStripeTypeEnum = map[string]ClientTenantStripeTypeEnum{
	"STRIPE_MANAGER_FUNCTION": ClientTenantStripeTypeStripeManagerFunction,
}

var mappingClientTenantStripeTypeEnumLowerCase = map[string]ClientTenantStripeTypeEnum{
	"stripe_manager_function": ClientTenantStripeTypeStripeManagerFunction,
}

// GetClientTenantStripeTypeEnumValues Enumerates the set of values for ClientTenantStripeTypeEnum
func GetClientTenantStripeTypeEnumValues() []ClientTenantStripeTypeEnum {
	values := make([]ClientTenantStripeTypeEnum, 0)
	for _, v := range mappingClientTenantStripeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClientTenantStripeTypeEnumStringValues Enumerates the set of values in String for ClientTenantStripeTypeEnum
func GetClientTenantStripeTypeEnumStringValues() []string {
	return []string{
		"STRIPE_MANAGER_FUNCTION",
	}
}

// GetMappingClientTenantStripeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientTenantStripeTypeEnum(val string) (ClientTenantStripeTypeEnum, bool) {
	enum, ok := mappingClientTenantStripeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
