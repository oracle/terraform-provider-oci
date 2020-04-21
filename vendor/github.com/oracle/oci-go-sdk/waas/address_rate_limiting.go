// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AddressRateLimiting The IP rate limiting configuration. Defines the amount of allowed requests from a unique IP address and the resulting block response code when that threshold is exceeded.
type AddressRateLimiting struct {

	// Enables or disables the address rate limiting Web Application Firewall feature.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The number of allowed requests per second from one IP address. If unspecified, defaults to `1`.
	AllowedRatePerAddress *int `mandatory:"false" json:"allowedRatePerAddress"`

	// The maximum number of requests allowed to be queued before subsequent requests are dropped. If unspecified, defaults to `10`.
	MaxDelayedCountPerAddress *int `mandatory:"false" json:"maxDelayedCountPerAddress"`

	// The response status code returned when a request is blocked. If unspecified, defaults to `503`. The list of available response codes: `200`, `201`, `202`, `204`, `206`, `300`, `301`, `302`, `303`, `304`, `307`, `400`, `401`, `403`, `404`, `405`, `408`, `409`, `411`, `412`, `413`, `414`, `415`, `416`, `422`, `444`, `499`, `500`, `501`, `502`, `503`, `504`, `507`.
	BlockResponseCode *int `mandatory:"false" json:"blockResponseCode"`
}

func (m AddressRateLimiting) String() string {
	return common.PointerString(m)
}
