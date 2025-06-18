// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisRedisClusterDetachOciCacheUserResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisRedisClusterDetachOciCacheUser,
		Read:     readRedisRedisClusterDetachOciCacheUser,
		Delete:   deleteRedisRedisClusterDetachOciCacheUser,
		Schema: map[string]*schema.Schema{
			// Required
			"oci_cache_users": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"redis_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createRedisRedisClusterDetachOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisRedisClusterDetachOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisRedisClusterDetachOciCacheUser(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteRedisRedisClusterDetachOciCacheUser(d *schema.ResourceData, m interface{}) error {
	return nil
}

type RedisRedisClusterDetachOciCacheUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.RedisClusterClient
	DisableNotFoundRetries bool
}

func (s *RedisRedisClusterDetachOciCacheUserResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("RedisRedisClusterDetachOciCacheUserResource-", RedisRedisClusterDetachOciCacheUserResource(), s.D)
}

func (s *RedisRedisClusterDetachOciCacheUserResourceCrud) Create() error {
	request := oci_redis.DetachOciCacheUsersRequest{}

	if ociCacheUsers, ok := s.D.GetOkExists("oci_cache_users"); ok {
		interfaces := ociCacheUsers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("oci_cache_users") {
			request.OciCacheUsers = tmp
		}
	}

	if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
		tmp := redisClusterId.(string)
		request.RedisClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.DetachOciCacheUsers(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRedisClusterDetachOciCacheUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RedisRedisClusterDetachOciCacheUserResourceCrud) getRedisClusterDetachOciCacheUserFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_redis.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	redisClusterDetachOciCacheUserId, err := redisClusterDetachOciCacheUserWaitForWorkRequest(workId, "cluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, redisClusterDetachOciCacheUserId)
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
	s.D.SetId(*redisClusterDetachOciCacheUserId)

	//return s.Get()
	return nil
}

func redisClusterDetachOciCacheUserWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func redisClusterDetachOciCacheUserWaitForWorkRequest(wId *string, entityType string, action oci_redis.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_redis.RedisClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "redis")
	retryPolicy.ShouldRetryOperation = redisClusterDetachOciCacheUserWorkRequestShouldRetryFunc(timeout)

	response := oci_redis.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromRedisRedisClusterDetachOciCacheUserWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRedisRedisClusterDetachOciCacheUserWorkRequest(client *oci_redis.RedisClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_redis.ActionTypeEnum) error {
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

func (s *RedisRedisClusterDetachOciCacheUserResourceCrud) SetData() error {
	return nil
}
