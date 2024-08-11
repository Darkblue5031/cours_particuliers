
# Exercice Pratique : Liste de Courses en CLI

## Objectif
Créer une application en ligne de commande (CLI) pour gérer une liste de courses. L'application doit permettre d'ajouter, afficher, rayer, supprimer, et réorganiser les éléments de la liste. Les éléments doivent être sauvegardés dans un fichier JSON pour persistance.

## Étape 1 : Initialisation du Projet

1. **Création du Projet et du Dépôt Git :**
   - Créez un nouveau dossier pour le projet, nommé `pro.cli_golang`.
   - Initialisez un dépôt Git dans ce dossier : `git init`.
   - Créez les branches nécessaires : 
     - `main`
     - `dev`
     - `partie2_ajout_affichage`
     - `partie3_rayer_supprimer`
     - `partie4_reorganisation`
     - `partie5_sauvegarde_chargement`
   - Assurez-vous de bien configurer un fichier `.gitignore` pour ignorer les fichiers inutiles comme les fichiers de configuration, les exécutables, etc.

2. **Architecture du Projet :**
   - Créez un dossier `cmd` pour le code principal.
   - Créez un dossier `pkg` pour les packages et les bibliothèques utilitaires.
   - Ajoutez un dossier `data` pour le fichier JSON qui stockera la liste de courses.
   - Un fichier `main.go` sera à la racine pour lancer l'application.

3. **Documentation :**
   - Assurez-vous que chaque fichier et fonction contiennent des `docstrings` pour expliquer leur rôle.

## Étape 2 : Ajout et Affichage des Éléments

1. **Ajout des Éléments :**
   - Créez une fonction `AddItem` dans un fichier `list.go` dans le dossier `pkg`.
   - Cette fonction doit ajouter un élément à une structure de données de type `slice` (une liste).
   - Le format pourrait ressembler à : 
     ```go
     type ShoppingList struct {
         Items []string
     }
     ```
   - Le fichier JSON (par exemple `shopping_list.json`) sera utilisé pour la persistance.

2. **Affichage des Éléments :**
   - Créez une fonction `ShowItems` pour afficher les éléments numérotés de la liste.
   - Intégrez cette fonction dans le fichier `main.go` pour appeler les différentes fonctionnalités.

## Étape 3 : Rayer ou Supprimer des Éléments

1. **Rayage des Éléments :**
   - Créez une fonction `MarkItem` qui marque un élément comme "fait" ou "rayé". 
   - Pour cela, vous pouvez utiliser une autre structure comme une map qui stocke un booléen pour chaque élément.

2. **Suppression des Éléments :**
   - Créez une fonction `DeleteItem` qui supprime un élément de la liste. 
   - Cette fonction doit mettre à jour la structure de données et sauvegarder la liste mise à jour dans le fichier JSON.

## Étape 4 : Réorganisation de la Liste

1. **Réorganisation :**
   - Implémentez une fonction `ReorderList` qui permet de modifier l'ordre des éléments.
   - Les utilisateurs pourraient, par exemple, échanger deux éléments ou réorganiser la liste manuellement.

2. **Pagination :**
   - Implémentez une option pour afficher la liste par pages, par exemple, 5 éléments à la fois.
   - Vous pouvez ajouter des options pour naviguer entre les pages.

## Étape 5 : Sauvegarde et Chargement des Éléments

1. **Sauvegarde Automatique :**
   - Créez une fonction `SaveToFile` qui enregistre la liste dans un fichier JSON.
   - Cette fonction sera appelée chaque fois qu'une modification est apportée à la liste.

2. **Chargement Automatique :**
   - Créez une fonction `LoadFromFile` pour charger la liste à partir du fichier JSON au démarrage de l'application.
   - Gérez les erreurs potentielles comme un fichier manquant ou corrompu.

## Étape 6 : Tests et Documentation

1. **Tests Unitaires :**
   - Créez un dossier `tests` pour vos tests.
   - Écrivez des tests unitaires pour chaque fonction, en utilisant des cas de test simples et complexes.
   - Assurez-vous que les tests couvrent la plupart des scénarios.

2. **Documenter le Projet :**
   - Créez un `README.md` pour expliquer comment utiliser l'application.
   - Fournissez des exemples de commandes, une description de chaque fonctionnalité, et comment exécuter les tests.

## Conseils pour un Étudiant

1. **Commencez Simplement :** Ne vous précipitez pas. Commencez par les fonctionnalités de base (ajout et affichage) avant d'aborder les fonctionnalités plus complexes.

2. **Lisez la Documentation :** Consultez les documentations de Go et des bibliothèques que vous utilisez, comme `encoding/json` pour manipuler les fichiers JSON.

3. **Testez Fréquemment :** Chaque nouvelle fonctionnalité que vous ajoutez doit être testée pour vérifier qu'elle fonctionne bien avec les fonctionnalités existantes.

4. **Utilisez Git de Manière Intensive :** Commitez souvent, utilisez des branches pour organiser votre travail, et n'hésitez pas à créer des branches temporaires pour expérimenter.

5. **Posez des Questions :** Si quelque chose n'est pas clair, cherchez des réponses dans la documentation ou demandez à vos pairs/enseignants.
