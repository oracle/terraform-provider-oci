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

// InvokeRawBinaryCommandDetails Binary data payload definition
type InvokeRawBinaryCommandDetails struct {

	// Device endpoint where request should be forwarded to.
	RequestEndpoint *string `mandatory:"true" json:"requestEndpoint"`

	// Specified duration by which to send the request by.
	RequestDuration *string `mandatory:"false" json:"requestDuration"`

	// Specified duration by which to receive the response by.
	ResponseDuration *string `mandatory:"false" json:"responseDuration"`

	// Device endpoint from which response is expected to come.
	ResponseEndpoint *string `mandatory:"false" json:"responseEndpoint"`

	// Mime content type of data encoded using base64, default is application/octet-stream
	RequestDataContentType *string `mandatory:"false" json:"requestDataContentType"`

	// base 64 encoded request data
	RequestData *string `mandatory:"false" json:"requestData"`
}

// GetRequestDuration returns RequestDuration
func (m InvokeRawBinaryCommandDetails) GetRequestDuration() *string {
	return m.RequestDuration
}

// GetResponseDuration returns ResponseDuration
func (m InvokeRawBinaryCommandDetails) GetResponseDuration() *string {
	return m.ResponseDuration
}

// GetRequestEndpoint returns RequestEndpoint
func (m InvokeRawBinaryCommandDetails) GetRequestEndpoint() *string {
	return m.RequestEndpoint
}

// GetResponseEndpoint returns ResponseEndpoint
func (m InvokeRawBinaryCommandDetails) GetResponseEndpoint() *string {
	return m.ResponseEndpoint
}

func (m InvokeRawBinaryCommandDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvokeRawBinaryCommandDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InvokeRawBinaryCommandDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInvokeRawBinaryCommandDetails InvokeRawBinaryCommandDetails
	s := struct {
		DiscriminatorParam string `json:"requestDataFormat"`
		MarshalTypeInvokeRawBinaryCommandDetails
	}{
		"BINARY",
		(MarshalTypeInvokeRawBinaryCommandDetails)(m),
	}

	return json.Marshal(&s)
}
