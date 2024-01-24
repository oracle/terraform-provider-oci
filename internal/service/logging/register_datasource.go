// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_logging_log", LoggingLogDataSource())
	tfresource.RegisterDatasource("oci_logging_log_group", LoggingLogGroupDataSource())
	tfresource.RegisterDatasource("oci_logging_log_groups", LoggingLogGroupsDataSource())
	tfresource.RegisterDatasource("oci_logging_log_saved_search", LoggingLogSavedSearchDataSource())
	tfresource.RegisterDatasource("oci_logging_log_saved_searches", LoggingLogSavedSearchesDataSource())
	tfresource.RegisterDatasource("oci_logging_logs", LoggingLogsDataSource())
	tfresource.RegisterDatasource("oci_logging_unified_agent_configuration", LoggingUnifiedAgentConfigurationDataSource())
	tfresource.RegisterDatasource("oci_logging_unified_agent_configurations", LoggingUnifiedAgentConfigurationsDataSource())
}
