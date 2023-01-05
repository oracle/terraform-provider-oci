---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connection"
sidebar_current: "docs-oci-resource-golden_gate-connection"
description: |-
  Provides the Connection resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_connection
This resource provides the Connection resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new Connection.


## Example Usage

```hcl
resource "oci_golden_gate_connection" "test_connection" {
	#Required
	compartment_id = var.compartment_id
	connection_type = var.connection_connection_type
	display_name = var.connection_display_name
	technology_type = var.connection_technology_type

	#Optional
	account_key = var.connection_account_key
	account_name = var.connection_account_name
	additional_attributes {

		#Optional
		name = var.connection_additional_attributes_name
		value = var.connection_additional_attributes_value
	}
	authentication_type = var.connection_authentication_type
	azure_tenant_id = oci_golden_gate_azure_tenant.test_azure_tenant.id
	bootstrap_servers {

		#Optional
		host = var.connection_bootstrap_servers_host
		port = var.connection_bootstrap_servers_port
		private_ip = var.connection_bootstrap_servers_private_ip
	}
	client_id = oci_golden_gate_client.test_client.id
	client_secret = var.connection_client_secret
	connection_string = var.connection_connection_string
	consumer_properties = var.connection_consumer_properties
	database_id = oci_database_database.test_database.id
	database_name = oci_database_database.test_database.name
	db_system_id = oci_database_db_system.test_db_system.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	description = var.connection_description
	endpoint = var.connection_endpoint
	freeform_tags = {"bar-key"= "value"}
	host = var.connection_host
	key_id = oci_kms_key.test_key.id
	key_store = var.connection_key_store
	key_store_password = var.connection_key_store_password
	nsg_ids = var.connection_nsg_ids
	password = var.connection_password
	port = var.connection_port
	private_ip = var.connection_private_ip
	private_key_file = var.connection_private_key_file
	private_key_passphrase = var.connection_private_key_passphrase
	producer_properties = var.connection_producer_properties
	public_key_fingerprint = var.connection_public_key_fingerprint
	region = var.connection_region
	sas_token = var.connection_sas_token
	security_protocol = var.connection_security_protocol
	session_mode = var.connection_session_mode
	ssl_ca = var.connection_ssl_ca
	ssl_cert = var.connection_ssl_cert
	ssl_crl = var.connection_ssl_crl
	ssl_key = var.connection_ssl_key
	ssl_key_password = var.connection_ssl_key_password
	ssl_mode = var.connection_ssl_mode
	stream_pool_id = oci_streaming_stream_pool.test_stream_pool.id
	subnet_id = oci_core_subnet.test_subnet.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
	trust_store = var.connection_trust_store
	trust_store_password = var.connection_trust_store_password
	url = var.connection_url
	user_id = oci_identity_user.test_user.id
	username = var.connection_username
	vault_id = oci_kms_vault.test_vault.id
	wallet = var.connection_wallet
}
```

## Argument Reference

The following arguments are supported:

* `account_key` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'. e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ== 
* `account_name` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name. 
* `additional_attributes` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) An array of name-value pair attribute entries. Used as additional parameters in connection string. 
	* `name` - (Required when connection_type=MYSQL | POSTGRESQL) (Updatable) The name of the property entry. 
	* `value` - (Required when connection_type=MYSQL | POSTGRESQL) (Updatable) The value of the property entry. 
* `authentication_type` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | KAFKA_SCHEMA_REGISTRY) (Updatable) Used authentication mechanism to access Azure Data Lake Storage. 
* `azure_tenant_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - (Applicable when connection_type=KAFKA) (Updatable) Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - (Required when connection_type=KAFKA) (Updatable) The name or address of a host. 
	* `port` - (Applicable when connection_type=KAFKA) (Updatable) The port of an endpoint usually specified for a connection. 
	* `private_ip` - (Applicable when connection_type=KAFKA) (Updatable) The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `client_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d 
* `client_secret` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure client secret (aka application password) for authentication. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | ORACLE) (Updatable) JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;' 
* `connection_type` - (Required) (Updatable) The connection type. 
* `consumer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the consumer.properties file. 
* `database_id` - (Applicable when connection_type=ORACLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `database_name` - (Required when connection_type=MYSQL | POSTGRESQL) (Updatable) The name of the database.
* `db_system_id` - (Applicable when connection_type=MYSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - (Applicable when connection_type=GOLDENGATE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name.
* `endpoint` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure Storage service endpoint. e.g: https://test.blob.core.windows.net 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `host` - (Required when connection_type=GOLDENGATE | MYSQL | POSTGRESQL) (Updatable) The name or address of a host. 
* `key_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer "Master" key being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this key to manage secrets. 
* `key_store` - (Applicable when connection_type=KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The base64 encoded content of the KeyStore file. 
* `key_store_password` - (Applicable when connection_type=KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The KeyStore password. 
* `nsg_ids` - (Optional) (Updatable) An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `password` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | KAFKA | KAFKA_SCHEMA_REGISTRY | MYSQL | ORACLE | POSTGRESQL) (Updatable) The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, and so on. 
* `port` - (Required when connection_type=GOLDENGATE | MYSQL | POSTGRESQL) (Updatable) The port of an endpoint usually specified for a connection. 
* `private_ip` - (Applicable when connection_type=GOLDENGATE | KAFKA_SCHEMA_REGISTRY | MYSQL | ORACLE | POSTGRESQL) (Updatable) The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `private_key_file` - (Required when connection_type=OCI_OBJECT_STORAGE) (Updatable) The base64 encoded content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
* `private_key_passphrase` - (Applicable when connection_type=OCI_OBJECT_STORAGE) (Updatable) The passphrase of the private key. 
* `producer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the producer.properties file. 
* `public_key_fingerprint` - (Required when connection_type=OCI_OBJECT_STORAGE) (Updatable) The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
* `region` - (Applicable when connection_type=OCI_OBJECT_STORAGE) (Updatable) The name of the region. e.g.: us-ashburn-1 
* `sas_token` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Credential that uses a shared access signature (SAS) to authenticate to an Azure Service. This property is required when 'authenticationType' is set to 'SHARED_ACCESS_SIGNATURE'. e.g.: ?sv=2020-06-08&ss=bfqt&srt=sco&sp=rwdlacupyx&se=2020-09-10T20:27:28Z&st=2022-08-05T12:27:28Z&spr=https&sig=C1IgHsiLBmTSStYkXXGLTP8it0xBrArcgCqOsZbXwIQ%3D 
* `security_protocol` - (Required when connection_type=KAFKA | MYSQL | POSTGRESQL) (Updatable) Security protocol for PostgreSQL.
* `session_mode` - (Applicable when connection_type=ORACLE) (Updatable) The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `ssl_ca` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Database Certificate - The base64 encoded content of mysql.pem file containing the server public key (for 1 and 2-way SSL). 
* `ssl_cert` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Certificate - The base64 encoded content of client-cert.pem file  containing the client public key (for 2-way SSL). 
* `ssl_crl` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Certificates revoked by certificate authorities (CA). Server certificate must not be on this list (for 1 and 2-way SSL). Note: This is an optional and that too only applicable if TLS/MTLS option is selected. 
* `ssl_key` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Key - The client-key.pem containing the client private key (for 2-way SSL). 
* `ssl_key_password` - (Applicable when connection_type=KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The password for the cert inside of the KeyStore. In case it differs from the KeyStore password, it should be provided. 
* `ssl_mode` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) SSL modes for PostgreSQL.
* `stream_pool_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `technology_type` - (Required) The Kafka (e.g. Confluent) Schema Registry technology type. 
* `tenancy_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `trust_store` - (Applicable when connection_type=KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The base64 encoded content of the TrustStore file. 
* `trust_store_password` - (Applicable when connection_type=KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The TrustStore password. 
* `url` - (Required when connection_type=KAFKA_SCHEMA_REGISTRY) (Updatable) Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081' 
* `user_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Object Storage. The user must have write access to the bucket they want to connect to. 
* `username` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | KAFKA | KAFKA_SCHEMA_REGISTRY | MYSQL | ORACLE | POSTGRESQL) (Updatable) The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on.
* `vault_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer vault being referenced. If provided, this will reference a vault which the customer will be required to ensure the policies are established to permit the GoldenGate Service to manage secrets contained within this vault. 
* `wallet` - (Applicable when connection_type=ORACLE) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `account_name` - Sets the Azure storage account name. 
* `additional_attributes` - An array of name-value pair attribute entries. Used as additional parameters in connection string. 
	* `name` - The name of the property entry. 
	* `value` - The value of the property entry. 
* `authentication_type` - Used authentication mechanism to access Azure Data Lake Storage. 
* `azure_tenant_id` - Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - The name or address of a host. 
	* `port` - The port of an endpoint usually specified for a connection. 
	* `private_ip` - The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `client_id` - Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
* `connection_type` - The connection type. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `database_name` - The name of the database. 
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name.
* `endpoint` - Azure Storage service endpoint. e.g: https://test.blob.core.windows.net
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `host` - The name or address of a host. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `ingress_ips` - List of ingress IP addresses, from where the GoldenGate deployment connects to this connection's privateIp. 
	* `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet. 
* `key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer "Master" key being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this key to manage secrets. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `port` - The port of an endpoint usually specified for a connection. 
* `private_ip` - The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `region` - The name of the region. e.g.: us-ashburn-1
* `security_protocol` - Security protocol for PostgreSQL.
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `ssl_mode` - SSL mode for PostgreSQL.
* `state` - Possible lifecycle states for connection. 
* `stream_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}`
* `technology_type` - The Kafka (e.g. Confluent) Schema Registry technology type. 
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `url` - Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081'
* `user_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Object Storage. The user must have write access to the bucket they want to connect to. 
* `username` - The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on. 
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer vault being referenced. If provided, this will reference a vault which the customer will be required to ensure the policies are established to permit the GoldenGate Service to manage secrets contained within this vault. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Connection
	* `update` - (Defaults to 20 minutes), when updating the Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Connection


## Import

Connections can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_connection.test_connection "id"
```

