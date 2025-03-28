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

// HealthChecksVantagePointSummary Information about a vantage point.
type HealthChecksVantagePointSummary struct {

	// The display name for the vantage point. Display names are determined by
	// the best information available and may change over time.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The organization on whose infrastructure this vantage point resides.
	// Provider names are not unique, as Oracle Cloud Infrastructure maintains
	// many vantage points in each major provider.
	ProviderName *string `mandatory:"false" json:"providerName"`

	// The unique, permanent name for the vantage point.
	Name *string `mandatory:"false" json:"name"`

	Geo *Geolocation `mandatory:"false" json:"geo"`

	// An array of objects that describe how traffic to this vantage point is
	// routed, including which prefixes and ASNs connect it to the internet.
	// The addresses are sorted from the most-specific to least-specific
	// prefix (the smallest network to largest network). When a prefix has
	// multiple origin ASNs (MOAS routing), they are sorted by weight
	// (highest to lowest). Weight is determined by the total percentage of
	// peers observing the prefix originating from an ASN. Only present if
	// `fields` includes `routing`. The field will be null if the address's
	// routing information is unknown.
	Routing []Routing `mandatory:"false" json:"routing"`
}

func (m HealthChecksVantagePointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthChecksVantagePointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
