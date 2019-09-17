// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScanDetails The Single Client Access Name (SCAN) details.
type ScanDetails struct {

	// The SCAN hostname.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The SCAN port. Default is 1521.
	Port *int `mandatory:"true" json:"port"`

	// The list of SCAN IP addresses. Three addresses should be provided.
	Ips []string `mandatory:"true" json:"ips"`
}

func (m ScanDetails) String() string {
	return common.PointerString(m)
}
