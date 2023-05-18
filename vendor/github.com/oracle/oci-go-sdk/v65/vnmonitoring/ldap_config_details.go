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

// LdapConfigDetails The detail of LDAP's authentication configuration.
type LdapConfigDetails struct {

	// The IP address of primary LDAP server.
	PrimaryServer *string `mandatory:"false" json:"primaryServer"`

	// The IP address of secondary LDAP server.
	SecondaryServer *string `mandatory:"false" json:"secondaryServer"`

	// Option for LDAP SSL once customer enables SSL for certain ClientVpn.
	// Allowed values:
	// * `NEVER`: Do not use SSL (the default setting).
	// * `ADAPTIVE`: Try using SSL, if that fails, use plain-text to get through authentication.
	// * `ALWAYS`: Always use SSL.
	UseSsl LdapConfigDetailsUseSslEnum `mandatory:"false" json:"useSsl,omitempty"`

	// Choose the authentication method once useSSL enabled.
	// Allowed values:
	// * `NEVER`: No peer certificate is required.
	// * `ALLOW`: Request a peer certificate, but session will not be aborted if certificate cannot be validated.
	// * `DEMAND`: A valid peer certificate is required, then session will be aborted if one is not provided.
	VerifySsl LdapConfigDetailsVerifySslEnum `mandatory:"false" json:"verifySsl,omitempty"`

	// Enable case-sensitivity or not in LDAP authentication.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// Whether to apply Anonymous bind or not.
	IsBindAnon *bool `mandatory:"false" json:"isBindAnon"`

	// The bind DN (Distinguished Name) includes the user and location of the
	// user in LDAP directory tree
	BindDN *string `mandatory:"false" json:"bindDN"`

	// The bind password is used to log in the LDAP server.
	BindPW *string `mandatory:"false" json:"bindPW"`

	// The starting point element helps LDAP service to navigate search scope.
	BaseDN *string `mandatory:"false" json:"baseDN"`

	// The username of client at attribute level.
	ClientUsername *string `mandatory:"false" json:"clientUsername"`

	// This additional requirement uses LDAP query syntax. E.g., to require that the user be a member of a particular LDAP group (specified by DN) use this filter:
	// memberOf=CN=VPN Users, CN=Users, DC=example, DC=net
	AdditionalRequirements *string `mandatory:"false" json:"additionalRequirements"`
}

func (m LdapConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LdapConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLdapConfigDetailsUseSslEnum(string(m.UseSsl)); !ok && m.UseSsl != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UseSsl: %s. Supported values are: %s.", m.UseSsl, strings.Join(GetLdapConfigDetailsUseSslEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLdapConfigDetailsVerifySslEnum(string(m.VerifySsl)); !ok && m.VerifySsl != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VerifySsl: %s. Supported values are: %s.", m.VerifySsl, strings.Join(GetLdapConfigDetailsVerifySslEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LdapConfigDetailsUseSslEnum Enum with underlying type: string
type LdapConfigDetailsUseSslEnum string

// Set of constants representing the allowable values for LdapConfigDetailsUseSslEnum
const (
	LdapConfigDetailsUseSslNever    LdapConfigDetailsUseSslEnum = "NEVER"
	LdapConfigDetailsUseSslAdaptive LdapConfigDetailsUseSslEnum = "ADAPTIVE"
	LdapConfigDetailsUseSslAlways   LdapConfigDetailsUseSslEnum = "ALWAYS"
)

var mappingLdapConfigDetailsUseSslEnum = map[string]LdapConfigDetailsUseSslEnum{
	"NEVER":    LdapConfigDetailsUseSslNever,
	"ADAPTIVE": LdapConfigDetailsUseSslAdaptive,
	"ALWAYS":   LdapConfigDetailsUseSslAlways,
}

var mappingLdapConfigDetailsUseSslEnumLowerCase = map[string]LdapConfigDetailsUseSslEnum{
	"never":    LdapConfigDetailsUseSslNever,
	"adaptive": LdapConfigDetailsUseSslAdaptive,
	"always":   LdapConfigDetailsUseSslAlways,
}

// GetLdapConfigDetailsUseSslEnumValues Enumerates the set of values for LdapConfigDetailsUseSslEnum
func GetLdapConfigDetailsUseSslEnumValues() []LdapConfigDetailsUseSslEnum {
	values := make([]LdapConfigDetailsUseSslEnum, 0)
	for _, v := range mappingLdapConfigDetailsUseSslEnum {
		values = append(values, v)
	}
	return values
}

// GetLdapConfigDetailsUseSslEnumStringValues Enumerates the set of values in String for LdapConfigDetailsUseSslEnum
func GetLdapConfigDetailsUseSslEnumStringValues() []string {
	return []string{
		"NEVER",
		"ADAPTIVE",
		"ALWAYS",
	}
}

// GetMappingLdapConfigDetailsUseSslEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLdapConfigDetailsUseSslEnum(val string) (LdapConfigDetailsUseSslEnum, bool) {
	enum, ok := mappingLdapConfigDetailsUseSslEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LdapConfigDetailsVerifySslEnum Enum with underlying type: string
type LdapConfigDetailsVerifySslEnum string

// Set of constants representing the allowable values for LdapConfigDetailsVerifySslEnum
const (
	LdapConfigDetailsVerifySslNever  LdapConfigDetailsVerifySslEnum = "NEVER"
	LdapConfigDetailsVerifySslAllow  LdapConfigDetailsVerifySslEnum = "ALLOW"
	LdapConfigDetailsVerifySslDemand LdapConfigDetailsVerifySslEnum = "DEMAND"
)

var mappingLdapConfigDetailsVerifySslEnum = map[string]LdapConfigDetailsVerifySslEnum{
	"NEVER":  LdapConfigDetailsVerifySslNever,
	"ALLOW":  LdapConfigDetailsVerifySslAllow,
	"DEMAND": LdapConfigDetailsVerifySslDemand,
}

var mappingLdapConfigDetailsVerifySslEnumLowerCase = map[string]LdapConfigDetailsVerifySslEnum{
	"never":  LdapConfigDetailsVerifySslNever,
	"allow":  LdapConfigDetailsVerifySslAllow,
	"demand": LdapConfigDetailsVerifySslDemand,
}

// GetLdapConfigDetailsVerifySslEnumValues Enumerates the set of values for LdapConfigDetailsVerifySslEnum
func GetLdapConfigDetailsVerifySslEnumValues() []LdapConfigDetailsVerifySslEnum {
	values := make([]LdapConfigDetailsVerifySslEnum, 0)
	for _, v := range mappingLdapConfigDetailsVerifySslEnum {
		values = append(values, v)
	}
	return values
}

// GetLdapConfigDetailsVerifySslEnumStringValues Enumerates the set of values in String for LdapConfigDetailsVerifySslEnum
func GetLdapConfigDetailsVerifySslEnumStringValues() []string {
	return []string{
		"NEVER",
		"ALLOW",
		"DEMAND",
	}
}

// GetMappingLdapConfigDetailsVerifySslEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLdapConfigDetailsVerifySslEnum(val string) (LdapConfigDetailsVerifySslEnum, bool) {
	enum, ok := mappingLdapConfigDetailsVerifySslEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
