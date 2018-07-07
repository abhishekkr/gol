
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
              "action": "debug"
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
              "action": "Echo"
            },
            {
              "selector": ["div", "a"],
              "attribute": "href",
            }
          ],
          "action": "debug"
        }
      ]
    }
  ]
}
```

* fetchs "uri"
* processes queries in queryflows, but since just one "query" has an action, action will be triggered only in that query processing
* after all "queries" processed for "queryflows", since "queryflows" has an action... it will be called with list of all results and index "1"

---

NOTE:

* if no action is passed, "debug" action will be called for it printing the result unless action name has been mentioned as "~"

---
