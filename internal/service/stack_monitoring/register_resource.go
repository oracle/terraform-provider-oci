// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_stack_monitoring_discovery_job", StackMonitoringDiscoveryJobResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resource", StackMonitoringMonitoredResourceResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_associate_monitored_resource", StackMonitoringMonitoredResourcesAssociateMonitoredResourceResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_list_member", StackMonitoringMonitoredResourcesListMemberResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_search", StackMonitoringMonitoredResourcesSearchResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_search_association", StackMonitoringMonitoredResourcesSearchAssociationResource())
}
