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
	OnesubscriptionOnesubscriptionCommitmentSingularDataSourceRepresentation = map[string]interface{}{
		"commitment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.commitment_id}`},
	}

	OnesubscriptionOnesubscriptionCommitmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscribed_service_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscribed_service_id}`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionCommitmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionCommitmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subsServId := utils.GetEnvSettingWithBlankDefault("subscribed_service_id")
	subsServIdVariableStr := fmt.Sprintf("variable \"subscribed_service_id\" { default = \"%s\" }\n", subsServId)

	commitmentId := utils.GetEnvSettingWithBlankDefault("commitment_id")
	commitmentIdVariableStr := fmt.Sprintf("variable \"commitment_id\"  { default = \"%s\" }\n", commitmentId)

	datasourceName := "data.oci_onesubscription_commitments.test_commitments"
	singularDatasourceName := "data.oci_onesubscription_commitment.test_commitment"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_commitments", "test_commitments", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionCommitmentDataSourceRepresentation) +
				compartmentIdVariableStr + subsServIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_service_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "commitments.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.available_amount"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.funded_allocation_value"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.subscribed_service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.time_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "commitments.0.used_amount"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_commitment", "test_commitment", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionCommitmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + commitmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "commitment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_amount"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "funded_allocation_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quantity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used_amount"),
			),
		},
	})
}
