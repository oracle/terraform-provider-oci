// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data Plane Integration
//
// <b>Microsoft Azure:</b> <br>
// <b>Oracle Azure Connector Resource:</b>:&nbsp;&nbsp;The Oracle Azure Connector Resource is used to install the Azure Arc Server on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
//  The supported method to install the Azure Arc Server (Azure Identity) on the Exadata VM cluster:
// <ul>
//  <li>Using a Bearer Access Token</li>
// </ul>
// <b>Oracle Azure Blob Container Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Container Resource is used to capture the details of an Azure Blob Container.
// This resource can then be reused across multiple Exadata VM clusters in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) to mount the Azure container.
// <b>Oracle Azure Blob Mount Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Mount Resource is used to mount an Azure Blob Container on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// It relies on both the Oracle Azure Connector and the Oracle Azure Blob Container Resource to perform the mount operation.
// <b>Discover Azure Vaults and Keys Resource:</b>&nbsp;&nbsp;The Discover Oracle Azure Vaults and Azure Keys Resource is used to discover Azure Vaults and the associated encryption keys available in your Azure project.
// <b>Oracle Azure Vault:</b>&nbsp;&nbsp;The Oracle Azure Vault Resource is used to manage Azure Vaults within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Oracle Azure Key:</b>&nbsp;&nbsp;Oracle Azure Key Resource is used to register and manage a Oracle Azure Key Key within Oracle Cloud Infrastructure (OCI) under an associated Azure Vault.
// <br>
// <b>Google Cloud:</b><br>
// <b>Oracle Google Cloud Connector Resource:</b>&nbsp;&nbsp;The Oracle Google Cloud Connector Resource is used to install the Google Cloud Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Discover Google Key Rings and Keys Resource:</b>&nbsp;&nbsp;The Discover Google Key Rings and Keys Resource is used to discover Google Cloud Key Rings and the associated encryption keys available in your Google Cloud project.
// <b>Google Key Rings Resource:</b>&nbsp;&nbsp;The Google Key Rings Resource is used to register and manage Google Cloud Key Rings within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Google Key Resource:</b>&nbsp;&nbsp;The Google Key Resource is used to register and manage a Google Cloud Key within Oracle Cloud Infrastructure (OCI) under an associated Google Key Ring.
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArcAgentNodes Azure Arc Agent node details.
type ArcAgentNodes struct {

	// Host name or Azure Arc Agent name.
	HostName *string `mandatory:"false" json:"hostName"`

	// Host ID.
	HostId *string `mandatory:"false" json:"hostId"`

	// Current Arc Agent Version installed on this node of Oracle Cloud VM Cluster.
	CurrentArcAgentVersion *string `mandatory:"false" json:"currentArcAgentVersion"`

	// The current status of the Azure Arc Agent resource.
	Status ArcAgentNodesStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Time when the Azure Arc Agent's status was checked RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
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
