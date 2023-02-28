# Go DDNS Client
Now only supprt tencentcloud DNSPod

### Config file example
```json
{
    "UpdateInterval": 600,
    "PubIPUrls": [
        "https://ident.me",
        "https://api.ipify.org"
    ],
    "Domains": [
        {
            "Domain": "abc.com",
            "Provider": "tencent",
            "SecretId": "your tencent secret id",
            "SecretKey": "your tencent secret key",
            "SubDomains": [
                "test"
            ]
        }
    ]
}
```
* UpdateInterval: dns update interval (unit: seconds)
* PubIPUrls: the urls you can get public ip from response body


### TencentCloud
User permission
```json
{
    "version": "2.0",
    "statement": [
        {
            "effect": "allow",
            "action": [
                "dnspod:DescribeRecordList",
                "dnspod:ModifyRecord"
            ],
            "resource": [
                "*"
            ]
        }
    ]
}
```
