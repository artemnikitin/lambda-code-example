package handler

import (
	"log"

	"github.com/artemnikitin/delex-code-example/clients"
)

// TerminateUnusedInstancesHandler represents Lambda entry point for termination of unused CI instances
func TerminateUnusedInstancesHandler() error {
	ec2, err := clients.GetEC2Client()
	if err != nil {
		log.Println(err)
		return err
	}

	ids, err := ec2.GetInstancesForCI()
	if err != nil {
		log.Println(err)
		return err
	}

	err = ec2.TerminateInstance(ids)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
