// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MulticloudExternalLocationMappingMetadataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name_list}`},
		"subscription_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.subscription_id}`},
	}

	MulticloudExternalLocationMappingMetadataResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudExternalLocationMappingMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudExternalLocationMappingMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_compartment_id")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name_list\" { \n type = list(string)\n }\n")

	datasourceName := "data.oci_multicloud_external_location_mapping_metadata.test_external_location_mapping_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + subscriptionServiceNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_external_location_mapping_metadata", "test_external_location_mapping_metadata", acctest.Required, acctest.Create, MulticloudExternalLocationMappingMetadataDataSourceRepresentation) +
				MulticloudExternalLocationMappingMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_service_name.#"),

				resource.TestCheckResourceAttr(datasourceName, "external_location_mapping_metadatum_summary_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.#"),

				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.oci_logical_ad"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.oci_physical_ad"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.oci_region"),

				resource.TestCheckResourceAttr(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.0.csp_physical_az"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.0.csp_physical_az_display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.0.csp_region"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.0.csp_region_display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_mapping_metadatum_summary_collection.0.items.0.external_location.0.service_name"),
			),
		},
	})
}
