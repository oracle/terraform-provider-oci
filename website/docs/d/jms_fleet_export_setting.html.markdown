---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_export_setting"
sidebar_current: "docs-oci-datasource-jms-fleet_export_setting"
description: |-
  Provides details about a specific Fleet Export Setting in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_export_setting
This data source provides details about a specific Fleet Export Setting resource in Oracle Cloud Infrastructure Jms service.

Returns export setting for the specified fleet.

## Example Usage

```hcl
data "oci_jms_fleet_export_setting" "test_fleet_export_setting" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `export_duration` - The duration of data to be exported for fleets. 
* `export_frequency` - Schedule at which data will be exported. 
* `export_resources` - Resource to export data associated from the fleets. 
* `export_setting_key` - The internal identifier of the export setting. 
* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet. 
* `is_cross_region_acknowledged` - Acknowledgement for cross region target bucket configuration. 
* `is_enabled` - ExportSetting flag to store enabled or disabled status.
* `target_bucket_name` - The name of the bucket where data will be exported. 
* `target_bucket_namespace` - The namespace of the bucket where data will be exported. 
* `target_bucket_region` - The id of the region of the target bucket. 
* `time_created` - The creation date and time of the export setting (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `time_last_modified` - The update date and time of the export setting (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

