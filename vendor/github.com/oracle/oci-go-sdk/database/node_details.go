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

// NodeDetails Node details associated with a network.
type NodeDetails struct {

	// The node IP address.
	Ip *string `mandatory:"true" json:"ip"`

	// The node host name.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The node virtual IP (VIP) host name.
	VipHostname *string `mandatory:"false" json:"vipHostname"`

	// The node virtual IP (VIP) address.
	Vip *string `mandatory:"false" json:"vip"`
}

func (m NodeDetails) String() string {
	return common.PointerString(m)
}
