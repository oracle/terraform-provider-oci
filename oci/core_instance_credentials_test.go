// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	InstanceCredentialResourceConfig = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		OciWindowsImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_instance_credentials", "test_instance_credentials", Required, Create, instanceCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceCredentialResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "password"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "username"),
			),
		},
	})
}
