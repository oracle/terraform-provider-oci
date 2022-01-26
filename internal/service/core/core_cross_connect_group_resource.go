// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreCrossConnectGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreCrossConnectGroup,
		Read:     readCoreCrossConnectGroup,
		Update:   updateCoreCrossConnectGroup,
		Delete:   deleteCoreCrossConnectGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"customer_reference_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"macsec_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"state": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"encryption_cipher": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"primary_key": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"connectivity_association_key_secret_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"connectivity_association_name_secret_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Required
									"connectivity_association_key_secret_version": {
										Type:     schema.TypeString,
										Required: true,
									},
									"connectivity_association_name_secret_version": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
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

func createCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreCrossConnectGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CrossConnectGroup
	DisableNotFoundRetries bool
}

func (s *CoreCrossConnectGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreCrossConnectGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminating),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminated),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) Create() error {
	request := oci_core.CreateCrossConnectGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if customerReferenceName, ok := s.D.GetOkExists("customer_reference_name"); ok {
		tmp := customerReferenceName.(string)
		request.CustomerReferenceName = &tmp
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

	if macsecProperties, ok := s.D.GetOkExists("macsec_properties"); ok {
		if tmpList := macsecProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "macsec_properties", 0)
			tmp, err := s.mapToCreateMacsecProperties(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MacsecProperties = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Get() error {
	request := oci_core.GetCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	if customerReferenceName, ok := s.D.GetOkExists("customer_reference_name"); ok {
		tmp := customerReferenceName.(string)
		request.CustomerReferenceName = &tmp
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

	if macsecProperties, ok := s.D.GetOkExists("macsec_properties"); ok {
		if tmpList := macsecProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "macsec_properties", 0)
			tmp, err := s.mapToUpdateMacsecProperties(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MacsecProperties = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnectGroup(context.Background(), request)
	return err
}

func (s *CoreCrossConnectGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomerReferenceName != nil {
		s.D.Set("customer_reference_name", *s.Res.CustomerReferenceName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MacsecProperties != nil {
		s.D.Set("macsec_properties", []interface{}{MacsecPropertiesToMap(s.Res.MacsecProperties)})
	} else {
		s.D.Set("macsec_properties", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) mapToCreateMacsecKey(fieldKeyFormat string) (oci_core.CreateMacsecKey, error) {
	result := oci_core.CreateMacsecKey{}

	if connectivityAssociationKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_key_secret_id")); ok {
		tmp := connectivityAssociationKeySecretId.(string)
		result.ConnectivityAssociationKeySecretId = &tmp
	}

	if connectivityAssociationNameSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_name_secret_id")); ok {
		tmp := connectivityAssociationNameSecretId.(string)
		result.ConnectivityAssociationNameSecretId = &tmp
	}

	return result, nil
}

func (s *CoreCrossConnectGroupResourceCrud) mapToUpdateMacsecKey(fieldKeyFormat string) (oci_core.UpdateMacsecKey, error) {
	result := oci_core.UpdateMacsecKey{}

	if connectivityAssociationKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_key_secret_id")); ok {
		tmp := connectivityAssociationKeySecretId.(string)
		result.ConnectivityAssociationKeySecretId = &tmp
	}

	if connectivityAssociationNameSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_name_secret_id")); ok {
		tmp := connectivityAssociationNameSecretId.(string)
		result.ConnectivityAssociationNameSecretId = &tmp
	}

	if connectivityAssociationKeySecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_key_secret_version")); ok {
		tmp := connectivityAssociationKeySecretVersion.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert connectivityAssociationKeySecretVersion string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ConnectivityAssociationKeySecretVersion = &tmpInt64
	}

	if connectivityAssociationNameSecretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connectivity_association_name_secret_version")); ok {
		tmp := connectivityAssociationNameSecretVersion.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert connectivityAssociationNameSecretVersion string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ConnectivityAssociationNameSecretVersion = &tmpInt64
	}

	return result, nil
}

func (s *CoreCrossConnectGroupResourceCrud) mapToCreateMacsecProperties(fieldKeyFormat string) (oci_core.CreateMacsecProperties, error) {
	result := oci_core.CreateMacsecProperties{}

	if encryptionCipher, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_cipher")); ok {
		result.EncryptionCipher = oci_core.MacsecEncryptionCipherEnum(encryptionCipher.(string))
	}

	if primaryKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_key")); ok {
		if tmpList := primaryKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "primary_key"), 0)
			tmp, err := s.mapToCreateMacsecKey(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert primary_key, encountered error: %v", err)
			}
			result.PrimaryKey = &tmp
		}
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		result.State = oci_core.MacsecStateEnum(state.(string))
	}

	return result, nil
}

func (s *CoreCrossConnectGroupResourceCrud) mapToUpdateMacsecProperties(fieldKeyFormat string) (oci_core.UpdateMacsecProperties, error) {
	result := oci_core.UpdateMacsecProperties{}

	if encryptionCipher, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_cipher")); ok {
		result.EncryptionCipher = oci_core.MacsecEncryptionCipherEnum(encryptionCipher.(string))
	}

	if primaryKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_key")); ok {
		if tmpList := primaryKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "primary_key"), 0)
			tmp, err := s.mapToUpdateMacsecKey(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert primary_key, encountered error: %v", err)
			}
			result.PrimaryKey = &tmp
		}
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		result.State = oci_core.MacsecStateEnum(state.(string))
	}

	return result, nil
}

func (s *CoreCrossConnectGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeCrossConnectGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.CrossConnectGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeCrossConnectGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
