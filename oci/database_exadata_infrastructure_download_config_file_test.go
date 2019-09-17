// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	exadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"base64_encode_content":     Representation{repType: Optional, create: `true`},
	}

	ExadataInfrastructureDownloadConfigFileResourceConfig = ExadataInfrastructureRequiredOnlyResource
)

func TestDatabaseExadataInfrastructureDownloadConfigFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureDownloadConfigFileResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_exadata_infrastructure_download_config_file.test_exadata_infrastructure_download_config_file"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure_download_config_file", "test_exadata_infrastructure_download_config_file", Required, Create, exadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExadataInfrastructureDownloadConfigFileResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
			// verify singular datasource with base64 encoded content
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_exadata_infrastructure_download_config_file", "test_exadata_infrastructure_download_config_file", Optional, Create, exadataInfrastructureDownloadConfigFileSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExadataInfrastructureDownloadConfigFileResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
