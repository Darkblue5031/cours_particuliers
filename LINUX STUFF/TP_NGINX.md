
# Exercice : Utilisation et Configuration de Nginx

**Objectif :** Cet exercice vise à vous familiariser avec l'installation, la configuration et les bonnes pratiques d'utilisation de Nginx. Vous apprendrez à utiliser des fichiers de configuration additionnels, le dossier __`sites-enabled`__, et le répertoire __`conf.d/`__ pour éviter la duplication des lignes de configuration. La configuration se fera par adresse IP (pas de nom de domaine) et inclura plusieurs pages ainsi que des pages d'erreur personnalisées.

## Prérequis

- Une machine virtuelle ou un serveur Linux avec Nginx installé (ex: debian).
- Droits sudo.

## Étape 0 : Création des Pages pour les Sites

Créez une page index pour trois sites dans __``/var/www/``__.

Faire de même pour les pages d'erreur, mais elles seront communes pour les deux sites.

## Étape 1 : Installation de Nginx

Si Nginx n'est pas déjà installé, commencez par l'installer.

## Étape 2 : Création des Fichiers de Configuration des Sites

Créez un fichier de configuration par site dans __`/etc/nginx/sites-available/`__ :


## Étape 3 : Activer les Configurations de Sites

Activez les configurations des deux premiers sites en créant des liens symboliques vers les fichiers dans __`sites-enabled`__ :

## Étape 4 : Vérification et Redémarrage de Nginx

Vérifiez la configuration de Nginx :

```bash
sudo nginx -t
```

Si la vérification est réussie, redémarrez Nginx pour appliquer les modifications :

```bash
sudo systemctl restart nginx
```

## Étape 8 : Test des Sites et des Pages d'Erreur

Ouvrez votre navigateur et accédez aux sites. Testez également les pages d'erreur en accédant à des URL inexistantes.

## Bonnes Pratiques

- **Organisation des fichiers :** Utilisez __`sites-available`__ et __`sites-enabled`__ pour une gestion claire et modulaire des configurations.
- **Éviter les redondances :** Utilisez des fichiers de configuration séparés et des inclusions si nécessaire pour éviter de dupliquer des blocs de configuration.
- **Pages d'erreur personnalisées :** Créez des pages d'erreur spécifiques pour offrir une meilleure expérience utilisateur en cas de problème.
- **Réutilisation des configurations communes :** Le répertoire __`conf.d/`__ permet de centraliser les configurations réutilisables pour éviter la duplication.
- **Validation et redémarrage contrôlés :** Validez toujours la configuration avec __`nginx -t`__ avant de redémarrer le service pour éviter les interruptions.

En suivant cet exercice, vous aurez configuré Nginx pour servir plusieurs sites web avec des pages d'erreur personnalisées, tout en respectant les bonnes pratiques pour une gestion propre et efficace de vos configurations.
