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
	ContainerInstancesContainerInstancesContainerInstanceShapeSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	ContainerInstancesContainerInstancesContainerInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	ContainerInstancesContainerInstanceShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: container_instances/default
func TestContainerInstancesContainerInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerInstancesContainerInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_container_instances_container_instance_shapes.test_container_instance_shapes"
	singularDatasourceName := "data.oci_container_instances_container_instance_shape.test_container_instance_shape"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_container_instances_container_instance_shapes", "test_container_instance_shapes", acctest.Required, acctest.Create, ContainerInstancesContainerInstancesContainerInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerInstancesContainerInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "container_instance_shape_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "container_instance_shape_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_container_instances_container_instance_shape", "test_container_instance_shape", acctest.Required, acctest.Create, ContainerInstancesContainerInstancesContainerInstanceShapeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerInstancesContainerInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
			),
		},
	})
}
