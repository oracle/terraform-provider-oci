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
	CloudGuardCloudGuardProblemEntitySingularDataSourceRepresentation = map[string]interface{}{
		"problem_id": acctest.Representation{RepType: acctest.Required, Create: `${var.problem_id}`},
	}

	CloudGuardCloudGuardProblemEntityDataSourceRepresentation = map[string]interface{}{
		"problem_id": acctest.Representation{RepType: acctest.Required, Create: `${var.problem_id}`},
	}

	CloudGuardProblemEntityResourceConfig = ""
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardProblemEntityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardProblemEntityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	problemId := utils.GetEnvSettingWithBlankDefault("problem_ocid")
	problemIdVariableStr := fmt.Sprintf("variable \"problem_id\" { default = \"%s\" }\n", problemId)

	ProblemEntityName := "data.oci_cloud_guard_problem_entities.test_problem_entities"
	singularProblemEntityName := "data.oci_cloud_guard_problem_entity.test_problem_entity"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify problem entities
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_problem_entities", "test_problem_entities", acctest.Required, acctest.Create, CloudGuardCloudGuardProblemEntityDataSourceRepresentation) +
				compartmentIdVariableStr + problemIdVariableStr + CloudGuardProblemEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(ProblemEntityName, "problem_id"),
				resource.TestCheckResourceAttrSet(ProblemEntityName, "problem_entity_collection.#"),
				resource.TestCheckResourceAttr(ProblemEntityName, "problem_entity_collection.0.items.#", "2"),
				resource.TestCheckResourceAttr(ProblemEntityName, "problem_entity_collection.0.items.0.problem_id", problemId),
			),
		},
		// verify singular problem entity
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_problem_entity", "test_problem_entity", acctest.Required, acctest.Create, CloudGuardCloudGuardProblemEntitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + problemIdVariableStr + CloudGuardProblemEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularProblemEntityName, "problem_id"),
				resource.TestCheckResourceAttr(singularProblemEntityName, "items.#", "2"),
				resource.TestCheckResourceAttr(singularProblemEntityName, "items.0.problem_id", problemId),
			),
		},
	})
}
