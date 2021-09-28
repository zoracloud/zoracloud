package controllers

const DefaultContainerPort = 8888
const DefaultServingPort = 80
const AnnotationRewriteURI = "notebooks.zoracloud.ai/http-rewrite-uri"
const AnnotationHeadersRequestSet = "notebooks.zoracloud.ai/http-headers-request-set"

// The default fsGroup of PodSecurityContext.
// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#podsecuritycontext-v1-core
const DefaultFSGroup = int64(100)
