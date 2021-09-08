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
	publicIpPoolCapacityRepresentation = map[string]interface{}{
		"public_ip_pool_id": Representation{repType: Required, create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
		"byoip_id":          Representation{repType: Required, create: `${var.byoip_range_id}`},
		"cidr_block":        Representation{repType: Required, create: `${var.public_ip_pool_cidr_block}`},
	}

	publicIpPoolCidrBlock            = getEnvSettingWithBlankDefault("public_ip_pool_cidr_block")
	publicIpPoolCidrBlockVariableStr = fmt.Sprintf("variable \"public_ip_pool_cidr_block\" { default = \"%s\" }\n", publicIpPoolCidrBlock)

	byoipRangeId            = getEnvSettingWithBlankDefault("byoip_range_ocid")
	byoipRangeIdVariableStr = fmt.Sprintf("variable \"byoip_range_id\" { default = \"%s\" }\n", byoipRangeId)

	PublicIpPoolAddCapacityResourceDependencies = publicIpPoolCidrBlockVariableStr + byoipRangeIdVariableStr + generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolRepresentation)
)

// issue-routing-tag: core/vcnip
func TestResourceCorePublicIpPoolCapacity_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpPoolCapacityResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_public_ip_pool_capacity.test_public_ip_pool_capacity"

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + PublicIpPoolAddCapacityResourceDependencies + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cidr_block", publicIpPoolCidrBlock),

				func(s *terraform.State) (err error) {
					_, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
