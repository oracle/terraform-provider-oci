# Private DNS RRSet examples

This directory creates a private view and zones, then demonstrates RRSets in
that view. Each RRSet uses `view_id` to select the private view.

## Setup

Use a test compartment and an OCI identity with permission to manage the example
resources. The provider uses API key authentication through these inputs:
`tenancy_ocid`, `user_ocid`, `fingerprint` and `private_key_path`. Also set
`compartment_ocid` and `region`. Supply these with `TF_VAR_` environment variables
or a local, untracked variables file. Keep credentials out of source control.
For other authentication methods, see the [provider authentication guide](https://registry.terraform.io/providers/oracle/oci/latest/docs#authentication).

Run from this directory:

```sh
terraform init
terraform plan
terraform apply
terraform plan
terraform destroy
```

## Single and multiple records

`rrset.tf` creates `single.<zone>` with one A record and `multi.<zone>` with two.
The top-level zone, domain and type identify the RRSet. Each `items` block repeats
the domain and type and supplies one record value and TTL. Use one resource for
all values in the same RRSet. The A records use documentation addresses.

Both A examples use `rrset_ttl`, which defaults to 300 seconds. Every record
within an RRSet uses the same TTL. The AAAA and CNAME examples specify their TTLs
directly.

Remove an `items` block to remove that record, then review the plan and apply.
An unchanged configuration should produce an empty plan. Import an existing
RRSet before managing it: creating over an RRSet that already contains records
fails with a conflict. Omitting `items` does not delete those existing records.

## Other record types and lookups

`rrset.tf` also demonstrates an AAAA record for an IPv6 address and a CNAME
record that points `sub.<zone>` to the zone apex. Its data sources read the
CNAME RRSet without managing its records. The `zones` output lists zone names
and IDs.

These examples cover a few common record types. See the [OCI supported record
types](https://docs.oracle.com/en-us/iaas/Content/DNS/Tasks/privatedns.htm) for the full list and record-data formats.

See also the [public RRSet example](../../global/rrset.tf).
