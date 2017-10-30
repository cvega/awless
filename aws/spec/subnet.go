package awsspec

import (
	"net"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateSubnet struct {
	_                string `action:"create" entity:"subnet" awsAPI:"ec2" awsCall:"CreateSubnet" awsInput:"ec2.CreateSubnetInput" awsOutput:"ec2.CreateSubnetOutput" awsDryRun:""`
	logger           *logger.Logger
	api              ec2iface.EC2API
	CIDR             *string `awsName:"CidrBlock" awsType:"awsstr" templateName:"cidr" required:""`
	VPC              *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc" required:""`
	AvailabilityZone *string `awsName:"AvailabilityZone" awsType:"awsstr" templateName:"availabilityzone"`
	Name             *string `templateName:"name"`
}

func (cmd *CreateSubnet) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateSubnet) ValidateCIDR() error {
	_, _, err := net.ParseCIDR(StringValue(cmd.CIDR))
	return err
}

func (cmd *CreateSubnet) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateSubnetOutput).Subnet.SubnetId)
}

func (cmd *CreateSubnet) AfterRun(ctx map[string]interface{}, output interface{}) error {
	return createNameTag(awssdk.String(cmd.ExtractResult(output)), cmd.Name, ctx)
}

type UpdateSubnet struct {
	_      string `action:"update" entity:"subnet" awsAPI:"ec2" awsCall:"ModifySubnetAttribute" awsInput:"ec2.ModifySubnetAttributeInput" awsOutput:"ec2.ModifySubnetAttributeOutput"`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"SubnetId" awsType:"awsstr" templateName:"id" required:""`
	Public *bool   `awsName:"MapPublicIpOnLaunch" awsType:"awsboolattribute" templateName:"public"`
}

func (cmd *UpdateSubnet) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type DeleteSubnet struct {
	_      string `action:"delete" entity:"subnet" awsAPI:"ec2" awsCall:"DeleteSubnet" awsInput:"ec2.DeleteSubnetInput" awsOutput:"ec2.DeleteSubnetOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"SubnetId" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteSubnet) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}