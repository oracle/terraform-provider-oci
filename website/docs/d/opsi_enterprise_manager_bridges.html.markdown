---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_enterprise_manager_bridges"
sidebar_current: "docs-oci-datasource-opsi-enterprise_manager_bridges"
description: |-
  Provides the list of Enterprise Manager Bridges in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_enterprise_manager_bridges
This data source provides the list of Enterprise Manager Bridges in Oracle Cloud Infrastructure Opsi service.

Gets a list of Operations Insights Enterprise Manager bridges. Either compartmentId or id must be specified.
When both compartmentId and compartmentIdInSubtree are specified, a list of bridges in that compartment and in all sub-compartments will be returned.


## Example Usage

```hcl
data "oci_opsi_enterprise_manager_bridges" "test_enterprise_manager_bridges" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.enterprise_manager_bridge_compartment_id_in_subtree
	display_name = var.enterprise_manager_bridge_display_name
	id = var.enterprise_manager_bridge_id
	state = var.enterprise_manager_bridge_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name.
* `id` - (Optional) Unique Enterprise Manager bridge identifier
* `state` - (Optional) Lifecycle states


## Attributes Reference

The following attributes are exported:

* `enterprise_manager_bridge_collection` - The list of enterprise_manager_bridge_collection.

### EnterpriseManagerBridge Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the Enterprise Manager bridge
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of Enterprise Manager Bridge
* `display_name` - User-friedly name of Enterprise Manager Bridge that does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Enterprise Manager bridge identifier
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `object_storage_bucket_name` - Object Storage Bucket Name
* `object_storage_bucket_status_details` - A message describing status of the object storage bucket of this resource. For example, it can be used to provide actionable information about the permission and content validity of the bucket.
* `object_storage_namespace_name` - Object Storage Namespace Name
* `state` - The current state of the Enterprise Manager bridge.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Enterprise Manager bridge was first created. An RFC3339 formatted datetime string
* `time_updated` - The time the Enterprise Manager bridge was updated. An RFC3339 formatted datetime string

