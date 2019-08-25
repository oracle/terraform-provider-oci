// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
	}

	InstanceCredentialResourceConfig = OciWindowsImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig
)

func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceCredentialResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_credentials", "test_instance_credentials", Required, Create, instanceCredentialSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "password"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "username"),
				),
			},
		},
	})
}
