# Gestion des Erreurs et Contrôle des Inputs en Go

Dans cette documentation, nous allons explorer la gestion des erreurs et le contrôle des inputs en Go, un langage de programmation efficace et concurrent.

## Sommaire

1. [Gestion des Erreurs](#gestion-des-erreurs)
    - [Introduction à la Gestion des Erreurs](#introduction-à-la-gestion-des-erreurs)
    - [Utilisation de `error` et `errors`](#utilisation-de-error-et-errors)
    - [Erreurs Personnalisées](#erreurs-personnalisées)
    - [Gestion des Erreurs avec `defer`, `panic` et `recover`](#gestion-des-erreurs-avec-defer-panic-et-recover)
2. [Contrôle des Inputs](#contrôle-des-inputs)
    - [Validation des Inputs](#validation-des-inputs)
    - [Gestion des Inputs Utilisateur](#gestion-des-inputs-utilisateur)

## Gestion des Erreurs

### Introduction à la Gestion des Erreurs

En Go, la gestion des erreurs se fait principalement à l'aide du type __`error`__. Contrairement à d'autres langages qui utilisent des exceptions, Go préfère un traitement explicite des erreurs.

### Utilisation de __`error`__ et __`errors`__

Le package __`errors`__ permet de créer des erreurs simples.

```go
package main

import (
    "errors"
    "log"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 0)
    if err != nil {
        log.Println("Error:", err)
    } else {
        log.Println("Result:", result)
    }
}
```
### Erreurs Personnalisées
Vous pouvez définir des erreurs personnalisées en implémentant l'interface error.

```go
package main

import (
    "log"
    "fmt"
)

type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func riskyOperation(value int) error {
    if value < 0 {
        return &CustomError{Code: 123, Message: "Negative value provided"}
    }
    return nil
}

func main() {
    err := riskyOperation(-1)
    if err != nil {
        log.Println(err)
    } else {
        log.Println("Operation succeeded")
    }
}
```

### Gestion des Erreurs avec __``defer``__, __``panic``__ et __``recover``__

En Go, les mécanismes defer, panic et recover permettent de gérer les erreurs de manière plus globale et de traiter les situations exceptionnelles. Voici une explication détaillée de chacun de ces mécanismes, ainsi que des exemples pour illustrer leur utilisation.

#### 1. defer
La fonction ``defer`` est utilisée pour programmer l'exécution d'une fonction juste avant que la fonction englobante ne se termine. Les appels defer sont exécutés dans l'ordre inverse de leur déclaration, c’est-à-dire que le dernier defer déclaré sera le premier exécuté.

__Utilisation typique : Nettoyage de ressources, fermeture de fichiers, déblocage de verrous, etc.__
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Erreur lors de la création du fichier:", err)
        return
    }
    // Déclenche la fermeture du fichier avant la fin de la fonction main
    defer file.Close()

    _, err = file.WriteString("Hello, World!")
    if err != nil {
        fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
        return
    }
    fmt.Println("Écriture réussie.")
}
```
#### 2. panic`

La fonction ``panic`` est utilisée pour interrompre l'exécution normale d'une fonction et commencer à remonter la pile des appels. Lorsqu'une fonction appelle panic, elle termine immédiatement son exécution et tous les defer dans cette fonction sont exécutés avant que la panique ne se propage.

__Utilisation typique : Erreurs graves et irrécupérables, comme une tentative d'accès à un index de tableau hors limites.__

```go
package main

import "fmt"

func divide(a, b int) int {
    if b == 0 {
        panic("division par zéro")
    }
    return a / b
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Récupéré de la panique :", r)
        }
    }()
    
    fmt.Println(divide(4, 2)) // Fonctionne correctement
    fmt.Println(divide(4, 0)) // Provoque une panique
}
```

#### 3. recover

La fonction ``recover`` est utilisée pour intercepter une panique et reprendre l'exécution normale du programme. Elle doit être appelée dans une fonction différée (defer). Si recover est appelé en dehors d'une fonction defer, il ne récupérera pas la panique.

__Utilisation typique : Gestion des erreurs inattendues et reprise du contrôle après une panique.__

```go
package main

import "fmt"

func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Récupéré de la panique dans riskyOperation:", r)
        }
    }()
    panic("quelque chose a mal tourné")
}

func main() {
    riskyOperation()
    fmt.Println("Programme continu après la panique")
}
```

#### Explications

- __``defer``__ : Assure que des opérations de nettoyage ou de déblocage sont effectuées même si une erreur se produit. Dans l'exemple de gestion de fichiers, defer file.Close() garantit que le fichier sera toujours fermé, même si une erreur survient lors de l'écriture.

- __``panic``__ : Indique une erreur critique que le programme ne peut pas récupérer. Cela interrompt l'exécution normale et provoque une panique, mais tous les appels defer en cours sont exécutés avant que le programme ne se termine.

- __``recover``__ : Permet de récupérer d'une panique et de continuer l'exécution normale du programme. Il est souvent utilisé pour gérer les erreurs inattendues et garantir que le programme ne se termine pas brusquement.

#### Remarques

- __Utilisation judicieuse__ : panic et recover doivent être utilisés avec parcimonie. Ils sont généralement réservés aux erreurs graves qui ne peuvent pas être traitées autrement.

- __Gestion des erreurs__ : Pour des erreurs prévisibles et récupérables, il est préférable d'utiliser les mécanismes d'erreur traditionnels en Go (error), plutôt que de recourir à panic et recover.
Ces mécanismes offrent des outils puissants pour gérer les erreurs et assurer la robustesse et la fiabilité des programmes Go.

### Contrôle des Inputs
#### Validation des Inputs

Valider les inputs est essentiel pour assurer la fiabilité et la sécurité de votre application.

```go
package main

import (
    "unicode"
    "log"
)

func isValidUsername(username string) bool {
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    for _, char := range username {
        if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
            return false
        }
    }
    return true
}

func main() {
    username := "user123"
    if isValidUsername(username) {
        log.Println("Valid username")
    } else {
        log.Println("Invalid username")
    }
}
```

#### Gestion des Inputs Utilisateur

Interagir avec l'utilisateur et gérer les inputs dans un programme Go peut se faire en utilisant bufio et os.

```go
package main

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    log.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    if name == "" {
        log.Println("Name cannot be empty")
        return
    }

    log.Printf("Hello, %s!\n", name)
}
```

En appliquant ces techniques, vous pourrez gérer les erreurs efficacement et assurer que les inputs de votre programme soient valides et sûrs.




