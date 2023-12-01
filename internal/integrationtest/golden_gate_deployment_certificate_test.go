// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

/*
 Note:
  This test will create a deployment and a certificate for it. Then it will create a new certificate (same content, but with new key) and delete the old one.
  Set the following environmentVariables in order to make it work:
 - for deployment creation
    TF_VAR_compartment_id=compartment_id_of_the_deployment
    TF_VAR_test_subnet_id=test_subnet_id
    TF_VAR_password=password

 - for certificate (variable should contain base64 encoded string)
	export TF_VAR_certificate_content=`cat /path/to/certificate.pem|base64`
*/

const (
	COMPARTMENT_ID              = "compartment_id"
	TEST_SUBNET_ID              = "test_subnet_id"
	PASSWORD                    = "password"
	TRUSTED_CERTIFICATE_CONTENT = "certificate_content"
)

var (
	DeploymentRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `1`},
		"deployment_type":         acctest.Representation{RepType: acctest.Required, Create: `DATABASE_ORACLE`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `TF_Certificate_test`},
		"is_auto_scaling_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.test_subnet_id}`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"is_public":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ogg_data": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"admin_password":  acctest.Representation{RepType: acctest.Required, Create: `${var.password}`},
			"admin_username":  acctest.Representation{RepType: acctest.Required, Create: `adminUsername`},
			"deployment_name": acctest.Representation{RepType: acctest.Required, Create: `depl_test_ggs_deployment_name`},
		}},
	}

	CertificateRepresentation = map[string]interface{}{
		"certificate_content": acctest.Representation{RepType: acctest.Required, Create: `${var.certificate_content}`},
		"deployment_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"key":                 acctest.Representation{RepType: acctest.Required, Create: `MyCertificate`, Update: `MyNewCertificate`},
	}

	CertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment_certificate.test_deployment_certificate.key}`},
		"deployment_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
	}

	CertificateDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: string(oci_golden_gate.CertificateLifecycleStateActive)},
	}
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(TEST_SUBNET_ID, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(TRUSTED_CERTIFICATE_CONTENT, t) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create, DeploymentRepresentation)

	var (
		certificateContent = utils.GetEnvSettingWithBlankDefault(TRUSTED_CERTIFICATE_CONTENT)
		resId              string
		resId2             string
	)

	resourceName := "oci_golden_gate_deployment_certificate.test_deployment_certificate"
	datasourceName := "data.oci_golden_gate_deployment_certificates.test_deployment_certificates"
	singularDatasourceName := "data.oci_golden_gate_deployment_certificate.test_deployment_certificate_singular"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create, DeploymentRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Create, CertificateRepresentation), "goldengate", "deploymentCertificate", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentCertificateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Create, CertificateRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_content", certificateContent),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "key", "MyCertificate"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify recreate (previously created has been deleted and a new one has been created)
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Update, CertificateRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_content", certificateContent),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "key", "MyNewCertificate"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource is not recreated, when it was supposed to be")
					}
					resId = resId2
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Update, CertificateRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_certificates", "test_deployment_certificates", acctest.Optional, acctest.Update, CertificateDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", string(oci_golden_gate.CertificateLifecycleStateActive)),
				resource.TestCheckResourceAttr(datasourceName, "certificate_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.0.key", "MyNewCertificate"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Update, CertificateRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate_singular", acctest.Required, acctest.Create, CertificateSingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "authority_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_content"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_ca"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_self_signed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "issuer"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "md5hash"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key", "MyNewCertificate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_key_algorithm"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_key_size"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serial"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sha1hash"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subject"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subject_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_valid_from"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_valid_to"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_certificate", "test_deployment_certificate", acctest.Required, acctest.Update, CertificateRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGoldenGateDeploymentCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_deployment_certificate" {
			noResourceFound = false
			request := oci_golden_gate.GetCertificateRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.CertificateKey = &value
			}

			if value, ok := rs.Primary.Attributes["deployment_id"]; ok {
				request.DeploymentId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetCertificate(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.CertificateLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
