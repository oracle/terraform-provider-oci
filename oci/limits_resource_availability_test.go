// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	resourceAvailabilitySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"limit_name":          Representation{repType: Required, create: `custom-image-count`},
		"service_name":        Representation{repType: Required, create: `${data.oci_limits_services.test_services.services.0.name}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	ResourceAvailabilityResourceConfig = AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_limits_services", "test_services", Required, Create, limitsServiceDataSourceRepresentation)
)

func TestLimitsResourceAvailabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsResourceAvailabilityResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	singularDatasourceName := "data.oci_limits_resource_availability.test_resource_availability"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_limits_resource_availability", "test_resource_availability", Required, Create, resourceAvailabilitySingularDataSourceRepresentation) +
					compartmentIdVariableStr + ResourceAvailabilityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "limit_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				),
			},
		},
	})
}
