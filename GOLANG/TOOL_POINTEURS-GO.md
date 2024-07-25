# Introduction aux Pointeurs en Go

**Objectif** : Comprendre ce qu'est un pointeur, comment les utiliser en Go, et pourquoi ils sont importants.

## 1. Qu'est-ce qu'un Pointeur ?

En programmation, un pointeur est une variable qui stocke l'adresse mémoire d'une autre variable. Au lieu de contenir directement une valeur, un pointeur "pointe" vers l'emplacement en mémoire où cette valeur est stockée.

## 2. Déclaration et Initialisation des Pointeurs

En Go, vous pouvez créer un pointeur en utilisant l'opérateur `*`, et vous pouvez obtenir l'adresse d'une variable avec l'opérateur `&`.

**Déclaration** :  
```go
var p *int // Déclare un pointeur vers un int
```


**Initialisation** :  
```go
var x int = 42
var p *int = &x // p contient l'adresse de x
```

## 3. Déférencement des Pointeurs

Pour accéder à la valeur pointée par un pointeur, vous utilisez l'opérateur `*` (appelé opérateur de déférencement).

```go
fmt.Println(*p) // Affiche la valeur stockée à l'adresse pointée par p
```

## 4. Pointeurs et Fonction

Les pointeurs sont très utiles pour passer des arguments à des fonctions lorsque vous voulez modifier les valeurs originales.

```go
func increment(x *int) {
    *x++ // Déférencement pour accéder et modifier la valeur
}

func main() {
    num := 5
    increment(&num) // Passe l'adresse de num à la fonction
    fmt.Println(num) // Affiche 6
}
```

## 5. Pointeurs et Structures

Les pointeurs sont également utilisés avec les structures pour éviter de copier toute la structure lors du passage à une fonction ou pour maintenir des références partagées.

```go
type Person struct {
    Name string
    Age  int
}

func updateAge(p *Person, newAge int) {
    p.Age = newAge
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    updateAge(&p, 31)
    fmt.Println(p.Age) // Affiche 31
}
```

## 6. Pointeurs Nuls

Un pointeur peut aussi ne pas pointer vers une adresse valide. Il est alors appelé pointeur nul, et il est initialisé par défaut à `nil`.

```go
var p *int
if p == nil {
    fmt.Println("Le pointeur est nul.")
}
```

## 7. Pointeurs et Mémoire

Il est important de comprendre que l'utilisation incorrecte des pointeurs peut mener à des erreurs comme les accès à des adresses invalides, ce qui peut provoquer des comportements indéfinis ou des plantages. Go gère beaucoup de problèmes de mémoire automatiquement, mais il est toujours bon de comprendre comment les pointeurs fonctionnent.

## Résumé

- Déclaration d'un pointeur : `var p *Type`
- Obtenir l'adresse d'une variable : `&variable`
- Déférencement pour accéder à la valeur : `*p`
- Utilisation dans les fonctions pour modifier les valeurs : Passez l'adresse de la variable à la fonction.
