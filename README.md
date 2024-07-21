# Application de bureau Cinéphoria

Cinéphoria Maintenance est une application desktop multiplateformes (Linux, FreeBSD, Windows, Darwin, Android) supportant de multiples architectures.

Vous pouvez télécharger la version de l'application qui correspond à votre OS dans l'onglet Release. Si vous souhaitez une version pour les OS d'Apple, vous devez la compiler vous-même sur une machine Apple.

##  Mode d'emploi

Le mode d'emploi complet est ici : https://github.com/bbalet/web-ecf/blob/main/docs/guide_utilisateur.pdf

En résumé, seul un utilisateur ayant un rôle employé peut utiliser l'application (login : employee@example.org ; mot de passe : employee).

## Exemples

Sur Linux :

![Cinéphoria sur Android](./docs/screenshot-linux.jpg)

Sur Windows :

![Cinéphoria sur Android](./docs/screenshot-windows.jpg)

Et même sur Android :

![Cinéphoria sur Android](./docs/screenshot-android.jpg)

## Construire localement avec Linux

Vous devez installer Go et les dépendances du projet (pour Linux) :

    $ sudo apt-get update && sudo apt-get install gcc libgl1-mesa-dev libegl1-mesa-dev libgles2-mesa-dev libx11-dev xorg-dev bc
    $ sudo apt install go
    $ cd /chemin/vers/mon/projet
    $ go mod tidy
    $ go build

## Cross-compilation

Construire pour une autre plateforme (par exemple Windows, Mac, Android...) avec les outils multiplateformes de Fyne (cf. https://docs.fyne.io/started/cross-compiling.html and https://github.com/fyne-io/fyne-cross).

    $ fyne-cross windows -arch=*
    $ fyne-cross linux -arch=*
    $ fyne-cross darwin -arch=*
    $ fyne-cross freebsd -arch=*
    $ fyne-cross ios
    $ fyne-cross android
