package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/op-server/pkg/logger"
	"github.com/op-server/pkg/server/handlers/cluster/v1/meta/paramtypes"
	"github.com/parnurzeal/gorequest"
	"k8s.io/client-go/util/homedir"
)

func main() {
	serverAddr := flag.String("op-server", "http://127.0.0.1:8080", "op server address")
	clusterName := flag.String("name", "", "Input Cluster Name")
	clusterRegion := flag.String("region", "", "Input Cluster Region")
	clusterAZ := flag.String("az", "", "Input cluster az")
	clusterDepartment := flag.String("department", "", "Input cluster department")
	clusterEnv := flag.String("env", "", "Input cluster environment")
	clusterConfPath := flag.String("conf_path", path.Join(homedir.HomeDir(), ".kube", "config"), "Input cluster kube conf path")
	clusterAttr := flag.String("attr", "", "Input cluster attribute")

	flag.Parse()

	config, err := getClusterConf(*clusterConfPath)
	if err != nil {
		log.Fatalf("load kube config fail,%v", err)
	}

	attrMap := splitClusterAttr(*clusterAttr)

	if err := register(*serverAddr, *clusterName, *clusterRegion, *clusterAZ, *clusterDepartment, *clusterEnv, string(config), attrMap); err != nil {
		log.Fatal(err.Error())
	}

}

func getClusterConf(path string) ([]byte, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return yamlFile, nil
}

func splitClusterAttr(attr string) map[string]string {
	result := make(map[string]string)
	tmpList := strings.Split(attr, ",")
	for _, item := range tmpList {
		kvPair := strings.Split(item, "=")
		if len(kvPair) != 2 {
			continue
		}
		result[kvPair[0]] = kvPair[1]
	}
	return result
}

func register(addr string, name string, region string, az string, department string, env string, conf string, attr map[string]string) error {
	param := paramtypes.ClusterMetaRequestParam{
		Name:        &[]string{name}[0],
		Region:      &[]string{region}[0],
		Az:          &[]string{az}[0],
		Department:  &[]string{department}[0],
		Environment: &[]string{env}[0],
		KubeConf:    conf,
		Attribute:   attr,
	}

	request := gorequest.New()
	resp, body, errs := request.Post(fmt.Sprintf("%s/api/op/cluster/v1/meta/register", addr)).
		AppendHeader("Content-Type", "application/json").Send(param).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).End()

	if errs != nil {
		return fmt.Errorf("%v", errs)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("register fail,status code is %d", resp.StatusCode)
	}
	logger.Infof("register cluster resp is %s", body)
	return nil
}
