---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_attribute_auto_activate_status"
sidebar_current: "docs-oci-datasource-apm_traces-attribute_auto_activate_status"
description: |-
  Provides details about a specific Attribute Auto Activate Status in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_attribute_auto_activate_status
This data source provides details about a specific Attribute Auto Activate Status resource in Oracle Cloud Infrastructure Apm Traces service.

Get autoactivation status for a private data key or public data key in the APM Domain.


## Example Usage

```hcl
data "oci_apm_traces_attribute_auto_activate_status" "test_attribute_auto_activate_status" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	data_key_type = var.attribute_auto_activate_status_data_key_type
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `data_key_type` - (Required) Data key type for which auto-activate needs to be turned on or off. 


## Attributes Reference

The following attributes are exported:

* `data_key` - Data key type for which auto-activate needs needs to be turned on or off. 
* `state` - State of autoactivation in this APM Domain.  If "ON" auto-activate is set to true, if "OFF" auto-activate is set to false. 

