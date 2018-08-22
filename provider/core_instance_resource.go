// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &TwoHours,
			Update: &TwoHours,
			Delete: &TwoHours,
		},
		Create: createInstance,
		Read:   readInstance,
		Update: updateInstance,
		Delete: deleteInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"create_vnic_details": {
				Type:     schema.TypeList,
				Optional: true,
				// This must be set to computed, since it's optional and required subnet_id param is being refreshed.
				// If this isn't computed, then that would always force a change on users who do not set create_vnic_details.
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"assign_public_ip": {
							// Change type from boolean to string because TF doesn't handle default
							// values for boolean nested objects correctly.
							Type:     schema.TypeString,
							Optional: true,
							// @CODEGEN 1/2018: Avoid breaking change by setting assign_public_ip to true by default.
							ForceNew: true,
							Default:  "true",
							ValidateFunc: func(v interface{}, k string) ([]string, []error) {
								// Verify that we can parse the string value as a bool value.
								var es []error
								if _, err := strconv.ParseBool(v.(string)); err != nil {
									es = append(es, fmt.Errorf("%s: cannot parse 'assign_public_ip' as bool: %v", k, err))
								}
								return nil, es
							},
							StateFunc: func(v interface{}) string {
								// ValidateFunc runs before StateFunc. Must be valid by now.
								b, _ := NormalizeBoolString(v.(string))
								return b
							},
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: definedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
							// @CODEGEN 6/2018: Remove ForceNew, this is updatable via vnic update
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
							// @CODEGEN 6/2018: Remove ForceNew, this is updatable via vnic update
						},
						"hostname_label": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"skip_source_dest_check": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"hostname_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"image": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: FieldDeprecatedAndOverridenByAnother("image", "source_details"),
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"preserve_boot_volume": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"bootVolume",
								"image",
							}, true),
						},

						// Optional
						"boot_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			// Add this computed boot_volume_id field even though it's not part of the API specs. This will make it easier to
			// discover the attached boot volume's ID; to preserve it for reattachment.
			"boot_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"launch_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"launch_options": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"boot_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_data_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"region": {
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
			// Legacy custom computed convenience values
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.BlockStorageClient = m.(*OracleClients).blockstorageClient

	return CreateResource(d, sync)
}

func readInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.BlockStorageClient = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

func updateInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.BlockStorageClient = m.(*OracleClients).blockstorageClient

	return UpdateResource(d, sync)
}

func deleteInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.BlockStorageClient = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type InstanceResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	BlockStorageClient     *oci_core.BlockstorageClient
	Res                    *oci_core.Instance
	DisableNotFoundRetries bool
}

func (s *InstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateProvisioning),
		string(oci_core.InstanceLifecycleStateStarting),
	}
}

func (s *InstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateRunning),
	}
}

func (s *InstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateTerminating),
	}
}

func (s *InstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateTerminated),
	}
}

func (s *InstanceResourceCrud) Create() error {
	request := oci_core.LaunchInstanceRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
			tmp, err := s.mapToCreateVnicDetailsInstance(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CreateVnicDetails = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if rawExtendedMetadata, ok := s.D.GetOkExists("extended_metadata"); ok {
		extendedMetadata := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		request.ExtendedMetadata = extendedMetadata
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.ImageId = &tmp
	}

	if ipxeScript, ok := s.D.GetOkExists("ipxe_script"); ok {
		tmp := ipxeScript.(string)
		request.IpxeScript = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = objectMapToStringMap(metadata.(map[string]interface{}))
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToInstanceSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.LaunchInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *InstanceResourceCrud) Get() error {
	request := oci_core.GetInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *InstanceResourceCrud) Update() error {
	request := oci_core.UpdateInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance

	// Check for changes in the create_vnic_details sub resource and separately update the vnic

	_, ok := s.D.GetOkExists("create_vnic_details")
	if !s.D.HasChange("create_vnic_details") || !ok {
		log.Printf("[DEBUG] No changes to primary VNIC. Instance ID: %q", s.Res.Id)
		return nil
	}

	vnic, err := s.getPrimaryVnic()
	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be found during instance update: %q (Instance ID: %q, State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
	updateVnicDetails, err := s.mapToUpdateVnicDetailsInstance(fieldKeyFormat)
	if err != nil {
		return err
	}

	vnicOpts := oci_core.UpdateVnicRequest{
		VnicId:            vnic.Id,
		UpdateVnicDetails: updateVnicDetails,
	}

	_, err = s.VirtualNetworkClient.UpdateVnic(context.Background(), vnicOpts)

	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be updated during instance update: %q (Instance ID: %q, State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	return nil
}

func (s *InstanceResourceCrud) Delete() error {
	request := oci_core.TerminateInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	if preserveBootVolume, ok := s.D.GetOkExists("preserve_boot_volume"); ok {
		tmp := preserveBootVolume.(bool)
		request.PreserveBootVolume = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.TerminateInstance(context.Background(), request)
	return err
}

func (s *InstanceResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	// Extended metadata (a json blob) may not return with the same node order in which it
	// was originally created, the solution is to not set it here after subsequent GETS to
	// prevent inadvertent diffs or destroy/creates
	// if s.Res.ExtendedMetadata != nil {
	// // extended_metadata is an arbitrarily structured json object, `objectToMap` would not work
	// 	s.D.Set("extended_metadata", []interface{}{objectToMap(s.Res.ExtendedMetadata)})
	// }

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		// @CODEGEN 1/2018: support legacy name "image"
		s.D.Set("image", *s.Res.ImageId)
	}

	if s.Res.IpxeScript != nil {
		s.D.Set("ipxe_script", *s.Res.IpxeScript)
	}

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
	}

	if s.Res.Metadata != nil {
		err := s.D.Set("metadata", s.Res.Metadata)
		if err != nil {
			log.Printf("error setting metadata %q", err)
		}
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	bootVolume, bootVolumeErr := s.getBootVolume()
	if bootVolumeErr != nil {
		log.Printf("[WARN] Could not get the boot volume: %q", bootVolumeErr)
	}

	if s.Res.SourceDetails != nil {
		var sourceDetailsFromConfig map[string]interface{}
		if details, ok := s.D.GetOkExists("source_details"); ok {
			if tmpList := details.([]interface{}); len(tmpList) > 0 {
				sourceDetailsFromConfig = tmpList[0].(map[string]interface{})
			}
		}
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceSourceDetailsToMap(&s.Res.SourceDetails, bootVolume, sourceDetailsFromConfig); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		err := s.D.Set("source_details", sourceDetailsArray)
		if err != nil {
			return err
		}
	} else {
		s.D.Set("source_details", nil)
	}

	if bootVolume != nil && bootVolume.Id != nil {
		s.D.Set("boot_volume_id", *bootVolume.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.LifecycleState == oci_core.InstanceLifecycleStateRunning {
		vnic, vnicError := s.getPrimaryVnic()
		if vnicError != nil || vnic == nil {
			log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q", vnicError)
		} else {
			s.D.Set("hostname_label", vnic.HostnameLabel)
			s.D.Set("public_ip", vnic.PublicIp)
			s.D.Set("private_ip", vnic.PrivateIp)
			s.D.Set("subnet_id", vnic.SubnetId)

			var createVnicDetails map[string]interface{}
			if details, ok := s.D.GetOkExists("create_vnic_details"); ok {
				if tmpList := details.([]interface{}); len(tmpList) > 0 {
					createVnicDetails = tmpList[0].(map[string]interface{})
				}
			}

			err := s.D.Set("create_vnic_details", []interface{}{vnicDetailsToMap(vnic, createVnicDetails)})
			if err != nil {
				log.Printf("[WARN] create_vnic_details could not be set: %q", err)
			}
		}
	}

	return nil
}

func (s *InstanceResourceCrud) mapToCreateVnicDetailsInstance(fieldKeyFormat string) (oci_core.CreateVnicDetails, error) {
	result := oci_core.CreateVnicDetails{}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(string)
		boolVal, err := strconv.ParseBool(tmp)
		if err != nil {
			return result, err
		}
		result.AssignPublicIp = &boolVal
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *InstanceResourceCrud) mapToUpdateVnicDetailsInstance(fieldKeyFormat string) (oci_core.UpdateVnicDetails, error) {
	result := oci_core.UpdateVnicDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	return result, nil
}

func vnicDetailsToMap(obj *oci_core.Vnic, createVnicDetails map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	// "assign_public_ip" isn't part of the VNIC's state & is only useful at creation time (and
	// subsequent force-new creations). So persist the user-defined value in the config & update it
	// when the user changes that value.
	if createVnicDetails != nil {
		assignPublicIP, _ := NormalizeBoolString(createVnicDetails["assign_public_ip"].(string)) // Must be valid.
		result["assign_public_ip"] = assignPublicIP
	} else {
		// Set to "true" in case "create_vnic_details" is ommited altogether & the default value for
		// "assign_public_ip" doesn't kick in.
		result["assign_public_ip"] = "true"
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SkipSourceDestCheck != nil {
		result["skip_source_dest_check"] = bool(*obj.SkipSourceDestCheck)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *InstanceResourceCrud) mapToInstanceSourceDetails(fieldKeyFormat string) (oci_core.InstanceSourceDetails, error) {
	var baseObject oci_core.InstanceSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("bootVolume"):
		details := oci_core.InstanceSourceViaBootVolumeDetails{}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.BootVolumeId = &tmp
		}
		baseObject = details
	case strings.ToLower("image"):
		details := oci_core.InstanceSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeSizeInGBs = &tmpInt64
		}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.ImageId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InstanceSourceDetailsToMap(obj *oci_core.InstanceSourceDetails, bootVolume *oci_core.BootVolume, sourceDetailsFromConfig map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceSourceViaBootVolumeDetails:
		result["source_type"] = "bootVolume"

		if v.BootVolumeId != nil {
			result["source_id"] = string(*v.BootVolumeId)
		}
	case oci_core.InstanceSourceViaImageDetails:
		result["source_type"] = "image"

		if v.BootVolumeSizeInGBs != nil {
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*v.BootVolumeSizeInGBs, 10)
		} else if bootVolume != nil && bootVolume.SizeInGBs != nil {
			// The service could omit the boot volume size in the InstanceSourceViaImageDetails, so use the boot volume
			// SizeInGBs property if that's the case.
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*bootVolume.SizeInGBs, 10)
		} else if sourceDetailsFromConfig != nil {
			// Last resort. If we can't query the boot volume size from service, use the config value.
			result["boot_volume_size_in_gbs"] = sourceDetailsFromConfig["boot_volume_size_in_gbs"]
		}

		if v.ImageId != nil {
			result["source_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func mapToExtendedMetadata(rm map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range rm {
		val := make(map[string]interface{})
		//Use the string value that was passed if it is not a valid JSON string
		if err := json.Unmarshal([]byte(v.(string)), &val); err == nil {
			result[k] = val
		} else {
			result[k] = v.(string)
		}
	}
	return result
}

func (s *InstanceResourceCrud) getPrimaryVnic() (*oci_core.Vnic, error) {
	request := oci_core.ListVnicAttachmentsRequest{
		CompartmentId: s.Res.CompartmentId,
		InstanceId:    s.Res.Id,
	}

	var attachments []oci_core.VnicAttachment

	for {
		result, err := s.Client.ListVnicAttachments(context.Background(), request)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, result.Items...)
		request.Page = result.OpcNextPage

		if request.Page == nil {
			break
		}
	}

	if len(attachments) < 1 {
		return nil, errors.New("No VNIC attachments found.")
	}

	for _, attachment := range attachments {
		if attachment.LifecycleState == oci_core.VnicAttachmentLifecycleStateAttached {
			request := oci_core.GetVnicRequest{VnicId: attachment.VnicId}
			response, _ := s.VirtualNetworkClient.GetVnic(context.Background(), request)
			vnic := &response.Vnic

			// Ignore errors on GetVnic, since we might not have permissions to view some secondary VNICs.
			if vnic != nil && vnic.IsPrimary != nil && *vnic.IsPrimary {
				return vnic, nil
			}
		}
	}

	return nil, errors.New("Primary VNIC not found.")
}

func (s *InstanceResourceCrud) getBootVolume() (*oci_core.BootVolume, error) {
	request := oci_core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: s.Res.AvailabilityDomain,
		CompartmentId:      s.Res.CompartmentId,
		InstanceId:         s.Res.Id,
	}

	response, err := s.Client.ListBootVolumeAttachments(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if len(response.Items) < 1 {
		return nil, fmt.Errorf("Could not find any attached boot volumes")
	}

	bootVolumeId := response.Items[0].BootVolumeId
	if bootVolumeId == nil {
		return nil, fmt.Errorf("Found a boot volume attachment with no boot volume ID")
	}

	bootVolumeRequest := oci_core.GetBootVolumeRequest{BootVolumeId: bootVolumeId}
	bootVolumeResponse, err := s.BlockStorageClient.GetBootVolume(context.Background(), bootVolumeRequest)
	if err != nil {
		return nil, err
	}

	return &bootVolumeResponse.BootVolume, nil
}
