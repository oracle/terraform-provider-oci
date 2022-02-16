// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WafTrafficDatum A time series of traffic data for the  Web Application Firewall configured for a policy.
type WafTrafficDatum struct {

	// The date and time the traffic was observed, rounded down to the start of the range, and expressed in RFC 3339 timestamp format.
	TimeObserved *common.SDKTime `mandatory:"false" json:"timeObserved"`

	// The number of seconds this data covers.
	TimeRangeInSeconds *int `mandatory:"false" json:"timeRangeInSeconds"`

	// The tenancy OCID of the data.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The compartment OCID of the data.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The policy OCID of the data.
	WaasPolicyId *string `mandatory:"false" json:"waasPolicyId"`

	// Traffic in bytes.
	TrafficInBytes *int `mandatory:"false" json:"trafficInBytes"`
}

func (m WafTrafficDatum) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WafTrafficDatum) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
