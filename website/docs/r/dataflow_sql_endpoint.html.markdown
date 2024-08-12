---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_sql_endpoint"
sidebar_current: "docs-oci-resource-dataflow-sql_endpoint"
description: |-
  Provides the Sql Endpoint resource in Oracle Cloud Infrastructure Data Flow service
---

# oci_dataflow_sql_endpoint
This resource provides the Sql Endpoint resource in Oracle Cloud Infrastructure Data Flow service.
## Note
Resource Discovery is not supported for this resource.

Create a new Sql Endpoint.

## Example Usage

```hcl
resource "oci_dataflow_sql_endpoint" "test_sql_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.sql_endpoint_display_name
	driver_shape = var.sql_endpoint_driver_shape
	executor_shape = var.sql_endpoint_executor_shape
	lake_id = oci_dataflow_lake.test_lake.id
	max_executor_count = var.sql_endpoint_max_executor_count
	metastore_id = oci_datacatalog_metastore.test_metastore.id
	min_executor_count = var.sql_endpoint_min_executor_count
	network_configuration {
		#Required
		network_type = var.sql_endpoint_network_configuration_network_type

		#Optional
		access_control_rules {

			#Optional
			ip_notation = var.sql_endpoint_network_configuration_access_control_rules_ip_notation
			value = var.sql_endpoint_network_configuration_access_control_rules_value
			vcn_ips = var.sql_endpoint_network_configuration_access_control_rules_vcn_ips
		}
		host_name_prefix = var.sql_endpoint_network_configuration_host_name_prefix
		nsg_ids = var.sql_endpoint_network_configuration_nsg_ids
		private_endpoint_ip = var.sql_endpoint_network_configuration_private_endpoint_ip
		public_endpoint_ip = var.sql_endpoint_network_configuration_public_endpoint_ip
		subnet_id = oci_core_subnet.test_subnet.id
		vcn_id = oci_core_vcn.test_vcn.id
	}
	sql_endpoint_version = var.sql_endpoint_sql_endpoint_version
	warehouse_bucket_uri = var.sql_endpoint_warehouse_bucket_uri

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sql_endpoint_description
	driver_shape_config {

		#Optional
		memory_in_gbs = var.sql_endpoint_driver_shape_config_memory_in_gbs
		ocpus = var.sql_endpoint_driver_shape_config_ocpus
	}
	executor_shape_config {

		#Optional
		memory_in_gbs = var.sql_endpoint_executor_shape_config_memory_in_gbs
		ocpus = var.sql_endpoint_executor_shape_config_ocpus
	}
	freeform_tags = {"Department"= "Finance"}
	spark_advanced_configurations = var.sql_endpoint_spark_advanced_configurations
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The identifier of the compartment used with the SQL Endpoint.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of CreateSQLEndpointDetails.
* `display_name` - (Required) (Updatable) The SQL Endpoint name, which can be changed.
* `driver_shape` - (Required) The shape of the SQL Endpoint driver instance.
* `driver_shape_config` - (Optional) This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - (Optional) The amount of memory used for the driver or executors. 
	* `ocpus` - (Optional) The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `executor_shape` - (Required) The shape of the SQL Endpoint worker instance.
* `executor_shape_config` - (Optional) This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - (Optional) The amount of memory used for the driver or executors. 
	* `ocpus` - (Optional) The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `lake_id` - (Required) Oracle Cloud Infrastructure lake OCID
* `max_executor_count` - (Required) (Updatable) The maximum number of executors.
* `metastore_id` - (Required) Metastore OCID
* `min_executor_count` - (Required) (Updatable) The minimum number of executors.
* `network_configuration` - (Required) The network configuration of a SQL Endpoint.
	* `access_control_rules` - (Applicable when network_type=SECURE_ACCESS) A list of SecureAccessControlRule's to which access is limited to
		* `ip_notation` - (Required when network_type=SECURE_ACCESS) The type of IP notation.
		* `value` - (Required when network_type=SECURE_ACCESS) The associated value of the selected IP notation.
		* `vcn_ips` - (Applicable when network_type=SECURE_ACCESS) A comma-separated IP or CIDR address for VCN OCID IP notation selection.
	* `host_name_prefix` - (Applicable when network_type=VCN) The host name prefix.
	* `network_type` - (Required) The type of network configuration.
	* `nsg_ids` - (Applicable when network_type=VCN) The OCIDs of Network Security Groups (NSGs).
	* `private_endpoint_ip` - (Applicable when network_type=VCN) Ip Address of private endpoint
	* `public_endpoint_ip` - (Applicable when network_type=SECURE_ACCESS) Ip Address of public endpoint
	* `subnet_id` - (Required when network_type=VCN) The VCN Subnet OCID.
	* `vcn_id` - (Required when network_type=VCN) The VCN OCID.
* `spark_advanced_configurations` - (Optional) (Updatable) The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `sql_endpoint_version` - (Required) The version of the SQL Endpoint.
* `warehouse_bucket_uri` - (Required) The warehouse bucket URI. It is a Oracle Cloud Infrastructure Object Storage bucket URI as defined here https://docs.oracle.com/en/cloud/paas/atp-cloud/atpud/object-storage-uris.html
* `state` - (Optional) (Updatable) The target state for the Sql Endpoint. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sql Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Sql Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Sql Endpoint


## Import

SqlEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_dataflow_sql_endpoint.test_sql_endpoint "id"
```

