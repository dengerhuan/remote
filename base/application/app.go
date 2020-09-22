package app

type Application interface {
	Run()
}

func RunApplication(app Application, params ...string) {
	// se env
	app.Run()
}
