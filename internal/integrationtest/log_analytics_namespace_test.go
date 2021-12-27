// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation2 = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	namespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOnBoardRepresentation = map[string]interface{}{
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOffBoardRepresentation = map[string]interface{}{
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	NameSpaceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, namespaceDataSourceRepresentation)

	NameSpaceSingularDataSourceDependencies = NameSpaceResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceResourceOnBoardRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceResource_basic(t *testing.T) {
	t.Skip("skipping test as onboarding tenancy is a one time operation only and cannot be done on a recurring basis")
	httpreplay.SetScenario("TestLogAnalyticsNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace.test_namespace"
	datasourceName := "data.oci_log_analytics_namespaces.test_namespaces"
	singularDatasourceName := "data.oci_log_analytics_namespace.test_namespace"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource also works as a dependency for next step
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, namespaceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttrSet(datasourceName, "namespace_collection.#"),
			),
		},
		// verify onboard
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceResourceOnBoardRepresentation) +
				compartmentIdVariableStr + NameSpaceResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation2) +
				compartmentIdVariableStr + NameSpaceSingularDataSourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_onboarded"),
			),
		},
		// verify offboard
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceResourceOffBoardRepresentation) +
				compartmentIdVariableStr + NameSpaceResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "false"),
			),
		},
	})
}
