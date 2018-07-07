
## golgoquery

### First Example

```
{
  "uriflows": [
    {
      "uri": "http://example.com/",
      "queryflows": [
        {
          "queries": [
            {
              "selector": ["div", "p"],
              "attribute": "text",
              "action": "Echo"
            },
            {
              "selector": ["div", "a"],
              "attribute": "href",
              "action": "Debug"
            }
          ]
        }
      ]
    }
  ]
}
```

* fetchs "uri"
* picks every query from "queryflows"
> * first "query", pulls innerHTML text from "div.p" selector element; passes list of results to Method mapped to "Dummy"
> * then next "query", pulls href from "div.a" selector element; passes list of results to Method mapped to "Dummy"

---

### Second Example

```
{
  "uriflows": [
    {
      "uri": "http://example.com/",
      "queryflows": [
        {
          "queries": [
            {
              "selector": ["div", "p"],
              "attribute": "text",
              "action": "Dummy"
            },
            {
              "selector": ["div", "a"],
              "attribute": "href",
              "action": "Dummy"
            }
          ]
        }
      ]
    }
  ]
}
```


---
