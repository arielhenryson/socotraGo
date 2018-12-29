package routes

type Route struct {
	Path string
}

var Routes = []Route{
	{
		Path: "/",
	},
	{
		Path: "/page2",
	},
	{
		Path: "/page3",
	},
}