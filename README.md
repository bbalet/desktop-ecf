# Application de bureau Cinéphoria

Vous devez installer Go et les dépendances du projet (pour Linux) :

    $ sudo apt-get update && sudo apt-get install gcc libgl1-mesa-dev libegl1-mesa-dev libgles2-mesa-dev libx11-dev xorg-dev bc
    $ sudo apt install go
    $ cd /chemin/vers/mon/projet
    $ go mod tidy
    $ go build

Notez que Windows n'a besoin d'aucune dépendance supplémentaire.

Construire localement :

    $ go build

Construire pour une autre plateforme (par exemple Windows, Mac, Android...) avec les outils multiplateformes de Fyne (cf. https://docs.fyne.io/started/cross-compiling.html).
