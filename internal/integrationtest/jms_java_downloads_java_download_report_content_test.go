// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsJavaDownloadsJavaDownloadReportContentSingularDataSourceRepresentation = map[string]interface{}{
		"java_download_report_id": acctest.Representation{
			RepType: acctest.Required,
			Create:  `${oci_jms_java_downloads_java_download_report.test_java_download_report_content_data.id}`},
	}

	JmsJavaDownloadsJavaDownloadReportContentResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_jms_java_downloads_java_download_report",
		"test_java_download_report_content_data",
		acctest.Optional,
		acctest.Create,
		JmsJavaDownloadsJavaDownloadReportRepresentation,
	)
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaDownloadReportContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaDownloadReportContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	singularDatasourceName := "data.oci_jms_java_downloads_java_download_report_content.test_java_download_report_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_report_content",
					"test_java_download_report_content",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadReportContentSingularDataSourceRepresentation,
				) +
				JmsJavaDownloadsJavaDownloadReportContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "java_download_report_id"),
			),
		},
	})
}
