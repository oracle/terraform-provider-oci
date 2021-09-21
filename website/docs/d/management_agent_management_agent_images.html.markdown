---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_images"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_images"
description: |-
  Provides the list of Management Agent Images in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_images
This data source provides the list of Management Agent Images in Oracle Cloud Infrastructure Management Agent service.

Get supported agent image information


## Example Usage

```hcl
data "oci_management_agent_management_agent_images" "test_management_agent_images" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	install_type = var.management_agent_image_install_type
	name = var.management_agent_image_name
	state = var.management_agent_image_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `install_type` - (Optional) A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
* `name` - (Optional) A filter to return only resources that match the entire platform name given.
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.


## Attributes Reference

The following attributes are exported:

* `management_agent_images` - The list of management_agent_images.

### ManagementAgentImage Reference

The following attributes are exported:

* `checksum` - Agent image content SHA256 Hash
* `id` - Agent image resource id
* `object_url` - Object storage URL for download
* `platform_name` - Agent image platform display name
* `platform_type` - Agent image platform type
* `size` - Agent image size in bytes
* `state` - The current state of Management Agent Image
* `version` - Agent image version

