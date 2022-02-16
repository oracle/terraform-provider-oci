package resourcediscovery

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v58/loganalytics"

	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/log_analytics"

	tf_logging "github.com/terraform-providers/terraform-provider-oci/internal/service/logging"

	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v58/networkloadbalancer"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"
	oci_objectstorage "github.com/oracle/oci-go-sdk/v58/objectstorage"

	tf_bds "github.com/terraform-providers/terraform-provider-oci/internal/service/bds"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
	tf_log_analytics "github.com/terraform-providers/terraform-provider-oci/internal/service/log_analytics"
	tf_network_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/network_load_balancer"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

var loadBalancerCertificateNameMap map[string]map[string]string // helper map to generate references for certificate names, stores certificate name to certificate name interpolation

func processDnsRrset(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {

	for _, record := range resources {
		if record.parent == nil {
			continue
		}
		// Populate config file from compositeId
		record.compartmentId = record.parent.compartmentId
		domain, rtype, zoneNameOrId, _, _, err := parseRrsetCompositeId(record.id)
		if err == nil {
			record.sourceAttributes["domain"] = domain
			record.sourceAttributes["rtype"] = rtype
			record.sourceAttributes["zone_name_or_id"] = zoneNameOrId
		}
	}
	return resources, nil
}

func parseRrsetCompositeId(compositeId string) (domain string, rtype string, zoneNameOrId string, scope string, viewId string, err error) {
	parts := strings.Split(compositeId, "/")
	match1, _ := regexp.MatchString("zoneNameOrId/.*/domain/.*/rtype/.*", compositeId)
	match2, _ := regexp.MatchString("zoneNameOrId/.*/domain/.*/rtype/.*/scope/.*/viewId/.*", compositeId)
	if match1 && len(parts) == 6 {
		zoneNameOrId, _ = url.PathUnescape(parts[1])
		domain, _ = url.PathUnescape(parts[3])
		rtype, _ = url.PathUnescape(parts[5])
	} else if match2 && len(parts) == 10 {
		zoneNameOrId, _ = url.PathUnescape(parts[1])
		domain, _ = url.PathUnescape(parts[3])
		rtype, _ = url.PathUnescape(parts[5])
		scope, _ = url.PathUnescape(parts[7])
		viewId, _ = url.PathUnescape(parts[9])
	} else {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	return
}

func findDnsRrset(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) (resources []*OCIResource, err error) {
	// Rrset is singular datasource only
	// and need to use GetZoneRecordsRequest to list all records
	zoneId := parent.id
	request := oci_dns.GetZoneRecordsRequest{}
	request.ZoneNameOrId = &zoneId
	response, err := ctx.clients.DnsClient().GetZoneRecords(context.Background(), request)

	if err != nil {
		return resources, err
	}

	for _, record := range response.Items {
		recordResource := resourcesMap[tfMeta.resourceClass]
		d := recordResource.TestResourceData()
		zoneId := parent.id
		domain := record.Domain
		rtype := record.Rtype
		d.SetId(getRrsetCompositeId(*domain, *rtype, zoneId))
		if err := recordResource.Read(d, ctx.clients); err != nil {
			rdError := &ResourceDiscoveryError{tfMeta.resourceClass, parent.terraformName, err, resourceGraph}
			ctx.addErrorToList(rdError)
			continue
		}
		resource := &OCIResource{
			compartmentId:    parent.compartmentId,
			sourceAttributes: convertResourceDataToMap(recordResource.Schema, d),
			rawResource:      record,
			TerraformResource: TerraformResource{
				id:             d.Id(),
				terraformClass: tfMeta.resourceClass,
				terraformName:  fmt.Sprintf("%s_%s", parent.parent.terraformName, *record.RecordHash),
			},
			getHclStringFn: getHclStringFromGenericMap,
			parent:         parent,
		}
		resources = append(resources, resource)
	}

	return resources, err
}

/*
// Custom functions to alter behavior of resource discovery and resource HCL representation

func getModelProvenanceId(resource *OCIResource) (string, error) {
	modelId := resource.parent.id

	return getModelProvenanceCompositeId(modelId), nil
}
*/

func processCorePublicIp(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	publicIps := []*OCIResource{}

	for _, publicIp := range resources {

		if lifeTime, exists := publicIp.sourceAttributes["lifetime"].(string); exists {
			// this is public IP created by NAT gateway
			if lifeTime == "EPHEMERAL" {
				continue
			}
		}
		publicIps = append(publicIps, publicIp)
	}

	return publicIps, nil
}

func processContainerengineNodePool(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, nodePool := range resources {
		// subnet_ids and quantity_per_subnet are deprecated and conflict with node_config_details
		if _, exists := nodePool.sourceAttributes["node_config_details"]; exists {
			if _, ok := nodePool.sourceAttributes["subnet_ids"]; ok {
				delete(nodePool.sourceAttributes, "subnet_ids")
			}
			if _, ok := nodePool.sourceAttributes["quantity_per_subnet"]; ok {
				delete(nodePool.sourceAttributes, "quantity_per_subnet")
			}
		}
	}
	return resources, nil
}

func processStreamingStream(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, streamingStream := range resources {
		// compartment_id conflict with stream_pool_id
		if _, exists := streamingStream.sourceAttributes["compartment_id"]; exists {
			if _, ok := streamingStream.sourceAttributes["stream_pool_id"]; ok {
				delete(streamingStream.sourceAttributes, "stream_pool_id")
			}
		}
	}
	return resources, nil
}
func processNosqlIndex(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, index := range resources {
		if index.parent == nil {
			continue
		}
		index.sourceAttributes["table_name_or_id"] = index.parent.id
	}
	return resources, nil
}

func processKmsKey(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["management_endpoint"] = resource.parent.sourceAttributes["management_endpoint"].(string)
		var resourceSchema *schema.ResourceData = resource.rawResource.(*schema.ResourceData)
		resource.sourceAttributes["id"] = resourceSchema.Id()
	}
	return resources, nil
}

func processKmsKeyVersion(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["management_endpoint"] = resource.parent.sourceAttributes["management_endpoint"].(string)
		resource.importId = resource.id
	}
	return resources, nil
}

// Custom functions to alter behavior of resource discovery and resource HCL representation

func processPrivateIps(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	privateIps := []*OCIResource{}

	for _, privateIp := range resources {

		if privateIp.hasFreeformTag(ResourceCreatedByInstancePool) {
			continue
		}

		// OKE will add tagging support, for now we rely on Automatic default tags for tenancies created after December 17, 2019
		if privateIp.hasDefinedTag(OracleTagsCreatedBy, OkeTagValue) {
			continue
		}

		privateIps = append(privateIps, privateIp)
	}

	return privateIps, nil
}

func processInstances(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	for _, instance := range resources {
		// Omit any resources that were launched by an instance pool. Those shouldn't be managed by Terraform as they are created
		// and managed through the instance pool resource instead.
		if instance.hasFreeformTag(ResourceCreatedByInstancePool) {
			continue
		}

		// OKE will add tagging support, for now we rely on Automatic default tags for tenancies created after December 17, 2019
		if instance.hasDefinedTag(OracleTagsCreatedBy, OkeTagValue) {
			continue
		}

		// Ensure the boot volume created by this instance can be referenced elsewhere by adding it to the reference map
		if bootVolumeId, exists := instance.sourceAttributes["boot_volume_id"]; exists {
			if bootVolumeIdStr, ok := bootVolumeId.(string); ok {
				refMapLock.Lock()
				referenceMap[bootVolumeIdStr] = tfHclVersion.getDoubleExpHclString(instance.getTerraformReference(), "boot_volume_id")
				refMapLock.Unlock()
			}
		}

		if rawSourceDetailsList, sourceDetailsExist := instance.sourceAttributes["source_details"]; sourceDetailsExist {
			if sourceDetailList, ok := rawSourceDetailsList.([]interface{}); ok && len(sourceDetailList) > 0 {
				if sourceDetails, ok := sourceDetailList[0].(map[string]interface{}); ok {
					if imageId, ok := instance.sourceAttributes["image"].(string); ok {
						sourceDetails["source_id"] = imageId

						// The image OCID may be different if it's in a different tenancy or region, add a variable for users to specify
						imageVarName := fmt.Sprintf("%s_source_image_id", instance.terraformName)
						vars[imageVarName] = fmt.Sprintf("\"%s\"", imageId)
						refMapLock.Lock()
						referenceMap[imageId] = tfHclVersion.getVarHclString(imageVarName)
						refMapLock.Unlock()
					}

					// Workaround for service limitation. Service returns 47GB size for boot volume but LaunchInstance can only
					// accept sizes 50GB and above. If such a situation arises, fall back to service default values for boot volume size.
					if bootVolumeSizeInGbs, exists := sourceDetails["boot_volume_size_in_gbs"]; exists {
						bootVolumeSize, err := strconv.ParseInt(bootVolumeSizeInGbs.(string), 10, 64)
						if err != nil {
							return resources, err
						}

						if bootVolumeSize < 50 {
							delete(sourceDetails, "boot_volume_size_in_gbs")
						}
					}
				}
			}
		}

		results = append(results, instance)
	}

	return results, nil
}

func filterSecondaryVnicAttachments(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	for _, attachment := range resources {
		// Filter out any primary vnics, as it's not necessary to Create separate TF resources for those.
		datasourceSchema := datasourcesMap["oci_core_vnic"]
		if vnicReadFn := datasourceSchema.Read; vnicReadFn != nil {
			d := datasourceSchema.TestResourceData()
			d.Set("vnic_id", attachment.sourceAttributes["vnic_id"].(string))
			if err := vnicReadFn(d, ctx.clients); err != nil {
				return results, err
			}

			if isPrimaryVnic, ok := d.GetOkExists("is_primary"); ok && isPrimaryVnic.(bool) {
				continue
			}
		}
		results = append(results, attachment)
	}

	return results, nil
}

func filterMysqlBackups(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	// Filter out Mysql Backups that are automatically created. We cannot operate on "Automatic" backups.
	for _, backup := range resources {
		sourceDetails, exists := backup.sourceAttributes["creation_type"]

		if exists && sourceDetails.(string) == "AUTOMATIC" {
			continue
		}

		results = append(results, backup)
	}

	return results, nil
}

// TODO: remove this when service fixes source
func processMysqlDbSystem(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, dbSystem := range resources {
		if source, exists := dbSystem.sourceAttributes["source"]; exists {
			if sourceList := source.([]interface{}); len(sourceList) > 0 {
				if sourceMap, ok := sourceList[0].(map[string]interface{}); ok {
					if sourceMap["source_type"].(string) == "NONE" {
						delete(dbSystem.sourceAttributes, "source")
					}
				}
			}
		}
	}

	return resources, nil
}

func filterSourcedBootVolumes(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	// Filter out boot volumes that don't have source details. We cannot Create boot volumes unless they have source details.
	for _, bootVolume := range resources {
		sourceDetails, exists := bootVolume.sourceAttributes["source_details"]
		if !exists {
			continue
		}

		if sourceDetailsList, ok := sourceDetails.([]interface{}); !ok || len(sourceDetailsList) == 0 {
			continue
		}

		results = append(results, bootVolume)
	}

	return results, nil
}

func processAvailabilityDomains(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for idx, ad := range resources {
		ad.sourceAttributes["index"] = idx + 1

		adName, ok := ad.sourceAttributes["name"].(string)
		if !ok || adName == "" {
			return resources, fmt.Errorf("[ERROR] availability domain at index '%v' has no name\n", idx)
		}
		refMapLock.Lock()
		referenceMap[adName] = tfHclVersion.getDataSourceHclString(ad.getTerraformReference(), "name")
		refMapLock.Unlock()
	}

	return resources, nil
}

func processObjectStorageNamespace(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, ns := range resources {
		namespaceName, ok := ns.sourceAttributes["namespace"].(string)
		if !ok || namespaceName == "" {
			return resources, fmt.Errorf("[ERROR] object storage namespace data source has no name\n")
		}
		refMapLock.Lock()
		referenceMap[namespaceName] = tfHclVersion.getDataSourceHclString(ns.getTerraformReference(), "namespace")
		refMapLock.Unlock()
	}

	return resources, nil
}

func getAvailabilityDomainHCLDatasource(builder *strings.Builder, ociRes *OCIResource, varMap map[string]string) error {
	builder.WriteString(fmt.Sprintf("data %s %s {\n", ociRes.terraformClass, ociRes.terraformName))

	builder.WriteString(fmt.Sprintf("compartment_id = %v\n", varMap[ociRes.compartmentId]))

	adIndex, ok := ociRes.sourceAttributes["index"]
	if !ok {
		return fmt.Errorf("[ERROR] no index found for availability domain '%s'", ociRes.getTerraformReference())
	}
	builder.WriteString(fmt.Sprintf("ad_number = \"%v\"\n", adIndex.(int)))
	builder.WriteString("}\n")

	return nil
}

func getObjectStorageNamespaceHCLDatasource(builder *strings.Builder, ociRes *OCIResource, varMap map[string]string) error {
	builder.WriteString(fmt.Sprintf("data %s %s {\n", ociRes.terraformClass, ociRes.terraformName))
	builder.WriteString(fmt.Sprintf("compartment_id = %v\n", varMap[ociRes.compartmentId]))
	builder.WriteString("}\n")

	return nil
}

func filterCustomImages(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	// Filter out official images that are predefined by Oracle. We cannot manage such images in Terraform.
	// Official images have a null or empty compartment ID.
	for _, image := range resources {
		compartmentId, exists := image.sourceAttributes["compartment_id"]
		if !exists {
			continue
		}

		if compartmentIdString, ok := compartmentId.(string); !ok || len(compartmentIdString) == 0 {
			continue
		}

		results = append(results, image)
	}

	return results, nil
}

func processVolumeGroups(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Replace the volume Group's source details volume list with the actual volume list
	// The source details only captures the list of volumes that were known when the Group was created.
	// Additional volumes may have been added since and should be part of the source_details that we generate.
	// TODO: This is a shortcoming that should be addressed by the service and/or the Terraform
	for _, group := range resources {
		volumeIdsRaw, exists := group.sourceAttributes["volume_ids"]
		if !exists {
			continue
		}

		if volumeIds, ok := volumeIdsRaw.([]interface{}); ok && len(volumeIds) > 0 {
			sourceDetailsRaw, detailsExist := group.sourceAttributes["source_details"]
			if !detailsExist {
				continue
			}

			sourceDetails := sourceDetailsRaw.([]interface{})[0].(map[string]interface{})
			sourceDetails["volume_ids"] = volumeIds
		}
	}

	return resources, nil
}

func processLoadBalancerBackendSets(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, backendSet := range resources {
		if backendSet.parent == nil {
			continue
		}

		backendSetName := backendSet.sourceAttributes["name"].(string)
		backendSet.id = tf_load_balancer.GetBackendSetCompositeId(backendSetName, backendSet.parent.id)
		backendSet.sourceAttributes["load_balancer_id"] = backendSet.parent.id
	}

	return resources, nil
}

func processLoadBalancerBackends(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, backend := range resources {
		if backend.parent == nil {
			continue
		}

		backend.id = tf_load_balancer.GetBackendCompositeId(backend.sourceAttributes["name"].(string), backend.parent.sourceAttributes["name"].(string), backend.parent.sourceAttributes["load_balancer_id"].(string))
		backend.sourceAttributes["load_balancer_id"] = backend.parent.sourceAttributes["load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !backend.parent.omitFromExport {
			backend.sourceAttributes["backendset_name"] = InterpolationString{
				resourceReference: backend.parent.getTerraformReference(),
				interpolation:     tfHclVersion.getDoubleExpHclString(backend.parent.getTerraformReference(), "name"),
				value:             backend.parent.sourceAttributes["name"].(string),
			}
		} else {
			backend.sourceAttributes["backendset_name"] = backend.parent.sourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func processLoadBalancerHostnames(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, hostname := range resources {
		if hostname.parent == nil {
			continue
		}

		hostname.id = tf_load_balancer.GetHostnameCompositeId(hostname.parent.id, hostname.sourceAttributes["name"].(string))
		hostname.sourceAttributes["load_balancer_id"] = hostname.parent.id
	}

	return resources, nil
}

func processLoadBalancerPathRouteSets(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, pathRouteSet := range resources {
		if pathRouteSet.parent == nil {
			continue
		}

		pathRouteSet.id = tf_load_balancer.GetPathRouteSetCompositeId(pathRouteSet.parent.id, pathRouteSet.sourceAttributes["name"].(string))
		pathRouteSet.sourceAttributes["load_balancer_id"] = pathRouteSet.parent.id
	}

	return resources, nil
}

func processLoadBalancerRoutingPolicies(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, routingPolicy := range resources {
		if routingPolicy.parent == nil {
			continue
		}

		routingPolicy.id = tf_load_balancer.GetLoadBalancerRoutingPolicyCompositeId(routingPolicy.parent.id, routingPolicy.sourceAttributes["name"].(string))
		routingPolicy.sourceAttributes["load_balancer_id"] = routingPolicy.parent.id
	}

	return resources, nil
}

func processLoadBalancerRuleSets(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, ruleSet := range resources {
		if ruleSet.parent == nil {
			continue
		}

		ruleSet.id = tf_load_balancer.GetRuleSetCompositeId(ruleSet.parent.id, ruleSet.sourceAttributes["name"].(string))
		ruleSet.sourceAttributes["load_balancer_id"] = ruleSet.parent.id
	}

	return resources, nil
}

func processLoadBalancerCertificates(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, certificate := range resources {
		if certificate.parent == nil {
			continue
		}

		certificate.id = tf_load_balancer.GetCertificateCompositeId(certificate.sourceAttributes["certificate_name"].(string), certificate.parent.id)
		certificate.sourceAttributes["load_balancer_id"] = certificate.parent.id

		// add certificate name and interpolation to loadBalancerCertificateNameMap
		if loadBalancerCertificateNameMap == nil {
			loadBalancerCertificateNameMap = make(map[string]map[string]string)
		}
		_, ok := loadBalancerCertificateNameMap[certificate.parent.id]
		if !ok {
			loadBalancerCertificateNameMap[certificate.parent.id] = make(map[string]string)
		}

		if certificateName, ok := certificate.sourceAttributes["certificate_name"].(string); ok {
			loadBalancerCertificateNameMap[certificate.parent.id][certificateName] = tfHclVersion.getDoubleExpHclString(certificate.getTerraformReference(), "certificate_name")
		}
	}

	return resources, nil
}

func processObjectStoragePreauthenticatedRequest(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["bucket"] = resource.parent.sourceAttributes["name"].(string)
		resource.sourceAttributes["namespace"] = resource.parent.sourceAttributes["namespace"].(string)

		// Check if time is already in RFC3339Nano format
		timeExpires, err := time.Parse(time.RFC3339Nano, resource.sourceAttributes["time_expires"].(string))
		if err != nil {
			// parse time using format in time.String()
			timeExpires, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", resource.sourceAttributes["time_expires"].(string))
			if err != nil {
				return resources, err
			}
			// Format to RFC3339Nano
			resource.sourceAttributes["time_expires"] = timeExpires.Format(time.RFC3339Nano)
		}

	}
	return resources, nil
}

func processAutonomousDatabaseSource(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.sourceAttributes["is_refreshable_clone"] == true {
			resource.sourceAttributes["source"] = "CLONE_TO_REFRESHABLE"
		}
	}
	return resources, nil
}

func processObjectStorageReplicationPolicy(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["bucket"] = resource.parent.sourceAttributes["name"].(string)
		resource.sourceAttributes["namespace"] = resource.parent.sourceAttributes["namespace"].(string)
	}
	return resources, nil
}

func processLogAnalyticsObjectCollectionRules(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		namespace := resource.sourceAttributes["namespace"].(string)
		logAnalyticsObjectCollectionRuleId := resource.id
		resource.importId = log_analytics.GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId, namespace)
	}

	return resources, nil
}

func processNetworkLoadBalancerBackendSets(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, backendSet := range resources {
		if backendSet.parent == nil {
			continue
		}

		backendSetName := backendSet.sourceAttributes["name"].(string)
		backendSet.id = tf_network_load_balancer.GetNlbBackendSetCompositeId(backendSetName, backendSet.parent.id)
		backendSet.sourceAttributes["network_load_balancer_id"] = backendSet.parent.id
	}

	return resources, nil
}

func processNetworkLoadBalancerBackends(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, backend := range resources {
		if backend.parent == nil {
			continue
		}

		backend.id = tf_network_load_balancer.GetNlbBackendCompositeId(backend.sourceAttributes["name"].(string), backend.parent.sourceAttributes["name"].(string), backend.parent.sourceAttributes["network_load_balancer_id"].(string))
		backend.sourceAttributes["network_load_balancer_id"] = backend.parent.sourceAttributes["network_load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !backend.parent.omitFromExport {
			backend.sourceAttributes["backend_set_name"] = InterpolationString{
				resourceReference: backend.parent.getTerraformReference(),
				interpolation:     tfHclVersion.getDoubleExpHclString(backend.parent.getTerraformReference(), "name"),
				value:             backend.parent.sourceAttributes["name"].(string),
			}
		} else {
			backend.sourceAttributes["backend_set_name"] = backend.parent.sourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func findNetworkLoadBalancerListeners(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	networkLoadBalancerId := parent.sourceAttributes["network_load_balancer_id"].(string)
	backendSetName := parent.sourceAttributes["name"].(string)

	request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}
	request.NetworkLoadBalancerId = &networkLoadBalancerId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

	response, err := ctx.clients.NetworkLoadBalancerClient().GetNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}

	listenerResource := resourcesMap[tfMeta.resourceClass]

	results := []*OCIResource{}
	for listenerName, listener := range response.NetworkLoadBalancer.Listeners {
		if *listener.DefaultBackendSetName != backendSetName {
			continue
		}

		d := listenerResource.TestResourceData()
		d.SetId(tf_network_load_balancer.GetNlbListenerCompositeId(listenerName, networkLoadBalancerId))

		// This calls into the listener resource's Read fn which has the unfortunate implementation of
		// calling GetNetworkLoadBalancer and looping through the listeners to find the expected one. So this entire method
		// may require O(n^^2) time. However, the benefits of having Read populate the ResourceData struct is better than duplicating it here.
		if err := listenerResource.Read(d, ctx.clients); err != nil {
			// add error to the errorList and continue discovering rest of the resources
			rdError := &ResourceDiscoveryError{tfMeta.resourceClass, parent.terraformName, err, resourceGraph}
			ctx.addErrorToList(rdError)
			continue
		}

		resource := &OCIResource{
			compartmentId:    parent.compartmentId,
			sourceAttributes: convertResourceDataToMap(listenerResource.Schema, d),
			rawResource:      listener,
			TerraformResource: TerraformResource{
				id:             d.Id(),
				terraformClass: tfMeta.resourceClass,
				terraformName:  fmt.Sprintf("%s_%s", parent.parent.terraformName, listenerName),
			},
			getHclStringFn: getHclStringFromGenericMap,
			parent:         parent,
		}

		if !parent.omitFromExport {
			resource.sourceAttributes["default_backend_set_name"] = InterpolationString{
				resourceReference: parent.getTerraformReference(),
				interpolation:     tfHclVersion.getDoubleExpHclString(parent.getTerraformReference(), "name"),
				value:             parent.sourceAttributes["name"].(string),
			}
		} else {
			resource.sourceAttributes["default_backend_set_name"] = parent.sourceAttributes["name"].(string)
		}
		results = append(results, resource)
	}

	return results, nil
}

func processNetworkLoadBalancerListeners(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, listener := range resources {
		if listener.parent == nil {
			continue
		}

		listenerName := listener.sourceAttributes["name"].(string)
		listener.id = tf_network_load_balancer.GetNlbListenerCompositeId(listenerName, listener.parent.sourceAttributes["network_load_balancer_id"].(string))
		listener.sourceAttributes["network_load_balancer_id"] = listener.parent.sourceAttributes["network_load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !listener.parent.omitFromExport {
			listener.sourceAttributes["default_backend_set_name"] = InterpolationString{
				resourceReference: listener.parent.getTerraformReference(),
				interpolation:     tfHclVersion.getDoubleExpHclString(listener.parent.getTerraformReference(), "name"),
				value:             listener.parent.sourceAttributes["name"].(string),
			}
		} else {
			listener.sourceAttributes["default_backend_set_name"] = listener.parent.sourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func findIdentityTags(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	// List on Tags does not return validator, and resource Read requires tagNamespaceId
	// which is also not returned in Summary response. Tags also do not have composite id in state.
	// Getting tags using ListTagsRequest and the calling tagResource.Read
	tagNamespaceId := parent.id
	request := oci_identity.ListTagsRequest{}

	request.TagNamespaceId = &tagNamespaceId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
	results := []*OCIResource{}

	response, err := ctx.clients.IdentityClient().ListTags(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.clients.IdentityClient().ListTags(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, tag := range response.Items {
		tagResource := resourcesMap[tfMeta.resourceClass]

		d := tagResource.TestResourceData()
		d.SetId(tf_identity.GetIdentityTagCompositeId(*tag.Name, parent.id))

		if err := tagResource.Read(d, ctx.clients); err != nil {
			rdError := &ResourceDiscoveryError{tfMeta.resourceClass, parent.terraformName, err, resourceGraph}
			ctx.addErrorToList(rdError)
			continue
		}

		resource := &OCIResource{
			compartmentId:    parent.compartmentId,
			sourceAttributes: convertResourceDataToMap(tagResource.Schema, d),
			rawResource:      tag,
			TerraformResource: TerraformResource{
				id:             d.Id(),
				terraformClass: tfMeta.resourceClass,
			},
			getHclStringFn: getHclStringFromGenericMap,
			parent:         parent,
		}

		if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, tagResource.Schema); err != nil {
			resource.terraformName = fmt.Sprintf("%s_%s", parent.parent.terraformName, *tag.Name)
		}

		results = append(results, resource)
	}

	return results, nil

}

func findLoadBalancerListeners(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	loadBalancerId := parent.sourceAttributes["load_balancer_id"].(string)
	backendSetName := parent.sourceAttributes["name"].(string)

	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

	response, err := ctx.clients.LoadBalancerClient().GetLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}

	listenerResource := resourcesMap[tfMeta.resourceClass]

	results := []*OCIResource{}
	for listenerName, listener := range response.LoadBalancer.Listeners {
		if *listener.DefaultBackendSetName != backendSetName {
			continue
		}

		d := listenerResource.TestResourceData()
		d.SetId(tf_load_balancer.GetListenerCompositeId(listenerName, loadBalancerId))

		// This calls into the listener resource's Read fn which has the unfortunate implementation of
		// calling GetLoadBalancer and looping through the listeners to find the expected one. So this entire method
		// may require O(n^^2) time. However, the benefits of having Read populate the ResourceData struct is better than duplicating it here.
		if err := listenerResource.Read(d, ctx.clients); err != nil {
			// add error to the errorList and continue discovering rest of the resources
			rdError := &ResourceDiscoveryError{tfMeta.resourceClass, parent.terraformName, err, resourceGraph}
			ctx.addErrorToList(rdError)
			continue
		}

		resource := &OCIResource{
			compartmentId:    parent.compartmentId,
			sourceAttributes: convertResourceDataToMap(listenerResource.Schema, d),
			rawResource:      listener,
			TerraformResource: TerraformResource{
				id:             d.Id(),
				terraformClass: tfMeta.resourceClass,
				terraformName:  fmt.Sprintf("%s_%s", parent.parent.terraformName, listenerName),
			},
			getHclStringFn: getHclStringFromGenericMap,
			parent:         parent,
		}

		if !parent.omitFromExport {
			resource.sourceAttributes["default_backend_set_name"] = InterpolationString{
				resourceReference: parent.getTerraformReference(),
				interpolation:     tfHclVersion.getDoubleExpHclString(parent.getTerraformReference(), "name"),
				value:             parent.sourceAttributes["name"].(string),
			}
		} else {
			resource.sourceAttributes["default_backend_set_name"] = parent.sourceAttributes["name"].(string)
		}
		results = append(results, resource)
	}

	return results, nil
}

func findLogAnalyticsObjectCollectionRules(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	// List on LogAnalyticsObjectCollectionRules requires namespaceName path parameter.
	// Getting namespace from ObjectStorage.GetNamespace API before calling ListLogAnalyticsObjectCollectionRules API.
	results := []*OCIResource{}

	namespaceRequest := oci_objectstorage.GetNamespaceRequest{}
	namespaceResponse, err := ctx.clients.ObjectStorageClient().GetNamespace(context.Background(), namespaceRequest)
	if err != nil {
		return results, err
	}
	namespace := namespaceResponse.Value
	request := oci_log_analytics.ListLogAnalyticsObjectCollectionRulesRequest{}

	request.NamespaceName = namespace
	request.CompartmentId = ctx.CompartmentId
	request.LifecycleState = oci_log_analytics.ListLogAnalyticsObjectCollectionRulesLifecycleStateActive

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

	response, err := ctx.clients.LogAnalyticsClient().ListLogAnalyticsObjectCollectionRules(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.clients.LogAnalyticsClient().ListLogAnalyticsObjectCollectionRules(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, logAnalyticsObjectCollectionRule := range response.Items {
		logAnalyticsObjectCollectionRuleResource := resourcesMap[tfMeta.resourceClass]

		d := logAnalyticsObjectCollectionRuleResource.TestResourceData()
		d.SetId(tf_log_analytics.GetLogAnalyticsObjectCollectionRuleCompositeId(*logAnalyticsObjectCollectionRule.Id, *namespace))

		if err := logAnalyticsObjectCollectionRuleResource.Read(d, ctx.clients); err != nil {
			rdError := &ResourceDiscoveryError{tfMeta.resourceClass, parent.terraformName, err, resourceGraph}
			ctx.addErrorToList(rdError)
			continue
		}

		resource := &OCIResource{
			compartmentId:    *ctx.CompartmentId,
			sourceAttributes: convertResourceDataToMap(logAnalyticsObjectCollectionRuleResource.Schema, d),
			rawResource:      logAnalyticsObjectCollectionRule,
			TerraformResource: TerraformResource{
				id:             d.Id(),
				terraformClass: tfMeta.resourceClass,
			},
			getHclStringFn: getHclStringFromGenericMap,
			parent:         parent,
		}

		if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, logAnalyticsObjectCollectionRuleResource.Schema); err != nil {
			resource.terraformName = fmt.Sprintf("%s_%s", parent.parent.terraformName, *logAnalyticsObjectCollectionRule.Name)
		}

		results = append(results, resource)
	}

	return results, nil

}

func processLoadBalancerListeners(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {

	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		if sslConfiguration, ok := resource.sourceAttributes["ssl_configuration"].([]interface{}); ok && len(sslConfiguration) > 0 {
			if sslConfig, ok := sslConfiguration[0].(map[string]interface{}); ok {
				if certificateName, ok := sslConfig["certificate_name"]; ok {
					// check if we have expected ResourceIds set, is load balancer certificate id expected
					if ctx.expectedResourceIds != nil && len(ctx.expectedResourceIds) > 0 {
						certificateId := tf_load_balancer.GetCertificateCompositeId(certificateName.(string), resource.sourceAttributes["load_balancer_id"].(string))
						if _, ok = ctx.expectedResourceIds[certificateId]; !ok {
							continue
						}
					}
					sslConfig["certificate_name"] = InterpolationString{
						resource.parent.getTerraformReference(),
						loadBalancerCertificateNameMap[resource.parent.parent.id][sslConfig["certificate_name"].(string)],
						sslConfig["certificate_name"].(string),
					}
				}
			}
		}
	}
	return resources, nil
}

func processTagDefinitions(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			resource.importId = fmt.Sprintf("tagNamespaces/%s/tags/%s", resource.sourceAttributes["tag_namespace_id"], resource.sourceAttributes["name"].(string))
			continue
		}

		resource.sourceAttributes["tag_namespace_id"] = resource.parent.id
		resource.importId = fmt.Sprintf("tagNamespaces/%s/tags/%s", resource.parent.id, resource.sourceAttributes["name"].(string))
		resource.id = resource.importId
	}
	return resources, nil
}

func processNetworkSecurityGroupRules(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}

		resource.sourceAttributes["network_security_group_id"] = resource.parent.id
	}
	return resources, nil
}

func processDrgRouteDistributionStatements(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["drg_route_distribution_id"] = resource.parent.id
	}
	return resources, nil
}

func processDrgRouteTableRouteRules(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		resource.sourceAttributes["drg_route_table_id"] = resource.parent.id
	}
	return resources, nil
}

func filterPrimaryDbHomes(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// No need to filter if db homes are in vm cluster
	if len(resources) > 0 && resources[0].parent != nil && resources[0].parent.terraformClass == "oci_database_vm_cluster" {
		return resources, nil
	}
	results := []*OCIResource{}
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		// If we found a db home that matches the db system's primary home, then don't return it as part of result
		if dbSystem := resource.parent; dbSystem != nil {
			if dbHomes, ok := dbSystem.sourceAttributes["db_home"].([]interface{}); ok && len(dbHomes) > 0 {
				if primaryDbHome, ok := dbHomes[0].(map[string]interface{}); ok {
					if primaryDbHomeId, ok := primaryDbHome["id"]; ok && primaryDbHomeId.(string) == resource.id {
						continue
					}
				}
			}
		}
		// Fix db version to remove the PSU date from versions with 18+ major version
		if dbVersion, ok := resource.sourceAttributes["db_version"].(string); ok {
			resource.sourceAttributes["db_version"] = getValidDbVersion(dbVersion)
		}
		results = append(results, resource)
	}
	return results, nil
}

func filterPrimaryDatabases(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		// Only return database resources that don't match the database ID of the dbHome resource.
		if databases, ok := resource.parent.sourceAttributes["database"].([]interface{}); ok && len(databases) > 0 {
			if primaryDatabase, ok := databases[0].(map[string]interface{}); ok {
				if primaryDatabaseId, ok := primaryDatabase["id"]; ok && primaryDatabaseId.(string) != resource.id {
					results = append(results, resource)
				}
			}
		}
	}
	return results, nil
}

func processIdentityAuthenticationPolicies(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Add composite id as the resource's import ID
	for _, resource := range resources {
		resource.importId = tf_identity.GetAuthenticationPolicyCompositeId(resource.compartmentId)
		resource.id = resource.importId
	}
	return resources, nil
}

func processDefaultSecurityLists(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Default security lists need to be handled as default resources
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}

		if resource.id == resource.parent.sourceAttributes["default_security_list_id"].(string) {
			resource.sourceAttributes["manage_default_resource_id"] = resource.id
			resource.TerraformResource.terraformClass = "oci_core_default_security_list"

			// Don't use references to parent resources if they will be omitted from final result
			if !resource.parent.omitFromExport {
				resource.TerraformResource.terraformReferenceIdString = fmt.Sprintf("%s.%s", resource.parent.getTerraformReference(), "default_security_list_id")
			}
		}
	}
	return resources, nil
}

func processDefaultRouteTables(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Default route tables need to be handled as default resources
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}

		if resource.id == resource.parent.sourceAttributes["default_route_table_id"].(string) {
			resource.sourceAttributes["manage_default_resource_id"] = resource.id
			resource.TerraformResource.terraformClass = "oci_core_default_route_table"

			// Don't use references to parent resources if they will be omitted from final result
			if !resource.parent.omitFromExport {
				resource.TerraformResource.terraformReferenceIdString = fmt.Sprintf("%s.%s", resource.parent.getTerraformReference(), "default_route_table_id")
			}
		}
	}
	return resources, nil
}

func processDefaultDhcpOptions(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Default dhcp options need to be handled as default resources
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}

		if resource.id == resource.parent.sourceAttributes["default_dhcp_options_id"].(string) {
			resource.sourceAttributes["manage_default_resource_id"] = resource.id
			resource.TerraformResource.terraformClass = "oci_core_default_dhcp_options"

			// Don't use references to parent resources if they will be omitted from final result
			if !resource.parent.omitFromExport {
				resource.TerraformResource.terraformReferenceIdString = fmt.Sprintf("%s.%s", resource.parent.getTerraformReference(), "default_dhcp_options_id")
			}
		}
	}
	return resources, nil
}

func processDbSystems(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Fix db version to remove the PSU date from versions with 18+ major version
	for _, resource := range resources {
		if dbHomes, ok := resource.sourceAttributes["db_home"].([]interface{}); ok {
			if dbHome, ok := dbHomes[0].(map[string]interface{}); ok {
				if dbVersion, ok := dbHome["db_version"].(string); ok {
					dbHome["db_version"] = getValidDbVersion(dbVersion)
				}
			}
		}
	}
	return resources, nil
}

func processDatabases(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Fix database db version to remove the PSU date from versions with 18+ major version
	for _, resource := range resources {
		if databases, ok := resource.sourceAttributes["database"].([]interface{}); ok {
			if database, ok := databases[0].(map[string]interface{}); ok {
				if dbVersion, ok := database["db_version"].(string); ok {
					database["db_version"] = getValidDbVersion(dbVersion)
				}
			}
		}
	}
	return resources, nil
}

func processDatabaseExadataInfrastructures(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// Remove weeks_of_month if there is no item in response
	for _, resource := range resources {
		if maintenanceWindow, ok := resource.sourceAttributes["maintenance_window"].([]interface{}); ok {
			if mWindow, ok := maintenanceWindow[0].(map[string]interface{}); ok {
				if weeksOfMonth, ok := mWindow["weeks_of_month"].([]interface{}); ok && len(weeksOfMonth) == 0 {
					delete(mWindow, "weeks_of_month")
				}
			}
		}
	}
	return resources, nil
}

func getValidDbVersion(dbVersion string) string {
	/*
		For 11.2.0.4, 12.1.0.2 and 12.2.0.1, the PSU is added as the 5th digit. So when the customer specifies either of these,
		service will be returning 11.2.0.4.xxxxxx where the last part is the PSU version.
		For 18.0.0.0 and 19.0.0.0 onwards, the second digit specifies the PSU version and the fifth digit specifies the date for that PSU.
		(The PSU-date pair change hand in hand)
		* For pre 18 versions, service returns 5th digit in response and 5 digit version is valid for Create
		* For 18+ versions, service will return PSU date but only 4 digit version is valid for Create.
		* Resource discovery will keep only 4 digits in config and dbVersionDiffSuppress will handle the diff
	*/
	parts := strings.Split(dbVersion, ".")
	if strings.Compare(parts[0], "18") == 1 {
		return strings.Join(parts[0:4], ".")
	}
	return dbVersion
}

func getLogId(resource *OCIResource) (string, error) {
	logId, ok := resource.sourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find log_id for Log")
	}
	logGroupId := resource.parent.id
	return tf_logging.GetLogCompositeId(logGroupId, logId), nil
}

func processCoreVcns(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	// remove deprecated cidr_block field from discovered vcns,
	// either cidr_block or cidr_blocks should be specified in config
	// service returns the cidr_block value in cidr_blocks field
	for _, resource := range resources {
		if _, ok := resource.sourceAttributes["cidr_block"].(string); ok {
			delete(resource.sourceAttributes, "cidr_block")
		}
	}
	return resources, nil
}

func processCertificateAuthorities(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		certificateAuthorityConfigMap := map[string]interface{}{}
		if configType, ok := resource.sourceAttributes["config_type"].(string); ok {
			certificateAuthorityConfigMap["config_type"] = configType
		}

		if subjects, ok := resource.sourceAttributes["subject"].([]interface{}); ok {
			if subject, ok := subjects[0].(map[string]interface{}); ok {
				certificateAuthorityConfigMap["subject"] = []interface{}{subject}
			}
		}

		if issuerCertificateAuthorityId, ok := resource.sourceAttributes["issuer_certificate_authority_id"].(string); ok {
			certificateAuthorityConfigMap["issuer_certificate_authority_id"] = issuerCertificateAuthorityId
		}

		if signingAlgorithm, ok := resource.sourceAttributes["signing_algorithm"].(string); ok {
			certificateAuthorityConfigMap["signing_algorithm"] = signingAlgorithm
		}

		if currentVersions, ok := resource.sourceAttributes["current_version"].([]interface{}); ok {
			if currentVersion, ok := currentVersions[0].(map[string]interface{}); ok {
				if validity, ok := currentVersion["validity"].([]interface{}); ok {
					validityMap := map[string]interface{}{}
					if timeOfValidityNotAfter, ok := validity[0].(map[string]interface{})["time_of_validity_not_after"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotAfter.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_after"] = tmp.Format(time.RFC3339Nano)
						}
					}
					if timeOfValidityNotBefore, ok := validity[0].(map[string]interface{})["time_of_validity_not_before"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotBefore.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_before"] = tmp.Format(time.RFC3339Nano)
						}
					}

					certificateAuthorityConfigMap["validity"] = []interface{}{validityMap}
				}

				if versionName, ok := currentVersion["version_name"].(string); ok {
					certificateAuthorityConfigMap["version_name"] = versionName
				}
			}
		}

		resource.sourceAttributes["certificate_authority_config"] = []interface{}{certificateAuthorityConfigMap}
	}

	return resources, nil
}

func processCertificates(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		certificateConfigMap := map[string]interface{}{}
		if configType, ok := resource.sourceAttributes["config_type"].(string); ok {
			certificateConfigMap["config_type"] = configType
		}

		if profileType, ok := resource.sourceAttributes["certificate_profile_type"].(string); ok {
			certificateConfigMap["certificate_profile_type"] = profileType
		}

		if csrPem, ok := resource.sourceAttributes["csr_pem"].(string); ok {
			certificateConfigMap["csr_pem"] = csrPem
		}

		if issuerCertificateAuthorityId, ok := resource.sourceAttributes["issuer_certificate_authority_id"].(string); ok {
			certificateConfigMap["issuer_certificate_authority_id"] = issuerCertificateAuthorityId
		}

		if keyAlgorithm, ok := resource.sourceAttributes["key_algorithm"].(string); ok {
			certificateConfigMap["key_algorithm"] = keyAlgorithm
		}

		if signatureAlgorithm, ok := resource.sourceAttributes["signature_algorithm"].(string); ok {
			certificateConfigMap["signature_algorithm"] = signatureAlgorithm
		}

		if subjects, ok := resource.sourceAttributes["subject"].([]interface{}); ok {
			if subject, ok := subjects[0].(map[string]interface{}); ok {
				certificateConfigMap["subject"] = []interface{}{subject}
			}
		}

		if currentVersions, ok := resource.sourceAttributes["current_version"].([]interface{}); ok {
			if currentVersion, ok := currentVersions[0].(map[string]interface{}); ok {
				if validity, ok := currentVersion["validity"].([]interface{}); ok {
					validityMap := map[string]interface{}{}
					if timeOfValidityNotAfter, ok := validity[0].(map[string]interface{})["time_of_validity_not_after"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotAfter.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_after"] = tmp.Format(time.RFC3339Nano)
						}
					}
					if timeOfValidityNotBefore, ok := validity[0].(map[string]interface{})["time_of_validity_not_before"]; ok {
						// Check if time is already in RFC3339Nano format
						tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
						if err != nil {
							// parse time using format in time.String()
							tmp, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeOfValidityNotBefore.(string))
							if err != nil {
								return resources, err
							}
							// Format to RFC3339Nano
							validityMap["time_of_validity_not_before"] = tmp.Format(time.RFC3339Nano)
						}
					}

					certificateConfigMap["validity"] = []interface{}{validityMap}
				}

				if versionName, ok := currentVersion["version_name"].(string); ok {
					certificateConfigMap["version_name"] = versionName
				}

				if subjectAlternativeNames, ok := currentVersion["subject_alternative_names"].([]interface{}); ok {
					tmp := []interface{}{}
					for _, item := range subjectAlternativeNames {
						tmp = append(tmp, item)
					}
					certificateConfigMap["subject_alternative_names"] = tmp
				}
			}
		}

		resource.sourceAttributes["certificate_config"] = []interface{}{certificateConfigMap}
	}

	return resources, nil
}

func processBdsInstanceMetastoreConfigs(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		if resource.parent == nil {
			continue
		}
		metastoreConfigId := resource.id
		bdsInstanceId := resource.parent.id
		resource.importId = tf_bds.GetBdsInstanceMetastoreConfigCompositeId(bdsInstanceId, metastoreConfigId)
	}
	return resources, nil
}
