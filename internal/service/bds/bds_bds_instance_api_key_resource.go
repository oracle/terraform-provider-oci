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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BdsBdsInstanceApiKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceApiKey,
		Update:   updateBdsBdsInstanceApiKey,
		Read:     readBdsBdsInstanceApiKey,
		Delete:   deleteBdsBdsInstanceApiKey,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_alias": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},
			"passphrase": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"default_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pemfilepath": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
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

func createBdsBdsInstanceApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func updateBdsBdsInstanceApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func readBdsBdsInstanceApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func deleteBdsBdsInstanceApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceApiKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsApiKey
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceApiKeyResourceCrud) ID() string {
	return GetBdsInstanceApiKeyCompositeId(*s.Res.Id, s.D.Get("bds_instance_id").(string))
}

func (s *BdsBdsInstanceApiKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsApiKeyLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceApiKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsApiKeyLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceApiKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsApiKeyLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceApiKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsApiKeyLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceApiKeyResourceCrud) Create() error {
	request := oci_bds.CreateBdsApiKeyRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if defaultRegion, ok := s.D.GetOkExists("default_region"); ok {
		tmp := defaultRegion.(string)
		request.DefaultRegion = &tmp
	}

	if keyAlias, ok := s.D.GetOkExists("key_alias"); ok {
		tmp := keyAlias.(string)
		request.KeyAlias = &tmp
	}

	if passphrase, ok := s.D.GetOkExists("passphrase"); ok {
		tmp := passphrase.(string)
		request.Passphrase = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getBdsInstanceApiKeyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceApiKeyResourceCrud) setIdFromWorkRequest(workId *string) {
	var identifier_str string
	var identifier *string
	var err error

	workRequestResponse := oci_bds.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
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
		apiKeyId, _, err := parseBdsInstanceApiKeyCompositeId(*identifier)
		if err == nil && apiKeyId != "" {
			s.D.SetId(apiKeyId)
		}
	}
}

func (s *BdsBdsInstanceApiKeyResourceCrud) getBdsInstanceApiKeyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceApiKeyId, err := bdsInstanceApiKeyWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceApiKeyId)

	return s.Get()
}

func bdsInstanceApiKeyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsInstanceApiKeyWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceApiKeyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromBdsBdsInstanceApiKeyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceApiKeyWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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

func (s *BdsBdsInstanceApiKeyResourceCrud) Get() error {
	request := oci_bds.GetBdsApiKeyRequest{}

	tmp := s.D.Id()
	request.ApiKeyId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	apiKeyId, bdsInstanceId, err := parseBdsInstanceApiKeyCompositeId(s.D.Id())
	if err == nil {
		request.ApiKeyId = &apiKeyId
		request.BdsInstanceId = &bdsInstanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsApiKey
	return nil
}

func (s *BdsBdsInstanceApiKeyResourceCrud) Delete() error {
	request := oci_bds.DeleteBdsApiKeyRequest{}

	tmp := s.D.Id()
	request.ApiKeyId = &tmp

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteBdsApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceApiKeyWaitForWorkRequest(workId, "bds",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceApiKeyResourceCrud) SetData() error {

	apiKeyId, bdsInstanceId, err := parseBdsInstanceApiKeyCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(apiKeyId)
		s.D.Set("bds_instance_id", &bdsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DefaultRegion != nil {
		s.D.Set("default_region", *s.Res.DefaultRegion)
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	if s.Res.KeyAlias != nil {
		s.D.Set("key_alias", *s.Res.KeyAlias)
	}

	if s.Res.Pemfilepath != nil {
		s.D.Set("pemfilepath", *s.Res.Pemfilepath)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}

func GetBdsInstanceApiKeyCompositeId(apiKeyId string, bdsInstanceId string) string {
	apiKeyId = url.PathEscape(apiKeyId)
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/apiKeys/" + apiKeyId
	return compositeId
}

func parseBdsInstanceApiKeyCompositeId(compositeId string) (apiKeyId string, bdsInstanceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/apiKeys/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	apiKeyId, _ = url.PathUnescape(parts[3])

	return
}
