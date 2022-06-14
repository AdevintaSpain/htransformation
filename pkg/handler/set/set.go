package set

import (
	"net/http"

	"github.mpi-internal.com/devops-re--htransformation/pkg/types"
)

func Handle(_ http.ResponseWriter, req *http.Request, rule types.Rule) {
	req.Header.Set(rule.Header, rule.Value)
}
