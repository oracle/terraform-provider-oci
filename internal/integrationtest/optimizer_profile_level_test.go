// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	profileLevelSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"recommendation_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	profileLevelDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"recommendation_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	ProfileLevelResourceConfig = RecommendationResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, recommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerProfileLevelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerProfileLevelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_profile_levels.test_profile_levels"
	singularDatasourceName := "data.oci_optimizer_profile_level.test_profile_level"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_profile_levels", "test_profile_levels", acctest.Required, acctest.Create, profileLevelDataSourceRepresentation) +
				compartmentIdVariableStr + ProfileLevelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_level_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_level_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_profile_level", "test_profile_level", acctest.Required, acctest.Create, profileLevelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ProfileLevelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
