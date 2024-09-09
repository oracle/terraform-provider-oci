---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_generate_on_prem_connector_configuration"
sidebar_current: "docs-oci-resource-data_safe-generate_on_prem_connector_configuration"
description: |-
  Provides the Generate On Prem Connector Configuration resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_generate_on_prem_connector_configuration
This resource provides the Generate On Prem Connector Configuration resource in Oracle Cloud Infrastructure Data Safe service.

Creates and downloads the configuration of the specified on-premises connector.


## Example Usage

```hcl
resource "oci_data_safe_generate_on_prem_connector_configuration" "test_generate_on_prem_connector_configuration" {
	#Required
	on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
	password = var.generate_on_prem_connector_configuration_password
}
```

## Argument Reference

The following arguments are supported:

* `on_prem_connector_id` - (Required) The OCID of the on-premises connector.
* `password` - (Required) The password to encrypt the keys inside the wallet included as part of the configuration. The password must be between 12 and 30 characters long and must contain atleast 1 uppercase, 1 lowercase, 1 numeric, and 1 special character.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Generate On Prem Connector Configuration
	* `update` - (Defaults to 20 minutes), when updating the Generate On Prem Connector Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Generate On Prem Connector Configuration


## Import

GenerateOnPremConnectorConfiguration can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_generate_on_prem_connector_configuration.test_generate_on_prem_connector_configuration "id"
```

