// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DigitalTwinAdapterInboundEnvelope Payload containing device-specific metadata and optional value mappings used to interpret
// or transform that metadata. This structure includes the device endpoint, the actual payload,
// and an optional envelope mapping that applies JQ (https://stedolan.github.io/jq/) expressions
// to extract or reshape the data as needed.
type DigitalTwinAdapterInboundEnvelope struct {

	// The device endpoint.
	ReferenceEndpoint *string `mandatory:"true" json:"referenceEndpoint"`

	ReferencePayload DigitalTwinAdapterPayload `mandatory:"false" json:"referencePayload"`

	EnvelopeMapping *DigitalTwinAdapterEnvelopeMapping `mandatory:"false" json:"envelopeMapping"`
}

func (m DigitalTwinAdapterInboundEnvelope) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DigitalTwinAdapterInboundEnvelope) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DigitalTwinAdapterInboundEnvelope) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ReferencePayload  digitaltwinadapterpayload          `json:"referencePayload"`
		EnvelopeMapping   *DigitalTwinAdapterEnvelopeMapping `json:"envelopeMapping"`
		ReferenceEndpoint *string                            `json:"referenceEndpoint"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ReferencePayload.UnmarshalPolymorphicJSON(model.ReferencePayload.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ReferencePayload = nn.(DigitalTwinAdapterPayload)
	} else {
		m.ReferencePayload = nil
	}

	m.EnvelopeMapping = model.EnvelopeMapping

	m.ReferenceEndpoint = model.ReferenceEndpoint

	return
}
