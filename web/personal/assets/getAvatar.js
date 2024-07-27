document.addEventListener('DOMContentLoaded', function() {
    fetch("/data/api/avatar/",
{
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    method: "POST",
    body: JSON.stringify({a: 1, b: 2})
})
.then(function(res) {
    if (!res.ok) {
      // Если статус ответа не ок, создаём и бросаем ошибку с сообщением о статусе
      document.getElementById("wsbAvai").src = './assets/avatar.jpeg';

      throw new Error(`HTTP error! Status: ${res.status}`);

    }
    return res.json(); // Предполагаем, что ответ в формате JSON
  })
  .then(function(res) {
    if (!res.avatar) {
      // Если параметр "avatar" отсутствует в ответе, бросаем ошибку
      document.getElementById("wsbAvai").src = './assets/avatar.jpeg';

      throw new Error("Expected parameter 'avatar' not found in response.");

    }
    avatar = res.avatar; // Извлекаем значение "avatar" и сохраняем его в переменную
    console.log(avatar); // Выводим значение "avatar" в консоль

    // Устанавливаем источник изображения
    document.getElementById("huAva").src = `/storage/user/avatars/${avatar}`;
    document.getElementById("wsbAvai").src = `/storage/user/avatars/${avatar}`;
  })
  .catch(function(err) {
    console.error(err); // Выводим ошибку в консоль

    // Устанавливаем альтернативное изображение
    document.getElementById("wsbAvai").src = './assets/avatar.jpeg';
  });
})