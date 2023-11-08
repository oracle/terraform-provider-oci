---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_drs_files"
sidebar_current: "docs-oci-datasource-jms-fleet_drs_files"
description: |-
  Provides the list of Fleet Drs Files in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_drs_files
This data source provides the list of Fleet Drs Files in Oracle Cloud Infrastructure Jms service.

List the details about the created DRS files in the Fleet.

## Example Usage

```hcl
data "oci_jms_fleet_drs_files" "test_fleet_drs_files" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `drs_file_collection` - The list of drs_file_collection.

### FleetDrsFile Reference

The following attributes are exported:

* `bucket` - The Object Storage bucket name where the DRS file is located.
* `checksum_type` - The checksum type for the DRS file in Object Storage.
* `checksum_value` - The checksum value for the DRS file in Object Storage.
* `drs_file_key` - The unique identifier of the DRS file in Object Storage.
* `drs_file_name` - The name of the DRS file in Object Store.
* `is_default` - To check if the DRS file is the detfault ones.
* `namespace` - The namespace for Object Storage.

