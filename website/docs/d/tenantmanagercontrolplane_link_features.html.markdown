---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_link_features"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-link_features"
description: |-
  Provides the list of Link Features in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_link_features
This data source provides the list of Link Features in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a list of link features.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_link_features" "test_link_features" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `link_features_collection` - The list of link_features_collection.

### LinkFeature Reference

The following attributes are exported:

* `items` - Array containing LinkFeature items.
	* `description` - Description of the feature.
	* `display_name` - Display name of the feature.
	* `feature` - The feature associated with this link. Default value is CORE.
	* `partner_service_console_url` - ConsoleUrl of the feature.
	* `user_guide_url` - UserGuideUrl of the feature.

