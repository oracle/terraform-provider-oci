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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisOciCacheUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisOciCacheUser,
		Read:     readRedisOciCacheUser,
		Update:   updateRedisOciCacheUser,
		Delete:   deleteRedisOciCacheUser,
		Schema: map[string]*schema.Schema{
			// Required
			"acl_string": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authentication_mode": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"authentication_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IAM",
								"PASSWORD",
							}, true),
						},

						// Optional
						"hashed_passwords": {
							Type:      schema.TypeList,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createRedisOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readRedisOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.ReadResource(sync)
}

func updateRedisOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRedisOciCacheUser(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheUserClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RedisOciCacheUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.OciCacheUserClient
	Res                    *oci_redis.OciCacheUser
	RedisClusterClient     *oci_redis.RedisClusterClient
	DisableNotFoundRetries bool
}

func (s *RedisOciCacheUserResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RedisOciCacheUserResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_redis.OciCacheUserLifecycleStateCreating),
	}
}

func (s *RedisOciCacheUserResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_redis.OciCacheUserLifecycleStateActive),
	}
}

func (s *RedisOciCacheUserResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_redis.OciCacheUserLifecycleStateDeleting),
	}
}

func (s *RedisOciCacheUserResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_redis.OciCacheUserLifecycleStateDeleted),
	}
}

func (s *RedisOciCacheUserResourceCrud) Create() error {
	request := oci_redis.CreateOciCacheUserRequest{}

	if aclString, ok := s.D.GetOkExists("acl_string"); ok {
		tmp := aclString.(string)
		request.AclString = &tmp
	}

	if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
		if tmpList := authenticationMode.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "authentication_mode", 0)
			tmp, err := s.mapToAuthenticationMode(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AuthenticationMode = tmp
		}
	}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_redis.OciCacheUserStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.CreateOciCacheUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_redis.GetWorkRequestResponse{}
	workRequestResponse, err = s.RedisClusterClient.GetWorkRequest(context.Background(),
		oci_redis.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "ocicacheuser") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getOciCacheUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RedisOciCacheUserResourceCrud) getOciCacheUserFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_redis.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	ociCacheUserId, err := ociCacheUserWaitForWorkRequest(workId, "ocicacheuser",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.RedisClusterClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, ociCacheUserId)
		_, cancelErr := s.RedisClusterClient.CancelWorkRequest(context.Background(),
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
	s.D.SetId(*ociCacheUserId)

	return s.Get()
}

func ociCacheUserWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func ociCacheUserWaitForWorkRequest(wId *string, entityType string, action oci_redis.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_redis.RedisClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "redis")
	retryPolicy.ShouldRetryOperation = ociCacheUserWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromRedisOciCacheUserWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRedisOciCacheUserWorkRequest(client *oci_redis.RedisClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_redis.ActionTypeEnum) error {
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

func (s *RedisOciCacheUserResourceCrud) Get() error {
	request := oci_redis.GetOciCacheUserRequest{}

	tmp := s.D.Id()
	request.OciCacheUserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.GetOciCacheUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OciCacheUser
	return nil
}

func (s *RedisOciCacheUserResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_redis.UpdateOciCacheUserRequest{}

	if aclString, ok := s.D.GetOkExists("acl_string"); ok {
		tmp := aclString.(string)
		request.AclString = &tmp
	}

	if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
		if tmpList := authenticationMode.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "authentication_mode", 0)
			tmp, err := s.mapToAuthenticationMode(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AuthenticationMode = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.OciCacheUserId = &tmp

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_redis.OciCacheUserStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.UpdateOciCacheUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOciCacheUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RedisOciCacheUserResourceCrud) Delete() error {
	request := oci_redis.DeleteOciCacheUserRequest{}

	tmp := s.D.Id()
	request.OciCacheUserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.DeleteOciCacheUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := ociCacheUserWaitForWorkRequest(workId, "ocicacheuser",
		oci_redis.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.RedisClusterClient)
	return delWorkRequestErr
}

func (s *RedisOciCacheUserResourceCrud) SetData() error {
	if s.Res.AclString != nil {
		s.D.Set("acl_string", *s.Res.AclString)
	}

	if s.Res.AuthenticationMode != nil {
		authenticationModeArray := []interface{}{}
		authenticationModeMap := AuthenticationModeToMap(&s.Res.AuthenticationMode)

		// Special handling for PASSWORD authentication with empty hashed_passwords
		if authType, ok := authenticationModeMap["authentication_type"].(string); ok &&
			strings.EqualFold(authType, "PASSWORD") {

			// Check if hashed_passwords is empty from API
			hashedPasswords, ok := authenticationModeMap["hashed_passwords"].([]string)
			if !ok || len(hashedPasswords) == 0 {
				// Check if we have existing values in state
				if existingAuth, ok := s.D.GetOk("authentication_mode"); ok {
					if existingAuthList := existingAuth.([]interface{}); len(existingAuthList) > 0 {
						if existingAuthMap, ok := existingAuthList[0].(map[string]interface{}); ok {
							if existingPasswords, ok := existingAuthMap["hashed_passwords"]; ok {
								// Preserve existing hashed_passwords from state
								authenticationModeMap["hashed_passwords"] = existingPasswords
							}
						}
					}
				}
			}
		}

		authenticationModeArray = append(authenticationModeArray, authenticationModeMap)
		s.D.Set("authentication_mode", authenticationModeArray)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

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

func (s *RedisOciCacheUserResourceCrud) mapToAuthenticationMode(fieldKeyFormat string) (oci_redis.AuthenticationMode, error) {
	var baseObject oci_redis.AuthenticationMode
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("IAM"):
		details := oci_redis.IamAuthenticationMode{}
		baseObject = details
	case strings.ToLower("PASSWORD"):
		details := oci_redis.PasswordAuthenticationMode{}
		if hashedPasswords, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hashed_passwords")); ok {
			interfaces := hashedPasswords.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hashed_passwords")) {
				details.HashedPasswords = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func AuthenticationModeToMap(obj *oci_redis.AuthenticationMode) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_redis.IamAuthenticationMode:
		result["authentication_type"] = "IAM"
	case oci_redis.PasswordAuthenticationMode:
		result["authentication_type"] = "PASSWORD"

		// Include hashed_passwords even if empty
		result["hashed_passwords"] = v.HashedPasswords
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func OciCacheUserSummaryToMap(obj oci_redis.OciCacheUserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["authentication_type"] = string(obj.AuthenticationType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *RedisOciCacheUserResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_redis.ChangeOciCacheUserCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OciCacheUserId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ChangeOciCacheUserCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOciCacheUserFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
