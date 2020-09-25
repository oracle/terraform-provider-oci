// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UpdateInternalDnsRecordDetails This structure is used when updating DnsRecord for internal clients.
type UpdateInternalDnsRecordDetails struct {

	// Time to live value for the DnsRecord, according to RFC 1035 (https://tools.ietf.org/html/rfc1035).
	// Defaults to 86400.
	Ttl *int `mandatory:"false" json:"ttl"`

	// Value for the DnsRecord.
	// -*A:* One or more IPv4 addresses. Enter addresses on separate lines.
	Value *string `mandatory:"false" json:"value"`
}

func (m UpdateInternalDnsRecordDetails) String() string {
	return common.PointerString(m)
}
