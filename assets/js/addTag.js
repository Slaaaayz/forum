var tag1 = document.getElementById("tag1")
var tag2 = document.getElementById("tag2")
var tag3 = document.getElementById("tag3")
var tag4 = document.getElementById("tag4")
var tag5 = document.getElementById("tag5")
var submit = document.getElementById("submittag")

var IndexTag = 1
console.log("indexdebut : ",IndexTag)


window.addEventListener("load", function(){
    console.log("chargement")
    tag2.style.display="none"
    tag3.style.display="none"
    tag4.style.display="none"
    tag5.style.display="none"
});

function Submittag(){
    
    console.log(IndexTag)
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
