# https://orahub.oci.oraclecorp.com/ateam/jobs/-/raw/master/python_jobs/job_std_logs.py

import oci
import datetime
import os
import time
import sys
from oci.loggingingestion import LoggingClient

OCI_RESOURCE_PRINCIPAL_VERSION = "OCI_RESOURCE_PRINCIPAL_VERSION"
JOB_RUN_OCID_KEY = "JOB_RUN_OCID"
LOG_OBJECT_OCID_KEY = "LOG_OCID"

class Job:
    def __init__(self):
        # Auto switch between local and job
        rp_version = os.environ.get(OCI_RESOURCE_PRINCIPAL_VERSION, "UNDEFINED")
        if rp_version == "UNDEFINED":
            # RUN LOCAL TEST
            self.signer = oci.config.from_file("~/.oci/config", "DEFAULT")
            self.log_client = LoggingClient(config=self.signer)
        else:
            # RUN AS JOB
            self.signer = oci.auth.signers.get_resource_principals_signer()
            self.log_client = LoggingClient(config={}, signer=self.signer)
try:
    job = Job()

    try:
        print(
            "Start logging for job run: {}".format(
                os.environ.get(JOB_RUN_OCID_KEY, "UNDEFINED")
            )
        )
        print("Current timestamp in UTC: {}".format(str(datetime.datetime.utcnow())))

        print("Delay 5s")

        time.sleep(5)

        print("Log after delay...")

        print("Job Done.")
    except Exception as e:
        print(e)
        raise e
except Exception as e:
    print(e)
    raise e
