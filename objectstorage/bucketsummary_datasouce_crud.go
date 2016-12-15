package objectstorage

import (
	"time"
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
	"github.com/hashicorp/terraform/helper/schema"
)

type BucketsummaryDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListBuckets
}


func (s *BucketsummaryDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	namespace := s.D.Get("namespace").(string)
	page := s.D.Get("page").(string)

	opts := &baremetal.ListBucketsOptions{}
	opts.Page = page
	pageListOption := &baremetal.PageListOptions{}
	pageListOption.Page = page

	s.Res = &baremetal.ListBuckets{BucketSummaries: []baremetal.BucketSummary{}}

	for {
		var list *baremetal.ListBuckets
		if list, e = s.Client.ListBuckets(compartmentID, baremetal.Namespace(namespace), opts); e != nil {
			break
		}

		s.Res.BucketSummaries = append(s.Res.BucketSummaries, list.BucketSummaries...)

		// TODO: Add optionsSetNextListBucketPageOptions
		break
		pageListOption.Page = list.NextPage

		if hasNextPage := options.SetNextPageOption(list.NextPage, pageListOption); !hasNextPage {
			break
		}
	}
	return
}

func (s * BucketsummaryDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, r := range s.Res.BucketSummaries {
			res := map[string]interface{}{
				"namespace":       r.Namespace,
				"name":            r.Name,
				"compartment_id":  r.CompartmentID,
				"created_by":      r.CreatedBy,
				"time_created":    r.TimeCreated.String(),
				"etag":            r.ETag,
			}
			resources = append(resources, res)
		}
		s.D.Set("bucketsummary", resources)
	}
	return
}
