---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_cluster"
sidebar_current: "docs-oci-resource-database_management-external_cluster"
description: |-
  Provides the External Cluster resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_cluster
This resource provides the External Cluster resource in Oracle Cloud Infrastructure Database Management service.

Updates the external cluster specified by `externalClusterId`.


## Example Usage

```hcl
resource "oci_database_management_external_cluster" "test_external_cluster" {
	#Required
	external_cluster_id = oci_database_management_external_cluster.test_external_cluster.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	external_connector_id = oci_database_management_external_connector.test_external_connector.id
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `external_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster.
* `external_connector_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the external cluster defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the external cluster. The name does not have to be unique.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the cluster is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `grid_home` - The directory in which Oracle Grid Infrastructure is installed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster.
* `is_flex_cluster` - Indicates whether the cluster is Oracle Flex Cluster or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `network_configurations` - The list of network address configurations of the external cluster.
	* `network_number` - The network number.
	* `network_type` - The network type.
	* `subnet` - The subnet for the network.
* `ocr_file_location` - The location of the Oracle Cluster Registry (OCR).
* `scan_configurations` - The list of Single Client Access Name (SCAN) configurations of the external cluster.
	* `network_number` - The network number from which SCAN VIPs are obtained.
	* `scan_name` - The name of the SCAN listener.
	* `scan_port` - The port number of the SCAN listener.
	* `scan_protocol` - The protocol of the SCAN listener.
* `state` - The current lifecycle state of the external cluster.
* `time_created` - The date and time the external cluster was created.
* `time_updated` - The date and time the external cluster was last updated.
* `version` - The cluster version.
* `vip_configurations` - The list of Virtual IP (VIP) configurations of the external cluster.
	* `address` - The VIP name or IP address.
	* `network_number` - The network number from which VIPs are obtained.
	* `node_name` - The name of the node with the VIP.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Cluster
	* `update` - (Defaults to 20 minutes), when updating the External Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the External Cluster


## Import

ExternalClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_cluster.test_external_cluster "id"
```

