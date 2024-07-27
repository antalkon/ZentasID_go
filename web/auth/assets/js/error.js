function errorVisible(err){
    document.getElementById("errAlert").style.display = "flex"
    document.getElementById("errTxt").textContent = err

    setTimeout(() => {
        document.getElementById("errAlert").style.display = "none"
    }, 5000); // Ждем 3 секунды
}