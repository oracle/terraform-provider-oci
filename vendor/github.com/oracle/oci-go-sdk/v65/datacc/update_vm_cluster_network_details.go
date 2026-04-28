// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVmClusterNetworkDetails Details for an VM cluster network.
type UpdateVmClusterNetworkDetails struct {

	// The user-friendly name for the VM cluster network. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The SCAN details.
	Scans []ScanDetails `mandatory:"false" json:"scans"`

	// The listener TCP/IP port.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The listener TCP/IP SSL port. Default is 2484.
	ListenerPortSsl *int `mandatory:"false" json:"listenerPortSsl"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServers []string `mandatory:"false" json:"dnsServers"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServers []string `mandatory:"false" json:"ntpServers"`

	// Details of the client and backup networks.
	VmNetworks []VmNetworkDetails `mandatory:"false" json:"vmNetworks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateVmClusterNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVmClusterNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
