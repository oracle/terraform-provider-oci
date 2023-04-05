// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HashKey The hash key used to calculate for the consistent hash load balancer like RingHashLoadBalancer or
// MaglevLoadBalancer. Only one of the httpCookie, httpHeaderName, isSourceIp or queryParameterName may be provided.
type HashKey struct {
	HttpCookie *HttpCookie `mandatory:"false" json:"httpCookie"`

	// The name of the request header that will be used to obtain the hash key.
	HttpHeaderName *string `mandatory:"false" json:"httpHeaderName"`

	// If true, the source ip address of the request will be used as hash key.
	IsSourceIp *bool `mandatory:"false" json:"isSourceIp"`

	// The name of the URL query parameter that will be used to obtain the hash key.
	QueryParameterName *string `mandatory:"false" json:"queryParameterName"`
}

func (m HashKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HashKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
