// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGatewayDetails Information about the new gateway.
type CreateGatewayDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be
	// accessible on a private IP address on the subnet.
	// Example: `PUBLIC` or `PRIVATE`
	EndpointType GatewayEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which
	// related resources are created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An array of Network Security Groups OCIDs associated with this API Gateway.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource which can be
	// empty string.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	ResponseCacheDetails ResponseCacheDetails `mandatory:"false" json:"responseCacheDetails"`

	// Locks associated with this resource.
	Locks []AddResourceLockDetails `mandatory:"false" json:"locks"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of CA bundles that should be used on the Gateway for TLS validation.
	CaBundles []CaBundle `mandatory:"false" json:"caBundles"`

	// Determines whether the gateway has an IPv4 or IPv6 address assigned to it, or both.
	// `IPV4` means the gateway will only have an IPv4 address assigned to it, and `IPV6` means the gateway will
	// only have an `IPv6` address assigned to it. `DUAL_STACK` means the gateway will have both an IPv4 and IPv6
	// address assigned to it.
	// Example: `IPV4` or `IPV6` or `DUAL_STACK`
	IpMode GatewayIpModeEnum `mandatory:"false" json:"ipMode,omitempty"`

	Ipv6AddressConfiguration *Ipv6AddressConfiguration `mandatory:"false" json:"ipv6AddressConfiguration"`

	Ipv4AddressConfiguration *Ipv4AddressConfiguration `mandatory:"false" json:"ipv4AddressConfiguration"`
}

func (m CreateGatewayDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGatewayDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGatewayEndpointTypeEnum(string(m.EndpointType)); !ok && m.EndpointType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EndpointType: %s. Supported values are: %s.", m.EndpointType, strings.Join(GetGatewayEndpointTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingGatewayIpModeEnum(string(m.IpMode)); !ok && m.IpMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpMode: %s. Supported values are: %s.", m.IpMode, strings.Join(GetGatewayIpModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateGatewayDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		NetworkSecurityGroupIds  []string                          `json:"networkSecurityGroupIds"`
		CertificateId            *string                           `json:"certificateId"`
		ResponseCacheDetails     responsecachedetails              `json:"responseCacheDetails"`
		Locks                    []AddResourceLockDetails          `json:"locks"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		CaBundles                []cabundle                        `json:"caBundles"`
		IpMode                   GatewayIpModeEnum                 `json:"ipMode"`
		Ipv6AddressConfiguration *Ipv6AddressConfiguration         `json:"ipv6AddressConfiguration"`
		Ipv4AddressConfiguration *Ipv4AddressConfiguration         `json:"ipv4AddressConfiguration"`
		CompartmentId            *string                           `json:"compartmentId"`
		EndpointType             GatewayEndpointTypeEnum           `json:"endpointType"`
		SubnetId                 *string                           `json:"subnetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.NetworkSecurityGroupIds = make([]string, len(model.NetworkSecurityGroupIds))
	copy(m.NetworkSecurityGroupIds, model.NetworkSecurityGroupIds)
	m.CertificateId = model.CertificateId

	nn, e = model.ResponseCacheDetails.UnmarshalPolymorphicJSON(model.ResponseCacheDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseCacheDetails = nn.(ResponseCacheDetails)
	} else {
		m.ResponseCacheDetails = nil
	}

	m.Locks = make([]AddResourceLockDetails, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CaBundles = make([]CaBundle, len(model.CaBundles))
	for i, n := range model.CaBundles {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.CaBundles[i] = nn.(CaBundle)
		} else {
			m.CaBundles[i] = nil
		}
	}
	m.IpMode = model.IpMode

	m.Ipv6AddressConfiguration = model.Ipv6AddressConfiguration

	m.Ipv4AddressConfiguration = model.Ipv4AddressConfiguration

	m.CompartmentId = model.CompartmentId

	m.EndpointType = model.EndpointType

	m.SubnetId = model.SubnetId

	return
}
