// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AccessPolicyTarget Target of the access policy. This can either be the source or the destination of the traffic.
type AccessPolicyTarget interface {
}

type accesspolicytarget struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *accesspolicytarget) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraccesspolicytarget accesspolicytarget
	s := struct {
		Model Unmarshaleraccesspolicytarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *accesspolicytarget) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VIRTUAL_SERVICE":
		mm := VirtualServiceAccessPolicyTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALL_VIRTUAL_SERVICES":
		mm := AllVirtualServicesAccessPolicyTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTERNAL_SERVICE":
		mm := ExternalServiceAccessPolicyTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INGRESS_GATEWAY":
		mm := IngressGatewayAccessPolicyTarget{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m accesspolicytarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m accesspolicytarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccessPolicyTargetTypeEnum Enum with underlying type: string
type AccessPolicyTargetTypeEnum string

// Set of constants representing the allowable values for AccessPolicyTargetTypeEnum
const (
	AccessPolicyTargetTypeAllVirtualServices AccessPolicyTargetTypeEnum = "ALL_VIRTUAL_SERVICES"
	AccessPolicyTargetTypeVirtualService     AccessPolicyTargetTypeEnum = "VIRTUAL_SERVICE"
	AccessPolicyTargetTypeExternalService    AccessPolicyTargetTypeEnum = "EXTERNAL_SERVICE"
	AccessPolicyTargetTypeIngressGateway     AccessPolicyTargetTypeEnum = "INGRESS_GATEWAY"
)

var mappingAccessPolicyTargetTypeEnum = map[string]AccessPolicyTargetTypeEnum{
	"ALL_VIRTUAL_SERVICES": AccessPolicyTargetTypeAllVirtualServices,
	"VIRTUAL_SERVICE":      AccessPolicyTargetTypeVirtualService,
	"EXTERNAL_SERVICE":     AccessPolicyTargetTypeExternalService,
	"INGRESS_GATEWAY":      AccessPolicyTargetTypeIngressGateway,
}

var mappingAccessPolicyTargetTypeEnumLowerCase = map[string]AccessPolicyTargetTypeEnum{
	"all_virtual_services": AccessPolicyTargetTypeAllVirtualServices,
	"virtual_service":      AccessPolicyTargetTypeVirtualService,
	"external_service":     AccessPolicyTargetTypeExternalService,
	"ingress_gateway":      AccessPolicyTargetTypeIngressGateway,
}

// GetAccessPolicyTargetTypeEnumValues Enumerates the set of values for AccessPolicyTargetTypeEnum
func GetAccessPolicyTargetTypeEnumValues() []AccessPolicyTargetTypeEnum {
	values := make([]AccessPolicyTargetTypeEnum, 0)
	for _, v := range mappingAccessPolicyTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessPolicyTargetTypeEnumStringValues Enumerates the set of values in String for AccessPolicyTargetTypeEnum
func GetAccessPolicyTargetTypeEnumStringValues() []string {
	return []string{
		"ALL_VIRTUAL_SERVICES",
		"VIRTUAL_SERVICE",
		"EXTERNAL_SERVICE",
		"INGRESS_GATEWAY",
	}
}

// GetMappingAccessPolicyTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessPolicyTargetTypeEnum(val string) (AccessPolicyTargetTypeEnum, bool) {
	enum, ok := mappingAccessPolicyTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
