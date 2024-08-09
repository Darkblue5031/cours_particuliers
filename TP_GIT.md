
# Exercice Pratique : Maîtriser Git

Pour vérifier votre compréhension de Git, suivez les étapes ci-dessous. Cet exercice vous permettra de pratiquer les commandes de base de Git, y compris la création de dépôts, les commits, les branches, les pull requests, et l'utilisation du fichier `.gitignore`.

## Étapes de l'exercice

### 1. Créez un Nouveau Projet

1. Créez un nouveau répertoire pour votre projet :
    ```sh
    mkdir exercice-git
    cd exercice-git
    ```
2. Initialisez un dépôt Git :
    ```sh
    git init
    ```

### 2. Ajoutez des Fichiers

1. Créez un fichier `README.md` et ajoutez-y une description de votre projet :
    ```sh
    echo "# Exercice Git" > README.md
    ```
2. Ajoutez ce fichier à l'index de Git et effectuez un commit :
    ```sh
    git add README.md
    git commit -m "Ajout du fichier README.md"
    ```

### 3. Utilisez le Fichier `.gitignore`

1. Créez un fichier `.gitignore` pour exclure certains fichiers du dépôt. Ajoutez les lignes suivantes :
    ```sh
    echo "node_modules/" > .gitignore
    echo "*.log" >> .gitignore
    ```

__Explique ce que vont faire ces lignes dans le `.gitignore`__

2. Ajoutez le fichier `.gitignore` à l'index de Git et effectuez un commit :
    ```sh
    git add .gitignore
    git commit -m "Ajout du fichier .gitignore"
    ```

### 4. Créez une Nouvelle Branche

1. Créez une nouvelle branche nommée `develop` :
    ```sh
    git branch develop
    ```
2. Basculez sur la branche `develop` :
    ```sh
    git checkout develop
    ```

### 5. Modifiez et Commitez

1. Créez un fichier `app.js` et ajoutez-y du code JavaScript de base :
    ```sh
    echo "console.log('Hello, Git!');" > app.js
    ```
2. Ajoutez ce fichier à l'index de Git et effectuez un commit :
    ```sh
    git add app.js
    git commit -m "Ajout du fichier app.js"
    ```

### 6. Stashezz et Récupérez les Modifications

1. Modifiez le fichier `README.md` en ajoutant une ligne supplémentaire.
2. Utilisez la commande `stash` pour sauvegarder vos modifications sans les commiter :
    ```sh
    git stash
    ```
3. Vérifiez que vos modifications ont été stashezzées et que votre branche `develop` est propre.
4. Récupérez vos modifications avec la commande `stash pop` :
    ```sh
    git stash pop
    ```

### 7. Fusionnez les Branches

1. Basculez sur la branche `main` :
    ```sh
    git checkout main
    ```
2. Fusionnez les modifications de la branche `develop` :
    ```sh
    git merge develop
    ```

### 8. Poussez les Modifications vers GitHub

1. Créez un nouveau dépôt sur GitHub (ne pas initialiser avec un `README.md`).
2. Ajoutez le dépôt distant à votre dépôt local :
    ```sh
    git remote add origin https://github.com/votre-utilisateur/exercice-git.git
    ```
3. Poussez vos modifications vers GitHub :
    ```sh
    git push -u origin main
    ```

### 9. Créez une Pull Request

1. Allez sur GitHub et créez une nouvelle branche `feature-1` à partir de `main`.
2. Cloner cette branche dans votre dépôt local :
    ```sh
    git checkout -b feature-1
    ```
3. Modifiez le fichier `app.js` en ajoutant une nouvelle fonctionnalité.
4. Ajoutez et commitez vos modifications :
    ```sh
    git add app.js
    git commit -m "Ajout de la nouvelle fonctionnalité dans app.js"
    ```
5. Poussez la branche `feature-1` vers GitHub :
    ```sh
    git push origin feature-1
    ```
6. Sur GitHub, créez une Pull Request pour fusionner `feature-1` dans `main`.

### 10. Révision et Fusion de la Pull Request

1. Consultez la Pull Request sur GitHub, passez en revue les modifications et fusionnez-la dans `main`.
2. Supprimez la branche `feature-1` après la fusion.

## Conclusion

Cet exercice couvre les bases de Git, y compris la création de dépôts, les commits, les branches, les pull requests, et l'utilisation de `.gitignore`. En complétant cet exercice, vous aurez une bonne compréhension de Git et serez prêt à l'utiliser dans vos projets de développement.
