// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/data_safe"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/bds"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/data_connectivity"
	tf_datascience "github.com/terraform-providers/terraform-provider-oci/internal/service/datascience"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/devops"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_log_analytics "github.com/terraform-providers/terraform-provider-oci/internal/service/log_analytics"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/apm_config"

	tf_logging "github.com/terraform-providers/terraform-provider-oci/internal/service/logging"

	tf_datacatalog "github.com/terraform-providers/terraform-provider-oci/internal/service/datacatalog"

	tf_apm_synthetics "github.com/terraform-providers/terraform-provider-oci/internal/service/apm_synthetics"

	tf_nosql "github.com/terraform-providers/terraform-provider-oci/internal/service/nosql"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"

	tf_blockchain "github.com/terraform-providers/terraform-provider-oci/internal/service/blockchain"
	tf_database "github.com/terraform-providers/terraform-provider-oci/internal/service/database"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
	network_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/network_load_balancer"
	tf_objectstorage "github.com/terraform-providers/terraform-provider-oci/internal/service/objectstorage"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/osp_gateway"
	tf_usage_proxy "github.com/terraform-providers/terraform-provider-oci/internal/service/usage_proxy"
)

func init() {
	exportApmConfigConfigHints.getIdFn = getApmConfigConfigId
	exportApmSyntheticsScriptHints.getIdFn = getApmSyntheticsScriptId
	exportApmSyntheticsMonitorHints.getIdFn = getApmSyntheticsMonitorId
	exportApmSyntheticsDedicatedVantagePointHints.getIdFn = getApmSyntheticsDedicatedVantagePointId
	exportBlockchainPeerHints.getIdFn = getBlockchainPeerId
	exportBlockchainOsnHints.getIdFn = getBlockchainOsnId
	exportBudgetAlertRuleHints.getIdFn = getBudgetAlertRuleId
	exportCoreInstancePoolInstanceHints.getIdFn = getCoreInstancePoolInstanceId
	exportCoreNetworkSecurityGroupSecurityRuleHints.getIdFn = getCoreNetworkSecurityGroupSecurityRuleId
	exportCoreDrgRouteTableRouteRuleHints.getIdFn = getCoreDrgRouteTableRouteRuleId
	exportDataConnectivityRegistryConnectionHints.getIdFn = getDataConnectivityRegistryConnectionId
	exportDataConnectivityRegistryDataAssetHints.getIdFn = getDataConnectivityRegistryDataAssetId
	exportDataConnectivityRegistryFolderHints.getIdFn = getDataConnectivityRegistryFolderId
	exportDataSafeMaskingPoliciesMaskingColumnHints.getIdFn = getDataSafeMaskingPoliciesMaskingColumnId
	exportDataSafeSensitiveDataModelsSensitiveColumnHints.getIdFn = getDataSafeSensitiveDataModelsSensitiveColumnId
	exportDataSafeDiscoveryJobsResultHints.getIdFn = getDataSafeDiscoveryJobsResultId
	exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints.getIdFn = getDatabaseAutonomousContainerDatabaseDataguardAssociationId
	exportDatabaseVmClusterNetworkHints.getIdFn = getDatabaseVmClusterNetworkId
	exportDatacatalogDataAssetHints.getIdFn = getDatacatalogDataAssetId
	exportDatacatalogConnectionHints.getIdFn = getDatacatalogConnectionId
	exportDatascienceModelProvenanceHints.getIdFn = getDatascienceModelProvenanceId
	exportDevopsRepositoryRefHints.getIdFn = getDevopsRepositoryRefId
	exportDnsRrsetHints.getIdFn = getDnsRrsetId
	exportIdentityApiKeyHints.getIdFn = getIdentityApiKeyId
	exportIdentityAuthenticationPolicyHints.getIdFn = getIdentityAuthenticationPolicyId
	exportIdentityAuthTokenHints.getIdFn = getIdentityAuthTokenId
	exportIdentityCustomerSecretKeyHints.getIdFn = getIdentityCustomerSecretKeyId
	exportIdentityIdpGroupMappingHints.getIdFn = getIdentityIdpGroupMappingId
	exportIdentitySmtpCredentialHints.getIdFn = getIdentitySmtpCredentialId
	exportIdentitySwiftPasswordHints.getIdFn = getIdentitySwiftPasswordId
	exportIdentityDbCredentialHints.getIdFn = getIdentityDbCredentialId
	exportKmsKeyHints.getIdFn = getKmsKeyId
	exportKmsKeyVersionHints.getIdFn = getKmsKeyVersionId
	exportLoadBalancerBackendHints.getIdFn = getLoadBalancerBackendId
	exportLoadBalancerBackendSetHints.getIdFn = getLoadBalancerBackendSetId
	exportLoadBalancerCertificateHints.getIdFn = getLoadBalancerCertificateId
	exportLoadBalancerHostnameHints.getIdFn = getLoadBalancerHostnameId
	exportLoadBalancerListenerHints.getIdFn = getLoadBalancerListenerId
	exportLoadBalancerPathRouteSetHints.getIdFn = getLoadBalancerPathRouteSetId
	exportLoadBalancerLoadBalancerRoutingPolicyHints.getIdFn = getLoadBalancerLoadBalancerRoutingPolicyId
	exportLoadBalancerRuleSetHints.getIdFn = getLoadBalancerRuleSetId
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.getIdFn = getLogAnalyticsLogAnalyticsObjectCollectionRuleId
	exportLogAnalyticsNamespaceScheduledTaskHints.getIdFn = getLogAnalyticsNamespaceScheduledTaskId
	exportLoggingLogHints.getIdFn = getLoggingLogId
	exportNetworkLoadBalancerBackendSetHints.getIdFn = getNetworkLoadBalancerBackendSetId
	exportNetworkLoadBalancerBackendHints.getIdFn = getNetworkLoadBalancerBackendId
	exportNetworkLoadBalancerListenerHints.getIdFn = getNetworkLoadBalancerListenerId
	exportNosqlIndexHints.getIdFn = getNosqlIndexId
	exportObjectStorageBucketHints.getIdFn = getObjectStorageBucketId
	exportObjectStorageObjectLifecyclePolicyHints.getIdFn = getObjectStorageObjectLifecyclePolicyId
	exportObjectStorageObjectHints.getIdFn = getObjectStorageObjectId
	exportObjectStoragePreauthenticatedRequestHints.getIdFn = getObjectStoragePreauthenticatedRequestId
	exportObjectStorageReplicationPolicyHints.getIdFn = getObjectStorageReplicationPolicyId
	exportOspGatewaySubscriptionHints.getIdFn = getOspGatewaySubscriptionId
	exportUsageProxySubscriptionRedeemableUserHints.getIdFn = getUsageProxySubscriptionRedeemableUserId
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmConfigConfigId(resource *OCIResource) (string, error) {

	configId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find configId for ApmConfig Config")
	}
	return apm_config.GetConfigCompositeId(configId), nil
}

func getApmSyntheticsScriptId(resource *OCIResource) (string, error) {

	scriptId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scriptId for ApmSynthetics Script")
	}
	return apm_synthetics.GetScriptCompositeId(scriptId), nil
}

func getApmSyntheticsMonitorId(resource *OCIResource) (string, error) {

	monitorId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find monitorId for ApmSynthetics Monitor")
	}
	return apm_synthetics.GetMonitorCompositeId(monitorId), nil
}

func getApmSyntheticsDedicatedVantagePointId(resource *OCIResource) (string, error) {

	dedicatedVantagePointId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dedicatedVantagePointId for ApmSynthetics DedicatedVantagePoint")
	}
	return apm_synthetics.GetDedicatedVantagePointCompositeId(dedicatedVantagePointId), nil
}

func getArtifactsContainerRepositoryId(resource *OCIResource) (string, error) {

	repositoryId, ok := resource.sourceAttributes["repository_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find repositoryId for Artifacts ContainerRepository")
	}
	return artifacts.GetContainerRepositoryCompositeId(repositoryId), nil
}

func getArtifactsContainerImageSignatureId(resource *OCIResource) (string, error) {

	imageSignatureId, ok := resource.sourceAttributes["image_signature_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find imageSignatureId for Artifacts ContainerImageSignature")
	}
	return artifacts.GetContainerImageSignatureCompositeId(imageSignatureId), nil
}

func getBdsAutoScalingConfigurationId(resource *OCIResource) (string, error) {

	autoScalingConfigurationId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find autoScalingConfigurationId for Bds AutoScalingConfiguration")
	}
	bdsInstanceId := resource.parent.id
	return bds.GetAutoScalingConfigurationCompositeId(autoScalingConfigurationId, bdsInstanceId), nil
}

func getBlockchainPeerId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	peerId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerId for Blockchain Peer")
	}
	return tf_blockchain.GetPeerCompositeId(blockchainPlatformId, peerId), nil
}

func getBlockchainOsnId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	osnId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find osnId for Blockchain Osn")
	}
	return tf_blockchain.GetOsnCompositeId(blockchainPlatformId, osnId), nil
}

func getBudgetAlertRuleId(resource *OCIResource) (string, error) {

	alertRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find alertRuleId for Budget AlertRule")
	}
	budgetId := resource.parent.id
	return budget.GetAlertRuleCompositeId(alertRuleId, budgetId), nil
}

func getCoreInstancePoolInstanceId(resource *OCIResource) (string, error) {

	instancePoolId := resource.parent.id
	return core.GetInstancePoolInstanceCompositeId(instancePoolId), nil
}

func getCoreNetworkSecurityGroupSecurityRuleId(resource *OCIResource) (string, error) {

	networkSecurityGroupId, ok := resource.sourceAttributes["network_security_group_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find networkSecurityGroupId for Core NetworkSecurityGroupSecurityRule")
	}
	return core.GetNetworkSecurityGroupSecurityRuleCompositeId(networkSecurityGroupId), nil
}

func getCoreDrgRouteTableRouteRuleId(resource *OCIResource) (string, error) {

	drgRouteTableId, ok := resource.sourceAttributes["drg_route_table_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find drgRouteTableId for Core DrgRouteTableRouteRule")
	}
	return core.GetDrgRouteTableRouteRuleCompositeId(drgRouteTableId), nil
}

func getDataConnectivityRegistryConnectionId(resource *OCIResource) (string, error) {

	connectionKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find connectionKey for DataConnectivity RegistryConnection")
	}
	registryId := resource.parent.sourceAttributes["registry_id"].(string)
	return data_connectivity.GetRegistryConnectionCompositeId(connectionKey, registryId), nil
}

func getDataConnectivityRegistryDataAssetId(resource *OCIResource) (string, error) {

	dataAssetKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for DataConnectivity RegistryDataAsset")
	}
	registryId := resource.parent.id
	return data_connectivity.GetRegistryDataAssetCompositeId(dataAssetKey, registryId), nil
}

func getDataConnectivityRegistryFolderId(resource *OCIResource) (string, error) {

	folderKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find folderKey for DataConnectivity RegistryFolder")
	}
	registryId := resource.parent.id
	return data_connectivity.GetRegistryFolderCompositeId(folderKey, registryId), nil
}

func getDataSafeMaskingPoliciesMaskingColumnId(resource *OCIResource) (string, error) {

	maskingColumnKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find maskingColumnKey for DataSafe MaskingPoliciesMaskingColumn")
	}
	maskingPolicyId := resource.parent.id
	return data_safe.GetMaskingPoliciesMaskingColumnCompositeId(maskingColumnKey, maskingPolicyId), nil
}

func getDataSafeSensitiveDataModelsSensitiveColumnId(resource *OCIResource) (string, error) {

	sensitiveColumnKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find sensitiveColumnKey for DataSafe SensitiveDataModelsSensitiveColumn")
	}
	sensitiveDataModelId := resource.parent.id
	return data_safe.GetSensitiveDataModelsSensitiveColumnCompositeId(sensitiveColumnKey, sensitiveDataModelId), nil
}

func getDataSafeDiscoveryJobsResultId(resource *OCIResource) (string, error) {

	discoveryJobId := resource.parent.id
	resultKey, ok := resource.sourceAttributes["result_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find resultKey for DataSafe DiscoveryJobsResult")
	}
	return data_safe.GetDiscoveryJobsResultCompositeId(discoveryJobId, resultKey), nil
}

func getDatabaseAutonomousContainerDatabaseDataguardAssociationId(resource *OCIResource) (string, error) {

	autonomousContainerDatabaseDataguardAssociationId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find autonomousContainerDatabaseDataguardAssociationId for Database AutonomousContainerDatabaseDataguardAssociation")
	}
	autonomousContainerDatabaseId := resource.parent.id
	return tf_database.GetAutonomousContainerDatabaseDataguardAssociationCompositeId(autonomousContainerDatabaseDataguardAssociationId, autonomousContainerDatabaseId), nil
}

func getDatabaseVmClusterNetworkId(resource *OCIResource) (string, error) {

	exadataInfrastructureId := resource.parent.id
	vmClusterNetworkId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find vmClusterNetworkId for Database VmClusterNetwork")
	}
	return database.GetVmClusterNetworkCompositeId(exadataInfrastructureId, vmClusterNetworkId), nil
}

func getDatacatalogDataAssetId(resource *OCIResource) (string, error) {

	catalogId := resource.parent.id
	dataAssetKey, ok := resource.sourceAttributes["data_asset_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog DataAsset")
	}
	return datacatalog.GetDataAssetCompositeId(catalogId, dataAssetKey), nil
}

func getDatacatalogConnectionId(resource *OCIResource) (string, error) {

	catalogId, ok := resource.sourceAttributes["catalog_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find catalogId for Datacatalog Connection")
	}
	connectionKey, ok := resource.sourceAttributes["connection_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find connectionKey for Datacatalog Connection")
	}
	dataAssetKey, ok := resource.sourceAttributes["data_asset_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog Connection")
	}
	return datacatalog.GetConnectionCompositeId(catalogId, connectionKey, dataAssetKey), nil
}

func getDatascienceModelProvenanceId(resource *OCIResource) (string, error) {

	modelId := resource.parent.id
	return datascience.GetModelProvenanceCompositeId(modelId), nil
}

func getDevopsRepositoryRefId(resource *OCIResource) (string, error) {

	refName, ok := resource.sourceAttributes["ref_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find refName for Devops RepositoryRef")
	}
	repositoryId := resource.parent.id
	return devops.GetRepositoryRefCompositeId(refName, repositoryId), nil
}

func getDnsRrsetId(resource *OCIResource) (string, error) {

	domain, ok := resource.sourceAttributes["domain"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find domain for Dns Rrset")
	}
	rtype, ok := resource.sourceAttributes["rtype"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find rtype for Dns Rrset")
	}
	zoneNameOrId, ok := resource.sourceAttributes["zone_name_or_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find zoneNameOrId for Dns Rrset")
	}
	return dns.GetRrsetCompositeId(domain, rtype, zoneNameOrId), nil
}

func getIdentityApiKeyId(resource *OCIResource) (string, error) {

	fingerprint, ok := resource.sourceAttributes["fingerprint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find fingerprint for Identity ApiKey")
	}
	userId := resource.parent.id
	return identity.GetApiKeyCompositeId(fingerprint, userId), nil
}

func getIdentityAuthenticationPolicyId(resource *OCIResource) (string, error) {

	compartmentId, ok := resource.sourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Identity AuthenticationPolicy")
	}
	return identity.GetAuthenticationPolicyCompositeId(compartmentId), nil
}

func getIdentityAuthTokenId(resource *OCIResource) (string, error) {

	authTokenId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find authTokenId for Identity AuthToken")
	}
	userId := resource.parent.id
	return identity.GetAuthTokenCompositeId(authTokenId, userId), nil
}

func getIdentityCustomerSecretKeyId(resource *OCIResource) (string, error) {

	customerSecretKeyId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find customerSecretKeyId for Identity CustomerSecretKey")
	}
	userId := resource.parent.id
	return identity.GetCustomerSecretKeyCompositeId(customerSecretKeyId, userId), nil
}

func getIdentityIdpGroupMappingId(resource *OCIResource) (string, error) {

	identityProviderId := resource.parent.id
	mappingId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find mappingId for Identity IdpGroupMapping")
	}
	return identity.GetIdpGroupMappingCompositeId(identityProviderId, mappingId), nil
}

func getIdentitySmtpCredentialId(resource *OCIResource) (string, error) {

	smtpCredentialId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find smtpCredentialId for Identity SmtpCredential")
	}
	userId := resource.parent.id
	return identity.GetSmtpCredentialCompositeId(smtpCredentialId, userId), nil
}

func getIdentitySwiftPasswordId(resource *OCIResource) (string, error) {

	swiftPasswordId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find swiftPasswordId for Identity SwiftPassword")
	}
	userId := resource.parent.id
	return identity.GetSwiftPasswordCompositeId(swiftPasswordId, userId), nil
}

func getIdentityDbCredentialId(resource *OCIResource) (string, error) {

	tagName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find tagName for Identity Tag")
	}
	userId := resource.parent.id
	return tf_identity.GetDbCredentialCompositeId(dbCredentialId, userId), nil
}

func getKmsKeyId(resource *OCIResource) (string, error) {

	keyId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find keyId for Kms Key")
	}
	return kms.GetKeyCompositeId(keyId), nil
}

func getKmsKeyVersionId(resource *OCIResource) (string, error) {

	keyId := resource.parent.id
	keyVersionId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find keyVersionId for Kms KeyVersion")
	}
	return kms.GetKeyVersionCompositeId(keyId, keyVersionId), nil
}

func getLoadBalancerBackendId(resource *OCIResource) (string, error) {

	backendName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendName for LoadBalancer Backend")
	}
	backendsetName, ok := resource.sourceAttributes["backendset_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendsetName for LoadBalancer Backend")
	}
	loadBalancerId, ok := resource.sourceAttributes["load_balancer_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find loadBalancerId for LoadBalancer Backend")
	}
	return load_balancer.GetBackendCompositeId(backendName, backendsetName, loadBalancerId), nil
}

func getLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for LoadBalancer BackendSet")
	}
	loadBalancerId := resource.parent.id
	return load_balancer.GetBackendSetCompositeId(backendSetName, loadBalancerId), nil
}

func getLoadBalancerCertificateId(resource *OCIResource) (string, error) {

	certificateName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find certificateName for LoadBalancer Certificate")
	}
	loadBalancerId := resource.parent.id
	return load_balancer.GetCertificateCompositeId(certificateName, loadBalancerId), nil
}

func getLoadBalancerHostnameId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer Hostname")
	}
	return load_balancer.GetHostnameCompositeId(loadBalancerId, name), nil
}

func getLoadBalancerListenerId(resource *OCIResource) (string, error) {

	listenerName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for LoadBalancer Listener")
	}
	loadBalancerId, ok := resource.sourceAttributes["load_balancer_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find loadBalancerId for LoadBalancer Listener")
	}
	return load_balancer.GetListenerCompositeId(listenerName, loadBalancerId), nil
}

func getLoadBalancerPathRouteSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	pathRouteSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find pathRouteSetName for LoadBalancer PathRouteSet")
	}
	return load_balancer.GetPathRouteSetCompositeId(loadBalancerId, pathRouteSetName), nil
}

func getLoadBalancerLoadBalancerRoutingPolicyId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	routingPolicyName, ok := resource.sourceAttributes["routing_policy_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find routingPolicyName for LoadBalancer LoadBalancerRoutingPolicy")
	}
	return load_balancer.GetLoadBalancerRoutingPolicyCompositeId(loadBalancerId, routingPolicyName), nil
}

func getLoadBalancerRuleSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer RuleSet")
	}
	return load_balancer.GetRuleSetCompositeId(loadBalancerId, name), nil
}

func getLogAnalyticsLogAnalyticsObjectCollectionRuleId(resource *OCIResource) (string, error) {

	logAnalyticsObjectCollectionRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find logAnalyticsObjectCollectionRuleId for LogAnalytics LogAnalyticsObjectCollectionRule")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsObjectCollectionRule")
	}
	return log_analytics.GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId, namespace), nil
}

func getLogAnalyticsNamespaceScheduledTaskId(resource *OCIResource) (string, error) {

	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceScheduledTask")
	}
	scheduledTaskId, ok := resource.sourceAttributes["scheduled_task_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scheduledTaskId for LogAnalytics NamespaceScheduledTask")
	}
	return log_analytics.GetNamespaceScheduledTaskCompositeId(namespace, scheduledTaskId), nil
}

func getLoggingLogId(resource *OCIResource) (string, error) {

	logGroupId := resource.parent.id
	logId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find logId for Logging Log")
	}
	return logging.GetLogCompositeId(logGroupId, logId), nil
}

func getMysqlMysqlBackupId(resource *OCIResource) (string, error) {

	backupId, ok := resource.sourceAttributes["backup_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backupId for Mysql MysqlBackup")
	}
	return mysql.GetMysqlBackupCompositeId(backupId), nil
}

func getMysqlMysqlDbSystemId(resource *OCIResource) (string, error) {

	dbSystemId, ok := resource.sourceAttributes["db_system_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dbSystemId for Mysql MysqlDbSystem")
	}
	return mysql.GetMysqlDbSystemCompositeId(dbSystemId), nil
}

func getNetworkLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer BackendSet")
	}
	networkLoadBalancerId := resource.parent.id
	return network_load_balancer.GetBackendSetCompositeId(backendSetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerBackendId(resource *OCIResource) (string, error) {

	backendName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendName for NetworkLoadBalancer Backend")
	}
	backendSetName := resource.parent.id
	networkLoadBalancerId := resource.parent.id
	return network_load_balancer.GetBackendCompositeId(backendName, backendSetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerListenerId(resource *OCIResource) (string, error) {

	listenerName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for NetworkLoadBalancer Listener")
	}
	networkLoadBalancerId := resource.parent.id
	return network_load_balancer.GetListenerCompositeId(listenerName, networkLoadBalancerId), nil
}

func getNosqlIndexId(resource *OCIResource) (string, error) {

	indexName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find indexName for Nosql Index")
	}
	tableNameOrId, ok := resource.sourceAttributes["table_name_or_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find tableNameOrId for Nosql Index")
	}
	return nosql.GetIndexCompositeId(indexName, tableNameOrId), nil
}

func getObjectStorageBucketId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Bucket")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Bucket")
	}
	return object_storage.GetBucketCompositeId(bucket, namespace), nil
}

func getObjectStorageObjectLifecyclePolicyId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ObjectLifecyclePolicy")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage ObjectLifecyclePolicy")
	}
	return object_storage.GetObjectLifecyclePolicyCompositeId(bucket, namespace), nil
}

func getObjectStorageObjectId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Object")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Object")
	}
	object, ok := resource.sourceAttributes["object"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find object for ObjectStorage Object")
	}
	return object_storage.GetObjectCompositeId(bucket, namespace, object), nil
}

func getObjectStoragePreauthenticatedRequestId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage PreauthenticatedRequest")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage PreauthenticatedRequest")
	}
	parId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find parId for ObjectStorage PreauthenticatedRequest")
	}
	return object_storage.GetPreauthenticatedRequestCompositeId(bucket, namespace, parId), nil
}

func getObjectStorageReplicationPolicyId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ReplicationPolicy")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage ReplicationPolicy")
	}
	replicationId, ok := resource.sourceAttributes["replication_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find replicationId for ObjectStorage ReplicationPolicy")
	}
	return object_storage.GetReplicationPolicyCompositeId(bucket, namespace, replicationId), nil
}

func getOspGatewaySubscriptionId(resource *OCIResource) (string, error) {

	subscriptionId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find subscriptionId for OspGateway Subscription")
	}
	return osp_gateway.GetSubscriptionCompositeId(subscriptionId), nil
}

func getUsageProxySubscriptionRedeemableUserId(resource *OCIResource) (string, error) {

	subscriptionId := resource.parent.id
	return usage_proxy.GetSubscriptionRedeemableUserCompositeId(subscriptionId), nil
}
