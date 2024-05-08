// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
/*
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

//to download the file.
//should this be allowed and exposed through TF?

// /adhocQuerie/{adhocQueryID}/result/content

var (
	CloudGuardAdhocQueryResultContentSingularDataSourceRepresentation = map[string]interface{}{
		"adhoc_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_adhoc_query.test_adhoc_query.id}`},
	}

	CloudGuardAdhocQueryResultContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Create, CloudGuardAdhocQueryRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardAdhocQueryResultContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardAdhocQueryResultContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_cloud_guard_adhoc_query_result_content.test_adhoc_query_result_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		//to download the file.
		//should this be allowed and exposed through TF?
		{
			Config: config + compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_adhoc_query_result_content", "test_adhoc_query_result_content", acctest.Required, acctest.Create, CloudGuardAdhocQueryResultContentSingularDataSourceRepresentation) +
				CloudGuardAdhocQueryResultContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "adhoc_query_id"),
			),
		},
	})
}*/
