// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

func LoadBalancerListenerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerListener,
		Read:     readLoadBalancerListener,
		Update:   updateLoadBalancerListener,
		Delete:   deleteLoadBalancerListener,
		Schema: map[string]*schema.Schema{
			// Required
			"default_backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"connection_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idle_timeout_in_seconds": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Optional
						"backend_tcp_proxy_protocol_version": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"hostname_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"path_route_set_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_set_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ssl_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Optional
						"certificate_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"certificate_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cipher_suite_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"has_session_resumption": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"protocols": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"server_order_preference": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trusted_certificate_authority_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"verify_depth": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
						"verify_peer_certificate": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						// Computed
					},
				},
			},

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	return tfresource.ReadResource(sync)
}

func updateLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerListenerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Listener
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerListenerResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return GetListenerCompositeId(s.D.Get("name").(string), s.D.Get("load_balancer_id").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerListenerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerListenerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerListenerResourceCrud) Create() error {
	request := oci_load_balancer.CreateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_configuration", 0)
			tmp, err := s.mapToConnectionConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	// patch for backward compatibility
	request.HostnameNames = []string{}
	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.HostnameNames = tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("routing_policy_name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	if ruleSetNames, ok := s.D.GetOkExists("rule_set_names"); ok {
		interfaces := ruleSetNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("rule_set_names") {
			request.RuleSetNames = tmp
		}
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssl_configuration", 0)
			tmp, err := s.mapToSSLConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateListener(context.Background(), request)
	if err != nil {
		return err
	}

	var compositeId string
	compositeId = GetListenerCompositeId(s.D.Get("name").(string), s.D.Get("load_balancer_id").(string))
	s.D.SetId(compositeId)
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerListenerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID,name}
	_, stillWorking, err := loadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		listenerName, loadBalancerId, err := parseListenerCompositeId(s.D.Id())
		if err == nil {
			s.D.Set("name", &listenerName)
			s.D.Set("load_balancer_id", &loadBalancerId)
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	res, e := s.GetListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	if e == nil {
		s.Res = res
	}
	return
}

func (s *LoadBalancerListenerResourceCrud) GetListener(loadBalancerID, name string) (*oci_load_balancer.Listener, error) {
	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerID
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}
	lb := &response.LoadBalancer
	if lb != nil && lb.Listeners != nil {
		if l, ok := lb.Listeners[name]; ok {
			if l.Name != nil && *l.Name == name {
				return &l, nil
			}
		}
	}
	return nil, fmt.Errorf("Listener %s on load balancer %s does not exist", name, loadBalancerID)
}

func (s *LoadBalancerListenerResourceCrud) Update() error {
	request := oci_load_balancer.UpdateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_configuration", 0)
			tmp, err := s.mapToConnectionConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("hostname_names") {
			request.HostnameNames = tmp
		}
	}

	if listenerName, ok := s.D.GetOkExists("name"); ok {
		tmp := listenerName.(string)
		request.ListenerName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("routing_policy_name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	if ruleSetNames, ok := s.D.GetOkExists("rule_set_names"); ok {
		interfaces := ruleSetNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("rule_set_names") {
			request.RuleSetNames = tmp
		}
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssl_configuration", 0)
			tmp, err := s.mapToSSLConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateListener(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerListenerResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteListenerRequest{}

	if listenerName, ok := s.D.GetOkExists("name"); ok {
		tmp := listenerName.(string)
		request.ListenerName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteListener(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerListenerResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	listenerName, loadBalancerId, err := parseListenerCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &listenerName)
		s.D.Set("load_balancer_id", &loadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ConnectionConfiguration != nil {
		s.D.Set("connection_configuration", []interface{}{ConnectionConfigurationToMap(s.Res.ConnectionConfiguration)})
	} else {
		s.D.Set("connection_configuration", []interface{}{})
	}
	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
	}
	if s.Res.HostnameNames != nil {
		s.D.Set("hostname_names", s.Res.HostnameNames)
	} else {
		s.D.Set("hostname_names", []interface{}{})
	}
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}
	if s.Res.PathRouteSetName != nil {
		s.D.Set("path_route_set_name", *s.Res.PathRouteSetName)
	}
	if s.Res.RuleSetNames != nil {
		s.D.Set("rule_set_names", s.Res.RuleSetNames)
	}
	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}
	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}
	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", []interface{}{SSLConfigurationToMap(s.Res.SslConfiguration)})
	} else {
		s.D.Set("ssl_configuration", []interface{}{})
	}

	return nil
}

func GetListenerCompositeId(listenerName string, loadBalancerId string) string {
	listenerName = url.PathEscape(listenerName)
	loadBalancerId = url.PathEscape(loadBalancerId)
	compositeId := "loadBalancers/" + loadBalancerId + "/listeners/" + listenerName
	return compositeId
}

func parseListenerCompositeId(compositeId string) (listenerName string, loadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/listeners/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	listenerName, _ = url.PathUnescape(parts[3])

	return
}

func (s *LoadBalancerListenerResourceCrud) mapToConnectionConfiguration(fieldKeyFormat string) (oci_load_balancer.ConnectionConfiguration, error) {
	result := oci_load_balancer.ConnectionConfiguration{}

	if backendTcpProxyProtocolVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_tcp_proxy_protocol_version")); ok {
		tmp := backendTcpProxyProtocolVersion.(int)
		// Terraform v11 will auto assign nil value to 0 which is invalid value
		// this check will remove backend_tcp_proxy_protocol_version in the request
		if tmp != 0 {
			result.BackendTcpProxyProtocolVersion = &tmp
		}
	}

	if idleTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idle_timeout_in_seconds")); ok {
		tmp := idleTimeoutInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert idleTimeoutInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.IdleTimeout = &tmpInt64
	}

	return result, nil
}

func ConnectionConfigurationToMap(obj *oci_load_balancer.ConnectionConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackendTcpProxyProtocolVersion != nil {
		result["backend_tcp_proxy_protocol_version"] = int(*obj.BackendTcpProxyProtocolVersion)
	}

	if obj.IdleTimeout != nil {
		result["idle_timeout_in_seconds"] = strconv.FormatInt(*obj.IdleTimeout, 10)
	}

	return result
}

func (s *LoadBalancerListenerResourceCrud) mapToSSLConfigurationDetails(fieldKeyFormat string) (oci_load_balancer.SslConfigurationDetails, error) {
	result := oci_load_balancer.SslConfigurationDetails{}

	if certificateIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_ids")); ok {
		interfaces := certificateIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "certificate_ids")) {
			result.CertificateIds = tmp
		}
	}

	if certificateName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_name")); ok {
		tmp := certificateName.(string)
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "certificate_name")) {
			if tmp == "" {
				result.CertificateName = nil
			} else {
				result.CertificateName = &tmp
			}
		}
	}

	if cipherSuiteName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cipher_suite_name")); ok {
		tmp := cipherSuiteName.(string)
		result.CipherSuiteName = &tmp
	}

	if protocols, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocols")); ok {
		interfaces := protocols.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = fmt.Sprintf("%s", interfaces[i])
		}
		result.Protocols = tmp
	}

	if serverOrderPreference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_order_preference")); ok {
		result.ServerOrderPreference = oci_load_balancer.SslConfigurationDetailsServerOrderPreferenceEnum(serverOrderPreference.(string))
	}

	if hasSessionResumption, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "has_session_resumption")); ok {
		tmp := hasSessionResumption.(bool)
		result.HasSessionResumption = &tmp
	}

	if trustedCertificateAuthorityIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trusted_certificate_authority_ids")); ok {
		interfaces := trustedCertificateAuthorityIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "trusted_certificate_authority_ids")) {
			result.TrustedCertificateAuthorityIds = tmp
		}
	}

	if verifyDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_depth")); ok {
		tmp := verifyDepth.(int)
		result.VerifyDepth = &tmp
	}

	if verifyPeerCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_peer_certificate")); ok {
		tmp := verifyPeerCertificate.(bool)
		result.VerifyPeerCertificate = &tmp
	}

	return result, nil
}

// @CODEGEN 08/2018 - Method SSLConfigurationDetailsToMap is available in load_balancer_backend_set_resource.go
