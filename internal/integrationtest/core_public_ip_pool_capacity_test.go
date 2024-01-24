// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PublicIpPoolCapacityRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation)

	publicIpPoolCapacityRepresentation = map[string]interface{}{
		"public_ip_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
		"byoip_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.byoip_range_id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `${var.public_ip_pool_cidr_block}`},
	}

	publicIpPoolCidrBlock            = utils.GetEnvSettingWithBlankDefault("public_ip_pool_cidr_block")
	publicIpPoolCidrBlockVariableStr = fmt.Sprintf("variable \"public_ip_pool_cidr_block\" { default = \"%s\" }\n", publicIpPoolCidrBlock)

	byoipRangeId            = utils.GetEnvSettingWithBlankDefault("byoip_range_ocid")
	byoipRangeIdVariableStr = fmt.Sprintf("variable \"byoip_range_id\" { default = \"%s\" }\n", byoipRangeId)

	PublicIpPoolAddCapacityResourceDependencies = publicIpPoolCidrBlockVariableStr + byoipRangeIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Required, acctest.Create, CorePublicPoolRepresentation)
)

// issue-routing-tag: core/vcnip
func TestResourceCorePublicIpPoolCapacity_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpPoolCapacityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_public_ip_pool_capacity.test_public_ip_pool_capacity"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + PublicIpPoolAddCapacityResourceDependencies + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", publicIpPoolCidrBlock),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + PublicIpPoolCapacityRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
