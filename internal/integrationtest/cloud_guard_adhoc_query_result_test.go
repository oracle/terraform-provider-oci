// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
/*
import (
	"fmt"
	//"fmt"
	"github.com/oracle/terraform-provider-oci/internal/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	CloudGuardAdhocQueryResultDataSourceRepresentation = map[string]interface{}{
		"adhoc_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_adhoc_query.test_adhoc_query.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	CloudGuardAdhocQueryResultResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Create, CloudGuardAdhocQueryRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardAdhocQueryResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardAdhocQueryResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_guard_adhoc_query_results.test_adhoc_query_results"

	//acctest.SaveConfigContent("", "", "", t)

	/*acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardAdhocQueryResourceDependencies+
	acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Create, CloudGuardAdhocQueryRepresentation), "cloudguard", "adhocQuery", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_adhoc_query_results", "test_adhoc_query_results", acctest.Required, acctest.Create, CloudGuardAdhocQueryResultDataSourceRepresentation) + CloudGuardAdhocQueryResultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "adhoc_query_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "adhoc_query_result_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "adhoc_query_result_collection.0.items.#", "0"),
			),
		},
	})
}*/
