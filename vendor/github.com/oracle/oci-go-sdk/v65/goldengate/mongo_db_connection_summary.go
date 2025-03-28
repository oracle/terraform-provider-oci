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

// MongoDbConnectionSummary Summary of the MongoDB Connection.
type MongoDbConnectionSummary struct {

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

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Indicates that sensitive attributes are provided via Secrets.
	DoesUseSecretIds *bool `mandatory:"false" json:"doesUseSecretIds"`

	// MongoDB connection string.
	// e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The username Oracle GoldenGate uses to connect to the database.
	// This username must already exist and be available by the database to be connected to.
	Username *string `mandatory:"false" json:"username"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Autonomous Json Database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect the associated database.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
	// - The content of a .pem file containing the client private key (for 2-way SSL).
	// Note: When provided, 'tlsCertificateKeyFile' field must not be provided.
	TlsCertificateKeyFileSecretId *string `mandatory:"false" json:"tlsCertificateKeyFileSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file.
	// Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided.
	TlsCertificateKeyFilePasswordSecretId *string `mandatory:"false" json:"tlsCertificateKeyFilePasswordSecretId"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	RoutingMethod RoutingMethodEnum `mandatory:"false" json:"routingMethod,omitempty"`

	// The MongoDB technology type.
	TechnologyType MongoDbConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Security Type for MongoDB.
	SecurityProtocol MongoDbConnectionSecurityProtocolEnum `mandatory:"false" json:"securityProtocol,omitempty"`
}

// GetId returns Id
func (m MongoDbConnectionSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m MongoDbConnectionSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m MongoDbConnectionSummary) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m MongoDbConnectionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m MongoDbConnectionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MongoDbConnectionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m MongoDbConnectionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLifecycleState returns LifecycleState
func (m MongoDbConnectionSummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m MongoDbConnectionSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m MongoDbConnectionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m MongoDbConnectionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetVaultId returns VaultId
func (m MongoDbConnectionSummary) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m MongoDbConnectionSummary) GetKeyId() *string {
	return m.KeyId
}

// GetIngressIps returns IngressIps
func (m MongoDbConnectionSummary) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

// GetNsgIds returns NsgIds
func (m MongoDbConnectionSummary) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m MongoDbConnectionSummary) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m MongoDbConnectionSummary) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetLocks returns Locks
func (m MongoDbConnectionSummary) GetLocks() []ResourceLock {
	return m.Locks
}

// GetDoesUseSecretIds returns DoesUseSecretIds
func (m MongoDbConnectionSummary) GetDoesUseSecretIds() *bool {
	return m.DoesUseSecretIds
}

func (m MongoDbConnectionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MongoDbConnectionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMongoDbConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetMongoDbConnectionTechnologyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMongoDbConnectionSecurityProtocolEnum(string(m.SecurityProtocol)); !ok && m.SecurityProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityProtocol: %s. Supported values are: %s.", m.SecurityProtocol, strings.Join(GetMongoDbConnectionSecurityProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MongoDbConnectionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMongoDbConnectionSummary MongoDbConnectionSummary
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeMongoDbConnectionSummary
	}{
		"MONGODB",
		(MarshalTypeMongoDbConnectionSummary)(m),
	}

	return json.Marshal(&s)
}
