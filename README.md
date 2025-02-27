# HPA Demo

This demo uses a slightly modified version of the [HorizontalPodAutoscaler Walkthrough](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) example from the Kuberentes documentation.
It creates the following resources:

* `hpa-demo-app` Deployment
* `hpa-demo-service` Service
* `hpa-demo` HorizontalPodAutoscaler with minReplicas=5, maxReplicas=40 and targetCPUUtilizationPercentage=70

To create a cluster, run
```
gcloud beta container clusters \
create-auto hpa-demo --region us-central1 \
--cluster-version "1.32.1-gke.1729000" \
--release-channel "rapid" \
--hpa-profile performance
```

To deploy, run
```
kubectl apply -f manifests
```

For generating load, the `load.sh` script is provided.  It uses the simple load test application [hey](https://github.com/rakyll/hey) to generate load for 2 minutes with 10 concurrent clients.  You can adjust the parameters in the script as you see fit.
