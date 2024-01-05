// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	FunctionsFunctionsPbfListingVersionSingularDataSourceRepresentation = map[string]interface{}{
		"pbf_listing_version_id": acctest.Representation{RepType: acctest.Required, Create: `${var.pbf_listing_version_id}`},
	}

	FunctionsFunctionsPbfListingVersionDataSourceRepresentation = map[string]interface{}{
		"pbf_listing_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.pbf_listing_id}`},
		"is_current_version":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"name":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_version_name}`},
		"pbf_listing_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_version_id}`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
)

// issue-routing-tag: functions/default
func TestFunctionsPbfListingVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsPbfListingVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	pbfListingId := utils.GetEnvSettingWithBlankDefault("pbf_listing_id")
	pbfListingIdVariableStr := fmt.Sprintf("variable \"pbf_listing_id\" { default = \"%s\" }\n", pbfListingId)
	pbfListingVersionId := utils.GetEnvSettingWithBlankDefault("pbf_listing_version_id")
	pbfListingVersionIdVariableStr := fmt.Sprintf("variable \"pbf_listing_version_id\" { default = \"%s\" }\n", pbfListingVersionId)
	pbfListingVersionName := utils.GetEnvSettingWithBlankDefault("pbf_listing_version_name")
	pbfListingVersionNameVariableStr := fmt.Sprintf("variable \"pbf_listing_version_name\" { default = \"%s\" }\n", pbfListingVersionName)

	datasourceName := "data.oci_functions_pbf_listing_versions.test_pbf_listing_versions"
	singularDatasourceName := "data.oci_functions_pbf_listing_version.test_pbf_listing_version"

	acctest.SaveConfigContent("", "", "", t)

	pbfListingVersionDataSourceRepresentationWithCurrentVersionOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "pbf_listing_version_id", "state"}, FunctionsFunctionsPbfListingVersionDataSourceRepresentation)
	pbfListingVersionDataSourceRepresentationWithNameOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"is_current_version", "pbf_listing_version_id", "state"}, FunctionsFunctionsPbfListingVersionDataSourceRepresentation)
	pbfListingVersionDataSourceRepresentationWithVersionIdOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"is_current_version", "name", "state"}, FunctionsFunctionsPbfListingVersionDataSourceRepresentation)
	pbfListingVersionDataSourceRepresentationWithStateOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"is_current_version", "name", "pbf_listing_version_id"}, FunctionsFunctionsPbfListingVersionDataSourceRepresentation)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + pbfListingIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_versions", "test_pbf_listing_versions", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingVersionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckNoResourceAttr(datasourceName, "is_current_version"),
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_version_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),

				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pbf_listing_versions_collection.0.items.#"),
			),
		},
		{
			Config: config + pbfListingIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_versions", "test_pbf_listing_versions", acctest.Optional, acctest.Create, pbfListingVersionDataSourceRepresentationWithCurrentVersionOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckResourceAttr(datasourceName, "is_current_version", "true"),
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_version_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),

				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.0.items.#", "1"),
			),
		},
		{
			Config: config + pbfListingIdVariableStr + pbfListingVersionNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_versions", "test_pbf_listing_versions", acctest.Optional, acctest.Create, pbfListingVersionDataSourceRepresentationWithNameOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckNoResourceAttr(datasourceName, "is_current_version"),
				resource.TestCheckResourceAttr(datasourceName, "name", pbfListingVersionName),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_version_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),

				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.0.items.#", "1"),
			),
		},
		{
			Config: config + pbfListingIdVariableStr + pbfListingVersionIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_versions", "test_pbf_listing_versions", acctest.Optional, acctest.Create, pbfListingVersionDataSourceRepresentationWithVersionIdOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckNoResourceAttr(datasourceName, "is_current_version"),
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_version_id", pbfListingVersionId),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),

				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.0.items.#", "1"),
			),
		},
		{
			Config: config + pbfListingIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_versions", "test_pbf_listing_versions", acctest.Optional, acctest.Create, pbfListingVersionDataSourceRepresentationWithStateOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckNoResourceAttr(datasourceName, "is_current_version"),
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_version_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_versions_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pbf_listing_versions_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + pbfListingVersionIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_version", "test_pbf_listing_version", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingVersionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "pbf_listing_version_id", pbfListingVersionId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "change_summary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", pbfListingVersionId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "requirements.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "triggers.#", "1"),
			),
		},
	})
}
