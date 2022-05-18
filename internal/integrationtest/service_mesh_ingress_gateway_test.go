// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	IngressGatewayRequiredOnlyResource = IngressGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Required, acctest.Create, ingressGatewayRepresentation)

	IngressGatewayResourceConfig = IngressGatewayResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Update, ingressGatewayRepresentation)

	ingressGatewaySingularDataSourceRepresentation = map[string]interface{}{
		"ingress_gateway_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_ingress_gateway.test_ingress_gateway.id}`},
	}

	ingressGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_ingress_gateway.test_ingress_gateway.id}`},
		"mesh_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_mesh.mesh1.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ingressGatewayDataSourceFilterRepresentation}}
	ingressGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_ingress_gateway.test_ingress_gateway.id}`}},
	}

	ingressGatewayRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"hosts":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ingressGatewayHostsRepresentation},
		"mesh_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_mesh.mesh1.id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"access_logging": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayAccessLoggingRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"mtls":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayMtlsRepresentation},
	}
	ingressGatewayHostsRepresentation = map[string]interface{}{
		"listeners": acctest.RepresentationGroup{RepType: acctest.Required, Group: ingressGatewayHostsListenersRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"hostnames": acctest.Representation{RepType: acctest.Optional, Create: []string{`hostnames`}, Update: []string{`hostnames2`}},
	}
	ingressGatewayAccessLoggingRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	ingressGatewayMtlsRepresentation = map[string]interface{}{
		"maximum_validity": acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `51`},
	}
	ingressGatewayHostsListenersRepresentation = map[string]interface{}{
		"port":     acctest.Representation{RepType: acctest.Required, Create: `8090`, Update: `8091`},
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TLS_PASSTHROUGH`},
		"tls":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayHostsListenersTlsRepresentation},
	}
	ingressGatewayHostsListenersTlsRepresentation = map[string]interface{}{
		"mode":               acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"client_validation":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayHostsListenersTlsClientValidationRepresentation},
		"server_certificate": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayHostsListenersTlsServerCertificateRepresentation},
	}
	ingressGatewayHostsListenersTlsClientValidationRepresentation = map[string]interface{}{
		"subject_alternate_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`subjectAlternateNames`}, Update: []string{`subjectAlternateNames2`}},
		"trusted_ca_bundle":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ingressGatewayHostsListenersTlsClientValidationTrustedCaBundleRepresentation},
	}
	ingressGatewayHostsListenersTlsServerCertificateRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `OCI_CERTIFICATES`, Update: `LOCAL_FILE`},
		"certificate_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.certificate_id}`},
		"secret_name":    acctest.Representation{RepType: acctest.Optional, Update: `${oci_vault_secret.secret_1.secret_name}`},
	}
	ingressGatewayHostsListenersTlsClientValidationTrustedCaBundleRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `OCI_CERTIFICATES`, Update: `LOCAL_FILE`},
		"ca_bundle_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_ca_bundle.ca_bundle_1.id}`},
		"secret_name":  acctest.Representation{RepType: acctest.Optional, Update: `${oci_vault_secret.secret_1.secret_name}`},
	}

	meshSecretSecretContentRepresentation = map[string]interface{}{
		"content_type": acctest.Representation{RepType: acctest.Required, Create: `BASE64`},
		"content":      acctest.Representation{RepType: acctest.Required, Create: `PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg==`},
		"name":         acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"stage":        acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
	}
	meshSecretSecretRulesRepresentation = map[string]interface{}{
		"rule_type":                                     acctest.Representation{RepType: acctest.Required, Create: `SECRET_EXPIRY_RULE`, Update: `SECRET_REUSE_RULE`},
		"is_enforced_on_deleted_secret_versions":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_secret_content_retrieval_blocked_on_expiry": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"secret_version_expiry_interval":                acctest.Representation{RepType: acctest.Optional, Create: `P3D`},
		"time_of_absolute_expiry":                       acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339)},
	}

	IngressGatewayResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "mesh1", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(meshRepresentation, map[string]interface{}{
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "ca_bundle_1", acctest.Required, acctest.Create, caBundleRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "secret_1", acctest.Required, acctest.Create, secretRepresentation)
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshIngressGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshIngressGatewayResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	certificateId := utils.GetEnvSettingWithBlankDefault("certificate_id")
	certificateIdVariableStr := fmt.Sprintf("variable \"certificate_id\" { default = \"%s\" }\n", certificateId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("key_id")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_service_mesh_ingress_gateway.test_ingress_gateway"
	datasourceName := "data.oci_service_mesh_ingress_gateways.test_ingress_gateways"
	singularDatasourceName := "data.oci_service_mesh_ingress_gateway.test_ingress_gateway"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+certificateIdVariableStr+certificateAuthorityIdVariableStr+vaultIdVariableStr+keyIdVariableStr+compartmentIdVariableStr+IngressGatewayResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Create, ingressGatewayRepresentation), "servicemesh", "ingressGateway", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshIngressGatewayDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + certificateIdVariableStr + certificateAuthorityIdVariableStr + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Required, acctest.Create, ingressGatewayRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.port", "8090"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateIdVariableStr + certificateAuthorityIdVariableStr + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + certificateIdVariableStr + certificateAuthorityIdVariableStr + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ingressGatewayRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.hostnames.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.port", "8090"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.subject_alternate_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.ca_bundle_id"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.type", "OCI_CERTIFICATES"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.type", "OCI_CERTIFICATES"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "mtls.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.maximum_validity", "50"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + certificateIdVariableStr + certificateAuthorityIdVariableStr + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ingressGatewayRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.hostnames.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.port", "8090"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.subject_alternate_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.ca_bundle_id"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.type", "OCI_CERTIFICATES"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.type", "OCI_CERTIFICATES"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "mtls.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.maximum_validity", "50"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + certificateIdVariableStr + certificateAuthorityIdVariableStr + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Update, acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"hosts.listeners.tls.client_validation.trusted_ca_bundle.ca_bundle_id", "hosts.listeners.tls.server_certificate.certificate_id"}, ingressGatewayRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.hostnames.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.port", "8091"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.protocol", "TLS_PASSTHROUGH"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.subject_alternate_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.secret_name"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.type", "LOCAL_FILE"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.mode", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.secret_name"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.type", "LOCAL_FILE"),
				resource.TestCheckResourceAttr(resourceName, "hosts.0.name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "mtls.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.maximum_validity", "51"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_ingress_gateways", "test_ingress_gateways", acctest.Optional, acctest.Update, ingressGatewayDataSourceRepresentation) +
				vaultIdVariableStr + keyIdVariableStr + certificateIdVariableStr + certificateAuthorityIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Update, acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"hosts.listeners.tls.client_validation.trusted_ca_bundle.ca_bundle_id", "hosts.listeners.tls.server_certificate.certificate_id"}, ingressGatewayRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "mesh_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "ingress_gateway_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ingress_gateway_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Required, acctest.Create, ingressGatewaySingularDataSourceRepresentation) +
				vaultIdVariableStr + keyIdVariableStr + certificateIdVariableStr + certificateAuthorityIdVariableStr + compartmentIdVariableStr + IngressGatewayResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "test_ingress_gateway", acctest.Optional, acctest.Update, acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"hosts.listeners.tls.client_validation.trusted_ca_bundle.ca_bundle_id", "hosts.listeners.tls.server_certificate.certificate_id"}, ingressGatewayRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ingress_gateway_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "access_logging.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.hostnames.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.port", "8091"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.protocol", "TLS_PASSTHROUGH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.client_validation.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.client_validation.0.subject_alternate_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.client_validation.0.trusted_ca_bundle.0.type", "LOCAL_FILE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.mode", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.server_certificate.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.listeners.0.tls.0.server_certificate.0.type", "LOCAL_FILE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosts.0.name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mtls.0.certificate_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mtls.0.maximum_validity", "51"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + IngressGatewayRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshIngressGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_ingress_gateway" {
			noResourceFound = false
			request := oci_service_mesh.GetIngressGatewayRequest{}

			tmp := rs.Primary.ID
			request.IngressGatewayId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetIngressGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.IngressGatewayLifecycleStateDeleted): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ServiceMeshIngressGateway") {
		resource.AddTestSweepers("ServiceMeshIngressGateway", &resource.Sweeper{
			Name:         "ServiceMeshIngressGateway",
			Dependencies: acctest.DependencyGraph["ingressGateway"],
			F:            sweepServiceMeshIngressGatewayResource,
		})
	}
}

func sweepServiceMeshIngressGatewayResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	ingressGatewayIds, err := getIngressGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, ingressGatewayId := range ingressGatewayIds {
		if ok := acctest.SweeperDefaultResourceId[ingressGatewayId]; !ok {
			deleteIngressGatewayRequest := oci_service_mesh.DeleteIngressGatewayRequest{}

			deleteIngressGatewayRequest.IngressGatewayId = &ingressGatewayId

			deleteIngressGatewayRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteIngressGateway(context.Background(), deleteIngressGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting IngressGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", ingressGatewayId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ingressGatewayId, ingressGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				ingressGatewaySweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getIngressGatewayIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IngressGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listIngressGatewaysRequest := oci_service_mesh.ListIngressGatewaysRequest{}
	listIngressGatewaysRequest.CompartmentId = &compartmentId
	active := "ACTIVE"
	listIngressGatewaysRequest.LifecycleState = &active
	listIngressGatewaysResponse, err := serviceMeshClient.ListIngressGateways(context.Background(), listIngressGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IngressGateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ingressGateway := range listIngressGatewaysResponse.Items {
		id := *ingressGateway.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IngressGatewayId", id)
	}
	return resourceIds, nil
}

func ingressGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ingressGatewayResponse, ok := response.Response.(oci_service_mesh.GetIngressGatewayResponse); ok {
		return ingressGatewayResponse.LifecycleState != oci_service_mesh.IngressGatewayLifecycleStateDeleted
	}
	return false
}

func ingressGatewaySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetIngressGateway(context.Background(), oci_service_mesh.GetIngressGatewayRequest{
		IngressGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
