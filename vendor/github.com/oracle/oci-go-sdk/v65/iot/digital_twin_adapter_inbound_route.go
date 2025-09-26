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

// DigitalTwinAdapterInboundRoute Defines how inbound device payloads should be routed and mapped within a digital twin context.
// Routes are evaluated in the order they are defined, and only the first matching
// condition is processed. A final default route (with
// a condition that always evaluates to true) is recommended for fallback handling.
type DigitalTwinAdapterInboundRoute struct {

	// A boolean expression used to determine whether the following transformation
	// should be processed for the incoming payload. This expression is typically based
	// on fields defined at the inbound Envelope and is evaluated before applying the `payloadMapping`.
	Condition *string `mandatory:"true" json:"condition"`

	ReferencePayload DigitalTwinAdapterPayload `mandatory:"false" json:"referencePayload"`

	// A set of key-value JQ expressions used to transform the incoming payload into a shape
	// compatible with the digital twin model's context or schema.
	// The keys are target fields (in the digital twin model), and values are JQ expressions
	// pointing to data in the reference payload.
	// Example:
	// Given payload:
	// {
	//   "time": "<timestamp>",
	//   "temp": 65,
	//   "hum": 55
	// }
	// And mapping:
	// {
	//   "temperature": "$.temp",
	//   "humidity": "$.hum",
	//   "timeObserved": "$.time"
	// }
	// The output will be:
	// {
	//   "temperature": 65,
	//   "humidity": 55,
	//   "timeObserved": "<timestamp>"
	// }
	PayloadMapping map[string]string `mandatory:"false" json:"payloadMapping"`

	// Meaningful write up about the inbound route.
	Description *string `mandatory:"false" json:"description"`
}

func (m DigitalTwinAdapterInboundRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DigitalTwinAdapterInboundRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DigitalTwinAdapterInboundRoute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ReferencePayload digitaltwinadapterpayload `json:"referencePayload"`
		PayloadMapping   map[string]string         `json:"payloadMapping"`
		Description      *string                   `json:"description"`
		Condition        *string                   `json:"condition"`
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

	m.PayloadMapping = model.PayloadMapping

	m.Description = model.Description

	m.Condition = model.Condition

	return
}
