# Ytracker - Documentation

Ce projet nommé "Ytracker, votre solution de traçage de célébrités." est le projet Groupie Tracker de Luka GARCIA, Maxence GILLES et Éric MAGNE.

## Sommaire

- [I. A propos du projet](#i-a-propos-du-projet)
- [II. Utilisation de Ytracker](#ii-utilisation-de-ytracker)
- [III. Réalisation des consignes & bonus](#iii-réalisation-des-consignes-bonus)
    - [A. Consignes requises](#a-consignes-requises)
    - [B. Bonus 1 : search-bar](#b-bonus-1-search-bar)
    - [C. Bonus 2 : filters](#c-bonus-2-filters)
    - [D. Bonus 3 : vizualisations](#d-bonus-3-vizualisations)
- [IV. Crédits](#iv-crédits)


## I. A propos du projet

Le projet Groupie Tracker est un projet dont le but est de créer une interface web vous permettant de suivre différents artistes et groupes en utilisant une [API](https://groupietrackers.herokuapp.com/api).

## II. Utilisation de Ytracker

Pour lancer le projet, vous avez seulement besoin de cloner ou télécharger ce repository, ouvrir un terminal et vous diriger dans le dossier contenant le projet.
Ici, vous n'avez plus qu'à exécuter la commande `go run server/main.go`, et utiliser votre navigateur pour vous rentre sur le lien suivant : [http://localhost:8080](http://localhost:8080).

## III. Réalisation des consignes & bonus

### A. Consignes requises

La consigne principale était de créer une interface web avec le langage de programmation Golang en utilisant une API.

---

### B. Bonus 1 : search-bar

La barre de recherche présente sur le site vous permet de rechercher vos artistes favoris afin de les retrouver plus facilement.

Vous pouvez rechercher par : Nom de l'artiste, Nom du groupe, Nom d'un des membres du groupe.

**À noter :** votre recherche n'est pas sensible à la casse, vous pouvez tout aussi bien chercher `Queen` ou `queen`.

---

### C. Bonus 2 : filters

Les filtres sur la pages contenant tous les artistes vous permettent de retrouver facilement vos artistes préférés en utilisant des caractéristiques qui leurs sont propres.

Vous pouvez trier les artistes par : Nombre de membres, Date de création, Date du premier album.

**À noter :** le changement des artistes présents sur la page change de manière dynamique ! Vous n'avez pas besoin de recharger la page.

---

### D. Bonus 3 : vizualisations

Une présentation moderne et contemporaine, avec une retour utilisateur optimal est un élément essentiel. C'est pourquoi notre projet respecte du mieux possible les 8 Règles d'Or de l'Interface Graphique de Schneiderman.

## IV. Crédits

Le projet Ytracker a été conçu par Luka GARCIA, Maxence GILLES et Éric MAGNE (avec l'aide précieuse des mentors qui nous ont accompagnés).