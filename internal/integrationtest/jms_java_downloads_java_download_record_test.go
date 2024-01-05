// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	jdRecordTimeStarted = time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339)
	jdRecordTimeEnded   = time.Now().UTC().Format(time.RFC3339)

	JmsJavaDownloadsJavaDownloadRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"time_end":       acctest.Representation{RepType: acctest.Optional, Create: jdRecordTimeEnded},
		"time_start":     acctest.Representation{RepType: acctest.Optional, Create: jdRecordTimeStarted},
	}
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaDownloadRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaDownloadRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_jms_java_downloads_java_download_records.test_java_download_records"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_records",
					"test_java_download_records",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadRecordDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "time_end", jdRecordTimeEnded),
				resource.TestCheckResourceAttr(datasourceName, "time_start", jdRecordTimeStarted),

				resource.TestCheckResourceAttrSet(datasourceName, "java_download_record_collection.#"),
			),
		},
	})
}
