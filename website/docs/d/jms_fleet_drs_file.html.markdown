---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_drs_file"
sidebar_current: "docs-oci-datasource-jms-fleet_drs_file"
description: |-
  Provides details about a specific Fleet Drs File in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_drs_file
This data source provides details about a specific Fleet Drs File resource in Oracle Cloud Infrastructure Jms service.

Get the detail about the created DRS file in the Fleet.

## Example Usage

```hcl
data "oci_jms_fleet_drs_file" "test_fleet_drs_file" {
	#Required
	drs_file_key = var.fleet_drs_file_drs_file_key
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `drs_file_key` - (Required) The unique identifier of the DRS File in Object Storage.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `bucket` - The Object Storage bucket name where the DRS file is located.
* `checksum_type` - The checksum type for the DRS file in Object Storage.
* `checksum_value` - The checksum value for the DRS file in Object Storage.
* `drs_file_key` - The unique identifier of the DRS file in Object Storage.
* `drs_file_name` - The name of the DRS file in Object Store.
* `is_default` - To check if the DRS file is the detfault ones.
* `namespace` - The namespace for Object Storage.

