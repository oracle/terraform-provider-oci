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
	MulticloudMulticloudsubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.root_compartment_id}`},
	}

	MulticloudMulticloudsubscriptionResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudMulticloudsubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudMulticloudsubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_root_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"root_compartment_id\" {}\n")

	datasourceName := "data.oci_multicloud_multicloudsubscriptions.test_multicloudsubscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_multicloudsubscriptions", "test_multicloudsubscriptions", acctest.Required, acctest.Create, MulticloudMulticloudsubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + MulticloudMulticloudsubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_subscription_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_subscription_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_subscription_collection.0.items.0.classic_subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_subscription_collection.0.items.0.subscription_id"),
			),
		},
	})
}
