// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vn_monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_vn_monitoring_path_analyzer_test", VnMonitoringPathAnalyzerTestDataSource())
	tfresource.RegisterDatasource("oci_vn_monitoring_path_analyzer_tests", VnMonitoringPathAnalyzerTestsDataSource())
}
