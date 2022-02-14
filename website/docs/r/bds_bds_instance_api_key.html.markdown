---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_api_key"
sidebar_current: "docs-oci-resource-bds-bds_instance_api_key"
description: |-
  Provides the Bds Instance Api Key resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_api_key
This resource provides the Bds Instance Api Key resource in Oracle Cloud Infrastructure Big Data Service service.

Create an API key on behalf of the specified user.


## Example Usage

```hcl
resource "oci_bds_bds_instance_api_key" "test_bds_instance_api_key" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	key_alias = var.bds_instance_api_key_key_alias
	passphrase = var.bds_instance_api_key_passphrase
	user_id = oci_identity_user.test_user.id

	#Optional
	default_region = var.bds_instance_api_key_default_region
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `default_region` - (Optional) The name of the region to establish the Object Storage endpoint. See https://docs.oracle.com/en-us/iaas/api/#/en/identity/20160918/Region/ for additional information.
* `key_alias` - (Required) User friendly identifier used to uniquely differentiate between different API keys associated with this Big Data Service cluster. Only ASCII alphanumeric characters with no spaces allowed.
* `passphrase` - (Required) Base64 passphrase used to secure the private key which will be created on user behalf.
* `user_id` - (Required) The OCID of the user for whom this new generated API key pair will be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Bds Instance Api Key
* `update` - (Defaults to 20 minutes), when updating the Bds Instance Api Key
* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Api Key


## Import

BdsInstanceApiKeys can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_api_key.test_bds_instance_api_key "bdsInstances/{bdsInstanceId}/apiKeys/{apiKeyId}" 
```
