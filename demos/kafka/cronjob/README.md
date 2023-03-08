## `px-kafka` CronJob

This CronJob enables a time-based scale up / down of the `shipping` deployment for the [Consumer Rebalancing demo scenario](https://docs.px.dev/tutorials/pixie-101/kafka-monitoring/#consumer-rebalancing-duration) in the `px-kafka` demo.

Every 10 minutes, the `shipping` deployment will scale to 3 replicas followed by a scale down to 1 replica 10 minutes later.

```
kubectl apply -f setup.yaml
kubectl apply -f cronjob.yaml
```