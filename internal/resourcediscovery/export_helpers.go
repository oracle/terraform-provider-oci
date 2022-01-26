// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/bds"

	tf_datascience "github.com/terraform-providers/terraform-provider-oci/internal/service/datascience"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/devops"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_log_analytics "github.com/terraform-providers/terraform-provider-oci/internal/service/log_analytics"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/apm_config"

	tf_logging "github.com/terraform-providers/terraform-provider-oci/internal/service/logging"

	tf_datacatalog "github.com/terraform-providers/terraform-provider-oci/internal/service/datacatalog"

	tf_apm_synthetics "github.com/terraform-providers/terraform-provider-oci/internal/service/apm_synthetics"

	tf_nosql "github.com/terraform-providers/terraform-provider-oci/internal/service/nosql"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"

	tf_blockchain "github.com/terraform-providers/terraform-provider-oci/internal/service/blockchain"
	tf_database "github.com/terraform-providers/terraform-provider-oci/internal/service/database"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
	network_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/network_load_balancer"
	tf_objectstorage "github.com/terraform-providers/terraform-provider-oci/internal/service/objectstorage"
	tf_usage_proxy "github.com/terraform-providers/terraform-provider-oci/internal/service/usage_proxy"
)

func init() {
	exportApmConfigConfigHints.getIdFn = getApmConfigConfigId
	exportArtifactsContainerRepositoryHints.getIdFn = getArtifactsContainerRepositoryId
	exportArtifactsContainerImageSignatureHints.getIdFn = getArtifactsContainerImageSignatureId
	exportBdsBdsInstanceApiKeyHints.getIdFn = getBdsBdsInstanceApiKeyId
	exportArtifactsRepositoryHints.getIdFn = getArtifactsRepositoryId
	exportApmSyntheticsScriptHints.getIdFn = getApmSyntheticsScriptId
	exportApmSyntheticsMonitorHints.getIdFn = getApmSyntheticsMonitorId
	exportBlockchainPeerHints.getIdFn = getBlockchainPeerId
	exportBlockchainOsnHints.getIdFn = getBlockchainOsnId
	exportBudgetAlertRuleHints.getIdFn = getBudgetAlertRuleId
	exportCoreInstancePoolInstanceHints.getIdFn = getCoreInstancePoolInstanceId
	exportCoreNetworkSecurityGroupSecurityRuleHints.getIdFn = getCoreNetworkSecurityGroupSecurityRuleId
	exportCoreDrgRouteTableRouteRuleHints.getIdFn = getCoreDrgRouteTableRouteRuleId
	exportDatabaseAutonomousContainerDatabaseDataguardAssociationHints.getIdFn = getDatabaseAutonomousContainerDatabaseDataguardAssociationId
	exportDatabaseVmClusterNetworkHints.getIdFn = getDatabaseVmClusterNetworkId
	exportDatacatalogDataAssetHints.getIdFn = getDatacatalogDataAssetId
	exportDatacatalogConnectionHints.getIdFn = getDatacatalogConnectionId
	exportDatascienceModelProvenanceHints.getIdFn = getDatascienceModelProvenanceId
	exportDevopsRepositoryRefHints.getIdFn = getDevopsRepositoryRefId
	exportDnsRrsetHints.getIdFn = getDnsRrsetId
	exportIdentityApiKeyHints.getIdFn = getIdentityApiKeyId
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
	exportUsageProxySubscriptionRedeemableUserHints.getIdFn = getUsageProxySubscriptionRedeemableUserId
	exportOnsNotificationTopicHints.getIdFn = getOnsNotificationTopicId
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmConfigConfigId(resource *OCIResource) (string, error) {

	configId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find configId for ApmConfig Config")
	}
	apmDomainId, ok := resource.sourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmConfig Config")
	}
	return apm_config.GetConfigCompositeId(configId, apmDomainId), nil
}

func getApmSyntheticsScriptId(resource *OCIResource) (string, error) {

	scriptId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scriptId for ApmSynthetics Script")
	}
	apmDomainId, ok := resource.sourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmSynthetics Script")
	}
	return tf_apm_synthetics.GetScriptCompositeId(scriptId, apmDomainId), nil
}

func getApmSyntheticsMonitorId(resource *OCIResource) (string, error) {

	monitorId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find monitorId for ApmSynthetics Monitor")
	}
	apmDomainId, ok := resource.sourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmSynthetics Monitor")
	}
	return tf_apm_synthetics.GetMonitorCompositeId(monitorId, apmDomainId), nil
}

func getArtifactsContainerRepositoryId(resource *OCIResource) (string, error) {

	repositoryId, ok := resource.sourceAttributes["repository_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find repositoryId for Artifacts ContainerRepository")
	}
	return repositoryId, nil

}

func getArtifactsContainerImageSignatureId(resource *OCIResource) (string, error) {

	imageSignatureId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find imageSignatureId for Artifacts ContainerImageSignature")
	}
	return imageSignatureId, nil
}

func getArtifactsRepositoryId(resource *OCIResource) (string, error) {
	repositoryId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find repositoryId for Artifacts Respository")
	}
	return repositoryId, nil
}

func getBdsBdsInstanceApiKeyId(resource *OCIResource) (string, error) {

	apiKeyId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apiKeyId for Bds BdsInstanceApiKey")
	}
	bdsInstanceId := resource.parent.id
	return bds.GetBdsInstanceApiKeyCompositeId(apiKeyId, bdsInstanceId), nil
}

func getBlockchainPeerId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	peerId, ok := resource.sourceAttributes["peer_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerId for Blockchain Peer")
	}
	return tf_blockchain.GetPeerCompositeId(blockchainPlatformId, peerId), nil
}

func getBlockchainOsnId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	osnId, ok := resource.sourceAttributes["osn_key"].(string)
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
	instanceId := resource.sourceAttributes["instance_id"].(string)
	return tf_core.GetInstancePoolInstanceCompositeId(instancePoolId, instanceId), nil
}

func getCoreNetworkSecurityGroupSecurityRuleId(resource *OCIResource) (string, error) {

	networkSecurityGroupId := resource.parent.id
	securityRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for Core NetworkSecurityGroupSecurityRule")
	}
	return tf_core.GetNetworkSecurityGroupSecurityRuleCompositeId(networkSecurityGroupId, securityRuleId), nil
}

func getCoreDrgRouteTableRouteRuleId(resource *OCIResource) (string, error) {

	drgRouteTableId := resource.parent.id
	drgRouteRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find drgRouteTableId for Core DrgRouteTableRouteRule")
	}
	return tf_core.GetDrgRouteTableRouteRuleCompositeId(drgRouteTableId, drgRouteRuleId), nil
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
	return tf_database.GetVmClusterNetworkCompositeId(exadataInfrastructureId, vmClusterNetworkId), nil
}

func getDatacatalogDataAssetId(resource *OCIResource) (string, error) {

	catalogId := resource.parent.id
	dataAssetKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog DataAsset")
	}
	return tf_datacatalog.GetDataAssetCompositeId(catalogId, dataAssetKey), nil
}

func getDatacatalogConnectionId(resource *OCIResource) (string, error) {

	catalogId, ok := resource.parent.sourceAttributes["catalog_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find catalogId for Datacatalog Connection")
	}
	connectionKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find connectionKey for Datacatalog Connection")
	}
	dataAssetKey, ok := resource.sourceAttributes["data_asset_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog Connection")
	}
	return tf_datacatalog.GetConnectionCompositeId(catalogId, connectionKey, dataAssetKey), nil
}

func getDatascienceModelProvenanceId(resource *OCIResource) (string, error) {

	modelId := resource.parent.id
	return tf_datascience.GetModelProvenanceCompositeId(modelId), nil
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
	zoneNameOrId := resource.parent.id
	return getRrsetCompositeId(domain, rtype, zoneNameOrId), nil
}

func getRrsetCompositeId(domain string, rtype string, zoneNameOrId string) string {
	domain = url.PathEscape(domain)
	rtype = url.PathEscape(rtype)
	zoneNameOrId = url.PathEscape(zoneNameOrId)
	compositeId := "zoneNameOrId/" + zoneNameOrId + "/domain/" + domain + "/rtype/" + rtype
	return compositeId
}

func getIdentityApiKeyId(resource *OCIResource) (string, error) {

	fingerprint, ok := resource.sourceAttributes["fingerprint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find fingerprint for Identity ApiKey")
	}
	userId := resource.parent.id
	return tf_identity.GetApiKeyCompositeId(fingerprint, userId), nil
}

func getIdentityAuthenticationPolicyId(resource *OCIResource) (string, error) {

	compartmentId, ok := resource.sourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Identity AuthenticationPolicy")
	}
	return tf_identity.GetAuthenticationPolicyCompositeId(compartmentId), nil
}

func getIdentityAuthTokenId(resource *OCIResource) (string, error) {

	authTokenId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find authTokenId for Identity AuthToken")
	}
	userId := resource.parent.id
	return tf_identity.GetAuthTokenCompositeId(authTokenId, userId), nil
}

func getIdentityCustomerSecretKeyId(resource *OCIResource) (string, error) {

	customerSecretKeyId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find customerSecretKeyId for Identity CustomerSecretKey")
	}
	userId := resource.parent.id
	return tf_identity.GetCustomerSecretKeyCompositeId(customerSecretKeyId, userId), nil
}

func getIdentityIdpGroupMappingId(resource *OCIResource) (string, error) {

	identityProviderId := resource.parent.id
	mappingId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find mappingId for Identity IdpGroupMapping")
	}
	return tf_identity.GetIdpGroupMappingCompositeId(identityProviderId, mappingId), nil
}

func getIdentitySmtpCredentialId(resource *OCIResource) (string, error) {

	smtpCredentialId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find smtpCredentialId for Identity SmtpCredential")
	}
	userId := resource.parent.id
	return tf_identity.GetSmtpCredentialCompositeId(smtpCredentialId, userId), nil
}

func getIdentitySwiftPasswordId(resource *OCIResource) (string, error) {

	swiftPasswordId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find swiftPasswordId for Identity SwiftPassword")
	}
	userId := resource.parent.id
	return tf_identity.GetSwiftPasswordCompositeId(swiftPasswordId, userId), nil
}

func getIdentityDbCredentialId(resource *OCIResource) (string, error) {

	dbCredentialId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dbCredentialId for Identity DbCredential")
	}
	userId := resource.parent.id
	return tf_identity.GetDbCredentialCompositeId(dbCredentialId, userId), nil
}

func getKmsKeyId(resource *OCIResource) (string, error) {
	managementEndpoint, ok := resource.parent.sourceAttributes["management_endpoint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find management_endpoint for Index id")
	}
	var keyId string
	// observed that Id is not always available in sourceAttributes - refer export_compartment.go->findResourcesGeneric() to visualize below docs
	// resource.sourceAttributes has the id in the cases where getKmsKeyId is called with LIST data source response, because list SetData() sets the Id, but this is only done temporarily to populate compositeID
	// When getKmsKeyId is called for resource, resource.sourceAttributes is not set yet,(so far we used LIST response to get composite Id) but we can get the real ocid after Read because Id was set in the method kms_key_resource.go->readKmsKey()
	switch resource.rawResource.(type) {
	case *schema.ResourceData:
		// 	rawResource from resource read response
		var resourceSchema *schema.ResourceData = resource.rawResource.(*schema.ResourceData)
		keyId = resourceSchema.Id()
	case map[string]interface{}:
		// 	rawResource from LIST data source read response
		var resourceMap map[string]interface{} = resource.rawResource.(map[string]interface{})
		keyId = resourceMap["id"].(string)
	}
	return tf_kms.GetCompositeKeyId(managementEndpoint, keyId), nil
}

func getKmsKeyVersionId(resource *OCIResource) (string, error) {

	managementEndpoint, ok := resource.parent.sourceAttributes["management_endpoint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find management_endpoint for Kms KeyVersion")
	}
	keyId := resource.parent.sourceAttributes["id"].(string)
	keyVersionId, ok := resource.sourceAttributes["key_version_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find keyVersionId for Kms KeyVersion")
	}
	return tf_kms.GetCompositeKeyVersionId(managementEndpoint, keyId, keyVersionId), nil
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
	return tf_load_balancer.GetBackendCompositeId(backendName, backendsetName, loadBalancerId), nil
}

func getLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for LoadBalancer BackendSet")
	}
	loadBalancerId := resource.parent.id
	return tf_load_balancer.GetBackendSetCompositeId(backendSetName, loadBalancerId), nil
}

func getLoadBalancerCertificateId(resource *OCIResource) (string, error) {

	certificateName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find certificateName for LoadBalancer Certificate")
	}
	loadBalancerId := resource.parent.id
	return tf_load_balancer.GetCertificateCompositeId(certificateName, loadBalancerId), nil
}

func getLoadBalancerHostnameId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer Hostname")
	}
	return tf_load_balancer.GetHostnameCompositeId(loadBalancerId, name), nil
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
	return tf_load_balancer.GetListenerCompositeId(listenerName, loadBalancerId), nil
}

func getLoadBalancerPathRouteSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	pathRouteSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find pathRouteSetName for LoadBalancer PathRouteSet")
	}
	return tf_load_balancer.GetPathRouteSetCompositeId(loadBalancerId, pathRouteSetName), nil
}

func getLoadBalancerLoadBalancerRoutingPolicyId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	routingPolicyName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find routingPolicyName for LoadBalancer LoadBalancerRoutingPolicy")
	}
	return tf_load_balancer.GetLoadBalancerRoutingPolicyCompositeId(loadBalancerId, routingPolicyName), nil
}

func getLoadBalancerRuleSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer RuleSet")
	}
	return tf_load_balancer.GetRuleSetCompositeId(loadBalancerId, name), nil
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
	return tf_log_analytics.GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId, namespace), nil
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
	return tf_log_analytics.GetNamespaceScheduledTaskCompositeId(namespace, scheduledTaskId), nil
}

func getLoggingLogId(resource *OCIResource) (string, error) {

	logGroupId := resource.parent.id
	logId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find logId for Logging Log")
	}
	return tf_logging.GetLogCompositeId(logGroupId, logId), nil
}

func getNetworkLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer BackendSet")
	}
	networkLoadBalancerId := resource.parent.id
	return network_load_balancer.GetNlbBackendSetCompositeId(backendSetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerBackendId(resource *OCIResource) (string, error) {
	backendName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendName for NetworkLoadBalancer Backend")
	}
	backendsetName, ok := resource.parent.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer Backend")
	}
	networkLoadBalancerId := resource.parent.parent.id
	return network_load_balancer.GetNlbBackendCompositeId(backendName, backendsetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerListenerId(resource *OCIResource) (string, error) {

	listenerName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for NetworkLoadBalancer Listener")
	}
	networkLoadBalancerId := resource.parent.parent.id
	return network_load_balancer.GetNlbListenerCompositeId(listenerName, networkLoadBalancerId), nil
}

func getNosqlIndexId(resource *OCIResource) (string, error) {

	indexName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find indexName for Nosql Index")
	}
	tableNameOrId := resource.parent.id

	return tf_nosql.GetIndexCompositeId(indexName, tableNameOrId), nil
}

func getObjectStorageBucketId(resource *OCIResource) (string, error) {

	bucket, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Bucket")
	}
	namespace, ok := resource.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Bucket")
	}
	return tf_objectstorage.GetBucketCompositeId(bucket, namespace), nil
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
	return tf_objectstorage.GetObjectLifecyclePolicyCompositeId(bucket, namespace), nil
}

func getObjectStorageObjectId(resource *OCIResource) (string, error) {

	bucket, ok := resource.parent.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Object")
	}
	namespace, ok := resource.parent.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Object")
	}
	object, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find object for ObjectStorage Object")
	}
	return tf_objectstorage.GetObjectCompositeId(bucket, namespace, object), nil
}

func getObjectStoragePreauthenticatedRequestId(resource *OCIResource) (string, error) {

	bucket, ok := resource.parent.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage PreauthenticatedRequest")
	}
	namespace, ok := resource.parent.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage PreauthenticatedRequest")
	}
	parId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find parId for ObjectStorage PreauthenticatedRequest")
	}
	return tf_objectstorage.GetPreauthenticatedRequestCompositeId(bucket, namespace, parId), nil
}

func getObjectStorageReplicationPolicyId(resource *OCIResource) (string, error) {

	bucket, ok := resource.parent.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ReplicationPolicy")
	}
	namespace, ok := resource.parent.sourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage ReplicationPolicy")
	}
	replicationId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find replicationId for ObjectStorage ReplicationPolicy")
	}
	return tf_objectstorage.GetReplicationPolicyCompositeId(bucket, namespace, replicationId), nil
}

func getUsageProxySubscriptionRedeemableUserId(resource *OCIResource) (string, error) {

	subscriptionId := resource.parent.id
	tenancyId, ok := resource.parent.sourceAttributes["tenancy_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ReplicationPolicy")
	}
	return tf_usage_proxy.GetSubscriptionRedeemableUserCompositeId(subscriptionId, tenancyId), nil
}

func getOnsNotificationTopicId(resource *OCIResource) (string, error) {
	id, ok := resource.sourceAttributes["topic_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find topic id for ons notification topic")
	}
	return id, nil
}
