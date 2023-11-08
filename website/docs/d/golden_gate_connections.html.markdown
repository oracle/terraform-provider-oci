---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connections"
sidebar_current: "docs-oci-datasource-golden_gate-connections"
description: |-
  Provides the list of Connections in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_connections
This data source provides the list of Connections in Oracle Cloud Infrastructure Golden Gate service.

Lists the Connections in the compartment.


## Example Usage

```hcl
data "oci_golden_gate_connections" "test_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	assignable_deployment_id = oci_golden_gate_deployment.test_deployment.id
	assignable_deployment_type = var.connection_assignable_deployment_type
	assigned_deployment_id = oci_golden_gate_deployment.test_deployment.id
	connection_type = var.connection_connection_type
	display_name = var.connection_display_name
	state = var.connection_state
	technology_type = var.connection_technology_type
}
```

## Argument Reference

The following arguments are supported:

* `assignable_deployment_id` - (Optional) Filters for compatible connections which can be, but currently not assigned to the deployment specified by its id. 
* `assignable_deployment_type` - (Optional) Filters for connections which can be assigned to the latest version of the specified deployment type. 
* `assigned_deployment_id` - (Optional) The OCID of the deployment which for the connection must be assigned. 
* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `connection_type` - (Optional) The array of connection types. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `state` - (Optional) A filter to return only connections having the 'lifecycleState' given. 
* `technology_type` - (Optional) The array of technology types. 


## Attributes Reference

The following attributes are exported:

* `connection_collection` - The list of connection_collection.

### Connection Reference

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
    * MONGODB: MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'
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
	* ELASTICSEARCH, KAFKA, MICROSOFT_SQLSERVER, MYSQL, POSTGRESQL, REDIS
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
* `user_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure user who will access the Oracle NoSQL database/Object Storage. The user must have write access to the table they want to connect to.
* `username` - The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivity requirements defined in it.
* `vault_id` - Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault.
