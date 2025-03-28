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

// StaticPublicKey A static public key which is used to verify the JWT signature.
type StaticPublicKey interface {

	// A unique key ID. This key will be used to verify the signature of a
	// JWT with matching "kid".
	GetKid() *string
}

type staticpublickey struct {
	JsonData []byte
	Kid      *string `mandatory:"true" json:"kid"`
	Format   string  `json:"format"`
}

// UnmarshalJSON unmarshals json
func (m *staticpublickey) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstaticpublickey staticpublickey
	s := struct {
		Model Unmarshalerstaticpublickey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kid = s.Model.Kid
	m.Format = s.Model.Format

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *staticpublickey) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Format {
	case "JSON_WEB_KEY":
		mm := JsonWebKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PEM":
		mm := PemEncodedPublicKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for StaticPublicKey: %s.", m.Format)
		return *m, nil
	}
}

// GetKid returns Kid
func (m staticpublickey) GetKid() *string {
	return m.Kid
}

func (m staticpublickey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m staticpublickey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StaticPublicKeyFormatEnum Enum with underlying type: string
type StaticPublicKeyFormatEnum string

// Set of constants representing the allowable values for StaticPublicKeyFormatEnum
const (
	StaticPublicKeyFormatJsonWebKey StaticPublicKeyFormatEnum = "JSON_WEB_KEY"
	StaticPublicKeyFormatPem        StaticPublicKeyFormatEnum = "PEM"
)

var mappingStaticPublicKeyFormatEnum = map[string]StaticPublicKeyFormatEnum{
	"JSON_WEB_KEY": StaticPublicKeyFormatJsonWebKey,
	"PEM":          StaticPublicKeyFormatPem,
}

var mappingStaticPublicKeyFormatEnumLowerCase = map[string]StaticPublicKeyFormatEnum{
	"json_web_key": StaticPublicKeyFormatJsonWebKey,
	"pem":          StaticPublicKeyFormatPem,
}

// GetStaticPublicKeyFormatEnumValues Enumerates the set of values for StaticPublicKeyFormatEnum
func GetStaticPublicKeyFormatEnumValues() []StaticPublicKeyFormatEnum {
	values := make([]StaticPublicKeyFormatEnum, 0)
	for _, v := range mappingStaticPublicKeyFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetStaticPublicKeyFormatEnumStringValues Enumerates the set of values in String for StaticPublicKeyFormatEnum
func GetStaticPublicKeyFormatEnumStringValues() []string {
	return []string{
		"JSON_WEB_KEY",
		"PEM",
	}
}

// GetMappingStaticPublicKeyFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStaticPublicKeyFormatEnum(val string) (StaticPublicKeyFormatEnum, bool) {
	enum, ok := mappingStaticPublicKeyFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
