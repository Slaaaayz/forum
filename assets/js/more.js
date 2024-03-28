var moreElements = document.querySelectorAll('.more');

moreElements.forEach(function(element) {
    var btn = element.querySelector('.more-btn');
    var menu = element.querySelector('.more-menu');
    var visible = false;

    function showMenu(e) {
        e.preventDefault();
        if (!visible) {
            visible = true;
            element.classList.add('show-more-menu');
            menu.setAttribute('aria-hidden', false);
            document.addEventListener('click', hideMenu, false);
        }
    }

    function hideMenu(e) {
        if (btn.contains(e.target)) {
            return;
        }
        if (visible) {
            visible = false;
            element.classList.remove('show-more-menu');
            menu.setAttribute('aria-hidden', true);
            document.removeEventListener('mousedown', hideMenu);
        }
    }
    btn.addEventListener('click', showMenu, false);
});

function Copy(button) {
    var commentElement = button.parentNode.parentNode.parentNode.parentNode.parentNode.parentNode.parentNode.querySelector('.reponse .lareponse');
    var commentText = commentElement.innerText;
    var tempTextArea = document.createElement('textarea');
    tempTextArea.value = commentText;
    document.body.appendChild(tempTextArea);
    tempTextArea.select();
    tempTextArea.setSelectionRange(0, 99999);
    document.execCommand('copy');
    document.body.removeChild(tempTextArea);
}

function Edit(button){
    var commentElement = button.closest('.Profil').querySelector('.reponse .lareponse');
    var endedit = button.parentNode.parentNode.parentNode.parentNode.parentNode.parentNode.querySelector('.endedit');
    endedit.style.display = "block"
    commentElement.contentEditable = "true";
}
function Endedit(button){
    var commentElement = button.parentNode.parentNode.parentNode.querySelector('.reponse .lareponse');
    button.style.display = "none"
    commentElement.contentEditable = "false";
    var newmess = commentElement.innerText
    console.log(document.location.href.split("/")[5])
    console.log(document.location.href)
    url = '/faq/question/' + document.location.href.split("/")[5]
    var data = {
        type:"edit",
        mess: newmess,
        id : button.value,
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

function Signaler(button){
    url = '/faq/question/' + document.location.href.split("/")[5]
    var data = {
        type:"signaler",
        mess: "",
        id : button.value,
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