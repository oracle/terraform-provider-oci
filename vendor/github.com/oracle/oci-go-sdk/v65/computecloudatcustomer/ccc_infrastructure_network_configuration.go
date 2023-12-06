// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructureNetworkConfiguration Configuration information for the Compute Cloud@Customer infrastructure. This
// network configuration information cannot be updated and is retrieved from the data center.
// The information will only be available
// after the connectionState is transitioned to CONNECTED.
type CccInfrastructureNetworkConfiguration struct {

	// Information about the management nodes that are provisioned in the Compute Cloud@Customer
	// infrastructure.
	ManagementNodes []CccInfrastructureManagementNode `mandatory:"false" json:"managementNodes"`

	// Uplink port speed defined in gigabytes per second.
	// All uplink ports must have identical speed.
	UplinkPortSpeedInGbps *int `mandatory:"false" json:"uplinkPortSpeedInGbps"`

	// Number of uplink ports per spine switch. Connectivity is identical on both spine switches.
	// For example, if input is two 100 gigabyte ports; then port-1 and port-2 on both spines will be configured.
	UplinkPortCount *int `mandatory:"false" json:"uplinkPortCount"`

	// The virtual local area network (VLAN) maximum transmission unit (MTU) size
	// for the uplink ports.
	UplinkVlanMtu *int `mandatory:"false" json:"uplinkVlanMtu"`

	// Netmask of the subnet that the Compute Cloud@Customer infrastructure is
	// connected to.
	UplinkNetmask *string `mandatory:"false" json:"uplinkNetmask"`

	// The port forward error correction (FEC) setting for the uplink port on the
	// Compute Cloud@Customer infrastructure.
	UplinkPortForwardErrorCorrection CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum `mandatory:"false" json:"uplinkPortForwardErrorCorrection,omitempty"`

	// Domain name to be used as the base domain for the internal network and by
	// public facing services.
	UplinkDomain *string `mandatory:"false" json:"uplinkDomain"`

	// Uplink gateway in the datacenter network that the Compute Cloud@Customer
	// connects to.
	UplinkGatewayIp *string `mandatory:"false" json:"uplinkGatewayIp"`

	// Addresses of the network spine switches.
	SpineIps []string `mandatory:"false" json:"spineIps"`

	// The spine switch public virtual IP (VIP). Traffic routed to the
	// Compute Cloud@Customer infrastructure and
	// and virtual cloud networks (VCNs) should have this address as next hop.
	SpineVip *string `mandatory:"false" json:"spineVip"`

	// The hostname corresponding to the virtual IP (VIP) address of the management nodes.
	MgmtVipHostname *string `mandatory:"false" json:"mgmtVipHostname"`

	// The IP address used as the virtual IP (VIP) address of the management nodes.
	MgmtVipIp *string `mandatory:"false" json:"mgmtVipIp"`

	// The domain name system (DNS) addresses that the Compute Cloud@Customer infrastructure
	// uses for the data center network.
	DnsIps []string `mandatory:"false" json:"dnsIps"`

	InfrastructureRoutingStatic *CccInfrastructureRoutingStaticDetails `mandatory:"false" json:"infrastructureRoutingStatic"`

	InfrastructureRoutingDynamic *CccInfrastructureRoutingDynamicDetails `mandatory:"false" json:"infrastructureRoutingDynamic"`
}

func (m CccInfrastructureNetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructureNetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum(string(m.UplinkPortForwardErrorCorrection)); !ok && m.UplinkPortForwardErrorCorrection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UplinkPortForwardErrorCorrection: %s. Supported values are: %s.", m.UplinkPortForwardErrorCorrection, strings.Join(GetCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum Enum with underlying type: string
type CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum string

// Set of constants representing the allowable values for CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum
const (
	CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionAuto                    CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = "AUTO"
	CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionFireCodeFec             CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = "FIRE_CODE_FEC"
	CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonConsortium16 CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = "REED_SOLOMON_CONSORTIUM_16"
	CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonFec          CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = "REED_SOLOMON_FEC"
	CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonIeee         CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = "REED_SOLOMON_IEEE"
)

var mappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum = map[string]CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum{
	"AUTO":                       CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionAuto,
	"FIRE_CODE_FEC":              CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionFireCodeFec,
	"REED_SOLOMON_CONSORTIUM_16": CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonConsortium16,
	"REED_SOLOMON_FEC":           CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonFec,
	"REED_SOLOMON_IEEE":          CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonIeee,
}

var mappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumLowerCase = map[string]CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum{
	"auto":                       CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionAuto,
	"fire_code_fec":              CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionFireCodeFec,
	"reed_solomon_consortium_16": CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonConsortium16,
	"reed_solomon_fec":           CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonFec,
	"reed_solomon_ieee":          CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionReedSolomonIeee,
}

// GetCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumValues Enumerates the set of values for CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum
func GetCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumValues() []CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum {
	values := make([]CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum, 0)
	for _, v := range mappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum {
		values = append(values, v)
	}
	return values
}

// GetCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumStringValues Enumerates the set of values in String for CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum
func GetCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumStringValues() []string {
	return []string{
		"AUTO",
		"FIRE_CODE_FEC",
		"REED_SOLOMON_CONSORTIUM_16",
		"REED_SOLOMON_FEC",
		"REED_SOLOMON_IEEE",
	}
}

// GetMappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum(val string) (CccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnum, bool) {
	enum, ok := mappingCccInfrastructureNetworkConfigurationUplinkPortForwardErrorCorrectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
