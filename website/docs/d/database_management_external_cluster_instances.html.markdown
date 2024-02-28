---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_cluster_instances"
sidebar_current: "docs-oci-datasource-database_management-external_cluster_instances"
description: |-
  Provides the list of External Cluster Instances in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_cluster_instances
This data source provides the list of External Cluster Instances in Oracle Cloud Infrastructure Database Management service.

Lists the cluster instances in the specified external cluster.

## Example Usage

```hcl
data "oci_database_management_external_cluster_instances" "test_external_cluster_instances" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.external_cluster_instance_display_name
	external_cluster_id = oci_database_management_external_cluster.test_external_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.
* `external_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster.


## Attributes Reference

The following attributes are exported:

* `external_cluster_instance_collection` - The list of external_cluster_instance_collection.

### ExternalClusterInstance Reference

The following attributes are exported:

* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external cluster instance.
* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cluster instance. The name does not have to be unique.
* `external_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster that the cluster instance belongs to.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the cluster instance is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the cluster instance is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster instance.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `node_role` - The role of the cluster node.
* `state` - The current lifecycle state of the external cluster instance.
* `time_created` - The date and time the external cluster instance was created.
* `time_updated` - The date and time the external cluster instance was last updated.

