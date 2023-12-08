package load_balancer

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportLoadBalancerBackendHints.GetIdFn = getLoadBalancerBackendId
	exportLoadBalancerBackendSetHints.GetIdFn = getLoadBalancerBackendSetId
	exportLoadBalancerCertificateHints.GetIdFn = getLoadBalancerCertificateId
	exportLoadBalancerHostnameHints.GetIdFn = getLoadBalancerHostnameId
	exportLoadBalancerListenerHints.GetIdFn = getLoadBalancerListenerId
	exportLoadBalancerPathRouteSetHints.GetIdFn = getLoadBalancerPathRouteSetId
	exportLoadBalancerLoadBalancerRoutingPolicyHints.GetIdFn = getLoadBalancerLoadBalancerRoutingPolicyId
	exportLoadBalancerRuleSetHints.GetIdFn = getLoadBalancerRuleSetId
	exportLoadBalancerSslCipherSuiteHints.GetIdFn = getLoadBalancerSslCipherSuiteId
	exportLoadBalancerBackendHints.ProcessDiscoveredResourcesFn = processLoadBalancerBackends
	exportLoadBalancerBackendSetHints.ProcessDiscoveredResourcesFn = processLoadBalancerBackendSets
	exportLoadBalancerCertificateHints.ProcessDiscoveredResourcesFn = processLoadBalancerCertificates
	exportLoadBalancerHostnameHints.ProcessDiscoveredResourcesFn = processLoadBalancerHostnames
	exportLoadBalancerListenerHints.FindResourcesOverrideFn = findLoadBalancerListeners
	exportLoadBalancerListenerHints.ProcessDiscoveredResourcesFn = processLoadBalancerListeners
	exportLoadBalancerPathRouteSetHints.ProcessDiscoveredResourcesFn = processLoadBalancerPathRouteSets
	exportLoadBalancerRuleSetHints.ProcessDiscoveredResourcesFn = processLoadBalancerRuleSets
	exportLoadBalancerLoadBalancerRoutingPolicyHints.ProcessDiscoveredResourcesFn = processLoadBalancerRoutingPolicies
	tf_export.RegisterCompartmentGraphs("load_balancer", loadBalancerResourceGraph)
	tf_export.RegisterRelatedResourcesGraph("oci_load_balancer_backend_set", relatedLoadBalancerBackendSet)
	tf_export.RegisterRelatedResourcesGraph("oci_load_balancer_load_balancer", relatedLoadBalancerLoadBalancer)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processLoadBalancerListeners(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {

	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		if sslConfiguration, ok := resource.SourceAttributes["ssl_configuration"].([]interface{}); ok && len(sslConfiguration) > 0 {
			if sslConfig, ok := sslConfiguration[0].(map[string]interface{}); ok {
				if certificateName, ok := sslConfig["certificate_name"]; ok {
					// check if we have expected ResourceIds set, is load balancer certificate id expected
					if ctx.ExpectedResourceIds != nil && len(ctx.ExpectedResourceIds) > 0 {
						certificateId := GetCertificateCompositeId(certificateName.(string), resource.SourceAttributes["load_balancer_id"].(string))
						if _, ok = ctx.ExpectedResourceIds[certificateId]; !ok {
							continue
						}
					}
					sslConfig["certificate_name"] = tf_export.InterpolationString{
						ResourceReference: resource.Parent.GetTerraformReference(),
						Interpolation:     tf_export.LoadBalancerCertificateNameMap[resource.Parent.Parent.Id][sslConfig["certificate_name"].(string)],
						Value:             sslConfig["certificate_name"].(string),
					}
				}
			}
		}
	}
	return resources, nil
}

func findLoadBalancerListeners(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	loadBalancerId := parent.SourceAttributes["load_balancer_id"].(string)
	backendSetName := parent.SourceAttributes["name"].(string)

	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

	response, err := ctx.Clients.LoadBalancerClient().GetLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}

	listenerResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

	results := []*tf_export.OCIResource{}
	for listenerName, listener := range response.LoadBalancer.Listeners {
		if *listener.DefaultBackendSetName != backendSetName {
			continue
		}

		d := listenerResource.TestResourceData()
		d.SetId(GetListenerCompositeId(listenerName, loadBalancerId))

		// This calls into the listener resource's Read fn which has the unfortunate implementation of
		// calling GetLoadBalancer and looping through the listeners to find the expected one. So this entire method
		// may require O(n^^2) time. However, the benefits of having Read populate the ResourceData struct is better than duplicating it here.
		if err := listenerResource.Read(d, ctx.Clients); err != nil {
			// add error to the errorList and continue discovering rest of the resources
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(listenerResource.Schema, d),
			RawResource:      listener,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
				TerraformName:  fmt.Sprintf("%s_%s", parent.Parent.TerraformName, listenerName),
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if !parent.OmitFromExport {
			resource.SourceAttributes["default_backend_set_name"] = tf_export.InterpolationString{
				ResourceReference: parent.GetTerraformReference(),
				Interpolation:     tf_export.TfHclVersionvar.GetDoubleExpHclString(parent.GetTerraformReference(), "name"),
				Value:             parent.SourceAttributes["name"].(string),
			}
		} else {
			resource.SourceAttributes["default_backend_set_name"] = parent.SourceAttributes["name"].(string)
		}
		resource.TerraformName = tf_export.GetValidUniqueTerraformName(resource.TerraformName)
		results = append(results, resource)
	}

	return results, nil
}

func processLoadBalancerBackendSets(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, backendSet := range resources {
		if backendSet.Parent == nil {
			continue
		}

		backendSetName := backendSet.SourceAttributes["name"].(string)
		backendSet.Id = GetBackendSetCompositeId(backendSetName, backendSet.Parent.Id)
		backendSet.SourceAttributes["load_balancer_id"] = backendSet.Parent.Id
	}

	return resources, nil
}

func processLoadBalancerBackends(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, backend := range resources {
		if backend.Parent == nil {
			continue
		}

		backend.Id = GetBackendCompositeId(backend.SourceAttributes["name"].(string), backend.Parent.SourceAttributes["name"].(string), backend.Parent.SourceAttributes["load_balancer_id"].(string))
		backend.SourceAttributes["load_balancer_id"] = backend.Parent.SourceAttributes["load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !backend.Parent.OmitFromExport {
			backend.SourceAttributes["backendset_name"] = tf_export.InterpolationString{
				ResourceReference: backend.Parent.GetTerraformReference(),
				Interpolation:     tf_export.TfHclVersionvar.GetDoubleExpHclString(backend.Parent.GetTerraformReference(), "name"),
				Value:             backend.Parent.SourceAttributes["name"].(string),
			}
		} else {
			backend.SourceAttributes["backendset_name"] = backend.Parent.SourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func processLoadBalancerHostnames(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, hostname := range resources {
		if hostname.Parent == nil {
			continue
		}

		hostname.Id = GetHostnameCompositeId(hostname.Parent.Id, hostname.SourceAttributes["name"].(string))
		hostname.SourceAttributes["load_balancer_id"] = hostname.Parent.Id
	}

	return resources, nil
}

func processLoadBalancerPathRouteSets(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, pathRouteSet := range resources {
		if pathRouteSet.Parent == nil {
			continue
		}

		pathRouteSet.Id = GetPathRouteSetCompositeId(pathRouteSet.Parent.Id, pathRouteSet.SourceAttributes["name"].(string))
		pathRouteSet.SourceAttributes["load_balancer_id"] = pathRouteSet.Parent.Id
	}

	return resources, nil
}

func processLoadBalancerRoutingPolicies(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, routingPolicy := range resources {
		if routingPolicy.Parent == nil {
			continue
		}

		routingPolicy.Id = GetLoadBalancerRoutingPolicyCompositeId(routingPolicy.Parent.Id, routingPolicy.SourceAttributes["name"].(string))
		routingPolicy.SourceAttributes["load_balancer_id"] = routingPolicy.Parent.Id
	}

	return resources, nil
}

func processLoadBalancerRuleSets(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, ruleSet := range resources {
		if ruleSet.Parent == nil {
			continue
		}

		ruleSet.Id = GetRuleSetCompositeId(ruleSet.Parent.Id, ruleSet.SourceAttributes["name"].(string))
		ruleSet.SourceAttributes["load_balancer_id"] = ruleSet.Parent.Id
	}

	return resources, nil
}

func processLoadBalancerCertificates(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, certificate := range resources {
		if certificate.Parent == nil {
			continue
		}

		certificate.Id = GetCertificateCompositeId(certificate.SourceAttributes["certificate_name"].(string), certificate.Parent.Id)
		certificate.SourceAttributes["load_balancer_id"] = certificate.Parent.Id

		// add certificate name and interpolation to loadBalancerCertificateNameMap
		if tf_export.LoadBalancerCertificateNameMap == nil {
			tf_export.LoadBalancerCertificateNameMap = make(map[string]map[string]string)
		}
		_, ok := tf_export.LoadBalancerCertificateNameMap[certificate.Parent.Id]
		if !ok {
			tf_export.LoadBalancerCertificateNameMap[certificate.Parent.Id] = make(map[string]string)
		}

		if certificateName, ok := certificate.SourceAttributes["certificate_name"].(string); ok {
			tf_export.LoadBalancerCertificateNameMap[certificate.Parent.Id][certificateName] = tf_export.TfHclVersionvar.GetDoubleExpHclString(certificate.GetTerraformReference(), "certificate_name")
		}
	}

	return resources, nil
}

func getLoadBalancerBackendId(resource *tf_export.OCIResource) (string, error) {

	backendName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendName for LoadBalancer Backend")
	}
	backendsetName, ok := resource.SourceAttributes["backendset_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendsetName for LoadBalancer Backend")
	}
	loadBalancerId, ok := resource.SourceAttributes["load_balancer_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find loadBalancerId for LoadBalancer Backend")
	}
	return GetBackendCompositeId(backendName, backendsetName, loadBalancerId), nil
}

func getLoadBalancerBackendSetId(resource *tf_export.OCIResource) (string, error) {

	backendSetName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for LoadBalancer BackendSet")
	}
	loadBalancerId := resource.Parent.Id
	return GetBackendSetCompositeId(backendSetName, loadBalancerId), nil
}

func getLoadBalancerCertificateId(resource *tf_export.OCIResource) (string, error) {

	certificateName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find certificateName for LoadBalancer Certificate")
	}
	loadBalancerId := resource.Parent.Id
	return GetCertificateCompositeId(certificateName, loadBalancerId), nil
}

func getLoadBalancerHostnameId(resource *tf_export.OCIResource) (string, error) {

	loadBalancerId := resource.Parent.Id
	name, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer Hostname")
	}
	return GetHostnameCompositeId(loadBalancerId, name), nil
}

func getLoadBalancerListenerId(resource *tf_export.OCIResource) (string, error) {

	listenerName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for LoadBalancer Listener")
	}
	loadBalancerId, ok := resource.SourceAttributes["load_balancer_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find loadBalancerId for LoadBalancer Listener")
	}
	return GetListenerCompositeId(listenerName, loadBalancerId), nil
}

func getLoadBalancerPathRouteSetId(resource *tf_export.OCIResource) (string, error) {

	loadBalancerId := resource.Parent.Id
	pathRouteSetName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find pathRouteSetName for LoadBalancer PathRouteSet")
	}
	return GetPathRouteSetCompositeId(loadBalancerId, pathRouteSetName), nil
}

func getLoadBalancerLoadBalancerRoutingPolicyId(resource *tf_export.OCIResource) (string, error) {

	loadBalancerId := resource.Parent.Id
	routingPolicyName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find routingPolicyName for LoadBalancer LoadBalancerRoutingPolicy")
	}
	return GetLoadBalancerRoutingPolicyCompositeId(loadBalancerId, routingPolicyName), nil
}

func getLoadBalancerRuleSetId(resource *tf_export.OCIResource) (string, error) {

	loadBalancerId := resource.Parent.Id
	name, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer RuleSet")
	}
	return GetRuleSetCompositeId(loadBalancerId, name), nil
}

func getLoadBalancerSslCipherSuiteId(resource *tf_export.OCIResource) (string, error) {

	loadBalancerId := resource.Parent.Id
	name, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find name for LoadBalancer SslCipherSuite")
	}
	return getSslCipherSuiteCompositeId(loadBalancerId, name), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportLoadBalancerBackendHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_backend",
	DatasourceClass:      "oci_load_balancer_backends",
	DatasourceItemsAttr:  "backends",
	ResourceAbbreviation: "backend",
}

var exportLoadBalancerBackendSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_backend_set",
	DatasourceClass:      "oci_load_balancer_backend_sets",
	DatasourceItemsAttr:  "backendsets",
	ResourceAbbreviation: "backend_set",
}

var exportLoadBalancerCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_certificate",
	DatasourceClass:      "oci_load_balancer_certificates",
	DatasourceItemsAttr:  "certificates",
	ResourceAbbreviation: "certificate",
}

var exportLoadBalancerHostnameHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_hostname",
	DatasourceClass:      "oci_load_balancer_hostnames",
	DatasourceItemsAttr:  "hostnames",
	ResourceAbbreviation: "hostname",
}

var exportLoadBalancerListenerHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_listener",
	ResourceAbbreviation: "listener",
}

var exportLoadBalancerLoadBalancerHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_load_balancer",
	DatasourceClass:      "oci_load_balancer_load_balancers",
	DatasourceItemsAttr:  "load_balancers",
	ResourceAbbreviation: "load_balancer",
	DiscoverableLifecycleStates: []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateActive),
	},
}

var exportLoadBalancerPathRouteSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_path_route_set",
	DatasourceClass:      "oci_load_balancer_path_route_sets",
	DatasourceItemsAttr:  "path_route_sets",
	ResourceAbbreviation: "path_route_set",
}

var exportLoadBalancerLoadBalancerRoutingPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_load_balancer_routing_policy",
	DatasourceClass:      "oci_load_balancer_load_balancer_routing_policies",
	DatasourceItemsAttr:  "routing_policies",
	ResourceAbbreviation: "load_balancer_routing_policy",
}

var exportLoadBalancerRuleSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_rule_set",
	DatasourceClass:      "oci_load_balancer_rule_sets",
	DatasourceItemsAttr:  "rule_sets",
	ResourceAbbreviation: "rule_set",
}

var exportLoadBalancerSslCipherSuiteHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_load_balancer_ssl_cipher_suite",
	DatasourceClass:      "oci_load_balancer_ssl_cipher_suites",
	DatasourceItemsAttr:  "ssl_cipher_suites",
	ResourceAbbreviation: "ssl_cipher_suite",
}

var loadBalancerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLoadBalancerLoadBalancerHints},
	},
	"oci_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportLoadBalancerBackendHints,
			DatasourceQueryParams: map[string]string{
				"backendset_name":  "name",
				"load_balancer_id": "load_balancer_id",
			},
		},
		{TerraformResourceHints: exportLoadBalancerListenerHints},
	},
	"oci_load_balancer_load_balancer": {
		// certificates have to be discovered before listeners in order to populate
		// the references for certificate_name in listeners (dependency)
		// If moving to parallel execution in future, this dependency needs to be maintained
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerHostnameHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerLoadBalancerRoutingPolicyHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerPathRouteSetHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerRuleSetHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerSslCipherSuiteHints,
			DatasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
	},
}

var relatedLoadBalancerBackendSet = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportLoadBalancerBackendHints,
		DatasourceQueryParams: map[string]string{
			"backendset_name":  "name",
			"load_balancer_id": "load_balancer_id",
		},
	},
	{TerraformResourceHints: exportLoadBalancerListenerHints},
}

var relatedLoadBalancerLoadBalancer = []tf_export.TerraformResourceAssociation{
	// certificates have to be discovered before listeners in order to populate
	// the references for certificate_name in listeners (dependency)
	// If moving to parallel execution in future, this dependency needs to be maintained
	{
		TerraformResourceHints: exportLoadBalancerCertificateHints,
		DatasourceQueryParams: map[string]string{
			"load_balancer_id": "id",
		},
	},
	{
		TerraformResourceHints: exportLoadBalancerBackendSetHints,
		DatasourceQueryParams: map[string]string{
			"load_balancer_id": "id",
		},
	},
}
