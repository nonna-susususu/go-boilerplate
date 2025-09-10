package health

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return new(HealthHandler)
}
