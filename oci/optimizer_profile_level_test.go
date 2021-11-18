// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	profileLevelSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"recommendation_name":       Representation{RepType: Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	profileLevelDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"recommendation_name":       Representation{RepType: Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	ProfileLevelResourceConfig = RecommendationResourceDependencies + GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerProfileLevelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerProfileLevelResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_profile_levels.test_profile_levels"
	singularDatasourceName := "data.oci_optimizer_profile_level.test_profile_level"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_profile_levels", "test_profile_levels", Required, Create, profileLevelDataSourceRepresentation) +
				compartmentIdVariableStr + ProfileLevelResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_optimizer_profile_level", "test_profile_level", Required, Create, profileLevelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ProfileLevelResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
