---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_configuration"
sidebar_current: "docs-oci-datasource-service_catalog-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Service Catalog service.

Get the detail of whether the tenancy is in service catalog mode or not.

## Example Usage

```hcl
data "oci_service_catalog_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.


## Attributes Reference

The following attributes are exported:

* `is_service_catalog_mode` - mode of tenancy

