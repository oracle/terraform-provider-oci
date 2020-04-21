// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// StickyCookieLoadBalancingMethod An object that represents the `sticky-cookie` load balancing method and its properties.
type StickyCookieLoadBalancingMethod struct {

	// The name of the cookie used to track the persistence.
	// Can contain any US-ASCII character except separator or control character.
	Name *string `mandatory:"false" json:"name"`

	// The domain for which the cookie is set, defaults to WAAS policy domain.
	Domain *string `mandatory:"false" json:"domain"`

	// The time for which a browser should keep the cookie in seconds.
	// Empty value will cause the cookie to expire at the end of a browser session.
	ExpirationTimeInSeconds *int `mandatory:"false" json:"expirationTimeInSeconds"`
}

func (m StickyCookieLoadBalancingMethod) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StickyCookieLoadBalancingMethod) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStickyCookieLoadBalancingMethod StickyCookieLoadBalancingMethod
	s := struct {
		DiscriminatorParam string `json:"method"`
		MarshalTypeStickyCookieLoadBalancingMethod
	}{
		"STICKY_COOKIE",
		(MarshalTypeStickyCookieLoadBalancingMethod)(m),
	}

	return json.Marshal(&s)
}
