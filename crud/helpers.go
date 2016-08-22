package crud

import (
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const FiveMinutes time.Duration = 5 * time.Minute

func CreateResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
		return
	}
	d.SetId(sync.Id())
	sync.SetData()

	stateful, ok := sync.(StatefullyCreatedResource)
	if ok {
		pending := stateful.CreatedPending()
		target := stateful.CreatedTarget()
		e = waitForStateRefresh(stateful, pending, target)
	}

	return
}

func ReadResource(sync ResourceReader) (e error) {
	if e = sync.Get(); e != nil {
		return
	}
	sync.SetData()

	return
}

func UpdateResource(d *schema.ResourceData, sync ResourceUpdater) (e error) {
	d.Partial(true)
	if e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)
	sync.SetData()

	return
}

func DeleteResource(sync ResourceDeleter) (e error) {
	if e = sync.Delete(); e != nil {
		return
	}

	stateful, ok := sync.(StatefullyDeletedResource)
	if ok {
		pending := stateful.DeletedPending()
		target := stateful.DeletedTarget()
		e = waitForStateRefresh(stateful, pending, target)
	}

	return
}

func stateRefreshFunc(sync StatefulResource) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if e = sync.Get(); e != nil {
			return nil, "", e
		}
		return sync, sync.State(), e
	}
}

func waitForStateRefresh(sync StatefulResource, pending, target []string) (e error) {
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFunc(sync),
		Timeout: FiveMinutes,
	}

	if _, e = stateConf.WaitForState(); e != nil {
		return
	}

	sync.SetData()

	return
}
