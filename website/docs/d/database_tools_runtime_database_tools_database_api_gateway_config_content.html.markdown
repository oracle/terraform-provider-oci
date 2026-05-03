---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_content"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_content"
description: |-
  Provides details about a specific Database Tools Database Api Gateway Config Content in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_content
This data source provides details about a specific Database Tools Database Api Gateway Config Content resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get the content of a Database Tools database API gateway config

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_content" "test_database_tools_database_api_gateway_config_content" {
	#Required
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.


## Attributes Reference

The following attributes are exported:


