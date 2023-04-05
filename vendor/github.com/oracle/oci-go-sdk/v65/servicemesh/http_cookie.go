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

// HttpCookie The HTTP cookie that will be used as the hash key for RingHashLoadBalancer or MaglevLoadBalancer.
// If the cookie doesn't exist, a new cookie will be generated.
type HttpCookie struct {

	// The name of the cookie that will be used to obtain the hash key. Generate new cookie if the cookie is
	// not present.
	Name *string `mandatory:"true" json:"name"`

	// Duration of cookie in seconds. This will be used to set the expiry time of a new cookie when it is
	// generated. If ttlInS is present and zero, the generated cookie will be a session cookie.
	TtlInS *int64 `mandatory:"false" json:"ttlInS"`

	// Path of cookie. This will be used to set the path of the generated cookie.
	Path *string `mandatory:"false" json:"path"`
}

func (m HttpCookie) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpCookie) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
