package clients

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"

	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EC2Client used to encapsulate EC2 specifics
type EC2Client struct {
	client ec2iface.EC2API
}

// GetEC2Client is a constructor for EC2Client
func GetEC2Client() (*EC2Client, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cfg.Region = endpoints.EuWest1RegionID
	return &EC2Client{client: ec2.New(cfg)}, nil
}

// GetInstancesForCI returns list of instances using for CI which are running right now
func (v *EC2Client) GetInstancesForCI() ([]string, error) {
	req := v.client.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	resp, err := req.Send()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ids := make([]string, 0)
	for _, i := range resp.Reservations {
		for _, j := range i.Instances {
			name := getInstanceName(j.Tags)
			if name == "" || strings.Contains(name, "CI node") {
				if *j.State.Code == 16 {
					ids = append(ids, *j.InstanceId)
				}
			}
		}
	}

	return ids, nil
}

// TerminateInstance terminates EC2 instance using provided instance ID
func (v *EC2Client) TerminateInstance(IDs []string) error {
	if len(IDs) == 0 {
		log.Println("List of IDs is empty!")
		return nil
	}
	req := v.client.TerminateInstancesRequest(&ec2.TerminateInstancesInput{
		InstanceIds: IDs,
	})
	_, err := req.Send()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getInstanceName(tags []ec2.Tag) string {
	for _, v := range tags {
		if *v.Key == "Name" {
			return *v.Value
		}
	}
	return ""
}
