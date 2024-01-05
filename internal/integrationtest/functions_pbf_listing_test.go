// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FunctionsFunctionsPbfListingSingularDataSourceRepresentation = map[string]interface{}{
		"pbf_listing_id": acctest.Representation{RepType: acctest.Required, Create: `${var.pbf_listing_id}`},
	}

	FunctionsFunctionsPbfListingDataSourceRepresentation = map[string]interface{}{
		"name":             acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_name}`},
		"name_contains":    acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_name_fragment}`},
		"name_starts_with": acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_name_prefix}`},
		"pbf_listing_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.pbf_listing_id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"trigger":          acctest.Representation{RepType: acctest.Optional, Create: []string{`HTTP`}},
	}
)

// issue-routing-tag: functions/default
func TestFunctionsPbfListingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsPbfListingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	pbfListingName := utils.GetEnvSettingWithBlankDefault("pbf_listing_name")
	pbfListingNameVariableStr := fmt.Sprintf("variable \"pbf_listing_name\" { default = \"%s\" }\n", pbfListingName)
	pbfListingNameFragment := utils.GetEnvSettingWithBlankDefault("pbf_listing_name_fragment")
	pbfListingNameFragmentVariableStr := fmt.Sprintf("variable \"pbf_listing_name_fragment\" { default = \"%s\" }\n", pbfListingNameFragment)
	pbfListingNamePrefix := utils.GetEnvSettingWithBlankDefault("pbf_listing_name_prefix")
	pbfListingNamePrefixVariableStr := fmt.Sprintf("variable \"pbf_listing_name_prefix\" { default = \"%s\" }\n", pbfListingNamePrefix)
	pbfListingId := utils.GetEnvSettingWithBlankDefault("pbf_listing_id")
	pbfListingIdVariableStr := fmt.Sprintf("variable \"pbf_listing_id\" { default = \"%s\" }\n", pbfListingId)

	datasourceName := "data.oci_functions_pbf_listings.test_pbf_listings"
	singularDatasourceName := "data.oci_functions_pbf_listing.test_pbf_listing"

	acctest.SaveConfigContent("", "", "", t)

	pbfListingDataSourceRepresentationWithNameOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name_contains", "name_starts_with", "pbf_listing_id", "state", "trigger"}, FunctionsFunctionsPbfListingDataSourceRepresentation)
	pbfListingDataSourceRepresentationWithNameContainsOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "name_starts_with", "pbf_listing_id", "state", "trigger"}, FunctionsFunctionsPbfListingDataSourceRepresentation)
	pbfListingDataSourceRepresentationWithNameStartsWithOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "name_contains", "pbf_listing_id", "state", "trigger"}, FunctionsFunctionsPbfListingDataSourceRepresentation)
	pbfListingDataSourceRepresentationWithIdOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "name_contains", "name_starts_with", "state", "trigger"}, FunctionsFunctionsPbfListingDataSourceRepresentation)
	pbfListingDataSourceRepresentationWithStateOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "name_contains", "name_starts_with", "pbf_listing_id", "trigger"}, FunctionsFunctionsPbfListingDataSourceRepresentation)
	pbfListingDataSourceRepresentationWithTriggerOnly := acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"name", "name_contains", "name_starts_with", "pbf_listing_id", "state"}, FunctionsFunctionsPbfListingDataSourceRepresentation)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pbf_listings_collection.0.items.#"),
			),
		},
		{
			Config: config + pbfListingNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithNameOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name", pbfListingName),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.0.name", pbfListingName),
			),
		},
		{
			Config: config + pbfListingNameFragmentVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithNameContainsOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", pbfListingNameFragment),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "pbf_listings_collection.0.items.0.name", regexp.MustCompile(fmt.Sprintf("^.*%s.*$", pbfListingNameFragment))),
			),
		},
		{
			Config: config + pbfListingNamePrefixVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithNameStartsWithOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckResourceAttr(datasourceName, "name_starts_with", pbfListingNamePrefix),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "pbf_listings_collection.0.items.0.name", regexp.MustCompile(fmt.Sprintf("^%s.*$", pbfListingNamePrefix))),
			),
		},
		{
			Config: config + pbfListingIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithIdOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listing_id", pbfListingId),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.0.items.0.id", pbfListingId),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithStateOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckNoResourceAttr(datasourceName, "trigger.#"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pbf_listings_collection.0.items.#"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Optional, acctest.Create, pbfListingDataSourceRepresentationWithTriggerOnly),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_contains"),
				resource.TestCheckNoResourceAttr(datasourceName, "name_starts_with"),
				resource.TestCheckNoResourceAttr(datasourceName, "pbf_listing_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "state"),
				resource.TestCheckResourceAttr(datasourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "trigger.0", "HTTP"),
				resource.TestCheckResourceAttr(datasourceName, "pbf_listings_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pbf_listings_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + pbfListingIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing", "test_pbf_listing", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "pbf_listing_id", pbfListingId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", pbfListingId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "publisher_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "triggers.#", "1"),
			),
		},
	})
}
