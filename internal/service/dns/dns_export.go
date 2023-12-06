package dns

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDnsRrsetHints.GetIdFn = getDnsRrsetId
	exportDnsRrsetHints.FindResourcesOverrideFn = findDnsRrset
	exportDnsRrsetHints.ProcessDiscoveredResourcesFn = processDnsRrset
	exportDnsResolverEndpointHints.GetIdFn = getDnsResolverEndpointId
	exportDnsResolverHints.GetHCLStringOverrideFn = getHclStringForDnsResolver
	exportDnsResolverEndpointHints.GetHCLStringOverrideFn = getHclStringForDnsResolverEndpoint
	exportDnsZoneHints.FindResourcesOverrideFn = findDnsZones
	tf_export.RegisterCompartmentGraphs("dns", dnsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getDnsResolverEndpointId(resource *tf_export.OCIResource) (string, error) {

	resolverEndpointName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find resolverEndpointName for Dns ResolverEndpoint")
	}
	resolverId := resource.Parent.Id
	return GetResolverEndpointCompositeId(resolverEndpointName, resolverId), nil
}

func processDnsRrset(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {

	for _, record := range resources {
		if record.Parent == nil {
			continue
		}
		// Populate config file from compositeId
		record.CompartmentId = record.Parent.CompartmentId
		domain, rtype, zoneNameOrId, _, _, err := parseRrsetCompositeId(record.Id)
		if err == nil {
			record.SourceAttributes["domain"] = domain
			record.SourceAttributes["rtype"] = rtype
			record.SourceAttributes["zone_name_or_id"] = zoneNameOrId
		}
	}
	return resources, nil
}

func getDnsRrsetTerraformName(parentTerraformName string, domain string, rtype string) string {

	terraformName := fmt.Sprintf("%s_%s_%s", parentTerraformName, strings.Replace(strings.Replace(domain, "-", "--", -1), ".", "-", -1), rtype)
	reg := regexp.MustCompile(`[^a-zA-Z0-9\-\_]+`)
	terraformName = reg.ReplaceAllString(terraformName, "-")
	terraformName = tf_export.CheckDuplicateResourceName(terraformName)

	return terraformName
}

func findDnsRrset(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) (resources []*tf_export.OCIResource, err error) {
	// Rrset is singular datasource only
	// and need to use GetZoneRecordsRequest to list all records
	zoneId := parent.Id
	request := oci_dns.GetZoneRecordsRequest{}
	request.ZoneNameOrId = &zoneId
	response, err := ctx.Clients.DnsClient().GetZoneRecords(context.Background(), request)

	if err != nil {
		return resources, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.DnsClient().GetZoneRecords(context.Background(), request)
		if err != nil {
			return resources, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	rrsetMap := map[string]map[string][]oci_dns.Record{}

	for _, record := range response.Items {
		if _, ok := rrsetMap[*record.Domain]; !ok {
			rrsetMap[*record.Domain] = map[string][]oci_dns.Record{}
		}
		if _, ok := rrsetMap[*record.Domain][*record.Rtype]; !ok {
			rrsetMap[*record.Domain][*record.Rtype] = []oci_dns.Record{}
		}
		rrsetMap[*record.Domain][*record.Rtype] = append(rrsetMap[*record.Domain][*record.Rtype], record)
	}

	for domain, domainMap := range rrsetMap {
		for rtype, rrset := range domainMap {
			recordResource := tf_export.ResourcesMap[tfMeta.ResourceClass]
			d := recordResource.TestResourceData()
			zoneId := parent.Id
			d.SetId(getRrsetCompositeId(domain, rtype, zoneId))
			if err := recordResource.Read(d, ctx.Clients); err != nil {
				rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
				ctx.AddErrorToList(rdError)
				continue
			}
			resourceHint, err := ctx.GetResourceHint(tfMeta.ResourceClass)
			if err != nil {
				rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
				ctx.AddErrorToList(rdError)
				continue
			}
			resource, err := tf_export.GetOciResource(d, recordResource.Schema, parent.CompartmentId, resourceHint, d.Id())
			if err != nil {
				rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
				ctx.AddErrorToList(rdError)
				continue
			}
			resource.TerraformName = getDnsRrsetTerraformName(parent.TerraformName, domain, rtype)
			resource.RawResource = rrset
			resource.Parent = parent

			resources = append(resources, resource)
		}
	}

	return resources, err
}

func getDnsRrsetId(resource *tf_export.OCIResource) (string, error) {

	domain, ok := resource.SourceAttributes["domain"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find domain for Dns Rrset")
	}
	rtype, ok := resource.SourceAttributes["rtype"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find rtype for Dns Rrset")
	}
	zoneNameOrId := resource.Parent.Id
	return getRrsetCompositeId(domain, rtype, zoneNameOrId), nil
}

func findDnsZones(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) (resources []*tf_export.OCIResource, err error) {
	if tfMeta.DatasourceQueryParams == nil {
		tfMeta.DatasourceQueryParams = map[string]string{}
	}
	// Setting the "scope" field to the special value "ALL" will
	// result in terraform fetching both global and private zones
	// when populating the "oci_dns_zones" data source
	tfMeta.DatasourceQueryParams["scope"] = "'ALL'"
	return tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
}

var getHclStringForDnsResolver = func(builder *strings.Builder, ociRes *tf_export.OCIResource, interpolationMap map[string]string) error {

	// Map endpoint name to endpoint reference
	resolverEndpointRefsMap := make(map[string]string)

	if resolverEndpoints, ok := ociRes.RawResource.(*schema.ResourceData).Get("endpoints").([]interface{}); ok {
		for _, resolverEndpoint := range resolverEndpoints {
			resolverEndpointMap, ok := resolverEndpoint.(map[string]interface{})
			if !ok {
				continue
			}
			name, ok := resolverEndpointMap["name"].(string)
			if !ok {
				continue
			}
			resolverId := ociRes.TerraformResource.Id
			resolverEndpointId := GetResolverEndpointCompositeId(name, resolverId)
			resolverEndpointRef := interpolationMap[resolverEndpointId]
			resolverEndpointRefsMap[name] = fmt.Sprintf("%s.name", strings.TrimSuffix(resolverEndpointRef, ".id"))
		}

		// Replace resolver rules endpoint names with references
		if resolverRules, haveResolverRules := ociRes.SourceAttributes["rules"]; haveResolverRules {
			if resolverRulesList, ok := resolverRules.([]map[string]interface{}); ok {
				for _, rule := range resolverRulesList {
					if sourceEndpointName, haveSourceEndpointName := rule["source_endpoint_name"]; haveSourceEndpointName {
						if sourceEndpointNameStr, ok := sourceEndpointName.(string); ok {
							endpointRef := resolverEndpointRefsMap[sourceEndpointNameStr]
							rule["source_endpoint_name"] = tf_export.InterpolationString{
								ResourceReference: strings.TrimSuffix(endpointRef, ".name"),
								Interpolation:     endpointRef,
								Value:             sourceEndpointNameStr,
							}
						}
					}
				}
			}
		}
	}

	return tf_export.GetHclStringFromGenericMap(builder, ociRes, interpolationMap)
}

var getHclStringForDnsResolverEndpoint = func(builder *strings.Builder, ociRes *tf_export.OCIResource, interpolationMap map[string]string) error {

	delete(interpolationMap, ociRes.Parent.Id)

	return tf_export.GetHclStringFromGenericMap(builder, ociRes, interpolationMap)
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDnsZoneHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dns_zone",
	DatasourceClass:        "oci_dns_zones",
	DatasourceItemsAttr:    "zones",
	ResourceAbbreviation:   "zone",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dns.ZoneLifecycleStateActive),
	},
}

var exportDnsSteeringPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dns_steering_policy",
	DatasourceClass:        "oci_dns_steering_policies",
	DatasourceItemsAttr:    "steering_policies",
	ResourceAbbreviation:   "steering_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dns.SteeringPolicyLifecycleStateActive),
	},
}

var exportDnsSteeringPolicyAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_dns_steering_policy_attachment",
	DatasourceClass:      "oci_dns_steering_policy_attachments",
	DatasourceItemsAttr:  "steering_policy_attachments",
	ResourceAbbreviation: "steering_policy_attachment",
	DiscoverableLifecycleStates: []string{
		string(oci_dns.SteeringPolicyAttachmentLifecycleStateActive),
	},
}

var exportDnsTsigKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dns_tsig_key",
	DatasourceClass:        "oci_dns_tsig_keys",
	DatasourceItemsAttr:    "tsig_keys",
	ResourceAbbreviation:   "tsig_key",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dns.TsigKeyLifecycleStateActive),
	},
}

var exportDnsRrsetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_dns_rrset",
	DatasourceClass:      "oci_dns_rrsets",
	DatasourceItemsAttr:  "rrsets",
	ResourceAbbreviation: "rrset",
}

var exportDnsResolverHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dns_resolver",
	DatasourceClass:        "oci_dns_resolvers",
	DatasourceItemsAttr:    "resolvers",
	ResourceAbbreviation:   "resolver",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dns.ResolverLifecycleStateActive),
	},
}

var exportDnsResolverEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dns_resolver_endpoint",
	DatasourceClass:        "oci_dns_resolver_endpoints",
	DatasourceItemsAttr:    "resolver_endpoints",
	ResourceAbbreviation:   "resolver_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dns.ResolverEndpointLifecycleStateActive),
	},
}

var exportDnsViewHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_dns_view",
	DatasourceClass:      "oci_dns_views",
	DatasourceItemsAttr:  "views",
	ResourceAbbreviation: "view",
	DiscoverableLifecycleStates: []string{
		string(oci_dns.ViewLifecycleStateActive),
	},
}

var dnsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDnsZoneHints},
		{TerraformResourceHints: exportDnsSteeringPolicyHints},
		{TerraformResourceHints: exportDnsSteeringPolicyAttachmentHints},
		{TerraformResourceHints: exportDnsTsigKeyHints},
		{TerraformResourceHints: exportDnsResolverHints},
		{TerraformResourceHints: exportDnsViewHints},
	},
	"oci_dns_zone": {
		{
			TerraformResourceHints: exportDnsRrsetHints,
			DatasourceQueryParams: map[string]string{
				"zone_name_or_id": "id",
			},
		},
	},
	"oci_dns_resolver": {
		{
			TerraformResourceHints: exportDnsResolverEndpointHints,
			DatasourceQueryParams: map[string]string{
				"resolver_id": "id",
			},
		},
	},
}
