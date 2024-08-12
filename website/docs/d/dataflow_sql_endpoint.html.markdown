---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_sql_endpoint"
sidebar_current: "docs-oci-datasource-dataflow-sql_endpoint"
description: |-
  Provides details about a specific Sql Endpoint in Oracle Cloud Infrastructure Data Flow service
---

# Data Source: oci_dataflow_sql_endpoint
This data source provides details about a specific Sql Endpoint resource in Oracle Cloud Infrastructure Data Flow service.

Retrieves a SQL Endpoint using a sqlEndpointId.

## Example Usage

```hcl
data "oci_dataflow_sql_endpoint" "test_sql_endpoint" {
	#Required
	sql_endpoint_id = oci_dataflow_sql_endpoint.test_sql_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `sql_endpoint_id` - (Required) The unique id of the SQL Endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of a compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the SQL Endpoint.
* `display_name` - The SQL Endpoint name, which can be changed.
* `driver_shape` - The shape of the SQL Endpoint driver instance.
* `driver_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - The amount of memory used for the driver or executors. 
	* `ocpus` - The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `executor_shape` - The shape of the SQL Endpoint executor instance.
* `executor_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - The amount of memory used for the driver or executors. 
	* `ocpus` - The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The provision identifier that is immutable on creation.
* `jdbc_endpoint_url` - The JDBC URL field. For example, jdbc:spark://{serviceFQDN}:443/default;SparkServerType=DFI
* `lake_id` - The OCID of Oracle Cloud Infrastructure Lake.
* `max_executor_count` - The maximum number of executors.
* `metastore_id` - The OCID of Oracle Cloud Infrastructure Hive Metastore. 
* `min_executor_count` - The minimum number of executors.
* `network_configuration` - The network configuration of a SQL Endpoint.
	* `access_control_rules` - A list of SecureAccessControlRule's to which access is limited to
		* `ip_notation` - The type of IP notation.
		* `value` - The associated value of the selected IP notation.
		* `vcn_ips` - A comma-separated IP or CIDR address for VCN OCID IP notation selection.
	* `host_name_prefix` - The host name prefix.
	* `network_type` - The type of network configuration.
	* `nsg_ids` - The OCIDs of Network Security Groups (NSGs).
	* `private_endpoint_ip` - Ip Address of private endpoint
	* `public_endpoint_ip` - Ip Address of public endpoint
	* `subnet_id` - The VCN Subnet OCID.
	* `vcn_id` - The VCN OCID.
* `spark_advanced_configurations` - The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `sql_endpoint_version` - The version of SQL Endpoint.
* `state` - The current state of the Sql Endpoint.
* `state_message` - A message describing the reason why the resource is in it's current state. Helps bubble up errors in state changes. For example, it can be used to provide actionable information for a resource in the Failed state.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the Sql Endpoint was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Sql Endpoint was updated. An RFC3339 formatted datetime string.
* `warehouse_bucket_uri` - The warehouse bucket URI. It is a Oracle Cloud Infrastructure Object Storage bucket URI as defined here https://docs.oracle.com/en/cloud/paas/atp-cloud/atpud/object-storage-uris.html

