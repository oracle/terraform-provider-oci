// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func init() {
	exportArtifactsContainerImageSignatureHints.getIdFn = getArtifactsContainerImageSignatureId
	exportBlockchainPeerHints.getIdFn = getBlockchainPeerId
	exportBlockchainOsnHints.getIdFn = getBlockchainOsnId
	exportBudgetAlertRuleHints.getIdFn = getBudgetAlertRuleId
	exportCoreInstancePoolInstanceHints.getIdFn = getCoreInstancePoolInstanceId
	exportCoreNetworkSecurityGroupSecurityRuleHints.getIdFn = getCoreNetworkSecurityGroupSecurityRuleId
	exportCoreDrgRouteTableRouteRuleHints.getIdFn = getCoreDrgRouteTableRouteRuleId
	exportCoreDrgRouteDistributionStatementHints.getIdFn = getCoreDrgRouteDistributionStatementId
	exportDatabaseVmClusterNetworkHints.getIdFn = getDatabaseVmClusterNetworkId
	exportDatacatalogDataAssetHints.getIdFn = getDatacatalogDataAssetId
	exportDatacatalogConnectionHints.getIdFn = getDatacatalogConnectionId
	exportDatascienceModelProvenanceHints.getIdFn = getDatascienceModelProvenanceId
	exportDnsRrsetHints.getIdFn = getDnsRrsetId
	exportIdentityApiKeyHints.getIdFn = getIdentityApiKeyId
	exportIdentityAuthTokenHints.getIdFn = getIdentityAuthTokenId
	exportIdentityCustomerSecretKeyHints.getIdFn = getIdentityCustomerSecretKeyId
	exportIdentityIdpGroupMappingHints.getIdFn = getIdentityIdpGroupMappingId
	exportIdentitySmtpCredentialHints.getIdFn = getIdentitySmtpCredentialId
	exportIdentitySwiftPasswordHints.getIdFn = getIdentitySwiftPasswordId
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
	exportNetworkLoadBalancerBackendSetHints.getIdFn = getNetworkLoadBalancerBackendSetId
	exportNetworkLoadBalancerBackendHints.getIdFn = getNetworkLoadBalancerBackendId
	exportNetworkLoadBalancerListenerHints.getIdFn = getNetworkLoadBalancerListenerId
	exportNosqlIndexHints.getIdFn = getNosqlIndexId
	exportObjectStorageBucketHints.getIdFn = getObjectStorageBucketId
	exportObjectStorageObjectLifecyclePolicyHints.getIdFn = getObjectStorageObjectLifecyclePolicyId
	exportObjectStorageObjectHints.getIdFn = getObjectStorageObjectId
	exportObjectStoragePreauthenticatedRequestHints.getIdFn = getObjectStoragePreauthenticatedRequestId
	exportObjectStorageReplicationPolicyHints.getIdFn = getObjectStorageReplicationPolicyId
	exportOnsNotificationTopicHints.getIdFn = getOnsNotificationTopicId

}

func getArtifactsContainerImageSignatureId(resource *OCIResource) (string, error) {

	imageSignatureId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find imageSignatureId for Artifacts ContainerImageSignature")
	}
	return imageSignatureId, nil
}

func getBlockchainPeerId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	peerId, ok := resource.sourceAttributes["peer_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerId for Blockchain Peer")
	}
	return getPeerCompositeId(blockchainPlatformId, peerId), nil
}

func getBlockchainOsnId(resource *OCIResource) (string, error) {

	blockchainPlatformId := resource.parent.id
	osnId, ok := resource.sourceAttributes["osn_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find osnId for Blockchain Osn")
	}
	return getOsnCompositeId(blockchainPlatformId, osnId), nil
}

func getBudgetAlertRuleId(resource *OCIResource) (string, error) {

	alertRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find alertRuleId for Budget AlertRule")
	}
	budgetId := resource.parent.id
	return getAlertRuleCompositeId(alertRuleId, budgetId), nil
}

func getCoreInstancePoolInstanceId(resource *OCIResource) (string, error) {

	instancePoolId := resource.parent.id
	instanceId := resource.sourceAttributes["instance_id"].(string)
	return getInstancePoolInstanceCompositeId(instancePoolId, instanceId), nil
}

func getCoreNetworkSecurityGroupSecurityRuleId(resource *OCIResource) (string, error) {

	networkSecurityGroupId := resource.parent.id
	securityRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for Core NetworkSecurityGroupSecurityRule")
	}
	return getNetworkSecurityGroupSecurityRuleCompositeId(networkSecurityGroupId, securityRuleId), nil
}

func getCoreDrgRouteTableRouteRuleId(resource *OCIResource) (string, error) {

	drgRouteTableId := resource.parent.id
	drgRouteRuleId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find drgRouteTableId for Core DrgRouteTableRouteRule")
	}
	return getDrgRouteTableRouteRuleCompositeId(drgRouteTableId, drgRouteRuleId), nil
}

func getCoreDrgRouteDistributionStatementId(resource *OCIResource) (string, error) {

	drgRouteDistributionId := resource.parent.id
	statementId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for Core DrgRouteDistributionStatementId")
	}
	return getDrgRouteDistributionStatementCompositeId(drgRouteDistributionId, statementId), nil
}

func getDatabaseVmClusterNetworkId(resource *OCIResource) (string, error) {

	exadataInfrastructureId := resource.parent.id
	vmClusterNetworkId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find vmClusterNetworkId for Database VmClusterNetwork")
	}
	return getVmClusterNetworkCompositeId(exadataInfrastructureId, vmClusterNetworkId), nil
}

func getDatacatalogDataAssetId(resource *OCIResource) (string, error) {

	catalogId := resource.parent.id
	dataAssetKey, ok := resource.sourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog DataAsset")
	}
	return getDataAssetCompositeId(catalogId, dataAssetKey), nil
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
	return getConnectionCompositeId(catalogId, connectionKey, dataAssetKey), nil
}

func getDatascienceModelProvenanceId(resource *OCIResource) (string, error) {

	modelId := resource.parent.id
	return getModelProvenanceCompositeId(modelId), nil
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

func getIdentityApiKeyId(resource *OCIResource) (string, error) {

	fingerprint, ok := resource.sourceAttributes["fingerprint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find fingerprint for Identity ApiKey")
	}
	userId := resource.parent.id
	return getApiKeyCompositeId(fingerprint, userId), nil
}

func getIdentityAuthenticationPolicyId(resource *OCIResource) (string, error) {

	compartmentId, ok := resource.sourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Identity AuthenticationPolicy")
	}
	return getAuthenticationPolicyCompositeId(compartmentId), nil
}

func getIdentityAuthTokenId(resource *OCIResource) (string, error) {

	authTokenId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find authTokenId for Identity AuthToken")
	}
	userId := resource.parent.id
	return getAuthTokenCompositeId(authTokenId, userId), nil
}

func getIdentityCustomerSecretKeyId(resource *OCIResource) (string, error) {

	customerSecretKeyId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find customerSecretKeyId for Identity CustomerSecretKey")
	}
	userId := resource.parent.id
	return getCustomerSecretKeyCompositeId(customerSecretKeyId, userId), nil
}

func getIdentityIdpGroupMappingId(resource *OCIResource) (string, error) {

	identityProviderId := resource.parent.id
	mappingId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find mappingId for Identity IdpGroupMapping")
	}
	return getIdpGroupMappingCompositeId(identityProviderId, mappingId), nil
}

func getIdentitySmtpCredentialId(resource *OCIResource) (string, error) {

	smtpCredentialId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find smtpCredentialId for Identity SmtpCredential")
	}
	userId := resource.parent.id
	return getSmtpCredentialCompositeId(smtpCredentialId, userId), nil
}

func getIdentitySwiftPasswordId(resource *OCIResource) (string, error) {

	swiftPasswordId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find swiftPasswordId for Identity SwiftPassword")
	}
	userId := resource.parent.id
	return getSwiftPasswordCompositeId(swiftPasswordId, userId), nil
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
	return getCompositeKeyId(managementEndpoint, keyId), nil
}

func getKmsKeyVersionId(resource *OCIResource) (string, error) {

	managementEndpoint, ok := resource.parent.sourceAttributes["management_endpoint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find management_endpoint for Kms KeyVersion")
	}
	keyId := resource.parent.id
	keyVersionId, ok := resource.sourceAttributes["key_version_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find keyVersionId for Kms KeyVersion")
	}
	return getCompositeKeyVersionId(managementEndpoint, keyId, keyVersionId), nil
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
	return getBackendCompositeId(backendName, backendsetName, loadBalancerId), nil
}

func getLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for LoadBalancer BackendSet")
	}
	loadBalancerId := resource.parent.id
	return getBackendSetCompositeId(backendSetName, loadBalancerId), nil
}

func getLoadBalancerCertificateId(resource *OCIResource) (string, error) {

	certificateName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find certificateName for LoadBalancer Certificate")
	}
	loadBalancerId := resource.parent.id
	return getCertificateCompositeId(certificateName, loadBalancerId), nil
}

func getLoadBalancerHostnameId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer Hostname")
	}
	return getHostnameCompositeId(loadBalancerId, name), nil
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
	return getListenerCompositeId(listenerName, loadBalancerId), nil
}

func getLoadBalancerPathRouteSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	pathRouteSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find pathRouteSetName for LoadBalancer PathRouteSet")
	}
	return getPathRouteSetCompositeId(loadBalancerId, pathRouteSetName), nil
}

func getLoadBalancerLoadBalancerRoutingPolicyId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	routingPolicyName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find routingPolicyName for LoadBalancer LoadBalancerRoutingPolicy")
	}
	return getLoadBalancerRoutingPolicyCompositeId(loadBalancerId, routingPolicyName), nil
}

func getLoadBalancerRuleSetId(resource *OCIResource) (string, error) {

	loadBalancerId := resource.parent.id
	name, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer RuleSet")
	}
	return getRuleSetCompositeId(loadBalancerId, name), nil
}

func getNetworkLoadBalancerBackendSetId(resource *OCIResource) (string, error) {

	backendSetName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer BackendSet")
	}
	networkLoadBalancerId := resource.parent.id
	return getNlbBackendSetCompositeId(backendSetName, networkLoadBalancerId), nil
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
	return getNlbBackendCompositeId(backendName, backendsetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerListenerId(resource *OCIResource) (string, error) {

	listenerName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for NetworkLoadBalancer Listener")
	}
	networkLoadBalancerId := resource.parent.parent.id
	return getNlbListenerCompositeId(listenerName, networkLoadBalancerId), nil
}

func getNosqlIndexId(resource *OCIResource) (string, error) {

	indexName, ok := resource.sourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find indexName for Nosql Index")
	}
	tableNameOrId := resource.parent.id

	return getIndexCompositeId(indexName, tableNameOrId), nil
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
	return getBucketCompositeId(bucket, namespace), nil
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
	return getObjectLifecyclePolicyCompositeId(bucket, namespace), nil
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
	return getObjectCompositeId(bucket, namespace, object), nil
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
	return getPreauthenticatedRequestCompositeId(bucket, namespace, parId), nil
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
	return getReplicationPolicyCompositeId(bucket, namespace, replicationId), nil
}

func getOnsNotificationTopicId(resource *OCIResource) (string, error) {
	id, ok := resource.sourceAttributes["topic_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find topic id for ons notification topic")
	}
	return id, nil
}
