package deleter

import (
	"net/http"

	"github.com/gustauperez/htransformation/pkg/types"
)

func Handle(_ http.ResponseWriter, req *http.Request, rule types.Rule) {
	req.Header.Del(rule.Header)
}
