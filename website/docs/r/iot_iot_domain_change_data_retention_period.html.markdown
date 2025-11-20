---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_iot_domain_change_data_retention_period"
sidebar_current: "docs-oci-resource-iot-iot_domain_change_data_retention_period"
description: |-
  Provides the Iot Domain Change Data Retention Period resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_iot_domain_change_data_retention_period
This resource provides the Iot Domain Change Data Retention Period resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/IotDomain/ChangeDataRetentionPeriod

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Updates Data Retention Period of the IoT Domain.


## Example Usage

```hcl
resource "oci_iot_iot_domain_change_data_retention_period" "test_iot_domain_change_data_retention_period" {
	#Required
	data_retention_period_in_days = var.iot_domain_change_data_retention_period_data_retention_period_in_days
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id
	type = var.iot_domain_change_data_retention_period_type
}
```

## Argument Reference

The following arguments are supported:

* `data_retention_period_in_days` - (Required) The duration (in days) for which data will be retained in the IoT domain.
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `type` - (Required) The type of data retention period to apply. Allowed values are RAW_DATA, REJECTED_DATA, HISTORIZED_DATA, and RAW_COMMAND_DATA. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Iot Domain Change Data Retention Period
	* `update` - (Defaults to 20 minutes), when updating the Iot Domain Change Data Retention Period
	* `delete` - (Defaults to 20 minutes), when destroying the Iot Domain Change Data Retention Period


## Import

Import is not supported for this resource.

