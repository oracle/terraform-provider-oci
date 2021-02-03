---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_configuration"
sidebar_current: "docs-oci-datasource-artifacts-container_configuration"
description: |-
  Provides details about a specific Container Configuration in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_configuration
This data source provides details about a specific Container Configuration resource in Oracle Cloud Infrastructure Artifacts service.

Get container configuration.

## Example Usage

```hcl
data "oci_artifacts_container_configuration" "test_container_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `is_repository_created_on_first_push` - Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment. 
* `namespace` - The tenancy namespace used in the container repository path.

