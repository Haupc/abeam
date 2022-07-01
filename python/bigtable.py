

"""Demonstrates how to connect to Cloud Bigtable and run some basic operations.

Prerequisites:

- Create a Cloud Bigtable cluster.
  https://cloud.google.com/bigtable/docs/creating-cluster
- Set your Google Application Default Credentials.
  https://developers.google.com/identity/protocols/application-default-credentials
"""

import argparse

import datetime

from google.cloud import bigtable
from google.cloud.bigtable import column_family
from google.cloud.bigtable import row_filters


def main(project_id, instance_id, table_id):
    client = bigtable.Client(project=project_id, admin=True)
    instance = client.instance(instance_id)

    print("Creating the {} table.".format(table_id))
    table = instance.table(table_id)
    column_family_id = 'balance'
    rows = []
    column = "amount"
    value = str(10**80+1)

    row = table.direct_row("test")
    row.set_cell(
        column_family_id, column, value, timestamp=datetime.datetime.utcnow()
    )
    rows.append(row)
    statuslist = table.mutate_rows(rows)
    print(statuslist)


if __name__ == "__main__":
    main("dev-krystal-wallet", "haupc-instance", "balance")
