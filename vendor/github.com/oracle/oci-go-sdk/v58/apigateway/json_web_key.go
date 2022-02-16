// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// JsonWebKey A JSON Web Key that represents the public key used for verifying the JWT signature.
type JsonWebKey struct {

	// A unique key ID. This key will be used to verify the signature of a
	// JWT with matching "kid".
	Kid *string `mandatory:"true" json:"kid"`

	// The algorithm intended for use with this key.
	Alg *string `mandatory:"true" json:"alg"`

	// The base64 url encoded modulus of the RSA public key represented
	// by this key.
	N *string `mandatory:"true" json:"n"`

	// The base64 url encoded exponent of the RSA public key represented
	// by this key.
	E *string `mandatory:"true" json:"e"`

	// The key type.
	Kty JsonWebKeyKtyEnum `mandatory:"true" json:"kty"`

	// The intended use of the public key.
	Use JsonWebKeyUseEnum `mandatory:"false" json:"use,omitempty"`

	// The operations for which this key is to be used.
	KeyOps []JsonWebKeyKeyOpsEnum `mandatory:"false" json:"key_ops,omitempty"`
}

//GetKid returns Kid
func (m JsonWebKey) GetKid() *string {
	return m.Kid
}

func (m JsonWebKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JsonWebKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJsonWebKeyKtyEnum(string(m.Kty)); !ok && m.Kty != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kty: %s. Supported values are: %s.", m.Kty, strings.Join(GetJsonWebKeyKtyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJsonWebKeyUseEnum(string(m.Use)); !ok && m.Use != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Use: %s. Supported values are: %s.", m.Use, strings.Join(GetJsonWebKeyUseEnumStringValues(), ",")))
	}
	for _, val := range m.KeyOps {
		if _, ok := GetMappingJsonWebKeyKeyOpsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyOps: %s. Supported values are: %s.", val, strings.Join(GetJsonWebKeyKeyOpsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JsonWebKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJsonWebKey JsonWebKey
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypeJsonWebKey
	}{
		"JSON_WEB_KEY",
		(MarshalTypeJsonWebKey)(m),
	}

	return json.Marshal(&s)
}

// JsonWebKeyKtyEnum Enum with underlying type: string
type JsonWebKeyKtyEnum string

// Set of constants representing the allowable values for JsonWebKeyKtyEnum
const (
	JsonWebKeyKtyRsa JsonWebKeyKtyEnum = "RSA"
)

var mappingJsonWebKeyKtyEnum = map[string]JsonWebKeyKtyEnum{
	"RSA": JsonWebKeyKtyRsa,
}

// GetJsonWebKeyKtyEnumValues Enumerates the set of values for JsonWebKeyKtyEnum
func GetJsonWebKeyKtyEnumValues() []JsonWebKeyKtyEnum {
	values := make([]JsonWebKeyKtyEnum, 0)
	for _, v := range mappingJsonWebKeyKtyEnum {
		values = append(values, v)
	}
	return values
}

// GetJsonWebKeyKtyEnumStringValues Enumerates the set of values in String for JsonWebKeyKtyEnum
func GetJsonWebKeyKtyEnumStringValues() []string {
	return []string{
		"RSA",
	}
}

// GetMappingJsonWebKeyKtyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJsonWebKeyKtyEnum(val string) (JsonWebKeyKtyEnum, bool) {
	mappingJsonWebKeyKtyEnumIgnoreCase := make(map[string]JsonWebKeyKtyEnum)
	for k, v := range mappingJsonWebKeyKtyEnum {
		mappingJsonWebKeyKtyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJsonWebKeyKtyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// JsonWebKeyUseEnum Enum with underlying type: string
type JsonWebKeyUseEnum string

// Set of constants representing the allowable values for JsonWebKeyUseEnum
const (
	JsonWebKeyUseSig JsonWebKeyUseEnum = "sig"
)

var mappingJsonWebKeyUseEnum = map[string]JsonWebKeyUseEnum{
	"sig": JsonWebKeyUseSig,
}

// GetJsonWebKeyUseEnumValues Enumerates the set of values for JsonWebKeyUseEnum
func GetJsonWebKeyUseEnumValues() []JsonWebKeyUseEnum {
	values := make([]JsonWebKeyUseEnum, 0)
	for _, v := range mappingJsonWebKeyUseEnum {
		values = append(values, v)
	}
	return values
}

// GetJsonWebKeyUseEnumStringValues Enumerates the set of values in String for JsonWebKeyUseEnum
func GetJsonWebKeyUseEnumStringValues() []string {
	return []string{
		"sig",
	}
}

// GetMappingJsonWebKeyUseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJsonWebKeyUseEnum(val string) (JsonWebKeyUseEnum, bool) {
	mappingJsonWebKeyUseEnumIgnoreCase := make(map[string]JsonWebKeyUseEnum)
	for k, v := range mappingJsonWebKeyUseEnum {
		mappingJsonWebKeyUseEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJsonWebKeyUseEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// JsonWebKeyKeyOpsEnum Enum with underlying type: string
type JsonWebKeyKeyOpsEnum string

// Set of constants representing the allowable values for JsonWebKeyKeyOpsEnum
const (
	JsonWebKeyKeyOpsVerify JsonWebKeyKeyOpsEnum = "verify"
)

var mappingJsonWebKeyKeyOpsEnum = map[string]JsonWebKeyKeyOpsEnum{
	"verify": JsonWebKeyKeyOpsVerify,
}

// GetJsonWebKeyKeyOpsEnumValues Enumerates the set of values for JsonWebKeyKeyOpsEnum
func GetJsonWebKeyKeyOpsEnumValues() []JsonWebKeyKeyOpsEnum {
	values := make([]JsonWebKeyKeyOpsEnum, 0)
	for _, v := range mappingJsonWebKeyKeyOpsEnum {
		values = append(values, v)
	}
	return values
}

// GetJsonWebKeyKeyOpsEnumStringValues Enumerates the set of values in String for JsonWebKeyKeyOpsEnum
func GetJsonWebKeyKeyOpsEnumStringValues() []string {
	return []string{
		"verify",
	}
}

// GetMappingJsonWebKeyKeyOpsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJsonWebKeyKeyOpsEnum(val string) (JsonWebKeyKeyOpsEnum, bool) {
	mappingJsonWebKeyKeyOpsEnumIgnoreCase := make(map[string]JsonWebKeyKeyOpsEnum)
	for k, v := range mappingJsonWebKeyKeyOpsEnum {
		mappingJsonWebKeyKeyOpsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJsonWebKeyKeyOpsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
