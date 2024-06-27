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
    return res.json(); // Предполагаем, что ответ в формате JSON
  })
  .then(function(res) {
    avatar = res.avatar; // Извлекаем значение "avatar" и сохраняем его в переменную
    console.log(avatar); // Выводим значение "avatar" в консоль
    document.getElementById("huAva").src = `/storage/user/avatars/${avatar}`
    document.getElementById("wsbAvai").src = `/storage/user/avatars/${avatar}`

  })
  .catch(function(err) {
    console.error(err); // Обрабатываем ошибки
  });
})