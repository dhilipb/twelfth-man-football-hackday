package handlers

const (
	appOkStatus  = "ok"
	appErrStatus = "error"
)

type appResBody struct {
	Status string
}

type appErrorResBody struct {
	appResBody
	Message string
}

func NewAppErrResBody(message string) *appErrorResBody {
	return &appErrorResBody{appResBody{appOkStatus}, message}
}
