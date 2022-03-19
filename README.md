### Delete excess delete markers for versioned directory objects

This tool handles only excess delete markers for dir object entries

#### Remove duplicate delete markers on a prefix "prefix/to/em7" and delete extra delete markers

 ./dmdirdelete --bucket bucket --prefix prefix/to/em7 --endpoint http://minio:9000 --access-key minio --secret-key minio123 > /tmp/log

#### Perform a dry run to delete duplicate delete markers on a prefix "prefix/to/em7"
 ./dmdirdelete --bucket bucket --prefix prefix/to/em7 --endpoint http://minio:9000 --access-key minio --secret-key minio123 --fake

Example:
```
mc ls minio/bucket/prefix/to -r --versions
[2022-03-19 16:36:42 PDT]     0B STANDARD f926ea2e-4941-4028-9e37-faa505d6ccb9 v1 PUT em1/
[2022-03-19 14:33:01 PDT]     0B STANDARD 6385ec31-5c39-4ac0-93ee-7f0f63015b14 v2 DEL em3/
[2022-03-19 13:35:23 PDT]     0B STANDARD 8defc21d-fabb-473a-9bfa-05d6f6d7212d v1 PUT em3/
[2022-03-19 15:57:46 PDT]     0B STANDARD 9fa1103e-1868-4c6a-9eb9-df1e4c74a633 v1 PUT em5/
[2022-03-19 15:58:28 PDT]     0B STANDARD 7b081827-bb79-463a-9615-df4580594523 v2 DEL em6/
[2022-03-19 15:58:01 PDT]     0B STANDARD a4ed811f-96c2-4efb-9d51-33b5e178fd9b v1 PUT em6/
[2022-03-19 17:13:39 PDT]     0B STANDARD 9d86d5fd-9c79-46a3-95f4-1fd4c2275d66 v4 DEL em7/
[2022-03-19 17:13:38 PDT]     0B STANDARD 680bbe57-32f6-4b56-a27f-38bad66362a8 v3 DEL em7/
[2022-03-19 15:58:33 PDT]     0B STANDARD 85113983-b581-4fb3-a371-c7025c0e93b0 v2 DEL em7/
[2022-03-19 15:58:09 PDT]     0B STANDARD d99a91d1-b6ca-4041-a7bf-4c5452822f1b v1 PUT em7/
```

After running this tool on prefix "prefix/to/em7", you should see:
```
mc ls minio/bucket/prefix/to -r --versions
➜  sidekick mc ls sitea/bucket/prefix/to -r --versions  
[2022-03-19 16:36:42 PDT]     0B STANDARD f926ea2e-4941-4028-9e37-faa505d6ccb9 v1 PUT em1/
[2022-03-19 14:33:01 PDT]     0B STANDARD 6385ec31-5c39-4ac0-93ee-7f0f63015b14 v2 DEL em3/
[2022-03-19 13:35:23 PDT]     0B STANDARD 8defc21d-fabb-473a-9bfa-05d6f6d7212d v1 PUT em3/
[2022-03-19 15:57:46 PDT]     0B STANDARD 9fa1103e-1868-4c6a-9eb9-df1e4c74a633 v1 PUT em5/
[2022-03-19 15:58:28 PDT]     0B STANDARD 7b081827-bb79-463a-9615-df4580594523 v2 DEL em6/
[2022-03-19 15:58:01 PDT]     0B STANDARD a4ed811f-96c2-4efb-9d51-33b5e178fd9b v1 PUT em6/
[2022-03-19 15:58:33 PDT]     0B STANDARD 85113983-b581-4fb3-a371-c7025c0e93b0 v2 DEL em7/
[2022-03-19 15:58:09 PDT]     0B STANDARD d99a91d1-b6ca-4041-a7bf-4c5452822f1b v1 PUT em7/
```
