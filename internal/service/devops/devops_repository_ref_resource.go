// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
)

func DevopsRepositoryRefResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsRepositoryRef,
		Read:     readDevopsRepositoryRef,
		Update:   updateDevopsRepositoryRef,
		Delete:   deleteDevopsRepositoryRef,
		Schema: map[string]*schema.Schema{
			// Required
			"ref_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ref_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BRANCH",
					"TAG",
				}, true),
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"commit_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"full_ref_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDevopsRepositoryRef(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsRepositoryRef(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsRepositoryRef(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsRepositoryRef(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsRepositoryRefResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.RepositoryRef
	DisableNotFoundRetries bool
}

func (s *DevopsRepositoryRefResourceCrud) ID() string {
	repositoryRef := *s.Res
	return GetRepositoryRefCompositeId(*repositoryRef.GetRefName(), *repositoryRef.GetRepositoryId()) //TODO: is this unique
}

func (s *DevopsRepositoryRefResourceCrud) Create() error {
	request := oci_devops.PutRepositoryRefRequest{}
	err := s.populateTopLevelPolymorphicPutRepositoryRefRequest(&request)
	if err != nil {
		return err
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.PutRepositoryRef(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositoryRef
	return nil
}

func (s *DevopsRepositoryRefResourceCrud) Get() error {
	request := oci_devops.GetRefRequest{}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	refName, repositoryId, err := parseRepositoryRefCompositeId(s.D.Id())
	if err == nil {
		request.RefName = &refName
		request.RepositoryId = &repositoryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetRef(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositoryRef
	return nil
}

func (s *DevopsRepositoryRefResourceCrud) Update() error {
	request := oci_devops.PutRepositoryRefRequest{}
	err := s.populateTopLevelPolymorphicPutRepositoryRefRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.PutRepositoryRef(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositoryRef
	return nil
}

func (s *DevopsRepositoryRefResourceCrud) Delete() error {
	request := oci_devops.DeleteRefRequest{}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	_, err := s.Client.DeleteRef(context.Background(), request)
	return err
}

func (s *DevopsRepositoryRefResourceCrud) SetData() error {

	refName, repositoryId, err := parseRepositoryRefCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("ref_name", &refName)
		s.D.Set("repository_id", &repositoryId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_devops.RepositoryBranch:
		s.D.Set("ref_type", "BRANCH")

		if v.CommitId != nil {
			s.D.Set("commit_id", *v.CommitId)
		}

		if v.FullRefName != nil {
			s.D.Set("full_ref_name", *v.FullRefName)
		}

		if v.RefName != nil {
			s.D.Set("ref_name", *v.RefName)
		}

		if v.RepositoryId != nil {
			s.D.Set("repository_id", *v.RepositoryId)
		}
	case oci_devops.RepositoryTag:
		s.D.Set("ref_type", "TAG")

		if v.ObjectId != nil {
			s.D.Set("object_id", *v.ObjectId)
		}

		if v.FullRefName != nil {
			s.D.Set("full_ref_name", *v.FullRefName)
		}

		if v.RefName != nil {
			s.D.Set("ref_name", *v.RefName)
		}

		if v.RepositoryId != nil {
			s.D.Set("repository_id", *v.RepositoryId)
		}
	default:
		log.Printf("[WARN] Received 'ref_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetRepositoryRefCompositeId(refName string, repositoryId string) string {
	refName = url.PathEscape(refName)
	repositoryId = url.PathEscape(repositoryId)
	compositeId := "repositories/" + repositoryId + "/refs/" + refName
	return compositeId
}

func parseRepositoryRefCompositeId(compositeId string) (refName string, repositoryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("repositories/.*/refs/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	repositoryId, _ = url.PathUnescape(parts[1])
	refName, _ = url.PathUnescape(parts[3])

	return
}

func RepositoryRefSummaryToMap(obj oci_devops.RepositoryRefSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetRepositoryId() != nil {
		result["repository_id"] = string(*obj.GetRepositoryId())
	}

	if obj.GetRefName() != nil {
		result["ref_name"] = string(*obj.GetRefName())
	}

	if obj.GetFullRefName() != nil {
		result["full_ref_name"] = string(*obj.GetFullRefName())
	}

	switch v := (obj).(type) {
	case oci_devops.RepositoryBranchSummary:
		result["ref_type"] = "BRANCH"

		if v.CommitId != nil {
			result["commit_id"] = string(*v.CommitId)
		}
	case oci_devops.RepositoryTagSummary:
		result["ref_type"] = "TAG"

		if v.ObjectId != nil {
			result["object_id"] = string(*v.ObjectId)
		}
	default:
		log.Printf("[WARN] Received 'ref_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsRepositoryRefResourceCrud) populateTopLevelPolymorphicPutRepositoryRefRequest(request *oci_devops.PutRepositoryRefRequest) error {
	//discriminator
	refTypeRaw, ok := s.D.GetOkExists("ref_type")
	var refType string
	if ok {
		refType = refTypeRaw.(string)
	} else {
		refType = "" // default value
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if RefName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := RefName.(string)
		request.RefName = &tmp
	}

	switch strings.ToLower(refType) {
	case strings.ToLower("BRANCH"):
		details := oci_devops.PutRepositoryBranchDetails{}
		if commitId, ok := s.D.GetOkExists("commit_id"); ok {
			tmp := commitId.(string)
			details.CommitId = &tmp
		}
		request.PutRepositoryRefDetails = details
	case strings.ToLower("TAG"):
		details := oci_devops.PutRepositoryTagDetails{}
		if objectId, ok := s.D.GetOkExists("object_id"); ok {
			tmp := objectId.(string)
			details.ObjectId = &tmp
		}
		request.PutRepositoryRefDetails = details
	default:
		return fmt.Errorf("unknown ref_type '%v' was specified", refType)
	}
	return nil
}
