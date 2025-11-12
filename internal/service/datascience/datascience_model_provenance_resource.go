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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatascienceModelProvenanceWithContext,
		ReadContext:   readDatascienceModelProvenanceWithContext,
		UpdateContext: updateDatascienceModelProvenanceWithContext,
		DeleteContext: deleteDatascienceModelProvenanceWithContext,
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

func createDatascienceModelProvenanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatascienceModelProvenanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatascienceModelProvenanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceModelProvenanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatascienceModelProvenanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func (s *DatascienceModelProvenanceResourceCrud) CreateWithContext(ctx context.Context) error {
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

	response, err := s.Client.CreateModelProvenance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelProvenance
	return nil
}

func (s *DatascienceModelProvenanceResourceCrud) GetWithContext(ctx context.Context) error {
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

	response, err := s.Client.GetModelProvenance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelProvenance
	return nil
}

func (s *DatascienceModelProvenanceResourceCrud) UpdateWithContext(ctx context.Context) error {
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

	response, err := s.Client.UpdateModelProvenance(ctx, request)
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
