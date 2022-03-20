// Copyright (c) 2022 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint, accessKey, secretKey string
	bucket, prefix                 string
	debug, fake                    bool
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "https://play.min.io", "S3 endpoint URL")
	flag.StringVar(&accessKey, "access-key", "Q3AM3UQ867SPQQA43P2F", "S3 Access Key")
	flag.StringVar(&secretKey, "secret-key", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", "S3 Secret Key")
	flag.StringVar(&bucket, "bucket", "", "Select a specific bucket")
	flag.StringVar(&prefix, "prefix", "", "Select a prefix")
	flag.BoolVar(&debug, "debug", false, "Prints HTTP network calls to S3 endpoint")
	flag.BoolVar(&fake, "fake", false, "Do a dry run")

	flag.Parse()

	if endpoint == "" {
		log.Fatalln("Endpoint is not provided")
	}

	if accessKey == "" {
		log.Fatalln("Access key is not provided")
	}

	if secretKey == "" {
		log.Fatalln("Secret key is not provided")
	}

	if bucket == "" && prefix != "" {
		log.Fatalln("--prefix is specified without --bucket.")
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalln(err)
	}

	s3Client, err := minio.New(u.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: strings.EqualFold(u.Scheme, "https"),
	})
	if err != nil {
		log.Fatalln()
	}

	if debug {
		s3Client.TraceOn(os.Stderr)
	}

	opts := minio.ListObjectsOptions{
		WithVersions: true,
		Recursive:    true,
		Prefix:       prefix,
	}

	objectsCh := make(chan minio.ObjectInfo)
	sendForDeleteFn := func(ctx context.Context, object minio.ObjectInfo) {
		if fake {
			fmt.Println("DRY RUN:", bucket+"/"+object.Key+" "+object.VersionID, " will be deleted")
			return
		}
		fmt.Println("LOG:", bucket+"/"+object.Key+" "+object.VersionID, "deleted.")
		sent := false
		for !sent {
			select {
			case objectsCh <- object:
				sent = true
			case <-ctx.Done():
				return
			}
		}
	}
	ctx := context.Background()
	//var prev, older minio.ObjectInfo
	// Send object names that are needed to be removed to objectsCh
	markDelete := 0
	currPrefix := ""
	listCnt := 0
	var markDel minio.ObjectInfo
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		for object := range s3Client.ListObjects(ctx, bucket, opts) {
			if object.Err != nil {
				log.Fatalln("LIST error:", object.Err)
				continue
			}
			if strings.HasSuffix(object.Key, "/") {
				if object.IsLatest {
					markDel = object
					continue
				}
				if object.Key == markDel.Key {
					sendForDeleteFn(ctx, object)
					markDelete++
				}
			}
			listCnt++
			if strings.HasSuffix(object.Key, "/") {
				currPrefix = object.Key
			}
			if listCnt%1000 == 0 {
				log.Println("Listing at prefix: "+currPrefix+" Total listed:", listCnt, " # older versions of empty directories deleted", markDelete)
			}
		}

	}()
	errorCh := s3Client.RemoveObjects(context.Background(), bucket, objectsCh, minio.RemoveObjectsOptions{})
	for e := range errorCh {
		fmt.Println("Failed to remove " + e.ObjectName + e.VersionID + ", error: " + e.Err.Error())
	}
	log.Println(bucket+"/"+prefix, " Total listed :", listCnt, " # older versions of empty directories deleted=", markDelete)

}
