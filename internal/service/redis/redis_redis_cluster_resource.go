// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisRedisClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisRedisCluster,
		Read:     readRedisRedisCluster,
		Update:   updateRedisRedisCluster,
		Delete:   deleteRedisRedisCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"node_memory_in_gbs": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"software_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_endpoint_fqdn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_endpoint_ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"primary_endpoint_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicas_endpoint_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicas_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createRedisRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.ReadResource(sync)
}

func updateRedisRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRedisRedisCluster(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RedisRedisClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.RedisClusterClient
	Res                    *oci_redis.RedisCluster
	DisableNotFoundRetries bool
}

func (s *RedisRedisClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RedisRedisClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_redis.RedisClusterLifecycleStateCreating),
	}
}

func (s *RedisRedisClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_redis.RedisClusterLifecycleStateActive),
	}
}

func (s *RedisRedisClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_redis.RedisClusterLifecycleStateDeleting),
	}
}

func (s *RedisRedisClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_redis.RedisClusterLifecycleStateDeleted),
	}
}

func (s *RedisRedisClusterResourceCrud) Create() error {
	request := oci_redis.CreateRedisClusterRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	if nodeMemoryInGBs, ok := s.D.GetOkExists("node_memory_in_gbs"); ok {
		tmp, ok := nodeMemoryInGBs.(float32)
		if !ok {
			tmp = float32(nodeMemoryInGBs.(float64))
		}
		request.NodeMemoryInGBs = &tmp
	}

	if softwareVersion, ok := s.D.GetOkExists("software_version"); ok {
		request.SoftwareVersion = oci_redis.RedisClusterSoftwareVersionEnum(softwareVersion.(string))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.CreateRedisCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_redis.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_redis.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "cluster") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getRedisClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RedisRedisClusterResourceCrud) getRedisClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_redis.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	redisClusterId, err := redisClusterWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, redisClusterId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_redis.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*redisClusterId)

	return s.Get()
}

func redisClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "redis", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_redis.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func redisClusterWaitForWorkRequest(wId *string, entityType string, action oci_redis.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_redis.RedisClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "redis")
	retryPolicy.ShouldRetryOperation = redisClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_redis.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_redis.OperationStatusInProgress),
			string(oci_redis.OperationStatusAccepted),
			string(oci_redis.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_redis.OperationStatusSucceeded),
			string(oci_redis.OperationStatusFailed),
			string(oci_redis.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_redis.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_redis.OperationStatusFailed || response.Status == oci_redis.OperationStatusCanceled {
		return nil, getErrorFromRedisRedisClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRedisRedisClusterWorkRequest(client *oci_redis.RedisClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_redis.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_redis.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *RedisRedisClusterResourceCrud) Get() error {
	request := oci_redis.GetRedisClusterRequest{}

	tmp := s.D.Id()
	request.RedisClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.GetRedisCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RedisCluster
	return nil
}

func (s *RedisRedisClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request := oci_redis.UpdateRedisClusterRequest{}
		request.DefinedTags = convertedDefinedTags
		err = s.updateRedisCluster(request)
		if err != nil {
			return err
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		oldDisplayName, newDisplayName := s.D.GetChange("display_name")
		if oldDisplayName != "" && newDisplayName != "" {
			request := oci_redis.UpdateRedisClusterRequest{}
			tmp := displayName.(string)
			request.DisplayName = &tmp
			err := s.updateRedisCluster(request)
			if err != nil {
				return err
			}
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request := oci_redis.UpdateRedisClusterRequest{}
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		err := s.updateRedisCluster(request)
		if err != nil {
			return err
		}
	}

	if nodeMemoryInGBs, ok := s.D.GetOkExists("node_memory_in_gbs"); ok && s.D.HasChange("node_memory_in_gbs") {
		request := oci_redis.UpdateRedisClusterRequest{}
		tmp, ok := nodeMemoryInGBs.(float32)
		if !ok {
			tmp = float32(nodeMemoryInGBs.(float64))
		}
		request.NodeMemoryInGBs = &tmp
		err := s.updateRedisCluster(request)
		if err != nil {
			return err
		}
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok && s.D.HasChange("node_count") {
		request := oci_redis.UpdateRedisClusterRequest{}
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
		err := s.updateRedisCluster(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RedisRedisClusterResourceCrud) Delete() error {
	request := oci_redis.DeleteRedisClusterRequest{}

	tmp := s.D.Id()
	request.RedisClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.DeleteRedisCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := redisClusterWaitForWorkRequest(workId, "cluster",
		oci_redis.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *RedisRedisClusterResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NodeCollection != nil {
		s.D.Set("node_collection", []interface{}{NodeCollectionToMap(s.Res.NodeCollection)})
	} else {
		s.D.Set("node_collection", nil)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NodeMemoryInGBs != nil {
		s.D.Set("node_memory_in_gbs", *s.Res.NodeMemoryInGBs)
	}

	if s.Res.PrimaryEndpointIpAddress != nil {
		s.D.Set("primary_endpoint_ip_address", *s.Res.PrimaryEndpointIpAddress)
	}

	if s.Res.PrimaryFqdn != nil {
		s.D.Set("primary_fqdn", *s.Res.PrimaryFqdn)
	}

	if s.Res.ReplicasEndpointIpAddress != nil {
		s.D.Set("replicas_endpoint_ip_address", *s.Res.ReplicasEndpointIpAddress)
	}

	if s.Res.ReplicasFqdn != nil {
		s.D.Set("replicas_fqdn", *s.Res.ReplicasFqdn)
	}

	s.D.Set("software_version", s.Res.SoftwareVersion)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func NodeToMap(obj oci_redis.Node) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.PrivateEndpointFqdn != nil {
		result["private_endpoint_fqdn"] = string(*obj.PrivateEndpointFqdn)
	}

	if obj.PrivateEndpointIpAddress != nil {
		result["private_endpoint_ip_address"] = string(*obj.PrivateEndpointIpAddress)
	}

	return result
}

func NodeCollectionToMap(obj *oci_redis.NodeCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, NodeToMap(item))
	}
	result["items"] = items

	return result
}

func RedisClusterSummaryToMap(obj oci_redis.RedisClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.NodeCount != nil {
		result["node_count"] = int(*obj.NodeCount)
	}

	if obj.NodeMemoryInGBs != nil {
		result["node_memory_in_gbs"] = float32(*obj.NodeMemoryInGBs)
	}

	if obj.PrimaryEndpointIpAddress != nil {
		result["primary_endpoint_ip_address"] = string(*obj.PrimaryEndpointIpAddress)
	}

	if obj.PrimaryFqdn != nil {
		result["primary_fqdn"] = string(*obj.PrimaryFqdn)
	}

	if obj.ReplicasEndpointIpAddress != nil {
		result["replicas_endpoint_ip_address"] = string(*obj.ReplicasEndpointIpAddress)
	}

	if obj.ReplicasFqdn != nil {
		result["replicas_fqdn"] = string(*obj.ReplicasFqdn)
	}

	result["software_version"] = string(obj.SoftwareVersion)

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *RedisRedisClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_redis.ChangeRedisClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RedisClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ChangeRedisClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRedisClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RedisRedisClusterResourceCrud) updateRedisCluster(request oci_redis.UpdateRedisClusterRequest) error {
	tmp := s.D.Id()
	request.RedisClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.UpdateRedisCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRedisClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
