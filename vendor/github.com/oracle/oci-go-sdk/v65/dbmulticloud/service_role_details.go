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
// <br>
// <b>AWS</b>:<br>
// <b>Oracle AWS Connector Resource:</b>&nbsp;&nbsp;The Oracle AWS Connector Resource is used to install the AWS Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Google AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceRoleDetails Service Role details object.
type ServiceRoleDetails struct {

	// Amazon resource name AWSof the IAM role.
	RoleArn *string `mandatory:"true" json:"roleArn"`

	// Private endpoint of the AWS service.
	ServicePrivateEndpoint *string `mandatory:"true" json:"servicePrivateEndpoint"`

	// Type of service.
	ServiceType ServiceRoleDetailsServiceTypeEnum `mandatory:"true" json:"serviceType"`

	// Assume role  status.
	AssumeRoleStatus ServiceRoleDetailsAssumeRoleStatusEnum `mandatory:"false" json:"assumeRoleStatus,omitempty"`

	// List of all VMs where AWS Identity Connector is configured for Oracle DB Cloud VM Cluster.
	AwsNodes []AwsNodes `mandatory:"false" json:"awsNodes"`
}

func (m ServiceRoleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceRoleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceRoleDetailsServiceTypeEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetServiceRoleDetailsServiceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingServiceRoleDetailsAssumeRoleStatusEnum(string(m.AssumeRoleStatus)); !ok && m.AssumeRoleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssumeRoleStatus: %s. Supported values are: %s.", m.AssumeRoleStatus, strings.Join(GetServiceRoleDetailsAssumeRoleStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceRoleDetailsServiceTypeEnum Enum with underlying type: string
type ServiceRoleDetailsServiceTypeEnum string

// Set of constants representing the allowable values for ServiceRoleDetailsServiceTypeEnum
const (
	ServiceRoleDetailsServiceTypeKms ServiceRoleDetailsServiceTypeEnum = "KMS"
)

var mappingServiceRoleDetailsServiceTypeEnum = map[string]ServiceRoleDetailsServiceTypeEnum{
	"KMS": ServiceRoleDetailsServiceTypeKms,
}

var mappingServiceRoleDetailsServiceTypeEnumLowerCase = map[string]ServiceRoleDetailsServiceTypeEnum{
	"kms": ServiceRoleDetailsServiceTypeKms,
}

// GetServiceRoleDetailsServiceTypeEnumValues Enumerates the set of values for ServiceRoleDetailsServiceTypeEnum
func GetServiceRoleDetailsServiceTypeEnumValues() []ServiceRoleDetailsServiceTypeEnum {
	values := make([]ServiceRoleDetailsServiceTypeEnum, 0)
	for _, v := range mappingServiceRoleDetailsServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceRoleDetailsServiceTypeEnumStringValues Enumerates the set of values in String for ServiceRoleDetailsServiceTypeEnum
func GetServiceRoleDetailsServiceTypeEnumStringValues() []string {
	return []string{
		"KMS",
	}
}

// GetMappingServiceRoleDetailsServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceRoleDetailsServiceTypeEnum(val string) (ServiceRoleDetailsServiceTypeEnum, bool) {
	enum, ok := mappingServiceRoleDetailsServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ServiceRoleDetailsAssumeRoleStatusEnum Enum with underlying type: string
type ServiceRoleDetailsAssumeRoleStatusEnum string

// Set of constants representing the allowable values for ServiceRoleDetailsAssumeRoleStatusEnum
const (
	ServiceRoleDetailsAssumeRoleStatusConnected          ServiceRoleDetailsAssumeRoleStatusEnum = "CONNECTED"
	ServiceRoleDetailsAssumeRoleStatusDisconnected       ServiceRoleDetailsAssumeRoleStatusEnum = "DISCONNECTED"
	ServiceRoleDetailsAssumeRoleStatusPartiallyConnected ServiceRoleDetailsAssumeRoleStatusEnum = "PARTIALLY_CONNECTED"
	ServiceRoleDetailsAssumeRoleStatusUnknown            ServiceRoleDetailsAssumeRoleStatusEnum = "UNKNOWN"
)

var mappingServiceRoleDetailsAssumeRoleStatusEnum = map[string]ServiceRoleDetailsAssumeRoleStatusEnum{
	"CONNECTED":           ServiceRoleDetailsAssumeRoleStatusConnected,
	"DISCONNECTED":        ServiceRoleDetailsAssumeRoleStatusDisconnected,
	"PARTIALLY_CONNECTED": ServiceRoleDetailsAssumeRoleStatusPartiallyConnected,
	"UNKNOWN":             ServiceRoleDetailsAssumeRoleStatusUnknown,
}

var mappingServiceRoleDetailsAssumeRoleStatusEnumLowerCase = map[string]ServiceRoleDetailsAssumeRoleStatusEnum{
	"connected":           ServiceRoleDetailsAssumeRoleStatusConnected,
	"disconnected":        ServiceRoleDetailsAssumeRoleStatusDisconnected,
	"partially_connected": ServiceRoleDetailsAssumeRoleStatusPartiallyConnected,
	"unknown":             ServiceRoleDetailsAssumeRoleStatusUnknown,
}

// GetServiceRoleDetailsAssumeRoleStatusEnumValues Enumerates the set of values for ServiceRoleDetailsAssumeRoleStatusEnum
func GetServiceRoleDetailsAssumeRoleStatusEnumValues() []ServiceRoleDetailsAssumeRoleStatusEnum {
	values := make([]ServiceRoleDetailsAssumeRoleStatusEnum, 0)
	for _, v := range mappingServiceRoleDetailsAssumeRoleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceRoleDetailsAssumeRoleStatusEnumStringValues Enumerates the set of values in String for ServiceRoleDetailsAssumeRoleStatusEnum
func GetServiceRoleDetailsAssumeRoleStatusEnumStringValues() []string {
	return []string{
		"CONNECTED",
		"DISCONNECTED",
		"PARTIALLY_CONNECTED",
		"UNKNOWN",
	}
}

// GetMappingServiceRoleDetailsAssumeRoleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceRoleDetailsAssumeRoleStatusEnum(val string) (ServiceRoleDetailsAssumeRoleStatusEnum, bool) {
	enum, ok := mappingServiceRoleDetailsAssumeRoleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
