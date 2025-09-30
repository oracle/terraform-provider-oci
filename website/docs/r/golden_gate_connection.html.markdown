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
	access_key_id = oci_kms_key.test_key.id
	account_key = var.connection_account_key
	account_key_secret_id = oci_vault_secret.test_secret.id
	account_name = var.connection_account_name
	additional_attributes {

		#Optional
		name = var.connection_additional_attributes_name
		value = var.connection_additional_attributes_value
	}
	authentication_mode = var.connection_authentication_mode
	authentication_type = var.connection_authentication_type
	azure_authority_host = var.connection_azure_authority_host
	azure_tenant_id = oci_golden_gate_azure_tenant.test_azure_tenant.id
	bootstrap_servers {

		#Optional
		host = var.connection_bootstrap_servers_host
		port = var.connection_bootstrap_servers_port
		private_ip = var.connection_bootstrap_servers_private_ip
	}
	catalog {
		#Required
		catalog_type = var.connection_catalog_catalog_type

		#Optional
		branch = var.connection_catalog_branch
		client_id = oci_golden_gate_client.test_client.id
		client_secret_secret_id = oci_vault_secret.test_secret.id
		glue_id = oci_golden_gate_glue.test_glue.id
		name = var.connection_catalog_name
		principal_role = var.connection_catalog_principal_role
		properties_secret_id = oci_vault_secret.test_secret.id
		uri = var.connection_catalog_uri
	}
	client_id = oci_golden_gate_client.test_client.id
	client_secret = var.connection_client_secret
	client_secret_secret_id = oci_vault_secret.test_secret.id
	cluster_id = oci_containerengine_cluster.test_cluster.id
	cluster_placement_group_id = oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id
	connection_factory = var.connection_connection_factory
	connection_string = var.connection_connection_string
	connection_url = var.connection_connection_url
	consumer_properties = var.connection_consumer_properties
	core_site_xml = var.connection_core_site_xml
	database_id = oci_database_database.test_database.id
	database_name = oci_database_database.test_database.name
	db_system_id = oci_database_db_system.test_db_system.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	description = var.connection_description
	does_use_secret_ids = var.connection_does_use_secret_ids
	endpoint = var.connection_endpoint
	fingerprint = var.connection_fingerprint
	freeform_tags = {"bar-key"= "value"}
	host = var.connection_host
	jndi_connection_factory = var.connection_jndi_connection_factory
	jndi_initial_context_factory = var.connection_jndi_initial_context_factory
	jndi_provider_url = var.connection_jndi_provider_url
	jndi_security_credentials = var.connection_jndi_security_credentials
	jndi_security_credentials_secret_id = oci_vault_secret.test_secret.id
	jndi_security_principal = var.connection_jndi_security_principal
	key_id = oci_kms_key.test_key.id
	key_store = var.connection_key_store
	key_store_password = var.connection_key_store_password
	key_store_secret_id = oci_vault_secret.test_secret.id
	key_store_password_secret_id = oci_vault_secret.test_secret.id
	locks {
		#Required
		type = var.connection_locks_type

		#Optional
		message = var.connection_locks_message
	}
	nsg_ids = var.connection_nsg_ids
	password = var.connection_password
	password_secret_id = oci_vault_secret.test_secret.id
	port = var.connection_port
	private_ip = var.connection_private_ip
	private_key_file = var.connection_private_key_file
	private_key_file_secret_id = oci_vault_secret.test_secret.id
	private_key_passphrase = var.connection_private_key_passphrase
	private_key_passphrase_secret_id = oci_vault_secret.test_secret.id
	producer_properties = var.connection_producer_properties
	public_key_fingerprint = var.connection_public_key_fingerprint
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
	region = var.connection_region
	routing_method = var.connection_routing_method
	sas_token = var.connection_sas_token
	sas_token_secret_id = oci_vault_secret.test_secret.id
	secret_access_key = var.connection_secret_access_key
	secret_access_key_secret_id = oci_vault_secret.test_secret.id
	security_attributes = var.connection_security_attributes
	security_protocol = var.connection_security_protocol
	servers = var.connection_servers
	service_account_key_file = var.connection_service_account_key_file
	service_account_key_file_secret_id = oci_vault_secret.test_secret.id
	session_mode = var.connection_session_mode
	should_use_jndi = var.connection_should_use_jndi
	should_use_resource_principal = var.connection_should_use_resource_principal
	should_validate_server_certificate = var.connection_should_validate_server_certificate
	ssl_ca = var.connection_ssl_ca
	ssl_cert = var.connection_ssl_cert
	ssl_client_keystash = var.connection_ssl_client_keystash
	ssl_client_keystash_secret_id = oci_vault_secret.test_secret.id
	ssl_client_keystoredb = var.connection_ssl_client_keystoredb
	ssl_client_keystoredb_secret_id = oci_vault_secret.test_secret.id
	ssl_crl = var.connection_ssl_crl
	ssl_key = var.connection_ssl_key
	ssl_key_password = var.connection_ssl_key_password
	ssl_key_password_secret_id = oci_vault_secret.test_secret.id
	ssl_key_secret_id = oci_vault_secret.test_secret.id
	ssl_mode = var.connection_ssl_mode
	ssl_server_certificate = var.connection_ssl_server_certificate
	storage {
		#Required
		storage_type = var.connection_storage_storage_type

		#Optional
		access_key_id = oci_kms_key.test_key.id
		account_key_secret_id = oci_vault_secret.test_secret.id
		account_name = var.connection_storage_account_name
		bucket = var.connection_storage_bucket
		container = var.connection_storage_container
		endpoint = var.connection_storage_endpoint
		project_id = oci_ai_anomaly_detection_project.test_project.id
		region = var.connection_storage_region
		scheme_type = var.connection_storage_scheme_type
		secret_access_key_secret_id = oci_vault_secret.test_secret.id
		service_account_key_file_secret_id = oci_vault_secret.test_secret.id
	}
	storage_credential_name = var.connection_storage_credential_name
	stream_pool_id = oci_streaming_stream_pool.test_stream_pool.id
	subnet_id = oci_core_subnet.test_subnet.id
	subscription_id = oci_onesubscription_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
	tenant_id = oci_golden_gate_tenant.test_tenant.id
	tls_ca_file = var.connection_tls_ca_file
	tls_certificate_key_file = var.connection_tls_certificate_key_file
	tls_certificate_key_file_password = var.connection_tls_certificate_key_file_password
	tls_certificate_key_file_password_secret_id = oci_vault_secret.test_secret.id
	tls_certificate_key_file_secret_id = oci_vault_secret.test_secret.id
	trust_store = var.connection_trust_store
	trust_store_password = var.connection_trust_store_password
	trust_store_password_secret_id = oci_vault_secret.test_secret.id
	trust_store_secret_id = oci_vault_secret.test_secret.id
	url = var.connection_url
	user_id = oci_identity_user.test_user.id
	username = var.connection_username
	vault_id = oci_kms_vault.test_vault.id
	wallet = var.connection_wallet
	wallet_secret_id = oci_vault_secret.test_secret.id
	trigger_refresh = true
}
```

## Argument Reference

The following arguments are supported:

* `access_key_id` - (Required when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) Access key ID to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" 
* `account_key` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'. e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ== Deprecated: This field is deprecated and replaced by "accountKeySecretId". This field will be removed after February 15 2026.
* `account_key_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. Note: When provided, 'accountKey' field must not be provided.
* `account_name` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name. 
* `additional_attributes` - (Applicable when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) An array of name-value pair attribute entries. Used as additional parameters in connection string. 
	* `name` - (Required when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name of the property entry. 
	* `value` - (Required when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The value of the property entry. 
* `authentication_mode` - (Applicable when connection_type=ORACLE) (Updatable) Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections, when a databaseId is provided. The default value is MTLS. 
* `authentication_type` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA_SCHEMA_REGISTRY | REDIS | SNOWFLAKE) (Updatable) Authentication type for Java Message Service.  If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required. 
* `azure_authority_host` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The endpoint used for authentication with Microsoft Entra ID (formerly Azure Active Directory). Default value: https://login.microsoftonline.com When connecting to a non-public Azure Cloud, the endpoint must be provided, eg:
	* Azure China: https://login.chinacloudapi.cn/
	* Azure US Government: https://login.microsoftonline.us/ 
* `azure_tenant_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f 
* `bootstrap_servers` - (Applicable when connection_type=KAFKA) (Updatable) Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - (Required when connection_type=KAFKA) (Updatable) The name or address of a host. 
	* `port` - (Applicable when connection_type=KAFKA) (Updatable) The port of an endpoint usually specified for a connection. 
	* `private_ip` - (Applicable when connection_type=KAFKA) (Updatable) Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.

		The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `catalog` - (Required when connection_type=ICEBERG) (Updatable) The information about a new catalog of given type used in an Iceberg connection. 
	* `branch` - (Required when catalog_type=NESSIE) (Updatable) The active branch of the Nessie catalog from which Iceberg reads and writes table metadata.
	* `catalog_type` - (Required) (Updatable) The catalog type. 
	* `client_id` - (Required when catalog_type=POLARIS) (Updatable) The OAuth client ID used for authentication.
	* `client_secret_secret_id` - (Required when catalog_type=POLARIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect to Snowflake platform. 
	* `glue_id` - (Required when catalog_type=GLUE) (Updatable) The AWS Glue Catalog ID where Iceberg tables are registered.
	* `name` - (Required when catalog_type=POLARIS) (Updatable) The catalog name within Polaris where Iceberg tables are registered.
	* `principal_role` - (Required when catalog_type=POLARIS) (Updatable) The Snowflake role used to access Polaris.
	* `properties_secret_id` - (Required when catalog_type=REST) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the configuration file containing additional properties for the REST catalog. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
	* `uri` - (Required when catalog_type=NESSIE | POLARIS | REST) (Updatable) The URL endpoint for the Polaris API. e.g.: 'https://<your-snowflake-account>.snowflakecomputing.com/polaris/api/catalog' 
* `client_id` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable) Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d 
* `client_secret` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable) Azure client secret (aka application password) for authentication. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1 Deprecated: This field is deprecated and replaced by "clientSecretSecretId". This field will be removed after February 15 2026. 
* `client_secret_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored. Only applicable for authenticationType == OAUTH_M2M. Note: When provided, 'clientSecret' field must not be provided. 
* `cluster_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kafka cluster being referenced from Oracle Cloud Infrastructure Streaming with Apache Kafka. 
* `cluster_placement_group_id` - (Optional) The OCID(https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource. Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud subscription id is provided. Otherwise the cluster placement group must not be provided. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA' 
* `connection_string` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | MONGODB | ORACLE) (Updatable) Connection string. AZURE_SYNAPSE_ANALYTICS e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;', MONGODB e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'. 
* `connection_type` - (Required) (Updatable) The connection type. 
* `connection_url` - (Required when connection_type=AMAZON_REDSHIFT | DATABRICKS | JAVA_MESSAGE_SERVICE | ORACLE_AI_DATA_PLATFORM | SNOWFLAKE) (Updatable) Connection URL. e.g.: 'jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb' 
* `consumer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the consumer.properties file. 
* `core_site_xml` - (Required when connection_type=HDFS) (Updatable) The base64 encoded content of the Hadoop Distributed File System configuration file (core-site.xml). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `database_id` - (Applicable when connection_type=MONGODB | ORACLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Autonomous Json Database. 
* `database_name` - (Required when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name of the database. 
* `db_system_id` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - (Applicable when connection_type=GOLDENGATE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name. 
* `does_use_secret_ids` - (Optional) (Updatable) Indicates that sensitive attributes are provided via Secrets. 
* `endpoint` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3 | AZURE_DATA_LAKE_STORAGE | MICROSOFT_FABRIC) (Updatable) The endpoint URL of the 3rd party cloud service. e.g.: 'https://kinesis.us-east-1.amazonaws.com' If not provided, GoldenGate will default to the default endpoint in the `region`. 
* `fingerprint` - (Applicable when connection_type=ELASTICSEARCH) (Updatable) Fingerprint required by TLS security protocol. E.g.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c' 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `host` - (Required when connection_type=DB2 |GENERIC | GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name or address of a host. In case of Generic connection type host and port separated by colon. Example: `"server.example.com:1234"`
	For multiple hosts, provide a comma separated list. Example: `"server1.example.com:1000,server1.example.com:2000"` 
* `jndi_connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory' 
* `jndi_initial_context_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The implementation of javax.naming.spi.InitialContextFactory interface that the client uses to obtain initial naming context. e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory' 
* `jndi_provider_url` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000' 
* `jndi_security_credentials` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The password associated to the principal. Deprecated: This field is deprecated and replaced by "jndiSecurityCredentialsSecretId". This field will be removed after February 15 2026. 
* `jndi_security_credentials_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the security credentials are stored associated to the principal. Note: When provided, 'jndiSecurityCredentials' field must not be provided. 
* `jndi_security_principal` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) Specifies the identity of the principal (user) to be authenticated. e.g.: 'admin2' 
* `key_id` - (Optional) (Updatable) Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `key_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the KeyStore file. Deprecated: This field is deprecated and replaced by "keyStoreSecretId". This field will be removed after February 15 2026. 
* `key_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The KeyStore password. Deprecated: This field is deprecated and replaced by "keyStorePasswordSecretId". This field will be removed after February 15 2026. 
* `key_store_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl KeyStore password is stored. Note: When provided, 'keyStorePassword' field must not be provided. 
* `key_store_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the KeyStore file is stored. Note: When provided, 'keyStore' field must not be provided.
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `type` - (Required) Type of the lock.
* `nsg_ids` - (Optional) (Updatable) An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `password` - (Applicable when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DATABRICKS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. Deprecated: This field is deprecated and replaced by "passwordSecretId". This field will be removed after February 15 2026. 
* `password_secret_id` - (Applicable when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DATABRICKS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. If secretId is used plaintext field must not be provided. Note: When provided, 'password' field must not be provided. 
* `port` - (Required when connection_type=DB2 | GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The port of an endpoint usually specified for a connection. 
* `private_ip` - (Applicable when connection_type=GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MYSQL | ORACLE | POSTGRESQL) (Updatable) Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.

	The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `private_key_file` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The base64 encoded content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Deprecated: This field is deprecated and replaced by "privateKeyFileSecretId". This field will be removed after February 15 2026. 
* `private_key_file_secret_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Note: When provided, 'privateKeyFile' field must not be provided. 
* `private_key_passphrase` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL | SNOWFLAKE) (Updatable) Password if the private key file is encrypted. Deprecated: This field is deprecated and replaced by "privateKeyPassphraseSecretId". This field will be removed after February 15 2026. 
* `private_key_passphrase_secret_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password for the private key file. Note: When provided, 'privateKeyPassphrase' field must not be provided. 
* `producer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the producer.properties file. 
* `public_key_fingerprint` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
* `redis_cluster_id` - (Applicable when connection_type=REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster. 
* `region` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3 | OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The name of the AWS region where the bucket is created. If not provided, GoldenGate will default to 'us-west-2'. Note: this property will become mandatory after May 20, 2026. 
* `routing_method` - (Optional) (Updatable) Controls the network traffic direction to the target: SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.  SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet. DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected. 
* `sas_token` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Credential that uses a shared access signature (SAS) to authenticate to an Azure Service. This property is required when 'authenticationType' is set to 'SHARED_ACCESS_SIGNATURE'. e.g.: ?sv=2020-06-08&ss=bfqt&srt=sco&sp=rwdlacupyx&se=2020-09-10T20:27:28Z&st=2022-08-05T12:27:28Z&spr=https&sig=C1IgHsiLBmTSStYkXXGLTP8it0xBrArcgCqOsZbXwIQ%3D Deprecated: This field is deprecated and replaced by "sasTokenSecretId". This field will be removed after February 15 2026. 
* `sas_token_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored. Note: When provided, 'sasToken' field must not be provided. 
* `secret_access_key` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) Secret access key to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" Deprecated: This field is deprecated and replaced by "secretAccessKeySecretId". This field will be removed after February 15 2026. 
* `secret_access_key_secret_id` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the secret access key is stored. Note: When provided, 'secretAccessKey' field must not be provided. 
* `security_attributes` - (Optional) (Updatable) Security attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `security_protocol` - (Required when connection_type=DB2 | ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA | MICROSOFT_SQLSERVER | MONGODB | MYSQL | POSTGRESQL | REDIS) (Updatable) Security protocol for Java Message Service. If not provided, default is PLAIN. Optional until 2024-06-27, in the release after it will be made required. 
* `servers` - (Required when connection_type=ELASTICSEARCH | REDIS) (Updatable) Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional.  If port is not specified, it defaults to 9200. Used for establishing the initial connection to the Elasticsearch cluster. Example: `"server1.example.com:4000,server2.example.com:4000"` 
* `service_account_key_file` - (Applicable when connection_type=GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE | GOOGLE_PUBSUB) (Updatable) The base64 encoded content of the service account key file containing the credentials required to use Google Cloud Storage. Deprecated: This field is deprecated and replaced by "serviceAccountKeyFileSecretId". This field will be removed after February 15 2026. 
* `service_account_key_file_secret_id` - (Applicable when connection_type=GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE | GOOGLE_PUBSUB) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage. Note: When provided, 'serviceAccountKeyFile' field must not be provided. 
* `session_mode` - (Applicable when connection_type=ORACLE) (Updatable) The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `should_use_jndi` - (Required when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) If set to true, Java Naming and Directory Interface (JNDI) properties should be provided. 
* `should_use_resource_principal` - (Applicable when connection_type=KAFKA | OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) Specifies that the user intends to authenticate to the instance using a resource principal. Applicable only for Oracle Cloud Infrastructure Streaming connections. Only available from 23.9.0.0.0 GoldenGate versions. Note: When specified, 'username'/'password'/'passwordSecretId' fields must not be provided. Default: false 
* `should_validate_server_certificate` - (Applicable when connection_type=MICROSOFT_SQLSERVER) (Updatable) If set to true, the driver validates the certificate that is sent by the database server. 
* `ssl_ca` - (Applicable when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The base64 encoded certificate of the trusted certificate authorities (Trusted CA) for PostgreSQL.  The supported file formats are .pem and .crt. It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_cert` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Certificate - The base64 encoded content of a .pem or .crt file containing the client public key (for 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_client_keystash` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded keystash file which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Deprecated: This field is deprecated and replaced by "sslClientKeystashSecretId". This field will be removed after February 15 2026. 
* `ssl_client_keystash_secret_id` - (Applicable when connection_type=DB2) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystash file is stored,  which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Note: When provided, 'sslClientKeystash' field must not be provided. 
* `ssl_client_keystoredb` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded keystore file created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Deprecated: This field is deprecated and replaced by "sslClientKeystoredbSecretId". This field will be removed after February 15 2026. 
* `ssl_client_keystoredb_secret_id` - (Applicable when connection_type=DB2) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystore file stored,  which created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Note: When provided, 'sslClientKeystoredb' field must not be provided. 
* `ssl_crl` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). Note: This is an optional property and only applicable if TLS/MTLS option is selected. It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_key` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Key - The base64 encoded content of a .pem or .crt file containing the client private key (for 2-way SSL). Deprecated: This field is deprecated and replaced by "sslKeySecretId". This field will be removed after February 15 2026. 
* `ssl_key_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The password for the cert inside of the KeyStore. In case it differs from the KeyStore password, it should be provided. Deprecated: This field is deprecated and replaced by "sslKeyPasswordSecretId". This field will be removed after February 15 2026. 
* `ssl_key_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored for the cert inside of the Keystore. In case it differs from the KeyStore password, it should be provided. Note: When provided, 'sslKeyPassword' field must not be provided. 
* `ssl_key_secret_id` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key
	* The content of a .pem or .crt file containing the client private key (for 2-way SSL). Note: When provided, 'sslKey' field must not be provided. 
* `ssl_mode` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) SSL modes for PostgreSQL.
* `ssl_server_certificate` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded file which contains the self-signed server certificate / Certificate Authority (CA) certificate. It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `storage` - (Required when connection_type=ICEBERG) (Updatable) The information about a new storage of given type used in an Iceberg connection. 
	* `access_key_id` - (Required when storage_type=AMAZON_S3) (Updatable) Access key ID to access the Amazon S3 bucket. 
	* `account_key_secret_id` - (Required when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. 
	* `account_name` - (Required when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name. 
	* `bucket` - (Required when storage_type=AMAZON_S3 | GOOGLE_CLOUD_STORAGE) (Updatable) Google Cloud Storage bucket where Iceberg stores metadata and data files.
	* `container` - (Required when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The Azure Blob Storage container where Iceberg tables are stored.
	* `endpoint` - (Applicable when storage_type=AMAZON_S3 | AZURE_DATA_LAKE_STORAGE) (Updatable) The Azure Blob Storage endpoint where Iceberg data is stored. e.g.: 'https://my-azure-storage-account.blob.core.windows.net' 
	* `project_id` - (Required when storage_type=GOOGLE_CLOUD_STORAGE) (Updatable) The Google Cloud Project where the bucket exists.
	* `region` - (Required when storage_type=AMAZON_S3) (Updatable) The AMAZON region where the S3 bucket is hosted. e.g.: 'us-east-2' 
	* `scheme_type` - (Required when storage_type=AMAZON_S3) (Updatable) The scheme of the storage. 
	* `secret_access_key_secret_id` - (Required when storage_type=AMAZON_S3) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored. 
	* `service_account_key_file_secret_id` - (Required when storage_type=GOOGLE_CLOUD_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage. 
	* `storage_type` - (Required) (Updatable) The storage type used in the Iceberg connection. 
* `storage_credential_name` - (Applicable when connection_type=DATABRICKS) (Updatable) Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS. 
* `stream_pool_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection. 
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `technology_type` - (Required) The Kafka (e.g. Confluent) Schema Registry technology type. 
* `tenancy_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `tenant_id` - (Required when connection_type=MICROSOFT_FABRIC) (Updatable) Azure tenant ID of the application. e.g.: 14593954-d337-4a61-a364-9f758c64f97f 
* `tls_ca_file` - (Applicable when connection_type=MONGODB) (Updatable) Database Certificate - The base64 encoded content of a .pem file, containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `tls_certificate_key_file` - (Applicable when connection_type=MONGODB) (Updatable) Client Certificate - The base64 encoded content of a .pem file, containing the client public key (for 2-way SSL). Deprecated: This field is deprecated and replaced by "tlsCertificateKeyFileSecretId". This field will be removed after February 15 2026. 
* `tls_certificate_key_file_password` - (Applicable when connection_type=MONGODB) (Updatable) Client Certificate key file password. Deprecated: This field is deprecated and replaced by "tlsCertificateKeyFilePasswordSecretId". This field will be removed after February 15 2026. 
* `tls_certificate_key_file_password_secret_id` - (Applicable when connection_type=MONGODB) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file. Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided. 
* `tls_certificate_key_file_secret_id` - (Applicable when connection_type=MONGODB) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
	* The content of a .pem file containing the client private key (for 2-way SSL). Note: When provided, 'tlsCertificateKeyFile' field must not be provided. 
* `trust_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the TrustStore file. Deprecated: This field is deprecated and replaced by "trustStoreSecretId". This field will be removed after February 15 2026. 
* `trust_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The TrustStore password. Deprecated: This field is deprecated and replaced by "trustStorePasswordSecretId". This field will be removed after February 15 2026. 
* `trust_store_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl TrustStore password is stored. Note: When provided, 'trustStorePassword' field must not be provided. 
* `trust_store_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the TrustStore file is stored. Note: When provided, 'trustStore' field must not be provided. 
* `url` - (Required when connection_type=KAFKA_SCHEMA_REGISTRY) (Updatable) Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081' 
* `user_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database. The user must have write access to the table they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint. 
* `username` - (Required when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it. 
* `vault_id` - (Optional) (Updatable) Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault. 
* `wallet` - (Applicable when connection_type=ORACLE) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database. This attribute is expected to be base64 encoded. Deprecated: This field is deprecated and replaced by "walletSecretId". This field will be removed after February 15 2026. 
* `wallet_secret_id` - (Applicable when connection_type=ORACLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the wallet file is stored.  The wallet contents Oracle GoldenGate uses to make connections to a database. Note: When provided, 'wallet' field must not be provided. 
* `trigger_refresh` - (Optional) (Updatable) If value is true, it triggers connection refresh action and this attribute change will always show up in the "update" plan and will apply steps in order to refresh secrets and dependent service properties (such as ADB connection strings, wallets, etc..).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_key_id` - Access key ID to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" 
* `account_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. Note: When provided, 'accountKey' field must not be provided. 
* `account_name` - Sets the Azure storage account name. 
* `additional_attributes` - An array of name-value pair attribute entries. Used as additional parameters in connection string. 
	* `name` - The name of the property entry. 
	* `value` - The value of the property entry. 
* `authentication_mode` - Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections, when a databaseId is provided. The default value is MTLS. 
* `authentication_type` - Used authentication mechanism to access Databricks. Required fields by authentication types:
	* PERSONAL_ACCESS_TOKEN: username is always 'token', user must enter password
	* OAUTH_M2M: user must enter clientId and clientSecret 
* `azure_authority_host` - The endpoint used for authentication with Microsoft Entra ID (formerly Azure Active Directory). Default value: https://login.microsoftonline.com When connecting to a non-public Azure Cloud, the endpoint must be provided, eg:
	* Azure China: https://login.chinacloudapi.cn/
	* Azure US Government: https://login.microsoftonline.us/ 
* `azure_tenant_id` - Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f 
* `bootstrap_servers` - Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - The name or address of a host. 
	* `port` - The port of an endpoint usually specified for a connection. 
	* `private_ip` - Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.

		The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `catalog` - Represents the catalog of given type used in an Iceberg connection. 
	* `branch` - The active branch of the Nessie catalog from which Iceberg reads and writes table metadata.
	* `catalog_type` - The catalog type. 
	* `client_id` - The OAuth client ID used for authentication.
	* `client_secret_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect to Snowflake platform. 
	* `glue_id` - The AWS Glue Catalog ID where Iceberg tables are registered.
	* `name` - The catalog name within Polaris where Iceberg tables are registered.
	* `principal_role` - The Snowflake role used to access Polaris.
	* `properties_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the configuration file containing additional properties for the REST catalog. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
	* `uri` - The URL endpoint for the Polaris API. e.g.: 'https://<your-snowflake-account>.snowflakecomputing.com/polaris/api/catalog' 
* `client_id` - Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d 
* `client_secret_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored. Only applicable for authenticationType == OAUTH_M2M. Note: When provided, 'clientSecret' field must not be provided. 
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kafka cluster being referenced from Oracle Cloud Infrastructure Streaming with Apache Kafka. 
* `cluster_placement_group_id` - The OCID(https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource. Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud subscription id is provided. Otherwise the cluster placement group must not be provided. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_factory` - The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA' 
* `connection_string` - JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;' 
* `connection_type` - The connection type. 
* `connection_url` - JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>' 
* `consumer_properties` - The base64 encoded content of the consumer.properties file. 
* `core_site_xml` - The base64 encoded content of the Hadoop Distributed File System configuration file (core-site.xml). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Autonomous Json Database. 
* `database_name` - The name of the database. 
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `does_use_secret_ids` - Indicates that sensitive attributes are provided via Secrets. 
* `endpoint` - Service endpoint. Optional for Microsoft Fabric, default value: https://onelake.dfs.fabric.microsoft.com, for Azure Storage e.g: https://test.blob.core.windows.net, for Amazon S3 e.g.: 'https://s3.amazonaws.com', for Amazon Kinesis e.g.: 'https://kinesis.us-east-1.amazonaws.com'
* `fingerprint` - Fingerprint required by TLS security protocol. E.g.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c' 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `host` - Host and port separated by colon. Example: `"server.example.com:1234"`

	For multiple hosts, provide a comma separated list. Example: `"server1.example.com:1000,server1.example.com:2000"` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `ingress_ips` - List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.  Customers may optionally set up ingress security rules to restrict traffic from these IP addresses. 
	* `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet. 
* `jndi_connection_factory` - The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory' 
* `jndi_initial_context_factory` - The implementation of javax.naming.spi.InitialContextFactory interface that the client uses to obtain initial naming context. e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory' 
* `jndi_provider_url` - The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000' 
* `jndi_security_credentials_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the security credentials are stored associated to the principal. Note: When provided, 'jndiSecurityCredentials' field must not be provided. 
* `jndi_security_principal` - Specifies the identity of the principal (user) to be authenticated. e.g.: 'admin2' 
* `key_id` - Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `key_store_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl KeyStore password is stored. Note: When provided, 'keyStorePassword' field must not be provided. 
* `key_store_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the KeyStore file is stored. Note: When provided, 'keyStore' field must not be provided. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. If secretId is used plaintext field must not be provided. Note: When provided, 'password' field must not be provided. 
* `port` - The port of an endpoint usually specified for a connection. 
* `private_ip` - Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.

	The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `private_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Note: When provided, 'privateKeyFile' field must not be provided. 
* `private_key_passphrase_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password for the private key file. Note: When provided, 'privateKeyPassphrase' field must not be provided. 
* `producer_properties` - The base64 encoded content of the producer.properties file. 
* `public_key_fingerprint` - The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
* `redis_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster. 
* `region` - The name of the AWS region where the bucket is created. If not provided, GoldenGate will default to 'us-west-2'. Note: this property will become mandatory after May 20, 2026. 
* `routing_method` - Controls the network traffic direction to the target: SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.  SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet. DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected. 
* `sas_token_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored. Note: When provided, 'sasToken' field must not be provided. 
* `secret_access_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the secret access key is stored. Note: When provided, 'secretAccessKey' field must not be provided. 
* `security_attributes` - Security attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `security_protocol` - Security protocol for Java Message Service. If not provided, default is PLAIN. Optional until 2024-06-27, in the release after it will be made required. 
* `servers` - Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional.  If port is not specified, it defaults to 9200. Used for establishing the initial connection to the Elasticsearch cluster. Example: `"server1.example.com:4000,server2.example.com:4000"` 
* `service_account_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage. Note: When provided, 'serviceAccountKeyFile' field must not be provided. 
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `should_use_jndi` - If set to true, Java Naming and Directory Interface (JNDI) properties should be provided. 
* `should_use_resource_principal` - Specifies that the user intends to authenticate to the instance using a resource principal. Applicable only for Oracle Cloud Infrastructure Streaming connections. Only available from 23.9.0.0.0 GoldenGate versions. Note: When specified, 'username'/'password'/'passwordSecretId' fields must not be provided. Default: false 
* `should_validate_server_certificate` - If set to true, the driver validates the certificate that is sent by the database server. 
* `ssl_ca` - The base64 encoded certificate of the trusted certificate authorities (Trusted CA).  The supported file formats are .pem and .crt. In case of MYSQL and POSTGRESQL connections it is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_cert` - Client Certificate - The base64 encoded content of a .pem or .crt file containing the client public key (for 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_client_keystash_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystash file is stored,  which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Note: When provided, 'sslClientKeystash' field must not be provided. 
* `ssl_client_keystoredb_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystore file stored,  which created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

	Note: When provided, 'sslClientKeystoredb' field must not be provided. 
* `ssl_crl` - The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). Note: This is an optional property and only applicable if TLS/MTLS option is selected. It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `ssl_key_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored for the cert inside of the Keystore. In case it differs from the KeyStore password, it should be provided. Note: When provided, 'sslKeyPassword' field must not be provided. 
* `ssl_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key
	* The content of a .pem or .crt file containing the client private key (for 2-way SSL). Note: When provided, 'sslKey' field must not be provided. 
* `ssl_mode` - SSL mode for PostgreSQL.
* `ssl_server_certificate` - The base64 encoded file which contains the self-signed server certificate / Certificate Authority (CA) certificate. It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `state` - Possible lifecycle states for connection. 
* `storage` - Represents the storage of given type used in an Iceberg connection. 
	* `access_key_id` - Access key ID to access the Amazon S3 bucket. 
	* `account_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. 
	* `account_name` - Sets the Azure storage account name. 
	* `bucket` - Google Cloud Storage bucket where Iceberg stores metadata and data files.
	* `container` - The Azure Blob Storage container where Iceberg tables are stored.
	* `endpoint` - The Azure Blob Storage endpoint where Iceberg data is stored. e.g.: 'https://my-azure-storage-account.blob.core.windows.net' 
	* `project_id` - The Google Cloud Project where the bucket exists.
	* `region` - The AMAZON region where the S3 bucket is hosted. e.g.: 'us-east-2' 
	* `scheme_type` - The scheme of the storage.
	* `secret_access_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored. 
	* `service_account_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage. 
	* `storage_type` - The storage type used in the Iceberg connection. 
* `storage_credential_name` - Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS. 
* `stream_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection. 
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `technology_type` - The Kafka (e.g. Confluent) Schema Registry technology type. 
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `tenant_id` - Azure tenant ID of the application. e.g.: 14593954-d337-4a61-a364-9f758c64f97f 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `tls_ca_file` - Database Certificate - The base64 encoded content of a .pem file, containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified. 
* `tls_certificate_key_file_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file. Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided. 
* `tls_certificate_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
	* The content of a .pem file containing the client private key (for 2-way SSL). Note: When provided, 'tlsCertificateKeyFile' field must not be provided. 
* `trust_store_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl TrustStore password is stored. Note: When provided, 'trustStorePassword' field must not be provided. 
* `trust_store_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the TrustStore file is stored. Note: When provided, 'trustStore' field must not be provided. 
* `url` - Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081' 
* `user_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database. The user must have write access to the table they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint. 
* `username` - The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it. 
* `vault_id` - Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault. 
* `wallet_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the wallet file is stored.  The wallet contents Oracle GoldenGate uses to make connections to a database. Note: When provided, 'wallet' field must not be provided.

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

