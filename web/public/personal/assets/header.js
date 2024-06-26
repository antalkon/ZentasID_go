let progileOpen = false
let messageOpen = false
const profileWindow = document.getElementById("profileWindow")
const messageWindow = document.getElementById("MessageWindow")

const profile = document.getElementById("profile")
const message = document.getElementById("message")

const OpenImg = document.getElementById("OpenImg")
const MessageImg = document.getElementById("messageImg")

function handleProfileClick() {
    if (progileOpen === false) {
            profileWindow.style.display = "flex"
            profile.style.zIndex = "101"
            OpenImg.src = "./assets/close.png"
            progileOpen = true
    }
    else if (progileOpen === true) {
        profileWindow.style.display = "none"
        profile.style.zIndex = "2"
        OpenImg.src = "./assets/down.png"
        progileOpen = false
    }
}


function handleNotifyClick(){
    if (messageOpen === false) {
        messageWindow.style.display = "flex"
        message.style.zIndex = "101"
        MessageImg.src = "./assets/close.png"
        messageOpen = true

    }
    else if (messageOpen === true) {
        messageWindow.style.display = "none"
        message.style.zIndex = "2"
        MessageImg.src = "./assets/business.png"
        messageOpen = false
    }
}
// Найти кнопку по ID и добавить обработчик события
profile.addEventListener('click', handleProfileClick);
message.addEventListener('click', handleNotifyClick);
