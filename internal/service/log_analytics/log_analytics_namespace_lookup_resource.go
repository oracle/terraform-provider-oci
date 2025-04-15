// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func LogAnalyticsNamespaceLookupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createLogAnalyticsNamespaceLookup,
		Read:   readLogAnalyticsNamespaceLookup,
		Update: updateLogAnalyticsNamespaceLookup,
		Delete: deleteLogAnalyticsNamespaceLookup,
		Schema: map[string]*schema.Schema{
			// Required
			"lookup_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"register_lookup_file": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"categories": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      categoriesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
						"is_system": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"char_encoding": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_match_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      lookupFieldsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"common_field_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_match_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_common_field": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"match_operator": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"position": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_hidden": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"max_matches": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"active_edit_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"canonical_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"edit_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_built_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_reference_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"referring_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"canonical_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"chunks_processed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"failure_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"filename": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_chunks": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceLookup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceLookup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespaceLookup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceLookup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceLookupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsLookup
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) ID() string {
	return GetNamespaceLookupCompositeId(s.D.Get("lookup_name").(string), s.D.Get("namespace").(string))
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) Create() error {
	registerRequest := oci_log_analytics.RegisterLookupRequest{}
	var namespaceName string
	var lookupName string

	if charEncoding, ok := s.D.GetOkExists("char_encoding"); ok {
		tmp := charEncoding.(string)
		registerRequest.CharEncoding = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		registerRequest.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		registerRequest.Description = &tmp
	}

	if isHidden, ok := s.D.GetOkExists("is_hidden"); ok {
		tmp := isHidden.(bool)
		registerRequest.IsHidden = &tmp
	}

	if lkName, ok := s.D.GetOkExists("lookup_name"); ok {
		lookupName = lkName.(string)
		registerRequest.Name = &lookupName
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		registerRequest.NamespaceName = &namespaceName
	}

	if registerLkFile, ok := s.D.GetOkExists("register_lookup_file"); ok {
		registerLookupFile := registerLkFile.(string)

		var definedTags map[string]map[string]interface{}
		var freeformTags map[string]string
		var isDefinedTagsSet bool
		var isFreeformTagsSet bool
		var deferr error

		if defTags, ok := s.D.GetOkExists("defined_tags"); ok {
			definedTags, deferr = tfresource.MapToDefinedTags(defTags.(map[string]interface{}))
			if deferr != nil {
				return deferr
			}
			isDefinedTagsSet = true
		}

		if freefTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			freeformTags = tfresource.ObjectMapToStringMap(freefTags.(map[string]interface{}))
			isFreeformTagsSet = true
		}

		if strings.HasSuffix(registerLookupFile, ".csv") && (isDefinedTagsSet || isFreeformTagsSet) {
			// Read the contents of the file as bytes
			contents, err := ioutil.ReadFile(registerLookupFile)
			if err != nil {
				return fmt.Errorf("The specified lookup file is not available: %q", err)
			}
			// Encode the csv content to base64 string
			encodedCsvContent := base64.StdEncoding.EncodeToString(contents)
			// Create the JSON payload with encoded csv content and tags
			registerLookupPayload := RegisterLookupPayload{
				CsvContent:   &encodedCsvContent,
				DefinedTags:  definedTags,
				FreeformTags: freeformTags,
			}
			// Marshal the JSON to bytes
			registerLookupPayloadContents, err := json.Marshal(registerLookupPayload)
			if err != nil {
				return err
			}
			// Set the JSON bytes as ReadCloser in request
			registerRequest.RegisterLookupContentFileBody = ioutil.NopCloser(bytes.NewReader(registerLookupPayloadContents))
		} else {
			contents, err := ioutil.ReadFile(registerLookupFile)
			if err != nil {
				return fmt.Errorf("The specified lookup file is not available: %q", err)
			}
			registerRequest.RegisterLookupContentFileBody = ioutil.NopCloser(bytes.NewReader(contents))
		}
	}

	if lkType, ok := s.D.GetOkExists("type"); ok {
		registerRequest.Type = oci_log_analytics.RegisterLookupTypeEnum(lkType.(string))
	}

	registerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.RegisterLookup(context.Background(), registerRequest)
	if err != nil {
		return err
	}

	// Wait until GetLookup returns successful status
	lookup, getLookupErr := s.getNamespaceLookupFromGetLookup(&namespaceName, &lookupName, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"), s.D.Timeout(schema.TimeoutDelete))
	if getLookupErr != nil {
		return getLookupErr
	}

	updateRequest := oci_log_analytics.UpdateLookupRequest{}
	updateRequest.LookupName = &lookupName
	updateRequest.NamespaceName = &namespaceName
	needsUpdate := false

	if categories, ok := s.D.GetOkExists("categories"); ok {
		interfaces := categories.(*schema.Set).List()
		tmp := make([]oci_log_analytics.LogAnalyticsCategory, len(interfaces))
		for i := range interfaces {
			stateDataIndex := categoriesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "categories", stateDataIndex)
			converted, err := s.mapToLogAnalyticsCategory(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("categories") {
			updateRequest.Categories = tmp
			needsUpdate = true
		}
	}

	if defaultMatchValue, ok := s.D.GetOkExists("default_match_value"); ok {
		tmp := defaultMatchValue.(string)
		updateRequest.DefaultMatchValue = &tmp
		needsUpdate = true
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.(*schema.Set).List()
		tmp := make([]oci_log_analytics.LogAnalyticsLookupFields, len(interfaces))
		for i := range interfaces {
			stateDataIndex := lookupFieldsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fields", stateDataIndex)
			converted, err := s.mapToLogAnalyticsLookupFields(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			updateRequest.Fields = tmp
			needsUpdate = true
		}
	}

	if maxMatches, ok := s.D.GetOkExists("max_matches"); ok {
		tmp := maxMatches.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert maxMatches string: %s to an int64 and encountered error: %v", tmp, err)
		}
		updateRequest.MaxMatches = &tmpInt64
		needsUpdate = true
	}

	if needsUpdate {
		updateRequest.Description = lookup.Description
		updateRequest.DefinedTags = lookup.DefinedTags
		updateRequest.FreeformTags = lookup.FreeformTags

		updateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

		updateResponse, err := s.Client.UpdateLookup(context.Background(), updateRequest)
		if err != nil {
			return err
		}

		s.Res = &updateResponse.LogAnalyticsLookup
	} else {
		s.D.SetId(GetNamespaceLookupCompositeId(lookupName, namespaceName))
		return s.Get()
	}
	return nil
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) getNamespaceLookupFromGetLookup(namespaceName *string, lookupName *string, retryPolicy *oci_common.RetryPolicy,
	timeout time.Duration) (*oci_log_analytics.LogAnalyticsLookup, error) {

	// Wait until lookup status becomes successful
	lookup, err := namespaceLookupWaitForGetLookup(namespaceName, lookupName, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return nil, err
	}

	return lookup, nil
}

func namespaceLookupWaitForGetLookup(namespaceName *string, lookupName *string,
	timeout time.Duration, disableFoundRetries bool, client *oci_log_analytics.LogAnalyticsClient) (*oci_log_analytics.LogAnalyticsLookup, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "log_analytics")
	retryPolicy.ShouldRetryOperation = namespaceLookupWorkRequestShouldRetryFunc(timeout, false)

	response := oci_log_analytics.GetLookupResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			"inProgress",
		},
		Target: []string{
			"successful",
			"failed",
			"failedNoReplay",
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetLookup(context.Background(),
				oci_log_analytics.GetLookupRequest{
					NamespaceName: namespaceName,
					LookupName:    lookupName,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			lookup := &response.LogAnalyticsLookup
			return lookup, *lookup.StatusSummary.Status, err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// Check if lookup status became successful, return error if not
	if *response.LogAnalyticsLookup.StatusSummary.Status != "successful" {
		return nil, getErrorFromLogAnalyticsNamespaceGetLookup()
	}

	return &response.LogAnalyticsLookup, nil
}

func getErrorFromLogAnalyticsNamespaceGetLookup() error {
	getLookupErr := fmt.Errorf("lookup status did not become successful")
	return getLookupErr
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) getNamespaceLookupFromWorkRequest(namespaceName *string, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := namespaceLookupWaitForWorkRequest(namespaceName, workId, "log_analytics",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func namespaceLookupWorkRequestShouldRetryFunc(timeout time.Duration, isConfigWorkRequest bool) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "log_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set or if lookup status is not in progress
		if isConfigWorkRequest {
			if workRequestResponse, ok := response.Response.(oci_log_analytics.GetConfigWorkRequestResponse); ok {
				return workRequestResponse.TimeFinished == nil
			}
		} else {
			if getLookupResponse, ok := response.Response.(oci_log_analytics.GetLookupResponse); ok {
				return *getLookupResponse.LogAnalyticsLookup.StatusSummary.Status == "inProgress"
			}
		}
		return false
	}
}

func namespaceLookupWaitForWorkRequest(namespaceName *string, wId *string, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_log_analytics.LogAnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "log_analytics")
	retryPolicy.ShouldRetryOperation = namespaceLookupWorkRequestShouldRetryFunc(timeout, true)

	response := oci_log_analytics.GetConfigWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateInProgress),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateAccepted),
		},
		Target: []string{
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateSucceeded),
			string(oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetConfigWorkRequest(context.Background(),
				oci_log_analytics.GetConfigWorkRequestRequest{
					NamespaceName: namespaceName,
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.LogAnalyticsConfigWorkRequest
			return wr, string(wr.LifecycleState), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.LogAnalyticsConfigWorkRequest.LifecycleState == oci_log_analytics.LogAnalyticsConfigWorkRequestLifecycleStateFailed {
		return nil, getErrorFromLogAnalyticsNamespaceLookupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return nil, nil
}

func getErrorFromLogAnalyticsNamespaceLookupWorkRequest(client *oci_log_analytics.LogAnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeEnum) error {
	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s", *workId, entityType, action)
	return workRequestErr
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) Get() error {
	request := oci_log_analytics.GetLookupRequest{}

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		request.LookupName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	lookupName, namespace, err := parseNamespaceLookupCompositeId(s.D.Id())
	if err == nil {
		request.LookupName = &lookupName
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetLookup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsLookup
	return nil
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_log_analytics.UpdateLookupRequest{}

	if categories, ok := s.D.GetOkExists("categories"); ok {
		interfaces := categories.(*schema.Set).List()
		tmp := make([]oci_log_analytics.LogAnalyticsCategory, len(interfaces))
		for i := range interfaces {
			stateDataIndex := categoriesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "categories", stateDataIndex)
			converted, err := s.mapToLogAnalyticsCategory(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("categories") {
			request.Categories = tmp
		}
	}

	if defaultMatchValue, ok := s.D.GetOkExists("default_match_value"); ok {
		tmp := defaultMatchValue.(string)
		request.DefaultMatchValue = &tmp
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

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.(*schema.Set).List()
		tmp := make([]oci_log_analytics.LogAnalyticsLookupFields, len(interfaces))
		for i := range interfaces {
			stateDataIndex := lookupFieldsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fields", stateDataIndex)
			converted, err := s.mapToLogAnalyticsLookupFields(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		request.LookupName = &tmp
	}

	if maxMatches, ok := s.D.GetOkExists("max_matches"); ok {
		tmp := maxMatches.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert maxMatches string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MaxMatches = &tmpInt64
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateLookup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsLookup
	return nil
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteLookupRequest{}

	var namespaceName string

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		request.LookupName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName = namespace.(string)
		request.NamespaceName = &namespaceName
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DeleteLookup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	delWorkRequestErr := s.getNamespaceLookupFromWorkRequest(&namespaceName, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics"),
		oci_log_analytics.LogAnalyticsConfigWorkRequestOperationTypeDeleteLookup, s.D.Timeout(schema.TimeoutDelete))
	return delWorkRequestErr
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) SetData() error {

	lookupName, namespace, err := parseNamespaceLookupCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("lookup_name", &lookupName)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ActiveEditVersion != nil {
		s.D.Set("active_edit_version", strconv.FormatInt(*s.Res.ActiveEditVersion, 10))
	}

	if s.Res.CanonicalLink != nil {
		s.D.Set("canonical_link", *s.Res.CanonicalLink)
	}

	categories := []interface{}{}
	for _, item := range s.Res.Categories {
		categories = append(categories, LogAnalyticsCategoryToMap(item))
	}
	s.D.Set("categories", schema.NewSet(categoriesHashCodeForSets, categories))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EditVersion != nil {
		s.D.Set("edit_version", strconv.FormatInt(*s.Res.EditVersion, 10))
	}

	fields := []interface{}{}
	for _, item := range s.Res.Fields {
		fields = append(fields, LookupFieldToMap(item))
	}
	s.D.Set("fields", schema.NewSet(lookupFieldsHashCodeForSets, fields))

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsBuiltIn != nil {
		s.D.Set("is_built_in", strconv.FormatInt(*s.Res.IsBuiltIn, 10))
	}

	if s.Res.IsHidden != nil {
		s.D.Set("is_hidden", *s.Res.IsHidden)
	}

	if s.Res.Id != nil {
		s.D.Set("lookup_id", *s.Res.Id)
	}

	if s.Res.LookupDisplayName != nil {
		s.D.Set("lookup_display_name", *s.Res.LookupDisplayName)
	}

	if s.Res.LookupReference != nil {
		s.D.Set("lookup_reference", strconv.FormatInt(*s.Res.LookupReference, 10))
	}

	if s.Res.LookupReferenceString != nil {
		s.D.Set("lookup_reference_string", *s.Res.LookupReferenceString)
	}

	if s.Res.ReferringSources != nil {
		s.D.Set("referring_sources", []interface{}{AutoLookupsToMap(s.Res.ReferringSources)})
	} else {
		s.D.Set("referring_sources", nil)
	}

	if registerLookupFile, ok := s.D.GetOkExists("register_lookup_file"); ok {
		s.D.Set("register_lookup_file", registerLookupFile)
	} else {
		if s.Res.StatusSummary.Filename != nil {
			s.D.Set("register_lookup_file", *s.Res.StatusSummary.Filename)
		}
	}

	if s.Res.StatusSummary != nil {
		s.D.Set("status_summary", []interface{}{StatusSummaryToMap(s.Res.StatusSummary)})
	} else {
		s.D.Set("status_summary", nil)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func GetNamespaceLookupCompositeId(lookupName string, namespace string) string {
	lookupName = url.PathEscape(lookupName)
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/lookups/" + lookupName
	return compositeId
}

func parseNamespaceLookupCompositeId(compositeId string) (lookupName string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/lookups/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	lookupName, _ = url.PathUnescape(parts[3])

	return
}

func AutoLookupsToMap(obj *oci_log_analytics.AutoLookups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanonicalLink != nil {
		result["canonical_link"] = string(*obj.CanonicalLink)
	}

	if obj.TotalCount != nil {
		result["total_count"] = strconv.FormatInt(*obj.TotalCount, 10)
	}

	return result
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) mapToLogAnalyticsCategory(fieldKeyFormat string) (oci_log_analytics.LogAnalyticsCategory, error) {
	result := oci_log_analytics.LogAnalyticsCategory{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if isSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_system")); ok {
		tmp := isSystem.(bool)
		result.IsSystem = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func LogAnalyticsLookupToMap(obj oci_log_analytics.LogAnalyticsLookup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveEditVersion != nil {
		result["active_edit_version"] = strconv.FormatInt(*obj.ActiveEditVersion, 10)
	}

	if obj.CanonicalLink != nil {
		result["canonical_link"] = string(*obj.CanonicalLink)
	}

	categories := []interface{}{}
	for _, item := range obj.Categories {
		categories = append(categories, LogAnalyticsCategoryToMap(item))
	}
	result["categories"] = categories

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EditVersion != nil {
		result["edit_version"] = strconv.FormatInt(*obj.EditVersion, 10)
	}

	fields := []interface{}{}
	for _, item := range obj.Fields {
		fields = append(fields, LookupFieldToMap(item))
	}
	result["fields"] = fields

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsBuiltIn != nil {
		result["is_built_in"] = strconv.FormatInt(*obj.IsBuiltIn, 10)
	}

	if obj.IsHidden != nil {
		result["is_hidden"] = bool(*obj.IsHidden)
	}

	if obj.LookupDisplayName != nil {
		result["lookup_display_name"] = string(*obj.LookupDisplayName)
	}

	if obj.LookupReference != nil {
		result["lookup_reference"] = strconv.FormatInt(*obj.LookupReference, 10)
	}

	if obj.LookupReferenceString != nil {
		result["lookup_reference_string"] = string(*obj.LookupReferenceString)
	}

	if obj.ReferringSources != nil {
		result["referring_sources"] = []interface{}{AutoLookupsToMap(obj.ReferringSources)}
	}

	if obj.StatusSummary != nil {
		result["status_summary"] = []interface{}{StatusSummaryToMap(obj.StatusSummary)}
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) mapToLogAnalyticsLookupFields(fieldKeyFormat string) (oci_log_analytics.LogAnalyticsLookupFields, error) {
	result := oci_log_analytics.LogAnalyticsLookupFields{}

	if commonFieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "common_field_name")); ok {
		tmp := commonFieldName.(string)
		result.CommonFieldName = &tmp
	}

	if defaultMatchValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_match_value")); ok {
		tmp := defaultMatchValue.(string)
		result.DefaultMatchValue = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if isCommonField, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_common_field")); ok {
		tmp := isCommonField.(bool)
		result.IsCommonField = &tmp
	}

	if matchOperator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_operator")); ok {
		tmp := matchOperator.(string)
		result.MatchOperator = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if position, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "position")); ok {
		tmp := position.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert position string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.Position = &tmpInt64
	}

	return result, nil
}

func LookupFieldToMap(obj oci_log_analytics.LookupField) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommonFieldName != nil {
		result["common_field_name"] = string(*obj.CommonFieldName)
	}

	if obj.DefaultMatchValue != nil {
		result["default_match_value"] = string(*obj.DefaultMatchValue)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsCommonField != nil {
		result["is_common_field"] = bool(*obj.IsCommonField)
	}

	if obj.MatchOperator != nil {
		result["match_operator"] = string(*obj.MatchOperator)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Position != nil {
		result["position"] = strconv.FormatInt(*obj.Position, 10)
	}

	return result
}

func StatusSummaryToMap(obj *oci_log_analytics.StatusSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChunksProcessed != nil {
		result["chunks_processed"] = strconv.FormatInt(*obj.ChunksProcessed, 10)
	}

	if obj.FailureDetails != nil {
		result["failure_details"] = string(*obj.FailureDetails)
	}

	if obj.Filename != nil {
		result["filename"] = string(*obj.Filename)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TotalChunks != nil {
		result["total_chunks"] = strconv.FormatInt(*obj.TotalChunks, 10)
	}

	return result
}

func categoriesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}

	return utils.GetStringHashcode(buf.String())
}

func lookupFieldsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}

	return utils.GetStringHashcode(buf.String())
}

func (s *LogAnalyticsNamespaceLookupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeLookupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		changeCompartmentRequest.LookupName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeLookupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

type RegisterLookupPayload struct {
	CsvContent   *string                           `mandatory:"true" json:"csvContent"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags,omitempty"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags,omitempty"`
}
