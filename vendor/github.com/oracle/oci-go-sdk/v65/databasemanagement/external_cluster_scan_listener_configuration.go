// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalClusterScanListenerConfiguration The details of a SCAN listener in an external cluster.
type ExternalClusterScanListenerConfiguration struct {

	// The name of the SCAN listener.
	ScanName *string `mandatory:"false" json:"scanName"`

	// The network number from which SCAN VIPs are obtained.
	NetworkNumber *int `mandatory:"false" json:"networkNumber"`

	// The port number of the SCAN listener.
	ScanPort *int `mandatory:"false" json:"scanPort"`

	// The protocol of the SCAN listener.
	ScanProtocol ExternalClusterScanListenerConfigurationScanProtocolEnum `mandatory:"false" json:"scanProtocol,omitempty"`
}

func (m ExternalClusterScanListenerConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalClusterScanListenerConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExternalClusterScanListenerConfigurationScanProtocolEnum(string(m.ScanProtocol)); !ok && m.ScanProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanProtocol: %s. Supported values are: %s.", m.ScanProtocol, strings.Join(GetExternalClusterScanListenerConfigurationScanProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalClusterScanListenerConfigurationScanProtocolEnum Enum with underlying type: string
type ExternalClusterScanListenerConfigurationScanProtocolEnum string

// Set of constants representing the allowable values for ExternalClusterScanListenerConfigurationScanProtocolEnum
const (
	ExternalClusterScanListenerConfigurationScanProtocolTcp  ExternalClusterScanListenerConfigurationScanProtocolEnum = "TCP"
	ExternalClusterScanListenerConfigurationScanProtocolTcps ExternalClusterScanListenerConfigurationScanProtocolEnum = "TCPS"
)

var mappingExternalClusterScanListenerConfigurationScanProtocolEnum = map[string]ExternalClusterScanListenerConfigurationScanProtocolEnum{
	"TCP":  ExternalClusterScanListenerConfigurationScanProtocolTcp,
	"TCPS": ExternalClusterScanListenerConfigurationScanProtocolTcps,
}

var mappingExternalClusterScanListenerConfigurationScanProtocolEnumLowerCase = map[string]ExternalClusterScanListenerConfigurationScanProtocolEnum{
	"tcp":  ExternalClusterScanListenerConfigurationScanProtocolTcp,
	"tcps": ExternalClusterScanListenerConfigurationScanProtocolTcps,
}

// GetExternalClusterScanListenerConfigurationScanProtocolEnumValues Enumerates the set of values for ExternalClusterScanListenerConfigurationScanProtocolEnum
func GetExternalClusterScanListenerConfigurationScanProtocolEnumValues() []ExternalClusterScanListenerConfigurationScanProtocolEnum {
	values := make([]ExternalClusterScanListenerConfigurationScanProtocolEnum, 0)
	for _, v := range mappingExternalClusterScanListenerConfigurationScanProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalClusterScanListenerConfigurationScanProtocolEnumStringValues Enumerates the set of values in String for ExternalClusterScanListenerConfigurationScanProtocolEnum
func GetExternalClusterScanListenerConfigurationScanProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingExternalClusterScanListenerConfigurationScanProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalClusterScanListenerConfigurationScanProtocolEnum(val string) (ExternalClusterScanListenerConfigurationScanProtocolEnum, bool) {
	enum, ok := mappingExternalClusterScanListenerConfigurationScanProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
