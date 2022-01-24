# Commands

## Build

- go build FILE_PATH

## Run

- go run FILE_PATH

# VAR declaration

- var VAR_NAME TYPE
  > example: var name string = "Gustavo"

> if you don't set a value in the variable declaration, the program
> assumes the value as empty

## Types

- string
- float32
- float64
- int
- array || slices

### Array

example:

```
  var sites [6]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.google.com.br/"
	sites[2] = "https://www.youtube.com/"
	sites[3] = "https://www.facebook.com/"
	sites[4] = "https://www.alura.com.br/"
```

### Slices

example of declaration:

```
  nomes := []string{"Douglas", "Daniel", "Bernardo"}

```

- len()
- cap()
- append(nomes, "Aparecida")

# IF

aceita somente valores booleanos nos testes.

example:

````

if option == 1 {
fmt.Println("Start the program")
} else if option == 2 {
fmt.Println("Show Logs")
} else if option == 0 {
fmt.Println("Exit")
} else {
fmt.Println("Invalid option")
}

```

# Switch

Não é necessário utilizar o Break

example:

```

switch option {
case 1:
fmt.Println("Start the program")
case 2:
fmt.Println("Show Logs")
case 0:
fmt.Println("Exit")
default:
fmt.Println("Invalid option")
}

```

# for

# Packages

## fmt

- Println
- Scanf
- Scan

### o uso do &

O uso do "&" é para usar o PONTEIRO da variável assim podendo usar o terminal para atribuir um valor a uma variável
example:

```

var option int
fmt.Scanf("%d", &option)
or
fmt.Scan(&option)

```

## reflect

- TypeOf

## os

- Exit example: os.Exit(0) - Success || os.Exit(-1) - Fail
```
````
