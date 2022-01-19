---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_metastore_config"
sidebar_current: "docs-oci-resource-bds-bds_instance_metastore_config"
description: |-
  Provides the Bds Instance Metastore Config resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_metastore_config
This resource provides the Bds Instance Metastore Config resource in Oracle Cloud Infrastructure Big Data Service service.

Create and activate external metastore configuration.


## Example Usage

```hcl
resource "oci_bds_bds_instance_metastore_config" "test_bds_instance_metastore_config" {
	#Required
	bds_api_key_id = oci_identity_api_key.test_api_key.id
	bds_api_key_passphrase = var.bds_instance_metastore_config_bds_api_key_passphrase
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_metastore_config_cluster_admin_password
	metastore_id = oci_datacatalog_metastore.test_metastore.id

	#Optional
	display_name = var.bds_instance_metastore_config_display_name
}
```

## Argument Reference

The following arguments are supported:

* `bds_api_key_id` - (Required) (Updatable) The ID of BDS Api Key used for Data Catalog metastore integration.
* `bds_api_key_passphrase` - (Required) (Updatable) Base-64 encoded passphrase of the BDS Api Key.
* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) (Updatable) Base-64 encoded password for the cluster admin user.
* `display_name` - (Optional) (Updatable) The display name of the metastore configuration
* `metastore_id` - (Required) The OCID of the Data Catalog metastore.
* `activate_trigger` - (Optional) (Updatable) An optional integer, when flipped triggers activation of metastore config.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_api_key_id` - The ID of BDS API Key used for metastore configuration. Set only if metastore's type is EXTERNAL.
* `display_name` - The display name of metastore configuration
* `id` - The ID of the metastore configuration
* `metastore_id` - The OCID of the Data Catalog metastore. Set only if metastore's type is EXTERNAL.
* `metastore_type` - The type of the metastore in the metastore configuration.
* `state` - the lifecycle state of the metastore configuration.
* `time_created` - The time when the configuration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time when the configuration was updated, shown as an RFC 3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Metastore Config
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Metastore Config
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Metastore Config


## Import

BdsInstanceMetastoreConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config "bdsInstances/{bdsInstanceId}/metastoreConfigs/{metastoreConfigId}" 
```

