// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabricksConnection Represents the metadata of a Databricks Connection.
type DatabricksConnection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Connection URL.
	// e.g.: 'jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb'
	ConnectionUrl *string `mandatory:"true" json:"connectionUrl"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.
	// Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	IngressIps []IngressIpDetails `mandatory:"false" json:"ingressIps"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Indicates that sensitive attributes are provided via Secrets.
	DoesUseSecretIds *bool `mandatory:"false" json:"doesUseSecretIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The OCID(/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource.
	// Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud
	// subscription id is provided. Otherwise the cluster placement group must not be provided.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// Security attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// OAuth client id, only applicable for authenticationType == OAUTH_M2M
	ClientId *string `mandatory:"false" json:"clientId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
	// Only applicable for authenticationType == OAUTH_M2M.
	// Note: When provided, 'clientSecret' field must not be provided.
	ClientSecretSecretId *string `mandatory:"false" json:"clientSecretSecretId"`

	// Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS.
	StorageCredentialName *string `mandatory:"false" json:"storageCredentialName"`

	// The Databricks technology type.
	TechnologyType DatabricksConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Used authentication mechanism to access Databricks.
	// Required fields by authentication types:
	// - PERSONAL_ACCESS_TOKEN: username is always 'token', user must enter password
	// - OAUTH_M2M: user must enter clientId and clientSecret
	AuthenticationType DatabricksConnectionAuthenticationTypeEnum `mandatory:"true" json:"authenticationType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`
}

// GetId returns Id
func (m DatabricksConnection) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m DatabricksConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DatabricksConnection) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m DatabricksConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m DatabricksConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DatabricksConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DatabricksConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m DatabricksConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabricksConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabricksConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabricksConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLocks returns Locks
func (m DatabricksConnection) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVaultId returns VaultId
func (m DatabricksConnection) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m DatabricksConnection) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m DatabricksConnection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m DatabricksConnection) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m DatabricksConnection) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m DatabricksConnection) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m DatabricksConnection) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

// GetSubscriptionId returns SubscriptionId
func (m DatabricksConnection) GetSubscriptionId() *string {
	return m.SubscriptionId
}

// GetClusterPlacementGroupId returns ClusterPlacementGroupId
func (m DatabricksConnection) GetClusterPlacementGroupId() *string {
	return m.ClusterPlacementGroupId
}

// GetSecurityAttributes returns SecurityAttributes
func (m DatabricksConnection) GetSecurityAttributes() map[string]map[string]interface{} {
	return m.SecurityAttributes
}

func (m DatabricksConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabricksConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabricksConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetDatabricksConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabricksConnectionAuthenticationTypeEnum(string(m.AuthenticationType)); !ok && m.AuthenticationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticationType: %s. Supported values are: %s.", m.AuthenticationType, strings.Join(GetDatabricksConnectionAuthenticationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabricksConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabricksConnection DatabricksConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeDatabricksConnection
	}{
		"DATABRICKS",
		(MarshalTypeDatabricksConnection)(m),
	}

	return json.Marshal(&s)
}

// DatabricksConnectionTechnologyTypeEnum Enum with underlying type: string
type DatabricksConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for DatabricksConnectionTechnologyTypeEnum
const (
	DatabricksConnectionTechnologyTypeDatabricks DatabricksConnectionTechnologyTypeEnum = "DATABRICKS"
)

var mappingDatabricksConnectionTechnologyTypeEnum = map[string]DatabricksConnectionTechnologyTypeEnum{
	"DATABRICKS": DatabricksConnectionTechnologyTypeDatabricks,
}

var mappingDatabricksConnectionTechnologyTypeEnumLowerCase = map[string]DatabricksConnectionTechnologyTypeEnum{
	"databricks": DatabricksConnectionTechnologyTypeDatabricks,
}

// GetDatabricksConnectionTechnologyTypeEnumValues Enumerates the set of values for DatabricksConnectionTechnologyTypeEnum
func GetDatabricksConnectionTechnologyTypeEnumValues() []DatabricksConnectionTechnologyTypeEnum {
	values := make([]DatabricksConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingDatabricksConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabricksConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for DatabricksConnectionTechnologyTypeEnum
func GetDatabricksConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"DATABRICKS",
	}
}

// GetMappingDatabricksConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabricksConnectionTechnologyTypeEnum(val string) (DatabricksConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingDatabricksConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabricksConnectionAuthenticationTypeEnum Enum with underlying type: string
type DatabricksConnectionAuthenticationTypeEnum string

// Set of constants representing the allowable values for DatabricksConnectionAuthenticationTypeEnum
const (
	DatabricksConnectionAuthenticationTypePersonalAccessToken DatabricksConnectionAuthenticationTypeEnum = "PERSONAL_ACCESS_TOKEN"
	DatabricksConnectionAuthenticationTypeOauthM2m            DatabricksConnectionAuthenticationTypeEnum = "OAUTH_M2M"
)

var mappingDatabricksConnectionAuthenticationTypeEnum = map[string]DatabricksConnectionAuthenticationTypeEnum{
	"PERSONAL_ACCESS_TOKEN": DatabricksConnectionAuthenticationTypePersonalAccessToken,
	"OAUTH_M2M":             DatabricksConnectionAuthenticationTypeOauthM2m,
}

var mappingDatabricksConnectionAuthenticationTypeEnumLowerCase = map[string]DatabricksConnectionAuthenticationTypeEnum{
	"personal_access_token": DatabricksConnectionAuthenticationTypePersonalAccessToken,
	"oauth_m2m":             DatabricksConnectionAuthenticationTypeOauthM2m,
}

// GetDatabricksConnectionAuthenticationTypeEnumValues Enumerates the set of values for DatabricksConnectionAuthenticationTypeEnum
func GetDatabricksConnectionAuthenticationTypeEnumValues() []DatabricksConnectionAuthenticationTypeEnum {
	values := make([]DatabricksConnectionAuthenticationTypeEnum, 0)
	for _, v := range mappingDatabricksConnectionAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabricksConnectionAuthenticationTypeEnumStringValues Enumerates the set of values in String for DatabricksConnectionAuthenticationTypeEnum
func GetDatabricksConnectionAuthenticationTypeEnumStringValues() []string {
	return []string{
		"PERSONAL_ACCESS_TOKEN",
		"OAUTH_M2M",
	}
}

// GetMappingDatabricksConnectionAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabricksConnectionAuthenticationTypeEnum(val string) (DatabricksConnectionAuthenticationTypeEnum, bool) {
	enum, ok := mappingDatabricksConnectionAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
