---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_cluster_instance"
sidebar_current: "docs-oci-datasource-database_management-external_cluster_instance"
description: |-
  Provides details about a specific External Cluster Instance in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_cluster_instance
This data source provides details about a specific External Cluster Instance resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the external cluster instance specified by `externalClusterInstanceId`.


## Example Usage

```hcl
data "oci_database_management_external_cluster_instance" "test_external_cluster_instance" {
	#Required
	external_cluster_instance_id = oci_database_management_external_cluster_instance.test_external_cluster_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `external_cluster_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external cluster instance.


## Attributes Reference

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

