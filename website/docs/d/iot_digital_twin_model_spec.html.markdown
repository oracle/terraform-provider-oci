---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_model_spec"
sidebar_current: "docs-oci-datasource-iot-digital_twin_model_spec"
description: |-
  Provides details about a specific Digital Twin Model Spec in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_model_spec
This data source provides details about a specific Digital Twin Model Spec resource in Oracle Cloud Infrastructure Iot service.

Retrieves the spec of digital twin model identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_digital_twin_model_spec" "test_digital_twin_model_spec" {
	#Required
	digital_twin_model_id = oci_iot_digital_twin_model.test_digital_twin_model.id
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin model.


## Attributes Reference

The following attributes are exported:


