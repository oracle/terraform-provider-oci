// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	// representation to generate addon_option dep data source
	ContainerengineAddonOptionDepDataSourceRepresentation = map[string]interface{}{
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
	}

	// representation for the real testing
	ContainerengineAddonOptionDataSourceRepresentation = map[string]interface{}{
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
		"addon_name":         acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_containerengine_addon_options.test_adddon_options_dep.addon_options[0].name}`},
	}

	// all dependencies for the data source test with required input
	ContainerengineAddonOptionDataSourceDependenciesForRequired = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation)

	// all dependencies for the data source test with optional input
	ContainerengineAddonOptionDataSourceDependenciesForOptional = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addon_options", "test_adddon_options_dep", acctest.Required, acctest.Create, ContainerengineAddonOptionDepDataSourceRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineAddonOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineAddonOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceNameForRequired := "data.oci_containerengine_addon_options.test_addon_options"
	datasourceNameForOptional := "data.oci_containerengine_addon_options.test_addon_options_with_name"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required
		{
			Config: config + ContainerengineAddonOptionDataSourceDependenciesForRequired +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addon_options", "test_addon_options", acctest.Required, acctest.Create, ContainerengineAddonOptionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "kubernetes_version"),

				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.#"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.addon_group"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.addon_schema_version"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.description"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.is_essential"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.name"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.state"),
				resource.TestCheckResourceAttrSet(datasourceNameForRequired, "addon_options.0.time_created"),
				resource.TestCheckResourceAttr(datasourceNameForRequired, "addon_options.0.versions.#", "1"),
			),
		},
		// verify datasource with optional
		{
			Config: config + ContainerengineAddonOptionDataSourceDependenciesForRequired + ContainerengineAddonOptionDataSourceDependenciesForOptional +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addon_options", "test_addon_options_with_name", acctest.Optional, acctest.Create, ContainerengineAddonOptionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_name"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "kubernetes_version"),

				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.#"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.addon_group"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.addon_schema_version"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.description"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.is_essential"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.name"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.state"),
				resource.TestCheckResourceAttrSet(datasourceNameForOptional, "addon_options.0.time_created"),
				resource.TestCheckResourceAttr(datasourceNameForOptional, "addon_options.0.versions.#", "1"),
			),
		},
	})
}
