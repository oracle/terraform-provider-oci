// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InternalBastionRequiredOnlyResource = InternalBastionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, InternalBastionRepresentation)

	InternalBastionResourceConfig = InternalBastionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create, InternalBastionRepresentation)

	InternalBastionSingularDataSourceRepresentation = map[string]interface{}{
		"bastion_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
	}

	internalBastionName = utils.RandomString(15, utils.CharsetWithoutDigits)

	InternalBastionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"bastion_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"bastion_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: internalBastionName},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: InternalBastionDataSourceFilterRepresentation}}
	InternalBastionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bastion_bastion.test_bastion.id}`}},
	}

	InternalBastionRepresentation = map[string]interface{}{
		"bastion_type":                  acctest.Representation{RepType: acctest.Required, Create: `INTERNAL`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_subnet_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"name":                          acctest.Representation{RepType: acctest.Required, Create: internalBastionName},
		"phone_book_entry":              acctest.Representation{RepType: acctest.Required, Create: `OCIBastion`},
		"static_jump_host_ip_addresses": acctest.Representation{RepType: acctest.Optional, Create: []string{`10.0.0.3`}},
	}

	InternalBastionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: bastion/default
func TestBastionBastionResource_internal(t *testing.T) {
	httpreplay.SetScenario("TestBastionBastionResource_internal")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_bastion.test_bastion"
	datasourceName := "data.oci_bastion_bastions.test_bastions"
	singularDatasourceName := "data.oci_bastion_bastion.test_bastion"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InternalBastionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create, InternalBastionRepresentation), "bastion", "bastion", t)

	acctest.ResourceTest(t, testAccCheckBastionBastionDestroy, []resource.TestStep{

		// verify Create
		{
			Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, InternalBastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "INTERNAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + InternalBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create, InternalBastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "INTERNAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", internalBastionName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "phone_book_entry", "OCIBastion"),
				resource.TestCheckResourceAttr(resourceName, "static_jump_host_ip_addresses.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_bastions", "test_bastions", acctest.Optional, acctest.Create, InternalBastionDataSourceRepresentation) +
				compartmentIdVariableStr + InternalBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, InternalBastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
				resource.TestCheckResourceAttr(datasourceName, "bastion_lifecycle_state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", internalBastionName),

				resource.TestCheckResourceAttr(datasourceName, "bastions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.bastion_type", "INTERNAL"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.name", internalBastionName),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, InternalBastionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InternalBastionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bastion_type", "INTERNAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", internalBastionName),
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
	})
}
