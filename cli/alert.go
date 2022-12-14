package cli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	"k8s.io/client-go/kubernetes"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Response struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
}
type Labels struct {
	AdminAlert  string `json:"admin_alert"`
	Alertgroup  string `json:"alertgroup"`
	Alertname   string `json:"alertname"`
	ClusterName string `json:"cluster_name"`
	Namespace   string `json:"namespace"`
	Pod         string `json:"pod"`
}

type Data struct {
	Labels Labels `json:"labels,omitempty"`
}

var (
	Client HTTPClient
)

func GetVMAlertBackendSize(server string, clientset *kubernetes.Clientset) (string, string) {
	// Initialisation of GET request
	res, err := http.Get(server)
	if err != nil {
		log.Fatal().Msgf("Error in GET request %s ", err)
	}

	// Closing request
	defer res.Body.Close()

	// Reading Body content
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal().Msgf("Error in reading body %s ", err)
	}

	// Serialising return of Body into JSON
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal().Msgf("Error in reading body %s ", err)
	}

	// Init var alertFiring with false by default

	// Parsing Json return to match Alertname with haproxyBackendSizeDivergence
	for _, alerts := range response.Data {
		if (alerts.Labels.Alertname == "haproxyBackendSizeDivergence") && (len(alerts.Labels.Pod) > 0) {
			log.Info().Msgf("Alert %s is firing on pod %s deletion ongoing", alerts.Labels.Alertname, alerts.Labels.Pod)
			podName := alerts.Labels.Pod
			namespace := alerts.Labels.Namespace
			//Proceeding to the deletion of pod if alert is firing
			//DeletePod(podName, clientset, namespace)
			return podName, namespace
		} else {
			log.Info().Msgf("No pod in state of backendsize divergence")
			continue
		}
	}
	return "", ""
}
