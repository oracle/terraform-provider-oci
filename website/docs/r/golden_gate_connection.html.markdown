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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/goldengate/latest/Connection

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/goldengate

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
	auth_details {
		#Required
		auth_type = var.connection_auth_details_auth_type

		#Optional
		api_key = var.connection_auth_details_api_key
		api_key_secret_id = oci_vault_secret.test_secret.id
		base_url = var.connection_auth_details_base_url
		key_fingerprint = var.connection_auth_details_key_fingerprint
		region = var.connection_auth_details_region
		tenancy_id = oci_identity_tenancy.test_tenancy.id
		user_id = oci_identity_user.test_user.id
	}
	authentication_mode = var.connection_authentication_mode
	authentication_type = var.connection_authentication_type
	azure_authority_host = var.connection_azure_authority_host
	azure_tenant_id = oci_golden_gate_azure_tenant.test_azure_tenant.id
	bootstrap_servers {

		#Optional
		host = var.connection_bootstrap_servers_host
		port = var.connection_bootstrap_servers_port
	}
	catalog {
		#Required
		catalog_type = var.connection_catalog_catalog_type

		#Optional
		branch = var.connection_catalog_branch
		client_id = oci_golden_gate_client.test_client.id
		client_secret = var.connection_catalog_client_secret
		client_secret_secret_id = oci_vault_secret.test_secret.id
		glue_id = oci_golden_gate_glue.test_glue.id
		name = var.connection_catalog_name
		principal_role = var.connection_catalog_principal_role
		properties = var.connection_catalog_properties
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
	key_store_password_secret_id = oci_vault_secret.test_secret.id
	key_store_secret_id = oci_vault_secret.test_secret.id
	locks {
		#Required
		type = var.connection_locks_type

		#Optional
		message = var.connection_locks_message
	}
	max_input_chars = var.connection_max_input_chars
	model_key = var.connection_model_key
	nsg_ids = var.connection_nsg_ids
	password = var.connection_password
	password_secret_id = oci_vault_secret.test_secret.id
	port = var.connection_port
	private_key_file = var.connection_private_key_file
	private_key_file_secret_id = oci_vault_secret.test_secret.id
	private_key_passphrase = var.connection_private_key_passphrase
	private_key_passphrase_secret_id = oci_vault_secret.test_secret.id
	producer_properties = var.connection_producer_properties
	provider_type = var.connection_provider_type
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
		account_key = var.connection_storage_account_key
		account_key_secret_id = oci_vault_secret.test_secret.id
		account_name = var.connection_storage_account_name
		bucket = var.connection_storage_bucket
		container = var.connection_storage_container
		endpoint = var.connection_storage_endpoint
		project_id = oci_ai_document_project.test_project.id
		region = var.connection_storage_region
		scheme_type = var.connection_storage_scheme_type
		secret_access_key = var.connection_storage_secret_access_key
		secret_access_key_secret_id = oci_vault_secret.test_secret.id
		service_account_key_file = var.connection_storage_service_account_key_file
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
}
```

## Argument Reference

The following arguments are supported:

* `access_key_id` - (Required when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) Access key ID for Amazon connection types.
    * AMAZON_KINESIS: Access key ID to access Amazon Kinesis.
    * AMAZON_S3: Access key ID to access the Amazon S3 bucket.
      Note: Despite the "Id" suffix, this value is not an Oracle Cloud Infrastructure OCID.
* `account_key` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'. e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ== Deprecated: This field is deprecated and replaced by "accountKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `account_key_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. Note: When provided, 'accountKey' field must not be provided.
* `account_name` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name.
* `additional_attributes` - (Applicable when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) An array of name-value pair attribute entries. Used as additional parameters in connection string.
    * `name` - (Required when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name of the property entry.
    * `value` - (Required when connection_type=DB2 | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The value of the property entry.
* `auth_details` - (Applicable when connection_type=AI_MODEL) (Updatable) The information about new authentication details for an AI Model connection.
    * `api_key` - (Optional) (Updatable) API key for the AI model connection. Deprecated: This field is deprecated and replaced by "apiKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
    * `api_key_secret_id` - (Optional) (Updatable) API key secret OCID for the AI model connection.
    * `auth_type` - (Required) (Updatable) Authentication type used by the AI model connection.
    * `base_url` - (Applicable when auth_type=API_KEY) (Updatable) Base URL of the AI model endpoint. If not specified, the default base URL for the selected AI provider will be used.
    * `key_fingerprint` - (Required when auth_type=OCI_GEN_AI) (Updatable) Oracle Cloud Infrastructure Generative AI key fingerprint.
    * `region` - (Applicable when auth_type=OCI_GEN_AI) (Updatable) The name of the region. e.g.: us-ashburn-1 If the region is not provided, backend will default to the default region.
    * `tenancy_id` - (Applicable when auth_type=OCI_GEN_AI) (Updatable) Oracle Cloud Infrastructure Generative AI tenancy OCID. If this value is not provided, or is updated to an empty value, it defaults to the tenancy OCID of the user who is executing the operation.
    * `user_id` - (Applicable when auth_type=OCI_GEN_AI) (Updatable) Oracle Cloud Infrastructure Generative AI user OCID. If this value is not provided, or is updated to an empty value, it defaults to the OCID of the user who is executing the operation.
* `authentication_mode` - (Applicable when connection_type=ORACLE) (Updatable) Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections, when a databaseId is provided. The default value is MTLS.
* `authentication_type` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA_SCHEMA_REGISTRY | REDIS | SNOWFLAKE) (Updatable) Used authentication mechanism to be provided for the following connection types:
    * AZURE_DATA_LAKE_STORAGE, ELASTICSEARCH, KAFKA_SCHEMA_REGISTRY, REDIS, SNOWFLAKE
    * JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
    * DATABRICKS - Required fields by authentication types:
        * PERSONAL_ACCESS_TOKEN: username is always 'token', user must enter password
        * OAUTH_M2M: user must enter clientId and clientSecret
* `azure_authority_host` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The endpoint used for authentication with Microsoft Entra ID (formerly Azure Active Directory). Default value: https://login.microsoftonline.com When connecting to a non-public Azure Cloud, the endpoint must be provided, eg:
    * Azure China: https://login.chinacloudapi.cn/
    * Azure US Government: https://login.microsoftonline.us/
* `azure_tenant_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - (Applicable when connection_type=KAFKA) (Updatable) Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"`
    * `host` - (Required when connection_type=KAFKA) (Updatable) The name or address of a host.
    * `port` - (Applicable when connection_type=KAFKA) (Updatable) The port of an endpoint usually specified for a connection.
* `catalog` - (Required when connection_type=ICEBERG) (Updatable) The information about a new catalog of given type used in an Iceberg connection.
    * `branch` - (Required when catalog_type=NESSIE) (Updatable) The active branch of the Nessie catalog from which Iceberg reads and writes table metadata.
    * `catalog_type` - (Required) (Updatable) The catalog type.
    * `client_id` - (Required when catalog_type=POLARIS) (Updatable) The OAuth client ID used for authentication.
    * `client_secret` - (Applicable when catalog_type=POLARIS) (Updatable) Client secret required to connect to Polaris.
    * `client_secret_secret_id` - (Applicable when catalog_type=POLARIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect to Polaris.
    * `glue_id` - (Required when catalog_type=GLUE) (Updatable) The AWS Glue Catalog ID where Iceberg tables are registered.
    * `name` - (Required when catalog_type=POLARIS) (Updatable) The catalog name within Polaris where Iceberg tables are registered.
    * `principal_role` - (Required when catalog_type=POLARIS) (Updatable) The Snowflake role used to access Polaris.
    * `properties` - (Applicable when catalog_type=REST) (Updatable) The base64 encoded content of the configuration file containing additional properties for the REST catalog.
    * `properties_secret_id` - (Applicable when catalog_type=REST) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the configuration file containing additional properties for the REST catalog. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
    * `uri` - (Required when catalog_type=NESSIE | POLARIS | REST) (Updatable)
        * NESSIE: Nessie URI. e.g.: 'http://<nessie-server>.com:10001/api/v2'
        * POLARIS: The URL endpoint for the Polaris API. e.g.: 'https://<your-snowflake-account>.snowflakecomputing.com/polaris/api/catalog'
        * REST: The base URL for the REST Catalog API. e.g.: 'https://my-rest-catalog.example.com/api/v1'
* `client_id` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable)
    * AZURE_DATA_LAKE_STORAGE: Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
    * DATABRICKS: OAuth client id, only applicable for authenticationType == OAUTH_M2M.
    * MICROSOFT_FABRIC: Azure client ID of the application. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
* `client_secret` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable)
    * AZURE_DATA_LAKE_STORAGE: Azure client secret (aka application password) for authentication. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1
    * DATABRICKS: OAuth client secret, only applicable for authenticationType == OAUTH_M2M.
    * MICROSOFT_FABRIC: Client secret associated with the client id.
      Deprecated: This field is deprecated and replaced by "clientSecretSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `client_secret_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE | DATABRICKS | MICROSOFT_FABRIC) (Updatable)
    * AZURE_DATA_LAKE_STORAGE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
    * DATABRICKS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored. Only applicable for authenticationType == OAUTH_M2M.
    * MICROSOFT_FABRIC: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
      Note: When provided, 'clientSecret' field must not be provided.
* `cluster_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kafka cluster being referenced from Oracle Cloud Infrastructure Streaming with Apache Kafka.
* `cluster_placement_group_id` - (Optional) The OCID(https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource. Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud subscription id is provided. Otherwise the cluster placement group must not be provided.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
* `connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA'
* `connection_string` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | MONGODB | ORACLE) (Updatable)
    * ORACLE: Connect descriptor or Easy Connect Naming method used to connect to a database.
    * MONGODB: MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
    * AZURE_SYNAPSE_ANALYTICS: JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
* `connection_type` - (Required) (Updatable) The connection type.
* `connection_url` - (Required when connection_type=AMAZON_REDSHIFT | DATABRICKS | JAVA_MESSAGE_SERVICE | ORACLE_AI_DATA_PLATFORM | SNOWFLAKE) (Updatable)
    * JAVA_MESSAGE_SERVICE: Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'
    * SNOWFLAKE: JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>'
    * AMAZON_REDSHIFT: Connection URL. e.g.: 'jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb'
    * DATABRICKS: Connection URL. e.g.: 'jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb'
    * ORACLE_AI_DATA_PLATFORM: Connection URL. It must start with 'jdbc:spark://'
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

  Deprecated: This field is deprecated. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  When set to `true`, all sensitive information must be provided as Oracle Cloud Infrastructure Vault secrets using the corresponding `*SecretId` attributes of the connection (for example, `passwordSecretId`). Plain-text sensitive attributes (for example, `password`) must not be used. This ensures that sensitive information remains stored and managed in the customer's Oracle Cloud Infrastructure Vault rather than by the GoldenGate service.

  When set to false, sensitive information must be provided in the corresponding plain-text attributes (for example, `password`) rather than in secret OCID attributes. In this mode, the sensitive information is stored by the GoldenGate service. If `vaultId` and `keyId` are not specified, the GoldenGate service uses Oracle-managed encryption keys to encrypt the stored data.

  If `vaultId` and `keyId` are provided, the specified customer-managed key is used.
* `endpoint` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3 | AZURE_DATA_LAKE_STORAGE | GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE | MICROSOFT_FABRIC) (Updatable)
    * AMAZON_KINESIS: The endpoint URL of the Amazon Kinesis service. e.g.: 'https://kinesis.us-east-1.amazonaws.com' If not provided, GoldenGate will default to 'https://kinesis.<region>.amazonaws.com'.
    * AMAZON_S3: The Amazon Endpoint for S3. e.g.: 'https://my-bucket.s3.us-east-1.amazonaws.com' If not provided, GoldenGate will default to 'https://s3.<region>.amazonaws.com'.
    * AZURE_DATA_LAKE_STORAGE: Azure Storage service endpoint. e.g: https://test.blob.core.windows.net
    * GOOGLE_BIGQUERY: A legal URL to connect to BigQuery including scheme, server name and port, if not the default port. Default: https://bigquery.googleapis.com
    * GOOGLE_CLOUD_STORAGE: A legal URL to connect to Google Cloud Storage including scheme, server name and port, if not the default port. Default: https://storage.googleapis.com
    * MICROSOFT_FABRIC: Optional Microsoft Fabric service endpoint. Default value: https://onelake.dfs.fabric.microsoft.com
* `fingerprint` - (Applicable when connection_type=ELASTICSEARCH) (Updatable) Fingerprint required by TLS security protocol. Eg.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c'
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}`
* `host` - (Required when connection_type=DB2 | GENERIC | GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) Host and port separated by colon. Example: `"server.example.com:1234"`

  For multiple hosts, provide a comma separated list. Example: `"server1.example.com:1000,server1.example.com:2000"`
* `jndi_connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory'
* `jndi_initial_context_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The implementation of javax.naming.spi.InitialContextFactory interface that the client uses to obtain initial naming context. e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory'
* `jndi_provider_url` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000'
* `jndi_security_credentials` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The password associated to the principal. Deprecated: This field is deprecated and replaced by "jndiSecurityCredentialsSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `jndi_security_credentials_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the security credentials are stored associated to the principal. Note: When provided, 'jndiSecurityCredentials' field must not be provided.
* `jndi_security_principal` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) Specifies the identity of the principal (user) to be authenticated. e.g.: 'admin2'
* `key_id` - (Optional) (Updatable) References the Oracle Cloud Infrastructure Vault key in the Oracle Cloud Infrastructure Vault identified by `vaultId`.

  Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  The GoldenGate service uses this key to encrypt sensitive information (for example, `password`) that is provided in plain-text connection attributes through the API. This field is applicable only when `doesUseSecretIds` is set to `false`. If both `vaultId` and `keyId` are provided, the GoldenGate service uses the specified customer-managed key to encrypt the sensitive data. If neither `vaultId` nor `keyId` is provided, the GoldenGate service uses Oracle-managed encryption keys.
* `key_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the KeyStore file. Deprecated: This field is deprecated and replaced by "keyStoreSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `key_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The KeyStore password. Deprecated: This field is deprecated and replaced by "keyStorePasswordSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `key_store_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable)
    * JAVA_MESSAGE_SERVICE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the KeyStore password is stored.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka KeyStore password is stored.
    * KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl KeyStore password is stored.
    * REDIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis KeyStore password is stored.
      Note: When provided, 'keyStorePassword' field must not be provided.
* `key_store_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the KeyStore file is stored. Note: When provided, 'keyStore' field must not be provided.
* `locks` - (Optional) Locks associated with this resource.
    * `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked.
    * `type` - (Required) Type of the lock.
* `max_input_chars` - (Applicable when connection_type=AI_MODEL) (Updatable) Maximum number of input characters supported by this AI model connection.
* `model_key` - (Required when connection_type=AI_MODEL) (Updatable) AI model identifier.
* `nsg_ids` - (Optional) (Updatable) An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
* `password` - (Applicable when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DATABRICKS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. Deprecated: This field is deprecated and replaced by "passwordSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `password_secret_id` - (Applicable when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DATABRICKS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. If secretId is used plaintext field must not be provided. Note: When provided, 'password' field must not be provided.
* `port` - (Required when connection_type=DB2 | GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The port of an endpoint usually specified for a connection.
* `private_key_file` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The base64 encoded content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Deprecated: This field is deprecated and replaced by "privateKeyFileSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `private_key_file_secret_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Note: When provided, 'privateKeyFile' field must not be provided.
* `private_key_passphrase` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) Password if the private key file is encrypted. Deprecated: This field is deprecated and replaced by "privateKeyPassphraseSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `private_key_passphrase_secret_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password for the private key file. Note: When provided, 'privateKeyPassphrase' field must not be provided.
* `producer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the producer.properties file.
* `provider_type` - (Applicable when connection_type=AI_MODEL) AI Provider type used by the AI Model Connection.
* `public_key_fingerprint` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
* `redis_cluster_id` - (Applicable when connection_type=REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster.
* `region` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3 | OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The name of the region:
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM, ORACLE_NOSQL - OCI region, e.g.: 'us-ashburn-1'. If not provided, backend will default to the default region.
    * AMAZON_KINESIS - AWS region, e.g.: 'us-west-1'. If not provided, GoldenGate will default to 'us-west-1'. Note: this property will become mandatory after July 30, 2026.
    * AMAZON_S3 - AWS region where the bucket is created, e.g.: 'us-west-2'. If not provided, GoldenGate will default to 'us-west-2'. Note: this property will become mandatory after May 20, 2026.
* `routing_method` - (Optional) (Updatable) Controls the network traffic direction to the target: SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet. DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected. SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.

  Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `sas_token` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Credential that uses a shared access signature (SAS) to authenticate to an Azure Service. This property is required when 'authenticationType' is set to 'SHARED_ACCESS_SIGNATURE'. e.g.: ?sv=2020-06-08&ss=bfqt&srt=sco&sp=rwdlacupyx&se=2020-09-10T20:27:28Z&st=2022-08-05T12:27:28Z&spr=https&sig=C1IgHsiLBmTSStYkXXGLTP8it0xBrArcgCqOsZbXwIQ%3D Deprecated: This field is deprecated and replaced by "sasTokenSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `sas_token_secret_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored. Note: When provided, 'sasToken' field must not be provided.
* `secret_access_key` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable)
    * AMAZON_KINESIS: Secret access key to access the Amazon Kinesis.
    * AMAZON_S3: Secret access key to access the Amazon S3 bucket.
      Deprecated: This field is deprecated and replaced by "secretAccessKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `secret_access_key_secret_id` - (Applicable when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable)
    * AMAZON_KINESIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the secret access key is stored.
    * AMAZON_S3: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
      Note: When provided, 'secretAccessKey' field must not be provided.
* `security_attributes` - (Optional) (Updatable) Security attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
* `security_protocol` - (Required when connection_type=DB2 | ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA | MICROSOFT_SQLSERVER | MONGODB | MYSQL | POSTGRESQL | REDIS) (Updatable)
    * DB2: Security protocol for the DB2 database.
    * ELASTICSEARCH: Security protocol for Elasticsearch.
    * JAVA_MESSAGE_SERVICE: Security protocol for Java Message Service. If not provided, default is PLAIN. Optional until 2024-06-27, in the release after it will be made required.
    * KAFKA: Security Type for Kafka.
    * MICROSOFT_SQLSERVER: Security Type for Microsoft SQL Server.
    * MONGODB: Security Type for MongoDB.
    * MYSQL: Security Type for MySQL.
    * POSTGRESQL: Security protocol for PostgreSQL.
    * REDIS: Security protocol for Redis.
* `servers` - (Required when connection_type=ELASTICSEARCH | REDIS) (Updatable)
    * ELASTICSEARCH: Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional. If port is not specified, it defaults to 9200. Used for establishing the initial connection to the Elasticsearch cluster. Example: `"server1.example.com:4000,server2.example.com:4000"`
    * REDIS: Comma separated list of Redis server addresses, specified as host:port entries, where :port is optional. If port is not specified, it defaults to 6379. Used for establishing the initial connection to the Redis cluster. Example: `"server1.example.com:6379,server2.example.com:6379"`
* `service_account_key_file` - (Applicable when connection_type=GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE | GOOGLE_PUBSUB) (Updatable)
    * GOOGLE_BIGQUERY: The base64 encoded content of the service account key file containing the credentials required to use Google BigQuery.
    * GOOGLE_CLOUD_STORAGE: The base64 encoded content of the service account key file containing the credentials required to use Google Cloud Storage.
    * GOOGLE_PUBSUB: The base64 encoded content of the service account key file containing the credentials required to use Google PubSub.
      Deprecated: This field is deprecated and replaced by "serviceAccountKeyFileSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `service_account_key_file_secret_id` - (Applicable when connection_type=GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE | GOOGLE_PUBSUB) (Updatable)
    * GOOGLE_BIGQUERY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google BigQuery.
    * GOOGLE_CLOUD_STORAGE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage.
    * GOOGLE_PUBSUB: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google PubSub.
      Note: When provided, 'serviceAccountKeyFile' field must not be provided.
* `session_mode` - (Applicable when connection_type=ORACLE) (Updatable) Specifies the session mode for the database connection. Use REDIRECT only for RAC databases with SCAN listeners that return IP addresses. For RAC databases with SCAN listeners that return FQDNs, and for all other Oracle database technologies, use DIRECT. In RAC deployments, SCAN listeners redirects a connection to a specific database node, identified by either IP address or FQDN. It is recommended to configure RAC with FQDN-based SCAN listeners.

  The default is DIRECT, except when databaseId is provided and the discovered database relies on the SCAN listener. In this case, the default is REDIRECT.

  Deprecated: Defaulting to the REDIRECT session mode will be removed after April 21, 2027.
* `should_use_jndi` - (Required when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
* `should_use_resource_principal` - (Applicable when connection_type=KAFKA | OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable)
    * KAFKA: Specifies that the user intends to authenticate to the instance using a resource principal. Applicable only for Oracle Cloud Infrastructure Streaming connections. Only available from 23.9.0.0.0 GoldenGate versions. Note: When specified, 'username'/'password'/'passwordSecretId' fields must not be provided. Default: false
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM, ORACLE_NOSQL: Specifies that the user intends to authenticate to the instance using a resource principal. Default: false
* `should_validate_server_certificate` - (Applicable when connection_type=MICROSOFT_SQLSERVER) (Updatable) If set to true, the driver validates the certificate that is sent by the database server.
* `ssl_ca` - (Applicable when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable)
    * MICROSOFT_SQLSERVER: Database Certificate - The base64 encoded content of a .pem or .crt file containing the server public key (for 1-way SSL).
    * MYSQL: Database Certificate - The base64 encoded content of a .pem or .crt file containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded certificate of the trusted certificate authorities (Trusted CA) for PostgreSQL. The supported file formats are .pem and .crt. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_cert` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable)
    * MYSQL: Client Certificate - The base64 encoded content of a .pem or .crt file containing the client public key (for 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded certificate of the PostgreSQL server. The supported file formats are .pem and .crt. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_client_keystash` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded keystash file which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Deprecated: This field is deprecated and replaced by "sslClientKeystashSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `ssl_client_keystash_secret_id` - (Applicable when connection_type=DB2) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystash file is stored,  which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Note: When provided, 'sslClientKeystash' field must not be provided.
* `ssl_client_keystoredb` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded keystore file created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Deprecated: This field is deprecated and replaced by "sslClientKeystoredbSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `ssl_client_keystoredb_secret_id` - (Applicable when connection_type=DB2) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystore file stored,  which created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Note: When provided, 'sslClientKeystoredb' field must not be provided.
* `ssl_crl` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable)
    * MYSQL: The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). Note: This is an optional property and only applicable if TLS/MTLS option is selected. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_key` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable)
    * MYSQL: Client Key - The base64 encoded content of a .pem or .crt file containing the client private key (for 2-way SSL).
    * POSTGRESQL: The base64 encoded private key of the PostgreSQL server. The supported file formats are .pem and .crt.
      Deprecated: This field is deprecated and replaced by "sslKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `ssl_key_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The password for the cert inside of the KeyStore. In case it differs from the KeyStore password, it should be provided. Deprecated: This field is deprecated and replaced by "sslKeyPasswordSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `ssl_key_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable)
    * JAVA_MESSAGE_SERVICE, KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored for the cert inside of the Keystore. In case it differs from the KeyStore password, it should be provided.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl Key password is stored.
      Note: When provided, 'sslKeyPassword' field must not be provided.
* `ssl_key_secret_id` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable)
    * MYSQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key - The content of a .pem or .crt file containing the client private key (for 2-way SSL).
    * POSTGRESQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the private key of the PostgreSQL server. The supported file formats are .pem and .crt.
      Note: When provided, 'sslKey' field must not be provided.
* `ssl_mode` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable)
    * MYSQL: SSL modes for MySQL.
    * POSTGRESQL: SSL modes for PostgreSQL.
* `ssl_server_certificate` - (Applicable when connection_type=DB2) (Updatable) The base64 encoded file which contains the self-signed server certificate / Certificate Authority (CA) certificate. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `storage` - (Required when connection_type=ICEBERG) (Updatable) The information about a new storage of given type used in an Iceberg connection.
    * `access_key_id` - (Required when storage_type=AMAZON_S3 | OCI_OBJECT_STORAGE_S3_API) (Updatable)
        * AMAZON_S3: Access key ID to access the Amazon S3 bucket.
        * OCI_OBJECT_STORAGE_S3_API: Access Key ID from the Oracle Cloud Infrastructure IAM user's Customer Secret Key pair used to authenticate to Oracle Cloud Infrastructure Object Storage via the S3 Compatibility API.
          Note: Despite the "Id" suffix, this value is not an Oracle Cloud Infrastructure OCID.
    * `account_key` - (Applicable when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'. e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ== Deprecated: This field is deprecated and replaced by "accountKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
    * `account_key_secret_id` - (Applicable when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored.
    * `account_name` - (Required when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name.
    * `bucket` - (Required when storage_type=AMAZON_S3 | GOOGLE_CLOUD_STORAGE | OCI_OBJECT_STORAGE_S3_API) (Updatable)
        * AMAZON_S3: S3 bucket where Iceberg stores metadata and data files.
        * GOOGLE_CLOUD_STORAGE: Google Cloud Storage bucket where Iceberg stores metadata and data files.
        * OCI_OBJECT_STORAGE_S3_API: Target Oracle Cloud Infrastructure Object Storage bucket name where Iceberg stores table metadata and data files.
    * `container` - (Required when storage_type=AZURE_DATA_LAKE_STORAGE) (Updatable) The Azure Blob Storage container where Iceberg tables are stored.
    * `endpoint` - (Applicable when storage_type=AMAZON_S3 | AZURE_DATA_LAKE_STORAGE | GOOGLE_CLOUD_STORAGE | OCI_OBJECT_STORAGE_S3_API; required when storage_type=OCI_OBJECT_STORAGE_S3_API) (Updatable)
        * AMAZON_S3: The endpoint URL of the Amazon S3 storage service. e.g.: 'https://s3.amazonaws.com'
        * AZURE_DATA_LAKE_STORAGE: The Azure Blob Storage endpoint where Iceberg data is stored. e.g.: 'https://my-azure-storage-account.blob.core.windows.net'
        * GOOGLE_CLOUD_STORAGE: A legal URL to connect to Google Cloud Storage including scheme, server name and port, if not the default port. Default: https://storage.googleapis.com
        * OCI_OBJECT_STORAGE_S3_API: Oracle Cloud Infrastructure Object Storage S3 Compatibility API endpoint URL. Format: "https://<namespace>.compat.objectstorage.<region>.<domain>" Example: "https://mynamespace.compat.objectstorage.us-ashburn-1.oraclecloud.com"
    * `project_id` - (Required when storage_type=GOOGLE_CLOUD_STORAGE) (Updatable) The Google Cloud Project where the bucket exists.
    * `region` - (Required when storage_type=AMAZON_S3) (Updatable) The AMAZON region where the S3 bucket is hosted. e.g.: 'us-east-2'
    * `scheme_type` - (Required when storage_type=AMAZON_S3) (Updatable) The scheme of the storage.
    * `secret_access_key` - (Applicable when storage_type=AMAZON_S3 | OCI_OBJECT_STORAGE_S3_API) (Updatable)
        * AMAZON_S3: Secret access key to access the Amazon S3 bucket.
        * OCI_OBJECT_STORAGE_S3_API: Secret Access Key from the Oracle Cloud Infrastructure IAM user's Customer Secret Key pair used to authenticate to Oracle Cloud Infrastructure Object Storage via the S3 Compatibility API.
          Deprecated: This field is deprecated and replaced by "secretAccessKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
    * `secret_access_key_secret_id` - (Applicable when storage_type=AMAZON_S3 | OCI_OBJECT_STORAGE_S3_API) (Updatable)
        * AMAZON_S3: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
        * OCI_OBJECT_STORAGE_S3_API: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key used for Oracle Cloud Infrastructure Object Storage S3 Compatibility authentication is stored.
    * `service_account_key_file` - (Applicable when storage_type=GOOGLE_CLOUD_STORAGE) (Updatable) The base64 encoded content of the service account key file containing the credentials required to use Google Cloud Storage. Deprecated: This field is deprecated and replaced by "serviceAccountKeyFileSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
    * `service_account_key_file_secret_id` - (Applicable when storage_type=GOOGLE_CLOUD_STORAGE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage.
    * `storage_type` - (Required) (Updatable) The storage type used in the Iceberg connection.
* `storage_credential_name` - (Applicable when connection_type=DATABRICKS) (Updatable) Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS.
* `stream_pool_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced.
* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `technology_type` - (Required) The technology type.
* `tenancy_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy.
* `tenant_id` - (Required when connection_type=MICROSOFT_FABRIC) (Updatable) Azure tenant ID of the application. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `tls_ca_file` - (Applicable when connection_type=MONGODB) (Updatable) Database Certificate - The base64 encoded content of a .pem file, containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `tls_certificate_key_file` - (Applicable when connection_type=MONGODB) (Updatable) Client Certificate - The base64 encoded content of a .pem file, containing the client public key (for 2-way SSL). Deprecated: This field is deprecated and replaced by "tlsCertificateKeyFileSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `tls_certificate_key_file_password` - (Applicable when connection_type=MONGODB) (Updatable) Client Certificate key file password. Deprecated: This field is deprecated and replaced by "tlsCertificateKeyFilePasswordSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `tls_certificate_key_file_password_secret_id` - (Applicable when connection_type=MONGODB) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file. Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided.
* `tls_certificate_key_file_secret_id` - (Applicable when connection_type=MONGODB) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
    * The content of a .pem file containing the client private key (for 2-way SSL). Note: When provided, 'tlsCertificateKeyFile' field must not be provided.
* `trust_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the TrustStore file. Deprecated: This field is deprecated and replaced by "trustStoreSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `trust_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The TrustStore password. Deprecated: This field is deprecated and replaced by "trustStorePasswordSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `trust_store_password_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable)
    * JAVA_MESSAGE_SERVICE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the TrustStore password is stored.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka TrustStore password is stored.
    * KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl TrustStore password is stored.
    * REDIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis TrustStore password is stored.
      Note: When provided, 'trustStorePassword' field must not be provided.
* `trust_store_secret_id` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the TrustStore file is stored. Note: When provided, 'trustStore' field must not be provided.
* `url` - (Required when connection_type=KAFKA_SCHEMA_REGISTRY) (Updatable) Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081'
* `user_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_AI_DATA_PLATFORM | ORACLE_NOSQL) (Updatable)
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access Object Storage. The user must have write access to the bucket they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint.
    * ORACLE_NOSQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database. The user must have write access to the table they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint.
* `username` - (Required when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | DB2 | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it.
* `vault_id` - (Optional) (Updatable) References the Oracle Cloud Infrastructure Vault that contains the customer-managed encryption key identified by `keyId`.

  Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  This field is applicable only when `doesUseSecretIds` is set to `false`. If `vaultId` is provided, `keyId` must also be provided.
* `wallet` - (Applicable when connection_type=ORACLE) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database. This attribute is expected to be base64 encoded. Deprecated: This field is deprecated and replaced by "walletSecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `wallet_secret_id` - (Applicable when connection_type=ORACLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the wallet file is stored.  The wallet contents Oracle GoldenGate uses to make connections to a database. Note: When provided, 'wallet' field must not be provided.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_key_id` - Access key ID for Amazon connection types.
    * AMAZON_KINESIS: Access key ID to access Amazon Kinesis.
    * AMAZON_S3: Access key ID to access the Amazon S3 bucket.
      Note: Despite the "Id" suffix, this value is not an Oracle Cloud Infrastructure OCID.
* `account_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. Note: When provided, 'accountKey' field must not be provided.
* `account_name` - Sets the Azure storage account name.
* `additional_attributes` - An array of name-value pair attribute entries. Used as additional parameters in connection string.
    * `name` - The name of the property entry.
    * `value` - The value of the property entry.
* `auth_details` - Represents authentication details for an AI Model connection.
    * `api_key` - API key for the AI model connection. Deprecated: This field is deprecated and replaced by "apiKeySecretId". This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
    * `api_key_secret_id` - API key secret OCID for the AI model connection.
    * `auth_type` - Authentication type used by the AI model connection.
    * `base_url` - Base URL of the AI model endpoint. If not specified, the default base URL for the selected AI provider will be used.
    * `key_fingerprint` - Oracle Cloud Infrastructure Generative AI key fingerprint.
    * `region` - The name of the region. e.g.: us-ashburn-1 If the region is not provided, backend will default to the default region.
    * `tenancy_id` - Oracle Cloud Infrastructure Generative AI tenancy OCID. If this value is not provided, or is updated to an empty value, it defaults to the tenancy OCID of the user who is executing the operation.
    * `user_id` - Oracle Cloud Infrastructure Generative AI user OCID. If this value is not provided, or is updated to an empty value, it defaults to the OCID of the user who is executing the operation.
* `authentication_mode` - Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections, when a databaseId is provided. The default value is MTLS.
* `authentication_type` - Used authentication mechanism to be provided for the following connection types:
    * AZURE_DATA_LAKE_STORAGE, ELASTICSEARCH, KAFKA_SCHEMA_REGISTRY, REDIS, SNOWFLAKE
    * JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
    * DATABRICKS - Required fields by authentication types:
        * PERSONAL_ACCESS_TOKEN: username is always 'token', user must enter password
        * OAUTH_M2M: user must enter clientId and clientSecret
* `azure_authority_host` - The endpoint used for authentication with Microsoft Entra ID (formerly Azure Active Directory). Default value: https://login.microsoftonline.com When connecting to a non-public Azure Cloud, the endpoint must be provided, eg:
    * Azure China: https://login.chinacloudapi.cn/
    * Azure US Government: https://login.microsoftonline.us/
* `azure_tenant_id` - Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"`
    * `host` - The name or address of a host.
    * `port` - The port of an endpoint usually specified for a connection.
    * `private_ip` - This property is not available when creating connections. For existing deprecated connections having this value set, the value cannot be updated; set it to empty.

      For deprecated connections created with this field in the past, either the private IP had to be specified in the connectionString or host field, or the host name had to be resolvable in the target VCN.
* `catalog` - Represents the catalog of given type used in an Iceberg connection.
    * `branch` - The active branch of the Nessie catalog from which Iceberg reads and writes table metadata.
    * `catalog_type` - The catalog type.
    * `client_id` - The OAuth client ID used for authentication.
    * `client_secret_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect to Polaris.
    * `glue_id` - The AWS Glue Catalog ID where Iceberg tables are registered.
    * `name` - The catalog name within Polaris where Iceberg tables are registered.
    * `principal_role` - The Snowflake role used to access Polaris.
    * `properties_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the configuration file containing additional properties for the REST catalog. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
    * `uri`
        * NESSIE: Nessie URI. e.g.: 'http://<nessie-server>.com:10001/api/v2'
        * POLARIS: The URL endpoint for the Polaris API. e.g.: 'https://<your-snowflake-account>.snowflakecomputing.com/polaris/api/catalog'
        * REST: The base URL for the REST Catalog API. e.g.: 'https://my-rest-catalog.example.com/api/v1'
* `client_id`
    * AZURE_DATA_LAKE_STORAGE: Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
    * DATABRICKS: OAuth client id, only applicable for authenticationType == OAUTH_M2M.
    * MICROSOFT_FABRIC: Azure client ID of the application. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
* `client_secret_secret_id`
    * AZURE_DATA_LAKE_STORAGE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
    * DATABRICKS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored. Only applicable for authenticationType == OAUTH_M2M.
    * MICROSOFT_FABRIC: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored.
      Note: When provided, 'clientSecret' field must not be provided.
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kafka cluster being referenced from Oracle Cloud Infrastructure Streaming with Apache Kafka.
* `cluster_placement_group_id` - The OCID(https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group for the resource. Only applicable for multicloud subscriptions. The cluster placement group id must be provided when a multicloud subscription id is provided. Otherwise the cluster placement group must not be provided.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
* `connection_factory` - The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA'
* `connection_string`
    * ORACLE: Connect descriptor or Easy Connect Naming method used to connect to a database.
    * MONGODB: MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
    * AZURE_SYNAPSE_ANALYTICS: JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
* `connection_type` - The connection type.
* `connection_url`
    * JAVA_MESSAGE_SERVICE: Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'
    * SNOWFLAKE: JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>'
    * AMAZON_REDSHIFT: Connection URL. e.g.: 'jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb'
    * DATABRICKS: Connection URL. e.g.: 'jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb'
    * ORACLE_AI_DATA_PLATFORM: Connection URL. It must start with 'jdbc:spark://'
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

  Deprecated: This field is deprecated. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  When set to `true`, all sensitive information must be provided as Oracle Cloud Infrastructure Vault secrets using the corresponding `*SecretId` attributes of the connection (for example, `passwordSecretId`). Plain-text sensitive attributes (for example, `password`) must not be used. This ensures that sensitive information remains stored and managed in the customer's Oracle Cloud Infrastructure Vault rather than by the GoldenGate service.

  When set to false, sensitive information must be provided in the corresponding plain-text attributes (for example, `password`) rather than in secret OCID attributes. In this mode, the sensitive information is stored by the GoldenGate service. If `vaultId` and `keyId` are not specified, the GoldenGate service uses Oracle-managed encryption keys to encrypt the stored data.

  If `vaultId` and `keyId` are provided, the specified customer-managed key is used.
* `endpoint`
    * AMAZON_KINESIS: The endpoint URL of the Amazon Kinesis service. e.g.: 'https://kinesis.us-east-1.amazonaws.com' If not provided, GoldenGate will default to 'https://kinesis.<region>.amazonaws.com'.
    * AMAZON_S3: The Amazon Endpoint for S3. e.g.: 'https://my-bucket.s3.us-east-1.amazonaws.com' If not provided, GoldenGate will default to 'https://s3.<region>.amazonaws.com'.
    * AZURE_DATA_LAKE_STORAGE: Azure Storage service endpoint. e.g: https://test.blob.core.windows.net
    * GOOGLE_BIGQUERY: A legal URL to connect to BigQuery including scheme, server name and port, if not the default port. Default: https://bigquery.googleapis.com
    * GOOGLE_CLOUD_STORAGE: A legal URL to connect to Google Cloud Storage including scheme, server name and port, if not the default port. Default: https://storage.googleapis.com
    * MICROSOFT_FABRIC: Optional Microsoft Fabric service endpoint. Default value: https://onelake.dfs.fabric.microsoft.com
* `fingerprint` - Fingerprint required by TLS security protocol. Eg.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c'
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
* `key_id` - References the Oracle Cloud Infrastructure Vault key in the Oracle Cloud Infrastructure Vault identified by `vaultId`.

  Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  The GoldenGate service uses this key to encrypt sensitive information (for example, `password`) that is provided in plain-text connection attributes through the API. This field is applicable only when `doesUseSecretIds` is set to `false`. If both `vaultId` and `keyId` are provided, the GoldenGate service uses the specified customer-managed key to encrypt the sensitive data. If neither `vaultId` nor `keyId` is provided, the GoldenGate service uses Oracle-managed encryption keys.
* `key_store_password_secret_id`
    * JAVA_MESSAGE_SERVICE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the KeyStore password is stored.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka KeyStore password is stored.
    * KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl KeyStore password is stored.
    * REDIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis KeyStore password is stored.
      Note: When provided, 'keyStorePassword' field must not be provided.
* `key_store_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the KeyStore file is stored. Note: When provided, 'keyStore' field must not be provided.
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
* `locks` - Locks associated with this resource.
    * `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked.
    * `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
    * `time_created` - When the lock was created.
    * `type` - Type of the lock.
* `max_input_chars` - Maximum number of input characters supported by this AI model connection.
* `model_key` - AI model identifier.
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. If secretId is used plaintext field must not be provided. Note: When provided, 'password' field must not be provided.
* `port` - The port of an endpoint usually specified for a connection.
* `private_ip` - This property is not available when creating connections. For existing deprecated connections having this value set, the value cannot be updated; set it to empty.

  For deprecated connections created with this field in the past, either the private IP had to be specified in the connectionString or host field, or the host name had to be resolvable in the target VCN.
* `private_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Note: When provided, 'privateKeyFile' field must not be provided.
* `private_key_passphrase_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password for the private key file. Note: When provided, 'privateKeyPassphrase' field must not be provided.
* `producer_properties` - The base64 encoded content of the producer.properties file.
* `provider_type` - AI Provider type used by the AI Model Connection.
* `public_key_fingerprint` - The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
* `redis_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster.
* `region` - The name of the region:
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM, ORACLE_NOSQL - OCI region, e.g.: 'us-ashburn-1'. If not provided, backend will default to the default region.
    * AMAZON_KINESIS - AWS region, e.g.: 'us-west-1'. If not provided, GoldenGate will default to 'us-west-1'. Note: this property will become mandatory after July 30, 2026.
    * AMAZON_S3 - AWS region where the bucket is created, e.g.: 'us-west-2'. If not provided, GoldenGate will default to 'us-west-2'. Note: this property will become mandatory after May 20, 2026.
* `routing_method` - Controls the network traffic direction to the target: SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet. DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected. SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.

  Deprecated: SHARED_SERVICE_ENDPOINT is deprecated. Use another supported routingMethod value, or update existing connections to use a supported routing method. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
* `sas_token_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored. Note: When provided, 'sasToken' field must not be provided.
* `secret_access_key_secret_id`
    * AMAZON_KINESIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the secret access key is stored.
    * AMAZON_S3: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
      Note: When provided, 'secretAccessKey' field must not be provided.
* `security_attributes` - Security attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}`
* `security_protocol`
    * DB2: Security protocol for the DB2 database.
    * ELASTICSEARCH: Security protocol for Elasticsearch.
    * JAVA_MESSAGE_SERVICE: Security protocol for Java Message Service. If not provided, default is PLAIN. Optional until 2024-06-27, in the release after it will be made required.
    * KAFKA: Security Type for Kafka.
    * MICROSOFT_SQLSERVER: Security Type for Microsoft SQL Server.
    * MONGODB: Security Type for MongoDB.
    * MYSQL: Security Type for MySQL.
    * POSTGRESQL: Security protocol for PostgreSQL.
    * REDIS: Security protocol for Redis.
* `servers`
    * ELASTICSEARCH: Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional. If port is not specified, it defaults to 9200. Used for establishing the initial connection to the Elasticsearch cluster. Example: `"server1.example.com:4000,server2.example.com:4000"`
    * REDIS: Comma separated list of Redis server addresses, specified as host:port entries, where :port is optional. If port is not specified, it defaults to 6379. Used for establishing the initial connection to the Redis cluster. Example: `"server1.example.com:6379,server2.example.com:6379"`
* `service_account_key_file_secret_id`
    * GOOGLE_BIGQUERY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google BigQuery.
    * GOOGLE_CLOUD_STORAGE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage.
    * GOOGLE_PUBSUB: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google PubSub.
      Note: When provided, 'serviceAccountKeyFile' field must not be provided.
* `session_mode` - Specifies the session mode for the database connection. Use REDIRECT only for RAC databases with SCAN listeners that return IP addresses. For RAC databases with SCAN listeners that return FQDNs, and for all other Oracle database technologies, use DIRECT. In RAC deployments, SCAN listeners redirects a connection to a specific database node, identified by either IP address or FQDN. It is recommended to configure RAC with FQDN-based SCAN listeners.

  The default is DIRECT, except when databaseId is provided and the discovered database relies on the SCAN listener. In this case, the default is REDIRECT.

  Deprecated: Defaulting to the REDIRECT session mode will be removed after April 21, 2027.
* `should_use_jndi` - If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
* `should_use_resource_principal`
    * KAFKA: Specifies that the user intends to authenticate to the instance using a resource principal. Applicable only for Oracle Cloud Infrastructure Streaming connections. Only available from 23.9.0.0.0 GoldenGate versions. Note: When specified, 'username'/'password'/'passwordSecretId' fields must not be provided. Default: false
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM, ORACLE_NOSQL: Specifies that the user intends to authenticate to the instance using a resource principal. Default: false
* `should_validate_server_certificate` - If set to true, the driver validates the certificate that is sent by the database server.
* `ssl_ca`
    * MICROSOFT_SQLSERVER: Database Certificate - The base64 encoded content of a .pem or .crt file containing the server public key (for 1-way SSL).
    * MYSQL: Database Certificate - The base64 encoded content of a .pem or .crt file containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded certificate of the trusted certificate authorities (Trusted CA) for PostgreSQL. The supported file formats are .pem and .crt. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_cert`
    * MYSQL: Client Certificate - The base64 encoded content of a .pem or .crt file containing the client public key (for 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded certificate of the PostgreSQL server. The supported file formats are .pem and .crt. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_client_keystash_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystash file is stored,  which contains the encrypted password to the key database file. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Note: When provided, 'sslClientKeystash' field must not be provided.
* `ssl_client_keystoredb_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystore file stored,  which created at the client containing the server certificate / CA root certificate. This property is not supported for IBM Db2 for i, as client TLS mode is not available.

  Note: When provided, 'sslClientKeystoredb' field must not be provided.
* `ssl_crl`
    * MYSQL: The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). Note: This is an optional property and only applicable if TLS/MTLS option is selected. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
    * POSTGRESQL: The base64 encoded list of certificates revoked by the trusted certificate authorities (Trusted CA). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `ssl_key_password_secret_id`
    * JAVA_MESSAGE_SERVICE, KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored for the cert inside of the Keystore. In case it differs from the KeyStore password, it should be provided.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl Key password is stored.
      Note: When provided, 'sslKeyPassword' field must not be provided.
* `ssl_key_secret_id`
    * MYSQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key - The content of a .pem or .crt file containing the client private key (for 2-way SSL).
    * POSTGRESQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the private key of the PostgreSQL server. The supported file formats are .pem and .crt.
      Note: When provided, 'sslKey' field must not be provided.
* `ssl_mode`
    * MYSQL: SSL modes for MySQL.
    * POSTGRESQL: SSL modes for PostgreSQL.
* `ssl_server_certificate` - The base64 encoded file which contains the self-signed server certificate / Certificate Authority (CA) certificate. It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `state` - Possible lifecycle states for connection.
* `storage` - Represents the storage of given type used in an Iceberg connection.
    * `access_key_id`
        * AMAZON_S3: Access key ID to access the Amazon S3 bucket.
        * OCI_OBJECT_STORAGE_S3_API: Access Key ID from the Oracle Cloud Infrastructure IAM user's Customer Secret Key pair used to authenticate to Oracle Cloud Infrastructure Object Storage via the S3 Compatibility API.
          Note: Despite the "Id" suffix, this value is not an Oracle Cloud Infrastructure OCID.
    * `account_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored.
    * `account_name` - Sets the Azure storage account name.
    * `bucket`
        * AMAZON_S3: S3 bucket where Iceberg stores metadata and data files.
        * GOOGLE_CLOUD_STORAGE: Google Cloud Storage bucket where Iceberg stores metadata and data files.
        * OCI_OBJECT_STORAGE_S3_API: Target Oracle Cloud Infrastructure Object Storage bucket name where Iceberg stores table metadata and data files.
    * `container` - The Azure Blob Storage container where Iceberg tables are stored.
    * `endpoint`
        * AMAZON_S3: The endpoint URL of the Amazon S3 storage service. e.g.: 'https://s3.amazonaws.com'
        * AZURE_DATA_LAKE_STORAGE: The Azure Blob Storage endpoint where Iceberg data is stored. e.g.: 'https://my-azure-storage-account.blob.core.windows.net'
        * GOOGLE_CLOUD_STORAGE: A legal URL to connect to Google Cloud Storage including scheme, server name and port, if not the default port. Default: https://storage.googleapis.com
        * OCI_OBJECT_STORAGE_S3_API: Oracle Cloud Infrastructure Object Storage S3 Compatibility API endpoint URL. Format: "https://<namespace>.compat.objectstorage.<region>.<domain>" Example: "https://mynamespace.compat.objectstorage.us-ashburn-1.oraclecloud.com"
    * `project_id` - The Google Cloud Project where the bucket exists.
    * `region` - The AMAZON region where the S3 bucket is hosted. e.g.: 'us-east-2'
    * `scheme_type` - The scheme of the storage.
    * `secret_access_key_secret_id`
        * AMAZON_S3: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
        * OCI_OBJECT_STORAGE_S3_API: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key used for Oracle Cloud Infrastructure Object Storage S3 Compatibility authentication is stored.
    * `service_account_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which contains the credentials required to use Google Cloud Storage.
    * `storage_type` - The storage type used in the Iceberg connection.
* `storage_credential_name` - Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS.
* `stream_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection.
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}`
* `technology_type` - The technology type.
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy.
* `tenant_id` - Azure tenant ID of the application. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `tls_ca_file` - Database Certificate - The base64 encoded content of a .pem file, containing the server public key (for 1 and 2-way SSL). It is not included in GET responses if the `view=COMPACT` query parameter is specified.
* `tls_certificate_key_file_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password of the tls certificate key file. Note: When provided, 'tlsCertificateKeyFilePassword' field must not be provided.
* `tls_certificate_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the certificate key file of the mtls connection.
    * The content of a .pem file containing the client private key (for 2-way SSL). Note: When provided, 'tlsCertificateKeyFile' field must not be provided.
* `trust_store_password_secret_id`
    * JAVA_MESSAGE_SERVICE: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the TrustStore password is stored.
    * KAFKA: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka TrustStore password is stored.
    * KAFKA_SCHEMA_REGISTRY: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the kafka Ssl TrustStore password is stored.
    * REDIS: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Redis TrustStore password is stored.
      Note: When provided, 'trustStorePassword' field must not be provided.
* `trust_store_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the TrustStore file is stored. Note: When provided, 'trustStore' field must not be provided.
* `url` - Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081'
* `user_id`
    * OCI_OBJECT_STORAGE, ORACLE_AI_DATA_PLATFORM: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access Object Storage. The user must have write access to the bucket they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint.
    * ORACLE_NOSQL: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database. The user must have write access to the table they want to connect to. If the user is not provided, backend will default to the user who is calling the API endpoint.
* `username` - The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it.
* `vault_id` - References the Oracle Cloud Infrastructure Vault that contains the customer-managed encryption key identified by `keyId`.

  Deprecated: This field is deprecated for GoldenGate connections. Sensitive attributes should be provided using the corresponding Secret OCID attributes of the connection (for example, `passwordSecretId`) instead of plain-text attributes encrypted with `vaultId` and `keyId`. This change follows the GoldenGate "Plain Text Fields in Connections" deprecation: https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate

  This field is applicable only when `doesUseSecretIds` is set to `false`. If `vaultId` is provided, `keyId` must also be provided.
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
