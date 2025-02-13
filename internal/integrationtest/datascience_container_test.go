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
	DatascienceContainerDataSourceRepresentation = map[string]interface{}{
		"container_name":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_container.test_container.name}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"is_latest":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"tag_query_param":   acctest.Representation{RepType: acctest.Optional, Create: `tagQueryParam`},
		"target_workload":   acctest.Representation{RepType: acctest.Optional, Create: `MODEL_DEPLOYMENT`},
		"usage_query_param": acctest.Representation{RepType: acctest.Optional, Create: `INFERENCE`},
	}

	DatascienceContainerResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_containers", "test_containers", acctest.Required, acctest.Create, DatascienceContainerDataSourceRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceContainerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceContainerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_containers.test_containers"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + DatascienceContainerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "containers.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.container_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.family_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.is_latest"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "containers.0.tag"),
				resource.TestCheckResourceAttr(datasourceName, "containers.0.tag_configuration_list.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "containers.0.target_workloads.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "containers.0.usages.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "containers.0.workload_configuration_details_list.#", "1"),
			),
		},
	})
}
