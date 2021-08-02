# istio-routing-issue

## Overview

Here we have two services. Service A calls service B.

1. Container A listens on port 3000 and forwards request to Service B (http://istio-routing-issue-b.istio-routing-issue.svc.cluster.local).
2. Container B listens on port 3000 and 3001. On port 3000 it responds with `1`, on port 3001 it responds with `2`.
3. Container A responds with status code `200` if it gets `1` and with `400` if it gets `2` from service B.
4. VS of service B configured for mesh forwards traffic to port 3001.

Most of the time exactly this happens. But if you create a new `A` pod (or kill one to have a new one), wait for `istio-proxy` to become ready,
and then send the load (50rps was enough in my examples), you will see results like:

```bash
$ ./load.sh
Requests      [total, rate, throughput]         250, 50.22, 22.67
Duration      [total, attack, wait]             5.073s, 4.978s, 94.992ms
Latencies     [min, mean, 50, 90, 95, 99, max]  87.442ms, 116.672ms, 97.619ms, 161.549ms, 280.066ms, 324.486ms, 403.418ms
Bytes In      [total, mean]                     111452, 445.81
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           46.00%
Status Codes  [code:count]                      200:115  400:135  
Error Set:
400 Bad Request
```

So it means the traffic sometimes goes to a wrong port on service B.

## Environment

1. Istio 1.6.14 (outboundTrafficPolicy.mode = ALLOW_ANY)
2. Kubernetes v1.18.9

Others, which I don't think is meaningful for this problem:

1. Helm v3.4.2
2. Golang 1.15.6 for compiling services
3. Vegeta 12.8.3 for load testing

## How to reproduce

Build and deploy the containers into a Kubernetes cluster:

```bash
./init-namespace.sh
./store-images.sh
ENDPOINT=example.com ./deploy.sh
```

Load the service A with `ADDRESS=https://example.com/ ./load.sh` (set the address for service A first)
