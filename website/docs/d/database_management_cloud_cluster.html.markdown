---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_cluster"
sidebar_current: "docs-oci-datasource-database_management-cloud_cluster"
description: |-
  Provides details about a specific Cloud Cluster in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_cluster
This data source provides details about a specific Cloud Cluster resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the cloud cluster specified by `cloudClusterId`.


## Example Usage

```hcl
data "oci_database_management_cloud_cluster" "test_cloud_cluster" {
	#Required
	cloud_cluster_id = oci_database_management_cloud_cluster.test_cloud_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster.


## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the cloud cluster defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `cloud_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the cluster is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud cluster.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud cluster. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `grid_home` - The directory in which Oracle Grid Infrastructure is installed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster.
* `is_flex_cluster` - Indicates whether the cluster is Oracle Flex Cluster or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `network_configurations` - The list of network address configurations of the cloud cluster.
	* `network_number` - The network number.
	* `network_type` - The network type.
	* `subnet` - The subnet for the network.
* `ocr_file_location` - The location of the Oracle Cluster Registry (OCR).
* `scan_configurations` - The list of Single Client Access Name (SCAN) configurations of the cloud cluster.
	* `network_number` - The network number from which SCAN VIPs are obtained.
	* `scan_name` - The name of the SCAN listener.
	* `scan_port` - The port number of the SCAN listener.
	* `scan_protocol` - The protocol of the SCAN listener.
* `state` - The current lifecycle state of the cloud cluster.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud cluster was created.
* `time_updated` - The date and time the cloud cluster was last updated.
* `version` - The cluster version.
* `vip_configurations` - The list of Virtual IP (VIP) configurations of the cloud cluster.
	* `address` - The VIP name or IP address.
	* `network_number` - The network number from which VIPs are obtained.
	* `node_name` - The name of the node with the VIP.

