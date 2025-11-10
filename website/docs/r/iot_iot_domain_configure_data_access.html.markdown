---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domain_configure_data_access"
sidebar_current: "docs-oci-resource-iot-iot_domain_configure_data_access"
description: |-
  Provides the Iot Domain Configure Data Access resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_domain_configure_data_access
This resource provides the Iot Domain Configure Data Access resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/IotDomain/ConfigureDataAccess

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Updates an IoT domain Data Access.


## Example Usage

```hcl
resource "oci_iot_iot_domain_configure_data_access" "test_iot_domain_configure_data_access" {
	#Required
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id
	type = var.iot_domain_configure_data_access_type

	#Optional
	db_allow_listed_identity_group_names = var.iot_domain_configure_data_access_db_allow_listed_identity_group_names
	db_allowed_identity_domain_host = var.iot_domain_configure_data_access_db_allowed_identity_domain_host
	db_workspace_admin_initial_password = var.iot_domain_configure_data_access_db_workspace_admin_initial_password
}
```

## Argument Reference

The following arguments are supported:

* `db_allow_listed_identity_group_names` - (Required when type=DIRECT) List of IAM groups of form described in [here](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/mnqmn/#GUID-3634D6C9-A7F1-4875-9925-BAEA2D3C5197) that are allowed to directly connect to the data host.
* `db_allowed_identity_domain_host` - (Required when type=ORDS) Host name of identity domain that is used for authenticating connect to data host via ORDS.
* `db_workspace_admin_initial_password` - (Required when type=APEX) Initial admin password for APEX workspace associated with the IoT domain.
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `type` - (Required) configuration type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Domain Configure Data Access
	* `update` - (Defaults to 20 minutes), when updating the Iot Domain Configure Data Access
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Domain Configure Data Access


## Import

Import is not supported for this resource.

