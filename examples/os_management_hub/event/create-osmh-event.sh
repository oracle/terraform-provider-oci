#!/bin/bash

# set -x -e

# . ~/venv-amd64/bin/activate

response="$(oci raw-request --http-method POST --target-uri "https://osmh.us-ashburn-1.oci.oraclecloud.com/20220901/events" --request-body '
{
    "eventSummary": "Manually created event 2 for testing caused by <Yijiu>",
    "type": "EXPLOIT_ATTEMPT",
    "data": {
        "content": {
            "type": "EXPLOIT_ATTEMPT",
            "exploitDetectionLogContent": "exploitDetectionLogContent",
            "contentLocation": "/var/lib/oracle-cloud-agent/plugins/oci-alx/oops/reports/2488906d741741b3349615959ddfa7b3cd047d50.2020-07-14T01:38:28.000Z"
        },
        "count": 2,
        "additionalDetails": {
            "exploitCves": [
                "CVE-1234",
                "CVE-2222"
            ]
        }
    },
    "eventDetails": "Autonomous instance has an exploit",
    "timeOccurred": "2024-01-25T23:00:49.382Z",
    "timeCreated": "2023-12-05T03:31:09.844Z",
    "timeUpdated": "2023-12-04T22:52:50.411Z",
    "compartmentId": "ocid1.tenancy.oc1..aaaaaaaatajukaw53bfyxan7qlhje4u6v5hvfa3la7jbh7cok3qukw7rirbq",
    "resourceId": "ocid1.instance.oc1.iad.anuwcljt22fr4bicxlufr5l3xmrts2ttuvn2cmx3ozrhopindb25jgfwyg3q",
    "systemDetails": {
        "architecture": "X86_64",
        "kspliceEffectiveKernelVersion": "5.4.17-2102.201.3.el7uek.x86_64",
        "osFamily": "ORACLE_LINUX_8",
        "osName": "ORACLE_LINUX_SERVER",
        "osKernelRelease": "#2 SMP Fri Apr 23 09:05:55 PDT 2021",
        "osKernelVersion": "5.4.17-2102.201.3.el7uek.x86_64",
        "osSystemVersion": "8.2"
    },
    "isManagedByAutonomousLinux": true,
    "freeformTags": {
        "Department": "Finance"
    }
}
')"

id="$(echo "$response" | jq -r '.data.id')"

cat <<EOF
{
  "id": "$id"
}
EOF
