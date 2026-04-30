// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

var (
	BdsCertificateFlowClusterAdminPassword = "T3JhY2xlVGVhbVVTQSExMjM="

	BdsPrimaryOciCertificateConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `tfAccOciCertCfg1`},
		"certificate_type":         acctest.Representation{RepType: acctest.Required, Create: `OCI_CERTIFICATE`},
		"certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: `${var.certificate_authority_id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	BdsBdsInstanceCertificateFlowResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceBdsCertificateConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceBdsCertificateConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_compartment_ocid")
	if compartmentId == "" {
		compartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	}
	if compartmentId == "" {
		t.Skip("compartment_ocid must be set")
	}
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	if subnetId == "" {
		t.Skip("subnet_ocid must be set")
	}
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	secretId := utils.GetEnvSettingWithBlankDefault("secret_ocid")
	clusterAdminPassword := utils.GetEnvSettingWithDefault("bds_cluster_admin_password", BdsCertificateFlowClusterAdminPassword)
	if secretId == "" && clusterAdminPassword == "" {
		t.Skip("either secret_ocid or bds_cluster_admin_password must be set")
	}
	secretIdVariableStr := ""
	if secretId != "" {
		secretIdVariableStr = fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)
	}

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	if certificateAuthorityId == "" {
		certificateAuthorityId = utils.GetEnvSettingWithBlankDefault("issuer_ca_ocid")
	}
	if certificateAuthorityId == "" {
		t.Skip("certificate_authority_id or issuer_ca_ocid must be set")
	}
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	rpRepresentation := map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tfCertFlowRp`},
		"session_token_life_span_duration_in_hours": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"force_refresh_resource_principal_trigger":  acctest.Representation{RepType: acctest.Required, Create: `0`},
	}

	if secretId != "" {
		rpRepresentation["secret_id"] = acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`}
	} else {
		rpRepresentation["cluster_admin_password"] = acctest.Representation{RepType: acctest.Required, Create: clusterAdminPassword}
	}

	rpResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Required, acctest.Create, rpRepresentation)

	primaryCreateRepresentation := cloneRepresentationMap(BdsPrimaryOciCertificateConfigurationRepresentation)
	if secretId != "" {
		primaryCreateRepresentation["secret_id"] = acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`}
	} else {
		primaryCreateRepresentation["cluster_admin_password"] = acctest.Representation{RepType: acctest.Optional, Create: clusterAdminPassword}
	}

	primarySetDefaultRepresentation := cloneRepresentationMap(primaryCreateRepresentation)
	primarySetDefaultRepresentation["set_default_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primarySetDefaultRepresentation["timeouts"] = acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `45m`, Update: `45m`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
	}}

	primaryIssueRepresentation := cloneRepresentationMap(primaryCreateRepresentation)
	primaryIssueRepresentation["is_missing_nodes_only"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`}
	primaryIssueRepresentation["set_default_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primaryIssueRepresentation["issue_certificate_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primaryIssueRepresentation["timeouts"] = acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `45m`, Update: `45m`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
	}}

	primaryRenewRepresentation := cloneRepresentationMap(primaryCreateRepresentation)
	primaryRenewRepresentation["is_missing_nodes_only"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`}
	primaryRenewRepresentation["set_default_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primaryRenewRepresentation["issue_certificate_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primaryRenewRepresentation["renew_certificate_trigger"] = acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`}
	primaryRenewRepresentation["timeouts"] = acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `45m`, Update: `45m`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `30m`, Update: `30m`},
	}}

	primaryCertConfigResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance_bds_certificate_configuration", "test_primary_bds_certificate_configuration", acctest.Required, acctest.Create, primaryCreateRepresentation)

	setDefaultPrimaryResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance_bds_certificate_configuration", "test_primary_bds_certificate_configuration", acctest.Optional, acctest.Update, primarySetDefaultRepresentation)

	generatePrimaryResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance_bds_certificate_configuration", "test_primary_bds_certificate_configuration", acctest.Optional, acctest.Update, primaryIssueRepresentation)

	renewPrimaryResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_bds_bds_instance_bds_certificate_configuration", "test_primary_bds_certificate_configuration", acctest.Optional, acctest.Update, primaryRenewRepresentation)

	primaryResourceName := "oci_bds_bds_instance_bds_certificate_configuration.test_primary_bds_certificate_configuration"
	clusterResourceName := "oci_bds_bds_instance.test_bds_instance"
	rpResourceName := "oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration"

	baseConfig := config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceCertificateFlowResourceDependencies

	configWithRp := baseConfig + secretIdVariableStr + rpResource

	configWithPrimaryCertConfig := configWithRp + certificateAuthorityIdVariableStr + primaryCertConfigResource

	configWithPrimarySetDefault := configWithRp + certificateAuthorityIdVariableStr + setDefaultPrimaryResource

	configWithPrimaryGenerate := configWithRp + certificateAuthorityIdVariableStr + generatePrimaryResource

	configWithPrimaryRenew := configWithRp + certificateAuthorityIdVariableStr + renewPrimaryResource

	configWithPrimaryCertConfigAndRp := configWithRp + certificateAuthorityIdVariableStr + primaryCertConfigResource

	cleanupRemoveCertConfig := configWithRp
	cleanupRemoveRp := baseConfig
	cleanupRemoveCluster := config + compartmentIdVariableStr + subnetIdVariableStr

	acctest.SaveConfigContent(configWithPrimaryCertConfigAndRp, "bds", "bdsCertificateConfigurationFullTriggerFlow", t)

	var clusterId string
	var primaryOciCertConfigId string
	var selfSignedOciCertConfigId string
	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: baseConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(clusterResourceName, "id"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "bds_cluster_version_summary.#"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "cluster_details.#"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "cluster_version"),
				resource.TestCheckResourceAttr(clusterResourceName, "cluster_admin_password", BdsCertificateFlowClusterAdminPassword),
				resource.TestCheckResourceAttr(clusterResourceName, "is_secure", "true"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "state"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "time_created"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(clusterResourceName, "nodes.#"),
				func(s *terraform.State) error {
					var err error
					clusterId, err = acctest.FromInstanceState(s, clusterResourceName, "id")
					if err != nil {
						return err
					}

					selfSignedOciCertConfigId, err = findBdsCertificateConfigurationIdByType(clusterId, "SELF_SIGNED")
					if err != nil {
						return err
					}

					return nil
				},
			),
		},
		{
			Config: configWithRp,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(rpResourceName, "id"),
				resource.TestCheckResourceAttrSet(rpResourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(rpResourceName, "display_name", "tfCertFlowRp"),
				resource.TestCheckResourceAttr(rpResourceName, "session_token_life_span_duration_in_hours", "1"),
				resource.TestCheckResourceAttrSet(rpResourceName, "state"),
				resource.TestCheckResourceAttrSet(rpResourceName, "time_created"),
				resource.TestCheckResourceAttrSet(rpResourceName, "time_updated"),
				func(s *terraform.State) error {
					return waitForBdsInstanceActive(clusterId, 30*time.Minute)
				},
			),
		},
		{
			Config: configWithPrimaryCertConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(primaryResourceName, "id"),
				resource.TestCheckResourceAttr(primaryResourceName, "display_name", "tfAccOciCertCfg1"),
				resource.TestCheckResourceAttr(primaryResourceName, "certificate_type", "OCI_CERTIFICATE"),
				resource.TestCheckResourceAttr(primaryResourceName, "certificate_authority_id", certificateAuthorityId),
				resource.TestCheckResourceAttr(primaryResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(primaryResourceName, "state"),
				resource.TestCheckResourceAttrSet(primaryResourceName, "time_created"),
				resource.TestCheckResourceAttrSet(primaryResourceName, "time_updated"),
				func(s *terraform.State) error {
					compositeId, err := acctest.FromInstanceState(s, primaryResourceName, "id")
					if err != nil {
						return err
					}
					_, primaryOciCertConfigId, err = parseCertificateConfigurationCompositeIdForTest(compositeId)
					if err != nil {
						return err
					}
					return waitForBdsInstanceActive(clusterId, 30*time.Minute)
				},
			),
		},
		{
			Config: configWithPrimarySetDefault,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(primaryResourceName, "id"),
				resource.TestCheckResourceAttr(primaryResourceName, "set_default_trigger", "true"),
				func(s *terraform.State) error {
					if err := waitForBdsInstanceActive(clusterId, 30*time.Minute); err != nil {
						return err
					}
					return waitForBdsCertificateConfigurationDefault(clusterId, primaryOciCertConfigId, true, 30*time.Minute)
				},
			),
		},
		{
			Config: configWithPrimaryGenerate,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(primaryResourceName, "id"),
				resource.TestCheckResourceAttr(primaryResourceName, "issue_certificate_trigger", "true"),
				resource.TestCheckResourceAttr(primaryResourceName, "is_missing_nodes_only", "false"),
				func(s *terraform.State) error {
					return waitForBdsInstanceActive(clusterId, 45*time.Minute)
				},
			),
		},
		{
			Config: configWithPrimaryRenew,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(primaryResourceName, "id"),
				resource.TestCheckResourceAttr(primaryResourceName, "renew_certificate_trigger", "true"),
				func(s *terraform.State) error {
					return waitForBdsInstanceActive(clusterId, 45*time.Minute)
				},
			),
		},
		{
			Config: configWithPrimaryCertConfigAndRp,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) error {
					if selfSignedOciCertConfigId == "" {
						var err error
						selfSignedOciCertConfigId, err = waitForBdsCertificateConfigurationByType(clusterId, "SELF_SIGNED", 30*time.Minute)
						if err != nil {
							return err
						}
					}

					if err := setDefaultBdsCertificateConfigurationForTest(clusterId, selfSignedOciCertConfigId, clusterAdminPassword, secretId); err != nil {
						return err
					}

					if err := waitForBdsInstanceActive(clusterId, 45*time.Minute); err != nil {
						return err
					}

					return waitForBdsCertificateConfigurationDefault(clusterId, selfSignedOciCertConfigId, true, 30*time.Minute)
				},
			),
		},
		{
			Config: cleanupRemoveCertConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(clusterResourceName, "id"),
				resource.TestCheckResourceAttrSet(rpResourceName, "id"),
				func(s *terraform.State) error {
					return waitForBdsInstanceActive(clusterId, 45*time.Minute)
				},
			),
		},
		{
			Config: cleanupRemoveRp,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(clusterResourceName, "id"),
				func(s *terraform.State) error {
					return waitForBdsInstanceActive(clusterId, 45*time.Minute)
				},
			),
		},
		{
			Config: cleanupRemoveCluster,
		},
	})
}

func cloneRepresentationMap(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func parseCertificateConfigurationCompositeIdForTest(compositeId string) (bdsInstanceId string, bdsCertificateConfigurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	if len(parts) != 4 {
		return "", "", fmt.Errorf("illegal compositeId %s encountered", compositeId)
	}
	return parts[1], parts[3], nil
}

func findBdsCertificateConfigurationIdByType(bdsInstanceId, expectedType string) (string, error) {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return "", err
	}

	request := oci_bds.ListBdsCertificateConfigurationsRequest{
		BdsInstanceId: &bdsInstanceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: getCertificateFlowRetryPolicy(),
		},
	}

	for {
		response, err := bdsClient.ListBdsCertificateConfigurations(context.Background(), request)
		if err != nil {
			return "", err
		}

		for _, item := range response.Items {
			if strings.EqualFold(string(item.Type), expectedType) && item.Id != nil {
				return *item.Id, nil
			}
		}

		if response.OpcNextPage == nil || *response.OpcNextPage == "" {
			break
		}

		request.Page = response.OpcNextPage
	}

	return "", fmt.Errorf("could not find bds certificate configuration with type %s for instance %s", expectedType, bdsInstanceId)
}

func waitForBdsCertificateConfigurationDefault(bdsInstanceId, certConfigId string, expected bool, timeout time.Duration) error {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return err
	}

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		request := oci_bds.GetBdsCertificateConfigurationRequest{
			BdsInstanceId:                 &bdsInstanceId,
			BdsCertificateConfigurationId: &certConfigId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: getCertificateFlowRetryPolicy(),
			},
		}

		response, err := bdsClient.GetBdsCertificateConfiguration(context.Background(), request)
		if err != nil {
			lastErr = err
			time.Sleep(20 * time.Second)
			continue
		}

		actual := false
		if response.IsDefaultConfiguration != nil {
			actual = *response.IsDefaultConfiguration
		}

		if actual == expected {
			return nil
		}

		lastErr = fmt.Errorf("certificate configuration %s default flag is %t, expected %t", certConfigId, actual, expected)
		time.Sleep(20 * time.Second)
	}

	if lastErr != nil {
		return lastErr
	}
	return fmt.Errorf("timed out waiting for certificate configuration %s default flag to become %t", certConfigId, expected)
}

func waitForBdsCertificateConfigurationByType(bdsInstanceId, expectedType string, timeout time.Duration) (string, error) {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return "", err
	}

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		request := oci_bds.ListBdsCertificateConfigurationsRequest{
			BdsInstanceId: &bdsInstanceId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: getCertificateFlowRetryPolicy(),
			},
		}

		for {
			response, err := bdsClient.ListBdsCertificateConfigurations(context.Background(), request)
			if err != nil {
				lastErr = err
				break
			}

			for _, item := range response.Items {
				if strings.EqualFold(string(item.Type), expectedType) && item.Id != nil {
					return *item.Id, nil
				}
			}

			if response.OpcNextPage == nil || *response.OpcNextPage == "" {
				break
			}
			request.Page = response.OpcNextPage
		}

		lastErr = fmt.Errorf("could not find bds certificate configuration with type %s for instance %s", expectedType, bdsInstanceId)
		time.Sleep(20 * time.Second)
	}

	return "", lastErr
}

func waitForBdsInstanceActive(bdsInstanceId string, timeout time.Duration) error {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return err
	}

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		request := oci_bds.GetBdsInstanceRequest{
			BdsInstanceId: &bdsInstanceId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: getCertificateFlowRetryPolicy(),
			},
		}

		response, err := bdsClient.GetBdsInstance(context.Background(), request)
		if err != nil {
			lastErr = err
			time.Sleep(20 * time.Second)
			continue
		}

		state := string(response.LifecycleState)
		if strings.EqualFold(state, "ACTIVE") {
			return nil
		}

		lastErr = fmt.Errorf("bds instance %s is in state %s", bdsInstanceId, state)
		time.Sleep(20 * time.Second)
	}

	if lastErr != nil {
		return lastErr
	}
	return fmt.Errorf("timed out waiting for bds instance %s to become ACTIVE", bdsInstanceId)
}

func getCertificateFlowRetryPolicy() *oci_common.RetryPolicy {
	policy := oci_common.DefaultRetryPolicyWithoutEventualConsistency()
	policy.MaximumNumberAttempts = 8
	return &policy
}

func getCertificateFlowBdsClient() (*oci_bds.BdsClient, error) {
	configFilePath := os.Getenv("OCI_CONFIG_FILE")
	if configFilePath == "" {
		configFilePath = filepath.Join(os.Getenv("HOME"), ".oci", "config")
	}

	profile := os.Getenv("TF_VAR_config_file_profile")
	if profile == "" {
		profile = os.Getenv("config_file_profile")
	}
	if profile == "" {
		profile = os.Getenv("OCI_CLI_PROFILE")
	}
	if profile == "" {
		profile = "DEFAULT"
	}

	provider := oci_common.CustomProfileConfigProvider(configFilePath, profile)

	bdsClient, err := oci_bds.NewBdsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	region := os.Getenv("TF_VAR_region")
	if region == "" {
		region = os.Getenv("region")
	}
	if region != "" {
		bdsClient.SetRegion(region)
	}

	if hostOverride := getCertificateFlowBdsClientHostOverride(); hostOverride != "" {
		bdsClient.Host = hostOverride
	}

	return &bdsClient, nil
}

func getCertificateFlowBdsClientHostOverride() string {
	overrides := os.Getenv("CLIENT_HOST_OVERRIDES")
	if overrides == "" {
		return ""
	}

	for _, override := range strings.Split(overrides, ",") {
		parts := strings.SplitN(strings.TrimSpace(override), "=", 2)
		if len(parts) != 2 {
			continue
		}

		if strings.TrimSpace(parts[0]) == "oci_bds.BdsClient" {
			return strings.TrimSpace(parts[1])
		}
	}

	return ""
}

func setDefaultBdsCertificateConfigurationForTest(bdsInstanceId, certConfigId, clusterAdminPassword, secretId string) error {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return err
	}

	details := oci_bds.SetDefaultBdsCertificateConfigurationDetails{}
	if secretId != "" {
		details.SecretId = &secretId
	} else {
		details.ClusterAdminPassword = &clusterAdminPassword
	}

	request := oci_bds.SetDefaultBdsCertificateConfigurationRequest{
		BdsInstanceId:                                &bdsInstanceId,
		BdsCertificateConfigurationId:                &certConfigId,
		SetDefaultBdsCertificateConfigurationDetails: details,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: getCertificateFlowRetryPolicy(),
		},
	}

	response, err := bdsClient.SetDefaultBdsCertificateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	return waitForBdsWorkRequestSucceeded(response.OpcWorkRequestId, 45*time.Minute)
}

func waitForBdsWorkRequestSucceeded(workRequestId *string, timeout time.Duration) error {
	bdsClient, err := getCertificateFlowBdsClient()
	if err != nil {
		return err
	}

	deadline := time.Now().Add(timeout)
	var lastErr error

	for time.Now().Before(deadline) {
		response, err := bdsClient.GetWorkRequest(context.Background(), oci_bds.GetWorkRequestRequest{
			WorkRequestId: workRequestId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: getCertificateFlowRetryPolicy(),
			},
		})
		if err != nil {
			lastErr = err
			time.Sleep(20 * time.Second)
			continue
		}

		switch response.Status {
		case oci_bds.OperationStatusSucceeded:
			return nil
		case oci_bds.OperationStatusFailed, oci_bds.OperationStatusCanceled:
			return getBdsWorkRequestErrorsForTest(bdsClient, workRequestId)
		}

		time.Sleep(20 * time.Second)
	}

	if lastErr != nil {
		return lastErr
	}
	return fmt.Errorf("timed out waiting for work request %s to succeed", *workRequestId)
}

func getBdsWorkRequestErrorsForTest(client *oci_bds.BdsClient, workRequestId *string) error {
	response, err := client.ListWorkRequestErrors(context.Background(), oci_bds.ListWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: getCertificateFlowRetryPolicy(),
		},
	})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, item := range response.Items {
		if item.Message != nil {
			allErrs = append(allErrs, *item.Message)
		}
	}

	if len(allErrs) == 0 {
		return fmt.Errorf("work request %s failed", *workRequestId)
	}
	return fmt.Errorf("work request %s failed: %s", *workRequestId, strings.Join(allErrs, "; "))
}
