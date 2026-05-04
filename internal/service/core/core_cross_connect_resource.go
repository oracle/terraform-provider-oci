// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreCrossConnectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreCrossConnect,
		Read:     readCoreCrossConnect,
		Update:   updateCoreCrossConnect,
		Delete:   deleteCoreCrossConnect,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_speed_shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"far_cross_connect_or_cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"interface_down_timer_value_in_milliseconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_interface_hold_timer_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_qos_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"loa_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorized_agent": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiry_extension_count": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntAtLeast(0),
						},
					},
				},
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
						"is_unprotected_traffic_allowed": {
							Type:     schema.TypeBool,
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

									// Optional

									// Computed
									"connectivity_association_key_secret_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"connectivity_association_name_secret_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},

						// Computed
					},
				},
			},
			"near_cross_connect_or_cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"oci_physical_device_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_active": {
				Type:         schema.TypeBool,
				Optional:     true,
				ValidateFunc: tfresource.ValidateBoolInSlice([]bool{true}),
			},

			// Computed
			"oci_logical_device_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_name": {
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

func createCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	// Issue an Update if 'is_active' is set to true
	if _, ok := sync.D.GetOkExists("is_active"); ok {
		log.Printf("[DEBUG] CrossConnect resource is set to be active, calling 'Update' for the resource")
		if err := tfresource.UpdateResource(d, sync); err != nil {
			return err
		}
	}

	return readCoreCrossConnect(d, m)
}

func readCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreCrossConnectResourceCrud struct {
	tfresource.BaseCrud
	Client                               *oci_core.VirtualNetworkClient
	Res                                  *oci_core.CrossConnect
	Loa                                  *oci_core.LetterOfAuthority
	DisableNotFoundRetries               bool
	letterOfAuthorityUpdatedDuringCreate bool
}

func (s *CoreCrossConnectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreCrossConnectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminating),
	}
}

func (s *CoreCrossConnectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminated),
	}
}

func (s *CoreCrossConnectResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectResourceCrud) UpdatedTarget() []string {
	if _, ok := s.D.GetOkExists("is_active"); ok {
		log.Printf("[DEBUG] CrossConnect resource is set to be active, wait until the state is '%s'", string(oci_core.CrossConnectLifecycleStateProvisioned))
		return []string{
			string(oci_core.CrossConnectLifecycleStateProvisioned),
		}
	}

	return []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectResourceCrud) Create() error {
	request := oci_core.CreateCrossConnectRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		tmp := crossConnectGroupId.(string)
		request.CrossConnectGroupId = &tmp
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

	if farCrossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists("far_cross_connect_or_cross_connect_group_id"); ok {
		tmp := farCrossConnectOrCrossConnectGroupId.(string)
		request.FarCrossConnectOrCrossConnectGroupId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if interfaceDownTimerValueInMilliseconds, ok := s.D.GetOkExists("interface_down_timer_value_in_milliseconds"); ok {
		tmp := interfaceDownTimerValueInMilliseconds.(int)
		request.InterfaceDownTimerValueInMilliseconds = &tmp
	}

	if interfaceName, ok := s.D.GetOkExists("interface_name"); ok {
		tmp := interfaceName.(string)
		request.InterfaceName = &tmp
	}

	if isInterfaceHoldTimerEnabled, ok := s.D.GetOkExists("is_interface_hold_timer_enabled"); ok {
		tmp := isInterfaceHoldTimerEnabled.(bool)
		request.IsInterfaceHoldTimerEnabled = &tmp
	}

	if isQosEnabled, ok := s.D.GetOkExists("is_qos_enabled"); ok {
		tmp := isQosEnabled.(bool)
		request.IsQosEnabled = &tmp
	}

	if locationName, ok := s.D.GetOkExists("location_name"); ok {
		tmp := locationName.(string)
		request.LocationName = &tmp
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

	if nearCrossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists("near_cross_connect_or_cross_connect_group_id"); ok {
		tmp := nearCrossConnectOrCrossConnectGroupId.(string)
		request.NearCrossConnectOrCrossConnectGroupId = &tmp
	}

	if ociPhysicalDeviceName, ok := s.D.GetOkExists("oci_physical_device_name"); ok {
		tmp := ociPhysicalDeviceName.(string)
		request.OciPhysicalDeviceName = &tmp
	}

	if portSpeedShapeName, ok := s.D.GetOkExists("port_speed_shape_name"); ok {
		tmp := portSpeedShapeName.(string)
		request.PortSpeedShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	if isAttributeExplicitlyConfiguredInRawConfig(s.D, "loa_properties") {
		if err := s.updateLetterOfAuthorityPropertiesForCrossConnectId(*response.CrossConnect.Id, false); err != nil {
			return err
		}
		s.letterOfAuthorityUpdatedDuringCreate = true
	} else {
		if loa, err := s.getLetterOfAuthority(*response.CrossConnect.Id); err != nil {
			return err
		} else {
			s.Loa = loa
		}
	}
	return nil
}

func (s *CoreCrossConnectResourceCrud) Get() error {
	request := oci_core.GetCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	if loa, err := s.getLetterOfAuthority(*request.CrossConnectId); err != nil {
		return err
	} else {
		s.Loa = loa
	}
	return nil
}

func (s *CoreCrossConnectResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if interfaceDownTimerValueInMilliseconds, ok := s.D.GetOkExists("interface_down_timer_value_in_milliseconds"); ok {
		tmp := interfaceDownTimerValueInMilliseconds.(int)
		request.InterfaceDownTimerValueInMilliseconds = &tmp
	}

	// Cross Connect Resource can be set to 'Active' only once when the resource is 'PENDING_CUSTOMER' and not 'PROVISIONED'
	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		if state, ok := s.D.GetOkExists("state"); ok && state.(string) == string(oci_core.CrossConnectLifecycleStatePendingCustomer) {
			log.Printf("[DEBUG] Cross Connect is in a valid state: '%s' to be set to active", state.(string))
			tmp := isActive.(bool)
			request.IsActive = &tmp
		}
	}

	if isInterfaceHoldTimerEnabled, ok := s.D.GetOkExists("is_interface_hold_timer_enabled"); ok {
		tmp := isInterfaceHoldTimerEnabled.(bool)
		request.IsInterfaceHoldTimerEnabled = &tmp
	}

	if macsecProperties, ok := s.D.GetOk("macsec_properties"); ok {
		if s.shouldIncludeMacsecPropertiesInUpdate() {
			if tmpList := macsecProperties.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "macsec_properties", 0)
				tmp, err := s.mapToUpdateMacsecProperties(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.MacsecProperties = &tmp
			}
		} else {
			log.Printf("[DEBUG] omitting macsec_properties from UpdateCrossConnect request because cross_connect_group_id is set and macsec_properties is not explicitly configured")
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	if s.D.HasChange("loa_properties") && !s.letterOfAuthorityUpdatedDuringCreate {
		if err := s.updateLetterOfAuthorityProperties(true); err != nil {
			return err
		}
	}
	return nil
}

func (s *CoreCrossConnectResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnect(context.Background(), request)
	return err
}

func (s *CoreCrossConnectResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CrossConnectGroupId != nil {
		s.D.Set("cross_connect_group_id", *s.Res.CrossConnectGroupId)
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

	if s.Res.InterfaceDownTimerValueInMilliseconds != nil {
		s.D.Set("interface_down_timer_value_in_milliseconds", *s.Res.InterfaceDownTimerValueInMilliseconds)
	}

	if s.Res.InterfaceName != nil {
		s.D.Set("interface_name", *s.Res.InterfaceName)
	}

	if s.Res.IsInterfaceHoldTimerEnabled != nil {
		s.D.Set("is_interface_hold_timer_enabled", *s.Res.IsInterfaceHoldTimerEnabled)
	}

	if s.Res.IsQosEnabled != nil {
		s.D.Set("is_qos_enabled", *s.Res.IsQosEnabled)
	}

	if s.Res.LocationName != nil {
		s.D.Set("location_name", *s.Res.LocationName)
	}

	if s.Res.MacsecProperties != nil {
		s.D.Set("macsec_properties", []interface{}{MacsecPropertiesToMap(s.Res.MacsecProperties)})
	} else {
		s.D.Set("macsec_properties", nil)
	}

	if s.Res.OciLogicalDeviceName != nil {
		s.D.Set("oci_logical_device_name", *s.Res.OciLogicalDeviceName)
	}

	if s.Res.OciPhysicalDeviceName != nil {
		s.D.Set("oci_physical_device_name", *s.Res.OciPhysicalDeviceName)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.PortSpeedShapeName != nil {
		s.D.Set("port_speed_shape_name", *s.Res.PortSpeedShapeName)
	}

	if s.Loa != nil && (isAttributeExplicitlyConfiguredInRawConfig(s.D, "loa_properties") || hasMeaningfulLetterOfAuthorityProperties(s.Loa)) {
		s.D.Set("loa_properties", []interface{}{s.letterOfAuthorityPropertiesToMap(s.Loa)})
	} else {
		s.D.Set("loa_properties", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreCrossConnectResourceCrud) mapToCreateMacsecKey(fieldKeyFormat string) (oci_core.CreateMacsecKey, error) {
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

func (s *CoreCrossConnectResourceCrud) mapToUpdateMacsecKey(fieldKeyFormat string) (oci_core.UpdateMacsecKey, error) {
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

func MacsecKeyToMap(obj *oci_core.MacsecKey) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectivityAssociationKeySecretId != nil {
		result["connectivity_association_key_secret_id"] = string(*obj.ConnectivityAssociationKeySecretId)
	}

	if obj.ConnectivityAssociationKeySecretVersion != nil {
		result["connectivity_association_key_secret_version"] = strconv.FormatInt(*obj.ConnectivityAssociationKeySecretVersion, 10)
	}

	if obj.ConnectivityAssociationNameSecretId != nil {
		result["connectivity_association_name_secret_id"] = string(*obj.ConnectivityAssociationNameSecretId)
	}

	if obj.ConnectivityAssociationNameSecretVersion != nil {
		result["connectivity_association_name_secret_version"] = strconv.FormatInt(*obj.ConnectivityAssociationNameSecretVersion, 10)
	}

	return result
}

func (s *CoreCrossConnectResourceCrud) mapToCreateMacsecProperties(fieldKeyFormat string) (oci_core.CreateMacsecProperties, error) {
	result := oci_core.CreateMacsecProperties{}

	if encryptionCipher, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_cipher")); ok {
		result.EncryptionCipher = oci_core.MacsecEncryptionCipherEnum(encryptionCipher.(string))
	}

	if isUnprotectedTrafficAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unprotected_traffic_allowed")); ok {
		tmp := isUnprotectedTrafficAllowed.(bool)
		result.IsUnprotectedTrafficAllowed = &tmp
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

func (s *CoreCrossConnectResourceCrud) mapToUpdateMacsecProperties(fieldKeyFormat string) (oci_core.UpdateMacsecProperties, error) {
	result := oci_core.UpdateMacsecProperties{}

	if encryptionCipher, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_cipher")); ok {
		result.EncryptionCipher = oci_core.MacsecEncryptionCipherEnum(encryptionCipher.(string))
	}

	if isUnprotectedTrafficAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unprotected_traffic_allowed")); ok {
		tmp := isUnprotectedTrafficAllowed.(bool)
		result.IsUnprotectedTrafficAllowed = &tmp
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

func MacsecPropertiesToMap(obj *oci_core.MacsecProperties) map[string]interface{} {
	result := map[string]interface{}{}

	result["encryption_cipher"] = string(obj.EncryptionCipher)

	if obj.IsUnprotectedTrafficAllowed != nil {
		result["is_unprotected_traffic_allowed"] = bool(*obj.IsUnprotectedTrafficAllowed)
	}

	if obj.PrimaryKey != nil {
		result["primary_key"] = []interface{}{MacsecKeyToMap(obj.PrimaryKey)}
	}

	result["state"] = string(obj.State)

	return result
}

func (s *CoreCrossConnectResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeCrossConnectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.CrossConnectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeCrossConnectCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *CoreCrossConnectResourceCrud) getLetterOfAuthority(crossConnectId string) (*oci_core.LetterOfAuthority, error) {
	return getCrossConnectLetterOfAuthority(s.Client, crossConnectId, s.DisableNotFoundRetries)
}

func (s *CoreCrossConnectResourceCrud) updateLetterOfAuthorityProperties(requireChange bool) error {
	return s.updateLetterOfAuthorityPropertiesForCrossConnectId(s.D.Id(), requireChange)
}

func (s *CoreCrossConnectResourceCrud) updateLetterOfAuthorityPropertiesForCrossConnectId(crossConnectId string, requireChange bool) error {
	if !s.hasLetterOfAuthorityPropertiesBlock() {
		return nil
	}

	currentExtensionCount := 0
	if requireChange {
		oldRaw, _ := s.D.GetChange("loa_properties.0.expiry_extension_count")
		if oldRaw != nil {
			currentExtensionCount = oldRaw.(int)
		}
	}
	desiredExtensionCount := currentExtensionCount
	desiredExtensionCountRaw, desiredExtensionCountOk := s.D.GetOkExists("loa_properties.0.expiry_extension_count")
	if requireChange && !s.D.HasChange("loa_properties.0.expiry_extension_count") {
		desiredExtensionCountOk = false
	}
	if desiredExtensionCountOk {
		desiredExtensionCount = desiredExtensionCountRaw.(int)
	}

	if desiredExtensionCount < currentExtensionCount {
		return fmt.Errorf("loa_properties.expiry_extension_count cannot be decreased from %d to %d", currentExtensionCount, desiredExtensionCount)
	}

	extensionDelta := desiredExtensionCount - currentExtensionCount
	if extensionDelta > 1 {
		return fmt.Errorf("loa_properties.expiry_extension_count can only be increased by 1 per update, got increase from %d to %d", currentExtensionCount, desiredExtensionCount)
	}

	var authorizedAgentRaw interface{}
	var hasAuthorizedAgent bool
	if requireChange {
		authorizedAgentRaw, hasAuthorizedAgent = s.D.GetOkExists("loa_properties.0.authorized_agent")
		if !s.D.HasChange("loa_properties.0.authorized_agent") {
			hasAuthorizedAgent = false
		}
	} else {
		authorizedAgentRaw, hasAuthorizedAgent = s.D.GetOk("loa_properties.0.authorized_agent")
	}
	authorizedAgent := ""
	if hasAuthorizedAgent {
		authorizedAgent = authorizedAgentRaw.(string)
	}
	removeAuthorizedAgent := hasAuthorizedAgent && authorizedAgent == ""

	if extensionDelta == 0 && !hasAuthorizedAgent {
		return nil
	}

	if extensionDelta > 0 || hasAuthorizedAgent {
		request := s.newUpdateLetterOfAuthorityRequest(crossConnectId)
		if extensionDelta > 0 {
			shouldExtend := true
			request.ShouldExtend = &shouldExtend
		}
		if hasAuthorizedAgent {
			if removeAuthorizedAgent {
				shouldRemoveAuthorizedAgent := true
				request.ShouldRemoveAuthorizedAgent = &shouldRemoveAuthorizedAgent
			} else {
				request.AuthorizedAgent = &authorizedAgent
			}
		}
		if _, err := s.Client.UpdateCrossConnectLetterOfAuthority(context.Background(), request); err != nil {
			return err
		}
	}

	updatedLoa, err := s.getLetterOfAuthority(crossConnectId)
	if err != nil {
		return err
	}
	s.Loa = updatedLoa

	return nil
}

func (s *CoreCrossConnectResourceCrud) newUpdateLetterOfAuthorityRequest(crossConnectId string) oci_core.UpdateCrossConnectLetterOfAuthorityRequest {
	request := oci_core.UpdateCrossConnectLetterOfAuthorityRequest{}
	request.CrossConnectId = &crossConnectId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	return request
}

func (s *CoreCrossConnectResourceCrud) hasLetterOfAuthorityPropertiesBlock() bool {
	loaProperties, ok := s.D.GetOkExists("loa_properties")
	if !ok {
		return false
	}

	tmpList := loaProperties.([]interface{})
	if len(tmpList) == 0 || tmpList[0] == nil {
		return false
	}

	_, ok = tmpList[0].(map[string]interface{})
	return ok
}

func LetterOfAuthorityPropertiesToMap(obj *oci_core.LetterOfAuthority) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthorizedAgent != nil {
		result["authorized_agent"] = string(*obj.AuthorizedAgent)
	}

	return result
}

func (s *CoreCrossConnectResourceCrud) letterOfAuthorityPropertiesToMap(obj *oci_core.LetterOfAuthority) map[string]interface{} {
	result := LetterOfAuthorityPropertiesToMap(obj)

	if expiryExtensionCount, ok := s.D.GetOkExists("loa_properties.0.expiry_extension_count"); ok {
		result["expiry_extension_count"] = expiryExtensionCount.(int)
	}

	return result
}

func hasMeaningfulLetterOfAuthorityProperties(obj *oci_core.LetterOfAuthority) bool {
	if obj == nil {
		return false
	}

	if obj.AuthorizedAgent != nil && *obj.AuthorizedAgent != "" {
		return true
	}

	return false
}

func (s *CoreCrossConnectResourceCrud) shouldIncludeMacsecPropertiesInUpdate() bool {
	if !s.isPartOfCrossConnectGroup() {
		return true
	}

	if isAttributeExplicitlyConfiguredInRawConfig(s.D, "macsec_properties") {
		return true
	}

	return false
}

func (s *CoreCrossConnectResourceCrud) isPartOfCrossConnectGroup() bool {
	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		return crossConnectGroupId.(string) != ""
	}
	return false
}

func isAttributeExplicitlyConfiguredInRawConfig(d *schema.ResourceData, attributeName string) bool {
	rawConfig := d.GetRawConfig()
	if !rawConfig.IsKnown() || rawConfig.IsNull() || !rawConfig.Type().IsObjectType() {
		return false
	}

	rawAttribute := rawConfig.GetAttr(attributeName)
	if !rawAttribute.IsKnown() || rawAttribute.IsNull() {
		return false
	}

	if rawAttribute.Type().IsListType() || rawAttribute.Type().IsSetType() || rawAttribute.Type().IsTupleType() {
		return rawAttribute.LengthInt() > 0
	}

	if rawAttribute.Type().IsMapType() || rawAttribute.Type().IsObjectType() {
		return len(rawAttribute.AsValueMap()) > 0
	}

	return true
}
