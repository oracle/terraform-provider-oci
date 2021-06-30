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
	replicationStatusSingularDataSourceRepresentation = map[string]interface{}{
		"replication_id":      Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.replica_details[0].replication_id}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	ReplicationStatusResourceDependencies = KeyResourceDependencies
)

func TestKmsReplicationStatusResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsReplicationStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_replication_status.test_replication_status"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_replication_status", "test_replication_status", Required, Create, replicationStatusSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ReplicationStatusResourceDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
				),
			},
		},
	})
}
