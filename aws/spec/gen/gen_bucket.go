/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
// This file was automatically generated with go generate
package awsspec

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/wallix/awless/logger"
)

type CreateBucket struct {
	_      string `action:"create" entity:"bucket" awsAPI:"s3" awsCall:"CreateBucket" awsInput:"s3.CreateBucketInput" awsOutput:"s3.CreateBucketOutput"`
	logger *logger.Logger
	api    s3iface.S3API
	Name   *string `awsName:"Bucket" awsType:"awsstr" templateName:"name" required:""`
	Acl    *string `awsName:"ACL" awsType:"awsstr" templateName:"acl"`
}

func (cmd *CreateBucket) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateBucket) ExtractResult(i interface{}) string {
	return params["name"]
}

type UpdateBucket struct {
	_                string `action:"update" entity:"bucket" awsAPI:"s3"`
	logger           *logger.Logger
	api              s3iface.S3API
	Name             *struct{} `templateName:"name" required:""`
	Acl              *struct{} `templateName:"acl"`
	PublicWebsite    *struct{} `templateName:"public-website"`
	RedirectHostname *struct{} `templateName:"redirect-hostname"`
	IndexSuffix      *struct{} `templateName:"index-suffix"`
	EnforceHttps     *struct{} `templateName:"enforce-https"`
}

type DeleteBucket struct {
	_      string `action:"delete" entity:"bucket" awsAPI:"s3" awsCall:"DeleteBucket" awsInput:"s3.DeleteBucketInput" awsOutput:"s3.DeleteBucketOutput"`
	logger *logger.Logger
	api    s3iface.S3API
	Name   *string `awsName:"Bucket" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *DeleteBucket) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
