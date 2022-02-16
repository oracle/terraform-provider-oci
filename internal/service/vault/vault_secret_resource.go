// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_vault "github.com/oracle/oci-go-sdk/v58/vault"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func VaultSecretResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createVaultSecret,
		Read:     readVaultSecret,
		Update:   updateVaultSecret,
		Delete:   deleteVaultSecret,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_content": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"content_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BASE64",
							}, true),
						},

						// Optional
						"content": {
							Type: schema.TypeString,
							//Optional: true,
							//Computed: true,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stage": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"secret_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vault_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"key_id": {
				Type: schema.TypeString,
				//Optional: true,
				//Computed: true,
				Required: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"secret_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SECRET_EXPIRY_RULE",
								"SECRET_REUSE_RULE",
							}, true),
						},

						// Optional
						"is_enforced_on_deleted_secret_versions": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_secret_content_retrieval_blocked_on_expiry": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"secret_version_expiry_interval": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_of_absolute_expiry": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},

			// Computed
			"current_version_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
			"time_of_current_version_expiry": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVaultSecret(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.CreateResource(d, sync)
}

func readVaultSecret(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.ReadResource(sync)
}

func updateVaultSecret(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteVaultSecret(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.DeleteResource(d, sync)
}

type VaultSecretResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_vault.VaultsClient
	Res                    *oci_vault.Secret
	DisableNotFoundRetries bool
}

func (s *VaultSecretResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VaultSecretResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateCreating),
	}
}

func (s *VaultSecretResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateActive),
	}
}

func (s *VaultSecretResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateDeleting),
		string(oci_vault.SecretLifecycleStateSchedulingDeletion),
	}
}

func (s *VaultSecretResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_vault.SecretLifecycleStateDeleted),
		string(oci_vault.SecretLifecycleStatePendingDeletion),
	}
}

func (s *VaultSecretResourceCrud) Create() error {
	request := oci_vault.CreateSecretRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = metadata.(map[string]interface{})
	}

	if secretContent, ok := s.D.GetOkExists("secret_content"); ok {
		if tmpList := secretContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secret_content", 0)
			tmp, err := s.mapToSecretContentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SecretContent = tmp
		}
	}

	if secretName, ok := s.D.GetOkExists("secret_name"); ok {
		tmp := secretName.(string)
		request.SecretName = &tmp
	}

	if secretRules, ok := s.D.GetOkExists("secret_rules"); ok {
		interfaces := secretRules.([]interface{})
		tmp := make([]oci_vault.SecretRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secret_rules", stateDataIndex)
			converted, err := s.mapToSecretRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("secret_rules") {
			request.SecretRules = tmp
		}
	}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.CreateSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret
	return nil
}

func (s *VaultSecretResourceCrud) Get() error {
	request := oci_vault.GetSecretRequest{}

	tmp := s.D.Id()
	request.SecretId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.GetSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret
	return nil
}

func (s *VaultSecretResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_vault.UpdateSecretRequest{}

	if currentVersionNumber, ok := s.D.GetOkExists("current_version_number"); ok && s.D.HasChange("current_version_number") {
		tmp := currentVersionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert currentVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CurrentVersionNumber = &tmpInt64
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = metadata.(map[string]interface{})
	}

	if secretContent, ok := s.D.GetOkExists("secret_content"); ok && s.D.HasChange("secret_content") {
		if tmpList := secretContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secret_content", 0)
			tmp, err := s.mapToSecretContentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SecretContent = tmp
		}
	}

	tmp := s.D.Id()
	request.SecretId = &tmp

	if secretRules, ok := s.D.GetOkExists("secret_rules"); ok && s.D.HasChange("secret_rules") {
		interfaces := secretRules.([]interface{})
		tmp := make([]oci_vault.SecretRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "secret_rules", stateDataIndex)
			converted, err := s.mapToSecretRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("secret_rules") {
			request.SecretRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	response, err := s.Client.UpdateSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Secret
	return nil
}

func (s *VaultSecretResourceCrud) Delete() error {
	request := oci_vault.ScheduleSecretDeletionRequest{}

	tmp := s.D.Id()
	request.SecretId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	_, err := s.Client.ScheduleSecretDeletion(context.Background(), request)
	return err
}

func (s *VaultSecretResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentVersionNumber != nil {
		s.D.Set("current_version_number", strconv.FormatInt(*s.Res.CurrentVersionNumber, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.SecretName != nil {
		s.D.Set("secret_name", *s.Res.SecretName)
	}

	secretRules := []interface{}{}
	for _, item := range s.Res.SecretRules {
		secretRules = append(secretRules, SecretRuleToMap(item))
	}
	s.D.Set("secret_rules", secretRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfCurrentVersionExpiry != nil {
		s.D.Set("time_of_current_version_expiry", s.Res.TimeOfCurrentVersionExpiry.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func (s *VaultSecretResourceCrud) mapToSecretContentDetails(fieldKeyFormat string) (oci_vault.SecretContentDetails, error) {
	var baseObject oci_vault.SecretContentDetails
	//discriminator
	contentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content_type"))
	var contentType string
	if ok {
		contentType = contentTypeRaw.(string)
	} else {
		contentType = "" // default value
	}
	switch strings.ToLower(contentType) {
	case strings.ToLower("BASE64"):
		details := oci_vault.Base64SecretContentDetails{}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			tmp := content.(string)
			details.Content = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if stage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stage")); ok {
			details.Stage = oci_vault.SecretContentDetailsStageEnum(stage.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown content_type '%v' was specified", contentType)
	}
	return baseObject, nil
}

func SecretContentDetailsToMap(obj *oci_vault.SecretContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_vault.Base64SecretContentDetails:
		result["content_type"] = "BASE64"

		if v.Content != nil {
			result["content"] = string(*v.Content)
		}
	default:
		log.Printf("[WARN] Received 'content_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *VaultSecretResourceCrud) mapToSecretRule(fieldKeyFormat string) (oci_vault.SecretRule, error) {
	var baseObject oci_vault.SecretRule
	//discriminator
	ruleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type"))
	var ruleType string
	if ok {
		ruleType = ruleTypeRaw.(string)
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("SECRET_EXPIRY_RULE"):
		details := oci_vault.SecretExpiryRule{}
		if isSecretContentRetrievalBlockedOnExpiry, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secret_content_retrieval_blocked_on_expiry")); ok {
			tmp := isSecretContentRetrievalBlockedOnExpiry.(bool)
			details.IsSecretContentRetrievalBlockedOnExpiry = &tmp
		}
		if secretVersionExpiryInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_version_expiry_interval")); ok {
			tmp := secretVersionExpiryInterval.(string)
			details.SecretVersionExpiryInterval = &tmp
		}
		if timeOfAbsoluteExpiry, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_absolute_expiry")); ok {
			tmp, err := time.Parse(time.RFC3339, timeOfAbsoluteExpiry.(string))
			if err != nil {
				return details, err
			}
			details.TimeOfAbsoluteExpiry = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	case strings.ToLower("SECRET_REUSE_RULE"):
		details := oci_vault.SecretReuseRule{}
		if isEnforcedOnDeletedSecretVersions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enforced_on_deleted_secret_versions")); ok {
			tmp := isEnforcedOnDeletedSecretVersions.(bool)
			details.IsEnforcedOnDeletedSecretVersions = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown rule_type '%v' was specified", ruleType)
	}
	return baseObject, nil
}

func SecretRuleToMap(obj oci_vault.SecretRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_vault.SecretExpiryRule:
		result["rule_type"] = "SECRET_EXPIRY_RULE"

		if v.IsSecretContentRetrievalBlockedOnExpiry != nil {
			result["is_secret_content_retrieval_blocked_on_expiry"] = bool(*v.IsSecretContentRetrievalBlockedOnExpiry)
		}

		if v.SecretVersionExpiryInterval != nil {
			result["secret_version_expiry_interval"] = string(*v.SecretVersionExpiryInterval)
		}

		if v.TimeOfAbsoluteExpiry != nil {
			result["time_of_absolute_expiry"] = v.TimeOfAbsoluteExpiry.Format(time.RFC3339Nano)
		}
	case oci_vault.SecretReuseRule:
		result["rule_type"] = "SECRET_REUSE_RULE"

		if v.IsEnforcedOnDeletedSecretVersions != nil {
			result["is_enforced_on_deleted_secret_versions"] = bool(*v.IsEnforcedOnDeletedSecretVersions)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *VaultSecretResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_vault.ChangeSecretCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecretId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vault")

	_, err := s.Client.ChangeSecretCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
