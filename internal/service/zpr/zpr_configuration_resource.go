// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package zpr

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ZprConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
				// d.Id() here is the last argument passed to the `terraform import RESOURCE_TYPE.RESOURCE_NAME RESOURCE_ID` command
				// Here we use a function to parse the import ID (like the example above) to simplify our logic
				compartmentId, configurationId, err := parseResourceID(d.Id())

				if err != nil {
					return nil, err
				}

				d.Set("compartment_id", compartmentId)
				d.SetId(configurationId)

				return []*schema.ResourceData{d}, nil
			},
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createZprConfiguration,
		Read:     readZprConfiguration,
		Delete:   deleteZprConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"zpr_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
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

func parseResourceID(id string) (string, string, error) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected compartmentId/configurationId", id)
	}

	return parts[0], parts[1], nil
}

func createZprConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ZprConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.CreateResource(d, sync)
}

func readZprConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ZprConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.ReadResource(sync)
}

func deleteZprConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ZprConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_zpr.ZprClient
	Res                    *oci_zpr.Configuration
	DisableNotFoundRetries bool
}

func (s *ZprConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ZprConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_zpr.ConfigurationLifecycleStateCreating),
	}
}

func (s *ZprConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_zpr.ConfigurationLifecycleStateActive),
	}
}

func (s *ZprConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_zpr.ConfigurationLifecycleStateDeleting),
	}
}

func (s *ZprConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_zpr.ConfigurationLifecycleStateDeleted),
	}
}

func (s *ZprConfigurationResourceCrud) Create() error {
	request := oci_zpr.CreateConfigurationRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if zprStatus, ok := s.D.GetOkExists("zpr_status"); ok {
		request.ZprStatus = oci_zpr.ConfigurationZprStatusEnum(zprStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.CreateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_zpr.GetZprConfigurationWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetZprConfigurationWorkRequest(context.Background(),
		oci_zpr.GetZprConfigurationWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "configuration") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr"), oci_zpr.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ZprConfigurationResourceCrud) getConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_zpr.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	configurationId, err := configurationWaitForWorkRequest(workId, "configuration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*configurationId)

	return s.Get()
}

func configurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "zpr", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_zpr.GetZprConfigurationWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func configurationWaitForWorkRequest(wId *string, entityType string, action oci_zpr.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_zpr.ZprClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "zpr")
	retryPolicy.ShouldRetryOperation = configurationWorkRequestShouldRetryFunc(timeout)

	response := oci_zpr.GetZprConfigurationWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_zpr.WorkRequestStatusInProgress),
			string(oci_zpr.WorkRequestStatusAccepted),
			string(oci_zpr.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_zpr.WorkRequestStatusSucceeded),
			string(oci_zpr.WorkRequestStatusFailed),
			string(oci_zpr.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetZprConfigurationWorkRequest(context.Background(),
				oci_zpr.GetZprConfigurationWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_zpr.WorkRequestStatusFailed || response.Status == oci_zpr.WorkRequestStatusCanceled {
		return nil, getErrorFromZprConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromZprConfigurationWorkRequest(client *oci_zpr.ZprClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_zpr.ActionTypeEnum) error {
	response, err := client.ListZprConfigurationWorkRequestErrors(context.Background(),
		oci_zpr.ListZprConfigurationWorkRequestErrorsRequest{
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

func (s *ZprConfigurationResourceCrud) Get() error {
	request := oci_zpr.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "zpr")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *ZprConfigurationResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("zpr_status", s.Res.ZprStatus)

	return nil
}
