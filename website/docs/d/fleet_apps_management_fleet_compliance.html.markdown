---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_compliance"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_compliance"
description: |-
  Provides details about a specific Fleet Compliance in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_compliance
This data source provides details about a specific Fleet Compliance resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Retrieve compliance for a fleet.

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_compliance" "test_fleet_compliance" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) Unique Fleet identifier.


## Attributes Reference

The following attributes are exported:

* `compliance_state` - Compliance State.
* `confirmed_target_count` - Confirmed Target Count.
* `non_compliant_target_count` - Non Compliant Target Count.

