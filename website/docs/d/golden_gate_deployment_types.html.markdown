---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_types"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_types"
description: |-
  Provides the list of Deployment Types in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_types
This data source provides the list of Deployment Types in Oracle Cloud Infrastructure Golden Gate service.

Returns an array of DeploymentTypeDescriptor


## Example Usage

```hcl
data "oci_golden_gate_deployment_types" "test_deployment_types" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_type = var.deployment_type_deployment_type
	display_name = var.deployment_type_display_name
	ogg_version = var.deployment_type_ogg_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `deployment_type` - (Optional) The type of deployment, the value determines the exact 'type' of the service executed in the deployment. Default value is DATABASE_ORACLE. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `ogg_version` - (Optional) Allows to query by a specific GoldenGate version. 


## Attributes Reference

The following attributes are exported:

* `deployment_type_collection` - The list of deployment_type_collection.

### DeploymentType Reference

The following attributes are exported:

* `items` - Array of DeploymentTypeSummary 
	* `category` - The deployment category defines the broad separation of the deployment type into categories.  Currently the separation is 'DATA_REPLICATION' and 'STREAM_ANALYTICS'. 
	* `connection_types` - An array of connectionTypes. 
	* `deployment_type` - The type of deployment, the value determines the exact 'type' of service executed in the Deployment. NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of the equivalent 'DATABASE_ORACLE' value. 
	* `display_name` - An object's Display Name. 
	* `ogg_version` - Version of OGG 
	* `source_technologies` - List of the supported technologies generally.  The value is a freeform text string generally consisting of a description of the technology and optionally the speific version(s) support.  For example, [ "Oracle Database 19c", "Oracle Exadata", "OCI Streaming" ] 
	* `target_technologies` - List of the supported technologies generally.  The value is a freeform text string generally consisting of a description of the technology and optionally the speific version(s) support.  For example, [ "Oracle Database 19c", "Oracle Exadata", "OCI Streaming" ] 

