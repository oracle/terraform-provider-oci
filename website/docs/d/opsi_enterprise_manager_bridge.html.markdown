---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_enterprise_manager_bridge"
sidebar_current: "docs-oci-datasource-opsi-enterprise_manager_bridge"
description: |-
  Provides details about a specific Enterprise Manager Bridge in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_enterprise_manager_bridge
This data source provides details about a specific Enterprise Manager Bridge resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an Operations Insights Enterprise Manager bridge.

## Example Usage

```hcl
data "oci_opsi_enterprise_manager_bridge" "test_enterprise_manager_bridge" {
	#Required
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_manager_bridge_id` - (Required) Unique Enterprise Manager bridge identifier


## Attributes Reference

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

