var showsign = true
var divSignalement = document.getElementById("signalements")
var divNotifs = document.getElementById("notifs")


function showSignalement(){
    if (showsign == true){
        divSignalement.style.display = "block"
        divNotifs.style.display = "none"
        showsign = false
    }
}
function showNotif(){
    if (showsign == false){
        divNotifs.style.display = "block"
        divSignalement.style.display = "none"
        showsign = true
    }
}