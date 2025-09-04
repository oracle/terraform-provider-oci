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
	MulticloudExternalLocationSummariesMetadataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_service_name}`},
	}

	MulticloudExternalLocationSummariesMetadataResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudExternalLocationSummariesMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudExternalLocationSummariesMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_compartment_id")
	subscriptionServiceName := utils.GetEnvSettingWithBlankDefault("TF_VAR_subscription_service_name")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" {}\n")
	subscriptionServiceNameVariableStr := fmt.Sprintf("variable \"subscription_service_name\" {}\n")

	datasourceName := "data.oci_multicloud_external_location_summaries_metadata.test_external_location_summaries_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + subscriptionServiceNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_external_location_summaries_metadata", "test_external_location_summaries_metadata", acctest.Required, acctest.Create, MulticloudExternalLocationSummariesMetadataDataSourceRepresentation) +
				MulticloudExternalLocationSummariesMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_service_name", subscriptionServiceName),

				resource.TestCheckResourceAttr(datasourceName, "external_location_summaries_metadatum_summary_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_summaries_metadatum_summary_collection.0.items.#"),

				resource.TestCheckResourceAttr(datasourceName, "external_location_summaries_metadatum_summary_collection.0.items.0.external_location.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_summaries_metadatum_summary_collection.0.items.0.external_location.0.csp_region"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_location_summaries_metadatum_summary_collection.0.items.0.external_location.0.csp_region_display_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "external_location_summaries_metadatum_summary_collection.0.items.0.oci_region"),
			),
		},
	})
}
