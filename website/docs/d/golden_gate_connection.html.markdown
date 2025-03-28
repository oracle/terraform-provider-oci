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
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Connection. 


## Attributes Reference

The following attributes are exported:

* `access_key_id` - Access key ID to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" 
* `account_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored. Note: When provided, 'accountKey' field must not be provided. 
* `account_name` - Sets the Azure storage account name. 
* `additional_attributes` - An array of name-value pair attribute entries. Used as additional parameters in connection string. 
	* `name` - The name of the property entry. 
	* `value` - The value of the property entry. 
* `authentication_mode` - Authentication mode. It can be provided at creation of Oracle Autonomous Database Serverless connections, when a databaseId is provided. The default value is MTLS.
* `authentication_type` - Used authentication mechanism to be provided for the following connection types:
	* AZURE_DATA_LAKE_STORAGE, ELASTICSEARCH, KAFKA_SCHEMA_REGISTRY, REDIS, SNOWFLAKE
    * JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
    * DATABRICKS - Required fields by authentication types:
      * PERSONAL_ACCESS_TOKEN: username is always 'token', user must enter password
      * OAUTH_M2M: user must enter clientId and clientSecret
* `azure_tenant_id` - Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - The name or address of a host. 
	* `port` - The port of an endpoint usually specified for a connection.
	* `private_ip` - Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.
		The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `client_id` - Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
* `client_secret_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the client secret is stored. Note: When provided, 'clientSecret' field must not be provided. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
* `connection_factory` - The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA'
* `connection_string` 
	* ORACLE: Connect descriptor or Easy Connect Naming method used to connect to a database.
    * MONGODB: MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
    * AZURE_SYNAPSE_ANALYTICS: JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
* `connection_type` - The connection type.
* `consumer_properties` - The base64 encoded content of the consumer.properties file.
* `connection_url` 
	* JAVA_MESSAGE_SERVICE: Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'
    * SNOWFLAKE: JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>'
    * AMAZON_REDSHIFT: Connection URL. e.g.: 'jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb'
    * DATABRICKS: Connection URL. e.g.: 'jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb'
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced.
* `database_name` - The name of the database. 
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - Metadata about this specific object.  
* `does_use_secret_ids` - Indicates that sensitive attributes are provided via Secrets.
* `display_name` - An object's Display Name.
* `endpoint` - Optional Microsoft Fabric service endpoint. Default value: https://onelake.dfs.fabric.microsoft.com
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
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored. The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. Note: When provided, 'password' field must not be provided. 
* `port` - The port of an endpoint usually specified for a connection. 
* `private_ip` - Deprecated: this field will be removed in future versions. Either specify the private IP in the connectionString or host  field, or make sure the host name is resolvable in the target VCN.
	The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `private_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm Note: When provided, 'privateKeyFile' field must not be provided. 
* `private_key_passphrase_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password for the private key file. Note: When provided, 'privateKeyPassphrase' field must not be provided. 
* `producer_properties` - The base64 encoded content of the producer.properties file. 
* `redis_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Redis cluster. 
* `region` - The name of the region. e.g.: us-ashburn-1 If the region is not provided, backend will default to the default region.
* `routing_method` - Controls the network traffic direction to the target: SHARED_SERVICE_ENDPOINT: Traffic flows through the Goldengate Service's network to public hosts. Cannot be used for private targets.  SHARED_DEPLOYMENT_ENDPOINT: Network traffic flows from the assigned deployment's private endpoint through the deployment's subnet. DEDICATED_ENDPOINT: A dedicated private endpoint is created in the target VCN subnet for the connection. The subnetId is required when DEDICATED_ENDPOINT networking is selected.
* `sas_token_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the sas token is stored. Note: When provided, 'sasToken' field must not be provided.
* `secret_access_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the secret access key is stored. Note: When provided, 'secretAccessKey' field must not be provided.
* `security_protocol` - Security Protocol to be provided for the following connection types:
	* ELASTICSEARCH, KAFKA, MICROSOFT_SQLSERVER, MYSQL, POSTGRESQL, REDIS
	* JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
* `servers` - Comma separated list of server addresses, specified as host:port entries, where :port is optional. Example: `"server1.example.com:4000,server2.example.com:4000"`
     If port is not specified, a default value is set, in case of ELASTICSEARCH: 9200, for REDIS 6379.
* `service_account_key_file_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored, which containing the credentials required to use Google Cloud Storage. Note: When provided, 'serviceAccountKeyFile' field must not be provided.
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT.
* `should_use_jndi` - If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
* `should_use_resource_principal` - Indicates that the user intents to connect to the instance through resource principal.
* `should_validate_server_certificate` - If set to true, the driver validates the certificate that is sent by the database server.
* `ssl_ca` - Database Certificate - The base64 encoded content of a .pem or .crt file. containing the server public key (for 1-way SSL).
* `ssl_client_keystash_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystash file is stored,  which contains the encrypted password to the key database file. Note: When provided, 'sslClientKeystash' field must not be provided. 
* `ssl_client_keystoredb_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the keystore file stored,  which created at the client containing the server certificate / CA root certificate. Note: When provided, 'sslClientKeystoredb' field must not be provided. 
* `ssl_key_password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored for the cert inside of the Keystore. In case it differs from the KeyStore password, it should be provided. Note: When provided, 'sslKeyPassword' field must not be provided. 
* `ssl_key_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Client Key
	* The content of a .pem or .crt file containing the client private key (for 2-way SSL). Note: When provided, 'sslKey' field must not be provided.
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `ssl_mode` - SSL mode to be provided for the following connection types: MYSQL, POSTGRESQL.
* `state` - Possible lifecycle states for connection. 
* `storage_credential_name` - Optional. External storage credential name to access files on object storage such as ADLS Gen2, S3 or GCS. 
* `stream_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}`
* `technology_type` - The technology type. 
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `tenant_id` - Azure tenant ID of the application. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
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
