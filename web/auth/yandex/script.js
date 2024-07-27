YaSendSuggestToken('http:/localhost:8080/yandex')
window.YaAuthSuggest.init(
    {
      client_id: "2b090bacad034d99822de2b35e3d12da",
      response_type: "token",
      redirect_uri: "http://localhost:8080/auth/yandex/token"
    },
    "http://localhost",
    {
      view: "button",
      parentId: "buttonContainerId",
      buttonSize: 'l',
      buttonView: 'icon',
      buttonTheme: 'light',
      buttonBorderRadius: "7",
      buttonIcon: 'ya',
    }
  )
  .then(({handler}) => handler())
  .then(data => console.log('Сообщение с токеном', data))
  .catch(error => console.log('Обработка ошибки', error))
  