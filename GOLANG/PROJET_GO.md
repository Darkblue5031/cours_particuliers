# Exercice Pratique : Liste de Courses en CLI
## Objectif
Créer une application en ligne de commande (CLI) pour gérer une liste de courses. L'application doit permettre d'ajouter, afficher, rayer, supprimer, et réorganiser les éléments de la liste. Les éléments doivent être sauvegardés dans un fichier JSON pour persistance.

## Consignes

### 0. Rappel Bonnes pratiques

- Le projet doit être organisé par une achitechecture clair de dossiers et de fichiers.

- Le code doit comprendre des __`docstrings`__.

### 1. Initialisation du Projet
- le nom doit être : __`pro.cli_golang(1)`__.

- il doit y avoir une branche par partie ainsi qu'une branche de développement.

- un __`.gitignore`__ doit être correctement configuré.

- à chaque passage à la suivante, un merge sur la branche de développement doit être effectué.

- la branche main doit donc avoir moins de __`5`__ commits. 

### 2. Ajout et Affichage des Éléments

- Implémention d'une fonction pour ajouter des éléments à la liste.

- Implémention d'une fonction pour afficher les éléments de la liste, numérotés.

### 3. Rayer ou Supprimer des Éléments
- Implémentation d'une option pour rayer des éléments de la liste.

- Implémention d'une option pour supprimer des éléments de la liste.

### 4. Réorganisation de la Liste
- Implémention d'une fonction pour modifier l'ordre des éléments dans la liste.

- Ajout de fonctionnalités pour choisir comment afficher la liste :
    - Nombre d'éléments à afficher à la fois.
    - Systèmes de pages pour naviguer à travers la liste.

### 5. Sauvegarde et Chargement des Éléments

- Utilisation d'un fichier JSON pour sauvegarder les éléments de la liste.

- Implémention d'une fonction pour créer automatiquement le fichier JSON s'il n'existe pas.

- Gérez les erreurs potentielles lors de la création, de l'accès, ou de l'utilisation fichier JSON.

### 6. Bonus !

Réaliser des tests pour vérifier le bon fonctionnement du projet.

__N'oublie pas le `README.md` !__