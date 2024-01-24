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
	CoreCoreInstancePoolLoadBalancerAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"instance_pool_load_balancer_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.load_balancers.0.id}`},
	}

	CoreInstancePoolLoadBalancerAttachmentResourceConfig = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_configuration", "test_instance_configuration", acctest.Optional, acctest.Create, CoreInstancePoolConfigurationPoolRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Optional, acctest.Update, CoreInstancePoolRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: core/computeManagement
func TestCoreInstancePoolLoadBalancerAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolLoadBalancerAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_pool_load_balancer_attachment.test_instance_pool_load_balancer_attachment"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_pool_load_balancer_attachment", "test_instance_pool_load_balancer_attachment", acctest.Required, acctest.Create, CoreCoreInstancePoolLoadBalancerAttachmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreInstancePoolLoadBalancerAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_load_balancer_attachment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_selection"),
			),
		},
	})
}
