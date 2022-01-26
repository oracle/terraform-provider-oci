// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v56/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ServiceCatalogPrivateApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceCatalogPrivateApplication,
		Read:     readServiceCatalogPrivateApplication,
		Update:   updateServiceCatalogPrivateApplication,
		Delete:   deleteServiceCatalogPrivateApplication,
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
			"package_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"package_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"STACK",
							}, true),
						},
						"version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"zip_file_base64encoded": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"short_description": {
				Type:     schema.TypeString,
				Required: true,
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
			"logo_file_base64encoded": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"logo": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"content_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mime_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"package_type": {
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

func createServiceCatalogPrivateApplication(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceCatalogPrivateApplication(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateServiceCatalogPrivateApplication(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceCatalogPrivateApplication(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceCatalogPrivateApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_catalog.ServiceCatalogClient
	Res                    *oci_service_catalog.PrivateApplication
	DisableNotFoundRetries bool
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_service_catalog.PrivateApplicationLifecycleStateCreating),
	}
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_catalog.PrivateApplicationLifecycleStateActive),
	}
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_service_catalog.PrivateApplicationLifecycleStateDeleting),
	}
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_catalog.PrivateApplicationLifecycleStateDeleted),
	}
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) Create() error {
	request := oci_service_catalog.CreatePrivateApplicationRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logoFileBase64Encoded, ok := s.D.GetOkExists("logo_file_base64encoded"); ok {
		tmp := logoFileBase64Encoded.(string)
		request.LogoFileBase64Encoded = &tmp
	}

	if longDescription, ok := s.D.GetOkExists("long_description"); ok {
		tmp := longDescription.(string)
		request.LongDescription = &tmp
	}

	if packageDetails, ok := s.D.GetOkExists("package_details"); ok {
		if tmpList := packageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "package_details", 0)
			tmp, err := s.mapToCreatePrivateApplicationPackage(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PackageDetails = tmp
		}
	}

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.CreatePrivateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateApplication
	return nil
	//workId := response.OpcWorkRequestId
	//return s.getPrivateApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog"), oci_service_catalog.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) getPrivateApplicationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_service_catalog.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	privateApplicationId, err := privateApplicationWaitForWorkRequest(workId, "service_catalog",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		//log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, privateApplicationId)
		//_, cancelErr := s.Client.CancelWorkRequestRequest(context.Background(),
		//	oci_service_catalog.CancelWorkRequestRequest{
		//		WorkRequestId: workId,
		//		RequestMetadata: oci_common.RequestMetadata{
		//			RetryPolicy: retryPolicy,
		//		},
		//	})
		//if cancelErr != nil {
		//	log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		//}
		return err
	}
	s.D.SetId(*privateApplicationId)

	return s.Get()
}

func privateApplicationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "service_catalog", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_service_catalog.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func privateApplicationWaitForWorkRequest(wId *string, entityType string, action oci_service_catalog.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_service_catalog.ServiceCatalogClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "service_catalog")
	retryPolicy.ShouldRetryOperation = privateApplicationWorkRequestShouldRetryFunc(timeout)

	response := oci_service_catalog.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{},
		Target: []string{
			string(oci_service_catalog.ActionTypeCreated),
			string(oci_service_catalog.ActionTypeFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_service_catalog.GetWorkRequestRequest{
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
				identifier = res.EntityId
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_service_catalog.OperationStatusFailed {
		return nil, getErrorFromServiceCatalogPrivateApplicationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceCatalogPrivateApplicationWorkRequest(client *oci_service_catalog.ServiceCatalogClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_service_catalog.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_service_catalog.ListWorkRequestErrorsRequest{
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

func (s *ServiceCatalogPrivateApplicationResourceCrud) Get() error {
	request := oci_service_catalog.GetPrivateApplicationRequest{}

	tmp := s.D.Id()
	request.PrivateApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.GetPrivateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateApplication
	return nil
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_catalog.UpdatePrivateApplicationRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logoFileBase64Encoded, ok := s.D.GetOkExists("logo_file_base64encoded"); ok {
		tmp := logoFileBase64Encoded.(string)
		request.LogoFileBase64Encoded = &tmp
	}

	if longDescription, ok := s.D.GetOkExists("long_description"); ok {
		tmp := longDescription.(string)
		request.LongDescription = &tmp
	}

	tmp := s.D.Id()
	request.PrivateApplicationId = &tmp

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.UpdatePrivateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateApplication
	return nil

	//workId := response.OpcWorkRequestId
	//return s.getPrivateApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog"), oci_service_catalog.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) Delete() error {
	request := oci_service_catalog.DeletePrivateApplicationRequest{}

	tmp := s.D.Id()
	request.PrivateApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	_, err := s.Client.DeletePrivateApplication(context.Background(), request)
	return err

	//if err != nil {
	//	return err
	//}
	//
	//workId := response.OpcWorkRequestId
	//// Wait until it finishes
	//_, delWorkRequestErr := privateApplicationWaitForWorkRequest(workId, "service_catalog",
	//	oci_service_catalog.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	//return delWorkRequestErr
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) SetData() error {
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

	if s.Res.Logo != nil {
		s.D.Set("logo", []interface{}{SCUploadDataToMap(s.Res.Logo)})
	} else {
		s.D.Set("logo", nil)
	}

	if s.Res.LongDescription != nil {
		s.D.Set("long_description", *s.Res.LongDescription)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) mapToCreatePrivateApplicationPackage(fieldKeyFormat string) (oci_service_catalog.CreatePrivateApplicationPackage, error) {
	var baseObject oci_service_catalog.CreatePrivateApplicationPackage
	//discriminator
	packageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_type"))
	var packageType string
	if ok {
		packageType = packageTypeRaw.(string)
	} else {
		packageType = "" // default value
	}
	switch strings.ToLower(packageType) {
	case strings.ToLower("STACK"):
		details := oci_service_catalog.CreatePrivateApplicationStackPackage{}
		if zipFileBase64Encoded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "zip_file_base64encoded")); ok {
			tmp := zipFileBase64Encoded.(string)
			details.ZipFileBase64Encoded = &tmp
		}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown package_type '%v' was specified", packageType)
	}
	return baseObject, nil
}

func CreatePrivateApplicationPackageToMap(obj *oci_service_catalog.CreatePrivateApplicationPackage) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_service_catalog.CreatePrivateApplicationStackPackage:
		result["package_type"] = "STACK"

		if v.ZipFileBase64Encoded != nil {
			result["zip_file_base64encoded"] = string(*v.ZipFileBase64Encoded)
		}
	default:
		log.Printf("[WARN] Received 'package_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func PrivateApplicationSummaryToMap(obj oci_service_catalog.PrivateApplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Logo != nil {
		result["logo"] = []interface{}{SCUploadDataToMap(obj.Logo)}
	}

	result["package_type"] = string(obj.PackageType)

	if obj.ShortDescription != nil {
		result["short_description"] = string(*obj.ShortDescription)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func SCUploadDataToMap(obj *oci_service_catalog.UploadData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentUrl != nil {
		result["content_url"] = string(*obj.ContentUrl)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MimeType != nil {
		result["mime_type"] = string(*obj.MimeType)
	}

	return result
}

func (s *ServiceCatalogPrivateApplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_catalog.ChangePrivateApplicationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PrivateApplicationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	_, err := s.Client.ChangePrivateApplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
	//workId := response.OpcWorkRequestId
	//return s.getPrivateApplicationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog"), oci_service_catalog.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
