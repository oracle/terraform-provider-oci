// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.TwoHours,
			Update: &crud.TwoHours,
			Delete: &crud.TwoHours,
		},
		Create: createImage,
		Read:   readImage,
		Update: updateImage,
		Delete: deleteImage,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// TODO: New optional, polymorphic field. Skip for now.
			// Tracked here: https://jira.aka.lgl.grungy.us/browse/ORCH-678
			//
			// "image_source_details": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	Computed: true,
			// 	ForceNew: true,
			// 	MaxItems: 1,
			// 	MinItems: 1,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			// Required
			// 			"source_type": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 				ForceNew: true,
			// 			},

			// 			// Optional

			// 			// Computed
			// 		},
			// 	},
			// },
			"instance_id": {
				Type:     schema.TypeString,
				Required: true, // Changed from optional/computed to required till "imageSourceDetails" is supported.
				ForceNew: true,
			},

			// Computed
			"base_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_image_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system_version": {
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

func createImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.CreateResource(d, sync)
}

func readImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

func updateImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.UpdateResource(d, sync)
}

func deleteImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ImageResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.Image
	DisableNotFoundRetries bool
}

func (s *ImageResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ImageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateProvisioning),
		string(oci_core.ImageLifecycleStateImporting),
	}
}

func (s *ImageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateAvailable),
	}
}

func (s *ImageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDisabled),
	}
}

func (s *ImageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDeleted),
	}
}

func (s *ImageResourceCrud) Create() error {
	request := oci_core.CreateImageRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// TODO: New optional, polymorphic field. Skip for now.
	// Tracked here: https://jira.aka.lgl.grungy.us/browse/ORCH-678
	// if imageSourceDetails, ok := s.D.GetOkExists("image_source_details"); ok {
	// 	if tmpList := imageSourceDetails.([]interface{}); len(tmpList) > 0 {
	// 		tmp := mapToImageSourceDetails(tmpList[0].(map[string]interface{}))
	// 		request.ImageSourceDetails = &tmp
	// 	}
	// }

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	response, err := s.Client.CreateImage(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *ImageResourceCrud) Get() error {
	request := oci_core.GetImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	response, err := s.Client.GetImage(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *ImageResourceCrud) Update() error {
	request := oci_core.UpdateImageRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ImageId = &tmp

	response, err := s.Client.UpdateImage(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *ImageResourceCrud) Delete() error {
	request := oci_core.DeleteImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	_, err := s.Client.DeleteImage(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	return err
}

func (s *ImageResourceCrud) SetData() {
	if s.Res.BaseImageId != nil {
		s.D.Set("base_image_id", *s.Res.BaseImageId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateImageAllowed != nil {
		s.D.Set("create_image_allowed", *s.Res.CreateImageAllowed)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.OperatingSystem != nil {
		s.D.Set("operating_system", *s.Res.OperatingSystem)
	}

	if s.Res.OperatingSystemVersion != nil {
		s.D.Set("operating_system_version", *s.Res.OperatingSystemVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

}

// TODO: The following 2 functions are used by the ImageSourceDetails field which was skipped.
// Tracked here: https://jira.aka.lgl.grungy.us/browse/ORCH-678
// func mapToImageSourceDetails(raw map[string]interface{}) oci_core.ImageSourceDetails {
// 	result := oci_core.ImageSourceDetails{}

// 	if sourceType, ok := raw["source_type"]; ok {
// 		tmp := sourceType.(string)
// 		result.SourceType = &tmp
// 	}

// 	return result
// }

// func ImageSourceDetailsToMap(obj *oci_core.ImageSourceDetails) map[string]interface{} {
// 	result := map[string]interface{}{}

// 	if obj.SourceType != nil {
// 		result["source_type"] = string(*obj.SourceType)
// 	}

// 	return result
// }
