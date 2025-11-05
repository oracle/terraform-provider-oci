// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

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
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaPrivateServiceAccessResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createPsaPrivateServiceAccesWithContext,
		ReadContext:   readPsaPrivateServiceAccesWithContext,
		UpdateContext: updatePsaPrivateServiceAccesWithContext,
		DeleteContext: deletePsaPrivateServiceAccesWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_id": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ipv4ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"fqdns": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createPsaPrivateServiceAccesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccesResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readPsaPrivateServiceAccesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccesResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updatePsaPrivateServiceAccesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccesResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deletePsaPrivateServiceAccesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccesResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type PsaPrivateServiceAccesResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_psa.PrivateServiceAccessClient
	Res                    *oci_psa.PrivateServiceAccess
	DisableNotFoundRetries bool
}

func (s *PsaPrivateServiceAccesResourceCrud) ID() string {
	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		return privateServiceAccessId.(string)
	} else {
		return s.D.Id()
	}
}

func (s *PsaPrivateServiceAccesResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_psa.PrivateServiceAccessLifecycleStateCreating),
	}
}

func (s *PsaPrivateServiceAccesResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_psa.PrivateServiceAccessLifecycleStateActive),
	}
}

func (s *PsaPrivateServiceAccesResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_psa.PrivateServiceAccessLifecycleStateDeleting),
	}
}

func (s *PsaPrivateServiceAccesResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_psa.PrivateServiceAccessLifecycleStateDeleted),
	}
}

func (s *PsaPrivateServiceAccesResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_psa.CreatePrivateServiceAccessRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipv4Ip, ok := s.D.GetOkExists("ipv4ip"); ok {
		tmp := ipv4Ip.(string)
		request.Ipv4Ip = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if serviceId, ok := s.D.GetOkExists("service_id"); ok {
		tmp := serviceId.(string)
		request.ServiceId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa")

	response, err := s.Client.CreatePrivateServiceAccess(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivateServiceAccesFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa"), oci_psa.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *PsaPrivateServiceAccesResourceCrud) getPrivateServiceAccesFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_psa.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privateServiceAccesId, err := privateServiceAccesWaitForWorkRequest(ctx, workId, "privateserviceacces",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*privateServiceAccesId)

	return s.GetWithContext(ctx)
}

func privateServiceAccesWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "psa", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_psa.GetPsaWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func privateServiceAccesWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_psa.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_psa.PrivateServiceAccessClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "psa")
	retryPolicy.ShouldRetryOperation = privateServiceAccesWorkRequestShouldRetryFunc(timeout)

	response := oci_psa.GetPsaWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_psa.OperationStatusInProgress),
			string(oci_psa.OperationStatusAccepted),
			string(oci_psa.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_psa.OperationStatusSucceeded),
			string(oci_psa.OperationStatusFailed),
			string(oci_psa.OperationStatusCancelled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetPsaWorkRequest(ctx,
				oci_psa.GetPsaWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_psa.OperationStatusFailed || response.Status == oci_psa.OperationStatusCancelled {
		return nil, getErrorFromPsaPrivateServiceAccesWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromPsaPrivateServiceAccesWorkRequest(ctx context.Context, client *oci_psa.PrivateServiceAccessClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_psa.ActionTypeEnum) error {
	response, err := client.ListPsaWorkRequestErrors(ctx,
		oci_psa.ListPsaWorkRequestErrorsRequest{
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

func (s *PsaPrivateServiceAccesResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psa.GetPrivateServiceAccessRequest{}

	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		tmp := privateServiceAccessId.(string)
		request.PrivateServiceAccessId = &tmp
	} else {
		privateServiceAccessId := s.D.Id()
		request.PrivateServiceAccessId = &privateServiceAccessId
	}

	/*privateServiceAccessId, err := parsePrivateServiceAccesCompositeId(s.D.Id())
	if err == nil {
		request.PrivateServiceAccessId = &privateServiceAccessId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}
	*/
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa")

	response, err := s.Client.GetPrivateServiceAccess(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateServiceAccess
	return nil
}

func (s *PsaPrivateServiceAccesResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_psa.UpdatePrivateServiceAccessRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		tmp := privateServiceAccessId.(string)
		request.PrivateServiceAccessId = &tmp
	} else {
		privateServiceAccessId := s.D.Id()
		request.PrivateServiceAccessId = &privateServiceAccessId
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa")

	response, err := s.Client.UpdatePrivateServiceAccess(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivateServiceAccesFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa"), oci_psa.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *PsaPrivateServiceAccesResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_psa.DeletePrivateServiceAccessRequest{}

	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		tmp := privateServiceAccessId.(string)
		request.PrivateServiceAccessId = &tmp
	} else {
		privateServiceAccessId := s.D.Id()
		request.PrivateServiceAccessId = &privateServiceAccessId
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa")

	response, err := s.Client.DeletePrivateServiceAccess(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := privateServiceAccesWaitForWorkRequest(ctx, workId, "privateserviceacces",
		oci_psa.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *PsaPrivateServiceAccesResourceCrud) SetData() error {

	privateServiceAccessId, err := parsePrivateServiceAccesCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("private_service_access_id", &privateServiceAccessId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("fqdns", s.Res.Fqdns)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Ipv4Ip != nil {
		s.D.Set("ipv4ip", *s.Res.Ipv4Ip)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	if s.Res.ServiceId != nil {
		s.D.Set("service_id", *s.Res.ServiceId)
	}

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

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}

func GetPrivateServiceAccesCompositeId(privateServiceAccessId string) string {
	privateServiceAccessId = url.PathEscape(privateServiceAccessId)
	compositeId := "privateServiceAccess/" + privateServiceAccessId
	return compositeId
}

func parsePrivateServiceAccesCompositeId(compositeId string) (privateServiceAccessId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("privateServiceAccess/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	privateServiceAccessId, _ = url.PathUnescape(parts[1])

	return
}

func PrivateServiceAccessSummaryToMap(obj oci_psa.PrivateServiceAccessSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["fqdns"] = obj.Fqdns

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Ipv4Ip != nil {
		result["ipv4ip"] = string(*obj.Ipv4Ip)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SecurityAttributes != nil {
		result["security_attributes"] = tfresource.SecurityAttributesToMap(obj.SecurityAttributes)
	}

	if obj.ServiceId != nil {
		result["service_id"] = string(*obj.ServiceId)
	}

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

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	if obj.VnicId != nil {
		result["vnic_id"] = string(*obj.VnicId)
	}

	return result
}

func (s *PsaPrivateServiceAccesResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_psa.ChangePrivateServiceAccessCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		tmp := privateServiceAccessId.(string)
		changeCompartmentRequest.PrivateServiceAccessId = &tmp
	} else {
		privateServiceAccessId := s.D.Id()
		changeCompartmentRequest.PrivateServiceAccessId = &privateServiceAccessId
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa")

	response, err := s.Client.ChangePrivateServiceAccessCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPrivateServiceAccesFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psa"), oci_psa.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
