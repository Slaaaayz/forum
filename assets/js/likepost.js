function Like(element) {
    element.style.display = "none";
    var sibling = element.nextElementSibling;
    sibling.style.display = "block";
    console.log(sibling,"liked")
    console.log(sibling.dataset.value,"liked")
    var nblike = document.getElementById("nblike"+sibling.dataset.value)
    value = parseInt(nblike.innerText)
    nblike.innerText = value + 1
    url = '/forum/topic/post/' + document.location.href.split("/")[6]
    var data = {
        type:"likePost",
        mess: "",
        id : sibling.dataset.value,
    }
    console.log(url)
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        // Gérer la réponse du serveur si nécessaire
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi de la requête:', error);
    });

}

function UnLike(element) {
    element.style.display = "none";
    var sibling = element.previousElementSibling;
    sibling.style.display = "block";
    console.log(element,"liked")
    console.log(element.dataset.value,"liked")
    var nblike = document.getElementById("nblike"+sibling.dataset.value)
    value = parseInt(nblike.innerText)
    nblike.innerText = value - 1
    url = '/forum/topic/post/' + document.location.href.split("/")[6]
    var data = {
        type:"dislikePost",
        mess: "",
        id : sibling.dataset.value,
    }
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        // Gérer la réponse du serveur si nécessaire
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi de la requête:', error);
    });
}
