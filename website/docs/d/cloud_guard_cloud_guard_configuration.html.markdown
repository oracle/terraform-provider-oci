---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_cloud_guard_configuration"
sidebar_current: "docs-oci-datasource-cloud_guard-cloud_guard_configuration"
description: |-
  Provides details about a specific Cloud Guard Configuration in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_cloud_guard_configuration
This data source provides details about a specific Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns the configuration details for a Cloud Guard tenancy,
identified by root compartment OCID.


## Example Usage

```hcl
data "oci_cloud_guard_cloud_guard_configuration" "test_cloud_guard_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `reporting_region` - The reporting region
* `self_manage_resources` - Were Oracle-managed resources created by customer?
* `status` - Status of the Cloud Guard tenant

