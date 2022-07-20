// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
// Key points:
// 1. Use the specific URLs or construct the URL patterns you want to match using wildcard `*` and token characters `./`.
// 2. In the host component and path component don't enter consecutive wildcard `*` for example: `example.*.*.com`.
// 3. To match the exact domain and its subdirectories add a trailing slash `/` to your entry.
// 4. The firewall interprets entries that do not end in a `/` or `*` with an implicit asterisk `*` to their end, which further increases the potential matches.
// 5. The Use of multiple wildcards in a single pattern can impact the performance of the firewall.
// Domain/subdomains match examples:
// 1. `*.example.com.*` will match `sub1.sub2.example.com.info.us` and `sub1.example.com.us` and `sub1.example.com/subdirectory` but not match `sub1.example.com`.
// 2. `*.example.com` will match `sub1.example.com` and `sub1.sub2.example.com` and `sub1.example.com.au` but not match `example.com`.
// 3. `*.example.com/` will match `sub1.sub2.example.com` and `sub1.example.com` but not match `sub1.example.com.au`.
// 4. `example.com` will match `example.com` and `example.com.au` and `example.com.info.us` and `example.com/subdirectory`.
// 5. `example.com/` will match `example.com` and `example.com/foo` but not `example.com.info.us`.
// 6. `example.*.com` will match `example.sub1.com` and `example.sub1.sub2.com` and `example.sub1.com.au` and `example.sub1.com/subdirecroty`.
// 7. `example.*.com/` will match `example.sub1.com` and `example.sub1.sub2.com` and `example.sub1.com/subdirecroty` but not match `example.sub1.com.au`.
// 8. `example.com.*` will match `example.com.us` and `example.com.info.us` and `example.com.us/subdirectory` but not match `sub1.example.com`.
// Subdirectory or path match examples:
// 1. `example.com/*` will match `example.com/foo` and `example.com/bar` and any `example.com/subdirectory`.
// 2. `example.com/foo` will match `example.com/foo`.
// 3. `www.example.com/foo/*` will match `www.exampe.com/foo/subdiectory`, but not match `www.example.com/FOO` or `www.example.com/bar/subdirectory`.
// 4. `*.example.com/foo/*` will match `sub2.sub.example.com/foo/subdirectory` but not match `sub1.example.com/FOO` or `sub1.example.com/bar/subdirectory`.
// Other examples containing IP addresses in urls:
// 1. 103.12.14.122/ will match 103.12.14.122 and 103.12.14.122/subdirectory.
// 2. 103.12.14.122:8081/ will match 103.12.14.122:8081 and 103.12.14.122:8081/subdirectory.
// 3. 2607:9d80:4680:3f01:0000:0000:00d0:00c0/ will match 2607:9d80:4680:3f01:0000:0000:00d0:00c0 and 2607:9d80:4680:3f01:0000:0000:00d0:00c0/subdirecroty.
// 4. [2607:9d80:4680:3f01:0000:0000:00d0:00c0]:8081/ will match [2607:9d80:4680:3f01:0000:0000:00d0:00c0]:8081 and [2607:9d80:4680:3f01:0000:0000:00d0:00c0]:8081/subdirecroty.
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
