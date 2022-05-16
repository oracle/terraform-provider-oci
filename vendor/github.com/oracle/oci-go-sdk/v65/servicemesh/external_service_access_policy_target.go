// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ExternalServiceAccessPolicyTarget External service target that internal virtual services direct traffic to.
type ExternalServiceAccessPolicyTarget struct {

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
	Protocol ExternalServiceAccessPolicyTargetProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m ExternalServiceAccessPolicyTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalServiceAccessPolicyTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalServiceAccessPolicyTargetProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetExternalServiceAccessPolicyTargetProtocolEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalServiceAccessPolicyTarget) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalServiceAccessPolicyTarget ExternalServiceAccessPolicyTarget
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExternalServiceAccessPolicyTarget
	}{
		"EXTERNAL_SERVICE",
		(MarshalTypeExternalServiceAccessPolicyTarget)(m),
	}

	return json.Marshal(&s)
}

// ExternalServiceAccessPolicyTargetProtocolEnum Enum with underlying type: string
type ExternalServiceAccessPolicyTargetProtocolEnum string

// Set of constants representing the allowable values for ExternalServiceAccessPolicyTargetProtocolEnum
const (
	ExternalServiceAccessPolicyTargetProtocolHttp  ExternalServiceAccessPolicyTargetProtocolEnum = "HTTP"
	ExternalServiceAccessPolicyTargetProtocolHttps ExternalServiceAccessPolicyTargetProtocolEnum = "HTTPS"
	ExternalServiceAccessPolicyTargetProtocolTcp   ExternalServiceAccessPolicyTargetProtocolEnum = "TCP"
)

var mappingExternalServiceAccessPolicyTargetProtocolEnum = map[string]ExternalServiceAccessPolicyTargetProtocolEnum{
	"HTTP":  ExternalServiceAccessPolicyTargetProtocolHttp,
	"HTTPS": ExternalServiceAccessPolicyTargetProtocolHttps,
	"TCP":   ExternalServiceAccessPolicyTargetProtocolTcp,
}

var mappingExternalServiceAccessPolicyTargetProtocolEnumLowerCase = map[string]ExternalServiceAccessPolicyTargetProtocolEnum{
	"http":  ExternalServiceAccessPolicyTargetProtocolHttp,
	"https": ExternalServiceAccessPolicyTargetProtocolHttps,
	"tcp":   ExternalServiceAccessPolicyTargetProtocolTcp,
}

// GetExternalServiceAccessPolicyTargetProtocolEnumValues Enumerates the set of values for ExternalServiceAccessPolicyTargetProtocolEnum
func GetExternalServiceAccessPolicyTargetProtocolEnumValues() []ExternalServiceAccessPolicyTargetProtocolEnum {
	values := make([]ExternalServiceAccessPolicyTargetProtocolEnum, 0)
	for _, v := range mappingExternalServiceAccessPolicyTargetProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalServiceAccessPolicyTargetProtocolEnumStringValues Enumerates the set of values in String for ExternalServiceAccessPolicyTargetProtocolEnum
func GetExternalServiceAccessPolicyTargetProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"HTTPS",
		"TCP",
	}
}

// GetMappingExternalServiceAccessPolicyTargetProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalServiceAccessPolicyTargetProtocolEnum(val string) (ExternalServiceAccessPolicyTargetProtocolEnum, bool) {
	enum, ok := mappingExternalServiceAccessPolicyTargetProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
