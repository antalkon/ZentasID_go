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

        // Проверка успешности ответа
        // if (!response.ok) {
        //     throw new Error(`HTTP error! status: ${response.status}`);
        // }

        // Извлекаем данные из ответа
        const responseData = await response.json();

        // Обрабатываем сообщения в ответе
        if (responseData.message) {
            suesses_v(responseData.message)
            document.getElementById("regModal").style.display = "flex"
            console.log(`Message: ${responseData.message}`);
        }
        if (responseData.error) {
            error_v(responseData.error)
            console.error(`Error: ${responseData.error}`);
        }
        if (responseData.warn) {
            warning_v(responseData.warn)
            console.warn(`Warning: ${responseData.warn}`);
        }

        return responseData; // Возвращаем данные для дальнейшей обработки, если нужно

    } catch (error) {
        // Обработка ошибок
        console.error('Fetch error:', error);
    }
}



function Registration(){
    const name = document.getElementById("FirstName").value
    const surname = document.getElementById("LastName").value
    const phone = document.getElementById("Phone").value
    const email = document.getElementById("Email").value
    const date = document.getElementById("Date").value 
    const dateNew = formatDateToString(date)

    const url = "/auth/api/v1/registration"
    const data = {
        "userName": name,
        "userSurname": surname,
        "userBirthday": dateNew,
        "userEmail": email,
        "userPhone": phone

    }
    postRequest(url, data)

}
function formatDateToString(dateInput) {
    const date = new Date(dateInput);
    
    // Получаем день, месяц и год из даты
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Месяцы начинаются с 0
    const year = date.getFullYear();
    
    // Форматируем в строку DD.MM.YYYY
    return `${day}.${month}.${year}`;
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
  

  function toggleModal() {
    const modalContainer = document.getElementById('modal-container');
    modalContainer.__x.$data.showModal = !modalContainer.__x.$data.showModal;
}