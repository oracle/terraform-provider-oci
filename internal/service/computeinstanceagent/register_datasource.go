// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package computeinstanceagent

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_computeinstanceagent_instance_agent_plugin", ComputeinstanceagentInstanceAgentPluginDataSource())
	tfresource.RegisterDatasource("oci_computeinstanceagent_instance_agent_plugins", ComputeinstanceagentInstanceAgentPluginsDataSource())
	tfresource.RegisterDatasource("oci_computeinstanceagent_instance_available_plugins", ComputeinstanceagentInstanceAvailablePluginsDataSource())
}
