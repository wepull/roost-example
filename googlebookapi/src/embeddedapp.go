// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
//"github.com/ZB-io/roost-desktop/samples/googlebook-api/contrib/activity/log",
//"../contrib/activity/log",
const flogoJSON string = `{
  "name": "googlebook",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.1.0",
  "description": "",
  "imports": [
    "github.com/ZB-io/zbio-example/googlebookapi/src/contrib/activity/actreturn",
    "github.com/ZB-io/zbio-example/googlebookapi/src/contrib/activity/log",
    "github.com/ZB-io/zbio-example/googlebookapi/src/contrib/activity/rest",
    "github.com/ZB-io/zbio-example/googlebookapi/src/contrib/function/string",
    "github.com/ZB-io/zbio-example/googlebookapi/src/contrib/trigger/rest",
    "github.com/project-flogo/flow"
  ],
  "triggers": [
    {
      "id": "receive_http_message",
      "ref": "#rest",
      "name": "Receive HTTP Message",
      "description": "Simple REST Trigger",
      "settings": {
        "port": 9999
      },
      "handlers": [
        {
          "settings": {
            "method": "GET",
            "path": "/books/:isbn"
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:get_books"
            },
            "input": {
              "isbn": "=string.concat(\"isbn:\", $.pathParams.isbn)"
            },
            "output": {
              "code": "=$.code",
              "data": "=$.message"
            }
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:get_books",
      "data": {
        "name": "GetBooks",
        "metadata": {
          "input": [
            {
              "name": "isbn",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "code",
              "type": "integer"
            },
            {
              "name": "message",
              "type": "any"
            }
          ]
        },
        "tasks": [
          {
            "id": "log_2",
            "name": "Log Message",
            "description": "Simple Log Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=string.concat(\"Getting string book data for \", $flow.isbn)"
              }
            }
          },
          {
            "id": "rest_3",
            "name": "Invoke REST Service",
            "description": "Simple REST Activity",
            "activity": {
              "ref": "#rest",
              "input": {
                "queryParams": {
                  "mapping": {
                    "q": "=$flow.isbn"
                  }
                }
              },
              "settings": {
                "method": "GET",
                "uri": "https://www.googleapis.com/books/v1/volumes"
              }
            }
          },
          {
            "id": "actreturn_4",
            "name": "Return",
            "description": "Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "message": {
                    "mapping": {
                      "title": "=$activity[rest_3].data.items[0].volumeInfo.title",
                      "publishedDate": "=$activity[rest_3].data.items[0].volumeInfo.publishedDate",
                      "description": "=$activity[rest_3].data.items[0].volumeInfo.description"
                    }
                  },
                  "code": 200
                }
              }
            }
          },
          {
            "id": "log_8",
            "name": "Log",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "addDetails": false,
                "usePrint": false,
                "message": "=string.concat( \"Title: \",\n  $activity[rest_3].data.items[0].volumeInfo.title, \n  \" PublishedDate: \",\n  $activity[rest_3].data.items[0].volumeInfo.publishedDate,\n  \" Description: \",\n  $activity[rest_3].data.items[0].volumeInfo.description\n)"
              }
            }
          }
        ],
        "links": [
          {
            "from": "log_2",
            "to": "rest_3"
          },
          {
            "from": "rest_3",
            "to": "actreturn_4"
          },
          {
            "from": "rest_3",
            "to": "log_8",
            "type": "expression",
            "value": "true"
          }
        ]
      }
    }
  ]
}`
const engineJSON string = ``

func init() {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}
