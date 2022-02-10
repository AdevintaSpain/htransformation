package deleter

import (
	"net/http"

	"github.com/tommoulard/htransformation/pkg/types"
)

func Handle(_ http.ResponseWriter, req *http.Request, rule types.Rule) error {
	req.Header.Del(rule.Header)

	return nil
}