## CRD controller
This directory contains code for custom Kubernetes CRDs and controllers used by
the orchestration system:

* ScheduledWorkflow
* Viewer

The following are guidelines on developing and running these controllers.

### Prerequisites

First, we need to generate the API client code from the API specification.

Get the dependencies:

```
dep ensure
go get -u ./...
```

Generate the API client code from the API specification:

```
./hack/update-codegen.sh
```
