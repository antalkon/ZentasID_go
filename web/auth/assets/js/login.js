let emails; 
let phone;
function loginStep1Email() {
    const email = document.getElementById("y_email").value; 
    emails = email

    console.log("++")
  
    const data = {
      email: email
    };
    console.log(data)
  
    fetch('api/login/standart/step1', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(jsonResponse => {
        if (jsonResponse.error === 'error') {
            console.error('Ошибка в ответе:', jsonResponse.error);
            // Дополнительная обработка ошибки здесь
            throw new Error('Ошибка в ответе сервера');
          }
      console.log('Успешный ответ:', jsonResponse);  
      SuessCodeEmail()
      return jsonResponse; 
    })
    .catch(error => {
      console.error('Произошла ошибка:', error);
      errorVisible(error)
      throw error; 
    });
  }
  

let loginEmail;
let loginPhone;

function clckWaitEmail(){
    loginEmail.addEventListener('click', function() {
        console.log("Clicked");
    
        // Получаем значение поля ввода email и временного кода
        const code = parseInt(document.getElementById('tempCode').value, 10); // Преобразуем значение в целое число с основанием 10
    
        const data = {
        "email": emails,
        "code": code
        };
    
        console.log(data);
    
        fetch('api/login/standart/step2', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
        })
        .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
        })
        .then(jsonResponse => {
        if (jsonResponse.error === 'error') {
            console.error('Ошибка в ответе:', jsonResponse.error);
            // Дополнительная обработка ошибки здесь
            throw new Error('Ошибка в ответе сервера');
        }
        console.log('Успешный ответ:', jsonResponse);  
        window.location.replace('https://id.zentas.ru/');
        return jsonResponse; 
        })
        .catch(error => {
        console.error('Произошла ошибка:', error);
        errorVisible(error);
        throw error; 
        });
    });
    
}
function SuessCodeEmail(){
    mainl.innerHTML = '';
    lWind.style.height = '30%';
    lWind.style.minHeight = '300px';
    lWind.style.maxHeight = '350px';
    mainl.innerHTML = `
        <h1 class="lmTitle">4-х значный код</h1>
        <div class="inputFrame" id="frameLogin">
            <input type="text" class="emailInp"  id="tempCode" placeholder="Введите код">
        </div>
        
        <div class="loginBtn" id="loginemail" ">Вход</div>`
    loginEmail = document.getElementById('loginemail')
    clckWaitEmail()
}





function loginStep1Phone() {
    let phone = document.getElementById("phoneNum").value; 
    phone = phone.replace(/[()-]/g, ''); // Удаляем символы скобок и тире

    phones = phone

    console.log("++p")
  
    const data = {
      phone: phone
    };
    console.log(data)
  
    fetch('api/login/standart/step1', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(jsonResponse => {
        if (jsonResponse.error === 'error') {
            console.error('Ошибка в ответе:', jsonResponse.error);
            // Дополнительная обработка ошибки здесь
            throw new Error('Ошибка в ответе сервера');
          }
      console.log('Успешный ответ:', jsonResponse);  
      SuessCodePhone()
      return jsonResponse; 
    })
    .catch(error => {
      console.error('Произошла ошибка:', error);
      errorVisible(error)
      throw error; 
    });
  }
  function clckWaitPhone(){
    loginPhone.addEventListener('click', function() {
        console.log("Clicked");
    
        // Получаем значение поля ввода email и временного кода
        const code = parseInt(document.getElementById('tempCode').value, 10); // Преобразуем значение в целое число с основанием 10
    
        const data = {
        "phone": phones,
        "code": code
        };
    
        console.log(data);
    
        fetch('api/login/standart/step2', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
        })
        .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
        })
        .then(jsonResponse => {
        if (jsonResponse.error === 'error') {
            console.error('Ошибка в ответе:', jsonResponse.error);
            // Дополнительная обработка ошибки здесь
            throw new Error('Ошибка в ответе сервера');
        }
        console.log('Успешный ответ:', jsonResponse);  
        window.location.replace('https://id.zentas.ru/');
        return jsonResponse; 
        })
        .catch(error => {
        console.error('Произошла ошибка:', error);
        errorVisible(error);
        throw error; 
        });
    });
    
}
  function SuessCodePhone(){
    mainl.innerHTML = '';
    lWind.style.height = '30%';
    lWind.style.minHeight = '300px';
    lWind.style.maxHeight = '350px';
    mainl.innerHTML = `
        <h1 class="lmTitle">4-х значный код</h1>
        <div class="inputFrame" id="frameLogin">
            <input type="text" class="emailInp"  id="tempCode" placeholder="Введите код">
        </div>
        
        <div class="loginBtn" id="loginphine" ">Вход</div>`
    loginPhone = document.getElementById('loginphine')
    clckWaitPhone()
}