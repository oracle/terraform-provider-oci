// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NosqlConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNosqlConfiguration,
		Read:     readNosqlConfiguration,
		Update:   updateNosqlConfiguration,
		Delete:   deleteNosqlConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"HOSTED",
					"MULTI_TENANCY",
				}, true),
			},

			// Optional
			"is_opc_dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kms_key": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kms_key_state": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: allChangeSuppressFunction,
						},
						"kms_vault_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: allChangeSuppressFunction,
						},
						"time_updated": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: allChangeSuppressFunction,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createNosqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.CreateResource(d, sync)
}

func readNosqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

func updateNosqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNosqlConfiguration(d *schema.ResourceData, m interface{}) error {
	log.Printf("[WARNING] Deleting the entire configuration is not supported. " +
		"To remove specific settings, set the corresponding property to null.")
	return nil
}

type NosqlConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_nosql.NosqlClient
	Res                    *oci_nosql.Configuration
	DisableNotFoundRetries bool
}

func (s *NosqlConfigurationResourceCrud) ID() string {
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		name := GetConfigurationCompositeId(tmp)
		return name
	}
	return "global"
}

func (s *NosqlConfigurationResourceCrud) Create() error {
	if keyId, ok := s.D.GetOkExists("kms_key.0.id"); ok {
		if keyId != nil && keyId != "" {
			return s.updateConfiguration()
		}
	}
	return s.unassignKmsKey()
}

func (s *NosqlConfigurationResourceCrud) getConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_nosql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	configurationId, err := configurationWaitForWorkRequest(workId, "configuration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, configurationId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_nosql.DeleteWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
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
		if tfresource.ShouldRetry(response, false, "nosql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_nosql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func configurationWaitForWorkRequest(wId *string, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_nosql.NosqlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "nosql")
	retryPolicy.ShouldRetryOperation = configurationWorkRequestShouldRetryFunc(timeout)

	response := oci_nosql.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_nosql.WorkRequestStatusInProgress),
			string(oci_nosql.WorkRequestStatusAccepted),
			string(oci_nosql.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_nosql.WorkRequestStatusSucceeded),
			string(oci_nosql.WorkRequestStatusFailed),
			string(oci_nosql.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_nosql.GetWorkRequestRequest{
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
				tmp := GetConfigurationCompositeId(*response.WorkRequest.CompartmentId)
				identifier = &tmp
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_nosql.WorkRequestStatusFailed || response.Status == oci_nosql.WorkRequestStatusCanceled {
		return nil, getErrorFromNosqlConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNosqlConfigurationWorkRequest(client *oci_nosql.NosqlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_nosql.ListWorkRequestErrorsRequest{
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

func (s *NosqlConfigurationResourceCrud) Get() error {
	request := oci_nosql.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	compartmentId, err := parseConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.CompartmentId = &compartmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *NosqlConfigurationResourceCrud) Update() error {
	if keyId, ok := s.D.GetOkExists("kms_key.0.id"); ok && s.D.HasChange("kms_key.0.id") {
		if keyId == nil || keyId == "" {
			return s.unassignKmsKey()
		}
	}
	return s.updateConfiguration()
}

func (s *NosqlConfigurationResourceCrud) updateConfiguration() error {
	request := oci_nosql.UpdateConfigurationRequest{}
	err := s.populateTopLevelPolymorphicUpdateConfigurationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql"),
		oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NosqlConfigurationResourceCrud) unassignKmsKey() error {
	request := oci_nosql.UnassignKmsKeyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.UnassignKmsKey(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "nosql"),
		oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NosqlConfigurationResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	compartmentId, err := parseConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("compartment_id", &compartmentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_nosql.HostedConfiguration:
		s.D.Set("environment", "HOSTED")

		if v.KmsKey != nil {
			s.D.Set("kms_key", []interface{}{KmsKeyToMap(v.KmsKey)})
		} else {
			s.D.Set("kms_key", nil)
		}
	case oci_nosql.MultiTenancyConfiguration:
		s.D.Set("environment", "MULTI_TENANCY")
	default:
		log.Printf("[WARN] Received 'environment' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetConfigurationCompositeId(compartmentId string) string {
	compartmentId = url.PathEscape(compartmentId)
	compositeId := "configuration/compartmentId/" + compartmentId + ""
	return compositeId
}

func parseConfigurationCompositeId(compositeId string) (compartmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("configuration/compartmentId/.*", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	compartmentId, _ = url.PathUnescape(parts[2])
	return
}

func (s *NosqlConfigurationResourceCrud) mapToKmsKey(fieldKeyFormat string) (oci_nosql.KmsKey, error) {
	result := oci_nosql.KmsKey{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if kmsVaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_vault_id")); ok {
		if kmsVaultId != "" {
			tmp := kmsVaultId.(string)
			result.KmsVaultId = &tmp
		}
	}

	return result, nil
}

func KmsKeyToMap(obj *oci_nosql.KmsKey) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["kms_key_state"] = string(obj.KmsKeyState)

	if obj.KmsVaultId != nil {
		result["kms_vault_id"] = string(*obj.KmsVaultId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func (s *NosqlConfigurationResourceCrud) populateTopLevelPolymorphicUpdateConfigurationRequest(request *oci_nosql.UpdateConfigurationRequest) error {
	//discriminator
	environmentRaw, ok := s.D.GetOkExists("environment")
	var environment string
	if ok {
		environment = environmentRaw.(string)
	} else {
		environment = "" // default value
	}
	switch strings.ToLower(environment) {
	case strings.ToLower("HOSTED"):
		details := oci_nosql.UpdateHostedConfigurationDetails{}
		if kmsKey, ok := s.D.GetOkExists("kms_key"); ok {
			if tmpList := kmsKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kms_key", 0)
				tmp, err := s.mapToKmsKey(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.KmsKey = &tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			request.CompartmentId = &tmp
		}
		if isOpcDryRun, ok := s.D.GetOkExists("is_opc_dry_run"); ok {
			tmp := isOpcDryRun.(bool)
			request.IsOpcDryRun = &tmp
		}
		request.UpdateConfigurationDetails = details
	case strings.ToLower("MULTI_TENANCY"):
		details := oci_nosql.UpdateMultiTenancyConfigurationDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			request.CompartmentId = &tmp
		}
		if isOpcDryRun, ok := s.D.GetOkExists("is_opc_dry_run"); ok {
			tmp := isOpcDryRun.(bool)
			request.IsOpcDryRun = &tmp
		}
		request.UpdateConfigurationDetails = details
	default:
		return fmt.Errorf("unknown environment '%v' was specified", environment)
	}
	return nil
}

func allChangeSuppressFunction(k string, old string, new string, d *schema.ResourceData) bool {
	if d.Id() == "" {
		return false
	}
	log.Printf("[INFO] Ignoring change to %s, k=%s, old=%s, new=%s", d.Id(), k, old, new)
	return true
}
