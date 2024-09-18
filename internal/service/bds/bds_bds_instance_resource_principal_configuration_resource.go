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

func BdsBdsInstanceResourcePrincipalConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceResourcePrincipalConfiguration,
		Read:     readBdsBdsInstanceResourcePrincipalConfiguration,
		Update:   updateBdsBdsInstanceResourcePrincipalConfiguration,
		Delete:   deleteBdsBdsInstanceResourcePrincipalConfiguration,
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
				ForceNew:  true,
				Sensitive: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"session_token_life_span_duration_in_hours": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"force_refresh_resource_principal_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_token_expiry": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_token_refreshed": {
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

func createBdsBdsInstanceResourcePrincipalConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	return nil

}

func readBdsBdsInstanceResourcePrincipalConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstanceResourcePrincipalConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if _, ok := sync.D.GetOkExists("force_refresh_resource_principal_trigger"); ok && sync.D.HasChange("force_refresh_resource_principal_trigger") {
		oldRaw, newRaw := sync.D.GetChange("force_refresh_resource_principal_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ForceRefreshResourcePrincipal()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("force_refresh_resource_principal_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteBdsBdsInstanceResourcePrincipalConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceResourcePrincipalConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.ResourcePrincipalConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) ID() string {
	return GetBdsInstanceResourcePrincipalConfigurationCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.ResourcePrincipalConfigurationLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.ResourcePrincipalConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.ResourcePrincipalConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.ResourcePrincipalConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) Create() error {
	request := oci_bds.CreateResourcePrincipalConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sessionTokenLifeSpanDurationInHours, ok := s.D.GetOkExists("session_token_life_span_duration_in_hours"); ok {
		tmp := sessionTokenLifeSpanDurationInHours.(int)
		request.SessionTokenLifeSpanDurationInHours = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateResourcePrincipalConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceResourcePrincipalConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) getBdsInstanceResourcePrincipalConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceResourcePrincipalConfigurationId, err := bdsInstanceResourcePrincipalConfigurationWaitForWorkRequest(workId, "resourcePrincipalConfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceResourcePrincipalConfigurationId)

	return s.Get()
}

func bdsInstanceResourcePrincipalConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceResourcePrincipalConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceResourcePrincipalConfigurationWorkRequestShouldRetryFunc(timeout)

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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceResourcePrincipalConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceResourcePrincipalConfigurationWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) Get() error {
	request := oci_bds.GetResourcePrincipalConfigurationRequest{}

	tmp := s.D.Id()
	request.ResourcePrincipalConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsInstanceId, resourcePrincipalConfigurationId, err := parseBdsInstanceResourcePrincipalConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.ResourcePrincipalConfigurationId = &resourcePrincipalConfigurationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetResourcePrincipalConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourcePrincipalConfiguration
	return nil

}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) Update() error {
	request := oci_bds.UpdateResourcePrincipalConfigurationRequest{}

	tmp := s.D.Id()
	request.ResourcePrincipalConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sessionTokenLifeSpanDurationInHours, ok := s.D.GetOkExists("session_token_life_span_duration_in_hours"); ok {
		tmp := sessionTokenLifeSpanDurationInHours.(int)
		request.SessionTokenLifeSpanDurationInHours = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateResourcePrincipalConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceResourcePrincipalConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesInProgress, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) SetData() error {

	bdsInstanceId, resourcePrincipalConfigurationId, err := parseBdsInstanceResourcePrincipalConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(resourcePrincipalConfigurationId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BdsInstanceId != nil {
		s.D.Set("bds_instance_id", *s.Res.BdsInstanceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.SessionTokenLifeSpanDurationInHours != nil {
		s.D.Set("session_token_life_span_duration_in_hours", *s.Res.SessionTokenLifeSpanDurationInHours)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTokenExpiry != nil {
		s.D.Set("time_token_expiry", s.Res.TimeTokenExpiry.String())
	}

	if s.Res.TimeTokenRefreshed != nil {
		s.D.Set("time_token_refreshed", s.Res.TimeTokenRefreshed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetBdsInstanceResourcePrincipalConfigurationCompositeId(bdsInstanceId string, resourcePrincipalConfigurationId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	resourcePrincipalConfigurationId = url.PathEscape(resourcePrincipalConfigurationId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/resourcePrincipalConfigurations/" + resourcePrincipalConfigurationId
	return compositeId
}

func parseBdsInstanceResourcePrincipalConfigurationCompositeId(compositeId string) (bdsInstanceId string, resourcePrincipalConfigurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/resourcePrincipalConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	resourcePrincipalConfigurationId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) ForceRefreshResourcePrincipal() error {
	request := oci_bds.ForceRefreshResourcePrincipalRequest{}

	tmp := s.D.Id()
	request.ResourcePrincipalConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ForceRefreshResourcePrincipal(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceResourcePrincipalConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationResourceCrud) Delete() error {
	request := oci_bds.RemoveResourcePrincipalConfigurationRequest{}

	tmp := s.D.Id()
	request.ResourcePrincipalConfigurationId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RemoveResourcePrincipalConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceResourcePrincipalConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
