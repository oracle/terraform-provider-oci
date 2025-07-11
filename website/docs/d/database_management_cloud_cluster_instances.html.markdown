---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_cluster_instances"
sidebar_current: "docs-oci-datasource-database_management-cloud_cluster_instances"
description: |-
  Provides the list of Cloud Cluster Instances in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_cluster_instances
This data source provides the list of Cloud Cluster Instances in Oracle Cloud Infrastructure Database Management service.

Lists the cluster instances in the specified cloud cluster.

## Example Usage

```hcl
data "oci_database_management_cloud_cluster_instances" "test_cloud_cluster_instances" {

	#Optional
	cloud_cluster_id = oci_database_management_cloud_cluster.test_cloud_cluster.id
	compartment_id = var.compartment_id
	display_name = var.cloud_cluster_instance_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cloud_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `cloud_cluster_instance_collection` - The list of cloud_cluster_instance_collection.

### CloudClusterInstance Reference

The following attributes are exported:

* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
* `cloud_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster that the cluster instance belongs to.
* `cloud_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the cluster instance is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud cluster instance.
* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cluster instance. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the cluster instance is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud cluster instance.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `node_role` - The role of the cluster node.
* `state` - The current lifecycle state of the cloud cluster instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud cluster instance was created.
* `time_updated` - The date and time the cloud cluster instance was last updated.

