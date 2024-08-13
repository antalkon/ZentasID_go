fetch('/data/api/avatar', {
    method: 'POST',
    body: null
})
.then(response => response.json())  // Дожидаемся завершения асинхронного вызова и извлекаем JSON
.then(result => {
    console.log(result.avatar); // Выводим значение avatar в консоль
    
    if (result.avatar) {
        // Обновляем ссылку на аватар
        document.getElementById("ava").src = `/storage/user/avatars/${result.avatar}`;
    } else {
        throw new Error('Ошибка при загрузке аватара');
    }
})
.catch(error => {
    console.error('Ошибка:', error);
});
