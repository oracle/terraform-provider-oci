// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseExadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"base64_encode_content":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DatabaseExadataInfrastructureDownloadConfigFileResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseAutonomousExadataInfrastructureRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureDownloadConfigFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureDownloadConfigFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_exadata_infrastructure_download_config_file.test_exadata_infrastructure_download_config_file"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure_download_config_file", "test_exadata_infrastructure_download_config_file", acctest.Required, acctest.Create, DatabaseDatabaseExadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadataInfrastructureDownloadConfigFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure_download_config_file", "test_exadata_infrastructure_download_config_file", acctest.Optional, acctest.Create, DatabaseDatabaseExadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadataInfrastructureDownloadConfigFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
