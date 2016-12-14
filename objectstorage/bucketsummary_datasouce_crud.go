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

	opts := &baremetal.ListBucketsOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListBuckets{BucketSummaries: []baremetal.BucketSummary{}}

	for {
		var list *baremetal.ListBuckets

		if list, e = s.Client.ListBuckets(compartmentID, baremetal.Namespace(namespace), opts); e != nil {
			break
		}

		s.Res.BucketSummaries = append(s.Res.BucketSummaries, list.BucketSummaries...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}
	return
}

func (s * BucketsummaryDatasourceCrud) SetData() (e error) {
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
		s.D.Set("bucketsummaries", resources)

	}
	return
}
