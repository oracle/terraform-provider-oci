// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_option_id": acctest.Representation{RepType: acctest.Required, Create: `all`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	ContainerengineClusterOptionResourceConfig = ""
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_option.test_cluster_option"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterOptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_option_id"),

				resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
		// verify singular datasource with compartment_id
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Optional, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterOptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_option_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_pod_network_options.#", "2"),
				resource.TestMatchResourceAttr(singularDatasourceName, "kubernetes_versions.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
