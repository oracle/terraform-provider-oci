// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// Sysctl A kernel parameter to be set.
type Sysctl struct {

	// Name of the kernel parameter to set. Allowed values are:
	// - `kernel.shm_rmid_forced`
	// - `net.ipv4.ip_local_port_range`
	// - `net.ipv4.tcp_syncookies`
	// - `net.ipv4.ping_group_range`
	Name *string `mandatory:"false" json:"name"`

	// Value of the kernel parameter to set.
	Value *string `mandatory:"false" json:"value"`
}

func (m Sysctl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Sysctl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
