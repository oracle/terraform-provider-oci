// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArcAgentNodes Azure Arc Agent Node Details.
type ArcAgentNodes struct {

	// Host Name or Azure Arc Agent Name.
	HostName *string `mandatory:"false" json:"hostName"`

	// Host ID.
	HostId *string `mandatory:"false" json:"hostId"`

	// Current Arc Agent Version installed on this node of VM Cluster.
	CurrentArcAgentVersion *string `mandatory:"false" json:"currentArcAgentVersion"`

	// The current status of the Azure Arc Agent Resource.
	Status ArcAgentNodesStatusEnum `mandatory:"false" json:"status,omitempty"`

	// time when the Azure Arc Agent's status was checked RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeLastChecked *common.SDKTime `mandatory:"false" json:"timeLastChecked"`
}

func (m ArcAgentNodes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArcAgentNodes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingArcAgentNodesStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArcAgentNodesStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArcAgentNodesStatusEnum Enum with underlying type: string
type ArcAgentNodesStatusEnum string

// Set of constants representing the allowable values for ArcAgentNodesStatusEnum
const (
	ArcAgentNodesStatusConnected    ArcAgentNodesStatusEnum = "CONNECTED"
	ArcAgentNodesStatusDisconnected ArcAgentNodesStatusEnum = "DISCONNECTED"
	ArcAgentNodesStatusUnknown      ArcAgentNodesStatusEnum = "UNKNOWN"
)

var mappingArcAgentNodesStatusEnum = map[string]ArcAgentNodesStatusEnum{
	"CONNECTED":    ArcAgentNodesStatusConnected,
	"DISCONNECTED": ArcAgentNodesStatusDisconnected,
	"UNKNOWN":      ArcAgentNodesStatusUnknown,
}

var mappingArcAgentNodesStatusEnumLowerCase = map[string]ArcAgentNodesStatusEnum{
	"connected":    ArcAgentNodesStatusConnected,
	"disconnected": ArcAgentNodesStatusDisconnected,
	"unknown":      ArcAgentNodesStatusUnknown,
}

// GetArcAgentNodesStatusEnumValues Enumerates the set of values for ArcAgentNodesStatusEnum
func GetArcAgentNodesStatusEnumValues() []ArcAgentNodesStatusEnum {
	values := make([]ArcAgentNodesStatusEnum, 0)
	for _, v := range mappingArcAgentNodesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetArcAgentNodesStatusEnumStringValues Enumerates the set of values in String for ArcAgentNodesStatusEnum
func GetArcAgentNodesStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"UNKNOWN",
	}
}

// GetMappingArcAgentNodesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArcAgentNodesStatusEnum(val string) (ArcAgentNodesStatusEnum, bool) {
	enum, ok := mappingArcAgentNodesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
