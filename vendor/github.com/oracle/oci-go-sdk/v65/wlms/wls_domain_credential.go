// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WlsDomainCredential Details of the WebLogic and Node Manager credentials.
type WlsDomainCredential struct {

	// The type of credential.
	Type *string `mandatory:"true" json:"type"`

	// The strategy for passing the new credentials.
	Strategy WlsDomainCredentialStrategyEnum `mandatory:"false" json:"strategy,omitempty"`

	// The OCID for user secret.
	UserSecretId *string `mandatory:"false" json:"userSecretId"`

	// The OCID for password secret.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`
}

func (m WlsDomainCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WlsDomainCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWlsDomainCredentialStrategyEnum(string(m.Strategy)); !ok && m.Strategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Strategy: %s. Supported values are: %s.", m.Strategy, strings.Join(GetWlsDomainCredentialStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WlsDomainCredentialStrategyEnum Enum with underlying type: string
type WlsDomainCredentialStrategyEnum string

// Set of constants representing the allowable values for WlsDomainCredentialStrategyEnum
const (
	WlsDomainCredentialStrategyDomainConfig      WlsDomainCredentialStrategyEnum = "USE_DOMAIN_CONFIG"
	WlsDomainCredentialStrategyNodeManagerConfig WlsDomainCredentialStrategyEnum = "USE_NODE_MANAGER_CONFIG"
	WlsDomainCredentialStrategySecrets           WlsDomainCredentialStrategyEnum = "USE_SECRETS"
)

var mappingWlsDomainCredentialStrategyEnum = map[string]WlsDomainCredentialStrategyEnum{
	"USE_DOMAIN_CONFIG":       WlsDomainCredentialStrategyDomainConfig,
	"USE_NODE_MANAGER_CONFIG": WlsDomainCredentialStrategyNodeManagerConfig,
	"USE_SECRETS":             WlsDomainCredentialStrategySecrets,
}

var mappingWlsDomainCredentialStrategyEnumLowerCase = map[string]WlsDomainCredentialStrategyEnum{
	"use_domain_config":       WlsDomainCredentialStrategyDomainConfig,
	"use_node_manager_config": WlsDomainCredentialStrategyNodeManagerConfig,
	"use_secrets":             WlsDomainCredentialStrategySecrets,
}

// GetWlsDomainCredentialStrategyEnumValues Enumerates the set of values for WlsDomainCredentialStrategyEnum
func GetWlsDomainCredentialStrategyEnumValues() []WlsDomainCredentialStrategyEnum {
	values := make([]WlsDomainCredentialStrategyEnum, 0)
	for _, v := range mappingWlsDomainCredentialStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetWlsDomainCredentialStrategyEnumStringValues Enumerates the set of values in String for WlsDomainCredentialStrategyEnum
func GetWlsDomainCredentialStrategyEnumStringValues() []string {
	return []string{
		"USE_DOMAIN_CONFIG",
		"USE_NODE_MANAGER_CONFIG",
		"USE_SECRETS",
	}
}

// GetMappingWlsDomainCredentialStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWlsDomainCredentialStrategyEnum(val string) (WlsDomainCredentialStrategyEnum, bool) {
	enum, ok := mappingWlsDomainCredentialStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
