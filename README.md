s3v
===

Tool for S3 Versioning object


Build
-----

```
$ go get
$ go build
```

Usage
-----

```
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ ./s3v --bucket="<bucket>" --region="<region>" <subcommand> [args..]
```

Sub commands
------------

### ls

```
$ ./s3v --bucket="bucket1" --region="ap-northeast-1" ls
2015-07-14 03:36:32 +0000 UTC           development/testobject.json
```

バケットに保存されているファイルの一覧を表示します。

### log

```
$ ./s3v --bucket="bucket1" --region="ap-northeast-1" log development/testobject.json
[development/testobject.json]
2015-07-10 06:43:26 +0000 UTC           i04Sxt.tH06AgnuJAeBA35W5HU1jitlJ [LATEST]
2015-07-10 06:43:06 +0000 UTC           rv4_iJoO_.iXy6oM18CDaE4b7zpLBIwC
```

指定したオブジェクトのバージョンリストを表示します。

### diff

```
$ ./s3v --bucket="bucket1" --region="ap-northeast-1" diff i04Sxt.tH06AgnuJAeBA35W5HU1jitlJ rv4_iJoO_.iXy6oM18CDaE4b7zpLBIwC development/testobject.json
--- i04Sxt.tH06AgnuJAeBA35W5HU1jitlJ:development/testobject.json 2015-07-10 06:43:06 +0000 UTC
+++ rv4_iJoO_.iXy6oM18CDaE4b7zpLBIwC:development/testobject.json 2015-07-10 06:43:26 +0000 UTC
   {
       "version": 1,
-      "serial": 1,
+      "serial": 0,
       "modules": [
       ]
   }
```


