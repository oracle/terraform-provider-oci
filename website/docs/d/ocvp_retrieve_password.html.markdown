---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_retrieve_password"
sidebar_current: "docs-oci-datasource-ocvp-retrieve_password"
description: |-
Retrieves the SDDC password in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_retrieve_password
This data source retrieves the SDDC password in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.


## Example Usage

```hcl
data "oci_ocvp_retrieve_password" "test_password" {
	#Required
  sddc_id = var.compartment_id
  type    = var.password_type
}
```

## Argument Reference

The following arguments are supported:

* `sddc_id` - (Required) The [OCID](/Content/General/Concepts/identifiers.htm) of the SDDC.
* `type` - (Required) The SDDC password type.


## Attributes Reference

The following attributes are exported:

* `sddc_password` - SDDC vCenter/NSX/HCX password.
    * `passwordType` - SDDC password type.
    * `value` - SDDC vCenter/NSX/HCX password context.
