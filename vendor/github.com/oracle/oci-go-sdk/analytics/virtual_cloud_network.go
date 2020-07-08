// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VirtualCloudNetwork Virtual Cloud Network definition.
type VirtualCloudNetwork struct {

	// The Virtual Cloud Network OCID.
	Id *string `mandatory:"true" json:"id"`

	// Source IP addresses or IP address ranges igress rules.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`
}

func (m VirtualCloudNetwork) String() string {
	return common.PointerString(m)
}
