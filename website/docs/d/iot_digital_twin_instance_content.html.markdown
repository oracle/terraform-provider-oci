---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_instance_content"
sidebar_current: "docs-oci-datasource-iot-digital_twin_instance_content"
description: |-
  Provides details about a specific Digital Twin Instance Content in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_instance_content
This data source provides details about a specific Digital Twin Instance Content resource in Oracle Cloud Infrastructure Iot service.

Retrieves the latest snapshot data of digital twin instance identified by the specified OCID.


## Example Usage

```hcl
data "oci_iot_digital_twin_instance_content" "test_digital_twin_instance_content" {
	#Required
	digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id

	#Optional
	should_include_metadata = var.digital_twin_instance_content_should_include_metadata
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin instance. 
* `should_include_metadata` - (Optional) If set to true , digital twin instance metadata is included in the response.


## Attributes Reference

The following attributes are exported:


