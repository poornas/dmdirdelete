### Delete noncurrent versions for versioned directory objects

This tool delete all non current versions of dir objects.
The top most delete marker is expected to be cleaned up by lifecycle.

#### Remove noncurrent versions of directory object "prefix/to/dirname"

 ./delnoncurrent --bucket bucket --object prefix/to/dirname --endpoint http://minio:9000 --access-key minio --secret-key minio123 > /tmp/log

#### Perform a dry run to delete noncurrent versions of directory object "prefix/to/dirname"
 ./delnoncurrent --bucket bucket --object prefix/to/dirname --endpoint http://minio:9000 --access-key minio --secret-key minio123 --fake

Example:
```
mc ls minio/bucket/prefix/to/em7 -r --versions
[2022-03-19 17:13:39 PDT]     0B STANDARD 9d86d5fd-9c79-46a3-95f4-1fd4c2275d66 v4 DEL em7/
[2022-03-19 17:13:38 PDT]     0B STANDARD 680bbe57-32f6-4b56-a27f-38bad66362a8 v3 DEL em7/
[2022-03-19 15:58:33 PDT]     0B STANDARD 85113983-b581-4fb3-a371-c7025c0e93b0 v2 DEL em7/
[2022-03-19 15:58:09 PDT]     0B STANDARD d99a91d1-b6ca-4041-a7bf-4c5452822f1b v1 PUT em7/
```

After running this tool on object "prefix/to/dirname", you should see:
```
➜  mc ls sitea/bucket/prefix/to/em7 -r --versions  
[2022-03-19 15:58:33 PDT]     0B STANDARD 85113983-b581-4fb3-a371-c7025c0e93b0 v2 DEL em7/
```
