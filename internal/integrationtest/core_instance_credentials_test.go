// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	instanceCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	InstanceCredentialResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		utils.OciWindowsImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_credentials", "test_instance_credentials", acctest.Required, acctest.Create, instanceCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "password"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "username"),
			),
		},
	})
}
