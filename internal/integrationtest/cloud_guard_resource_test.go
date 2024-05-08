package integrationtest

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
/*
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

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
	CloudGuardResourceSingularDataSourceRepresentation = map[string]interface{}{
		"resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_resource.test_resource.id}`},
	}

	CloudGuardResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"cve_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_cve.test_cve.id}`},
		"cvss_score":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"cvss_score_greater_than":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"cvss_score_less_than":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"detector_rule_id_list":     acctest.Representation{RepType: acctest.Optional, Create: []string{`detectorRuleIdList`}},
		"detector_type":             acctest.Representation{RepType: acctest.Optional, Create: `IAAS_ACTIVITY_DETECTOR`},
		"region":                    acctest.Representation{RepType: acctest.Optional, Create: `region`},
		"risk_level":                acctest.Representation{RepType: acctest.Optional, Create: `riskLevel`},
		"risk_level_greater_than":   acctest.Representation{RepType: acctest.Optional, Create: `riskLevelGreaterThan`},
		"risk_level_less_than":      acctest.Representation{RepType: acctest.Optional, Create: `riskLevelLessThan`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	CloudGuardResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_detector_recipe", "test_detector_recipe", acctest.Required, acctest.Create, CloudGuardDetectorRecipeRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources_view", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_responder_recipe", "test_responder_recipe", acctest.Required, acctest.Create, CloudGuardResponderRecipeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_target", "test_target", acctest.Required, acctest.Create, CloudGuardTargetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Required, acctest.Create, IdentityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_guard_resources.test_resources_view"
	singularDatasourceName := "data.oci_cloud_guard_resource.test_resource"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources_view", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "cve_id"),
				resource.TestCheckResourceAttr(datasourceName, "cvss_score", "10"),
				resource.TestCheckResourceAttr(datasourceName, "cvss_score_greater_than", "10"),
				resource.TestCheckResourceAttr(datasourceName, "cvss_score_less_than", "10"),
				resource.TestCheckResourceAttr(datasourceName, "detector_rule_id_list.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "detector_type", "IAAS_ACTIVITY_DETECTOR"),
				resource.TestCheckResourceAttr(datasourceName, "region", "region"),
				resource.TestCheckResourceAttr(datasourceName, "risk_level", "riskLevel"),
				resource.TestCheckResourceAttr(datasourceName, "risk_level_greater_than", "riskLevelGreaterThan"),
				resource.TestCheckResourceAttr(datasourceName, "risk_level_less_than", "riskLevelLessThan"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resource", "test_resource", acctest.Required, acctest.Create, CloudGuardResourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "additional_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "open_ports_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "problem_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "risk_level"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_first_monitored"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_monitored"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vulnerability_count"),
			),
		},
	})
}*/
