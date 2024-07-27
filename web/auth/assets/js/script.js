const phoneBtn = document.getElementById('phone')
const emailBtn = document.getElementById('mail')
const createAcc = document.getElementById('createAcc')
const frame = document.getElementById('frameLogin')
const mainl = document.getElementById('lMain')
const lWind = document.getElementById('loginWindow')
let next2;
let next3;
let next4;
let next5;

let methidTrue = false;

let phoneNumber;
let emailAdress;
let userName;
let userSurname;
let birthdayDate;

phoneBtn.addEventListener('click', function(){
    phoneBtn.classList.add('metBtnsActive')
    emailBtn.classList.remove('metBtnsActive')
    methidTrue = true
    frame.innerHTML = ''
    frame.innerHTML = `
                                <div class="countyFlag">🇷🇺</div>
                        <input type="text" class="phoneInp" id="phoneNum" value="+7" placeholder="+7(000)000-00-00">

    `
    IMask(
        document.getElementById('phoneNum'),
        {
            mask: '+{7}(000)000-00-00'
        }
    )
})
emailBtn.addEventListener('click', function(){
    emailBtn.classList.add('metBtnsActive')
    phoneBtn.classList.remove('metBtnsActive')
    frame.innerHTML = ''
    frame.innerHTML = `
        <input type="email" class="emailInp" placeholder="Ваш email">
    `
    methidTrue = false

})

function consol(){
    let num = document.getElementById('phoneNum')
    console.log(num.value)
}


createAcc.addEventListener('click', function (){
    mainl.innerHTML = ''
    lWind.style.height = '30%';
    lWind.style.minHeight = '300px';
    lWind.style.maxHeight = '350px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Введите номер телефона</h1>
        <div class="inputFrame" id="frameLogin">
            <div class="countyFlag">🇷🇺</div>
            <input type="text" class="phoneInp" id="phoneNumReg" value="+7" placeholder="+7(000)000-00-00">   
        </div>
        <div class="loginBtn" id="next2" ">Продолжить</div>
    `
    IMask(
        document.getElementById('phoneNumReg'),
        {
            mask: '+{7}(000)000-00-00'
        }
    )
    next2 = document.getElementById('next2')
    nextPhase2()




})

function nextPhase2(){
    next2.addEventListener('click', function (){
        phoneNumber = document.getElementById('phoneNumReg').value
        mainl.innerHTML = '';
        lWind.style.height = '30%';
        lWind.style.minHeight = '300px';
        lWind.style.maxHeight = '350px';
        mainl.innerHTML = `
        <h1 class="lmTitle">Введите электронный адрес</h1>
        <div class="inputFrame" id="frameLogin">
            <input type="email" class="emailInp" id="emailAdress" placeholder="Ваш email"> 
        </div>
        <div class="loginBtn" id="next3" ">Продолжить</div>
        `;
        next3 = document.getElementById('next3');
        nextPhase3()

    });
}
function nextPhase3(){
    next3.addEventListener('click', function (){
        emailAdress = document.getElementById('emailAdress').value

        mainl.innerHTML = '';
        lWind.style.height = '40%';
        lWind.style.minHeight = '390px';
        lWind.style.maxHeight = '450px';
        mainl.innerHTML = `
        <h1 class="lmTitle">Как вас зовут?</h1>
        <div class="inputFrame" id="frameLogin">
            <input type="text" class="emailInp" id="userName" placeholder="Ваше имя">

        </div>
        <div class="inputFrame" id="frameLogin">
            <input type="text" class="emailInp" id="userSurname" placeholder="Ваша фамилия">

        </div>
        
        <div class="loginBtn" id="next4" ">Продолжить</div>
        `;
        next4 = document.getElementById('next4');
        nextPhase4()
    });
}
function nextPhase4(){


    next4.addEventListener('click', function (){
        userName = document.getElementById('userName').value
        userSurname = document.getElementById('userSurname').value
        mainl.innerHTML = '';
        lWind.style.height = '30%';
        lWind.style.minHeight = '300px';
        lWind.style.maxHeight = '350px';
        mainl.innerHTML = `
        <h1 class="lmTitle">Ваша дата рождения</h1>
        <div class="inputFrame" id="frameLogin">
            <input type="text" class="emailInp"  id="birthDate" placeholder="01.01.2024">
        </div>
        
        <div class="loginBtn" id="next5" ">Завершить</div>
        `;
        IMask(
            document.getElementById('birthDate'),
            {
                mask: Date,
                min: new Date(1950, 0, 1),
                max: new Date(2017, 0, 1),
                lazy: false
            }
        )
        next5 = document.getElementById('next5');
        nextPhase5()
    });
}
function sendLogin(){
    birthdayDate = document.getElementById('birthDate').value
    console.log(phoneNumber, emailAdress, userName, userSurname, birthdayDate)
    registrationFinal()


}
function nextPhase5(){
    next5.addEventListener('click', function (){
        sendLogin()

        mainl.innerHTML = '';
        lWind.style.height = '50%';
        lWind.style.minHeight = '450px';
        lWind.style.maxHeight = '500px';
        mainl.innerHTML = `
        <h1 class="lmTitle">Регистрация завершена.</h1>
        <img src="./assets/img/sues.png" alt="" class="suess">
        <p class="veridyEmailTxt">Ссылка на подтверждение регистрации <br> отправлена на ваш e-mail.</p>
        
        <div class="loginBtn" id="next5" onclick="location.reload()">Вход</div>
        `;
        next5 = document.getElementById('next5');

        // nextPhase5()
    });
}
function errorPhase() {
    mainl.innerHTML = '';
    lWind.style.height = '50%';
    lWind.style.minHeight = '450px';
    lWind.style.maxHeight = '500px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Произошла ошибка.</h1>
        <img src="./assets/img/erroe.png" alt="" class="suess">
        <p class="veridyEmailTxt">Пользователь с такими данными <br> уже зарегистрирован</p>
        
        <div class="loginBtn" id="next5" ">Вход</div>
        `;
}
function errorLogin() {
    mainl.innerHTML = '';
    lWind.style.height = '50%';
    lWind.style.minHeight = '450px';
    lWind.style.maxHeight = '500px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Не верный логин или код</h1>
        <img src="./assets/img/erroe.png" alt="" class="suess">
        <p class="veridyEmailTxt">Убедитесь что вы ввели верный <br> логин или код</p>
        
        <div class="loginBtn" id="next5" ">Вход</div>
        `;
}
function errorMethod() {
    mainl.innerHTML = '';
    lWind.style.height = '50%';
    lWind.style.minHeight = '450px';
    lWind.style.maxHeight = '500px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Временно недоступно</h1>
        <img src="./assets/img/temp.png" alt="" class="suess">
        <p class="veridyEmailTxt">Вход по QR временно недоступен. <br> Мы работаем над этим</p>
        
        <div class="loginBtn" id="next5" ">Вход</div>
        `;
}
function moreMethod() {
    mainl.innerHTML = '';
    lWind.style.height = '50%';
    lWind.style.minHeight = '450px';
    lWind.style.maxHeight = '500px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Внешние способы недоступны</h1>
        <img src="./assets/img/temp.png" alt="" class="suess">
        <p class="veridyEmailTxt">Изменились правила вход в Зентас ID. <br> Внешние способы больше недоступны</p>
        <div class="loginBtn" id="next5" ">Вход</div>
        `;
}
function support() {
    mainl.innerHTML = '';
    lWind.style.height = '60%';
    lWind.style.minHeight = '500px';
    lWind.style.maxHeight = '550px';
    mainl.innerHTML = `
        <h1 class="lmTitle">Справка и поддержка</h1>
        <img src="./assets/img/support.png" alt="" class="suess">
        <p class="veridyEmailTxt">Вы можете обратиться к Wiki или <br> написать в поддержку </p>
        <div class="loginBtn" id="next5" ">Wiki</div>
        <div class="loginBtn2" id="next5" ">Поддержка</div>
        `;
}

function codeLoginEmail(){
    console.log("test")
    if (methidTrue === true) {
        loginStep1Phone()
    }
    if (methidTrue === false) {
        loginStep1Email()
    }

}
document.getElementById("supportBtn").addEventListener("click", support);
document.getElementById("qrBtn").addEventListener("click", errorMethod);
document.getElementById("moreBtn").addEventListener("click", moreMethod);
document.getElementById("login1Btn").addEventListener("click", codeLoginEmail);
