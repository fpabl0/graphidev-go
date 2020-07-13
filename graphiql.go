package main

import (
	"fmt"
	"net/http"
)

/*
	==================================================================================
	THIS FILE WAS BASED ON https://github.com/tonyghita/graphql-go-example
	==================================================================================
*/

// GraphiQL is an in-browser IDE for exploring GraphiQL APIs.
// This handler returns GraphiQL when requested.
//
// For more information, see https://github.com/graphql/graphiql.
type GraphiQL struct {
	GraphqlURL string
}

func (h GraphiQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respond(w, errorJSON("only GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}

	w.Write(h.getWebUI())
}

func (h GraphiQL) getWebUI() []byte {
	resp := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/1.0.3/graphiql.min.css"/>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/1.0.3/graphiql.min.js"></script>
			<style>
				/* Content font size */
				.graphiql-container .CodeMirror {
					font-size: 20px !important;
				}
				/* Search box text size */
				.graphiql-container .search-box > input {
					font-size: 22px !important;
				}
				/* History text size */
				.graphiql-container .history-contents p {
					font-size: 18px !important;
				}
				/* Linter font size */
				.CodeMirror-lint-tooltip {
					font-size: 20px !important;
				}
				.CodeMirror-hint-information .content {
					font-size: 20px !important;
				}
				.CodeMirror-hints {
					font-size: 20px !important;
				}
				/* Titles and buttons font size (Prettify - History) */
				.graphiql-container,
				.graphiql-container button,
				.graphiql-container input {
					font-size: 22px !important;
				}
				/* Docs button text size */
				.graphiql-container .docExplorerShow,
				.graphiql-container .historyShow {
					font-size: 18px !important;
				}
				/* Docs category title font size */
				.graphiql-container .doc-category-title {
					font-size: 22px !important;
				}
				/* History box item height */
				.graphiql-container .history-contents .history-label {
					height: 20px !important;
				}
				/* GraphiQL logo size */
				.graphiql-container .title {
					font-size: 22px !important;
				}
				.graphiql-container .title em {
					font-size: 22px !important;
				}
			</style>
		</head>
		<body style="width: 100%%; height: 100%%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function fetchGQL(params) {
					return fetch("%s", {
						method: "POST",
						body: JSON.stringify(params),
						headers: {
							"Content-Type": "application/json"
						}
					}).then(function (resp) {
						return resp.text();
					}).then(function (body) {
						try {
							return JSON.parse(body);
						} catch (error) {
							return body;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: fetchGQL}),
					document.getElementById("graphiql")
				)
			</script>
		</body>
	</html>
	`, h.GraphqlURL)
	return []byte(resp)
}
