// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafediscoveryAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: `targetId`},
		"is_common":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"sensitive_data_model_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"sensitive_type_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeDiscoveryAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_analytics", "test_discovery_analytics", acctest.Required, acctest.Create, DataSafediscoveryAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeDiscoveryAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDiscoveryAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_discovery_analytics.test_discovery_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_analytics_collection.#"),
			),
		},
	})
}
