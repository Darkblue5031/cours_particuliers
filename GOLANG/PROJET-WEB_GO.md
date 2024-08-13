# Exercice Pratique : Liste de Courses en Web App avec Go

## Objectif
Créer une application web pour gérer une liste de courses. L'application doit permettre aux utilisateurs de s'inscrire, se connecter, et gérer des listes de courses. Les utilisateurs peuvent créer des listes privées ou partagées et définir les permissions d'accès. Les données doivent être persistées dans une base de données SQLite.

## Consignes

### 0. Rappel Bonnes Pratiques

- Le projet doit être organisé avec une architecture claire de dossiers et de fichiers.

- Le code doit comprendre des **docstrings**.

- Utilisez le contrôle de version Git avec une branche par fonctionnalité et une branche de développement.

- Configurez un fichier `.gitignore` approprié.

### 1. Initialisation du Projet

- le nom doit être : __`pro.web_golang(1)`__.

- il doit y avoir une branche par partie ainsi qu'une branche de développement.

- à chaque passage à la suivante, un merge sur la branche de développement doit être effectué.

- la branche main doit donc avoir moins de __`5`__ commits. 

- Configurez un fichier `.gitignore` pour exclure les fichiers inutiles comme les fichiers de configuration d'environnement, les dossiers temporaires, et les fichiers de base de données.

- Ajoutez un fichier `README.md` pour documenter le projet.

### 2. Structure des Dossiers

- `handlers/`: pour les gestionnaires de routes.
- `models/`: pour les définitions des structures de données.
- `static/`: pour les fichiers statiques (CSS, JS).
- `templates/`: pour les templates HTML.
- `database/`: pour la gestion de la base de données.

### 3. Backend avec Go
- __Les imports commençant par ``github.com/`` sont interdits__

- Utilisation de `net/http` pour gérer les requêtes HTTP.

- Utilisation de `encoding/json` pour la sérialisation/désérialisation JSON.

- Utilisation de SQLite pour la persistance des données.

- Gestion de différents utilisateurs et de listes de courses privées et partagées.

### 4. Frontend avec HTML/CSS/JS

- Création de l'interface utilisateur en HTML/CSS.

- Utilisation de JavaScript pour des interactions dynamiques, et des options de couleurs, de thèmes.

### 5. Sauvegarde et Chargement des Données

- Sauvegarde des données dans une base de données SQLite.

- Gestion des erreurs lors de la création, de l'accès, ou de l'utilisation de la base de données.

### 6. Gestion des Utilisateurs

- Implémentation d'une page de profil où l'utilisateur peut supprimer
 son compte ou modifier ses données.

- Les utilisateurs peuvent créer des listes de courses et définir qui peut y accéder (privé, partagé avec lecture seule, ou partagé avec modification).

- Possibilité de créer des adminitrateurs.