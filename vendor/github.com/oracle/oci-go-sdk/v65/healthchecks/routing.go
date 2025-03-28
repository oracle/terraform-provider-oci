// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Routing The routing information for a vantage point.
type Routing struct {

	// The registry label for `asn`, usually the name of the organization that
	// owns the ASN. May be omitted or null.
	AsLabel *string `mandatory:"false" json:"asLabel"`

	// The Autonomous System Number (ASN) identifying the organization
	// responsible for routing packets to `prefix`.
	Asn *int `mandatory:"false" json:"asn"`

	// An IP prefix (CIDR syntax) that is less specific than
	// `address`, through which `address` is routed.
	Prefix *string `mandatory:"false" json:"prefix"`

	// An integer between 0 and 100 used to select between multiple
	// origin ASNs when routing to `prefix`. Most prefixes have
	// exactly one origin ASN, in which case `weight` will be 100.
	Weight *int `mandatory:"false" json:"weight"`
}

func (m Routing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Routing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
