package main
import(
  //cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
  "os"
  "k8s.io/federation/pkg/kubefed/util"
  //fedclient "k8s.io/federation/client/clientset_generated/federation_clientset"
  //restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
//  "github.com/golang/glog"
	//"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
  "k8s.io/federation/cmd/federation-nginx-controller/pkg"
  "fmt"
  "time"
)


func main(){

//  loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
//	configOverrides := &clientcmd.ConfigOverrides{}
//  kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
  config := util.NewAdminConfig(clientcmd.NewDefaultPathOptions())
  fedClientSet, err := config.FederationClientset("kfed", os.Getenv("KUBECONFIG"))

  controller, err := controller.NewNGINXFedIngressController(fedClientSet, 10*time.Second)

  // glog.Infof("kubeconfig: %v+",kubeConfig)
  if err == nil{
    fmt.Printf("Controller: %v", controller)
  }
}
