---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_configuration"
sidebar_current: "docs-oci-datasource-license_manager-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure License Manager service.

Retrieves configuration for a compartment.

## Example Usage

```hcl
data "oci_license_manager_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the configuration is specified. 
* `email_ids` - The list of associated configuration email IDs.
* `time_created` - The time the configuration was created. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `time_updated` - The time the configuration was updated. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.

