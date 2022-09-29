---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_importable_compute_entities"
sidebar_current: "docs-oci-datasource-opsi-importable_compute_entities"
description: |-
  Provides the list of Importable Compute Entities in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_importable_compute_entities
This data source provides the list of Importable Compute Entities in Oracle Cloud Infrastructure Opsi service.

Gets a list of available compute intances running cloud agent to add a new hostInsight.  An Compute entity is "available"
and will be shown if all the following conditions are true:
   1. Compute is running OCA
   2. Oracle Cloud Infrastructure Management Agent is not enabled or If Oracle Cloud Infrastructure Management Agent is enabled
      2.1 The agent OCID is not already being used for an existing hostInsight.
      2.2 The agent availabilityStatus = 'ACTIVE'
      2.3 The agent lifecycleState = 'ACTIVE'


## Example Usage

```hcl
data "oci_opsi_importable_compute_entities" "test_importable_compute_entities" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `importable_compute_entity_summary_collection` - The list of importable_compute_entity_summary_collection.

### ImportableComputeEntity Reference

The following attributes are exported:

* `items` - Array of importable compute entity objects.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `compute_display_name` - The [Display Name](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Compute Instance
	* `compute_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
	* `entity_source` - Source of the importable agent entity.
	* `host_name` - The host name. The host name is unique amongst the hosts managed by the same management agent.
	* `platform_type` - Platform type. Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS]. Supported platformType(s) for MACS-managed cloud host insight: [LINUX]. Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS]. 

