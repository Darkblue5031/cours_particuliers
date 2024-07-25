# Comment Rédiger du Code Proprement

## Introduction
Rédiger du code proprement est essentiel pour tout développeur. Un code propre est plus facile à lire, à maintenir et à déboguer. Il permet aussi une collaboration plus efficace avec d'autres développeurs. Voici quelques conseils pour rédiger du code proprement.

## Utiliser des Noms Significatifs
- **Variables**: Choisissez des noms de variables clairs et significatifs. Par exemple, `nombre_etudiants` est plus explicite que `n`.

- **Fonctions**: Les noms de fonctions doivent indiquer leur but. Par exemple, `calculerMoyenne` est plus clair que `calc`.

## Types de Nommage

- **CamelCase**: Utilisé pour les noms de classes. Par exemple, `MaClasse`. (Peut être utilisé par les fonctions comme ci-dessus.)

- **snake_case**: Utilisé pour les noms de variables (et de fonctions). Par exemple, `ma_variable`.

- **CONSTANTE**: Utilisé pour les constantes, c'est-à-dire des variables dont la valeur ne change pas. Par exemple, `TAILLE_MAX`.

__Dans un projet, toutes les variables doivent respecter la même norme de nommage ; il en est de même pour les fonctions. Les objets et les constantes (ou variables utilisés comme des constantes) doivent TOUJOURS respecter leur norme de nommage.__

## Garder les Fonctions Simples
- Une fonction doit réaliser une seule tâche spécifique. Si une fonction fait trop de choses, divisez-la en plusieurs fonctions plus petites.

- Essayez de ne pas dépasser 20 à 30 lignes de code par fonction.

## Commentaires et Docstrings

### Intérêt des Docstrings

Les docstrings sont des chaînes de caractères utilisées pour documenter les modules, classes et fonctions en Python. Elles sont essentielles pour plusieurs raisons :

- **Documentation**: Elles fournissent une description claire de ce que fait le code, facilitant la compréhension pour d'autres développeurs (ou pour vous-même dans le futur).

- **Génération Automatique**: Elles permettent la génération automatique de documentation à l'aide d'outils comme Sphinx.

- **Outils de Développement**: Les IDE et autres outils peuvent utiliser les docstrings pour fournir des informations contextuelles pendant le codage.

### Fonctionnement des Docstrings

Une docstring est placée juste après la définition d'une fonction, d'une classe ou d'un module. Voici un exemple de docstring pour une fonction :

```python
def ajouter(a, b):
    """
    Ajoute deux nombres et renvoie le résultat.

    Args:
        a (int): Le premier nombre.
        b (int): Le deuxième nombre.

    Returns:
        int: La somme des deux nombres.
    """
    return a + b
```

Il est aussi courant de trouver en première ligne d'un fichier une docstring informant de son contenu.

#### Bonnes Pratiques pour les Docstrings

- __Courtes et Précises__: Décrivez brièvement ce que fait la fonction, ses arguments et ce qu'elle renvoie.

- __Style__: Utilisez des conventions de style comme celles de PEP 257 pour les docstrings.

- __Mise à Jour__: Assurez-vous de mettre à jour les docstrings si vous modifiez le code.

### Organisation du Code

- __Modules et Packages__: Divisez votre code en modules et packages pour mieux l'organiser. Chaque module doit contenir des fonctions et classes liées entre elles.

- __Structure de Répertoire__: Adoptez une structure de répertoire logique pour votre projet. Par exemple :

### PEP 20: The Zen (of Python)
La PEP 20, également connue sous le nom de The Zen of Python, est une liste de 19 aphorismes qui capturent l'esprit de la philosophie de Python. Voici quelques points clés :

- __Beau est mieux que laid.__: Un code lisible et élégant est préféré.

- __Simple est mieux que complexe.__: Préférez la simplicité à la complexité.

- __Lisible compte.__: La lisibilité du code est primordiale.

- __Il devrait y avoir une - et de préférence une seule - manière évidente de le faire.__: Python privilégie une approche unifiée et cohérente pour résoudre les problèmes.

Pour voir tous les aphorismes de la PEP 20, vous pouvez simplement exécuter __``import this``__ dans un interpréteur Python.

#### D'autres outils dédiés à python comme __``pylint``__ ou __``black``__ servent à améliorer la qualité des fichiers de code.

### Conclusion

Rédiger du code propre est une compétence essentielle pour tout développeur. En suivant les conseils ci-dessus, vous rendrez votre code plus lisible, maintenable et professionnel. N'oubliez pas d'utiliser les docstrings pour documenter votre code efficacement, de respecter les conventions de nommage, et de maintenir une organisation cohérente dans vos projets. La PEP 20 peut également servir de guide pour adopter les meilleures pratiques de codage en Python (entre autres).

Bon codage !