---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connection"
sidebar_current: "docs-oci-datasource-golden_gate-connection"
description: |-
  Provides details about a specific Connection in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_connection
This data source provides details about a specific Connection resource in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a Connection.


## Example Usage

```hcl
data "oci_golden_gate_connection" "test_connection" {
	#Required
	connection_id = oci_golden_gate_connection.test_connection.id

	#Optional
	view = var.connection_view
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Connection.
* `view` - (Optional) Selects the connection fields returned in connection details.


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
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced.
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
* `host` - The name or address of a host.
  In case of Generic connection type it represents the Host and port separated by colon. Example: `"server.example.com:1234"`
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
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. Note: When provided, 'password' field must not be provided.
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
