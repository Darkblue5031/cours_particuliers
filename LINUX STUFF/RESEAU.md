# Sujet de rattrapage réseau

Dans ce TP, vous allez revoir plusieurs notions de réseau que vous avez étudiées pendant le module réseau.

Le rendu attendu se fera au travers d'un repository **gitea** que vous appellerez ```reseau-mentor```.

Pour ce qui est des rendus, vous devrez rendre un fichier markdown (```README.md```) qui contiendra les réponses aux questions posées.

Tout rendu ne respectant pas ces consignes ou remis en retard ne sera pas corrigé.

Lorsque vous voyez un emoji 🌞, cela signifie qu'il faut répondre à une question et lorsqu'il s'agit d'un 🦈 c'est qu'il faut encoyer un PCAP.

## Sommaire

[Pré-requis](#sommaire)

[Partie 1 : Setup du patron](#partie-1--setup-du-patron)

[Partie 2 : Premier test réseau](#partie-2--premier-test-réseau)

[Partie 3 : Un petit peu de routage](#partie-3--un-petit-peu-de-routage)

[Partie 4 : Web](#partie-4--web)

[Partie 5 : DNS](#partie-5--dns)

## Pré-requis

- Oracle VirtualBox
- Rocky Linux minimal
- Connexion internet

## Partie 1 : Setup du patron

Pour cette partie, vous allez devoir créer une VM Rocky Linux minimal. Vous allez devoir configurer la VM avec 2 interfaces réseaux : une en NAT et l'autre en Host-Only. Vous allez devoir donner un hostname à la machine et configurer SSH pour administrer la machine.

Cette VM sera votre patron pour le reste du TP, faites bien attention à la configurer correctement.

🌞 Créez une VM avec VirtualBox, nommée ```Patron```.

🌞 Configurez la VM avec 2 interfaces réseaux. La première en NAT et la deuxième en Host-Only.

🌞 Installez Rocky Linux minimal sur la VM.

🌞 Définissez une IP statique à la VM.

🌞 Donnez un hostname à la machine.

🌞 Configurez SSH pour administrer la machine (avec échange de clés).

## Partie 2 : Premier test réseau

Pour cette partie, vous allez devoir tester la connectivité entre deux machines.

Liste des machines :

| Machine     | LAN `10.1.1.0/24` |
| --------- | -------------- |
| `etudiant` | `10.1.1.10`    |
| `mentor`   | `10.1.1.11`    |

🌞 Configurez correctement ces 2 VM.
- Donnez un hostname à chaque machine.
- Fournissez un cat des fichiers de configuration réseau.
- Fournissez un cat des fichiers de configuration hostname.

🌞 Testez la connectivité entre les deux machines avec un ping.

## Partie 3 : Un petit peu de routage

À présent, nous allons créer un deuxième réseau dans lequel nous allons ajouter une machine qui devra communiquer avec le premier réseau. Pour cela, vous allez devoir configurer un routeur qui fera le lien entre les deux réseaux. Vous pouvez réutiliser les VM précédentes.

Supprimez la carte réseau Host-Only sur les deux VM précédentes (seul le routeur doit avoir une carte Host-Only).

Liste des machines :

| Machine  | LAN 1 `10.1.1.0/24` | LAN 2 `10.1.2.0/24` |
| -------- | ------------------- | ------------------- |
| `routeur` | `10.1.1.254`        | `10.1.2.254`        |
| `mentor`  | `10.1.1.11`         | no                  |
| `etudiant` | no                 | `10.1.2.10`         |

🌞 Configurez correctement ces 3 VM.

🌞 Configurez le routeur pour qu'il puisse faire le lien entre les deux réseaux.

🌞 Testez la connectivité entre les machines :
- `etudiant` et `routeur`
- `mentor` et `routeur`
- `etudiant` et `mentor`

🦈 Faites des captures avec WireShark pour prouver que les machines communiquent bien entre elles.

🌞 Donnez à `etudiant` et `mentor` un accès à internet via le routeur.

## Partie 4 : Web

Pour cette partie, vous allez devoir installer un serveur web sur la machine `mentor`.

Nous allons installer un serveur web Apache, (cela ressemble à Nginx donc ne vous inquiétez pas, vous ne serez pas trop perdu).

🌞 Installez un serveur web Apache sur la machine `mentor`.

🌞 Créez un fichier HTML qui va afficher le nom de la machine (à vous de trouver où le mettre).

🌞 Créez un nouvel utilisateur nommé `apache` qui va gérer le serveur web (donnez-lui les accès aux dossiers/fichiers dont il a besoin).

🌞 Ouvrez le port firewall.

🌞 Testez l'accès au serveur web depuis la machine `mentor` à l'aide d'une commande curl.

🌞 Testez l'accès au serveur web depuis la machine `etudiant` à l'aide d'une commande curl.