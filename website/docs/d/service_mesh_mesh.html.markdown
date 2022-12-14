---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_mesh"
sidebar_current: "docs-oci-datasource-service_mesh-mesh"
description: |-
  Provides details about a specific Mesh in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_mesh
This data source provides details about a specific Mesh resource in Oracle Cloud Infrastructure Service Mesh service.

Gets a Mesh by identifier.

## Example Usage

```hcl
data "oci_service_mesh_mesh" "test_mesh" {
	#Required
	mesh_id = oci_service_mesh_mesh.test_mesh.id
}
```

## Argument Reference

The following arguments are supported:

* `mesh_id` - (Required) Unique Mesh identifier.


## Attributes Reference

The following attributes are exported:

* `certificate_authorities` - A list of certificate authority resources to use for creating leaf certificates for mTLS authentication. Currently we only support one certificate authority, but this may expand in future releases. Request with more than one certificate authority will be rejected. 
	* `id` - The OCID of the certificate authority resource.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `display_name` - A user-friendly name. The name does not have to be unique and can be changed after creation. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `mtls` - Sets a minimum level of mTLS authentication for all virtual services within the mesh.
	* `minimum` - DISABLED: No minimum virtual services within this mesh can use any mTLS authentication mode. PERMISSIVE: Virtual services within this mesh can use either PERMISSIVE or STRICT modes. STRICT: All virtual services within this mesh must use STRICT mode. 
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

