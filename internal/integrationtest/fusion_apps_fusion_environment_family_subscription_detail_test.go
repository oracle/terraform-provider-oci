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
	FusionAppsFusionAppsFusionEnvironmentFamilySubscriptionDetailSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
	}

	FusionAppsFusionEnvironmentFamilySubscriptionDetailResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentFamilySubscriptionDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentFamilySubscriptionDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_family_subscription_detail.test_fusion_environment_family_subscription_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family_subscription_detail", "test_fusion_environment_family_subscription_detail", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentFamilySubscriptionDetailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilySubscriptionDetailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_family_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "subscriptions.#", "1"),
			),
		},
	})
}
