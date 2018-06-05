// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VnicAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVnicAttachment,
		Read:     readVnicAttachment,
		Update:   updateVnicAttachment,
		Delete:   deleteVnicAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"create_vnic_details": {
				Type:     schema.TypeList,
				Required: true,
				// @CODEGEN 1/2018: Generator says create_vnic_details is a ForceNew property. Remove it to avoid
				// a breaking change with existing provider, which allows some vnic properties to be updated.
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
							Default:  "true",
							ForceNew: true,
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
								b, _ := crud.NormalizeBoolString(v.(string))
								return b
							},
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: definedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
							// @CODEGEN 6/2018: Remove ForceNew for this attribute, it can be updated.
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							// @CODEGEN 1/2018: Remove ForceNew for this attribute, it can be updated.
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
							// @CODEGEN 6/2018: Remove ForceNew for this attribute, it can be updated.
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							// @CODEGEN 1/2018: Remove ForceNew for this attribute, it can be updated.
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
							// @CODEGEN 1/2018: Remove Computed and ForceNew for this attribute, it can be updated
							// and it should be false by default to avoid a breaking change.
							Default: false,
						},

						// Computed
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nic_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_tag": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VnicAttachmentResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	Res                    *oci_core.VnicAttachment
	DisableNotFoundRetries bool
}

func (s *VnicAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VnicAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateAttaching),
	}
}

func (s *VnicAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateAttached),
	}
}

func (s *VnicAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateDetaching),
	}
}

func (s *VnicAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateDetached),
	}
}

func (s *VnicAttachmentResourceCrud) Create() error {
	request := oci_core.AttachVnicRequest{}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			tmp, err := mapToCreateVnicDetails(tmpList[0].(map[string]interface{}))
			if err != nil {
				return err
			}
			request.CreateVnicDetails = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if nicIndex, ok := s.D.GetOkExists("nic_index"); ok {
		tmp := nicIndex.(int)
		request.NicIndex = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachVnic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VnicAttachment
	return nil
}

// @CODEGEN 1/2018: Generator doesn't give us an Update method for VnicAttachment.
// However, the existing behavior allows vnic to be updated through the create_vnic_details.
// So keep this Update functionality in the provider.
func (s *VnicAttachmentResourceCrud) Update() error {
	// We should fetch the VnicAttachment in order to update
	// the state data after the update call.
	err := s.Get()
	if err != nil {
		return err
	}

	request := oci_core.UpdateVnicRequest{}

	if s.Res.VnicId != nil {
		request.VnicId = s.Res.VnicId
	}

	if !s.D.HasChange("create_vnic_details") {
		return nil
	}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			tmp, err := mapToUpdateVnicDetails(tmpList[0].(map[string]interface{}))
			if err != nil {
				return err
			}

			request.UpdateVnicDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err = s.VirtualNetworkClient.UpdateVnic(context.Background(), request)
	return err
}

func (s *VnicAttachmentResourceCrud) Get() error {
	request := oci_core.GetVnicAttachmentRequest{}

	tmp := s.D.Id()
	request.VnicAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVnicAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VnicAttachment
	return nil
}

func (s *VnicAttachmentResourceCrud) Delete() error {
	request := oci_core.DetachVnicRequest{}

	tmp := s.D.Id()
	request.VnicAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DetachVnic(context.Background(), request)
	return err
}

func (s *VnicAttachmentResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	if s.Res.NicIndex != nil {
		s.D.Set("nic_index", *s.Res.NicIndex)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VlanTag != nil {
		s.D.Set("vlan_tag", *s.Res.VlanTag)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	// @CODEGEN 1/2018: We need to refresh the vnic details after every refresh.
	request := oci_core.GetVnicRequest{}
	request.VnicId = s.Res.VnicId
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.VirtualNetworkClient.GetVnic(context.Background(), request)
	if err != nil {
		// VNIC might not be found when attaching or detaching.
		log.Printf("[DEBUG] VNIC not found during VNIC Attachment refresh. (VNIC ID: %q, Error: %q)", *request.VnicId, err)
		return
	}

	var createVnicDetails map[string]interface{}
	if details, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			createVnicDetails = tmpList[0].(map[string]interface{})
		}
	}

	if err := s.D.Set("create_vnic_details", []interface{}{VnicDetailsToMap(&response.Vnic, createVnicDetails)}); err != nil {
		log.Printf("Unable to refresh create_vnic_details. Error: %q", err)
	}
}

func mapToCreateVnicDetails(raw map[string]interface{}) (oci_core.CreateVnicDetails, error) {
	result := oci_core.CreateVnicDetails{}

	if assignPublicIp, ok := raw["assign_public_ip"]; ok {
		tmp := assignPublicIp.(string)
		boolVal, _ := strconv.ParseBool(tmp) // Must be valid.
		result.AssignPublicIp = &boolVal
	}

	if definedTags, ok := raw["defined_tags"]; ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := raw["display_name"]; ok && displayName != "" {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := raw["freeform_tags"]; ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := raw["hostname_label"]; ok && hostnameLabel != "" {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if privateIp, ok := raw["private_ip"]; ok && privateIp != "" {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := raw["skip_source_dest_check"]; ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := raw["subnet_id"]; ok && subnetId != "" {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func mapToUpdateVnicDetails(raw map[string]interface{}) (oci_core.UpdateVnicDetails, error) {
	result := oci_core.UpdateVnicDetails{}

	if definedTags, ok := raw["defined_tags"]; ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := raw["display_name"]; ok {
		tmp := displayName.(string)
		if tmp != "" {
			result.DisplayName = &tmp
		}
	}

	if freeformTags, ok := raw["freeform_tags"]; ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := raw["hostname_label"]; ok {
		tmp := hostnameLabel.(string)
		if tmp != "" {
			result.HostnameLabel = &tmp
		}
	}

	if skipSourceDestCheck, ok := raw["skip_source_dest_check"]; ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	return result, nil
}

func VnicDetailsToMap(obj *oci_core.Vnic, createVnicDetails map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	// "assign_public_ip" isn't part of the VNIC's state & is only useful at creation time (and
	// subsequent force-new creations). So persist the user-defined value in the config & update it
	// when the user changes that value.
	if createVnicDetails != nil {
		assignPublicIP, _ := crud.NormalizeBoolString(createVnicDetails["assign_public_ip"].(string)) // Must be valid.
		result["assign_public_ip"] = assignPublicIP
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
