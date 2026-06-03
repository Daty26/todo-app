package web_transport_http

import (
	"net/http"

	core_logger "github.com/Daty26/todo-app/internal/core/logger"
	core_http_response "github.com/Daty26/todo-app/internal/core/transport/http/response"
)

func (h*WebHTTPHandler) GetMainPage(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPesponseHandler(log, w)
	html, err := h.webService.GetMainPage()
	if err != nil{
		responseHandler.ErrorResponse(err, "failed to get index.html")
	}
	responseHandler.HTMLResponse(html)
	//
}