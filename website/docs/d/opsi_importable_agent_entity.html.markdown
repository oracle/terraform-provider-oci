---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_importable_agent_entity"
sidebar_current: "docs-oci-datasource-opsi-importable_agent_entity"
description: |-
  Provides details about a specific Importable Agent Entity in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_importable_agent_entity
This data source provides details about a specific Importable Agent Entity resource in Oracle Cloud Infrastructure Opsi service.

Gets a list of agent entities available to add a new hostInsight.  An agent entity is "available"
and will be shown if all the following conditions are true:
   1.  The agent OCID is not already being used for an existing hostInsight.
   2.  The agent availabilityStatus = 'ACTIVE'
   3.  The agent lifecycleState = 'ACTIVE'


## Example Usage

```hcl
data "oci_opsi_importable_agent_entity" "test_importable_agent_entity" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `items` - Array of importable agent entity objects.
	* `entity_source` - Source of the importable agent entity.
	* `host_name` - The host name. The host name is unique amongst the hosts managed by the same management agent.
	* `management_agent_display_name` - The [Display Name](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Management Agent
	* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	* `platform_type` - Platform type. Supported platformType(s) for MACS-managed external host insight: [LINUX]. Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX]. 

