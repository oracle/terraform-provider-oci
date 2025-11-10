---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domain_group_configure_data_access"
sidebar_current: "docs-oci-resource-iot-iot_domain_group_configure_data_access"
description: |-
  Provides the Iot Domain Group Configure Data Access resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_domain_group_configure_data_access
This resource provides the Iot Domain Group Configure Data Access resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/IotDomainGroup/ConfigureDataAccess

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Updates an IoT domain Group Data Access.


## Example Usage

```hcl
resource "oci_iot_iot_domain_group_configure_data_access" "test_iot_domain_group_configure_data_access" {
	#Required
	db_allow_listed_vcn_ids = var.iot_domain_group_configure_data_access_db_allow_listed_vcn_ids
	iot_domain_group_id = oci_iot_iot_domain_group.test_iot_domain_group.id
}
```

## Argument Reference

The following arguments are supported:

* `db_allow_listed_vcn_ids` - (Required) This is an array of VCN OCID (virtual cloud network Oracle Cloud ID) that is allowed to connect the data host.
* `iot_domain_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an IoT Domain Group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Domain Group Configure Data Access
	* `update` - (Defaults to 20 minutes), when updating the Iot Domain Group Configure Data Access
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Domain Group Configure Data Access


## Import

Import is not supported for this resource.

