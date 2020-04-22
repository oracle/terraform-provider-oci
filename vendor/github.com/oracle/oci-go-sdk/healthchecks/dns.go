// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Dns The DNS resolution results.
type Dns struct {

	// Total DNS resolution duration, in milliseconds. Calculated using `domainLookupEnd`
	// minus `domainLookupStart`.
	DomainLookupDuration *float64 `mandatory:"false" json:"domainLookupDuration"`

	// The addresses returned by DNS resolution.
	Addresses []string `mandatory:"false" json:"addresses"`
}

func (m Dns) String() string {
	return common.PointerString(m)
}
