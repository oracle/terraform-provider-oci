---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_audit_configuration"
sidebar_current: "docs-oci-datasource-audit-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Audit service
---

# Data Source: oci_audit_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Audit service.

Get the configuration

## Example Usage 

```hcl
data "oci_audit_configuration" "test_configuration" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) ID of the root compartment (tenancy)


## Attributes Reference

The following attributes are exported:

* `retention_period_days` - The retention period days

