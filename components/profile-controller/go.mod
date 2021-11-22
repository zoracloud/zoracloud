module github.com/zoracloud/zoracloud/components/profile-controller

go 1.16

require (
	github.com/aws/aws-sdk-go v1.42.9
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/tidwall/gjson v1.11.0
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	google.golang.org/api v0.60.0
	istio.io/api v0.0.0-20211119221920-e0ac4ca57eb8
	istio.io/client-go v1.12.0
	k8s.io/api v0.22.4
	k8s.io/apimachinery v0.22.4
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.6.3
)
