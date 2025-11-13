// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

func DatascienceModelProvenanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModelProvenance,
		Read:     readDatascienceModelProvenance,
		Update:   updateDatascienceModelProvenance,
		Delete:   deleteDatascienceModelProvenance,
		Schema: map[string]*schema.Schema{
			// Required
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"git_branch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"git_commit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"repository_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_dir": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"training_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"training_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createDatascienceModelProvenance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelProvenance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceModelProvenance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceModelProvenance(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatascienceModelProvenanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.ModelProvenance
	DisableNotFoundRetries bool
}

func (s *DatascienceModelProvenanceResourceCrud) ID() string {
	return GetModelProvenanceCompositeId(s.D.Get("model_id").(string))
}

func (s *DatascienceModelProvenanceResourceCrud) Create() error {
	request := oci_datascience.CreateModelProvenanceRequest{}

	if gitBranch, ok := s.D.GetOkExists("git_branch"); ok {
		tmp := gitBranch.(string)
		request.GitBranch = &tmp
	}

	if gitCommit, ok := s.D.GetOkExists("git_commit"); ok {
		tmp := gitCommit.(string)
		request.GitCommit = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if repositoryUrl, ok := s.D.GetOkExists("repository_url"); ok {
		tmp := repositoryUrl.(string)
		request.RepositoryUrl = &tmp
	}

	if scriptDir, ok := s.D.GetOkExists("script_dir"); ok {
		tmp := scriptDir.(string)
		request.ScriptDir = &tmp
	}

	if trainingId, ok := s.D.GetOkExists("training_id"); ok {
		tmp := trainingId.(string)
		request.TrainingId = &tmp
	}

	if trainingScript, ok := s.D.GetOkExists("training_script"); ok {
		tmp := trainingScript.(string)
		request.TrainingScript = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateModelProvenance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelProvenance
	return nil
}

func (s *DatascienceModelProvenanceResourceCrud) Get() error {
	request := oci_datascience.GetModelProvenanceRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	modelId, err := parseModelProvenanceCompositeId(s.D.Id())
	if err == nil {
		request.ModelId = &modelId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetModelProvenance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelProvenance
	return nil
}

func (s *DatascienceModelProvenanceResourceCrud) Update() error {
	request := oci_datascience.UpdateModelProvenanceRequest{}

	if gitBranch, ok := s.D.GetOkExists("git_branch"); ok {
		tmp := gitBranch.(string)
		request.GitBranch = &tmp
	}

	if gitCommit, ok := s.D.GetOkExists("git_commit"); ok {
		tmp := gitCommit.(string)
		request.GitCommit = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if repositoryUrl, ok := s.D.GetOkExists("repository_url"); ok {
		tmp := repositoryUrl.(string)
		request.RepositoryUrl = &tmp
	}

	if scriptDir, ok := s.D.GetOkExists("script_dir"); ok {
		tmp := scriptDir.(string)
		request.ScriptDir = &tmp
	}

	if trainingId, ok := s.D.GetOkExists("training_id"); ok {
		tmp := trainingId.(string)
		request.TrainingId = &tmp
	}

	if trainingScript, ok := s.D.GetOkExists("training_script"); ok {
		tmp := trainingScript.(string)
		request.TrainingScript = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateModelProvenance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelProvenance
	return nil
}

func (s *DatascienceModelProvenanceResourceCrud) SetData() error {

	modelId, err := parseModelProvenanceCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("model_id", &modelId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.GitBranch != nil {
		s.D.Set("git_branch", *s.Res.GitBranch)
	}

	if s.Res.GitCommit != nil {
		s.D.Set("git_commit", *s.Res.GitCommit)
	}

	if s.Res.RepositoryUrl != nil {
		s.D.Set("repository_url", *s.Res.RepositoryUrl)
	}

	if s.Res.ScriptDir != nil {
		s.D.Set("script_dir", *s.Res.ScriptDir)
	}

	if s.Res.TrainingId != nil {
		s.D.Set("training_id", *s.Res.TrainingId)
	}

	if s.Res.TrainingScript != nil {
		s.D.Set("training_script", *s.Res.TrainingScript)
	}

	return nil
}

func GetModelProvenanceCompositeId(modelId string) string {
	modelId = url.PathEscape(modelId)
	compositeId := "models/" + modelId + "/provenance"
	return compositeId
}

func parseModelProvenanceCompositeId(compositeId string) (modelId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("models/.*/provenance", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	modelId, _ = url.PathUnescape(parts[1])

	return
}
