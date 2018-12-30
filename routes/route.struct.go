package routes


type Route struct {
	Path       string
}


type RouteMethod interface {
	Controller()
}