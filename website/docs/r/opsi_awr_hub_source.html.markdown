---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub_source"
sidebar_current: "docs-oci-resource-opsi-awr_hub_source"
description: |-
  Provides the Awr Hub Source resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_awr_hub_source
This resource provides the Awr Hub Source resource in Oracle Cloud Infrastructure Opsi service.

Register Awr Hub source


## Example Usage

```hcl
resource "oci_opsi_awr_hub_source" "test_awr_hub_source" {
	#Required
	awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
	compartment_id = var.compartment_id
	name = var.awr_hub_source_name
	type = var.awr_hub_source_type

	#Optional
	associated_opsi_id = oci_opsi_associated_opsi.test_associated_opsi.id
	associated_resource_id = oci_usage_proxy_resource.test_resource.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `associated_opsi_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `associated_resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `awr_hub_id` - (Required) AWR Hub OCID
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) The name of the Awr Hub source database.
* `type` - (Required) (Updatable) source type of the database


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `associated_opsi_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `associated_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `awr_hub_id` - AWR Hub OCID
* `awr_hub_opsi_source_id` - The shorted string of the Awr Hub source database identifier.
* `awr_source_database_id` - DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hours_since_last_import` - Number of hours since last AWR snapshots import happened from the Source database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Awr Hub source database.
* `is_registered_with_awr_hub` - This is `true` if the source databse is registered with a Awr Hub, otherwise `false`
* `max_snapshot_identifier` - The maximum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
* `min_snapshot_identifier` - The minimum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
* `name` - The name of the Awr Hub source database.
* `source_mail_box_url` - Opsi Mailbox URL based on the Awr Hub and Awr Hub source.
* `state` - the current state of the source database
* `status` - Indicates the status of a source database in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_first_snapshot_generated` - The time at which the earliest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
* `time_last_snapshot_generated` - The time at which the latest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string
* `type` - source type of the database

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Awr Hub Source
	* `update` - (Defaults to 20 minutes), when updating the Awr Hub Source
	* `delete` - (Defaults to 20 minutes), when destroying the Awr Hub Source


## Import

AwrHubSources can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_awr_hub_source.test_awr_hub_source "id"
```

