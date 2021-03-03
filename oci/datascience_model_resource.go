// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/oracle/oci-go-sdk/v36/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_datascience "github.com/oracle/oci-go-sdk/v36/datascience"
)

func init() {
	RegisterResource("oci_datascience_model", DatascienceModelResource())
}

func DatascienceModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatascienceModel,
		Read:     readDatascienceModel,
		Update:   updateDatascienceModel,
		Delete:   deleteDatascienceModel,
		Schema: map[string]*schema.Schema{
			// Required
			"artifact_content_length": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     validateInt64TypeString,
				DiffSuppressFunc: int64StringDiffSuppressFunction,
			},
			"model_artifact": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"artifact_content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_datascience.ModelLifecycleStateActive),
					string(oci_datascience.ModelLifecycleStateInactive),
				}, true),
			},

			// Computed
			"artifact_content_md5": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"artifact_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"empty_model": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	var deactivateModel = false
	if state, ok := sync.D.GetOkExists("state"); ok {
		desiredState := oci_datascience.ModelLifecycleStateEnum(strings.ToUpper(state.(string)))
		if desiredState == oci_datascience.ModelLifecycleStateInactive {
			deactivateModel = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}
	if deactivateModel {
		if e := sync.DeactivateModel(); e != nil {
			return e
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateInactive)
	}
	if e := sync.CreateArtifact(); e != nil {
		return e
	}
	sync.D.Set("empty_model", false)
	return ReadResource(sync)
}

func readDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

func updateDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	// Activate/Deactivate Model
	activate, deactivate := false, false

	if sync.D.HasChange("state") {
		desiredState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.ModelLifecycleStateActive == oci_datascience.ModelLifecycleStateEnum(desiredState) {
			activate = true
		} else if oci_datascience.ModelLifecycleStateInactive == oci_datascience.ModelLifecycleStateEnum(desiredState) {
			deactivate = true
		}
	} else {
		currentState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.ModelLifecycleStateActive == oci_datascience.ModelLifecycleStateEnum(currentState) {
			activate = true
			deactivate = true
		}
	}

	if deactivate {
		if err := sync.DeactivateModel(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateInactive)
	}
	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	if activate {
		if err := sync.ActivateModel(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateActive)
	}
	return nil
}

func deleteDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type HeadModelArtifact struct {
	ContentLength      *int64
	ContentDisposition *string
	ContentMd5         *string
	LastModified       *common.SDKTime
}

type DatascienceModelResourceCrud struct {
	BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.Model
	ArtifactHeadRes        *HeadModelArtifact
	DisableNotFoundRetries bool
}

func (s *DatascienceModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceModelResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatascienceModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ModelLifecycleStateActive),
	}
}

func (s *DatascienceModelResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatascienceModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ModelLifecycleStateDeleted),
	}
}

func (s *DatascienceModelResourceCrud) Create() error {
	request := oci_datascience.CreateModelRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *DatascienceModelResourceCrud) Get() error {
	request := oci_datascience.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	if emptyModel, ok := s.D.GetOkExists("empty_model"); ok {
		tmp := emptyModel.(bool)
		if !tmp {
			err := s.GetArtifactHead()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DatascienceModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateModelRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *DatascienceModelResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteModel(context.Background(), request)
	return err
}

func (s *DatascienceModelResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return s.SetArtifactData()
}

func (s *DatascienceModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceModelResourceCrud) ActivateModel() error {
	request := oci_datascience.ActivateModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ActivateModel(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelResourceCrud) DeactivateModel() error {
	request := oci_datascience.DeactivateModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeactivateModel(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelLifecycleStateInactive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelResourceCrud) CreateArtifact() error {
	request := oci_datascience.CreateModelArtifactRequest{}

	if contentDisposition, ok := s.D.GetOkExists("artifact_content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("artifact_content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert Content-Length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if modelArtifact, ok := s.D.GetOkExists("model_artifact"); ok {
		tmp := modelArtifact.(string)
		var artifactReader io.Reader
		artifactReader, err := os.Open(tmp)
		if err != nil {
			return fmt.Errorf("the specified model_artifact is not available: %q", err)
		}
		request.ModelArtifact = ioutil.NopCloser(artifactReader)
	}

	request.ModelId = s.Res.Id

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateModelArtifact(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceModelResourceCrud) GetArtifactHead() error {
	request := oci_datascience.HeadModelArtifactRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")

	response, err := s.Client.HeadModelArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.ArtifactHeadRes = &HeadModelArtifact{
		ContentLength:      response.ContentLength,
		ContentDisposition: response.ContentDisposition,
		ContentMd5:         response.ContentMd5,
		LastModified:       response.LastModified,
	}
	return nil
}

func (s *DatascienceModelResourceCrud) SetArtifactData() error {
	if s.ArtifactHeadRes != nil {
		if s.ArtifactHeadRes.ContentDisposition != nil {
			s.D.Set("artifact_content_disposition", *s.ArtifactHeadRes.ContentDisposition)
		}

		if s.ArtifactHeadRes.ContentLength != nil {
			s.D.Set("artifact_content_length", *s.ArtifactHeadRes.ContentLength)
		}

		if s.ArtifactHeadRes.ContentMd5 != nil {
			s.D.Set("artifact_content_md5", *s.ArtifactHeadRes.ContentMd5)
		}

		if s.ArtifactHeadRes.LastModified != nil {
			s.D.Set("artifact_last_modified", s.ArtifactHeadRes.LastModified.String())
		}

		s.D.Set("empty_model", false)
	}

	return nil
}
