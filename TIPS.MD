# Commands

## Build

- go build FILE_PATH

## Run

- go run FILE_PATH

## MOD INIT

- go mod init MODULE_NAME

## MOD TIDY

- go mod tidy
  > install all project dependencies

## CLEAN

- go clean -cache
  > used to clean cache project
  > to clean only test cache -testcache
  > https://til.cybertec-postgresql.com/post/2019-11-07-How-to-turn-off-test-caching-for-golang/

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

accepts only boolean values in tests.

example:

```

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

examples:

```
Example 1:
  for i:=0; i<len(sites); i++ {
	//code...
	}
```

```
Example 2:
  // this way the range returns the i of the index and the item itself in the site variable
  for i, site := range sites {
		fmt.Println("2Site:", site)
	}
```

# Packages

## fmt

- Println
- Scanf
- Scan

### o uso do &

The use of "&" is to use the POINTER of the variable so you can use the terminal to assign a value to a variable
example:

```

var option int
fmt.Scanf("%d", &option)
or
fmt.Scan(&option)

```

## reflect

- TypeOf
  > reflect.TypeOf(VARIABLE_HERE)

## os

- Exit example: os.Exit(0) - Success || os.Exit(-1) - Fail
- Open ( for read file in (memory/pointer))
- ReadFile ( for read file in []byte) to show string(FUNC_HERE)
- OpenFile
- WriteString

## net/http

- Get(URL_HERE)

## strings

- TrimSpace()

## time

- Sleep
- Second

## bufio

- NewReader
  - ReadString

## ioutil

- ReadFile
