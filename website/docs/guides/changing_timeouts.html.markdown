---
layout: "oci"
page_title: "Changing Timeouts for the Service"
sidebar_current: "docs-oci-guide-changing_timeouts"
description: |-
  The Oracle Cloud Infrastructure provider. Changing Timeouts for the Service
---

## Timeout errors when waiting for a state change

_If the Terraform CLI gives an error message like:_

```
* oci_database_backup.mydb: timeout while waiting for state to become 'ACTIVE' (last state: 'CREATING', timeout: 15m0s)
```

Then the OCI service is indicating that the resource has not yet reached the expected state after polling for some time.

You may need to increase the operation timeout for your resource to continue polling for longer. See [Operation Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for details on how to do this.
