// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudClusterScanListenerConfiguration The details of a SCAN listener in a cloud cluster.
type CloudClusterScanListenerConfiguration struct {

	// The name of the SCAN listener.
	ScanName *string `mandatory:"false" json:"scanName"`

	// The network number from which SCAN VIPs are obtained.
	NetworkNumber *int `mandatory:"false" json:"networkNumber"`

	// The port number of the SCAN listener.
	ScanPort *int `mandatory:"false" json:"scanPort"`

	// The protocol of the SCAN listener.
	ScanProtocol CloudClusterScanListenerConfigurationScanProtocolEnum `mandatory:"false" json:"scanProtocol,omitempty"`
}

func (m CloudClusterScanListenerConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudClusterScanListenerConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCloudClusterScanListenerConfigurationScanProtocolEnum(string(m.ScanProtocol)); !ok && m.ScanProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanProtocol: %s. Supported values are: %s.", m.ScanProtocol, strings.Join(GetCloudClusterScanListenerConfigurationScanProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudClusterScanListenerConfigurationScanProtocolEnum Enum with underlying type: string
type CloudClusterScanListenerConfigurationScanProtocolEnum string

// Set of constants representing the allowable values for CloudClusterScanListenerConfigurationScanProtocolEnum
const (
	CloudClusterScanListenerConfigurationScanProtocolTcp  CloudClusterScanListenerConfigurationScanProtocolEnum = "TCP"
	CloudClusterScanListenerConfigurationScanProtocolTcps CloudClusterScanListenerConfigurationScanProtocolEnum = "TCPS"
)

var mappingCloudClusterScanListenerConfigurationScanProtocolEnum = map[string]CloudClusterScanListenerConfigurationScanProtocolEnum{
	"TCP":  CloudClusterScanListenerConfigurationScanProtocolTcp,
	"TCPS": CloudClusterScanListenerConfigurationScanProtocolTcps,
}

var mappingCloudClusterScanListenerConfigurationScanProtocolEnumLowerCase = map[string]CloudClusterScanListenerConfigurationScanProtocolEnum{
	"tcp":  CloudClusterScanListenerConfigurationScanProtocolTcp,
	"tcps": CloudClusterScanListenerConfigurationScanProtocolTcps,
}

// GetCloudClusterScanListenerConfigurationScanProtocolEnumValues Enumerates the set of values for CloudClusterScanListenerConfigurationScanProtocolEnum
func GetCloudClusterScanListenerConfigurationScanProtocolEnumValues() []CloudClusterScanListenerConfigurationScanProtocolEnum {
	values := make([]CloudClusterScanListenerConfigurationScanProtocolEnum, 0)
	for _, v := range mappingCloudClusterScanListenerConfigurationScanProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudClusterScanListenerConfigurationScanProtocolEnumStringValues Enumerates the set of values in String for CloudClusterScanListenerConfigurationScanProtocolEnum
func GetCloudClusterScanListenerConfigurationScanProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingCloudClusterScanListenerConfigurationScanProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudClusterScanListenerConfigurationScanProtocolEnum(val string) (CloudClusterScanListenerConfigurationScanProtocolEnum, bool) {
	enum, ok := mappingCloudClusterScanListenerConfigurationScanProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
