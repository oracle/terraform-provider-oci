// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OdaPrivateEndpointScanProxySummary Details pertaining to a scan proxy instance created for a scan listener FQDN/IPs
type OdaPrivateEndpointScanProxySummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Scan Proxy.
	Id *string `mandatory:"true" json:"id"`

	// Type indicating whether Scan listener is specified by its FQDN or list of IPs
	ScanListenerType OdaPrivateEndpointScanProxyScanListenerTypeEnum `mandatory:"true" json:"scanListenerType"`

	// The protocol used for communication between client, scanProxy and RAC's scan listeners
	Protocol OdaPrivateEndpointScanProxyProtocolEnum `mandatory:"true" json:"protocol"`

	// The FQDN/IPs and port information of customer's Real Application Cluster (RAC)'s SCAN listeners.
	ScanListenerInfos []ScanListenerInfo `mandatory:"true" json:"scanListenerInfos"`

	// The current state of the ODA Private Endpoint Scan Proxy.
	LifecycleState OdaPrivateEndpointScanProxyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m OdaPrivateEndpointScanProxySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaPrivateEndpointScanProxySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaPrivateEndpointScanProxyScanListenerTypeEnum(string(m.ScanListenerType)); !ok && m.ScanListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanListenerType: %s. Supported values are: %s.", m.ScanListenerType, strings.Join(GetOdaPrivateEndpointScanProxyScanListenerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdaPrivateEndpointScanProxyProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetOdaPrivateEndpointScanProxyProtocolEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOdaPrivateEndpointScanProxyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaPrivateEndpointScanProxyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
