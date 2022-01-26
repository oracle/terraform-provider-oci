// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateResolverEndpointDetails The body for updating an existing resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateResolverEndpointDetails interface {
}

type updateresolverendpointdetails struct {
	JsonData     []byte
	EndpointType string `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *updateresolverendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateresolverendpointdetails updateresolverendpointdetails
	s := struct {
		Model Unmarshalerupdateresolverendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateresolverendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := UpdateResolverVnicEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updateresolverendpointdetails) String() string {
	return common.PointerString(m)
}

// UpdateResolverEndpointDetailsEndpointTypeEnum Enum with underlying type: string
type UpdateResolverEndpointDetailsEndpointTypeEnum string

// Set of constants representing the allowable values for UpdateResolverEndpointDetailsEndpointTypeEnum
const (
	UpdateResolverEndpointDetailsEndpointTypeVnic UpdateResolverEndpointDetailsEndpointTypeEnum = "VNIC"
)

var mappingUpdateResolverEndpointDetailsEndpointType = map[string]UpdateResolverEndpointDetailsEndpointTypeEnum{
	"VNIC": UpdateResolverEndpointDetailsEndpointTypeVnic,
}

// GetUpdateResolverEndpointDetailsEndpointTypeEnumValues Enumerates the set of values for UpdateResolverEndpointDetailsEndpointTypeEnum
func GetUpdateResolverEndpointDetailsEndpointTypeEnumValues() []UpdateResolverEndpointDetailsEndpointTypeEnum {
	values := make([]UpdateResolverEndpointDetailsEndpointTypeEnum, 0)
	for _, v := range mappingUpdateResolverEndpointDetailsEndpointType {
		values = append(values, v)
	}
	return values
}
