---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_infrastructure_scale_option"
sidebar_current: "docs-oci-datasource-datacc-infrastructure_scale_option"
description: |-
  Provides details about a specific Infrastructure Scale Option in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_infrastructure_scale_option
This data source provides details about a specific Infrastructure Scale Option resource in Oracle Cloud Infrastructure Datacc service.

Get the available scale options for the infrastructure that has the specified
[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_datacc_infrastructure_scale_option" "test_infrastructure_scale_option" {
	#Required
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.


## Attributes Reference

The following attributes are exported:

* `possible_ssd_configurations` - The available scale options for the infrastructure.

