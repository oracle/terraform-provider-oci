// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateContainerDnsConfigDetails Allow customers to define DNS settings for containers. If this is not provided, the containers use
// the default DNS settings of the subnet.
type CreateContainerDnsConfigDetails struct {

	// IP address of a name server that the resolver should query, either an IPv4 address
	// (in dot notation), or an IPv6 address in colon (and possibly dot) notation. If null, uses
	// nameservers from subnet dhcpDnsOptions.
	Nameservers []string `mandatory:"false" json:"nameservers"`

	// Search list for host-name lookup. If null, we will use searches from subnet dhcpDnsOptios.
	Searches []string `mandatory:"false" json:"searches"`

	// Options allows certain internal resolver variables to be modified. Options are a list of objects in
	// https://man7.org/linux/man-pages/man5/resolv.conf.5.html. Examples: ["ndots:n", "edns0"].
	Options []string `mandatory:"false" json:"options"`
}

func (m CreateContainerDnsConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerDnsConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
