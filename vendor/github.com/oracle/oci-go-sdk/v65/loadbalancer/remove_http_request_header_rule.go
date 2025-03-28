// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveHttpRequestHeaderRule An object that represents the action of removing a header from a request. This rule applies only to HTTP listeners.
// If the same header appears more than once in the request, the load balancer removes all occurances of the specified header.
// **Note:** The system does not distinquish between underscore and dash characters in headers. That is, it treats
// `example_header_name` and `example-header-name` as identical. Oracle recommends that you do not rely on underscore
// or dash characters to uniquely distinguish header names.
type RemoveHttpRequestHeaderRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`
}

func (m RemoveHttpRequestHeaderRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveHttpRequestHeaderRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RemoveHttpRequestHeaderRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoveHttpRequestHeaderRule RemoveHttpRequestHeaderRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRemoveHttpRequestHeaderRule
	}{
		"REMOVE_HTTP_REQUEST_HEADER",
		(MarshalTypeRemoveHttpRequestHeaderRule)(m),
	}

	return json.Marshal(&s)
}
