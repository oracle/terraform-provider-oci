// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmInitialization Type representing initialization configuration of a virtual machine.
type OlvmInitialization struct {

	// Active Directory Organizational Unit.
	ActiveDirectoryOu *string `mandatory:"false" json:"activeDirectoryOu"`

	// Defines the values for the cloud-init protocol. This protocol decides how the cloud-init network parameters are formatted before being passed to the virtual machine in order to be processed by cloud-init.
	CloudInitNetworkProtocol OlvmInitializationCloudInitNetworkProtocolEnum `mandatory:"false" json:"cloudInitNetworkProtocol,omitempty"`

	Configuration *OlvmConfiguration `mandatory:"false" json:"configuration"`

	// Custom script that will be run when the VM starts.
	CustomScript *string `mandatory:"false" json:"customScript"`

	// DNS Search of the virtual machine.
	DnsSearch *string `mandatory:"false" json:"dnsSearch"`

	// DNS Servers of the virtual  machine.
	DnsServers *string `mandatory:"false" json:"dnsServers"`

	// Domain of the virtual machine.
	Domain *string `mandatory:"false" json:"domain"`

	// Host name of the virtual machine.
	HostName *string `mandatory:"false" json:"hostName"`

	// Input locale of the virtual machine.
	InputLocale *string `mandatory:"false" json:"inputLocale"`

	// The configuration of a virtual network interface.
	NicConfigurations []OlvmNicConfiguration `mandatory:"false" json:"nicConfigurations"`

	// Organization name.
	OrgName *string `mandatory:"false" json:"orgName"`

	// Indicates if new IDs should be regenerated.
	IsRegenerateIds *bool `mandatory:"false" json:"isRegenerateIds"`

	// Indicates if new SSH Keys should be regenerated.
	IsRegenerateSshKeys *bool `mandatory:"false" json:"isRegenerateSshKeys"`

	// System locale of the virtual machine.
	SystemLocale *string `mandatory:"false" json:"systemLocale"`

	// Timezone of the virtual machine.
	Timezone *string `mandatory:"false" json:"timezone"`

	// UI Language of the virtual machine.
	UiLanguage *string `mandatory:"false" json:"uiLanguage"`

	// User Locale of the virtual machine.
	UserLocale *string `mandatory:"false" json:"userLocale"`

	// User name of the virtual machine.
	Username *string `mandatory:"false" json:"username"`

	// Windows License Key of the virtual machine.
	WindowsLicenseKey *string `mandatory:"false" json:"windowsLicenseKey"`
}

func (m OlvmInitialization) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmInitialization) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmInitializationCloudInitNetworkProtocolEnum(string(m.CloudInitNetworkProtocol)); !ok && m.CloudInitNetworkProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloudInitNetworkProtocol: %s. Supported values are: %s.", m.CloudInitNetworkProtocol, strings.Join(GetOlvmInitializationCloudInitNetworkProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmInitializationCloudInitNetworkProtocolEnum Enum with underlying type: string
type OlvmInitializationCloudInitNetworkProtocolEnum string

// Set of constants representing the allowable values for OlvmInitializationCloudInitNetworkProtocolEnum
const (
	OlvmInitializationCloudInitNetworkProtocolEni               OlvmInitializationCloudInitNetworkProtocolEnum = "ENI"
	OlvmInitializationCloudInitNetworkProtocolOpenstackMetadata OlvmInitializationCloudInitNetworkProtocolEnum = "OPENSTACK_METADATA"
)

var mappingOlvmInitializationCloudInitNetworkProtocolEnum = map[string]OlvmInitializationCloudInitNetworkProtocolEnum{
	"ENI":                OlvmInitializationCloudInitNetworkProtocolEni,
	"OPENSTACK_METADATA": OlvmInitializationCloudInitNetworkProtocolOpenstackMetadata,
}

var mappingOlvmInitializationCloudInitNetworkProtocolEnumLowerCase = map[string]OlvmInitializationCloudInitNetworkProtocolEnum{
	"eni":                OlvmInitializationCloudInitNetworkProtocolEni,
	"openstack_metadata": OlvmInitializationCloudInitNetworkProtocolOpenstackMetadata,
}

// GetOlvmInitializationCloudInitNetworkProtocolEnumValues Enumerates the set of values for OlvmInitializationCloudInitNetworkProtocolEnum
func GetOlvmInitializationCloudInitNetworkProtocolEnumValues() []OlvmInitializationCloudInitNetworkProtocolEnum {
	values := make([]OlvmInitializationCloudInitNetworkProtocolEnum, 0)
	for _, v := range mappingOlvmInitializationCloudInitNetworkProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmInitializationCloudInitNetworkProtocolEnumStringValues Enumerates the set of values in String for OlvmInitializationCloudInitNetworkProtocolEnum
func GetOlvmInitializationCloudInitNetworkProtocolEnumStringValues() []string {
	return []string{
		"ENI",
		"OPENSTACK_METADATA",
	}
}

// GetMappingOlvmInitializationCloudInitNetworkProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmInitializationCloudInitNetworkProtocolEnum(val string) (OlvmInitializationCloudInitNetworkProtocolEnum, bool) {
	enum, ok := mappingOlvmInitializationCloudInitNetworkProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
