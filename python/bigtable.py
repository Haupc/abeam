

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

TAG_1M = '1m'
TAG_3M = '3m'
TAG_5M = '5m'
TAG_15M = '15m'
TAG_30M = '30m'
TAG_45M = '45m'
TAG_1H = '1h'
TAG_2H = '2h'
TAG_3H = '3h'
TAG_4H = '4h'
TAG_1D = '1d'

TAGS = [TAG_1M, TAG_3M, TAG_5M, TAG_15M, TAG_30M,
        TAG_45M, TAG_1H, TAG_2H, TAG_3H, TAG_4H, TAG_1D]


def createTable(project_id, instance_id):
    client = bigtable.Client(project=project_id, admin=True)
    for tag in TAGS:
        instance = client.instance(instance_id)

        tableName = 'trading_view_' + tag
        table = instance.table(tableName)

        table.truncate()
        # column_families = {'data_point': column_family.MaxVersionsGCRule(1)}
        # if not table.exists():
        #     table.create(column_families=column_families)
        #     print("Table {} created.".format(tableName))
        # else:
        #     print("Table {} already exists.".format(tableName))


def main(project_id, instance_id, table_id):
    client = bigtable.Client(project=project_id, admin=True)
    instance = client.instance(instance_id)

    table = instance.table(table_id)

    # check if table exists
    column_family_id = 'testcf'
    column_families = {
        column_family_id: column_family.MaxVersionsGCRule(1)}
    if not table.exists():
        table.create(column_families=column_families)
    else:
        print("Table {} already exists.".format(table_id))

    rows = []
    timenow = datetime.datetime.utcfromtimestamp(1657217056)
    print(timenow.tzinfo)
    column = "c1"
    value = 1
    row = table.direct_row("test2")
    row.set_cell(
        column_family_id, column, value, timestamp=timenow
    )

    column = "c2"
    value = "c2_value"
    row.set_cell(
        column_family_id, column, value, timestamp=timenow
    )

    rows.append(row)
    statuslist = table.mutate_rows(rows)
    print(statuslist)


if __name__ == "__main__":
    # main("dev-krystal-wallet", "haupc-instance", "test")
    createTable("dev-krystal-wallet", "haupc-instance")
