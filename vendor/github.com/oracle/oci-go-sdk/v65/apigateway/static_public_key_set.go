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

// StaticPublicKeySet A set of static public keys that will be used to verify the JWT signature.
type StaticPublicKeySet struct {

	// The set of static public keys.
	Keys []StaticPublicKey `mandatory:"false" json:"keys"`
}

func (m StaticPublicKeySet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StaticPublicKeySet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StaticPublicKeySet) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStaticPublicKeySet StaticPublicKeySet
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStaticPublicKeySet
	}{
		"STATIC_KEYS",
		(MarshalTypeStaticPublicKeySet)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *StaticPublicKeySet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Keys []staticpublickey `json:"keys"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Keys = make([]StaticPublicKey, len(model.Keys))
	for i, n := range model.Keys {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Keys[i] = nn.(StaticPublicKey)
		} else {
			m.Keys[i] = nil
		}
	}
	return
}
