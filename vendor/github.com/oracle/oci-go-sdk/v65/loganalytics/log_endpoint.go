// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogEndpoint An endpoint used to fetch logs.
type LogEndpoint struct {

	// The endpoint name.
	Name *string `mandatory:"true" json:"name"`

	Request *EndpointRequest `mandatory:"true" json:"request"`

	// The endpoint description.
	Description *string `mandatory:"false" json:"description"`

	// The endpoint model.
	Model *string `mandatory:"false" json:"model"`

	// The endpoint unique identifier.
	EndpointId *int64 `mandatory:"false" json:"endpointId"`

	Response *EndpointResponse `mandatory:"false" json:"response"`

	Credentials *EndpointCredentials `mandatory:"false" json:"credentials"`

	Proxy *EndpointProxy `mandatory:"false" json:"proxy"`

	// A flag indicating whether or not the endpoint is enabled for log collection.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The system flag. A value of false denotes a custom, or user
	// defined endpoint. A value of true denotes an Oracle defined endpoint.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// A list of endpoint properties.
	EndpointProperties []LogAnalyticsProperty `mandatory:"false" json:"endpointProperties"`
}

func (m LogEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
