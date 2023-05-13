# WA-02

Pub/Sub exercise

https://cloud.google.com/run/docs/tutorials/pubsub
https://github.com/GoogleCloudPlatform/golang-samples/tree/main/run/pubsub

```sh
gcloud pubsub topics create myRunTopic

gcloud builds submit --tag gcr.io/${PROJECT_ID}/pubsub

gcloud run deploy pubsub-tutorial --image gcr.io/${PROJECT_ID}/pubsub --no-allow-unauthenticated

gcloud iam service-accounts create cloud-run-pubsub-invoker \
--display-name "Cloud Run Pub/Sub Invoker"

gcloud run services add-iam-policy-binding pubsub-tutorial \
--member=serviceAccount:cloud-run-pubsub-invoker@${PROJECT_ID}.iam.gserviceaccount.com \
--role=roles/run.invoker

gcloud pubsub subscriptions create myRunSubscription --topic myRunTopic \
--ack-deadline=600 \
--push-endpoint=${SERVICE_URL}/ \
--push-auth-service-account=cloud-run-pubsub-invoker@${PROJECT_ID}.iam.gserviceaccount.com
```

```sh
gcloud pubsub topics publish myRunTopic --message "Runner"
```