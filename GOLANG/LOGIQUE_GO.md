### Exercice de logique Golang

#### L'exercice doit être rendu sur un répertoire github : 

- le nom doit être : __`ex.logique_golang(1)`__.

- il doit y avoir une branche par partie ainsi qu'une branche de développement.

- un __`.gitignore`__ doit être correctement configuré.

- à chaque passage à la suivante, un merge sur la branche de développement doit être effectué.

- la branche main doit donc avoir moins de __`5`__ commits. 


#### Partie I :
Créer la fonction, __`print_diammond()`__ qui print par défaut le motif suivant :

```
 *****
*******
 *****
  ***
   *
```

#### Partie II :

Ajouter des arguments qui permettent d'écrire __`go run logique.go --size 4 --motif #`__

```
 ###
#####
 ###
  #
 ```

__`--size`__ doit être remplaçable par __`-s`__
__`--motif`__ doit être remplaçable par __`-m`__

#### Partie III :
- La fonction doit, pour une taille à 3 retourner :

```
*****
 ***
  *
 ```

 - La fonction doit, pour une taille à 2 retourner :

```
 ***
  *
 ```

 - Pour toute taille inférieure, la fonction devra afficher un message adapté.

#### Partie IV :

Il est également demandé que si __`go run logique.go --help`__ ou __`-h`__ est utilisé, un "usage" doit être proposé.

Trouve au moins deux solutions afin d'encadrer l'usage de la fonction, tu devras justifier tes choix en commentaire.

__N'oublie pas le `README.md` !__
