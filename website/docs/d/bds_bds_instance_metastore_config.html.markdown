---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_metastore_config"
sidebar_current: "docs-oci-datasource-bds-bds_instance_metastore_config"
description: |-
  Provides details about a specific Bds Instance Metastore Config in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_metastore_config
This data source provides details about a specific Bds Instance Metastore Config resource in Oracle Cloud Infrastructure Big Data Service service.

Returns the BDS Metastore configuration information for the given ID.

## Example Usage

```hcl
data "oci_bds_bds_instance_metastore_config" "test_bds_instance_metastore_config" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	metastore_config_id = oci_apm_config_config.test_config.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `metastore_config_id` - (Required) The metastore configuration ID


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

