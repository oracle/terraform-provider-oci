// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

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

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func CoreInstancePoolInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreInstancePoolInstance,
		Read:     readCoreInstancePoolInstance,
		Delete:   deleteCoreInstancePoolInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_pool_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"decrement_size_on_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			// Optional
			"auto_terminate_instance_on_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_configuration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balancer_backends": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backend_health_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backend_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backend_set_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstancePoolInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreInstancePoolInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

func deleteCoreInstancePoolInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreInstancePoolInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstancePoolInstance
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreInstancePoolInstanceResourceCrud) ID() string {
	return GetInstancePoolInstanceCompositeId(s.D.Get("instance_pool_id").(string), s.D.Get("instance_id").(string))
}

func (s *CoreInstancePoolInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstancePoolInstanceLifecycleStateAttaching),
	}
}

func (s *CoreInstancePoolInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstancePoolInstanceLifecycleStateActive),
	}
}

func (s *CoreInstancePoolInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstancePoolInstanceLifecycleStateDetaching),
	}
}

func (s *CoreInstancePoolInstanceResourceCrud) Create() error {
	request := oci_core.AttachInstancePoolInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachInstancePoolInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.InstancePoolInstance

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "instancepool", oci_work_requests.WorkRequestResourceActionTypeRelated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CoreInstancePoolInstanceResourceCrud) Get() error {
	request := oci_core.GetInstancePoolInstanceRequest{}
	instancePoolId, instanceId, err := parseInstancePoolInstanceCompositeId(s.D.Id())
	if err == nil {
		request.InstancePoolId = &instancePoolId
		request.InstanceId = &instanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstancePoolInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstancePoolInstance
	return nil
}

func (s *CoreInstancePoolInstanceResourceCrud) Delete() error {
	request := oci_core.DetachInstancePoolInstanceRequest{}
	instancePoolId, instanceId, err := parseInstancePoolInstanceCompositeId(s.D.Id())
	if err == nil {
		request.InstancePoolId = &instancePoolId
		request.InstanceId = &instanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	if decrementSizeOnDelete, ok := s.D.GetOkExists("decrement_size_on_delete"); ok {
		tmp := decrementSizeOnDelete.(bool)
		request.IsDecrementSize = &tmp
	}

	if autoTerminateInstanceOnDelete, ok := s.D.GetOkExists("auto_terminate_instance_on_delete"); ok {
		tmp := autoTerminateInstanceOnDelete.(bool)
		request.IsAutoTerminate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DetachInstancePoolInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "instancepool", oci_work_requests.WorkRequestResourceActionTypeRelated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CoreInstancePoolInstanceResourceCrud) SetData() error {

	instancePoolId, instanceId, err := parseInstancePoolInstanceCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("instance_pool_id", &instancePoolId)
		s.D.Set("instance_id", &instanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	if s.Res.InstancePoolId != nil {
		s.D.Set("instance_pool_id", *s.Res.InstancePoolId)
	}

	loadBalancerBackends := []interface{}{}
	for _, item := range s.Res.LoadBalancerBackends {
		loadBalancerBackends = append(loadBalancerBackends, InstancePoolInstanceLoadBalancerBackendToMap(item))
	}
	s.D.Set("load_balancer_backends", loadBalancerBackends)

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.State != nil {
		s.D.Set("state", *s.Res.State)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetInstancePoolInstanceCompositeId(instancePoolId string, instanceId string) string {
	instanceId = url.PathEscape(instanceId)
	instancePoolId = url.PathEscape(instancePoolId)
	compositeId := "instancePools/" + instancePoolId + "/instances/" + instanceId
	return compositeId
}

func parseInstancePoolInstanceCompositeId(compositeId string) (instancePoolId string, instanceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("instancePools/.*/instances/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	instancePoolId, _ = url.PathUnescape(parts[1])
	instanceId, _ = url.PathUnescape(parts[3])

	return
}

func InstancePoolInstanceLoadBalancerBackendToMap(obj oci_core.InstancePoolInstanceLoadBalancerBackend) map[string]interface{} {
	result := map[string]interface{}{}

	result["backend_health_status"] = string(obj.BackendHealthStatus)

	if obj.BackendName != nil {
		result["backend_name"] = string(*obj.BackendName)
	}

	if obj.BackendSetName != nil {
		result["backend_set_name"] = string(*obj.BackendSetName)
	}

	if obj.LoadBalancerId != nil {
		result["load_balancer_id"] = string(*obj.LoadBalancerId)
	}

	return result
}
