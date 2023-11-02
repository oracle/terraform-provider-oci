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
	endpoint = var.connection_endpoint
	fingerprint = var.connection_fingerprint
	freeform_tags = {"bar-key"= "value"}
	host = var.connection_host
	jndi_connection_factory = var.connection_jndi_connection_factory
	jndi_initial_context_factory = var.connection_jndi_initial_context_factory
	jndi_provider_url = var.connection_jndi_provider_url
	jndi_security_credentials = var.connection_jndi_security_credentials
	jndi_security_principal = var.connection_jndi_security_principal
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
	secret_access_key = var.connection_secret_access_key
	security_protocol = var.connection_security_protocol
	servers = var.connection_servers
	service_account_key_file = var.connection_service_account_key_file
	session_mode = var.connection_session_mode
	should_use_jndi = var.connection_should_use_jndi
	should_validate_server_certificate = var.connection_should_validate_server_certificate
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

* `access_key_id` - (Required when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) Access key ID to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" 
* `account_key` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure storage account key. This property is required when 'authenticationType' is set to 'SHARED_KEY'. e.g.: pa3WbhVATzj56xD4DH1VjOUhApRGEGHvOo58eQJVWIzX+j8j4CUVFcTjpIqDSRaSa1Wo2LbWY5at+AStEgLOIQ== 
* `account_name` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Sets the Azure storage account name.
* `additional_attributes` - (Applicable when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) An array of name-value pair attribute entries. Used as additional parameters in connection string.
	* `name` - (Required when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name of the property entry.
	* `value` - (Required when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The value of the property entry.
* `authentication_type` - (Required when connection_type=AZURE_DATA_LAKE_STORAGE | ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA_SCHEMA_REGISTRY | REDIS | SNOWFLAKE) (Updatable) Authentication type for Java Message Service.  If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
* `azure_tenant_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - (Applicable when connection_type=KAFKA) (Updatable) Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"` 
	* `host` - (Required when connection_type=KAFKA) (Updatable) The name or address of a host. 
	* `port` - (Applicable when connection_type=KAFKA) (Updatable) The port of an endpoint usually specified for a connection. 
	* `private_ip` - (Applicable when connection_type=KAFKA) (Updatable) The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `client_id` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d 
* `client_secret` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure client secret (aka application password) for authentication. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA' 
* `connection_string` - (Required when connection_type=AZURE_SYNAPSE_ANALYTICS | MONGODB | ORACLE) (Updatable) Connection string. AZURE_SYNAPSE_ANALYTICS e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;', MONGODB e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'. 
* `connection_type` - (Required) (Updatable) The connection type. 
* `connection_url` - (Required when connection_type=AMAZON_REDSHIFT | JAVA_MESSAGE_SERVICE | SNOWFLAKE) (Updatable) JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>' 
* `consumer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the consumer.properties file. 
* `core_site_xml` - (Required when connection_type=HDFS) (Updatable) The base64 encoded content of the Hadoop Distributed File System configuration file (core-site.xml). 
* `database_id` - (Applicable when connection_type=MONGODB | ORACLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Autonomous Json Database. 
* `database_name` - (Required when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name of the database. 
* `db_system_id` - (Applicable when connection_type=MYSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - (Applicable when connection_type=GOLDENGATE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name.
* `endpoint` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Azure Storage service endpoint. e.g: https://test.blob.core.windows.net 
* `fingerprint` - (Applicable when connection_type=ELASTICSEARCH) (Updatable) Fingerprint required by TLS security protocol. Eg.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c' 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `host` - (Required when connection_type=GENERIC | GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The name or address of a host. In case of Generic connection type host and port separated by colon. Example: `"server.example.com:1234"`
	For multiple hosts, provide a comma separated list. Example: `"server1.example.com:1000,server1.example.com:2000"` 
* `jndi_connection_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory' 
* `jndi_initial_context_factory` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The implementation of javax.naming.spi.InitialContextFactory interface that the client uses to obtain initial naming context. e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory' 
* `jndi_provider_url` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000' 
* `jndi_security_credentials` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) The password associated to the principal. 
* `jndi_security_principal` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) Specifies the identity of the principal (user) to be authenticated. e.g.: 'admin2' 
* `key_id` - (Optional) (Updatable) Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `key_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the KeyStore file. 
* `key_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The KeyStore password. 
* `nsg_ids` - (Optional) (Updatable) An array of Network Security Group OCIDs used to define network access for either Deployments or Connections. 
* `password` - (Required when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. 
* `port` - (Required when connection_type=GOLDENGATE | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) The port of an endpoint usually specified for a connection. 
* `private_ip` - (Applicable when connection_type=GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MYSQL | ORACLE | POSTGRESQL) (Updatable) The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection. 
* `private_key_file` - (Required when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL | SNOWFLAKE) (Updatable) The base64 encoded content of private key file in PEM format. 
* `private_key_passphrase` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL | SNOWFLAKE) (Updatable) Password if the private key file is encrypted. 
* `producer_properties` - (Applicable when connection_type=KAFKA) (Updatable) The base64 encoded content of the producer.properties file. 
* `public_key_fingerprint` - (Required when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL) (Updatable) The fingerprint of the API Key of the user specified by the userId. See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm 
* `region` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL) (Updatable) The name of the region. e.g.: us-ashburn-1 
* `sas_token` - (Applicable when connection_type=AZURE_DATA_LAKE_STORAGE) (Updatable) Credential that uses a shared access signature (SAS) to authenticate to an Azure Service. This property is required when 'authenticationType' is set to 'SHARED_ACCESS_SIGNATURE'. e.g.: ?sv=2020-06-08&ss=bfqt&srt=sco&sp=rwdlacupyx&se=2020-09-10T20:27:28Z&st=2022-08-05T12:27:28Z&spr=https&sig=C1IgHsiLBmTSStYkXXGLTP8it0xBrArcgCqOsZbXwIQ%3D 
* `secret_access_key` - (Required when connection_type=AMAZON_KINESIS | AMAZON_S3) (Updatable) Secret access key to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret" 
* `security_protocol` - (Required when connection_type=ELASTICSEARCH | JAVA_MESSAGE_SERVICE | KAFKA | MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL | REDIS) (Updatable) Security protocol for Java Message Service. If not provided, default is PLAIN. Optional until 2024-06-27, in the release after it will be made required. 
* `servers` - (Required when connection_type=ELASTICSEARCH | REDIS) (Updatable) Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional.  If port is not specified, it defaults to 9200. Used for establishing the initial connection to the Elasticsearch cluster. Example: `"server1.example.com:4000,server2.example.com:4000"` 
* `service_account_key_file` - (Required when connection_type=GOOGLE_BIGQUERY | GOOGLE_CLOUD_STORAGE) (Updatable) The base64 encoded content of the service account key file containing the credentials required to use Google Cloud Storage. 
* `session_mode` - (Applicable when connection_type=ORACLE) (Updatable) The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `should_use_jndi` - (Required when connection_type=JAVA_MESSAGE_SERVICE) (Updatable) If set to true, Java Naming and Directory Interface (JNDI) properties should be provided. 
* `should_validate_server_certificate` - (Applicable when connection_type=MICROSOFT_SQLSERVER) (Updatable) If set to true, the driver validates the certificate that is sent by the database server. 
* `ssl_ca` - (Applicable when connection_type=MICROSOFT_SQLSERVER | MYSQL | POSTGRESQL) (Updatable) Database Certificate - The base64 encoded content of pem file containing the server public key (for 1-way SSL). 
* `ssl_cert` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Certificate - The base64 encoded content of client-cert.pem file  containing the client public key (for 2-way SSL). 
* `ssl_crl` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Certificates revoked by certificate authorities (CA). Server certificate must not be on this list (for 1 and 2-way SSL). Note: This is an optional and that too only applicable if TLS/MTLS option is selected. 
* `ssl_key` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) Client Key - The client-key.pem containing the client private key (for 2-way SSL). 
* `ssl_key_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY) (Updatable) The password for the cert inside of the KeyStore. In case it differs from the KeyStore password, it should be provided. 
* `ssl_mode` - (Applicable when connection_type=MYSQL | POSTGRESQL) (Updatable) SSL modes for PostgreSQL.
* `stream_pool_id` - (Applicable when connection_type=KAFKA) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced. 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `technology_type` - (Required) The Kafka (e.g. Confluent) Schema Registry technology type. 
* `tenancy_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy. 
* `trust_store` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The base64 encoded content of the TrustStore file. 
* `trust_store_password` - (Applicable when connection_type=JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | REDIS) (Updatable) The TrustStore password. 
* `url` - (Required when connection_type=KAFKA_SCHEMA_REGISTRY) (Updatable) Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081' 
* `user_id` - (Applicable when connection_type=OCI_OBJECT_STORAGE | ORACLE_NOSQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database. The user must have write access to the table they want to connect to. 
* `username` - (Required when connection_type=AMAZON_REDSHIFT | AZURE_SYNAPSE_ANALYTICS | ELASTICSEARCH | GOLDENGATE | JAVA_MESSAGE_SERVICE | KAFKA | KAFKA_SCHEMA_REGISTRY | MICROSOFT_SQLSERVER | MONGODB | MYSQL | ORACLE | POSTGRESQL | REDIS | SNOWFLAKE) (Updatable) The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it. 
* `vault_id` - (Optional) (Updatable) Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault. 
* `wallet` - (Applicable when connection_type=ORACLE) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded. 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_key_id` - Access key ID to access the Amazon S3 bucket. e.g.: "this-is-not-the-secret"
* `account_name` - Sets the Azure storage account name.
* `additional_attributes` - An array of name-value pair attribute entries. Used as additional parameters in connection string.
	* `name` - The name of the property entry.
	* `value` - The value of the property entry.
* `authentication_type` - Used authentication mechanism to be provided for the following connection types:
	* SNOWFLAKE, AZURE_DATA_LAKE_STORAGE, ELASTICSEARCH, KAFKA_SCHEMA_REGISTRY, REDIS
	* JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
* `azure_tenant_id` - Azure tenant ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 14593954-d337-4a61-a364-9f758c64f97f
* `bootstrap_servers` - Kafka bootstrap. Equivalent of bootstrap.servers configuration property in Kafka: list of KafkaBootstrapServer objects specified by host/port. Used for establishing the initial connection to the Kafka cluster. Example: `"server1.example.com:9092,server2.example.com:9092"`
	* `host` - The name or address of a host.
	* `port` - The port of an endpoint usually specified for a connection.
	* `private_ip` - The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `client_id` - Azure client ID of the application. This property is required when 'authenticationType' is set to 'AZURE_ACTIVE_DIRECTORY'. e.g.: 06ecaabf-8b80-4ec8-a0ec-20cbf463703d
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
* `connection_factory` - The of Java class implementing javax.jms.ConnectionFactory interface supplied by the Java Message Service provider. e.g.: 'com.stc.jmsjca.core.JConnectionFactoryXA'
* `connection_string`
	* ORACLE: Connect descriptor or Easy Connect Naming method used to connect to a database.
	* MONGODB: MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/records'
	* AZURE_SYNAPSE_ANALYTICS: JDBC connection string. e.g.: 'jdbc:sqlserver://<synapse-workspace>.sql.azuresynapse.net:1433;database=<db-name>;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'
* `connection_type` - The connection type.
* `connection_url`
	* JAVA_MESSAGE_SERVICE: Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'
	* SNOWFLAKE: JDBC connection URL. e.g.: 'jdbc:snowflake://<account_name>.snowflakecomputing.com/?warehouse=<warehouse-name>&db=<db-name>'
	* AMAZON_REDSHIFT: Connection URL. e.g.: 'jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb'
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
  In case of Generic connection type it represents the Host and port separated by colon. Example: `"server.example.com:1234"`
  For multiple hosts, provide a comma separated list. Example: `"server1.example.com:1000,server1.example.com:2000"`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced.
* `ingress_ips` - List of ingress IP addresses from where the GoldenGate deployment connects to this connection's privateIp.  Customers may optionally set up ingress security rules to restrict traffic from these IP addresses.
	* `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet.
* `jndi_connection_factory` - The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory'
* `jndi_initial_context_factory` - The implementation of javax.naming.spi.InitialContextFactory interface that the client uses to obtain initial naming context. e.g.: 'org.apache.activemq.jndi.ActiveMQInitialContextFactory'
* `jndi_provider_url` - The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000'
* `jndi_security_principal` - Specifies the identity of the principal (user) to be authenticated. e.g.: 'admin2'
* `key_id` - Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
* `port` - The port of an endpoint usually specified for a connection.
* `private_ip` - The private IP address of the connection's endpoint in the customer's VCN, typically a database endpoint or a big data endpoint (e.g. Kafka bootstrap server). In case the privateIp is provided, the subnetId must also be provided. In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible. In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
* `region` - The name of the region. e.g.: us-ashburn-1
* `security_protocol` - Security Protocol to be provided for the following connection types:
	ELASTICSEARCH, KAFKA, MICROSOFT_SQLSERVER, MYSQL, POSTGRESQL, REDIS
	* JAVA_MESSAGE_SERVICE - If not provided, default is NONE. Optional until 2024-06-27, in the release after it will be made required.
* `servers` - Comma separated list of server addresses, specified as host:port entries, where :port is optional. Example: `"server1.example.com:4000,server2.example.com:4000"`
  If port is not specified, a default value is set, in case of ELASTICSEARCH: 9200, for REDIS 6379.
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT.
* `should_use_jndi` - If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
* `should_validate_server_certificate` - If set to true, the driver validates the certificate that is sent by the database server.
* `ssl_ca` - Database Certificate - The base64 encoded content of pem file containing the server public key (for 1-way SSL).
* `ssl_mode` - SSL mode to be provided for the following connection types: MYSQL, POSTGRESQL.
* `state` - Possible lifecycle states for connection.
* `stream_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream pool being referenced.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}`
* `technology_type` - The technology type.
* `tenancy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Oracle Cloud Infrastructure tenancy.
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
* `url` - Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081'
* `user_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database/Object Storage. The user must have write access.
* `username` - The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivity requirements defined in it.
* `vault_id` - Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault.

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

