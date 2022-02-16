---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_metastore_configs"
sidebar_current: "docs-oci-datasource-bds-bds_instance_metastore_configs"
description: |-
  Provides the list of Bds Instance Metastore Configs in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_metastore_configs
This data source provides the list of Bds Instance Metastore Configs in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of metastore configurations ssociated with this Big Data Service cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_metastore_configs" "test_bds_instance_metastore_configs" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	bds_api_key_id = oci_identity_api_key.test_api_key.id
	display_name = var.bds_instance_metastore_config_display_name
	metastore_id = oci_datacatalog_metastore.test_metastore.id
	metastore_type = var.bds_instance_metastore_config_metastore_type
	state = var.bds_instance_metastore_config_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_api_key_id` - (Optional) The ID of the API key that is associated with the external metastore in the metastore configuration
* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `metastore_id` - (Optional) The OCID of the Data Catalog metastore in the metastore configuration
* `metastore_type` - (Optional) The type of the metastore in the metastore configuration
* `state` - (Optional) The lifecycle state of the metastore in the metastore configuration


## Attributes Reference

The following attributes are exported:

* `bds_metastore_configurations` - The list of bds_metastore_configurations.

### BdsInstanceMetastoreConfig Reference

The following attributes are exported:

* `bds_api_key_id` - The ID of BDS API Key used for metastore configuration. Set only if metastore's type is EXTERNAL.
* `display_name` - The display name of metastore configuration
* `id` - The ID of the metastore configuration
* `metastore_id` - The OCID of the Data Catalog metastore. Set only if metastore's type is EXTERNAL.
* `metastore_type` - The type of the metastore in the metastore configuration.
* `state` - the lifecycle state of the metastore configuration.
* `time_created` - The time when the configuration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time when the configuration was updated, shown as an RFC 3339 formatted datetime string.

