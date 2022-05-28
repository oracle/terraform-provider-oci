---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_configuration"
sidebar_current: "docs-oci-resource-license_manager-configuration"
description: |-
  Provides the Configuration resource in Oracle Cloud Infrastructure License Manager service
---

# oci_license_manager_configuration
This resource provides the Configuration resource in Oracle Cloud Infrastructure License Manager service.

Updates the configuration for the compartment.

## Example Usage

```hcl
resource "oci_license_manager_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
	email_ids = var.configuration_email_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration. 
* `email_ids` - (Required) (Updatable) List of email IDs associated with the configuration.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the configuration is specified. 
* `email_ids` - The list of associated configuration email IDs.
* `time_created` - The time the configuration was created. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `time_updated` - The time the configuration was updated. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Configuration
	* `update` - (Defaults to 20 minutes), when updating the Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Configuration


## Import

Configurations can be imported using the `id`, e.g.

```
$ terraform import oci_license_manager_configuration.test_configuration "configuration/compartmentId/{compartmentId}" 
```

