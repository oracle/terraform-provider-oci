// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation2 = map[string]interface{}{
		"namespace": Representation{RepType: Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	namespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOnBoardRepresentation = map[string]interface{}{
		"namespace":      Representation{RepType: Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   Representation{RepType: Required, Create: `true`},
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOffBoardRepresentation = map[string]interface{}{
		"namespace":      Representation{RepType: Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   Representation{RepType: Required, Create: `false`},
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	NameSpaceResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", Required, Create, namespaceDataSourceRepresentation)

	NameSpaceSingularDataSourceDependencies = NameSpaceResourceDependencies + GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOnBoardRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceResource_basic(t *testing.T) {
	t.Skip("skipping test as onboarding tenancy is a one time operation only and cannot be done on a recurring basis")
	httpreplay.SetScenario("TestLogAnalyticsNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	tenancyId := GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace.test_namespace"
	datasourceName := "data.oci_log_analytics_namespaces.test_namespaces"
	singularDatasourceName := "data.oci_log_analytics_namespace.test_namespace"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource also works as a dependency for next step
		{
			Config: config + compartmentIdVariableStr +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", Required, Create, namespaceDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttrSet(datasourceName, "namespace_collection.#"),
			),
		},
		// verify onboard
		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOnBoardRepresentation) +
				compartmentIdVariableStr + NameSpaceResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation2) +
				compartmentIdVariableStr + NameSpaceSingularDataSourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_onboarded"),
			),
		},
		// verify offboard
		{
			Config: config +
				GenerateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOffBoardRepresentation) +
				compartmentIdVariableStr + NameSpaceResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_onboarded", "false"),
			),
		},
	})
}
