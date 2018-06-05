// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func PathRouteSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPathRouteSet,
		Read:     readPathRouteSet,
		Update:   updatePathRouteSet,
		Delete:   deletePathRouteSet,
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

func createPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &PathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readPathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &PathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updatePathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &PathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deletePathRouteSet(d *schema.ResourceData, m interface{}) error {
	sync := &PathRouteSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type PathRouteSetResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.PathRouteSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *PathRouteSetResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *PathRouteSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *PathRouteSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *PathRouteSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *PathRouteSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *PathRouteSetResourceCrud) Create() error {
	request := oci_load_balancer.CreatePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.PathRoutes = []oci_load_balancer.PathRoute{}
	if pathRoutes, ok := s.D.GetOkExists("path_routes"); ok {
		interfaces := pathRoutes.([]interface{})
		tmp := make([]oci_load_balancer.PathRoute, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToPathRoute(toBeConverted.(map[string]interface{}))
		}
		request.PathRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreatePathRouteSet(context.Background(), request)
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

func (s *PathRouteSetResourceCrud) Get() error {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetPathRouteSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PathRouteSet
	return nil
}

func (s *PathRouteSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdatePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	request.PathRoutes = []oci_load_balancer.PathRoute{}
	if pathRoutes, ok := s.D.GetOkExists("path_routes"); ok {
		interfaces := pathRoutes.([]interface{})
		tmp := make([]oci_load_balancer.PathRoute, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToPathRoute(toBeConverted.(map[string]interface{}))
		}
		request.PathRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdatePathRouteSet(context.Background(), request)
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

func (s *PathRouteSetResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeletePathRouteSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeletePathRouteSet(context.Background(), request)

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

func (s *PathRouteSetResourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	pathRoutes := []interface{}{}
	for _, item := range s.Res.PathRoutes {
		pathRoutes = append(pathRoutes, PathRouteToMap(item))
	}
	s.D.Set("path_routes", pathRoutes)

}

func mapToPathMatchType(raw map[string]interface{}) oci_load_balancer.PathMatchType {
	result := oci_load_balancer.PathMatchType{}

	if matchType, ok := raw["match_type"]; ok && matchType != "" {
		tmp := oci_load_balancer.PathMatchTypeMatchTypeEnum(matchType.(string))
		result.MatchType = tmp
	}

	return result
}

func PathMatchTypeToMap(obj *oci_load_balancer.PathMatchType) map[string]interface{} {
	result := map[string]interface{}{}

	result["match_type"] = string(obj.MatchType)

	return result
}

func mapToPathRoute(raw map[string]interface{}) oci_load_balancer.PathRoute {
	result := oci_load_balancer.PathRoute{}

	if backendSetName, ok := raw["backend_set_name"]; ok && backendSetName != "" {
		tmp := backendSetName.(string)
		result.BackendSetName = &tmp
	}

	if path, ok := raw["path"]; ok && path != "" {
		tmp := path.(string)
		result.Path = &tmp
	}

	if pathMatchType, ok := raw["path_match_type"]; ok {
		if tmpList := pathMatchType.([]interface{}); len(tmpList) > 0 {
			tmp := mapToPathMatchType(tmpList[0].(map[string]interface{}))
			result.PathMatchType = &tmp
		}
	}

	return result
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
