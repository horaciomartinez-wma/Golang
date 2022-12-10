package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
   
   jsonData := []byte(`[
	{
		"socioNumber": 1,
		"nombre": "Sylvester",
		"apellidos": "Stallone",
		"fechaNacimiento": "01/01/2001",
		"fechaDeingreso": "Date.now()",
		"imagenPerfilURL": "https://upload.wikimedia.org/wikipedia/commons/thumb/8/84/Sylvester_Stallone_Cannes_2019.jpg/423px-Sylvester_Stallone_Cannes_2019.jpg"
	},
	{
		"socioNumber": 2,
		"nombre": "Ronnie",
		"apellidos": "Coleman",
		"fechaNacimiento": "13/05/1964",
		"fechaDeingreso": "Date.now()",
		"imagenPerfilURL": "https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/Ronnie_Coleman_Invited_Guest.jpg/397px-Ronnie_Coleman_Invited_Guest.jpg"
	},
	{
		"socioNumber": 3,
		"nombre": "Dwayne La Roca",
		"apellidos": "Johnson",
		"fechaNacimiento": "13/05/1964",
		"fechaDeingreso": "Date.now()",
		"imagenPerfilURL": "https://upload.wikimedia.org/wikipedia/commons/f/f1/Dwayne_Johnson_2%2C_2013.jpg"
	},
	{
		"socioNumber": 4,
		"nombre": "John",
		"apellidos": "Cena",
		"fechaNacimiento": "13/05/1964",
		"fechaDeingreso": "Date.now()",
		"imagenPerfilURL": "https://upload.wikimedia.org/wikipedia/commons/6/60/John_Cena_July_2018.jpg"
	}
]`)

/*
  jsonData, err := os.ReadFile("netlify/functions/hello/socios.json")
  if err != nil {log.Fatal(err)}
*/
  return &events.APIGatewayProxyResponse{
    StatusCode:        200,
    //Body:              "Hello, World!",
    Headers: map[string]string{
            "Content-Type":   "application/json",
            "Access-Control-Allow-Origin":   "*",
        },
    Body:              string(jsonData),
  }, nil
}

func main() {
  lambda.Start(handler)
}
