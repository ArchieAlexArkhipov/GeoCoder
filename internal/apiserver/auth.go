package apiserver

import (
  "net/http"
  "geocoder/internal/utils"
  "os"
)

var AuthenticationMiddleware = func(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    response := make(map[string] interface{})
    tokenHeader := r.Header.Get("Authorization")

    if tokenHeader == "" { // Токен отсутствует, возвращаем  403 http-код Unauthorized
      response = utils.Message(false, "Missing auth token")
      w.WriteHeader(http.StatusForbidden)
      w.Header().Add("Content-Type", "application/json")
      utils.Respond(w, response)

      return
    }

    authToken, exists := os.LookupEnv("AUTH_TOKEN")

    if !exists {
         response = utils.Message(false, "Authorization token not found!")
         w.WriteHeader(http.StatusForbidden)
         w.Header().Add("Content-Type", "application/json")
         utils.Respond(w, response)

         return
    }

    if authToken != tokenHeader { // Токен неправильный
      response = utils.Message(false, "Token is not valid")
      w.WriteHeader(http.StatusForbidden)
      w.Header().Add("Content-Type", "application/json")
      utils.Respond(w, response)

      return
    }

    next.ServeHTTP(w, r)
  });
}
