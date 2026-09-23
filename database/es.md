```bash
curl -XGET "http://10.143.160.156:9200/analysis_notice_state/_search?pretty" -u 'elasticcx:lDZ*QPCKYpGkwQFc51R-' -H "Content-Type: application/json" -d'
{
  "query": {
    "function_score": {
      "query": { "match_all": {} },
      "random_score": {}, 
      "boost_mode": "replace"
    }
  },
  "size": 1 
}
'

curl -XGET "http://10.143.160.156:9200/analysis_notice_state/_count?pretty" -u 'elasticcx:lDZ*QPCKYpGkwQFc51R-' -H "Content-Type: application/json" -d'
{
  "query": {
    "range": {
      "etl_time_": {
        "gte": "202507211500",
        "lte": "202507221600",
        "format": "yyyyMMddHHmmss"
      }
    }
  }
}
'
```

```bash
curl -u 'username:password' '10.143.160.156:9200/analysis_notice_state/_count'
```

