const deleteButton = document.getElementById('boutondel');
const deleteModal = document.getElementById('delete-modal');
// deleteButton.addEventListener('click', function() {
//     deleteModal.style.display = 'flex';
// });

function Delete(e){
    deleteModal.style.display = 'flex';
    var id = document.getElementById("getid").value = e.value
    console.log("truc :",e.value)
}

document.getElementById("confirm-delete").addEventListener('click',function() {
    var getid = document.getElementById("getid").value
    url = '/faq/question/' + document.location.href.split("/")[5]
    console.log("url : ",url)
    console.log("idcom : ",getid)
    var data = {
        type:"delete",
        mess: "",
        id : getid,
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
    window.location.reload();
});

document.getElementById("lacroix").addEventListener('click',function() {
    deleteModal.style.display = 'none';
});

document.getElementById('cancel-delete').addEventListener('click', function() {
    deleteModal.style.display = 'none';
});

