// Создаем асинхронную функцию для выполнения POST запроса
fetchData()
async function fetchData() {
    // URL для отправки запроса
    const url = '/data/api/info';

    // Данные для отправки в теле запроса
    const postData = {
        key1: 'value1',
        key2: 'value2'
    };

    // Выполняем запрос и ждем ответа
    const response = await fetch(url, {
        method: 'POST', // Метод запроса
        headers: {
            'Content-Type': 'application/json' // Устанавливаем тип контента
        },
        body: JSON.stringify(postData) // Преобразуем данные в строку JSON
    });

    // Проверяем, успешно ли выполнен запрос
    if (response.ok) {
        // Преобразуем ответ в JSON
        const data = await response.json();

        // Создаем объект для хранения данных
        const result = {
            id: data.DisplayID,
            email: data.email,
            name: data.name,
            surname: data.surname,

            // Добавляем другие поля по мере необходимости
        };
        document.getElementById("wsbName").textContent = result.name
        document.getElementById("wsbSurname").textContent = result.surname
        document.getElementById("hulName").textContent = result.name
        document.getElementById("hulSurname").textContent = result.surname
        document.getElementById("uID").textContent = result.DisplayID;


        // Возвращаем объект для дальнейшего использования
        return result;
    } else {
        // Обрабатываем ошибку запроса
        throw new Error('Ошибка при выполнении запроса: ' + response.status);
    }
}

