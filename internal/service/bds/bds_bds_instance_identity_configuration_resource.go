// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceIdentityConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceIdentityConfiguration,
		Read:     readBdsBdsInstanceIdentityConfiguration,
		Update:   updateBdsBdsInstanceIdentityConfiguration,
		Delete:   deleteBdsBdsInstanceIdentityConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"confidential_application_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identity_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"iam_user_sync_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_posix_attributes_addition_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"upst_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"master_encryption_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"activate_iam_user_sync_configuration_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"activate_upst_configuration_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"refresh_confidential_application_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"refresh_upst_token_exchange_keytab_trigger": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"iam_user_sync_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_posix_attributes_addition_required": {
							Type:     schema.TypeBool,
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
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"upst_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"keytab_content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"master_encryption_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret_id": {
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
						"time_token_exchange_keytab_last_refreshed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"token_exchange_principal_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createBdsBdsInstanceIdentityConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceIdentityConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}
	return nil

}

func readBdsBdsInstanceIdentityConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceIdentityConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstanceIdentityConfiguration(d *schema.ResourceData, m interface{}) error {

	sync := &BdsBdsInstanceIdentityConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if triggerVal, ok := sync.D.GetOkExists("activate_iam_user_sync_configuration_trigger"); ok && sync.D.HasChange("activate_iam_user_sync_configuration_trigger") {
		if triggerVal.(string) == "true" {
			err := sync.ActivateIamUserSyncConfiguration()

			if err != nil {
				return err
			}
		}
	}

	if triggerVal, ok := sync.D.GetOkExists("activate_upst_configuration_trigger"); ok && sync.D.HasChange("activate_upst_configuration_trigger") {
		if triggerVal.(string) == "true" {
			err := sync.ActivateUpstConfiguration()

			if err != nil {
				return err
			}
		}
	}

	if refreshTrigger, ok := sync.D.GetOkExists("refresh_confidential_application_trigger"); ok {
		if refreshTrigger.(string) == "true" {
			err := sync.RefreshConfidentialApplication()
			if err != nil {
				return err
			}
		}
	}

	if refreshTrigger, ok := sync.D.GetOkExists("refresh_upst_token_exchange_keytab_trigger"); ok {
		if refreshTrigger.(string) == "true" {
			err := sync.RefreshUpstTokenExchangeKeytab()
			if err != nil {
				return err
			}
		}
	}

	if _, ok := sync.D.GetOkExists("activate_iam_user_sync_configuration_trigger"); ok && sync.D.HasChange("activate_iam_user_sync_configuration_trigger") {
		triggerVal := sync.D.Get("activate_iam_user_sync_configuration_trigger")
		if triggerVal.(string) == "false" {
			err := sync.DeactivateIamUserSyncConfiguration()

			if err != nil {
				return err
			}
		}
	}

	if _, ok := sync.D.GetOkExists("activate_upst_configuration_trigger"); ok && sync.D.HasChange("activate_upst_configuration_trigger") {
		triggerVal := sync.D.Get("activate_upst_configuration_trigger")
		if triggerVal.(string) == "false" {
			err := sync.DeactivateUpstConfiguration()

			if err != nil {
				return err
			}
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteBdsBdsInstanceIdentityConfiguration(d *schema.ResourceData, m interface{}) error {

	sync := &BdsBdsInstanceIdentityConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceIdentityConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.IdentityConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) ID() string {
	return GetBdsInstanceIdentityConfigurationCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.IdentityConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.IdentityConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) Create() error {
	request := oci_bds.CreateIdentityConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if confidentialApplicationId, ok := s.D.GetOkExists("confidential_application_id"); ok {
		tmp := confidentialApplicationId.(string)
		request.ConfidentialApplicationId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if _, ok := s.D.GetOkExists("activate_iam_user_sync_configuration_trigger"); ok {
		triggerVal := s.D.Get("activate_iam_user_sync_configuration_trigger")
		if iamUserSyncConfigurationDetails, ok := s.D.GetOkExists("iam_user_sync_configuration_details"); ok && triggerVal.(string) == "true" {
			if tmpList := iamUserSyncConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "iam_user_sync_configuration_details", 0)
				tmp, err := s.mapToiamUserSyncConfigurationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.IamUserSyncConfigurationDetails = &tmp
			}
		}
	}

	if identityDomainId, ok := s.D.GetOkExists("identity_domain_id"); ok {
		tmp := identityDomainId.(string)
		request.IdentityDomainId = &tmp
	}

	if _, ok := s.D.GetOkExists("activate_upst_configuration_trigger"); ok {
		triggerVal := s.D.Get("activate_upst_configuration_trigger")
		if upstConfigurationDetails, ok := s.D.GetOkExists("upst_configuration_details"); ok && triggerVal.(string) == "true" {
			if tmpList := upstConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "upst_configuration_details", 0)
				tmp, err := s.mapToupstConfigurationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.UpstConfigurationDetails = &tmp
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateIdentityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) getBdsInstanceIdentityConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {
	// Wait until it finishes
	bdsInstanceIdentityConfigurationId, err := bdsInstanceIdentityConfigurationWaitForWorkRequest(workId, "bdsidentityconfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceIdentityConfigurationId)
	return s.Get()
}

func bdsInstanceIdentityConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceIdentityConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceIdentityConfigurationWorkRequestShouldRetryFunc(timeout)
	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceIdentityConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}
	return identifier, nil
}

func getErrorFromBdsBdsInstanceIdentityConfigurationWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) Get() error {
	request := oci_bds.GetIdentityConfigurationRequest{}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsInstanceId, identityConfigurationId, err := parseBdsInstanceIdentityConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.IdentityConfigurationId = &identityConfigurationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetIdentityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityConfiguration
	return nil
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) Update() error {
	request := oci_bds.UpdateIdentityConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if iamUserSyncConfigurationDetails, ok := s.D.GetOkExists("iam_user_sync_configuration_details"); ok {
		if tmpList := iamUserSyncConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "iam_user_sync_configuration_details", 0)
			tmp, err := s.mapToiamUserSyncConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IamUserSyncConfigurationDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if upstConfigurationDetails, ok := s.D.GetOkExists("upst_configuration_details"); ok {
		if tmpList := upstConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "upst_configuration_details", 0)
			tmp, err := s.mapToupstConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpstConfigurationDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateIdentityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) Delete() error {
	request := oci_bds.DeleteIdentityConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if trigger, ok := s.D.GetOkExists("activate_iam_user_sync_configuration_trigger"); ok {
		if trigger.(string) == "true" {
			err := s.DeactivateIamUserSyncConfiguration()
			if err != nil {
				return err
			}
		}
	}

	if trigger, ok := s.D.GetOkExists("activate_upst_configuration_trigger"); ok {
		if trigger.(string) == "true" {
			err := s.DeactivateUpstConfiguration()

			if err != nil {
				return err
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteIdentityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceIdentityConfigurationWaitForWorkRequest(workId, "bdsidentityconfig",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) SetData() error {
	bdsInstanceId, identityConfigurationId, err := parseBdsInstanceIdentityConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(identityConfigurationId)

	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ConfidentialApplicationId != nil {
		s.D.Set("confidential_application_id", *s.Res.ConfidentialApplicationId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IamUserSyncConfiguration != nil {
		s.D.Set("iam_user_sync_configuration", []interface{}{IamUserSyncConfigurationToMap(s.Res.IamUserSyncConfiguration)})
	} else {
		s.D.Set("iam_user_sync_configuration", nil)
	}

	if s.Res.IdentityDomainId != nil {
		s.D.Set("identity_domain_id", *s.Res.IdentityDomainId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpstConfiguration != nil {
		s.D.Set("upst_configuration", []interface{}{UpstConfigurationToMap(s.Res.UpstConfiguration)})
	} else {
		s.D.Set("upst_configuration", nil)
	}
	return nil
}

func GetBdsInstanceIdentityConfigurationCompositeId(bdsInstanceId string, identityConfigurationId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	identityConfigurationId = url.PathEscape(identityConfigurationId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/identityConfigurations/" + identityConfigurationId
	return compositeId
}

func parseBdsInstanceIdentityConfigurationCompositeId(compositeId string) (bdsInstanceId string, identityConfigurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/identityConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	identityConfigurationId, _ = url.PathUnescape(parts[3])
	return
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) ActivateIamUserSyncConfiguration() error {
	request := oci_bds.ActivateIamUserSyncConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if isPosixAttributesAdditionRequired, ok := s.D.GetOkExists("iam_user_sync_configuration_details.0.is_posix_attributes_addition_required"); ok {
		tmp := isPosixAttributesAdditionRequired.(bool)
		request.IsPosixAttributesAdditionRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ActivateIamUserSyncConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	val := s.D.Get("activate_iam_user_sync_configuration_trigger")
	s.D.Set("activate_iam_user_sync_configuration_trigger", val)
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))

}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) ActivateUpstConfiguration() error {
	request := oci_bds.ActivateUpstConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if masterEncryptionKeyId, ok := s.D.GetOkExists("upst_configuration_details.0.master_encryption_key_id"); ok {
		tmp := masterEncryptionKeyId.(string)
		request.MasterEncryptionKeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists("upst_configuration_details.0.vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ActivateUpstConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	val := s.D.Get("activate_upst_configuration_trigger")
	s.D.Set("activate_upst_configuration_trigger", val)
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))

}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) DeactivateIamUserSyncConfiguration() error {
	request := oci_bds.DeactivateIamUserSyncConfigurationRequest{}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeactivateIamUserSyncConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	workId := response.OpcWorkRequestId
	val := s.D.Get("activate_iam_user_sync_configuration_trigger")
	s.D.Set("activate_iam_user_sync_configuration_trigger", val)
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))

}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) DeactivateUpstConfiguration() error {
	request := oci_bds.DeactivateUpstConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeactivateUpstConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	workId := response.OpcWorkRequestId
	val := s.D.Get("activate_upst_configuration_trigger")
	s.D.Set("activate_upst_configuration_trigger", val)
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))

}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) RefreshConfidentialApplication() error {
	request := oci_bds.RefreshConfidentialApplicationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RefreshConfidentialApplication(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	workId := response.OpcWorkRequestId

	//s.D.Set("refresh_confidential_application_trigger", "false")
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))

}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) RefreshUpstTokenExchangeKeytab() error {
	request := oci_bds.RefreshUpstTokenExchangeKeytabRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	tmp := s.D.Id()
	request.IdentityConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RefreshUpstTokenExchangeKeytab(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
	workId := response.OpcWorkRequestId

	//s.D.Set("refresh_upst_token_exchange_keytab_trigger", "false")
	return s.getBdsInstanceIdentityConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutCreate))

}

func IamUserSyncConfigurationToMap(obj *oci_bds.IamUserSyncConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPosixAttributesAdditionRequired != nil {
		result["is_posix_attributes_addition_required"] = bool(*obj.IsPosixAttributesAdditionRequired)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}
	return result
}

func UpstConfigurationToMap(obj *oci_bds.UpstConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeytabContent != nil {
		result["keytab_content"] = string(*obj.KeytabContent)
	}

	if obj.MasterEncryptionKeyId != nil {
		result["master_encryption_key_id"] = string(*obj.MasterEncryptionKeyId)
	}

	if obj.SecretId != nil {
		result["secret_id"] = string(*obj.SecretId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeTokenExchangeKeytabLastRefreshed != nil {
		result["time_token_exchange_keytab_last_refreshed"] = obj.TimeTokenExchangeKeytabLastRefreshed.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TokenExchangePrincipalName != nil {
		result["token_exchange_principal_name"] = string(*obj.TokenExchangePrincipalName)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}
	return result
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) mapToiamUserSyncConfigurationDetails(fieldKeyFormat string) (oci_bds.IamUserSyncConfigurationDetails, error) {
	result := oci_bds.IamUserSyncConfigurationDetails{}

	if isPosixAttributesAdditionRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_posix_attributes_addition_required")); ok {
		tmp := isPosixAttributesAdditionRequired.(bool)
		result.IsPosixAttributesAdditionRequired = &tmp
	}
	return result, nil
}

func iamUserSyncConfigurationDetailsToMap(obj *oci_bds.IamUserSyncConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPosixAttributesAdditionRequired != nil {
		result["is_posix_attributes_addition_required"] = bool(*obj.IsPosixAttributesAdditionRequired)
	}
	return result
}

func (s *BdsBdsInstanceIdentityConfigurationResourceCrud) mapToupstConfigurationDetails(fieldKeyFormat string) (oci_bds.UpstConfigurationDetails, error) {
	result := oci_bds.UpstConfigurationDetails{}

	if masterEncryptionKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "master_encryption_key_id")); ok {
		tmp := masterEncryptionKeyId.(string)
		result.MasterEncryptionKeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}
	return result, nil
}

func upstConfigurationDetailsToMap(obj *oci_bds.UpstConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MasterEncryptionKeyId != nil {
		result["master_encryption_key_id"] = string(*obj.MasterEncryptionKeyId)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}
	return result
}
