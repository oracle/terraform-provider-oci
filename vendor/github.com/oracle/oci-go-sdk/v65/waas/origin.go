// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Origin A detailed description of your web application's origin host server. An origin must be defined to set up WAF rules.
type Origin struct {

	// The URI of the origin. Does not support paths. Port numbers should be specified in the `httpPort` and `httpsPort` fields.
	Uri *string `mandatory:"true" json:"uri"`

	// The HTTP port on the origin that the web application listens on. If unspecified, defaults to `80`. If `0` is specified - the origin is not used for HTTP traffic.
	HttpPort *int `mandatory:"false" json:"httpPort"`

	// The HTTPS port on the origin that the web application listens on. If unspecified, defaults to `443`. If `0` is specified - the origin is not used for HTTPS traffic.
	HttpsPort *int `mandatory:"false" json:"httpsPort"`

	// A list of HTTP headers to forward to your origin.
	CustomHeaders []Header `mandatory:"false" json:"customHeaders"`
}

func (m Origin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Origin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
