// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SingleSelectionSourcePolicy The single context variable in an incoming request to match against specified selection keys when dynamically routing and dynamically authenticating requests.
type SingleSelectionSourcePolicy struct {

	// String describing the context variable used as selector.
	Selector *string `mandatory:"true" json:"selector"`
}

func (m SingleSelectionSourcePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SingleSelectionSourcePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SingleSelectionSourcePolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleSelectionSourcePolicy SingleSelectionSourcePolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSingleSelectionSourcePolicy
	}{
		"SINGLE",
		(MarshalTypeSingleSelectionSourcePolicy)(m),
	}

	return json.Marshal(&s)
}
