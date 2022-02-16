// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreComputeImageCapabilitySchemaResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeImageCapabilitySchema,
		Read:     readCoreComputeImageCapabilitySchema,
		Update:   updateCoreComputeImageCapabilitySchema,
		Delete:   deleteCoreComputeImageCapabilitySchema,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_global_image_capability_schema_version_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schema_data": {
				Type:             schema.TypeMap,
				Required:         true,
				DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
				Elem:             schema.TypeString,
			},

			// Optional
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

			// Computed
			"compute_global_image_capability_schema_id": {
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

func createCoreComputeImageCapabilitySchema(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeImageCapabilitySchema(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeImageCapabilitySchema(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeImageCapabilitySchema(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreComputeImageCapabilitySchemaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeImageCapabilitySchema
	DisableNotFoundRetries bool
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) Create() error {
	request := oci_core.CreateComputeImageCapabilitySchemaRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeGlobalImageCapabilitySchemaVersionName, ok := s.D.GetOkExists("compute_global_image_capability_schema_version_name"); ok {
		tmp := computeGlobalImageCapabilitySchemaVersionName.(string)
		request.ComputeGlobalImageCapabilitySchemaVersionName = &tmp
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

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if schemaData, ok := s.D.GetOkExists("schema_data"); ok {
		schemaData, err := mapToSchemaData(schemaData.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("unable to convert schema_data, encountered error: %v", err)
		}
		request.SchemaData = schemaData
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateComputeImageCapabilitySchema(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeImageCapabilitySchema
	return nil
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) Get() error {
	request := oci_core.GetComputeImageCapabilitySchemaRequest{}

	tmp := s.D.Id()
	request.ComputeImageCapabilitySchemaId = &tmp

	if isMergeEnabled, ok := s.D.GetOkExists("is_merge_enabled"); ok {
		tmp := isMergeEnabled.(bool)
		request.IsMergeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeImageCapabilitySchema(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeImageCapabilitySchema
	return nil
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeImageCapabilitySchemaRequest{}

	tmp := s.D.Id()
	request.ComputeImageCapabilitySchemaId = &tmp

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

	if schemaData, ok := s.D.GetOkExists("schema_data"); ok {
		schemaData, err := mapToSchemaData(schemaData.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("unable to convert schema_data, encountered error: %v", err)
		}
		request.SchemaData = schemaData
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeImageCapabilitySchema(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeImageCapabilitySchema
	return nil
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) Delete() error {
	request := oci_core.DeleteComputeImageCapabilitySchemaRequest{}

	tmp := s.D.Id()
	request.ComputeImageCapabilitySchemaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteComputeImageCapabilitySchema(context.Background(), request)
	return err
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeGlobalImageCapabilitySchemaId != nil {
		s.D.Set("compute_global_image_capability_schema_id", *s.Res.ComputeGlobalImageCapabilitySchemaId)
	}

	if s.Res.ComputeGlobalImageCapabilitySchemaVersionName != nil {
		s.D.Set("compute_global_image_capability_schema_version_name", *s.Res.ComputeGlobalImageCapabilitySchemaVersionName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.SchemaData != nil {
		s.D.Set("schema_data", schemaDataToMap(s.Res.SchemaData))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreComputeImageCapabilitySchemaResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeImageCapabilitySchemaCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeImageCapabilitySchemaId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeComputeImageCapabilitySchemaCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

type imagecapabilityschemadescriptor struct {
	JsonData       []byte
	Source         oci_core.ImageCapabilitySchemaDescriptorSourceEnum
	DescriptorType string
}

func mapToSchemaData(rm map[string]interface{}) (map[string]oci_core.ImageCapabilitySchemaDescriptor, error) {
	result := make(map[string]oci_core.ImageCapabilitySchemaDescriptor)
	for k, v := range rm {
		var val imagecapabilityschemadescriptor
		if err := json.Unmarshal([]byte(v.(string)), &val); err == nil {
			if schemaData, err := UnmarshalPolymorphicJSON(val.DescriptorType, v); err == nil {
				result[k] = schemaData
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func UnmarshalPolymorphicJSON(descriptorType string, val interface{}) (oci_core.ImageCapabilitySchemaDescriptor, error) {
	var err error
	switch descriptorType {
	case "enumstring":
		mm := oci_core.EnumStringImageCapabilitySchemaDescriptor{}
		err = json.Unmarshal([]byte(val.(string)), &mm)
		return mm, err
	case "enuminteger":
		mm := oci_core.EnumIntegerImageCapabilityDescriptor{}
		err = json.Unmarshal([]byte(val.(string)), &mm)
		return mm, err
	case "boolean":
		mm := oci_core.BooleanImageCapabilitySchemaDescriptor{}
		err = json.Unmarshal([]byte(val.(string)), &mm)
		return mm, err
	default:
		return nil, nil
	}
}

func schemaDataToMap(genericMap map[string]oci_core.ImageCapabilitySchemaDescriptor) map[string]interface{} {
	result := map[string]interface{}{}

	for key, value := range genericMap {
		bytes, err := json.Marshal(value)
		if err != nil {
			continue
		}
		result[key] = string(bytes)
	}

	return result
}
