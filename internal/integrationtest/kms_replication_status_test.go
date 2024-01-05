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
	KmsKmsReplicationStatusSingularDataSourceRepresentation = map[string]interface{}{
		"replication_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.replica_details[0].replication_id}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	KmsKmsReplicationStatusResourceDependencies = KmsKeyResourceDependencies
)

// issue-routing-tag: kms/default
func TestKmsReplicationStatusResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsReplicationStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_replication_status.test_replication_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_replication_status", "test_replication_status", acctest.Required, acctest.Create, KmsKmsReplicationStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsKmsReplicationStatusResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
			),
		},
	})
}
