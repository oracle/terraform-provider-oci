// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateScanProxyDetails Details for updating a scan proxy instance created for a scan listener FQDN/IPs
type UpdateScanProxyDetails struct {

	// Type indicating whether Scan listener is specified by its FQDN or list of IPs
	ScanListenerType ScanProxyScanListenerTypeEnum `mandatory:"false" json:"scanListenerType,omitempty"`

	// The FQDN/IPs and port information of customer's Real Application Cluster (RAC)'s SCAN
	// listeners.
	ScanListenerInfo []ScanListenerInfo `mandatory:"false" json:"scanListenerInfo"`

	// The protocol used for communication between client, scanProxy and RAC's scan
	// listeners
	Protocol ScanProxyProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	ScanListenerWallet *WalletInfo `mandatory:"false" json:"scanListenerWallet"`
}

func (m UpdateScanProxyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScanProxyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScanProxyScanListenerTypeEnum(string(m.ScanListenerType)); !ok && m.ScanListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanListenerType: %s. Supported values are: %s.", m.ScanListenerType, strings.Join(GetScanProxyScanListenerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScanProxyProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetScanProxyProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
