async function postRequest(url, data) {
    try {
        // Отправляем запрос
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        const responseData = await response.json();

        // Обрабатываем сообщения в ответе
        if (responseData.message) {
            suesses_v(responseData.message)
            document.getElementById("regModal").style.display = "flex"
            console.log(`Message: ${responseData.message}`);
            return
        }
        if (responseData.error) {
            error_v(responseData.error)
            console.error(`Error: ${responseData.error}`);
            return
        }
        if (responseData.warn) {
            warning_v(responseData.warn)
            console.warn(`Warning: ${responseData.warn}`);
            return
        }
        

        return responseData; // Возвращаем данные для дальнейшей обработки, если нужно

    } catch (error) {
        // Обработка ошибок
        console.error('Fetch error:', error);
    }
}



function Login(){
    const email = document.getElementById("email").value

    const url = "/auth/api/v1/login/req"
    const data = {
        "email": email
    }
    postRequest(url, data)

}




function LoginFin(){
  const code = document.getElementById("code").value

  const url = "/auth/api/v1/login/final"
  const data = {
      "code": `${code}`
  }
  postRequest(url, data)

}





function suesses_v(message) {
    const alertBox = document.getElementById('alertBox_su');
    alertBox.classList.remove('hidden');
    document.getElementById('suText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }

  function error_v(message) {
    const alertBox = document.getElementById('alertBox_er');
    alertBox.classList.remove('hidden');
    document.getElementById('erText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }
  
  function warning_v(message) {
    const alertBox = document.getElementById('alertBox_wa');
    alertBox.classList.remove('hidden');
    document.getElementById('waText').textContent = message;
    
    setTimeout(() => {
      alertBox.classList.add('hidden');
    }, 3000); // 3000 milliseconds = 3 seconds
  }
  


  