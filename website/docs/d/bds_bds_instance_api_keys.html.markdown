---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_api_key"
sidebar_current: "docs-oci-datasource-bds-bds_instance_api_key"
description: |-
Provides details about a specific Bds Instance Api Key in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_api_key
This data source provides details about a specific Bds Instance Api Key resource in Oracle Cloud Infrastructure Big Data Service service.

Returns the user's API key information for the given ID.

## Example Usage

```hcl
data "oci_bds_bds_instance_api_key" "test_bds_instance_api_key" {
	#Required
	api_key_id = oci_identity_api_key.test_api_key.id
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `api_key_id` - (Required) The API key identifier.
* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

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
