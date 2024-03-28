let tags
var index = 0
fetch('/createFAQ', {
    headers: {
      'X-Requested-With': 'XMLHttpRequest'
    }
  })
    .then(response => response.json())
    .then(data => {
      tags = data
      tags.sort(compareByNbUsed);
    })
    .catch(error => console.error('Erreur lors de la récupération des tags:', error));
  
var tag1 = document.getElementById("tag1")
var tag2 = document.getElementById("tag2")
var tag3 = document.getElementById("tag3")
var tag4 = document.getElementById("tag4")
var tag5 = document.getElementById("tag5")
var submit = document.getElementById("submittag")
function compareByNbUsed(tag1, tag2) {
    return tag2.NbUsed - tag1.NbUsed;
}

var IndexTag = 1
    
    
window.addEventListener("load", function(){
        console.log("chargement")
        tag2.style.display="none"
        tag3.style.display="none"
        tag4.style.display="none"
        tag5.style.display="none"
    });
    
function Submittag(){
        switch(IndexTag){
            case 1:
                console.log("tag1",tag1.value)
                if(tag1.value!=""){
                tag2.style.display="block"
                IndexTag++
                }
                break
            case 2:
                if(tag2.value!=""){
                tag3.style.display="block"
                IndexTag++
                }
                break
            case 3:
                if(tag3.value!=""){
                tag4.style.display="block"
                IndexTag++
                }
                break
            case 4:
                if(tag4.value!=""){
                tag5.style.display="block"
                submit.style.display="none"
                IndexTag++
                }
                break
            default:
                break
        }
    }
    



function autocompleteFunction(input) {
    var autocompleteList = document.getElementById("autocomplete-list");
    var options = autocompleteList.options;

    if (input.value.length > 0) {
        autocompleteList.style.display = "block";
        options.length = 0;
        console.log(tags)
        for(var i = 0 ; i < tags.length;i++){
            console.log("tag name : ",tags[i].Name)
            console.log(input.value)
            if (tags[i].Name.includes(input.value)){
                var option = new Option(tags[i].Name+"("+tags[i].NbUsed+")",tags[i].Name);
                options.add(option);
            }
        }
    } else {
        autocompleteList.style.display = "none";
    }
}
function selectOption() {
  var input = document.getElementsByClassName("myInput")[IndexTag-1];
  console.log("input : ",input)
  console.log("index : ",IndexTag)
  console.log("value : ",input.value)
  var autocompleteList = document.getElementById("autocomplete-list");
  var selectedValue = autocompleteList.value;
  console.log("new value ",selectedValue)
  input.value = selectedValue;
  autocompleteList.style.display = "none";
}


function compareTagsByNbUsed(a, b) {
    if (a.NbUsed < b.NbUsed) {
        return 1; 
    } else if (a.NbUsed > b.NbUsed) {
        return -1;
    } else {
        return 0;
    }
}
