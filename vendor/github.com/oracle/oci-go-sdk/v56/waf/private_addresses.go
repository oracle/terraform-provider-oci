// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PrivateAddresses A pair of VCN OCID and private IP address prefix in CIDR notation.
type PrivateAddresses struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// A private IP address or CIDR IP address range.
	Addresses *string `mandatory:"true" json:"addresses"`
}

func (m PrivateAddresses) String() string {
	return common.PointerString(m)
}
