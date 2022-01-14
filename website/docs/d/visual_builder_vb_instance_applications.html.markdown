---
subcategory: "Visual Builder"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_visual_builder_vb_instance_applications"
sidebar_current: "docs-oci-datasource-visual_builder-vb_instance_applications"
description: |-
  Provides the list of published and staged applications of a Visual Builder Instance in Oracle Cloud Infrastructure Visual Builder service
---

# Data Source: oci_visual_builder_vb_instance_applications
This data source provides the list of published and staged applications of a Visual Builder Instance in Oracle Cloud Infrastructure Visual Builder service.

Returns a list of published and staged applications of a Visual Builder instance.


## Example Usage

```hcl
data "oci_visual_builder_vb_instance_applications" "test_vb_instance_applications" {
	#Required
	vb_instance_id = oci_visual_builder_vb_instance.test_vb_instance.id
	idcs_open_id = "idcs_open_id_value"
}
```

## Argument Reference

The following arguments are supported:

* `vb_instance_id` - (Required) Unique Vb Instance identifier.
* `idcs_open_id` - (Required) Encrypted IDCS Open ID token which allows access to Visual Builder REST apis


## Attributes Reference

The following attributes are exported:

* `application_summary_collection` - The list of application_summary_collection.

### VbInstance Reference

The following attributes are exported:

* `id` - The Visual Builder application identifier.
* `project_id` - The Visual Builder application project identifier.
* `state` - The state of visual builder application. Either LIVE or STAGED
* `version` - The Visual Builder application version
