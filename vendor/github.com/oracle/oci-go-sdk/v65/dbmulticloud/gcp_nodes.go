// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// <b>Microsoft Azure</b>:<br>
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
// <b>Google Cloud</b>:<br>
// 1. Oracle Google Cloud Connector Resource: This is for installing Google Identity in ExaCS VM Cluster.<br>
// 2. Discover Google Key-Rings and Keys Resource: This is for to discover Google Key-Rings.<br>
// 3. Google Key-Rings Resource: This is for to maintain Google Key-Rings in Oracle Cloud.<br>
// 4. Google Key Resource: This is for to maintain Google Key in Oracle Cloud for a Google Key-Ring.<br>
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GcpNodes GCP Identity Connector Node Details.
type GcpNodes struct {

	// Host Name or Identity Connector Name.
	HostName *string `mandatory:"false" json:"hostName"`

	// Host ID.
	HostId *string `mandatory:"false" json:"hostId"`

	// The current status of the GCP Identity Connector Resource.
	Status GcpNodesStatusEnum `mandatory:"false" json:"status,omitempty"`

	// time when the GCP Identity Connector's status was checked RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeLastChecked *common.SDKTime `mandatory:"false" json:"timeLastChecked"`
}

func (m GcpNodes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GcpNodes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGcpNodesStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetGcpNodesStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GcpNodesStatusEnum Enum with underlying type: string
type GcpNodesStatusEnum string

// Set of constants representing the allowable values for GcpNodesStatusEnum
const (
	GcpNodesStatusConnected    GcpNodesStatusEnum = "CONNECTED"
	GcpNodesStatusDisconnected GcpNodesStatusEnum = "DISCONNECTED"
	GcpNodesStatusUnknown      GcpNodesStatusEnum = "UNKNOWN"
)

var mappingGcpNodesStatusEnum = map[string]GcpNodesStatusEnum{
	"CONNECTED":    GcpNodesStatusConnected,
	"DISCONNECTED": GcpNodesStatusDisconnected,
	"UNKNOWN":      GcpNodesStatusUnknown,
}

var mappingGcpNodesStatusEnumLowerCase = map[string]GcpNodesStatusEnum{
	"connected":    GcpNodesStatusConnected,
	"disconnected": GcpNodesStatusDisconnected,
	"unknown":      GcpNodesStatusUnknown,
}

// GetGcpNodesStatusEnumValues Enumerates the set of values for GcpNodesStatusEnum
func GetGcpNodesStatusEnumValues() []GcpNodesStatusEnum {
	values := make([]GcpNodesStatusEnum, 0)
	for _, v := range mappingGcpNodesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGcpNodesStatusEnumStringValues Enumerates the set of values in String for GcpNodesStatusEnum
func GetGcpNodesStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"UNKNOWN",
	}
}

// GetMappingGcpNodesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGcpNodesStatusEnum(val string) (GcpNodesStatusEnum, bool) {
	enum, ok := mappingGcpNodesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
