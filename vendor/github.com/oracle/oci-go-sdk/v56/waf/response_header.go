// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResponseHeader A header field to add to a response.
type ResponseHeader struct {

	// The name of the header field.
	Name *string `mandatory:"true" json:"name"`

	// The value of the header field.
	Value *string `mandatory:"true" json:"value"`
}

func (m ResponseHeader) String() string {
	return common.PointerString(m)
}
