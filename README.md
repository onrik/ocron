# ocron

### Config example

```yaml
tasks:
  - name: backup
    spec: "0 0 * * *"
    env:
      AWS_ACCESS_KEY_ID: vdBj6WUReTWqpZCLdnC4pxjgs
      AWS_SECRET_ACCESS_KEY: gNwWM5bcdS8Nwo9o8veJNL7fYJHFXboDmuUoW8k5
      S3_ENDPOINT_URL: https://storage.googleapis.com
      S3_BUCKET: test
      TG_ADMIN: 123456789
      TG_TOKEN: "987654321:i4phwufPXqMPbDrtwqz9wohTptNP2ukrg8c"
    script:
      - tar -cf home.tar -C / home
      - s5cmd cp home.tar s3://${S3_BUCKET}/`date +'%Y/%m'`/home.tar
    on_success:
      - "curl -X POST -d '{\"chat_id\": \"${TG_ADMIN}\", \"text\": \"✅ backup success\"}' -H 'Content-Type: application/json' https://api.telegram.org/bot${TG_TOKEN}/sendMessage"
    on_error:
      - "curl -X POST -d '{\"chat_id\": \"${TG_ADMIN}\", \"text\": \"❌ backup error\"}' -H 'Content-Type: application/json' https://api.telegram.org/bot${TG_TOKEN}/sendMessage"
```
