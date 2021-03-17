// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation2 = map[string]interface{}{
		"namespace": Representation{repType: Required, create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	namespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOnBoardRepresentation = map[string]interface{}{
		"namespace":      Representation{repType: Required, create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   Representation{repType: Required, create: `true`},
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	namespaceResourceOffBoardRepresentation = map[string]interface{}{
		"namespace":      Representation{repType: Required, create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
		"is_onboarded":   Representation{repType: Required, create: `false`},
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	NameSpaceResourceDependencies = generateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", Required, Create, namespaceDataSourceRepresentation)

	NameSpaceSingularDataSourceDependencies = NameSpaceResourceDependencies + generateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOnBoardRepresentation)
)

func TestLogAnalyticsNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace.test_namespace"
	datasourceName := "data.oci_log_analytics_namespaces.test_namespaces"
	singularDatasourceName := "data.oci_log_analytics_namespace.test_namespace"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource also works as a dependency for next step
			{
				Config: config + compartmentIdVariableStr +
					generateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", Required, Create, namespaceDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttrSet(datasourceName, "namespace_collection.#"),
				),
			},
			// verify onboard
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOnBoardRepresentation) +
					compartmentIdVariableStr + NameSpaceResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "is_onboarded", "true"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation2) +
					compartmentIdVariableStr + NameSpaceSingularDataSourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_onboarded"),
				),
			},
			// verify offboard
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_log_analytics_namespace", "test_namespace", Required, Create, namespaceResourceOffBoardRepresentation) +
					compartmentIdVariableStr + NameSpaceResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "is_onboarded", "false"),
				),
			},
		},
	})
}
