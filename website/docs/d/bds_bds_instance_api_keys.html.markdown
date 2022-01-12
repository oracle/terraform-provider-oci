---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_api_keys"
sidebar_current: "docs-oci-datasource-bds-bds_instance_api_keys"
description: |-
  Provides the list of Bds Instance Api Keys in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_api_keys
This data source provides the list of Bds Instance Api Keys in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of all API keys associated with this Big Data Service cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_api_keys" "test_bds_instance_api_keys" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	display_name = var.bds_instance_api_key_display_name
	state = var.bds_instance_api_key_state
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The state of the API key.
* `user_id` - (Optional) The OCID of the user for whom the API key belongs.


## Attributes Reference

The following attributes are exported:

* `bds_api_keys` - The list of bds_api_keys.

### BdsInstanceApiKey Reference

The following attributes are exported:

* `default_region` - The name of the region to establish the Object Storage endpoint which was set as part of key creation operation. If no region was provided this will be set to be the same region where the cluster lives. Example us-phoenix-1 . 
* `fingerprint` - The fingerprint that corresponds to the public API key requested.
* `id` - Identifier of the user's API key.
* `key_alias` - User friendly identifier used to uniquely differentiate between different API keys. Only ASCII alphanumeric characters with no spaces allowed. 
* `pemfilepath` - The full path and file name of the private key used for authentication. This location will be automatically selected on the BDS local file system. 
* `state` - The current status of the API key.
* `tenant_id` - The OCID of your tenancy.
* `time_created` - The time the API key was created, shown as an RFC 3339 formatted datetime string.
* `user_id` - The user OCID for which this API key was created.

