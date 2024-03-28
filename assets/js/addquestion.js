var text = document.getElementById("textField");
text.style.display = "none";

function displayText() {
    text.style.display = "block";
    var placeholder = document.querySelector(".buttonquestion");
    placeholder.style.display = "none";
}
