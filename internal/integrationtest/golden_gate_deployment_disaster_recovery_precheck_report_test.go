// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	FAULT_DOMAIN         = "fault_domain"
	AVAILABILITY_DOMAIN  = "availability_domain"
	SOURCE_DEPLOYMENT_ID = "source_deployment_id"
	PASSWORD_SECRET_ID   = "password_secret_id"
)

var (
	GoldenGateDeploymentDisasterRecoveryPrecheckReportSingularDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_golden_gate_deployment.source_deployment.availability_domain}`},
		"deployment_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"fault_domain":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_golden_gate_deployment.source_deployment.fault_domain}`},
	}

	namespaceSingularDataSourceRepresentation = map[string]interface{}{
		//"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	sourceDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
	}

	oggDataRepresentationForGoldenGate = map[string]interface{}{
		"deployment_name":    acctest.Representation{RepType: acctest.Required, Create: `Test`},
		"credential_store":   acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
		"admin_username":     acctest.Representation{RepType: acctest.Required, Create: `oggadmin`, Update: `adminUsername2`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.password_secret_id}`, Update: `${var.password_secret_id_2}`},
	}
	GoldenGateDeploymentDisasterRecoveryPrecheckReportResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Optional,
		acctest.Create, acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
			"source_deployment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.source_deployment_id}`},
			"deployment_type":      acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_ORACLE`},
			"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: `${var.availability_domain}`},
			"fault_domain":         acctest.Representation{RepType: acctest.Optional, Create: `${var.fault_domain}`},
			"ogg_data":             acctest.RepresentationGroup{RepType: acctest.Required, Group: oggDataRepresentationForGoldenGate},
		})) +

		acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment", "source_deployment", acctest.Required, acctest.Create, sourceDeploymentSingularDataSourceRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, IdentityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentDisasterRecoveryPrecheckReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentDisasterRecoveryPrecheckReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr(AVAILABILITY_DOMAIN, t) +
		makeVariableStr(FAULT_DOMAIN, t) +
		makeVariableStr(SOURCE_DEPLOYMENT_ID, t) +
		makeVariableStr(PASSWORD_SECRET_ID, t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	password := utils.GetEnvSettingWithBlankDefault("password")
	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")

	availabilityDomain := utils.GetEnvSettingWithBlankDefault("availability_domain")
	faultDomain := utils.GetEnvSettingWithBlankDefault("fault_domain")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	passwordVariableStr := fmt.Sprintf("variable \"password\" { default = \"%s\" }\n", password)
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)
	availabilityDomainStr := fmt.Sprintf("variable \"availabilityDomain\" { default = \"%s\" }\n", availabilityDomain)
	faultDomainStr := fmt.Sprintf("variable \"faultDomain\" { default = \"%s\" }\n", faultDomain)

	singularDatasourceName := "data.oci_golden_gate_deployment_disaster_recovery_precheck_report.test_deployment_disaster_recovery_precheck_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_disaster_recovery_precheck_report", "test_deployment_disaster_recovery_precheck_report", acctest.Required, acctest.Create, GoldenGateDeploymentDisasterRecoveryPrecheckReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + passwordVariableStr + subnetIdVariableStr + availabilityDomainStr + faultDomainStr + GoldenGateDeploymentDisasterRecoveryPrecheckReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),

				resource.TestCheckResourceAttr(singularDatasourceName, "checks.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "precheck_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_precheck_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_precheck_started"),
			),
		},
	})
}
