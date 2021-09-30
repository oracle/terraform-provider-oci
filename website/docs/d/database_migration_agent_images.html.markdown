---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_agent_images"
sidebar_current: "docs-oci-datasource-database_migration-agent_images"
description: |-
  Provides the list of Agent Images in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_agent_images
This data source provides the list of Agent Images in Oracle Cloud Infrastructure Database Migration service.

Get details of the ODMS Agent Images available to install on-premises.

## Example Usage

```hcl
data "oci_database_migration_agent_images" "test_agent_images" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `agent_image_collection` - The list of agent_image_collection.

### AgentImage Reference

The following attributes are exported:

* `items` - Items in collection. 
	* `download_url` - URL to download Agent Image of the ODMS Agent. 
	* `version` - ODMS Agent Image version. 

