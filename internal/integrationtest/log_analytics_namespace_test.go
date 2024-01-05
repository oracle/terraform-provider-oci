// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	LogAnalyticsLogAnalyticsNamespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	LogAnalyticsLogAnalyticsNamespaceResourceOnBoardRepresentation = map[string]interface{}{
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	LogAnalyticsLogAnalyticsNamespaceResourceOffBoardRepresentation = map[string]interface{}{
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	LogAnalyticsLogAnalyticsNameSpaceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceDataSourceRepresentation)

	LogAnalyticsLogAnalyticsNameSpaceSingularDataSourceDependencies = LogAnalyticsLogAnalyticsNameSpaceResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceResourceOnBoardRepresentation)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttrSet(datasourceName, "namespace_collection.#"),
			),
		},
		// verify onboard
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceResourceOnBoardRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsNameSpaceResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsNameSpaceSingularDataSourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_onboarded"),
			),
		},
		// verify offboard
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceResourceOffBoardRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsNameSpaceResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "false"),
			),
		},
	})
}
