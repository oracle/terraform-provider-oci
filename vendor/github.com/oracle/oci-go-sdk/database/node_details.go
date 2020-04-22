// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NodeDetails Node details associated with a network.
type NodeDetails struct {

	// The node host name.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The node IP address.
	Ip *string `mandatory:"true" json:"ip"`

	// The node virtual IP (VIP) host name.
	VipHostname *string `mandatory:"false" json:"vipHostname"`

	// The node virtual IP (VIP) address.
	Vip *string `mandatory:"false" json:"vip"`
}

func (m NodeDetails) String() string {
	return common.PointerString(m)
}
