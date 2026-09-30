// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PolicyConnectionOption Types of connection supported by Data Safe.
type PolicyConnectionOption struct {

	// The connection type used to connect to the database. Allowed values:
	// - PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
	// - ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe.
	ConnectionType PolicyConnectionOptionConnectionTypeEnum `mandatory:"true" json:"connectionType"`

	// List of OCIDs required to establish the connection.
	// - For `PRIVATE_ENDPOINT`, provide the OCID(s) of Data Safe private endpoint(s).
	// - For `ONPREM_CONNECTOR`, provide the OCID(s) of on-premises connector(s).
	Identifiers []string `mandatory:"true" json:"identifiers"`
}

func (m PolicyConnectionOption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PolicyConnectionOption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPolicyConnectionOptionConnectionTypeEnum(string(m.ConnectionType)); !ok && m.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", m.ConnectionType, strings.Join(GetPolicyConnectionOptionConnectionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PolicyConnectionOptionConnectionTypeEnum Enum with underlying type: string
type PolicyConnectionOptionConnectionTypeEnum string

// Set of constants representing the allowable values for PolicyConnectionOptionConnectionTypeEnum
const (
	PolicyConnectionOptionConnectionTypePrivateEndpoint PolicyConnectionOptionConnectionTypeEnum = "PRIVATE_ENDPOINT"
	PolicyConnectionOptionConnectionTypeOnpremConnector PolicyConnectionOptionConnectionTypeEnum = "ONPREM_CONNECTOR"
)

var mappingPolicyConnectionOptionConnectionTypeEnum = map[string]PolicyConnectionOptionConnectionTypeEnum{
	"PRIVATE_ENDPOINT": PolicyConnectionOptionConnectionTypePrivateEndpoint,
	"ONPREM_CONNECTOR": PolicyConnectionOptionConnectionTypeOnpremConnector,
}

var mappingPolicyConnectionOptionConnectionTypeEnumLowerCase = map[string]PolicyConnectionOptionConnectionTypeEnum{
	"private_endpoint": PolicyConnectionOptionConnectionTypePrivateEndpoint,
	"onprem_connector": PolicyConnectionOptionConnectionTypeOnpremConnector,
}

// GetPolicyConnectionOptionConnectionTypeEnumValues Enumerates the set of values for PolicyConnectionOptionConnectionTypeEnum
func GetPolicyConnectionOptionConnectionTypeEnumValues() []PolicyConnectionOptionConnectionTypeEnum {
	values := make([]PolicyConnectionOptionConnectionTypeEnum, 0)
	for _, v := range mappingPolicyConnectionOptionConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyConnectionOptionConnectionTypeEnumStringValues Enumerates the set of values in String for PolicyConnectionOptionConnectionTypeEnum
func GetPolicyConnectionOptionConnectionTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT",
		"ONPREM_CONNECTOR",
	}
}

// GetMappingPolicyConnectionOptionConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyConnectionOptionConnectionTypeEnum(val string) (PolicyConnectionOptionConnectionTypeEnum, bool) {
	enum, ok := mappingPolicyConnectionOptionConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
