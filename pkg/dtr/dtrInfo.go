package dtr

import (
	"encoding/json"
	"fmt"

	"github.com/thebsdbox/diver/pkg/dtr/types"
)

func (c *Client) dtrClusterStatus() (*dtrTypes.DTRCluster, error) {

	url := fmt.Sprintf("%s/api/v0/meta/cluster_status?refresh_token=%s", c.DTRURL, c.Token)

	response, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}
	//log.Debugf("%v", string(response))
	var info dtrTypes.DTRCluster

	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

//ListReplicas -
func (c *Client) ListReplicas() error {
	cluster, err := c.dtrClusterStatus()
	if err != nil {
		return err
	}

	replicas := cluster.ReplicaHealth
	fmt.Printf("Replica\t \tStatus\n")

	for replica, status := range replicas {
		fmt.Printf("%s\t %s\n", replica, status)
	}
	return nil
}
