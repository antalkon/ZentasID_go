function registrationFinal(){
    const data = {
        "Email": emailAdress,
        "Phone": phoneNumber,
        "Name": userName,
        "Bitrhday": birthdayDate,
        "Surname": userSurname
      };

      fetch('api/reg', {
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
              errorLogin()
              // Дополнительная обработка ошибки здесь
              throw new Error('Ошибка в ответе сервера');
            }
        console.log('Успешный ответ:', jsonResponse);  

        return jsonResponse; 
      })
      .catch(error => {
        console.error('Произошла ошибка:', error);
        errorVisible(error)
        errorLogin()

        throw error; 
      });
    
}