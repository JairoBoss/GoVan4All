package controllers

import "github.com/hcastellanos-dev/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Company routes
	s.Router.HandleFunc("/companies", middlewares.SetMiddlewareJSON(s.CreateCompany)).Methods("POST")
	s.Router.HandleFunc("/companies", middlewares.SetMiddlewareJSON(s.GetCompanies)).Methods("GET")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareJSON(s.GetCompany)).Methods("GET")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCompany))).Methods("PUT")
	s.Router.HandleFunc("/companies/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCompany)).Methods("DELETE")

	// Driver routes
	s.Router.HandleFunc("/drivers", middlewares.SetMiddlewareJSON(s.CreateDriver)).Methods("POST")
	s.Router.HandleFunc("/drivers", middlewares.SetMiddlewareJSON(s.GetDrivers)).Methods("GET")
	s.Router.HandleFunc("/drivers/{id}", middlewares.SetMiddlewareJSON(s.GetDriver)).Methods("GET")
	s.Router.HandleFunc("/drivers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateDriver))).Methods("PUT")
	s.Router.HandleFunc("/drivers/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDriver)).Methods("DELETE")

	// Route routes
	s.Router.HandleFunc("/routes", middlewares.SetMiddlewareJSON(s.CreateRoute)).Methods("POST")
	s.Router.HandleFunc("/routes", middlewares.SetMiddlewareJSON(s.GetRoutes)).Methods("GET")
	s.Router.HandleFunc("/routes/{id}", middlewares.SetMiddlewareJSON(s.GetRoute)).Methods("GET")
	s.Router.HandleFunc("/routes/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateRoute))).Methods("PUT")
	s.Router.HandleFunc("/routes/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteRoute)).Methods("DELETE")

	// User routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

}
