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
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"regexp"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsKey,
		Read:     readKmsKey,
		Update:   updateKmsKey,
		Delete:   deleteKmsKey,
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
			"key_shape": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"algorithm": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"length": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"curve_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"management_endpoint": {
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
			"external_key_reference": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"external_key_id": {
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
			"protection_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"desired_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
				}, false),
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
						"restore_key_from_file_details": {
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
			"current_key_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_key_reference_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"external_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_key_version_id": {
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
			"restored_from_key_id": {
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
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.CreateResource(d, sync)
}

func readKmsKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		//Import use case:
		id := d.Id()
		regex, _ := regexp.Compile("^managementEndpoint/(.*)/keys/(.*)$")
		tokens := regex.FindStringSubmatch(id)
		if len(tokens) == 3 {
			endpoint = tokens[1]
			d.Set("management_endpoint", endpoint)
			d.SetId(tokens[2])
		} else {
			return fmt.Errorf("id %s should be format: managementEndpoint/{managementEndpoint}/keys/{keyId}", id)
		}
	}

	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

func updateKmsKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.UpdateResource(d, sync)
}

func deleteKmsKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.DeleteResource(d, sync)
}

// existing Id is key OCID and it is not the format that readKmsKey expects managementEndpoint/{managementEndpoint}/keys/{keyId}
// GetCompositeKeyId is only used for resource discovery and it returns the Id as expected by readKmsKey method
// terraform import oci_kms_key.test_key "managementEndpoint/{managementEndpoint}/keys/{keyId}"
func GetCompositeKeyId(managementEndpoint string, keyId string) string {
	compositeId := "managementEndpoint/" + managementEndpoint + "/keys/" + keyId
	return compositeId
}

type KmsKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsManagementClient
	Res                    *oci_kms.Key
	DisableNotFoundRetries bool
}

func (s *KmsKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KmsKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateCreating),
		string(oci_kms.KeyLifecycleStateEnabling),
		string(oci_kms.KeyLifecycleStateRestoring),
		string(oci_kms.KeyLifecycleStateUpdating),
	}
}

func (s *KmsKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabled),
	}
}

func (s *KmsKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateDisabled),
		string(oci_kms.KeyLifecycleStateDeleting),
		string(oci_kms.KeyLifecycleStateSchedulingDeletion),
	}
}

func (s *KmsKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateDeleted),
		string(oci_kms.KeyLifecycleStatePendingDeletion),
	}
}

func (s *KmsKeyResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabling),
		string(oci_kms.KeyLifecycleStateDisabling),
		string(oci_kms.KeyLifecycleStateUpdating),
		string(oci_kms.KeyLifecycleStateRestoring),
	}
}

func (s *KmsKeyResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabled),
		string(oci_kms.KeyLifecycleStateDisabled),
	}
}

func (s *KmsKeyResourceCrud) Create() error {
	if _, ok := s.D.GetOk("restore_from_file"); ok {
		err := s.RestoreKeyFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
		return s.UpdateKeyDetails()
	}
	if _, ok := s.D.GetOk("restore_from_object_store"); ok {
		err := s.RestoreKeyFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())

		return s.UpdateKeyDetails()
	}

	if desiredState, ok := s.D.GetOkExists("desired_state"); ok && !strings.EqualFold(desiredState.(string), "ENABLED") {
		return fmt.Errorf("oci_kms_keys can only be created in ENABLED state")
	}

	request := oci_kms.CreateKeyRequest{}

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

	if externalKeyReference, ok := s.D.GetOkExists("external_key_reference"); ok {
		if tmpList := externalKeyReference.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "external_key_reference", 0)
			tmp, err := s.mapToExternalKeyReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExternalKeyReference = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if keyShape, ok := s.D.GetOkExists("key_shape"); ok {
		if tmpList := keyShape.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_shape", 0)
			tmp, err := s.mapToKeyShape(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KeyShape = &tmp
		}
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_kms.CreateKeyDetailsProtectionModeEnum(protectionMode.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KmsKeyResourceCrud) Get() error {
	request := oci_kms.GetKeyRequest{}

	tmp := s.D.Id()
	request.KeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KmsKeyResourceCrud) Update() error {
	if _, ok := s.D.GetOk("restore_from_file"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreKeyFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	if _, ok := s.D.GetOk("restore_from_object_store"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreKeyFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	return s.UpdateKeyDetails()
}

func (s *KmsKeyResourceCrud) UpdateKeyDetails() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_kms.UpdateKeyRequest{}

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

	request.KeyId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key

	// Handle activation/deactivation here
	if desiredState, ok := s.D.GetOkExists("desired_state"); ok && !strings.EqualFold(desiredState.(string), s.D.Get("state").(string)) {
		desiredStateString := desiredState.(string)

		if desiredStateString == "ENABLED" {
			activationRequest := oci_kms.EnableKeyRequest{}
			tmpId := s.D.Id()
			activationRequest.KeyId = &tmpId

			activationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")
			activationResponse, err := s.Client.EnableKey(context.Background(), activationRequest)
			if err != nil {
				return err
			}
			s.Res = &activationResponse.Key
		} else if desiredStateString == "DISABLED" {
			deactivationRequest := oci_kms.DisableKeyRequest{}
			tmpId := s.D.Id()
			deactivationRequest.KeyId = &tmpId

			deactivationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")
			deactivationResponse, err := s.Client.DisableKey(context.Background(), deactivationRequest)
			if err != nil {
				return err
			}
			s.Res = &deactivationResponse.Key
		}
	}

	return nil
}

func (s *KmsKeyResourceCrud) Delete() error {
	request := oci_kms.ScheduleKeyDeletionRequest{}

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")
	tmp := s.D.Id()
	request.KeyId = &tmp

	_, err := s.Client.ScheduleKeyDeletion(context.Background(), request)
	return err
}

func (s *KmsKeyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentKeyVersion != nil {
		s.D.Set("current_key_version", *s.Res.CurrentKeyVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKeyReferenceDetails != nil {
		s.D.Set("external_key_reference_details", []interface{}{ExternalKeyReferenceDetailsToMap(s.Res.ExternalKeyReferenceDetails)})
	} else {
		s.D.Set("external_key_reference_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	s.D.Set("desired_state", s.Res.LifecycleState)

	if s.Res.KeyShape != nil {
		s.D.Set("key_shape", []interface{}{KeyShapeToMap(s.Res.KeyShape)})
	} else {
		s.D.Set("key_shape", nil)
	}

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	if s.Res.ReplicaDetails != nil {
		s.D.Set("replica_details", []interface{}{KeyReplicaDetailsToMap(s.Res.ReplicaDetails)})
	} else {
		s.D.Set("replica_details", nil)
	}

	if s.Res.RestoredFromKeyId != nil {
		s.D.Set("restored_from_key_id", *s.Res.RestoredFromKeyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func (s *KmsKeyResourceCrud) mapToExternalKeyReference(fieldKeyFormat string) (oci_kms.ExternalKeyReference, error) {
	result := oci_kms.ExternalKeyReference{}

	if externalKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_key_id")); ok {
		tmp := externalKeyId.(string)
		result.ExternalKeyId = &tmp
	}

	return result, nil
}

func ExternalKeyReferenceToMap(obj *oci_kms.ExternalKeyReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExternalKeyId != nil {
		result["external_key_id"] = string(*obj.ExternalKeyId)
	}

	return result
}

func ExternalKeyReferenceDetailsToMap(obj *oci_kms.ExternalKeyReferenceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExternalKeyId != nil {
		result["external_key_id"] = string(*obj.ExternalKeyId)
	}

	if obj.ExternalKeyVersionId != nil {
		result["external_key_version_id"] = string(*obj.ExternalKeyVersionId)
	}

	return result
}

func KeyReplicaDetailsToMap(obj *oci_kms.KeyReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReplicationId != nil {
		result["replication_id"] = string(*obj.ReplicationId)
	}

	return result
}

func (s *KmsKeyResourceCrud) mapToKeyShape(fieldKeyFormat string) (oci_kms.KeyShape, error) {
	result := oci_kms.KeyShape{}

	if algorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "algorithm")); ok {
		result.Algorithm = oci_kms.KeyShapeAlgorithmEnum(algorithm.(string))
	}

	if curveId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "curve_id")); ok {
		result.CurveId = oci_kms.KeyShapeCurveIdEnum(curveId.(string))
	}

	if length, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "length")); ok {
		tmp := length.(int)
		result.Length = &tmp
	}

	return result, nil
}

func KeyShapeToMap(obj *oci_kms.KeyShape) map[string]interface{} {
	result := map[string]interface{}{}

	result["algorithm"] = string(obj.Algorithm)

	result["curve_id"] = string(obj.CurveId)

	if obj.Length != nil {
		result["length"] = int(*obj.Length)
	}

	return result
}

func (s *KmsKeyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_kms.ChangeKeyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KeyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ChangeKeyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *KmsKeyResourceCrud) RestoreKeyFromObjectStore() error {
	request := oci_kms.RestoreKeyFromObjectStoreRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.RestoreKeyFromObjectStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KmsKeyResourceCrud) RestoreKeyFromFile() error {
	request := oci_kms.RestoreKeyFromFileRequest{}
	if restoreKeyFromFileDetails, ok := s.D.GetOk("restore_from_file.0.restore_key_from_file_details"); ok {
		decodedFileContent, _ := base64.StdEncoding.DecodeString(restoreKeyFromFileDetails.(string))
		request.RestoreKeyFromFileDetails = ioutil.NopCloser(bytes.NewBuffer(decodedFileContent))
	} else {
		request.RestoreKeyFromFileDetails = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
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

	response, err := s.Client.RestoreKeyFromFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KmsKeyResourceCrud) mapToBackupLocation(fieldKeyFormat string) (oci_kms.BackupLocation, error) {
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

func KeyBackupLocationToMap(obj *oci_kms.BackupLocation) map[string]interface{} {
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
