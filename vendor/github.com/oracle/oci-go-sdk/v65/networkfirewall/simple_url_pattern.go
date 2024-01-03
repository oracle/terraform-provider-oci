// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SimpleUrlPattern Pattern describing an http/https URL or set thereof
// as a concatenation of optional host component and optional path component.
// `*.example.com` will match http://example.com/ and https://foo.example.com/foo?bar.
// `www.example.com/foo*` will match https://www.example.com/foo and http://www.exampe.com/foobar and https://www.example.com/foo/bar?baz, but not
// http://sub.www.example.com/foo or https://www.example.com/FOO.
// `*.example.com/foo*` will match http://example.com/foo and https://sub2.sub.example.com/foo/bar?baz, but not http://example.com/FOO.
type SimpleUrlPattern struct {

	// A string consisting of a concatenation of optional host component and optional path component.
	// The host component may start with `*.` to match the case-insensitive domain and all its subdomains.
	// The path component must start with a `/`, and may end with `*` to match all paths of which it is a case-sensitive prefix.
	// A missing host component matches all request domains, and a missing path component matches all request paths.
	// An empty value matches all requests.
	Pattern *string `mandatory:"true" json:"pattern"`
}

func (m SimpleUrlPattern) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SimpleUrlPattern) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SimpleUrlPattern) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSimpleUrlPattern SimpleUrlPattern
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSimpleUrlPattern
	}{
		"SIMPLE",
		(MarshalTypeSimpleUrlPattern)(m),
	}

	return json.Marshal(&s)
}
