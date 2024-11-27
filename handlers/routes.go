package handlers

func addRoutes(s Server) {
    mux := s.router
    
    mux.Handle("GET /{$}", s.indexGet())
}
