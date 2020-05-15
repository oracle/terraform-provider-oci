// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PublicKeySet A set of Public Keys that will be used to verify the JWT signature.
type PublicKeySet interface {
}

type publickeyset struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *publickeyset) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpublickeyset publickeyset
	s := struct {
		Model Unmarshalerpublickeyset
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *publickeyset) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STATIC_KEYS":
		mm := StaticPublicKeySet{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOTE_JWKS":
		mm := RemoteJsonWebKeySet{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m publickeyset) String() string {
	return common.PointerString(m)
}

// PublicKeySetTypeEnum Enum with underlying type: string
type PublicKeySetTypeEnum string

// Set of constants representing the allowable values for PublicKeySetTypeEnum
const (
	PublicKeySetTypeStaticKeys PublicKeySetTypeEnum = "STATIC_KEYS"
	PublicKeySetTypeRemoteJwks PublicKeySetTypeEnum = "REMOTE_JWKS"
)

var mappingPublicKeySetType = map[string]PublicKeySetTypeEnum{
	"STATIC_KEYS": PublicKeySetTypeStaticKeys,
	"REMOTE_JWKS": PublicKeySetTypeRemoteJwks,
}

// GetPublicKeySetTypeEnumValues Enumerates the set of values for PublicKeySetTypeEnum
func GetPublicKeySetTypeEnumValues() []PublicKeySetTypeEnum {
	values := make([]PublicKeySetTypeEnum, 0)
	for _, v := range mappingPublicKeySetType {
		values = append(values, v)
	}
	return values
}
