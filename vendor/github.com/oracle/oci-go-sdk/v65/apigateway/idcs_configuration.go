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

// IdcsConfiguration IDCS infra stripe details.
type IdcsConfiguration struct {

	// This is the infra stripe name of the tenant.
	InfraTenantStripeName *string `mandatory:"true" json:"infraTenantStripeName"`

	ClientTenantStripe ClientTenantStripe `mandatory:"true" json:"clientTenantStripe"`
}

func (m IdcsConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdcsConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IdcsConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InfraTenantStripeName *string            `json:"infraTenantStripeName"`
		ClientTenantStripe    clienttenantstripe `json:"clientTenantStripe"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InfraTenantStripeName = model.InfraTenantStripeName

	nn, e = model.ClientTenantStripe.UnmarshalPolymorphicJSON(model.ClientTenantStripe.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ClientTenantStripe = nn.(ClientTenantStripe)
	} else {
		m.ClientTenantStripe = nil
	}

	return
}
