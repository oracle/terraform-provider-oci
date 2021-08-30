---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_console_connections"
sidebar_current: "docs-oci-datasource-core-instance_console_connections"
description: |-
  Provides the list of Instance Console Connections in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_console_connections
This data source provides the list of Instance Console Connections in Oracle Cloud Infrastructure Core service.

Lists the console connections for the specified compartment or instance.

For more information about instance console connections, see [Troubleshooting Instances Using Instance Console Connections](https://docs.cloud.oracle.com/iaas/Content/Compute/References/serialconsole.htm).


## Example Usage

```hcl
data "oci_core_instance_console_connections" "test_instance_console_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	instance_id = oci_core_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_id` - (Optional) The OCID of the instance.


## Attributes Reference

The following attributes are exported:

* `instance_console_connections` - The list of instance_console_connections.

### InstanceConsoleConnection Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment to contain the console connection.
* `connection_string` - The SSH connection string for the console connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `fingerprint` - The SSH public key's fingerprint for client authentication to the console connection.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the console connection.
* `instance_id` - The OCID of the instance the console connection connects to.
* `service_host_key_fingerprint` - The SSH public key's fingerprint for the console connection service host.
* `state` - The current state of the console connection.
* `vnc_connection_string` - The SSH connection string for the SSH tunnel used to connect to the console connection over VNC. 

