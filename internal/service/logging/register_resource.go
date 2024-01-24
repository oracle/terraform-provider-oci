// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_logging_log", LoggingLogResource())
	tfresource.RegisterResource("oci_logging_log_group", LoggingLogGroupResource())
	tfresource.RegisterResource("oci_logging_log_saved_search", LoggingLogSavedSearchResource())
	tfresource.RegisterResource("oci_logging_unified_agent_configuration", LoggingUnifiedAgentConfigurationResource())
}
