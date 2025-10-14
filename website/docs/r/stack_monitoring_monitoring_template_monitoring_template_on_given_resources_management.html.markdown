---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management"
sidebar_current: "docs-oci-resource-stack_monitoring-monitoring_template_monitoring_template_on_given_resources_management"
description: |-
  Provides the Monitoring Template Monitoring Template On Given Resources Management resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management
This resource provides the Monitoring Template Monitoring Template On Given Resources Management resource in Oracle Cloud Infrastructure Stack Monitoring service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/stack-monitoring/latest/MonitoringTemplateMonitoringTemplateOnGivenResourcesManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/stack_monitoring
Apply the Monitoring Template identified by the id

## Example Usage

```hcl
resource "oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management" "test_monitoring_template_monitoring_template_on_given_resources_management" {
	#Required
	monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template.id
	enable_monitoring_template_on_given_resources = var.enable_monitoring_template_on_given_resources
}
```

## Argument Reference

The following arguments are supported:

* `monitoring_template_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
* `enable_monitoring_template_on_given_resources` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitoring Template Monitoring Template On Given Resources Management
	* `update` - (Defaults to 20 minutes), when updating the Monitoring Template Monitoring Template On Given Resources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Monitoring Template Monitoring Template On Given Resources Management
