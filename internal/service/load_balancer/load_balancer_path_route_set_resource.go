// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"
)

func LoadBalancerPathRouteSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerPathRouteSet,
		Read:     readLoadBalancerPathRouteSet,
		Update:   updateLoadBalancerPathRouteSet,
		Delete:   deleteLoadBalancerPathRouteSet,
		Schema: map[string]*schema.Schema{
			// Required
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
			"path_routes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"backend_set_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path_match_type": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"match_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerPathRouteSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.PathRouteSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerPathRouteSetResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return GetPathRouteSetCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerPathRouteSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerPathRouteSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerPathRouteSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerPathRouteSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerPathRouteSetResourceCrud) Create() error {
	request := oci_load_balancer.CreatePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if pathRoutes, ok := s.D.GetOkExists("path_routes"); ok {
		interfaces := pathRoutes.([]interface{})
		tmp := make([]oci_load_balancer.PathRoute, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "path_routes", stateDataIndex)
			converted, err := s.mapToPathRoute(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("path_routes") {
			request.PathRoutes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreatePathRouteSet(context.Background(), request)
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

func (s *LoadBalancerPathRouteSetResourceCrud) Get() error {
	_, stillWorking, err := tfresource.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetPathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		loadBalancerId, pathRouteSetName, err := parsePathRouteSetCompositeId(s.D.Id())
		if err == nil {
			request.LoadBalancerId = &loadBalancerId
			request.PathRouteSetName = &pathRouteSetName
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetPathRouteSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PathRouteSet
	return nil
}

func (s *LoadBalancerPathRouteSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdatePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if pathRoutes, ok := s.D.GetOkExists("path_routes"); ok {
		interfaces := pathRoutes.([]interface{})
		tmp := make([]oci_load_balancer.PathRoute, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "path_routes", stateDataIndex)
			converted, err := s.mapToPathRoute(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("path_routes") {
			request.PathRoutes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdatePathRouteSet(context.Background(), request)
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

func (s *LoadBalancerPathRouteSetResourceCrud) Delete() error {
	request := oci_load_balancer.DeletePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeletePathRouteSet(context.Background(), request)
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

func (s *LoadBalancerPathRouteSetResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	loadBalancerId, pathRouteSetName, err := parsePathRouteSetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("load_balancer_id", &loadBalancerId)
		s.D.Set("name", &pathRouteSetName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	pathRoutes := []interface{}{}
	for _, item := range s.Res.PathRoutes {
		pathRoutes = append(pathRoutes, PathRouteToMap(item))
	}
	s.D.Set("path_routes", pathRoutes)

	return nil
}

func GetPathRouteSetCompositeId(loadBalancerId string, pathRouteSetName string) string {
	loadBalancerId = url.PathEscape(loadBalancerId)
	pathRouteSetName = url.PathEscape(pathRouteSetName)
	compositeId := "loadBalancers/" + loadBalancerId + "/pathRouteSets/" + pathRouteSetName
	return compositeId
}

func parsePathRouteSetCompositeId(compositeId string) (loadBalancerId string, pathRouteSetName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/pathRouteSets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	pathRouteSetName, _ = url.PathUnescape(parts[3])

	return
}

func (s *LoadBalancerPathRouteSetResourceCrud) mapToPathMatchType(fieldKeyFormat string) (oci_load_balancer.PathMatchType, error) {
	result := oci_load_balancer.PathMatchType{}

	if matchType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_type")); ok {
		result.MatchType = oci_load_balancer.PathMatchTypeMatchTypeEnum(matchType.(string))
	}

	return result, nil
}

func PathMatchTypeToMap(obj *oci_load_balancer.PathMatchType) map[string]interface{} {
	result := map[string]interface{}{}

	result["match_type"] = string(obj.MatchType)

	return result
}

func (s *LoadBalancerPathRouteSetResourceCrud) mapToPathRoute(fieldKeyFormat string) (oci_load_balancer.PathRoute, error) {
	result := oci_load_balancer.PathRoute{}

	if backendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_name")); ok {
		tmp := backendSetName.(string)
		result.BackendSetName = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	if pathMatchType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path_match_type")); ok {
		if tmpList := pathMatchType.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "path_match_type"), 0)
			tmp, err := s.mapToPathMatchType(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert path_match_type, encountered error: %v", err)
			}
			result.PathMatchType = &tmp
		}
	}

	return result, nil
}

func PathRouteToMap(obj oci_load_balancer.PathRoute) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackendSetName != nil {
		result["backend_set_name"] = string(*obj.BackendSetName)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.PathMatchType != nil {
		result["path_match_type"] = []interface{}{PathMatchTypeToMap(obj.PathMatchType)}
	}

	return result
}
