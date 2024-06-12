const VKID = window.VKIDSDK;

VKID.Config.set({
  app: 51944733, // Идентификатор приложения.
  redirectUrl: 'http://localhost/auth/api/login/vk', // Адрес для перехода после авторизации.
  state: 'dj29fnsadjsd82...' // Произвольная строка состояния приложения.
});


const oneTap = new VKID.OneTap();
const container = document.getElementById('VkIdSdkOneTap');

// Проверка наличия кнопки в разметке.
if (container) {
  // Отрисовка кнопки в контейнере с именем приложения APP_NAME, светлой темой и на русском языке.
  oneTap.render({ container: container, scheme: VKID.Scheme.LIGHT, lang: VKID.Languages.RUS });
}