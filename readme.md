# Coderhouse Challenge

**Challenge**: https://docs.google.com/document/d/1LvxgOpcpkeILxrVOR5j-YLPuxEEUEVEEYHbH09gepIo/edit



## Intrucciones de uso:
1. Crear un usuario indicando: Nombre de usuario, Contraseña y Moneda de las cuentas.
2. Autenticar a dicho usuario, el endpoint de autenticacion devolvera un **token**.
3. Ejecutar las demas requests agregando en el header “**token**” el valor provisto en el paso anterior.


`En el repositorio se podra encontrar una colección de postman con todas las requests con los campos y headers necesarios para su ejecución.
`

## Endpoints disponibles:

### Create User
-   **Method**: `POST`
-   **URL**: `/user`
-   **Body**:
```
{
	"username": "username",
	"password": "password",
	"currencies": ["USD", "ARS", “EUR”]
}
```

- **Body validations**:
    - username(string): required, min 8, max 32, alphanumeric
    - password(string): required, min 8, max 64
    - currencies([string]): required, min 1, possible values(“ARS”, “USD”, “EUR”)


----------

### Authorize User

-   **Method**: `POST`
-   **URL**: `/authorize`
-   **Body**:
 ```
 {
	"username": "username",
	"password": "password"
}
```

- **Body validations**:
    - username(string): required
    - password(string): required

- **Response**:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluaXN0cmF0b3IyIiwiZXhwIjoxNzI2MTgyODMwfQ.NdxoyOlMZd0TSt27RZtvTvUA7_VfEtw5MI_9sedvDYc"
}
```


----------

### Get Balance

-   **Method**: `GET`
-   **URL**: `/account/balance`
-   **Headers**:
    - token(string)
- **Response**:
```
[
	{
		"balance": 18,
		"currency": "USD"
	},
	{
		"balance": 0,
		"currency": "ARS"
	}
]
```

  ----------

### Deposit

-   **Method**: `POST`
-   **URL**: `/account/deposit`
-   **Headers**:
    - token(string)
-   **Body**:
```
{
    "amount": 123,
    "currency": "USD"
}
```

- **Body validations**:
  - amount(float): required, gt 0
  - currency(string): required, one of (“ARS”, “USD”, “EUR”)
- **Response**:
```
{
    "balance": 20,
    "currency": "USD"
}
```

----------

### Withdraw

-   **Method**: `POST`
-   **URL**: `/account/withdraw`
-   **Headers**:
    - token(string)
-   **Body**:
```
{
    "amount": 123,
    "currency": "USD"
}
```

- **Body validations**:
    - amount(float): required, gt 0
    - currency(string): required, one of (“ARS”, “USD”, “EUR”)
- **Response**:
```
{
    "balance": 20,
    "currency": "USD"
}
```

----------

### Get Transaction History

-   **Method**: `GET`
-   **URL**: `/account/transactions`
-   **Headers**:
    - token(string)
-   **Query Params**:
    - limit(int)
- **Response**:
```
[
    {
        "amount": 5,
        "type": "deposit",
        "partial_balance": 23,
        "date": "2024-09-12T20:16:41.7043418-03:00",
        "currency": "USD"
    },
    {
        "amount": 2,
        "type": "withdraw",
        "partial_balance": 18,
        "date": "2024-09-12T19:59:41.5225468-03:00",
        "currency": "USD"
    }
]
```