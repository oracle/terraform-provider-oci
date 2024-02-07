// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_management_agent_management_agent", ManagementAgentManagementAgentResource())
	tfresource.RegisterResource("oci_management_agent_management_agent_data_source", ManagementAgentManagementAgentDataSourceResource())
	tfresource.RegisterResource("oci_management_agent_management_agent_install_key", ManagementAgentManagementAgentInstallKeyResource())
}
