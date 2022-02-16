// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"
)

func LoadBalancerBackendSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerBackendSet,
		Read:     readLoadBalancerBackendSet,
		Update:   updateLoadBalancerBackendSet,
		Delete:   deleteLoadBalancerBackendSet,
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
				Required: true,
			},

			// Optional
			"lb_cookie_session_persistence_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cookie_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disable_fallback": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"domain": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: utils.ValidateNotEmptyString(),
						},
						"is_http_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_secure": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"max_age_in_seconds": {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
				// @CODEGEN: lb_cookie_session_persistence_configuration and session_persistence_configuration are mutually exclusive
				ConflictsWith: []string{"session_persistence_configuration"},
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
				// @CODEGEN: lb_cookie_session_persistence_configuration and session_persistence_configuration are mutually exclusive
				ConflictsWith: []string{"lb_cookie_session_persistence_configuration"},
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
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"certificate_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cipher_suite_name": {
							Type:     schema.TypeString,
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
							Computed: true,
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
			"backend": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      backendHashCodeForSets,
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
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerBackendSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.BackendSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

// The oci_loadbalancer_backend resource may implicitly modify this backend set and this could happen concurrently.
// Use a per-backend set mutex to synchronize accesses to the backend set.
func (s *LoadBalancerBackendSetResourceCrud) GetMutex() *sync.Mutex {
	return lbBackendSetMutexes.GetOrCreateBackendSetMutex(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
}

func (s *LoadBalancerBackendSetResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return GetBackendSetCompositeId(s.D.Get("name").(string), s.D.Get("load_balancer_id").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerBackendSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerBackendSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerBackendSetResourceCrud) Create() error {
	request := oci_load_balancer.CreateBackendSetRequest{}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HealthChecker = &tmp
		}
	}

	if lbCookieSessionPersistenceConfiguration, ok := s.D.GetOkExists("lb_cookie_session_persistence_configuration"); ok {
		if tmpList := lbCookieSessionPersistenceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lb_cookie_session_persistence_configuration", 0)
			tmp, err := s.mapToLBCookieSessionPersistenceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LbCookieSessionPersistenceConfiguration = &tmp
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
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "session_persistence_configuration", 0)
			tmp, err := s.mapToSessionPersistenceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SessionPersistenceConfiguration = &tmp
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

	response, err := s.Client.CreateBackendSet(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerBackendSetResourceCrud) Get() error {
	_, stillWorking, err := tfresource.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		backendSetName, loadBalancerId, err := parseBackendSetCompositeId(s.D.Id())
		if err == nil {
			request.BackendSetName = &backendSetName
			request.LoadBalancerId = &loadBalancerId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackendSet
	return nil
}

func (s *LoadBalancerBackendSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdateBackendSetRequest{}

	// @CODEGEN: Backends are marked computed in this resource, so will do a GET and include the results in the UPDATE, although they are not a required parameter
	// Side-note: There is a potential for a race condition if the backend are added at the same time outside Terraform
	err := s.Get()
	if err != nil {
		return err
	}

	backends := []interface{}{}
	for _, item := range s.Res.Backends {
		backends = append(backends, BackendToMap(item))
	}

	set := schema.NewSet(backendHashCodeForSets, backends)
	interfaces := set.List()
	tmp := make([]oci_load_balancer.BackendDetails, len(interfaces))
	for i := range interfaces {
		stateDataIndex := backendHashCodeForSets(interfaces[i])
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backend", stateDataIndex)
		converted, err := s.mapToBackendDetails(fieldKeyFormat)
		if err != nil {
			return err
		}
		tmp[i] = converted
	}
	request.Backends = tmp

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
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

	//@CODEGEN: Since lbCookieSessionPersistenceConfiguration and sessionPersistenceConfiguration are mutually exclusive,
	// when migrating from one persistence configuration to another we want to pick only the change coming from the config
	// For lists HasChange returns false if you remove the list block from config
	lbCookieSessionPersistenceConfigurationChanged := false
	lbCookieSessionPersistenceConfiguration, lbCookieSessionPersistenceConfigurationPresent := s.D.GetOkExists("lb_cookie_session_persistence_configuration")
	if lbCookieSessionPersistenceConfigurationPresent && s.D.HasChange("lb_cookie_session_persistence_configuration") {
		lbCookieSessionPersistenceConfigurationChanged = true
	}

	sessionPersistenceConfigurationChanged := false
	sessionPersistenceConfiguration, sessionPersistenceConfigurationPresent := s.D.GetOkExists("session_persistence_configuration")
	if sessionPersistenceConfigurationPresent && s.D.HasChange("session_persistence_configuration") {
		sessionPersistenceConfigurationChanged = true
	}

	if !sessionPersistenceConfigurationChanged && lbCookieSessionPersistenceConfigurationPresent && lbCookieSessionPersistenceConfiguration != nil {
		if tmpList := lbCookieSessionPersistenceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lb_cookie_session_persistence_configuration", 0)
			tmp, err := s.mapToLBCookieSessionPersistenceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LbCookieSessionPersistenceConfiguration = &tmp
		}
	}

	if !lbCookieSessionPersistenceConfigurationChanged && sessionPersistenceConfigurationPresent && sessionPersistenceConfiguration != nil {
		if tmpList := sessionPersistenceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "session_persistence_configuration", 0)
			tmp, err := s.mapToSessionPersistenceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SessionPersistenceConfiguration = &tmp
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

	response, err := s.Client.UpdateBackendSet(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerBackendSetResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteBackendSet(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerBackendSetResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	backendSetName, loadBalancerId, err := parseBackendSetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &backendSetName)
		s.D.Set("load_balancer_id", &loadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	backend := []interface{}{}
	for _, item := range s.Res.Backends {
		backend = append(backend, BackendToMap(item))
	}
	s.D.Set("backend", schema.NewSet(backendHashCodeForSets, backend))

	if s.Res.HealthChecker != nil {
		s.D.Set("health_checker", []interface{}{HealthCheckerToMap(s.Res.HealthChecker)})
	} else {
		s.D.Set("health_checker", nil)
	}

	if s.Res.LbCookieSessionPersistenceConfiguration != nil {
		s.D.Set("lb_cookie_session_persistence_configuration", []interface{}{LBCookieSessionPersistenceConfigurationDetailsToMap(s.Res.LbCookieSessionPersistenceConfiguration)})
	} else {
		s.D.Set("lb_cookie_session_persistence_configuration", nil)
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
		s.D.Set("session_persistence_configuration", nil)
	}

	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", []interface{}{SSLConfigurationToMap(s.Res.SslConfiguration)})
	} else {
		s.D.Set("ssl_configuration", nil)
	}

	return nil
}

func GetBackendSetCompositeId(backendSetName string, loadBalancerId string) string {
	backendSetName = url.PathEscape(backendSetName)
	loadBalancerId = url.PathEscape(loadBalancerId)
	compositeId := "loadBalancers/" + loadBalancerId + "/backendSets/" + backendSetName
	return compositeId
}

func parseBackendSetCompositeId(compositeId string) (backendSetName string, loadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/backendSets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	backendSetName, _ = url.PathUnescape(parts[3])

	return
}

func (s *LoadBalancerBackendSetResourceCrud) mapToBackendDetails(fieldKeyFormat string) (oci_load_balancer.BackendDetails, error) {
	result := oci_load_balancer.BackendDetails{}

	if backup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup")); ok {
		tmp := backup.(bool)
		result.Backup = &tmp
	}

	if drain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "drain")); ok {
		tmp := drain.(bool)
		result.Drain = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_address")); ok {
		tmp := ipAddress.(string)
		result.IpAddress = &tmp
	}

	if offline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "offline")); ok {
		tmp := offline.(bool)
		result.Offline = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if weight, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weight")); ok {
		tmp := weight.(int)
		result.Weight = &tmp
	}

	return result, nil
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

func (s *LoadBalancerBackendSetResourceCrud) mapToHealthCheckerDetails(fieldKeyFormat string) (oci_load_balancer.HealthCheckerDetails, error) {
	result := oci_load_balancer.HealthCheckerDetails{}

	if intervalMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_ms")); ok {
		tmp := intervalMs.(int)
		result.IntervalInMillis = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if responseBodyRegex, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_body_regex")); ok {
		tmp := responseBodyRegex.(string)
		result.ResponseBodyRegex = &tmp
	}

	if retries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retries")); ok {
		tmp := retries.(int)
		result.Retries = &tmp
	}

	if returnCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "return_code")); ok {
		tmp := returnCode.(int)
		result.ReturnCode = &tmp
	}

	if timeoutInMillis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_millis")); ok {
		tmp := timeoutInMillis.(int)
		result.TimeoutInMillis = &tmp
	}

	if urlPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url_path")); ok {
		tmp := urlPath.(string)
		result.UrlPath = &tmp
	}

	return result, nil
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

func (s *LoadBalancerBackendSetResourceCrud) mapToLBCookieSessionPersistenceConfigurationDetails(fieldKeyFormat string) (oci_load_balancer.LbCookieSessionPersistenceConfigurationDetails, error) {
	result := oci_load_balancer.LbCookieSessionPersistenceConfigurationDetails{}

	if cookieName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cookie_name")); ok {
		tmp := cookieName.(string)
		result.CookieName = &tmp
	}

	if disableFallback, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disable_fallback")); ok {
		tmp := disableFallback.(bool)
		result.DisableFallback = &tmp
	}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		//@Codegen: When not specified, an unwanted empty string is set for this attribute in terraform state. This check removes this unwanted value before sending request
		if tmp != "" {
			result.Domain = &tmp
		}
	}

	if isHttpOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_http_only")); ok {
		tmp := isHttpOnly.(bool)
		result.IsHttpOnly = &tmp
	}

	if isSecure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure")); ok {
		tmp := isSecure.(bool)
		result.IsSecure = &tmp
	}

	if maxAgeInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_age_in_seconds")); ok {
		tmp := maxAgeInSeconds.(int)
		//@Codegen: When not specified, an unwanted value of 0 is set for this attribute in terraform state. This check removes this unwanted value before sending request.
		if tmp > 0 {
			result.MaxAgeInSeconds = &tmp
		}
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func LBCookieSessionPersistenceConfigurationDetailsToMap(obj *oci_load_balancer.LbCookieSessionPersistenceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CookieName != nil {
		result["cookie_name"] = string(*obj.CookieName)
	}

	if obj.DisableFallback != nil {
		result["disable_fallback"] = bool(*obj.DisableFallback)
	}

	if obj.Domain != nil {
		result["domain"] = string(*obj.Domain)
	}

	if obj.IsHttpOnly != nil {
		result["is_http_only"] = bool(*obj.IsHttpOnly)
	}

	if obj.IsSecure != nil {
		result["is_secure"] = bool(*obj.IsSecure)
	}

	if obj.MaxAgeInSeconds != nil {
		result["max_age_in_seconds"] = int(*obj.MaxAgeInSeconds)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *LoadBalancerBackendSetResourceCrud) mapToSSLConfigurationDetails(fieldKeyFormat string) (oci_load_balancer.SslConfigurationDetails, error) {
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
			result.CertificateName = &tmp
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

	if protocols, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocols")); ok {
		interfaces := protocols.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = fmt.Sprintf("%s", interfaces[i])
		}
		result.Protocols = tmp
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

	if cipherSuiteName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cipher_suite_name")); ok {
		tmp := cipherSuiteName.(string)
		result.CipherSuiteName = &tmp
	}

	if serverOrderPreference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_order_preference")); ok {
		result.ServerOrderPreference = oci_load_balancer.SslConfigurationDetailsServerOrderPreferenceEnum(serverOrderPreference.(string))
	}

	return result, nil
}

func SSLConfigurationToMap(obj *oci_load_balancer.SslConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateIds != nil {
		result["certificate_ids"] = obj.CertificateIds
	}

	if obj.CertificateName != nil {
		result["certificate_name"] = string(*obj.CertificateName)
	}

	if obj.CipherSuiteName != nil {
		result["cipher_suite_name"] = string(*obj.CipherSuiteName)
	}

	result["protocols"] = obj.Protocols

	result["server_order_preference"] = string(obj.ServerOrderPreference)

	if obj.TrustedCertificateAuthorityIds != nil {
		result["trusted_certificate_authority_ids"] = obj.TrustedCertificateAuthorityIds
	}

	if obj.VerifyDepth != nil {
		result["verify_depth"] = int(*obj.VerifyDepth)
	}

	if obj.VerifyPeerCertificate != nil {
		result["verify_peer_certificate"] = bool(*obj.VerifyPeerCertificate)
	}

	if obj.Protocols != nil {
		result["protocols"] = obj.Protocols
	}

	return result
}

func (s *LoadBalancerBackendSetResourceCrud) mapToSessionPersistenceConfigurationDetails(fieldKeyFormat string) (oci_load_balancer.SessionPersistenceConfigurationDetails, error) {
	result := oci_load_balancer.SessionPersistenceConfigurationDetails{}

	if cookieName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cookie_name")); ok {
		tmp := cookieName.(string)
		result.CookieName = &tmp
	}

	if disableFallback, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disable_fallback")); ok {
		tmp := disableFallback.(bool)
		result.DisableFallback = &tmp
	}

	return result, nil
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

func backendHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if backup, ok := m["backup"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", backup))
	}
	if drain, ok := m["drain"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", drain))
	}
	if ipAddress, ok := m["ip_address"]; ok && ipAddress != "" {
		buf.WriteString(fmt.Sprintf("%v-", ipAddress))
	}
	if offline, ok := m["offline"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", offline))
	}
	if port, ok := m["port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", port))
	}
	if weight, ok := m["weight"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", weight))
	}
	return hashcode.String(buf.String())
}
