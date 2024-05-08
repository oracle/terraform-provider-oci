// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateConnectionDetails The information about a new Connection.
type CreateConnectionDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []AddResourceLockDetails

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	GetVaultId() *string

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	GetKeyId() *string

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	GetNsgIds() []string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
	GetSubnetId() *string

	// Controls the network traffic direction to the target:
	// SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.
	// SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet.
	// DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
	GetRoutingMethod() RoutingMethodEnum
}

type createconnectiondetails struct {
	JsonData       []byte
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Locks          []AddResourceLockDetails          `mandatory:"false" json:"locks"`
	VaultId        *string                           `mandatory:"false" json:"vaultId"`
	KeyId          *string                           `mandatory:"false" json:"keyId"`
	NsgIds         []string                          `mandatory:"false" json:"nsgIds"`
	SubnetId       *string                           `mandatory:"false" json:"subnetId"`
	RoutingMethod  RoutingMethodEnum                 `mandatory:"false" json:"routingMethod,omitempty"`
	DisplayName    *string                           `mandatory:"true" json:"displayName"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	ConnectionType string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *createconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconnectiondetails createconnectiondetails
	s := struct {
		Model Unmarshalercreateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Locks = s.Model.Locks
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.NsgIds = s.Model.NsgIds
	m.SubnetId = s.Model.SubnetId
	m.RoutingMethod = s.Model.RoutingMethod
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "POSTGRESQL":
		mm := CreatePostgresqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA_SCHEMA_REGISTRY":
		mm := CreateKafkaSchemaRegistryConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MICROSOFT_SQLSERVER":
		mm := CreateMicrosoftSqlserverConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JAVA_MESSAGE_SERVICE":
		mm := CreateJavaMessageServiceConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOOGLE_BIGQUERY":
		mm := CreateGoogleBigQueryConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_KINESIS":
		mm := CreateAmazonKinesisConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SNOWFLAKE":
		mm := CreateSnowflakeConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := CreateAzureDataLakeStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MONGODB":
		mm := CreateMongoDbConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3":
		mm := CreateAmazonS3ConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS":
		mm := CreateHdfsConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OBJECT_STORAGE":
		mm := CreateOciObjectStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB2":
		mm := CreateDb2ConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ELASTICSEARCH":
		mm := CreateElasticsearchConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_SYNAPSE_ANALYTICS":
		mm := CreateAzureSynapseConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REDIS":
		mm := CreateRedisConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := CreateMysqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC":
		mm := CreateGenericConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOOGLE_CLOUD_STORAGE":
		mm := CreateGoogleCloudStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA":
		mm := CreateKafkaConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := CreateOracleConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOLDENGATE":
		mm := CreateGoldenGateConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_REDSHIFT":
		mm := CreateAmazonRedshiftConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_NOSQL":
		mm := CreateOracleNosqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateConnectionDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m createconnectiondetails) GetLocks() []AddResourceLockDetails {
	return m.Locks
}

// GetVaultId returns VaultId
func (m createconnectiondetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m createconnectiondetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m createconnectiondetails) GetNsgIds() []string {
	return m.NsgIds
}

// GetSubnetId returns SubnetId
func (m createconnectiondetails) GetSubnetId() *string {
	return m.SubnetId
}

// GetRoutingMethod returns RoutingMethod
func (m createconnectiondetails) GetRoutingMethod() RoutingMethodEnum {
	return m.RoutingMethod
}

// GetDisplayName returns DisplayName
func (m createconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createconnectiondetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoutingMethodEnum(string(m.RoutingMethod)); !ok && m.RoutingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RoutingMethod: %s. Supported values are: %s.", m.RoutingMethod, strings.Join(GetRoutingMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
