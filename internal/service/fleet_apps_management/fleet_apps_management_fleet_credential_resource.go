// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementFleetCredentialResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementFleetCredential,
		Read:     readFleetAppsManagementFleetCredential,
		Update:   updateFleetAppsManagementFleetCredential,
		Delete:   deleteFleetAppsManagementFleetCredential,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_specifics": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_level": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"TARGET",
							}, true),
						},
						"resource_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"target": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"KEY_ENCRYPTION",
								"PLAIN_TEXT",
								"VAULT_SECRET",
							}, true),
						},

						// Optional
						"key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
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
			"user": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"KEY_ENCRYPTION",
								"PLAIN_TEXT",
								"VAULT_SECRET",
							}, true),
						},

						// Optional
						"key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
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

			// Optional

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

func createFleetAppsManagementFleetCredential(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementFleetCredential(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementFleetCredential(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementFleetCredential(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementFleetCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.FleetCredential
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) ID() string {
	return GetFleetCredentialCompositeId(*s.Res.Id, s.D.Get("fleet_id").(string))
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetCredentialLifecycleStateActive),
	}
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetCredentialLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateFleetCredentialRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if entitySpecifics, ok := s.D.GetOkExists("entity_specifics"); ok {
		if tmpList := entitySpecifics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entity_specifics", 0)
			tmp, err := s.mapToCredentialEntitySpecificDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EntitySpecifics = tmp
		}
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		if tmpList := password.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "password", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Password = tmp
		}
	}

	if user, ok := s.D.GetOkExists("user"); ok {
		if tmpList := user.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.User = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateFleetCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetCredential
	return nil
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) getFleetCredentialFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fleetCredentialId, err := fleetCredentialWaitForWorkRequest(workId, "fleetcredential",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fleetCredentialId)

	return s.Get()
}

func fleetCredentialWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fleetCredentialWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = fleetCredentialWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementFleetCredentialWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementFleetCredentialWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
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

func (s *FleetAppsManagementFleetCredentialResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetCredentialRequest{}

	tmp := s.D.Id()
	request.FleetCredentialId = &tmp

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	fleetCredentialId, fleetId, err := parseFleetCredentialCompositeId(s.D.Id())
	if err == nil {
		request.FleetCredentialId = &fleetCredentialId
		request.FleetId = &fleetId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetFleetCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetCredential
	return nil
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateFleetCredentialRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if entitySpecifics, ok := s.D.GetOkExists("entity_specifics"); ok {
		if tmpList := entitySpecifics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "entity_specifics", 0)
			tmp, err := s.mapToCredentialEntitySpecificDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EntitySpecifics = tmp
		}
	}

	tmp := s.D.Id()
	request.FleetCredentialId = &tmp

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		if tmpList := password.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "password", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Password = tmp
		}
	}

	if user, ok := s.D.GetOkExists("user"); ok {
		if tmpList := user.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user", 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.User = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateFleetCredential(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetCredentialFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteFleetCredentialRequest{}

	tmp := s.D.Id()
	request.FleetCredentialId = &tmp

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.DeleteFleetCredential(context.Background(), request)
	if err != nil {
		return err
	}

	// This call is synchronous, and returns an invalid WorkRequestId
	return nil
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) SetData() error {

	fleetCredentialId, fleetId, err := parseFleetCredentialCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(fleetCredentialId)
		s.D.Set("fleet_id", &fleetId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EntitySpecifics != nil {
		entitySpecificsArray := []interface{}{}
		if entitySpecificsMap := CredentialEntitySpecificDetailsToMap(&s.Res.EntitySpecifics); entitySpecificsMap != nil {
			entitySpecificsArray = append(entitySpecificsArray, entitySpecificsMap)
		}
		s.D.Set("entity_specifics", entitySpecificsArray)
	} else {
		s.D.Set("entity_specifics", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Password != nil {
		passwordArray := []interface{}{}
		if passwordMap := CredentialDetailsToMap(&s.Res.Password); passwordMap != nil {
			passwordArray = append(passwordArray, passwordMap)
		}
		s.D.Set("password", passwordArray)
	} else {
		s.D.Set("password", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.User != nil {
		userArray := []interface{}{}
		if userMap := CredentialDetailsToMap(&s.Res.User); userMap != nil {
			userArray = append(userArray, userMap)
		}
		s.D.Set("user", userArray)
	} else {
		s.D.Set("user", nil)
	}

	return nil
}

func GetFleetCredentialCompositeId(fleetCredentialId string, fleetId string) string {
	fleetCredentialId = url.PathEscape(fleetCredentialId)
	fleetId = url.PathEscape(fleetId)
	compositeId := "fleets/" + fleetId + "/fleetCredentials/" + fleetCredentialId
	return compositeId
}

func parseFleetCredentialCompositeId(compositeId string) (fleetCredentialId string, fleetId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fleets/.*/fleetCredentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fleetId, _ = url.PathUnescape(parts[1])
	fleetCredentialId, _ = url.PathUnescape(parts[3])
	return
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) mapToCredentialDetails(fieldKeyFormat string) (oci_fleet_apps_management.CredentialDetails, error) {
	var baseObject oci_fleet_apps_management.CredentialDetails
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("KEY_ENCRYPTION"):
		details := oci_fleet_apps_management.KeyEncryptionCredentialDetails{}
		if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if keyVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_version")); ok {
			tmp := keyVersion.(string)
			details.KeyVersion = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	case strings.ToLower("PLAIN_TEXT"):
		details := oci_fleet_apps_management.PlainTextCredentialDetails{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("VAULT_SECRET"):
		details := oci_fleet_apps_management.VaultSecretCredentialDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_version")); ok {
			tmp := secretVersion.(string)
			details.SecretVersion = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func CredentialDetailsToMap(obj *oci_fleet_apps_management.CredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.KeyEncryptionCredentialDetails:
		result["credential_type"] = "KEY_ENCRYPTION"

		if v.KeyId != nil {
			result["key_id"] = string(*v.KeyId)
		}

		if v.KeyVersion != nil {
			result["key_version"] = string(*v.KeyVersion)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}
	case oci_fleet_apps_management.PlainTextCredentialDetails:
		result["credential_type"] = "PLAIN_TEXT"

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_fleet_apps_management.VaultSecretCredentialDetails:
		result["credential_type"] = "VAULT_SECRET"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}

		if v.SecretVersion != nil {
			result["secret_version"] = string(*v.SecretVersion)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementFleetCredentialResourceCrud) mapToCredentialEntitySpecificDetails(fieldKeyFormat string) (oci_fleet_apps_management.CredentialEntitySpecificDetails, error) {
	var baseObject oci_fleet_apps_management.CredentialEntitySpecificDetails
	//discriminator
	credentialLevelRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_level"))
	var credentialLevel string
	if ok {
		credentialLevel = credentialLevelRaw.(string)
	} else {
		credentialLevel = "" // default value
	}
	switch strings.ToLower(credentialLevel) {
	case strings.ToLower("TARGET"):
		details := oci_fleet_apps_management.TargetCredentialEntitySpecificDetails{}
		if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
			tmp := resourceId.(string)
			details.ResourceId = &tmp
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			tmp := target.(string)
			details.Target = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_level '%v' was specified", credentialLevel)
	}
	return baseObject, nil
}

func CredentialEntitySpecificDetailsToMap(obj *oci_fleet_apps_management.CredentialEntitySpecificDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.TargetCredentialEntitySpecificDetails:
		result["credential_level"] = "TARGET"

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		if v.Target != nil {
			result["target"] = string(*v.Target)
		}
	default:
		log.Printf("[WARN] Received 'credential_level' of unknown type %v", *obj)
		return nil
	}

	return result
}

func FleetCredentialSummaryToMap(obj oci_fleet_apps_management.FleetCredentialSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EntitySpecifics != nil {
		entitySpecificsArray := []interface{}{}
		if entitySpecificsMap := CredentialEntitySpecificDetailsToMap(&obj.EntitySpecifics); entitySpecificsMap != nil {
			entitySpecificsArray = append(entitySpecificsArray, entitySpecificsMap)
		}
		result["entity_specifics"] = entitySpecificsArray
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Password != nil {
		passwordArray := []interface{}{}
		if passwordMap := CredentialDetailsToMap(&obj.Password); passwordMap != nil {
			passwordArray = append(passwordArray, passwordMap)
		}
		result["password"] = passwordArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.User != nil {
		userArray := []interface{}{}
		if userMap := CredentialDetailsToMap(&obj.User); userMap != nil {
			userArray = append(userArray, userMap)
		}
		result["user"] = userArray
	}

	return result
}
