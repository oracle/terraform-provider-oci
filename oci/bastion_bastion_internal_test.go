// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InternalBastionRequiredOnlyResource = InternalBastionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, InternalBastionRepresentation)

	InternalBastionResourceConfig = InternalBastionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create, InternalBastionRepresentation)

	InternalBastionSingularDataSourceRepresentation = map[string]interface{}{
		"bastion_id": Representation{repType: Required, create: `${oci_bastion_bastion.test_bastion.id}`},
	}

	InternalBastionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"bastion_id":              Representation{repType: Optional, create: `${oci_bastion_bastion.test_bastion.id}`},
		"bastion_lifecycle_state": Representation{repType: Optional, create: `ACTIVE`},
		"name":                    Representation{repType: Optional, create: `bastionterraformtest`},
		"filter":                  RepresentationGroup{Required, InternalBastionDataSourceFilterRepresentation}}
	InternalBastionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_bastion_bastion.test_bastion.id}`}},
	}

	InternalBastionRepresentation = map[string]interface{}{
		"bastion_type":                  Representation{repType: Required, create: `INTERNAL`},
		"compartment_id":                Representation{repType: Required, create: `${var.compartment_id}`},
		"target_subnet_id":              Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"name":                          Representation{repType: Required, create: `bastionterraformtest`},
		"phone_book_entry":              Representation{repType: Required, create: `OCIBastion`},
		"static_jump_host_ip_addresses": Representation{repType: Optional, create: []string{`10.0.0.3`}},
	}

	InternalBastionResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestBastionBastionResource_internal(t *testing.T) {
	httpreplay.SetScenario("TestBastionBastionResource_internal")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_bastion.test_bastion"
	datasourceName := "data.oci_bastion_bastions.test_bastions"
	singularDatasourceName := "data.oci_bastion_bastion.test_bastion"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+InternalBastionResourceDependencies+
		generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create, InternalBastionRepresentation), "bastion", "bastion", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBastionBastionDestroy,
		Steps: []resource.TestStep{

			// verify create
			{
				Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, InternalBastionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bastion_type", "INTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create, InternalBastionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bastion_type", "INTERNAL"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "bastionterraformtest"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "phone_book_entry", "OCIBastion"),
					resource.TestCheckResourceAttr(resourceName, "static_jump_host_ip_addresses.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_bastion_bastions", "test_bastions", Optional, Create, InternalBastionDataSourceRepresentation) +
					compartmentIdVariableStr + InternalBastionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, InternalBastionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
					resource.TestCheckResourceAttr(datasourceName, "bastion_lifecycle_state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "bastionterraformtest"),

					resource.TestCheckResourceAttr(datasourceName, "bastions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "bastions.0.bastion_type", "INTERNAL"),
					resource.TestCheckResourceAttr(datasourceName, "bastions.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "bastions.0.name", "bastionterraformtest"),
					resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_vcn_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, InternalBastionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InternalBastionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "bastion_type", "INTERNAL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "bastionterraformtest"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_vcn_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + InternalBastionResourceConfig,
			},

			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
