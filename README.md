# sample-sls

Sample-sls (Sample simple log service) is a streaming log analysis project that applies the dataflow model. It adopts the front-end and back-end separation architecture of go+vue and supports distributed deployment.

## Check Points

- [ ] Swagger for RESTFul API
- [x] Generate Log IDs Using **Snowflake** Algorithm
- [ ] Data Dashboard

## Flow

> LogStream -> `converter` -> `metadata` -> `formater` -> `publisher`

TestInput

```shell
printf "MESSAGE=hello world\nPRIORITY=6\n\n" | systemd-cat
```

Usage

```shell
tail -F /var/log/nginx/*.log       |\  # outputs log lines
  ssls -c ./config.yml             |\  # outputs Journal Export Format
  systemd-cat                          # send to local/remote journald
```