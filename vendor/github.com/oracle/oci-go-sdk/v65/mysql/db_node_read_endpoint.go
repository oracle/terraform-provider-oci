// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbNodeReadEndpoint A read endpoint for a DB node.
// Each DB node has its own read endpoint in addition to the shared-storage DB cluster-level read endpoint.
type DbNodeReadEndpoint struct {

	// The hostname for the endpoint. Used for DNS. The default value is automatically assigned.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address for the endpoint. The default value is automatically assigned.
	// This IP address must be an available IP address within the subnet's CIDR.
	// This is a dotted-quad style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`
}

func (m DbNodeReadEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNodeReadEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
