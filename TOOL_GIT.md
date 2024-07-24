
# Introduction à Git pour les Développeurs Débutants

## Qu'est-ce que Git ?

Git est un système de contrôle de version distribué. Il permet aux développeurs de suivre les modifications apportées au code source au fil du temps, de collaborer avec d'autres développeurs et de revenir à des versions antérieures du code si nécessaire.

## Installation de Git

Pour installer Git, suivez les instructions selon votre système d'exploitation :

### Windows

1. Téléchargez Git depuis [git-scm.com](https://git-scm.com/).
2. Exécutez l'installateur et suivez les instructions.

### macOS

1. Ouvrez Terminal.
2. Installez Git en utilisant Homebrew :
    ```sh
    brew install git
    ```

### Linux

1. Ouvrez Terminal.
2. Installez Git en utilisant votre gestionnaire de paquets :
    ```sh
    sudo apt-get install git   # Pour les distributions basées sur Debian/Ubuntu
    sudo yum install git       # Pour les distributions basées sur Fedora/RedHat
    ```

## Configurer Git

__En cas d'utilisation de plusieurs comptes git différents, une configuration plus adaptée est préférable ; mais c'est configuration est recommandée pour le compte Gitea__

Après avoir installé Git, configurez-le avec votre nom et votre adresse email :
```sh
git config --global user.name "Votre Nom"
git config --global user.email "votre.email@example.com"
```

## Créer un Répertoire Git

Pour commencer à utiliser Git, vous devez créer un nouveau dépôt (repository) ou cloner un dépôt existant.

### Initialiser un Nouveau Dépôt

1. Créez un nouveau répertoire pour votre projet :
    ```sh
    mkdir mon-projet
    cd mon-projet
    ```
2. Initialisez Git dans ce répertoire :
    ```sh
    git init
    ```

### Cloner un Dépôt Existant

Pour cloner un dépôt existant depuis GitHub :
```sh
git clone URL
```
__En cas d'utilisation de plusieurs comptes git différents, un clone SSH est favorisé ; mais indisponible sur Gitea__

## Travailler avec Git

### Ajouter des Fichiers au Dépôt

1. Créez un fichier et ajoutez-y du contenu.
2. Ajoutez ce fichier à l'index de Git (stage) :
    ```sh
    git add nom-du-fichier
    ```

### Effectuer un Commit

Un commit enregistre les modifications dans l'historique du dépôt.
```sh
git commit -m "Message de commit"
```

### Pousser les Modifications

Pour envoyer vos modifications à un dépôt distant (par exemple sur GitHub) :
```sh
git push
```

### Récupérer les Modifications

Pour récupérer les modifications depuis un dépôt distant :
```sh
git pull
```

### Gérer les Modifications Temporaires avec Stash

Si vous avez des modifications non enregistrées et que vous devez basculer sur une autre branche, vous pouvez utiliser stash :
```sh
git stash
```

__Cette commande restaure les fichiers locaux à l'était du précedent commit.__

Pour récupérer vos modifications locales :
```sh
git stash pop
```

## Travailler avec des Branches

### Créer et Basculer sur une Nouvelle Branche

1. Créez une nouvelle branche :
    ```sh
    git branch nouvelle-branche
    ```
2. Basculer sur cette branche :
    ```sh
    git checkout nouvelle-branche
    ```

### Fusionner une Branche

__L'utilisation de Pull Request est favorisée__

Pour fusionner les modifications d'une branche dans une autre :
1. Basculer sur la branche principale :
    ```sh
    git checkout main
    ```
2. Fusionner la nouvelle branche :
    ```sh
    git merge nouvelle-branche
    ```

## Faire une Pull Request

Une Pull Request permet de proposer vos modifications à un projet pour revue et inclusion dans la branche principale. Voici comment en créer une sur GitHub :

1. Poussez votre branche sur GitHub
2. Allez sur le dépôt GitHub et cliquez sur "New Pull Request".
3. Sélectionnez votre branche et suivez les instructions pour soumettre votre Pull Request.

## Conclusion

Git est un outil puissant pour gérer votre code et collaborer avec d'autres développeurs. En maîtrisant les bases de Git, vous serez en mesure de gérer efficacement vos projets de développement. N'hésitez pas à explorer les commandes avancées de Git et à consulter la documentation officielle pour approfondir vos connaissances.
