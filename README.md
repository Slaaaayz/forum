# <div align="center">FORUM (Bamboo)

## SOMMAIRE

- [I. Comment installer le Forum](#i-comment-installer-le-Forum)
- [II. Hebergement de Bamboo](#ii-hebergement-de-bamboo)
- [III. Fonctionnement du Forum](#iii-fonctionnement-de-bamboo)
- [IV. Page FAQ](#iv-page-faq)
- [V. Si le Forum ne fonctionne pas](#v-si-le-forum-ne-fonctionne-pas)

## I. Comment installer le Forum

Pour installer le Forum, commencez par cloner le repository sur votre ordinateur en utilisant le terminal et la commande suivante :

```bash
git clone https://ytrack.learn.ynov.com/git/nelio/forum.git
```
Ensuite, lancez le site via les commandes suivantes:

```go
cd forum
god mod init forum
god mod tidy
go run server/main.go
```

Pour accéder au site, ouvrez un navigateur et entrez l'adresse suivante : http://localhost:8080/

Pour run le serveur sous docker :
    
```bash
docker build -t forum:latest .
docker run -p 8080:8080 forum:latest
```

## II. Hebergement de Forum

Nous avons choisi d'héberger Bamboo. Pour accéder à Bamboo, veuillez vous connectez à l'adresse suivante : https://groupe9.etudiants.ynov-bordeaux.com/

## III. Fonctionnement du Forum

Le forum fonctionne comme un espace virtuel où l'utilisateur peut s'inscrire, créer un sujet de discussion et publier des messages. Ces messages peuvent engendrer des réponses et des discussions subséquentes entre les membres. Cela permet à l'utilisateur de partager ses connaissances, d'échanger des idées et de discuter de sujets qui l'intéressent. Le forum est souvent organisé en catégories et sous-catégories thématiques pour faciliter la navigation et la recherche d'informations pertinentes. Il favorise également l'interaction asynchrone, permettant à l'utilisateur de participer aux discussions à son propre rythme, sans contrainte de temps en temps réel.

## IV. Page FAQ

La page FAQ est une page qui contient des questions posées et leurs réponses. Elle est conçue pour aider les utilisateurs à trouver rapidement des informations. La page FAQ peut également inclure des liens vers d'autres ressources utiles, des guides d'utilisation ou des tutoriels pour aider les utilisateurs à résoudre leurs problèmes.

## V. Si le Forum ne fonctionne pas

Si le Forum ne fonctionne pas en local, suivez ces étapes :

1. Vérifiez l'installation de Go sur votre ordinateur.

2. Assurez-vous que le repository du Forum est correctement installé.

3. Ouvrez le dossier Forum sur VsCode.

4. Si le problème persiste, supprimez le dossier Forum et recommencez l'installation.

5. Si aucune de ces solutions ne fonctionne, contactez les auteurs du Forum.

## <div align="right">Les auteurs du Forum

<div align="right">NGUINGNANG Elio
<div align="right">MAREC Aymeric
<div align="right">CHORT Maxime
