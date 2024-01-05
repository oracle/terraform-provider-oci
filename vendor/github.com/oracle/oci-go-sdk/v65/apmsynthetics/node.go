// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Node Details of the network node.
type Node struct {

	// ID of the network node.
	Id *string `mandatory:"true" json:"id"`

	// IP address of the network node.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Display name of the network node.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Geographical information of the network node.
	GeoInfo *string `mandatory:"false" json:"geoInfo"`

	// Outgoing links from the network node.
	OutgoingLinks []string `mandatory:"false" json:"outgoingLinks"`

	// Number of consecutive anonymous network nodes.
	ConsecutiveAnonymousCount *int `mandatory:"false" json:"consecutiveAnonymousCount"`

	// Level of the network node.
	Level *int `mandatory:"false" json:"level"`

	// Average packet response time in milliseconds.
	AvgPacketResponseTimeInMs *float64 `mandatory:"false" json:"avgPacketResponseTimeInMs"`

	// Percentage of the average packet loss.
	AvgPacketLossPercent *float64 `mandatory:"false" json:"avgPacketLossPercent"`

	// Type of network node.
	Type NodeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m Node) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Node) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetNodeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodeTypeEnum Enum with underlying type: string
type NodeTypeEnum string

// Set of constants representing the allowable values for NodeTypeEnum
const (
	NodeTypeSource      NodeTypeEnum = "SOURCE"
	NodeTypeDestination NodeTypeEnum = "DESTINATION"
	NodeTypeAnonymous   NodeTypeEnum = "ANONYMOUS"
	NodeTypeInternal    NodeTypeEnum = "INTERNAL"
	NodeTypeDangling    NodeTypeEnum = "DANGLING"
)

var mappingNodeTypeEnum = map[string]NodeTypeEnum{
	"SOURCE":      NodeTypeSource,
	"DESTINATION": NodeTypeDestination,
	"ANONYMOUS":   NodeTypeAnonymous,
	"INTERNAL":    NodeTypeInternal,
	"DANGLING":    NodeTypeDangling,
}

var mappingNodeTypeEnumLowerCase = map[string]NodeTypeEnum{
	"source":      NodeTypeSource,
	"destination": NodeTypeDestination,
	"anonymous":   NodeTypeAnonymous,
	"internal":    NodeTypeInternal,
	"dangling":    NodeTypeDangling,
}

// GetNodeTypeEnumValues Enumerates the set of values for NodeTypeEnum
func GetNodeTypeEnumValues() []NodeTypeEnum {
	values := make([]NodeTypeEnum, 0)
	for _, v := range mappingNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeTypeEnumStringValues Enumerates the set of values in String for NodeTypeEnum
func GetNodeTypeEnumStringValues() []string {
	return []string{
		"SOURCE",
		"DESTINATION",
		"ANONYMOUS",
		"INTERNAL",
		"DANGLING",
	}
}

// GetMappingNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeTypeEnum(val string) (NodeTypeEnum, bool) {
	enum, ok := mappingNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
