// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessPolicyTargetDetails Target of the access policy. This can either be the source or the destination of the traffic.
type AccessPolicyTargetDetails interface {
}

type accesspolicytargetdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *accesspolicytargetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraccesspolicytargetdetails accesspolicytargetdetails
	s := struct {
		Model Unmarshaleraccesspolicytargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *accesspolicytargetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "EXTERNAL_SERVICE":
		mm := ExternalServiceAccessPolicyTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIRTUAL_SERVICE":
		mm := VirtualServiceAccessPolicyTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALL_VIRTUAL_SERVICES":
		mm := AllVirtualServicesAccessPolicyTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INGRESS_GATEWAY":
		mm := IngressGatewayAccessPolicyTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m accesspolicytargetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m accesspolicytargetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccessPolicyTargetDetailsTypeEnum Enum with underlying type: string
type AccessPolicyTargetDetailsTypeEnum string

// Set of constants representing the allowable values for AccessPolicyTargetDetailsTypeEnum
const (
	AccessPolicyTargetDetailsTypeAllVirtualServices AccessPolicyTargetDetailsTypeEnum = "ALL_VIRTUAL_SERVICES"
	AccessPolicyTargetDetailsTypeVirtualService     AccessPolicyTargetDetailsTypeEnum = "VIRTUAL_SERVICE"
	AccessPolicyTargetDetailsTypeExternalService    AccessPolicyTargetDetailsTypeEnum = "EXTERNAL_SERVICE"
	AccessPolicyTargetDetailsTypeIngressGateway     AccessPolicyTargetDetailsTypeEnum = "INGRESS_GATEWAY"
)

var mappingAccessPolicyTargetDetailsTypeEnum = map[string]AccessPolicyTargetDetailsTypeEnum{
	"ALL_VIRTUAL_SERVICES": AccessPolicyTargetDetailsTypeAllVirtualServices,
	"VIRTUAL_SERVICE":      AccessPolicyTargetDetailsTypeVirtualService,
	"EXTERNAL_SERVICE":     AccessPolicyTargetDetailsTypeExternalService,
	"INGRESS_GATEWAY":      AccessPolicyTargetDetailsTypeIngressGateway,
}

var mappingAccessPolicyTargetDetailsTypeEnumLowerCase = map[string]AccessPolicyTargetDetailsTypeEnum{
	"all_virtual_services": AccessPolicyTargetDetailsTypeAllVirtualServices,
	"virtual_service":      AccessPolicyTargetDetailsTypeVirtualService,
	"external_service":     AccessPolicyTargetDetailsTypeExternalService,
	"ingress_gateway":      AccessPolicyTargetDetailsTypeIngressGateway,
}

// GetAccessPolicyTargetDetailsTypeEnumValues Enumerates the set of values for AccessPolicyTargetDetailsTypeEnum
func GetAccessPolicyTargetDetailsTypeEnumValues() []AccessPolicyTargetDetailsTypeEnum {
	values := make([]AccessPolicyTargetDetailsTypeEnum, 0)
	for _, v := range mappingAccessPolicyTargetDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessPolicyTargetDetailsTypeEnumStringValues Enumerates the set of values in String for AccessPolicyTargetDetailsTypeEnum
func GetAccessPolicyTargetDetailsTypeEnumStringValues() []string {
	return []string{
		"ALL_VIRTUAL_SERVICES",
		"VIRTUAL_SERVICE",
		"EXTERNAL_SERVICE",
		"INGRESS_GATEWAY",
	}
}

// GetMappingAccessPolicyTargetDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessPolicyTargetDetailsTypeEnum(val string) (AccessPolicyTargetDetailsTypeEnum, bool) {
	enum, ok := mappingAccessPolicyTargetDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
