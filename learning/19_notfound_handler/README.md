# Deploying Go App to Google App Engine

## Prerequisites

### 1. Install Google Cloud SDK
Make sure you have the **Google Cloud SDK** installed. You can download it [here](https://cloud.google.com/sdk/docs/install).
If you're using macOS, install it via Homebrew:
```sh
brew install --cask google-cloud-sdk
```

### 2. Install Go
Ensure you have **Go** installed (version 1.12+). You can check with:
```sh
go version
```
If Go is not installed, download it from [Go’s official website](https://go.dev/dl/).

### 3. Authenticate with Google Cloud
Login to your Google Cloud account:
```sh
gcloud auth login
```
After authentication, choose the correct account.

### 4. Set Your Project
Ensure you are working with the correct Google Cloud project:
```sh
gcloud config set project [YOUR_PROJECT_ID]
```
To verify the project is set correctly:
```sh
gcloud config list
```

### 5. Enable Billing (Required for Deployment)
Your project must have a billing account attached. To enable billing:
- Visit [Google Cloud Billing](https://console.cloud.google.com/billing).
- Set up a billing account if you don’t have one.
- Link your project to the billing account:
  ```sh
  gcloud beta billing projects link [YOUR_PROJECT_ID] --billing-account=[YOUR_BILLING_ACCOUNT_ID]
  ```

## Steps to Deploy

### 1. Prepare `app.yaml`
Create or update your `app.yaml` file in your project folder:
```yaml
runtime: go121  # Use a supported Go version
service: default  # Optional, only needed for multiple services

automatic_scaling:
  max_instances: 10  # Set your desired maximum number of instances

handlers:
- url: /.*
  script: auto
  secure: always
```

### 2. Deploy Your Application
Ensure you are in the root folder of your Go project (where `app.yaml` is located), then deploy your app with:
```sh
gcloud app deploy
```

### 3. Choose a Region for App Engine (First-Time Only)
If this is your first deployment, you'll be asked to choose a region:
```sh
Please choose the region where you want your App Engine application located:

 [1] asia-east1
 [2] asia-east2
 [3] asia-northeast1
 ...
 [17] us-central
 [18] us-east1
 [19] us-east4
 [20] us-west1
 [21] us-west2
 [22] us-west3
 [23] us-west4
 [24] cancel

Please enter your numeric choice:  17
```

### 4. Deploying After Region Setup
After the region is set, deploy your app again:
```sh
gcloud app deploy
```

### 5. Browse the Deployed App
Once the deployment is complete, you can open your app in the browser:
```sh
gcloud app browse
```
Alternatively, go to:
```
https://[YOUR_PROJECT_ID].appspot.com
```

## Troubleshooting

- **Missing Billing Account**: If you see errors about missing billing, enable billing as described in Step 5.
- **Permissions Errors**: Ensure your account has the `roles/appengine.deployer` role. Grant it using:
  ```sh
  gcloud projects add-iam-policy-binding [YOUR_PROJECT_ID] \
    --member="user:[YOUR_EMAIL]" \
    --role="roles/appengine.deployer"
  ```
- **Logs and Debugging**: Check logs if your app isn't working correctly:
  ```sh
  gcloud app logs tail -s default
  ```

## Additional Notes

- **Scaling**: Adjust `automatic_scaling` in your `app.yaml` to control the number of instances.
- **Environment Variables**: Add environment variables in `app.yaml` under `env_variables`:
  ```yaml
  env_variables:
    MY_ENV_VAR: "value"
  ```
- **Multiple Services**: If your app has multiple services, specify them in `app.yaml`.

## Helpful Links
- [Google Cloud SDK Documentation](https://cloud.google.com/sdk/docs)
- [Google App Engine Documentation](https://cloud.google.com/appengine/docs)

