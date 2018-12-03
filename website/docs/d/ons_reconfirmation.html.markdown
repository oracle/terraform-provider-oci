---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_reconfirmation"
sidebar_current: "docs-oci-datasource-ons-reconfirmation"
description: |-
  Provides details about a specific Reconfirmation in Oracle Cloud Infrastructure Ons service
---

# Data Source: oci_ons_reconfirmation
This data source provides reconfirmation details about a specific Subscription resource in Oracle Cloud Infrastructure Ons service.

Gets the reconfirmation details for the specified subscription.


## Example Usage

```hcl
data "oci_ons_reconfirmation" "test_reconfirmation" {
	#Required
	id = "${var.reconfirmation_id}"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The [OCID](/iaas/Content/General/Concepts/identifiers.htm) of the subscription to reconfirm.

## Attributes Reference

The following attributes are exported:

* `url` - (Computed) The reconfirmation url protocol.