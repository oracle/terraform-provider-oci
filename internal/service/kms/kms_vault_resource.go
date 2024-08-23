// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsVaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsVault,
		Read:     readKmsVault,
		Update:   updateKmsVault,
		Delete:   deleteKmsVault,
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
			"vault_type": {
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
			"external_key_manager_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"external_vault_endpoint_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"oauth_metadata": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"client_app_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"client_app_secret": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"idcs_account_name_url": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

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
			"time_of_deletion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restore_from_object_store": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"restore_from_file"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destination": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BUCKET",
								"PRE_AUTHENTICATED_REQUEST_URI",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"restore_from_file": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"restore_from_object_store"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"restore_vault_from_file_details": {
							Type:     schema.TypeString,
							Required: true,
						},
						"content_length": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Optional
						"content_md5": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"restore_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_key_manager_metadata_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"external_vault_endpoint_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oauth_metadata_summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"client_app_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"idcs_account_name_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vendor": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_vault_replicable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replica_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"replication_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"restored_from_vault_id": {
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
		},
	}
}

func createKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.CreateResource(d, sync)
}

func readKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

func updateKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.DeleteResource(d, sync)
}

type KmsVaultResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.Vault
	DisableNotFoundRetries bool
}

func (s *KmsVaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KmsVaultResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateCreating),
		string(oci_kms.VaultLifecycleStateRestoring),
	}
}

func (s *KmsVaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateActive),
	}
}

func (s *KmsVaultResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleting),
		string(oci_kms.VaultLifecycleStateSchedulingDeletion),
	}
}

func (s *KmsVaultResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleted),
		string(oci_kms.VaultLifecycleStatePendingDeletion),
	}
}

func (s *KmsVaultResourceCrud) Create() error {
	if _, ok := s.D.GetOkExists("restore_from_file"); ok {
		err := s.RestoreVaultFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
		return s.UpdateVaultDetails()
	}
	if _, ok := s.D.GetOkExists("restore_from_object_store"); ok {
		err := s.RestoreVaultFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
		return s.UpdateVaultDetails()
	}

	request := oci_kms.CreateVaultRequest{}

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

	if externalKeyManagerMetadata, ok := s.D.GetOkExists("external_key_manager_metadata"); ok {
		if tmpList := externalKeyManagerMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_key_manager_metadata", 0)
			tmp, err := s.mapToExternalKeyManagerMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExternalKeyManagerMetadata = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vaultType, ok := s.D.GetOkExists("vault_type"); ok {
		request.VaultType = oci_kms.CreateVaultDetailsVaultTypeEnum(vaultType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Get() error {
	request := oci_kms.GetVaultRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Update() error {
	if _, ok := s.D.GetOk("restore_from_file"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreVaultFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	if _, ok := s.D.GetOk("restore_from_object_store"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreVaultFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	return s.UpdateVaultDetails()
}

func (s *KmsVaultResourceCrud) UpdateVaultDetails() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_kms.UpdateVaultRequest{}
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}
	tmp := s.D.Id()
	request.VaultId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Delete() error {
	request := oci_kms.ScheduleVaultDeletionRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ScheduleVaultDeletion(context.Background(), request)
	return err
}

func (s *KmsVaultResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoEndpoint != nil {
		s.D.Set("crypto_endpoint", *s.Res.CryptoEndpoint)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKeyManagerMetadataSummary != nil {
		s.D.Set("external_key_manager_metadata_summary", []interface{}{ExternalKeyManagerMetadataSummaryToMap(s.Res.ExternalKeyManagerMetadataSummary)})
	} else {
		s.D.Set("external_key_manager_metadata_summary", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.IsVaultReplicable != nil {
		s.D.Set("is_vault_replicable", *s.Res.IsVaultReplicable)
	}

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	if s.Res.ReplicaDetails != nil {
		s.D.Set("replica_details", []interface{}{VaultReplicaDetailsToMap(s.Res.ReplicaDetails)})
	} else {
		s.D.Set("replica_details", nil)
	}

	if s.Res.RestoredFromVaultId != nil {
		s.D.Set("restored_from_vault_id", *s.Res.RestoredFromVaultId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	s.D.Set("vault_type", s.Res.VaultType)

	return nil
}

func (s *KmsVaultResourceCrud) mapToExternalKeyManagerMetadata(fieldKeyFormat string) (oci_kms.ExternalKeyManagerMetadata, error) {
	result := oci_kms.ExternalKeyManagerMetadata{}

	if externalVaultEndpointUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_vault_endpoint_url")); ok {
		tmp := externalVaultEndpointUrl.(string)
		result.ExternalVaultEndpointUrl = &tmp
	}

	if oauthMetadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oauth_metadata")); ok {
		if tmpList := oauthMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "oauth_metadata"), 0)
			tmp, err := s.mapToOauthMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert oauth_metadata, encountered error: %v", err)
			}
			result.OauthMetadata = &tmp
		}
	}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	return result, nil
}

func ExternalKeyManagerMetadataToMap(obj *oci_kms.ExternalKeyManagerMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExternalVaultEndpointUrl != nil {
		result["external_vault_endpoint_url"] = string(*obj.ExternalVaultEndpointUrl)
	}

	if obj.OauthMetadata != nil {
		result["oauth_metadata"] = []interface{}{OauthMetadataToMap(obj.OauthMetadata)}
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	return result
}

func ExternalKeyManagerMetadataSummaryToMap(obj *oci_kms.ExternalKeyManagerMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExternalVaultEndpointUrl != nil {
		result["external_vault_endpoint_url"] = string(*obj.ExternalVaultEndpointUrl)
	}

	if obj.OauthMetadataSummary != nil {
		result["oauth_metadata_summary"] = []interface{}{OauthMetadataSummaryToMap(obj.OauthMetadataSummary)}
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	return result
}

func (s *KmsVaultResourceCrud) mapToOauthMetadata(fieldKeyFormat string) (oci_kms.OauthMetadata, error) {
	result := oci_kms.OauthMetadata{}

	if clientAppId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_app_id")); ok {
		tmp := clientAppId.(string)
		result.ClientAppId = &tmp
	}

	if clientAppSecret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_app_secret")); ok {
		tmp := clientAppSecret.(string)
		result.ClientAppSecret = &tmp
	}

	if idcsAccountNameUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_account_name_url")); ok {
		tmp := idcsAccountNameUrl.(string)
		result.IdcsAccountNameUrl = &tmp
	}

	return result, nil
}

func OauthMetadataToMap(obj *oci_kms.OauthMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientAppId != nil {
		result["client_app_id"] = string(*obj.ClientAppId)
	}

	if obj.ClientAppSecret != nil {
		result["client_app_secret"] = string(*obj.ClientAppSecret)
	}

	if obj.IdcsAccountNameUrl != nil {
		result["idcs_account_name_url"] = string(*obj.IdcsAccountNameUrl)
	}

	return result
}

func OauthMetadataSummaryToMap(obj *oci_kms.OauthMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientAppId != nil {
		result["client_app_id"] = string(*obj.ClientAppId)
	}

	if obj.IdcsAccountNameUrl != nil {
		result["idcs_account_name_url"] = string(*obj.IdcsAccountNameUrl)
	}

	return result
}

func VaultReplicaDetailsToMap(obj *oci_kms.VaultReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReplicationId != nil {
		result["replication_id"] = string(*obj.ReplicationId)
	}

	return result
}

func (s *KmsVaultResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_kms.ChangeVaultCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VaultId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ChangeVaultCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *KmsVaultResourceCrud) RestoreVaultFromObjectStore() error {
	request := oci_kms.RestoreVaultFromObjectStoreRequest{}

	if backupLocation, ok := s.D.GetOkExists("restore_from_object_store"); ok {
		if tmpList := backupLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "restore_from_object_store", 0)
			tmp, err := s.mapToBackupLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupLocation = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.RestoreVaultFromObjectStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) RestoreVaultFromFile() error {
	request := oci_kms.RestoreVaultFromFileRequest{}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if restoreVaultFromFileDetails, ok := s.D.GetOk("restore_from_file.0.restore_vault_from_file_details"); ok {
		decodedFileContent, _ := base64.StdEncoding.DecodeString(restoreVaultFromFileDetails.(string))
		request.RestoreVaultFromFileDetails = ioutil.NopCloser(bytes.NewBuffer(decodedFileContent))
	} else {
		request.RestoreVaultFromFileDetails = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	if contentLength, ok := s.D.GetOk("restore_from_file.0.content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if contentMd5, ok := s.D.GetOk("restore_from_file.0.content_md5"); ok {
		tmp := contentMd5.(string)
		request.ContentMd5 = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.RestoreVaultFromFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) mapToBackupLocation(fieldKeyFormat string) (oci_kms.BackupLocation, error) {
	var baseObject oci_kms.BackupLocation
	//discriminator
	destinationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination"))
	var destination string
	if ok {
		destination = destinationRaw.(string)
	} else {
		destination = "" // default value
	}
	switch strings.ToLower(destination) {
	case strings.ToLower("BUCKET"):
		details := oci_kms.BackupLocationBucket{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST_URI"):
		details := oci_kms.BackupLocationUri{}
		if uri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown destination '%v' was specified", destination)
	}
	return baseObject, nil
}

func VaultBackupLocationToMap(obj *oci_kms.BackupLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_kms.BackupLocationBucket:
		result["destination"] = "BUCKET"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	case oci_kms.BackupLocationUri:
		result["destination"] = "PRE_AUTHENTICATED_REQUEST_URI"

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'destination' of unknown type %v", *obj)
		return nil
	}

	return result
}
