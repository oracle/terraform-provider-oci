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
	MulticloudMulticloudpolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.policies_tenancy}`},
	}

	MulticloudMulticloudpolicyResourceConfig = ""
)

// issue-routing-tag: multicloud/default
func TestMulticloudMulticloudpolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMulticloudMulticloudpolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_policies_tenancy")
	compartmentIdVariableStr := fmt.Sprintf("variable \"policies_tenancy\" {}\n")

	datasourceName := "data.oci_multicloud_multicloudpolicies.test_multicloudpolicies"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_multicloud_multicloudpolicies", "test_multicloudpolicies", acctest.Required, acctest.Create, MulticloudMulticloudpolicyDataSourceRepresentation) +
				compartmentIdVariableStr + MulticloudMulticloudpolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.#"),
				// resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.groups.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.subscription_type"),

				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.0.compartment_id"),
				// resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.0.compartment_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "multicloud_policy_collection.0.items.0.policies.0.statements.#"),
			),
		},
	})
}
