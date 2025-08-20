// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HealthCheckMetaData HealthCheckMetaData, A combination of knob, port number and protocol used by SmartNics for the health checks.
type HealthCheckMetaData struct {

	// Whether health-check should be enabled for this VNIC.
	IsHealthCheckEnabled *bool `mandatory:"true" json:"isHealthCheckEnabled"`

	// The port number to use when performing health check on this VNIC.
	Port *int `mandatory:"true" json:"port"`

	// The type of protocol i.e. TCP, ARP or ALL accompanied by port number above to be used for health-check on this VNIC.
	Protocol HealthCheckMetaDataProtocolEnum `mandatory:"true" json:"protocol"`
}

func (m HealthCheckMetaData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthCheckMetaData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHealthCheckMetaDataProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetHealthCheckMetaDataProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HealthCheckMetaDataProtocolEnum Enum with underlying type: string
type HealthCheckMetaDataProtocolEnum string

// Set of constants representing the allowable values for HealthCheckMetaDataProtocolEnum
const (
	HealthCheckMetaDataProtocolArp HealthCheckMetaDataProtocolEnum = "ARP"
	HealthCheckMetaDataProtocolUdp HealthCheckMetaDataProtocolEnum = "UDP"
)

var mappingHealthCheckMetaDataProtocolEnum = map[string]HealthCheckMetaDataProtocolEnum{
	"ARP": HealthCheckMetaDataProtocolArp,
	"UDP": HealthCheckMetaDataProtocolUdp,
}

var mappingHealthCheckMetaDataProtocolEnumLowerCase = map[string]HealthCheckMetaDataProtocolEnum{
	"arp": HealthCheckMetaDataProtocolArp,
	"udp": HealthCheckMetaDataProtocolUdp,
}

// GetHealthCheckMetaDataProtocolEnumValues Enumerates the set of values for HealthCheckMetaDataProtocolEnum
func GetHealthCheckMetaDataProtocolEnumValues() []HealthCheckMetaDataProtocolEnum {
	values := make([]HealthCheckMetaDataProtocolEnum, 0)
	for _, v := range mappingHealthCheckMetaDataProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckMetaDataProtocolEnumStringValues Enumerates the set of values in String for HealthCheckMetaDataProtocolEnum
func GetHealthCheckMetaDataProtocolEnumStringValues() []string {
	return []string{
		"ARP",
		"UDP",
	}
}

// GetMappingHealthCheckMetaDataProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckMetaDataProtocolEnum(val string) (HealthCheckMetaDataProtocolEnum, bool) {
	enum, ok := mappingHealthCheckMetaDataProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
