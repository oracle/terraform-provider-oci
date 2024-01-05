// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalServiceAccessPolicyTargetDetails External service target that internal virtual services direct traffic to.
type ExternalServiceAccessPolicyTargetDetails struct {

	// The hostnames of the external service. Only applicable for HTTP and HTTPS protocols.
	// Wildcard hostnames are supported in the prefix form.
	// Examples of valid hostnames are "www.example.com", "*.example.com", "*.com", "*".
	// Hostname "*" can be used to allow all hosts.
	Hostnames []string `mandatory:"false" json:"hostnames"`

	// The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol.
	// All requests matching the given CIDR notation will pass through.
	// In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication.
	IpAddresses []string `mandatory:"false" json:"ipAddresses"`

	// Ports exposed by an external service. If left empty all ports will be allowed.
	Ports []int `mandatory:"false" json:"ports"`

	// Protocol of the external service
	Protocol ExternalServiceAccessPolicyTargetDetailsProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m ExternalServiceAccessPolicyTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalServiceAccessPolicyTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalServiceAccessPolicyTargetDetailsProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetExternalServiceAccessPolicyTargetDetailsProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalServiceAccessPolicyTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalServiceAccessPolicyTargetDetails ExternalServiceAccessPolicyTargetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExternalServiceAccessPolicyTargetDetails
	}{
		"EXTERNAL_SERVICE",
		(MarshalTypeExternalServiceAccessPolicyTargetDetails)(m),
	}

	return json.Marshal(&s)
}

// ExternalServiceAccessPolicyTargetDetailsProtocolEnum Enum with underlying type: string
type ExternalServiceAccessPolicyTargetDetailsProtocolEnum string

// Set of constants representing the allowable values for ExternalServiceAccessPolicyTargetDetailsProtocolEnum
const (
	ExternalServiceAccessPolicyTargetDetailsProtocolHttp  ExternalServiceAccessPolicyTargetDetailsProtocolEnum = "HTTP"
	ExternalServiceAccessPolicyTargetDetailsProtocolHttps ExternalServiceAccessPolicyTargetDetailsProtocolEnum = "HTTPS"
	ExternalServiceAccessPolicyTargetDetailsProtocolTcp   ExternalServiceAccessPolicyTargetDetailsProtocolEnum = "TCP"
)

var mappingExternalServiceAccessPolicyTargetDetailsProtocolEnum = map[string]ExternalServiceAccessPolicyTargetDetailsProtocolEnum{
	"HTTP":  ExternalServiceAccessPolicyTargetDetailsProtocolHttp,
	"HTTPS": ExternalServiceAccessPolicyTargetDetailsProtocolHttps,
	"TCP":   ExternalServiceAccessPolicyTargetDetailsProtocolTcp,
}

var mappingExternalServiceAccessPolicyTargetDetailsProtocolEnumLowerCase = map[string]ExternalServiceAccessPolicyTargetDetailsProtocolEnum{
	"http":  ExternalServiceAccessPolicyTargetDetailsProtocolHttp,
	"https": ExternalServiceAccessPolicyTargetDetailsProtocolHttps,
	"tcp":   ExternalServiceAccessPolicyTargetDetailsProtocolTcp,
}

// GetExternalServiceAccessPolicyTargetDetailsProtocolEnumValues Enumerates the set of values for ExternalServiceAccessPolicyTargetDetailsProtocolEnum
func GetExternalServiceAccessPolicyTargetDetailsProtocolEnumValues() []ExternalServiceAccessPolicyTargetDetailsProtocolEnum {
	values := make([]ExternalServiceAccessPolicyTargetDetailsProtocolEnum, 0)
	for _, v := range mappingExternalServiceAccessPolicyTargetDetailsProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalServiceAccessPolicyTargetDetailsProtocolEnumStringValues Enumerates the set of values in String for ExternalServiceAccessPolicyTargetDetailsProtocolEnum
func GetExternalServiceAccessPolicyTargetDetailsProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"HTTPS",
		"TCP",
	}
}

// GetMappingExternalServiceAccessPolicyTargetDetailsProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalServiceAccessPolicyTargetDetailsProtocolEnum(val string) (ExternalServiceAccessPolicyTargetDetailsProtocolEnum, bool) {
	enum, ok := mappingExternalServiceAccessPolicyTargetDetailsProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
