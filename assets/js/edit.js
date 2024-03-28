const paragraph = document.getElementById("com");
const edit_button = document.getElementById("edit");
const end_button = document.getElementById("endedit");
end_button.style.display = "none"

edit_button.addEventListener("click", function() {
end_button.style.display = "block"
  paragraph.contentEditable = true;
} );

end_button.addEventListener("click", function() {
  paragraph.contentEditable = false;
  end_button.style.display = "none"
} )

