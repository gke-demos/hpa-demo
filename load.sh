#!/bin/sh

kubectl run -i --tty --rm hey --image us-docker.pkg.dev/gke-demos-345619/hey/hey --restart=Never --  -c 20 -z 2m  http://hpa-demo-service