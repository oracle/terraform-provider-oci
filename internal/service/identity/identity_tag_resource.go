// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

func IdentityTagResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("15m"),
			Update: tfresource.GetTimeoutDuration("15m"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createIdentityTag,
		Read:   readIdentityTag,
		Update: updateIdentityTag,
		Delete: deleteIdentityTag,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"tag_namespace_id": {
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
			"is_cost_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"validator": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"validator_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ENUM",
							}, true),
						},

						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
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
		},
	}
}

func createIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityTag(d *schema.ResourceData, m interface{}) error {
	// prevent tag deletion when testing, as its a time consuming and sequential operation permitted one per tenancy.
	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		return nil
	}

	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityTagResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Tag
	DisableNotFoundRetries bool
}

func (s *IdentityTagResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityTagResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IdentityTagResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.TagLifecycleStateActive),
	}
}

func (s *IdentityTagResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.TagLifecycleStateDeleting),
	}
}

func (s *IdentityTagResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.TagLifecycleStateDeleted),
	}
}

func (s *IdentityTagResourceCrud) Create() error {
	request := oci_identity.CreateTagRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isCostTracking, ok := s.D.GetOkExists("is_cost_tracking"); ok {
		tmp := isCostTracking.(bool)
		request.IsCostTracking = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	if validator, ok := s.D.GetOkExists("validator"); ok {
		if tmpList := validator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validator", 0)
			tmp, err := s.mapToBaseTagDefinitionValidator(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Validator = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	contextToUse := context.Background()
	response, err := s.Client.CreateTag(contextToUse, request)
	if err == nil {
		s.Res = &response.Tag
		s.D.SetId(*s.Res.Id)
		//is_retired field is currently not supported in Create so Update to make server state same as config
		if updateError := s.Update(); updateError != nil {
			return updateError
		}
		return nil
	}

	// Tag definitions can't be deleted, so this is a work around here to react to collisions by
	// basically importing that pre-existing namespace into this plan if tags_import_if_exists
	// flag is set to 'true'. This is ONLY for TESTING and should not be used elsewhere.
	// Use 'terraform import' for existing tag definitions
	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists && strings.Contains(err.Error(), "TagDefinitionAlreadyExists") {
		// List all tag definitions using the datasource to find that tag definition which matches
		s.D.Set("tag_namespace_id", request.TagNamespaceId)
		s.D.Set("name", request.Name)
		dsCrud := &IdentityTagsDataSourceCrud{s.D, s.Client, nil}
		if dsErr := dsCrud.Get(); dsErr != nil {
			//return original error when datasource call fails
			return err
		}

		for _, tag := range dsCrud.Res.Items {
			if strings.EqualFold(*tag.Name, *request.Name) {
				s.D.SetId(*tag.Id)
				if updateError := s.Update(); updateError != nil {
					//Update to tags can only be done from home region, so do get in that case
					if getError := s.Get(); getError != nil {
						return getError
					}
				}
				return nil
			}
		}
	}

	return err

}

func (s *IdentityTagResourceCrud) Get() error {
	request := oci_identity.GetTagRequest{}

	tagName, tagNamespaceId, parseTagCompositeIdErr := parseTagCompositeId(s.D.Id())
	if parseTagCompositeIdErr == nil {
		request.TagName = &tagName
		request.TagNamespaceId = &tagNamespaceId
	}

	if tagName, ok := s.D.GetOkExists("name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tag
	if parseTagCompositeIdErr == nil {
		// Import sets the ID to composite ID and hence overwriting ID to OCID from response
		id := response.Tag.Id
		if id == nil {
			return fmt.Errorf("error : received null value for id attribute for request %s, id attribute cannot be null", *response.OpcRequestId)
		}
		s.D.SetId(*id)
	}
	return nil
}

func (s *IdentityTagResourceCrud) Update() error {
	request := oci_identity.UpdateTagRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isCostTracking, ok := s.D.GetOkExists("is_cost_tracking"); ok {
		tmp := isCostTracking.(bool)
		request.IsCostTracking = &tmp
	}

	if isRetired, ok := s.D.GetOkExists("is_retired"); ok {
		tmp := isRetired.(bool)
		request.IsRetired = &tmp
	}

	if tagName, ok := s.D.GetOkExists("name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	if validator, ok := s.D.GetOkExists("validator"); ok {
		if tmpList := validator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validator", 0)
			tmp, err := s.mapToBaseTagDefinitionValidator(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Validator = tmp
		} else {
			// For case when updating ENUM type to DEFAULT: validator absent in config but s.D.GetOkExists("validator") returns true
			var baseObject oci_identity.BaseTagDefinitionValidator
			details := oci_identity.DefaultTagDefinitionValidator{}
			baseObject = details
			request.Validator = baseObject
		}
	} else {
		// For testing only- When Update() is called from Create() and there is no validator in config (Required Create)
		// remove the validator for an imported tag as Step 0 of test expects tag without validator
		importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
		if importIfExists {
			var baseObject oci_identity.BaseTagDefinitionValidator
			details := oci_identity.DefaultTagDefinitionValidator{}
			baseObject = details
			request.Validator = baseObject
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tag
	return nil
}

func (s *IdentityTagResourceCrud) Delete() error {
	// retire the tag if not already retired.
	if isRetired, ok := s.D.GetOkExists("is_retired"); ok && (!isRetired.(bool) || s.D.HasChange("is_retired")) {

		tmp := true
		s.D.Set("is_retired", &tmp)

		if err := s.Update(); err != nil {
			return err
		}
	}

	request := oci_identity.DeleteTagRequest{}

	if tagName, ok := s.D.GetOkExists("name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")
	response, err := s.Client.DeleteTag(context.Background(), request)

	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := IdentityTaggingWaitForWorkRequest(workId, "identity",
		oci_identity.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func IdentityTaggingWaitForWorkRequest(workRequestId *string, entityType string, action oci_identity.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_identity.IdentityClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "identity")
	retryPolicy.ShouldRetryOperation = identityTagWorkRequestShouldRetryFunc(timeout)
	response := oci_identity.GetTaggingWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_identity.TaggingWorkRequestStatusInProgress),
			string(oci_identity.TaggingWorkRequestStatusAccepted),
			string(oci_identity.TaggingWorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_identity.TaggingWorkRequestStatusSucceeded),
			string(oci_identity.TaggingWorkRequestStatusFailed),
			string(oci_identity.TaggingWorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetTaggingWorkRequest(context.Background(), oci_identity.GetTaggingWorkRequestRequest{
				WorkRequestId: workRequestId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
			wr := &response.TaggingWorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}

	// Should not wait when in replay mode
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of resources that finished the operation
	// for tagging work request is currently used for deletion only and it return single identifier
	for _, res := range response.TaggingWorkRequest.Resources {
		if res.Identifier != nil {
			identifier = res.Identifier
			break
		}
	}

	if response.Status == oci_identity.TaggingWorkRequestStatusFailed || response.Status == oci_identity.TaggingWorkRequestStatusCanceled {
		return nil, getIdentityTaggingWorkRequestErrors(client, workRequestId, retryPolicy, entityType, action)
	}
	return identifier, nil
}

func getIdentityTaggingWorkRequestErrors(client *oci_identity.IdentityClient, workRequestId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_identity.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListTaggingWorkRequestErrors(context.Background(), oci_identity.ListTaggingWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workRequestId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *IdentityTagResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCostTracking != nil {
		s.D.Set("is_cost_tracking", *s.Res.IsCostTracking)
	}

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TagNamespaceId != nil {
		s.D.Set("tag_namespace_id", *s.Res.TagNamespaceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Validator != nil {
		validatorArray := []interface{}{}
		if validatorMap := BaseTagDefinitionValidatorToMap(&s.Res.Validator); validatorMap != nil {
			validatorArray = append(validatorArray, validatorMap)
		}
		s.D.Set("validator", validatorArray)
	} else {
		s.D.Set("validator", nil)
	}

	return nil
}

func (s *IdentityTagResourceCrud) mapToBaseTagDefinitionValidator(fieldKeyFormat string) (oci_identity.BaseTagDefinitionValidator, error) {
	var baseObject oci_identity.BaseTagDefinitionValidator
	//discriminator
	validatorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validator_type"))
	var validatorType string
	if ok {
		validatorType = validatorTypeRaw.(string)
	} else {
		validatorType = "" // default value
	}
	switch strings.ToLower(validatorType) {
	case strings.ToLower("DEFAULT"):
		details := oci_identity.DefaultTagDefinitionValidator{}
		baseObject = details
	case strings.ToLower("ENUM"):
		details := oci_identity.EnumTagDefinitionValidator{}
		details.Values = []string{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.Values = tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown validator_type '%v' was specified", validatorType)
	}
	return baseObject, nil
}

func BaseTagDefinitionValidatorToMap(obj *oci_identity.BaseTagDefinitionValidator) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_identity.DefaultTagDefinitionValidator:
		result["validator_type"] = "DEFAULT"
	case oci_identity.EnumTagDefinitionValidator:
		result["validator_type"] = "ENUM"

		result["values"] = v.Values
	default:
		log.Printf("[WARN] Received 'validator_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func parseTagCompositeId(compositeId string) (tagName string, tagNamespaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("tagNamespaces/.*/tags/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	tagNamespaceId = parts[1]
	tagName = parts[3]

	return
}

func GetIdentityTagCompositeId(tagName string, tagNamespaceId string) string {
	compositeId := "tagNamespaces/" + tagNamespaceId + "/tags/" + tagName
	return compositeId
}

func identityTagWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "identity", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_identity.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}
