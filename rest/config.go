package rest

const PORT = ":8802"
const URL = "http://localhost:8802"

type Response struct {
	Message string `json:"message"`
}
