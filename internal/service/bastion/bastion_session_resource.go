// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BastionSessionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBastionSession,
		Read:     readBastionSession,
		Update:   updateBastionSession,
		Delete:   deleteBastionSession,
		Schema: map[string]*schema.Schema{
			// Required
			"bastion_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"public_key_content": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"target_resource_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"session_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DYNAMIC_PORT_FORWARDING",
								"MANAGED_SSH",
								"PORT_FORWARDING",
							}, true),
						},

						// Optional
						"target_resource_fqdn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"target_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"target_resource_operating_system_user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"target_resource_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"target_resource_private_ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"target_resource_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"session_ttl_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"bastion_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bastion_public_host_key_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bastion_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createBastionSession(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.CreateResource(d, sync)
}

func readBastionSession(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

func updateBastionSession(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBastionSession(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BastionSessionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bastion.BastionClient
	Res                    *oci_bastion.Session
	DisableNotFoundRetries bool
}

func (s *BastionSessionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BastionSessionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bastion.SessionLifecycleStateCreating),
	}
}

func (s *BastionSessionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bastion.SessionLifecycleStateActive),
	}
}

func (s *BastionSessionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bastion.SessionLifecycleStateDeleting),
	}
}

func (s *BastionSessionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bastion.SessionLifecycleStateDeleted),
	}
}

func (s *BastionSessionResourceCrud) Create() error {
	request := oci_bastion.CreateSessionRequest{}

	if bastionId, ok := s.D.GetOkExists("bastion_id"); ok {
		tmp := bastionId.(string)
		request.BastionId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if keyDetails, ok := s.D.GetOkExists("key_details"); ok {
		if tmpList := keyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_details", 0)
			tmp, err := s.mapToPublicKeyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KeyDetails = &tmp
		}
	}

	if keyType, ok := s.D.GetOkExists("key_type"); ok {
		request.KeyType = oci_bastion.CreateSessionDetailsKeyTypeEnum(keyType.(string))
	}

	if sessionTtlInSeconds, ok := s.D.GetOkExists("session_ttl_in_seconds"); ok {
		tmp := sessionTtlInSeconds.(int)
		request.SessionTtlInSeconds = &tmp
	}

	if targetResourceDetails, ok := s.D.GetOkExists("target_resource_details"); ok {
		if tmpList := targetResourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_resource_details", 0)
			tmp, err := s.mapToCreateSessionTargetResourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetResourceDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.CreateSession(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSessionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion"), oci_bastion.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BastionSessionResourceCrud) getSessionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bastion.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	sessionId, err := sessionWaitForWorkRequest(workId, "session",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*sessionId)

	return s.Get()
}

func sessionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bastion", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bastion.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func sessionWaitForWorkRequest(wId *string, entityType string, action oci_bastion.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bastion.BastionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bastion")
	retryPolicy.ShouldRetryOperation = sessionWorkRequestShouldRetryFunc(timeout)

	response := oci_bastion.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bastion.OperationStatusInProgress),
			string(oci_bastion.OperationStatusAccepted),
			string(oci_bastion.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bastion.OperationStatusSucceeded),
			string(oci_bastion.OperationStatusFailed),
			string(oci_bastion.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bastion.GetWorkRequestRequest{
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
			log.Printf("[DEBUG] res action type %v and action type: %v\n", res.ActionType, action)
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}
	log.Printf("[DEBUG] identifier is %v status is %v \n", identifier, response.Status)
	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bastion.OperationStatusFailed || response.Status == oci_bastion.OperationStatusCanceled {
		return nil, getErrorFromBastionSessionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBastionSessionWorkRequest(client *oci_bastion.BastionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bastion.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bastion.ListWorkRequestErrorsRequest{
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

func (s *BastionSessionResourceCrud) Get() error {
	request := oci_bastion.GetSessionRequest{}

	tmp := s.D.Id()
	request.SessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.GetSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Session
	return nil
}

func (s *BastionSessionResourceCrud) Update() error {
	request := oci_bastion.UpdateSessionRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.SessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.UpdateSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Session
	return nil
}

func (s *BastionSessionResourceCrud) Delete() error {
	request := oci_bastion.DeleteSessionRequest{}

	tmp := s.D.Id()
	request.SessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.DeleteSession(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sessionWaitForWorkRequest(workId, "session",
		oci_bastion.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BastionSessionResourceCrud) SetData() error {
	if s.Res.BastionId != nil {
		s.D.Set("bastion_id", *s.Res.BastionId)
	}

	if s.Res.BastionName != nil {
		s.D.Set("bastion_name", *s.Res.BastionName)
	}

	if s.Res.BastionPublicHostKeyInfo != nil {
		s.D.Set("bastion_public_host_key_info", *s.Res.BastionPublicHostKeyInfo)
	}

	if s.Res.BastionUserName != nil {
		s.D.Set("bastion_user_name", *s.Res.BastionUserName)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.KeyDetails != nil {
		s.D.Set("key_details", []interface{}{PublicKeyDetailsToMap(s.Res.KeyDetails)})
	} else {
		s.D.Set("key_details", nil)
	}

	s.D.Set("key_type", s.Res.KeyType)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SessionTtlInSeconds != nil {
		s.D.Set("session_ttl_in_seconds", *s.Res.SessionTtlInSeconds)
	}

	s.D.Set("ssh_metadata", s.Res.SshMetadata)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetResourceDetails != nil {
		targetResourceDetailsArray := []interface{}{}
		if targetResourceDetailsMap := TargetResourceDetailsToMap(&s.Res.TargetResourceDetails); targetResourceDetailsMap != nil {
			targetResourceDetailsArray = append(targetResourceDetailsArray, targetResourceDetailsMap)
		}
		s.D.Set("target_resource_details", targetResourceDetailsArray)
	} else {
		s.D.Set("target_resource_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *BastionSessionResourceCrud) mapToCreateSessionTargetResourceDetails(fieldKeyFormat string) (oci_bastion.CreateSessionTargetResourceDetails, error) {
	var baseObject oci_bastion.CreateSessionTargetResourceDetails
	//discriminator
	sessionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "session_type"))
	var sessionType string
	if ok {
		sessionType = sessionTypeRaw.(string)
	} else {
		sessionType = "" // default value
	}
	switch strings.ToLower(sessionType) {
	case strings.ToLower("DYNAMIC_PORT_FORWARDING"):
		details := oci_bastion.CreateDynamicPortForwardingSessionTargetResourceDetails{}
		baseObject = details
	case strings.ToLower("MANAGED_SSH"):
		details := oci_bastion.ManagedSshSessionTargetResourceDetails{}
		if targetResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_id")); ok {
			tmp := targetResourceId.(string)
			details.TargetResourceId = &tmp
		}
		if targetResourceOperatingSystemUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_operating_system_user_name")); ok {
			tmp := targetResourceOperatingSystemUserName.(string)
			details.TargetResourceOperatingSystemUserName = &tmp
		}
		if targetResourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_port")); ok {
			tmp := targetResourcePort.(int)
			details.TargetResourcePort = &tmp
		}
		if targetResourcePrivateIpAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_private_ip_address")); ok {
			tmp := targetResourcePrivateIpAddress.(string)
			details.TargetResourcePrivateIpAddress = &tmp
		}
		baseObject = details
	case strings.ToLower("PORT_FORWARDING"):
		details := oci_bastion.CreatePortForwardingSessionTargetResourceDetails{}
		if targetResourceFqdn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_fqdn")); ok {
			tmp := targetResourceFqdn.(string)
			details.TargetResourceFqdn = &tmp
		}
		if targetResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_id")); ok {
			tmp := targetResourceId.(string)
			details.TargetResourceId = &tmp
		}
		if targetResourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_port")); ok {
			tmp := targetResourcePort.(int)
			details.TargetResourcePort = &tmp
		}
		if targetResourcePrivateIpAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_resource_private_ip_address")); ok {
			tmp := targetResourcePrivateIpAddress.(string)
			details.TargetResourcePrivateIpAddress = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown session_type '%v' was specified", sessionType)
	}
	return baseObject, nil
}

func TargetResourceDetailsToMap(obj *oci_bastion.TargetResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_bastion.DynamicPortForwardingSessionTargetResourceDetails:
		result["session_type"] = "DYNAMIC_PORT_FORWARDING"
	case oci_bastion.ManagedSshSessionTargetResourceDetails:
		result["session_type"] = "MANAGED_SSH"

		if v.TargetResourceId != nil {
			result["target_resource_id"] = string(*v.TargetResourceId)
		}

		if v.TargetResourceOperatingSystemUserName != nil {
			result["target_resource_operating_system_user_name"] = string(*v.TargetResourceOperatingSystemUserName)
		}

		if v.TargetResourcePort != nil {
			result["target_resource_port"] = int(*v.TargetResourcePort)
		}

		if v.TargetResourcePrivateIpAddress != nil {
			result["target_resource_private_ip_address"] = string(*v.TargetResourcePrivateIpAddress)
		}
	case oci_bastion.PortForwardingSessionTargetResourceDetails:
		result["session_type"] = "PORT_FORWARDING"

		if v.TargetResourceFqdn != nil {
			result["target_resource_fqdn"] = string(*v.TargetResourceFqdn)
		}

		if v.TargetResourceId != nil {
			result["target_resource_id"] = string(*v.TargetResourceId)
		}

		if v.TargetResourcePort != nil {
			result["target_resource_port"] = int(*v.TargetResourcePort)
		}

		if v.TargetResourcePrivateIpAddress != nil {
			result["target_resource_private_ip_address"] = string(*v.TargetResourcePrivateIpAddress)
		}
	default:
		log.Printf("[WARN] Received 'session_type' of unknown type %T", *obj)
		return nil
	}

	return result
}

func (s *BastionSessionResourceCrud) mapToPublicKeyDetails(fieldKeyFormat string) (oci_bastion.PublicKeyDetails, error) {
	result := oci_bastion.PublicKeyDetails{}

	if publicKeyContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_key_content")); ok {
		tmp := publicKeyContent.(string)
		result.PublicKeyContent = &tmp
	}

	return result, nil
}

func PublicKeyDetailsToMap(obj *oci_bastion.PublicKeyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PublicKeyContent != nil {
		result["public_key_content"] = string(*obj.PublicKeyContent)
	}

	return result
}
