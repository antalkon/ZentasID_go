let progileOpen = false
const profileWindow = document.getElementById("profileWindow")
const OpenImg = document.getElementById("OpenImg")

function handleProfileClick() {
    if (progileOpen === false) {
            profileWindow.style.display = "flex"
            OpenImg.src = "./assets/close.png"
            progileOpen = true

    }
    else if (progileOpen === true) {
        profileWindow.style.display = "none"
        OpenImg.src = "./assets/down.png"
        progileOpen = false




    }
}

// Найти кнопку по ID и добавить обработчик события
document.getElementById('profile').addEventListener('click', handleProfileClick);
