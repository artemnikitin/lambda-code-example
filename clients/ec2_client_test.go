package clients

import (
	"errors"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// EC2 mock
type ec2Mock struct {
	ec2iface.EC2API
	err                     error
	describeInstancesOutput *ec2.DescribeInstancesOutput
	terminateInstanceOutput *ec2.TerminateInstancesOutput
}

func (v *ec2Mock) DescribeInstancesRequest(*ec2.DescribeInstancesInput) ec2.DescribeInstancesRequest {
	return ec2.DescribeInstancesRequest{
		Request: &aws.Request{
			Data:  v.describeInstancesOutput,
			Error: v.err,
		},
	}
}

func (v *ec2Mock) TerminateInstancesRequest(*ec2.TerminateInstancesInput) ec2.TerminateInstancesRequest {
	return ec2.TerminateInstancesRequest{
		Request: &aws.Request{
			Data:  v.terminateInstanceOutput,
			Error: v.err,
		},
	}
}

// tests
func TestEC2Client_GetInstancesForCI(t *testing.T) {
	cases := []struct {
		name   string
		client *ec2Mock
		list   []string
		error  bool
	}{
		{
			name: "Positive case without name",
			client: &ec2Mock{
				describeInstancesOutput: &ec2.DescribeInstancesOutput{
					Reservations: []ec2.RunInstancesOutput{
						{
							Instances: []ec2.Instance{
								{
									InstanceId: aws.String(""),
									State:      &ec2.InstanceState{Code: aws.Int64(16)},
								},
							},
						},
					},
				},
			},
			list: []string{""},
		},
		{
			name: "Positive case with name",
			client: &ec2Mock{
				describeInstancesOutput: &ec2.DescribeInstancesOutput{
					Reservations: []ec2.RunInstancesOutput{
						{
							Instances: []ec2.Instance{
								{
									Tags: []ec2.Tag{
										{
											Key:   aws.String("Name"),
											Value: aws.String("some CI node is here"),
										},
									},
									InstanceId: aws.String(""),
									State:      &ec2.InstanceState{Code: aws.Int64(16)},
								},
							},
						},
					},
				},
			},
			list: []string{""},
		},
		{
			name:   "Network error",
			client: &ec2Mock{err: errors.New("")},
			error:  true,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			client := &EC2Client{
				client: v.client,
			}
			list, err := client.GetInstancesForCI()
			if v.error {
				assert.Error(t, err)
				log.Println(err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, v.list, list)
			}
		})
	}
}

func TestEC2Client_TerminateInstance(t *testing.T) {
	cases := []struct {
		name   string
		client *ec2Mock
		ids    []string
		error  bool
	}{
		{
			name:   "Positive case",
			client: &ec2Mock{},
			ids:    []string{""},
		},
		{
			name:   "Network error",
			client: &ec2Mock{err: errors.New("")},
			ids:    []string{""},
			error:  true,
		},
		{
			name:   "Empty list of ID",
			client: &ec2Mock{},
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			client := &EC2Client{
				client: v.client,
			}
			err := client.TerminateInstance(v.ids)
			if v.error {
				assert.Error(t, err)
				log.Println(err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
