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
	ServiceMeshServiceMeshProxyDetailSingularDataSourceRepresentation = map[string]interface{}{}

	ServiceMeshProxyDetailResourceConfig = ""
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshProxyDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshProxyDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_service_mesh_proxy_detail.test_proxy_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_proxy_detail", "test_proxy_detail", acctest.Required, acctest.Create, ServiceMeshServiceMeshProxyDetailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceMeshProxyDetailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(singularDatasourceName, "proxy_image"),
			),
		},
	})
}
