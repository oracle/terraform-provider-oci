// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// HttpHeaderRule An object that represents the advance http header options that allow the setting of http header size and allow/disallow
// invalid characters in the http headers.
// For example httpLargeHeaderSizeInKB=32, the http header could have 4 buffers of 32KBs each
// This rule applies only to HTTP listeners. No more than one `HttpHeaderRule` object can be present in
// a given listener.
type HttpHeaderRule struct {

	// Indicates whether or not invalid characters in client header fields will be allowed.
	// Valid names are composed of English letters, digits, hyphens and underscores.
	// If "true", invalid characters are allowed in the HTTP header.
	// If "false", invalid characters are not allowed in the HTTP header
	AreInvalidCharactersAllowed *bool `mandatory:"false" json:"areInvalidCharactersAllowed"`

	// The maximum size of each buffer used for reading http client request header.
	// This value indicates the maximum size allowed for each buffer.
	// The allowed values for buffer size are 8, 16, 32 and 64.
	HttpLargeHeaderSizeInKB *int `mandatory:"false" json:"httpLargeHeaderSizeInKB"`
}

func (m HttpHeaderRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HttpHeaderRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpHeaderRule HttpHeaderRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeHttpHeaderRule
	}{
		"HTTP_HEADER",
		(MarshalTypeHttpHeaderRule)(m),
	}

	return json.Marshal(&s)
}
