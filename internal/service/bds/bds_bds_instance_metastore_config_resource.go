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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BdsBdsInstanceMetastoreConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBdsBdsInstanceMetastoreConfigWithContext,
		ReadContext:   readBdsBdsInstanceMetastoreConfigWithContext,
		UpdateContext: updateBdsBdsInstanceMetastoreConfigWithContext,
		DeleteContext: deleteBdsBdsInstanceMetastoreConfigWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_api_key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bds_api_key_passphrase": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
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
			"metastore_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Optional
			"activate_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"metastore_type": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBdsBdsInstanceMetastoreConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceMetastoreConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBdsBdsInstanceMetastoreConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceMetastoreConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBdsBdsInstanceMetastoreConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceMetastoreConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBdsBdsInstanceMetastoreConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceMetastoreConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BdsBdsInstanceMetastoreConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsMetastoreConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) ID() string {
	return GetBdsInstanceMetastoreConfigCompositeId(s.D.Get("bds_instance_id").(string), *s.Res.Id)
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateCreating),
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateActivating),
	}
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_bds.CreateBdsMetastoreConfigurationRequest{}

	if bdsApiKeyId, ok := s.D.GetOkExists("bds_api_key_id"); ok {
		tmp := bdsApiKeyId.(string)
		request.BdsApiKeyId = &tmp
	}

	if bdsApiKeyPassphrase, ok := s.D.GetOkExists("bds_api_key_passphrase"); ok {
		tmp := bdsApiKeyPassphrase.(string)
		request.BdsApiKeyPassphrase = &tmp
	}

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

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsMetastoreConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(ctx, workId)
	return s.getBdsInstanceMetastoreConfigFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) setIdFromWorkRequest(ctx context.Context, workId *string) {
	var identifier_str string
	var identifier *string
	var err error
	workRequestResponse := oci_bds.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(ctx,
		oci_bds.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "bds") && res.EntityUri != nil {
				var pattern = `/\d*/`
				var regex = regexp.MustCompile(pattern)
				identifier_str = regex.ReplaceAllString(*res.EntityUri, "")
				identifier = &identifier_str
				break
			}
		}
	}
	if identifier != nil {
		_, metastoreConfigId, err := parseBdsInstanceMetastoreConfigCompositeId(*identifier)
		if err == nil && metastoreConfigId != "" {
			s.D.SetId(metastoreConfigId)
		}
	}
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) getBdsInstanceMetastoreConfigFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceMetastoreConfigId, err := bdsInstanceMetastoreConfigWaitForWorkRequest(ctx, workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceMetastoreConfigId)

	return s.GetWithContext(ctx)
}

func bdsInstanceMetastoreConfigWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceMetastoreConfigWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceMetastoreConfigWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			response, err = client.GetWorkRequest(ctx,
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	var identifier *string
	var identifier_str string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				var pattern = `/\d*/`
				var regex = regexp.MustCompile(pattern)
				identifier_str = regex.ReplaceAllString(*res.EntityUri, "")
				identifier = &identifier_str
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceMetastoreConfigWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceMetastoreConfigWorkRequest(ctx context.Context, client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.GetBdsMetastoreConfigurationRequest{}

	tmp := s.D.Id()
	request.MetastoreConfigId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsInstanceId, metastoreConfigId, err := parseBdsInstanceMetastoreConfigCompositeId(s.D.Id())
	if err == nil {
		request.BdsInstanceId = &bdsInstanceId
		request.MetastoreConfigId = &metastoreConfigId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsMetastoreConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsMetastoreConfiguration
	return nil
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) UpdateWithContext(ctx context.Context) error {

	if _, ok := s.D.GetOkExists("activate_trigger"); ok && s.D.HasChange("activate_trigger") {
		err := s.activateBdsInstanceMetastoreConfig(ctx)
		if err != nil {
			return err
		}
	}

	request := oci_bds.UpdateBdsMetastoreConfigurationRequest{}

	tmp := s.D.Id()
	request.MetastoreConfigId = &tmp

	if bdsApiKeyId, ok := s.D.GetOkExists("bds_api_key_id"); ok {
		tmp := bdsApiKeyId.(string)
		request.BdsApiKeyId = &tmp
	}

	if bdsApiKeyPassphrase, ok := s.D.GetOkExists("bds_api_key_passphrase"); ok {
		tmp := bdsApiKeyPassphrase.(string)
		request.BdsApiKeyPassphrase = &tmp
	}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateBdsMetastoreConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceMetastoreConfigFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_bds.DeleteBdsMetastoreConfigurationRequest{}

	tmp := s.D.Id()
	request.MetastoreConfigId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteBdsMetastoreConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceMetastoreConfigWaitForWorkRequest(ctx, workId, "bds",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) SetData() error {

	bdsInstanceId, metastoreConfigId, err := parseBdsInstanceMetastoreConfigCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
		s.D.SetId(metastoreConfigId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BdsApiKeyId != nil {
		s.D.Set("bds_api_key_id", *s.Res.BdsApiKeyId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	s.D.Set("metastore_type", s.Res.MetastoreType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetBdsInstanceMetastoreConfigCompositeId(bdsInstanceId string, metastoreConfigId string) string {
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	metastoreConfigId = url.PathEscape(metastoreConfigId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/metastoreConfigs/" + metastoreConfigId
	return compositeId
}

func parseBdsInstanceMetastoreConfigCompositeId(compositeId string) (bdsInstanceId string, metastoreConfigId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/metastoreConfigs/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	metastoreConfigId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceMetastoreConfigResourceCrud) activateBdsInstanceMetastoreConfig(ctx context.Context) error {

	request := oci_bds.ActivateBdsMetastoreConfigurationRequest{}

	tmp := s.D.Id()
	request.MetastoreConfigId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	metastoreType, ok := s.D.GetOkExists("metastore_type")
	if ok && metastoreType == "EXTERNAL" {
		if bdsApiKeyPassphrase, ok := s.D.GetOkExists("bds_api_key_passphrase"); ok {
			tmp := bdsApiKeyPassphrase.(string)
			request.BdsApiKeyPassphrase = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ActivateBdsMetastoreConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, updateWorkRequestErr := bdsInstanceMetastoreConfigWaitForWorkRequest(ctx, workId, "bds",
		oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if updateWorkRequestErr != nil {
		return updateWorkRequestErr
	}

	val := s.D.Get("activate_trigger")
	s.D.Set("activate_trigger", val)

	return nil
}
