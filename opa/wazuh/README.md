もしアラートが Ignore されているものであれば 400 を、そうでなければ 200 を返す。

```shell
> go run main.go

> curl -i -H 'Content-Type: application/json' -d @alert.json localhost:8080/eval
HTTP/1.1 100 Continue

HTTP/1.1 200 OK # 👈
Date: Wed, 06 Jan 2021 08:54:34 GMT
Content-Length: 0

> curl -i -H 'Content-Type: application/json' -d @should-be-ignore-alert.json localhost:8080/
eval
HTTP/1.1 100 Continue

HTTP/1.1 400 Bad Request # 👈
Date: Wed, 06 Jan 2021 08:54:47 GMT
Content-Length: 0
```

# Test with `opa eval`

```shell
> opa eval -i alert.json -d wazuh.rego 'data.wazuh.ignore'
{
  "result": [
    {
      "expressions": [
        {
          "value": false,
          "text": "data.wazuh.ignore",
          "location": {
            "row": 1,
            "col": 1
          }
        }
      ]
    }
  ]
}

> opa eval -i should-be-ignore-alert.json -d wazuh.rego 'data.wazuh.ignore'
{
  "result": [
    {
      "expressions": [
        {
          "value": true,
          "text": "data.wazuh.ignore",
          "location": {
            "row": 1,
            "col": 1
          }
        }
      ]
    }
  ]
}
```
