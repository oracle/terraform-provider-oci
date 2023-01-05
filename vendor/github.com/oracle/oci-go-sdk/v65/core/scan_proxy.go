// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScanProxy Details pertaining to a scan proxy instance created for a scan listener FQDN/IPs
type ScanProxy struct {

	// An Oracle-assigned unique identifier for a scanProxy within a privateEndpoint. You specify
	// this ID when you want to get or update or delete a scanProxy
	Id *string `mandatory:"true" json:"id"`

	// Type indicating whether Scan listener is specified by its FQDN or list of IPs
	ScanListenerType ScanProxyScanListenerTypeEnum `mandatory:"true" json:"scanListenerType"`

	// The FQDN/IPs and port information of customer's Real Application Cluster (RAC)'s SCAN
	// listeners.
	ScanListenerInfo []ScanListenerInfo `mandatory:"true" json:"scanListenerInfo"`

	// The port to which service DB client has to connect on scan proxy to initiate scan
	// connections.
	ScanProxyPort *int `mandatory:"true" json:"scanProxyPort"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint
	// associated with the reverse connection.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// The protocol used for communication between client, scanProxy and RAC's scan
	// listeners
	Protocol ScanProxyProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	// Type indicating whether Scan proxy is IP multiplexing based or Port multiplexing based.
	ScanMultiplexingType ScanProxyScanMultiplexingTypeEnum `mandatory:"false" json:"scanMultiplexingType,omitempty"`

	// The IP address in the service VCN to be used to reach the reverse connection SCAN proxy
	// service.
	ScanProxyIp *string `mandatory:"false" json:"scanProxyIp"`

	// The scan proxy instance's current state.
	LifecycleState ScanProxyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the scan proxy instance was created, in the format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	ScanListenerWallet *WalletInfo `mandatory:"false" json:"scanListenerWallet"`
}

func (m ScanProxy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScanProxy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScanProxyScanListenerTypeEnum(string(m.ScanListenerType)); !ok && m.ScanListenerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanListenerType: %s. Supported values are: %s.", m.ScanListenerType, strings.Join(GetScanProxyScanListenerTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScanProxyProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetScanProxyProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScanProxyScanMultiplexingTypeEnum(string(m.ScanMultiplexingType)); !ok && m.ScanMultiplexingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanMultiplexingType: %s. Supported values are: %s.", m.ScanMultiplexingType, strings.Join(GetScanProxyScanMultiplexingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScanProxyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScanProxyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScanProxyScanListenerTypeEnum Enum with underlying type: string
type ScanProxyScanListenerTypeEnum string

// Set of constants representing the allowable values for ScanProxyScanListenerTypeEnum
const (
	ScanProxyScanListenerTypeFqdn ScanProxyScanListenerTypeEnum = "FQDN"
	ScanProxyScanListenerTypeIp   ScanProxyScanListenerTypeEnum = "IP"
)

var mappingScanProxyScanListenerTypeEnum = map[string]ScanProxyScanListenerTypeEnum{
	"FQDN": ScanProxyScanListenerTypeFqdn,
	"IP":   ScanProxyScanListenerTypeIp,
}

var mappingScanProxyScanListenerTypeEnumLowerCase = map[string]ScanProxyScanListenerTypeEnum{
	"fqdn": ScanProxyScanListenerTypeFqdn,
	"ip":   ScanProxyScanListenerTypeIp,
}

// GetScanProxyScanListenerTypeEnumValues Enumerates the set of values for ScanProxyScanListenerTypeEnum
func GetScanProxyScanListenerTypeEnumValues() []ScanProxyScanListenerTypeEnum {
	values := make([]ScanProxyScanListenerTypeEnum, 0)
	for _, v := range mappingScanProxyScanListenerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScanProxyScanListenerTypeEnumStringValues Enumerates the set of values in String for ScanProxyScanListenerTypeEnum
func GetScanProxyScanListenerTypeEnumStringValues() []string {
	return []string{
		"FQDN",
		"IP",
	}
}

// GetMappingScanProxyScanListenerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScanProxyScanListenerTypeEnum(val string) (ScanProxyScanListenerTypeEnum, bool) {
	enum, ok := mappingScanProxyScanListenerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScanProxyProtocolEnum Enum with underlying type: string
type ScanProxyProtocolEnum string

// Set of constants representing the allowable values for ScanProxyProtocolEnum
const (
	ScanProxyProtocolTcp  ScanProxyProtocolEnum = "TCP"
	ScanProxyProtocolTcps ScanProxyProtocolEnum = "TCPS"
)

var mappingScanProxyProtocolEnum = map[string]ScanProxyProtocolEnum{
	"TCP":  ScanProxyProtocolTcp,
	"TCPS": ScanProxyProtocolTcps,
}

var mappingScanProxyProtocolEnumLowerCase = map[string]ScanProxyProtocolEnum{
	"tcp":  ScanProxyProtocolTcp,
	"tcps": ScanProxyProtocolTcps,
}

// GetScanProxyProtocolEnumValues Enumerates the set of values for ScanProxyProtocolEnum
func GetScanProxyProtocolEnumValues() []ScanProxyProtocolEnum {
	values := make([]ScanProxyProtocolEnum, 0)
	for _, v := range mappingScanProxyProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetScanProxyProtocolEnumStringValues Enumerates the set of values in String for ScanProxyProtocolEnum
func GetScanProxyProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingScanProxyProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScanProxyProtocolEnum(val string) (ScanProxyProtocolEnum, bool) {
	enum, ok := mappingScanProxyProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScanProxyScanMultiplexingTypeEnum Enum with underlying type: string
type ScanProxyScanMultiplexingTypeEnum string

// Set of constants representing the allowable values for ScanProxyScanMultiplexingTypeEnum
const (
	ScanProxyScanMultiplexingTypePort ScanProxyScanMultiplexingTypeEnum = "PORT"
	ScanProxyScanMultiplexingTypeIp   ScanProxyScanMultiplexingTypeEnum = "IP"
)

var mappingScanProxyScanMultiplexingTypeEnum = map[string]ScanProxyScanMultiplexingTypeEnum{
	"PORT": ScanProxyScanMultiplexingTypePort,
	"IP":   ScanProxyScanMultiplexingTypeIp,
}

var mappingScanProxyScanMultiplexingTypeEnumLowerCase = map[string]ScanProxyScanMultiplexingTypeEnum{
	"port": ScanProxyScanMultiplexingTypePort,
	"ip":   ScanProxyScanMultiplexingTypeIp,
}

// GetScanProxyScanMultiplexingTypeEnumValues Enumerates the set of values for ScanProxyScanMultiplexingTypeEnum
func GetScanProxyScanMultiplexingTypeEnumValues() []ScanProxyScanMultiplexingTypeEnum {
	values := make([]ScanProxyScanMultiplexingTypeEnum, 0)
	for _, v := range mappingScanProxyScanMultiplexingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScanProxyScanMultiplexingTypeEnumStringValues Enumerates the set of values in String for ScanProxyScanMultiplexingTypeEnum
func GetScanProxyScanMultiplexingTypeEnumStringValues() []string {
	return []string{
		"PORT",
		"IP",
	}
}

// GetMappingScanProxyScanMultiplexingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScanProxyScanMultiplexingTypeEnum(val string) (ScanProxyScanMultiplexingTypeEnum, bool) {
	enum, ok := mappingScanProxyScanMultiplexingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScanProxyLifecycleStateEnum Enum with underlying type: string
type ScanProxyLifecycleStateEnum string

// Set of constants representing the allowable values for ScanProxyLifecycleStateEnum
const (
	ScanProxyLifecycleStateProvisioning ScanProxyLifecycleStateEnum = "PROVISIONING"
	ScanProxyLifecycleStateAvailable    ScanProxyLifecycleStateEnum = "AVAILABLE"
	ScanProxyLifecycleStateUpdating     ScanProxyLifecycleStateEnum = "UPDATING"
	ScanProxyLifecycleStateTerminating  ScanProxyLifecycleStateEnum = "TERMINATING"
	ScanProxyLifecycleStateTerminated   ScanProxyLifecycleStateEnum = "TERMINATED"
	ScanProxyLifecycleStateFailed       ScanProxyLifecycleStateEnum = "FAILED"
)

var mappingScanProxyLifecycleStateEnum = map[string]ScanProxyLifecycleStateEnum{
	"PROVISIONING": ScanProxyLifecycleStateProvisioning,
	"AVAILABLE":    ScanProxyLifecycleStateAvailable,
	"UPDATING":     ScanProxyLifecycleStateUpdating,
	"TERMINATING":  ScanProxyLifecycleStateTerminating,
	"TERMINATED":   ScanProxyLifecycleStateTerminated,
	"FAILED":       ScanProxyLifecycleStateFailed,
}

var mappingScanProxyLifecycleStateEnumLowerCase = map[string]ScanProxyLifecycleStateEnum{
	"provisioning": ScanProxyLifecycleStateProvisioning,
	"available":    ScanProxyLifecycleStateAvailable,
	"updating":     ScanProxyLifecycleStateUpdating,
	"terminating":  ScanProxyLifecycleStateTerminating,
	"terminated":   ScanProxyLifecycleStateTerminated,
	"failed":       ScanProxyLifecycleStateFailed,
}

// GetScanProxyLifecycleStateEnumValues Enumerates the set of values for ScanProxyLifecycleStateEnum
func GetScanProxyLifecycleStateEnumValues() []ScanProxyLifecycleStateEnum {
	values := make([]ScanProxyLifecycleStateEnum, 0)
	for _, v := range mappingScanProxyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScanProxyLifecycleStateEnumStringValues Enumerates the set of values in String for ScanProxyLifecycleStateEnum
func GetScanProxyLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingScanProxyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScanProxyLifecycleStateEnum(val string) (ScanProxyLifecycleStateEnum, bool) {
	enum, ok := mappingScanProxyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
