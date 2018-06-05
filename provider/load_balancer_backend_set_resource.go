// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func BackendSetResource() *schema.Resource {
	return &schema.Resource{
		Create: createBackendSet,
		Read:   readBackendSet,
		Update: updateBackendSet,
		Delete: deleteBackendSet,
		Schema: map[string]*schema.Schema{
			// Required
			"health_checker": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"interval_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  30000,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"response_body_regex": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"retries": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3,
						},
						"return_code": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"timeout_in_millis": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3000,
						},
						"url_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"policy": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Optional
			"backend": {
				Type: schema.TypeList,
				//Optional: true, // @CODEGEN Having 2 ways to specify backends (this and backend resource) is bad because they will override each other. Leaving computed for now.
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"backup": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"drain": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"offline": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"session_persistence_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cookie_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"disable_fallback": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},

						// Computed
					},
				},
			},
			"ssl_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"certificate_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
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

func createBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updateBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type BackendSetResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.BackendSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

// The oci_loadbalancer_backend resource may implicitly modify this backend set and this could happen concurrently.
// Use a per-backend set mutex to synchronize accesses to the backend set.
func (s *BackendSetResourceCrud) GetMutex() *sync.Mutex {
	return lbBackendSetMutexes.GetOrCreateBackendSetMutex(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
}

func (s *BackendSetResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *BackendSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendSetResourceCrud) Create() error {
	request := oci_load_balancer.CreateBackendSetRequest{}

	/*  // @CODEGEN Having 2 ways to specify backends (this and backend resource) is bad because they will override each other. Leaving computed for now.
	request.Backends = []oci_load_balancer.BackendDetails{}
	if backend, ok := s.D.GetOkExists("backend"); ok {
		interfaces := backend.([]interface{})
		tmp := make([]oci_load_balancer.BackendDetails, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToBackendDetails(toBeConverted.(map[string]interface{}))
		}
		request.Backends = tmp
	}
	*/

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			tmp := mapToHealthCheckerDetails(tmpList[0].(map[string]interface{}))
			request.HealthChecker = &tmp
		}
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		tmp := policy.(string)
		request.Policy = &tmp
	}

	if sessionPersistenceConfiguration, ok := s.D.GetOkExists("session_persistence_configuration"); ok {
		if tmpList := sessionPersistenceConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSessionPersistenceConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SessionPersistenceConfiguration = &tmp
		}
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSSLConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *BackendSetResourceCrud) Get() error {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackendSet
	return nil
}

func (s *BackendSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdateBackendSetRequest{}

	/*  // @CODEGEN Having 2 ways to specify backends (this and backend resource) is bad because they will override each other. Reverting to old logic.
	request.Backends = []oci_load_balancer.BackendDetails{}
	if backend, ok := s.D.GetOkExists("backend"); ok {
		interfaces := backend.([]interface{})
		tmp := make([]oci_load_balancer.BackendDetails, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToBackendDetails(toBeConverted.(map[string]interface{}))
		}
		request.Backends = tmp
	}
	*/
	// This is hacky and a race condition, but works for now. Ideally backends are not a required parameter to a backendset update
	err := s.Get()
	if err != nil {
		return err
	}
	request.Backends = backendArrayToBackendDetailsArray(s.Res.Backends)

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			tmp := mapToHealthCheckerDetails(tmpList[0].(map[string]interface{}))
			request.HealthChecker = &tmp
		}
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		tmp := policy.(string)
		request.Policy = &tmp
	}

	if sessionPersistenceConfiguration, ok := s.D.GetOkExists("session_persistence_configuration"); ok {
		if tmpList := sessionPersistenceConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSessionPersistenceConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SessionPersistenceConfiguration = &tmp
		}
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSSLConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *BackendSetResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteBackendSet(context.Background(), request)

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *BackendSetResourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	backend := []interface{}{}
	for _, item := range s.Res.Backends {
		backend = append(backend, BackendToMap(item))
	}
	s.D.Set("backend", backend)

	if s.Res.HealthChecker != nil {
		s.D.Set("health_checker", []interface{}{HealthCheckerToMap(s.Res.HealthChecker)})
	} else {
		s.D.Set("health_checker", []interface{}{})
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Policy != nil {
		s.D.Set("policy", *s.Res.Policy)
	}

	if s.Res.SessionPersistenceConfiguration != nil {
		s.D.Set("session_persistence_configuration", []interface{}{SessionPersistenceConfigurationDetailsToMap(s.Res.SessionPersistenceConfiguration)})
	} else {
		s.D.Set("session_persistence_configuration", []interface{}{})
	}

	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", []interface{}{SSLConfigurationToMap(s.Res.SslConfiguration)})
	} else {
		s.D.Set("ssl_configuration", []interface{}{})
	}

}

func BackendToMap(obj oci_load_balancer.Backend) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Backup != nil {
		result["backup"] = bool(*obj.Backup)
	}

	if obj.Drain != nil {
		result["drain"] = bool(*obj.Drain)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Offline != nil {
		result["offline"] = bool(*obj.Offline)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func mapToHealthCheckerDetails(raw map[string]interface{}) oci_load_balancer.HealthCheckerDetails {
	result := oci_load_balancer.HealthCheckerDetails{}

	if intervalMs, ok := raw["interval_ms"]; ok {
		tmp := intervalMs.(int)
		result.IntervalInMillis = &tmp
	}

	if port, ok := raw["port"]; ok {
		tmp := port.(int)
		if tmp != 0 {
			result.Port = &tmp
		}
	}

	if protocol, ok := raw["protocol"]; ok && protocol != "" {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if responseBodyRegex, ok := raw["response_body_regex"]; ok && responseBodyRegex != "" {
		tmp := responseBodyRegex.(string)
		result.ResponseBodyRegex = &tmp
	}

	if retries, ok := raw["retries"]; ok {
		tmp := retries.(int)
		result.Retries = &tmp
	}

	if returnCode, ok := raw["return_code"]; ok {
		tmp := returnCode.(int)
		if tmp != 0 {
			result.ReturnCode = &tmp
		}
	}

	if timeoutInMillis, ok := raw["timeout_in_millis"]; ok {
		tmp := timeoutInMillis.(int)
		result.TimeoutInMillis = &tmp
	}

	if urlPath, ok := raw["url_path"]; ok && urlPath != "" {
		tmp := urlPath.(string)
		result.UrlPath = &tmp
	}

	return result
}

func HealthCheckerToMap(obj *oci_load_balancer.HealthChecker) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IntervalInMillis != nil {
		result["interval_ms"] = int(*obj.IntervalInMillis)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	if obj.ResponseBodyRegex != nil {
		result["response_body_regex"] = string(*obj.ResponseBodyRegex)
	}

	if obj.Retries != nil {
		result["retries"] = int(*obj.Retries)
	}

	if obj.ReturnCode != nil {
		result["return_code"] = int(*obj.ReturnCode)
	}

	if obj.TimeoutInMillis != nil {
		result["timeout_in_millis"] = int(*obj.TimeoutInMillis)
	}

	if obj.UrlPath != nil {
		result["url_path"] = string(*obj.UrlPath)
	}

	return result
}

func mapToSSLConfigurationDetails(raw map[string]interface{}) oci_load_balancer.SslConfigurationDetails {
	result := oci_load_balancer.SslConfigurationDetails{}

	if certificateName, ok := raw["certificate_name"]; ok && certificateName != "" {
		tmp := certificateName.(string)
		result.CertificateName = &tmp
	}

	if verifyDepth, ok := raw["verify_depth"]; ok {
		tmp := verifyDepth.(int)
		result.VerifyDepth = &tmp
	}

	if verifyPeerCertificate, ok := raw["verify_peer_certificate"]; ok {
		tmp := verifyPeerCertificate.(bool)
		result.VerifyPeerCertificate = &tmp
	}

	return result
}

func SSLConfigurationToMap(obj *oci_load_balancer.SslConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateName != nil {
		result["certificate_name"] = string(*obj.CertificateName)
	}

	if obj.VerifyDepth != nil {
		result["verify_depth"] = int(*obj.VerifyDepth)
	}

	if obj.VerifyPeerCertificate != nil {
		result["verify_peer_certificate"] = bool(*obj.VerifyPeerCertificate)
	}

	return result
}

func mapToSessionPersistenceConfigurationDetails(raw map[string]interface{}) oci_load_balancer.SessionPersistenceConfigurationDetails {
	result := oci_load_balancer.SessionPersistenceConfigurationDetails{}

	if cookieName, ok := raw["cookie_name"]; ok && cookieName != "" {
		tmp := cookieName.(string)
		result.CookieName = &tmp
	}

	if disableFallback, ok := raw["disable_fallback"]; ok {
		tmp := disableFallback.(bool)
		result.DisableFallback = &tmp
	}

	return result
}

func SessionPersistenceConfigurationDetailsToMap(obj *oci_load_balancer.SessionPersistenceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CookieName != nil {
		result["cookie_name"] = string(*obj.CookieName)
	}

	if obj.DisableFallback != nil {
		result["disable_fallback"] = bool(*obj.DisableFallback)
	}

	return result
}

func backendArrayToBackendDetailsArray(backends []oci_load_balancer.Backend) []oci_load_balancer.BackendDetails {
	backendDetailsArr := make([]oci_load_balancer.BackendDetails, len(backends))
	for i, backend := range backends {
		backendDetailsArr[i] = backendToBackendDetails(backend)
	}
	return backendDetailsArr
}

func backendToBackendDetails(backend oci_load_balancer.Backend) oci_load_balancer.BackendDetails {
	result := oci_load_balancer.BackendDetails{}

	result.Backup = backend.Backup
	result.Drain = backend.Drain
	result.IpAddress = backend.IpAddress
	result.Offline = backend.Offline
	result.Port = backend.Port
	result.Weight = backend.Weight

	return result
}
